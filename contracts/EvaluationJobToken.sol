// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./IScatterProtocol.sol";

contract EvaluationJobToken is ERC721URIStorage, Ownable {
    using Counters for Counters.Counter;
    mapping(uint256 => string) private _tokenURIs;
    mapping(string => address) private jobToAddress;
    mapping(string => string) evaluationJobToDataMapping;
    string baseURI;
    address payable public protocolDeployer;
    IScatterProtocol public scatterProtocolContract;

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

    constructor() ERC721("Scatter Protocol Evaluation Jobs", "SPEJ") {
        protocolDeployer = payable(msg.sender);
    }

    function setScatterContractAddress(
        address contractAddress
    ) public onlyOwner {
        scatterProtocolContract = IScatterProtocol(contractAddress);
    }

    function publishEvaluationJob(
        address recipient,
        string memory tokenURI
    ) external onlyScatterProtocolContract returns (uint256) {
        _tokenIds.increment();

        uint256 newItemId = _tokenIds.current();
        _safeMint(recipient, newItemId);
        setTokenURI(newItemId, tokenURI);
        jobToAddress[tokenURI] = recipient;
        return newItemId;
    }

    function publishEvaluationData(
        address dataPublisher,
        string memory jobURI,
        string memory evaluationDataURI
    ) external onlyScatterProtocolContract {
        require(
            jobToAddress[jobURI] == dataPublisher,
            "This evaluation job does not belong to the data publisher"
        );

        require(
            keccak256(bytes(evaluationJobToDataMapping[jobURI])) ==
                keccak256(bytes("")),
            "This job already has a mapping to it"
        );

        evaluationJobToDataMapping[jobURI] = evaluationDataURI;
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
        require(
            !evaluationScoreSet[requestorAddress][topicName][validatorAddress][
                trainerAddress
            ],
            "Cannot resubmit a score for a trainer that has previously been submitted"
        );

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

        evaluationScoreSet[requestorAddress][topicName][validatorAddress][
            trainerAddress
        ] = true;
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

    modifier onlyScatterProtocolContract() {
        require(
            msg.sender == address(scatterProtocolContract),
            "This method can only be called by the scatter protocol contract"
        );
        _;
    }
}
