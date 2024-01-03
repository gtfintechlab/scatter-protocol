// SPDX-License-Identifier: MIT

pragma solidity >=0.8.17;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./IScatterProtocol.sol";
import "./IVoteManager.sol";
import "./IEvaluationJobToken.sol";
import "hardhat/console.sol";

contract EvaluationJobToken is ERC721URIStorage, Ownable, IEvaluationJobToken {
    using Counters for Counters.Counter;
    mapping(uint256 => string) private _tokenURIs;
    mapping(address => mapping(string => string)) private jobToAddress;
    mapping(address => mapping(string => string)) evaluationJobToDataMapping;
    string baseURI;
    address payable public protocolDeployer;
    IScatterProtocol public scatterProtocolContract;
    IVoteManager public voteManagerContract;

    Counters.Counter private _tokenIds;

    // Mappings used to keep track of evaluaation metrics
    /*
        Example evaluationScore Mapping:
        {
            requestor address: {
                topic name: {
                    validator 1 address: {
                        trainer 1 address: score 1,
                        trainer 2 address: score 2
                    },
                    validator 2 address: {...}
                    }
                }
            }
        }
    */
    mapping(address => mapping(string => mapping(address => mapping(address => uint256)))) evaluationScore;
    mapping(address => mapping(string => mapping(address => mapping(address => bool)))) evaluationScoreSet;

    /*
        Example evaluationScore Mapping:
        {
            requestor address: {
                topic name: {
                    trainer address: [1, 2, 3]
                }
            }
        }
    */
    mapping(address => mapping(string => mapping(address => uint[]))) trainerScoresForJob;

    constructor() ERC721("Scatter Protocol Evaluation Jobs", "SPEJ") {
        protocolDeployer = payable(msg.sender);
    }

    function setScatterContractAddress(
        address contractAddress
    ) public onlyOwner {
        scatterProtocolContract = IScatterProtocol(contractAddress);
    }

    function setVoteManagerContract(address contractAddress) public onlyOwner {
        voteManagerContract = IVoteManager(contractAddress);
    }

    function publishEvaluationJob(
        address requestorAddress,
        string memory topicName,
        string memory tokenURI
    ) external onlyScatterProtocolContract returns (uint256) {
        _tokenIds.increment();

        uint256 newItemId = _tokenIds.current();
        _safeMint(requestorAddress, newItemId);
        setTokenURI(newItemId, tokenURI);
        jobToAddress[requestorAddress][topicName] = tokenURI;
        return newItemId;
    }

    function publishEvaluationData(
        address requestorAddress,
        string memory topicName,
        string memory jobURI,
        string memory evaluationDataURI
    ) external onlyScatterProtocolContract {
        string memory checkOwner = jobToAddress[requestorAddress][topicName];
        require(
            keccak256(bytes(checkOwner)) == keccak256(bytes(jobURI)),
            "This evaluation job does not belong to the data publisher"
        );

        require(
            keccak256(
                bytes(evaluationJobToDataMapping[requestorAddress][topicName])
            ) == keccak256(bytes("")),
            "This job already has a mapping to it"
        );

        evaluationJobToDataMapping[requestorAddress][
            topicName
        ] = evaluationDataURI;
    }

    function setTokenURI(uint256 tokenId, string memory tokenURI) internal {
        _setTokenURI(tokenId, tokenURI);
    }

    function getTokenURI(
        uint256 tokenId
    ) external view returns (string memory) {
        return tokenURI(tokenId);
    }

    function setBaseURI(string memory baseURI_) external {
        require(
            msg.sender == protocolDeployer,
            "Only the owner can access this method"
        );

        baseURI = baseURI_;
    }

    function _baseURI() internal view override returns (string memory) {
        return baseURI;
    }

    function submitEvaluationScore(
        address requestorAddress,
        string memory topicName,
        address trainerAddress,
        address validatorAddress,
        uint256 score
    ) external onlyScatterProtocolContract {
        if (
            evaluationScoreSet[requestorAddress][topicName][validatorAddress][
                trainerAddress
            ]
        ) {
            return;
        }

        require(score >= 0 && score <= 100, "Score must be between 0 and 100");
        require(
            scatterProtocolContract.isValidatorForTrainingJob(
                requestorAddress,
                topicName,
                validatorAddress
            ),
            "You are not an assigned validator for this training job"
        );
        evaluationScore[requestorAddress][topicName][validatorAddress][
            trainerAddress
        ] = score;

        uint[] storage trainerScores = trainerScoresForJob[requestorAddress][
            topicName
        ][trainerAddress];

        trainerScores.push(score);

        evaluationScoreSet[requestorAddress][topicName][validatorAddress][
            trainerAddress
        ] = true;
        voteManagerContract.submitScoreVote(
            requestorAddress,
            topicName,
            validatorAddress,
            trainerAddress
        );
    }

    function isEvaluationScoreSet(
        address requestorAddress,
        string memory topicName,
        address validatorAddress,
        address trainerAddress
    ) external view returns (bool) {
        return
            evaluationScoreSet[requestorAddress][topicName][validatorAddress][
                trainerAddress
            ];
    }

    function getAverageScoreForTrainerForJob(
        address requestorAddress,
        string memory topicName,
        address trainerAddress
    ) external view returns (uint) {
        uint[] memory trainerScores = trainerScoresForJob[requestorAddress][
            topicName
        ][trainerAddress];
        uint summedScore = 0;
        for (uint i = 0; i < trainerScores.length; i++) {
            summedScore += trainerScores[i];
        }

        if (trainerScores.length == 0) {
            return 0;
        }

        return summedScore / trainerScores.length;
    }

    modifier onlyScatterProtocolContract() {
        require(
            msg.sender == address(scatterProtocolContract),
            "This method can only be called by the scatter protocol contract"
        );
        _;
    }
}
