// SPDX-License-Identifier: MIT

pragma solidity >=0.8.17;

import "./IScatterProtocol.sol";
import "./IScatterToken.sol";
import "./IEvaluationJobToken.sol";
import "./Shared.sol";

import "@openzeppelin/contracts/access/Ownable.sol";

contract VoteManager is Ownable {
    IScatterProtocol scatterProtocolContract;
    IScatterToken scatterTokenContract;
    IEvaluationJobToken evaluationTokenContract;

    uint counterQuorum = 1;

    struct ValidationProposal {
        address requestorAddress;
        string topicName;
        uint validationThreshold;
    }

    mapping(address => mapping(string => ValidationProposal)) validationProposals;
    /**
        {
            requestor address: {
                topic name: {
                    validator address: {
                        trainer address: true / false (they vote for a specific model)
                    }
                }
            }
        }
     */
    mapping(address => mapping(string => mapping(address => mapping(address => bool)))) validatorHasVoted;
    /**
        {
            requestor address: {
                topic name: {
                    trainer address: vote count
                }
            }
        }
     */
    mapping(address => mapping(string => mapping(address => uint))) voteCounter;
    mapping(address => mapping(string => mapping(address => bool))) voteResult;
    mapping(address => mapping(string => mapping(address => bool))) modelIsValidated;

    event ModelAccepted(
        address requestorAddress,
        string topicName,
        address trainerAddress
    );
    event ModelRejected(
        address requestorAddress,
        string topicName,
        address trainerAddress
    );

    function setScatterProtocolContract(
        address contractAddress
    ) external onlyOwner {
        scatterProtocolContract = IScatterProtocol(contractAddress);
    }

    function setEvaluationJobTokenContract(
        address contractAddress
    ) external onlyOwner {
        evaluationTokenContract = IEvaluationJobToken(contractAddress);
    }

    function createValidationProposal(
        address requestorAddress,
        string memory topicName,
        uint256 validationThreshold
    ) external onlyScatterProtocol {
        validationProposals[requestorAddress][topicName] = ValidationProposal(
            requestorAddress,
            topicName,
            validationThreshold
        );
    }

    function submitScoreVote(
        address requestorAddress,
        string memory topicName,
        address validatorAddress,
        address trainerAddress
    ) external onlyEvaluationTokenContract {
        require(
            validatorHasVoted[requestorAddress][topicName][validatorAddress][
                trainerAddress
            ] == false,
            "Validator cannot vote twice for a single model on a validation proposal"
        );

        validatorHasVoted[requestorAddress][topicName][validatorAddress][
            trainerAddress
        ] = true;

        ValidationProposal memory proposal = validationProposals[
            requestorAddress
        ][topicName];

        voteCounter[requestorAddress][topicName][trainerAddress] += 1;

        uint voteCount = voteCounter[requestorAddress][topicName][
            trainerAddress
        ];

        if (voteCount == counterQuorum) {
            uint averageScore = evaluationTokenContract
                .getAverageScoreForTrainerForJob(
                    requestorAddress,
                    topicName,
                    trainerAddress
                );

            modelIsValidated[requestorAddress][topicName][
                trainerAddress
            ] = true;

            if (averageScore >= proposal.validationThreshold) {
                emit ModelAccepted(requestorAddress, topicName, trainerAddress);
                voteResult[requestorAddress][topicName][trainerAddress] = true;
            } else {
                emit ModelRejected(requestorAddress, topicName, trainerAddress);
                voteResult[requestorAddress][topicName][trainerAddress] = false;
            }
        }
    }

    function getModelValidationStatus(
        address requestorAddress,
        string memory topicName,
        address trainerAddress
    ) external view returns (ValidationStatus) {
        if (!modelIsValidated[requestorAddress][topicName][trainerAddress]) {
            return ValidationStatus.ModelNotValidated;
        }

        if (voteResult[requestorAddress][topicName][trainerAddress]) {
            return ValidationStatus.ValidModel;
        } else {
            return ValidationStatus.InvalidModel;
        }
    }

    modifier onlyScatterProtocol() {
        require(
            msg.sender == address(scatterProtocolContract),
            "This method can only be called by the training job token contract"
        );
        _;
    }

    modifier onlyEvaluationTokenContract() {
        require(
            msg.sender == address(evaluationTokenContract),
            "This method can only be called by the evaluation token contract"
        );
        _;
    }
}
