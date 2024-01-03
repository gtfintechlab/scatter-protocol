// SPDX-License-Identifier: MIT

pragma solidity >=0.8.17;

import "./IScatterProtocol.sol";
import "./IScatterToken.sol";
import "./IEvaluationJobToken.sol";
import "./IVoteManager.sol";
import "./Shared.sol";

import "@openzeppelin/contracts/access/Ownable.sol";
import "hardhat/console.sol";

contract VoteManager is Ownable, IVoteManager {
    IScatterProtocol scatterProtocolContract;
    IScatterToken scatterTokenContract;
    IEvaluationJobToken evaluationTokenContract;

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

    // A validator can refrain from joining a trianing job if the evaluation job is deemed to be malicious
    // If this happens, they put a portion of their stake on the line
    // If a quorum of the validators agree, then the validator will still be rewarded and the requester will lose its reward
    // If they disagree, then the validator loses a portion of its stake
    // requestor address --> topicName --> validator address --> true/false
    mapping(address => mapping(string => mapping(address => bool))) maliciousEvaluations;

    // A place for challengers to challenge the maliciousness of specific nodes
    mapping(address => mapping(string => mapping(address => mapping(address => bool)))) validatorChallenge;
    mapping(address => mapping(string => mapping(address => uint256))) validatorChallengeMaliceCount;
    mapping(address => mapping(string => mapping(address => uint256))) validatorChallengeTotalCount;

    mapping(address => mapping(string => mapping(address => mapping(address => bool)))) trainerChallenge;
    mapping(address => mapping(string => mapping(address => uint256))) trainerChallengeMaliceCount;
    mapping(address => mapping(string => mapping(address => uint256))) trainerChallengeTotalCount;

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

    function refrainFromValidation(
        address requestorAddress,
        string memory topicName,
        address validatorAddress
    ) external onlyScatterProtocol {
        require(
            validatorHasVoted[requestorAddress][topicName][validatorAddress][
                address(0x0)
            ] == false,
            "Validator cannot mark a job as malicious after having submitted a score for it"
        );

        maliciousEvaluations[requestorAddress][topicName][
            validatorAddress
        ] = true;

        address[] memory trainers = scatterProtocolContract
            .getTrainersForFederatedJob(requestorAddress, topicName);
        for (uint256 i = 0; i < trainers.length; i++) {
            address trainerAddress = trainers[i];
            validatorHasVoted[requestorAddress][topicName][validatorAddress][
                trainerAddress
            ] = true;
            voteCounter[requestorAddress][topicName][trainerAddress] += 1;
        }
    }

    function isMaliciousValidationJob(
        address requestorAddress,
        string memory topicName
    ) internal view returns (bool) {
        address[] memory validators = scatterProtocolContract
            .getValidatorsForFederatedJob(requestorAddress, topicName);

        uint maliceCounter = 0;
        for (uint256 i = 0; i < validators.length; i++) {
            address validatorAddress = validators[i];

            if (
                maliciousEvaluations[requestorAddress][topicName][
                    validatorAddress
                ]
            ) {
                maliceCounter += 1;
            }
        }

        // If 2/3 of validators agree that it is a malicious job, then return true
        return maliceCounter >= (2 * validators.length) / 3;
    }

    function hasChallengedNode(
        address requestorAddress,
        string memory topicName,
        address nodeToChallenge,
        address challengerAddress
    ) external view returns (bool) {
        return
            validatorChallenge[requestorAddress][topicName][nodeToChallenge][
                challengerAddress
            ] ||
            trainerChallenge[requestorAddress][topicName][nodeToChallenge][
                challengerAddress
            ];
    }

    function submitValidatorChallenge(
        address requestorAddress,
        string memory topicName,
        address validatorAddress,
        bool isMalicious,
        address challengerAddress
    ) external onlyScatterProtocol {
        if (
            validatorChallenge[requestorAddress][topicName][validatorAddress][
                challengerAddress
            ]
        ) {
            return;
        }
        validatorChallenge[requestorAddress][topicName][validatorAddress][
            challengerAddress
        ] = true;

        if (isMalicious) {
            validatorChallengeMaliceCount[requestorAddress][topicName][
                validatorAddress
            ] += 1;
        }
        validatorChallengeTotalCount[requestorAddress][topicName][
            validatorAddress
        ] += 1;
    }

    function submitTrainerChallenge(
        address requestorAddress,
        string memory topicName,
        address trainerAddress,
        bool isMalicious,
        address challengerAddress
    ) external onlyScatterProtocol {
        if (
            trainerChallenge[requestorAddress][topicName][trainerAddress][
                challengerAddress
            ]
        ) {
            return;
        }

        trainerChallenge[requestorAddress][topicName][trainerAddress][
            challengerAddress
        ] = true;

        if (isMalicious) {
            trainerChallengeMaliceCount[requestorAddress][topicName][
                trainerAddress
            ] += 1;
        }
        trainerChallengeTotalCount[requestorAddress][topicName][
            trainerAddress
        ] += 1;
    }

    function isChallengeSuccessfulValidator(
        address requestorAddress,
        string memory topicName,
        address validatorAddress
    ) external view returns (bool) {
        uint256 malicious = validatorChallengeMaliceCount[requestorAddress][
            topicName
        ][validatorAddress];

        uint256 total = validatorChallengeTotalCount[requestorAddress][
            topicName
        ][validatorAddress];

        return malicious != 0 && malicious >= (total * 2) / 3;
    }

    function isChallengeSuccessfulTrainer(
        address requestorAddress,
        string memory topicName,
        address trainerAddress
    ) external view returns (bool) {
        uint256 malicious = trainerChallengeMaliceCount[requestorAddress][
            topicName
        ][trainerAddress];

        uint256 total = trainerChallengeTotalCount[requestorAddress][topicName][
            trainerAddress
        ];
        return malicious != 0 && malicious >= (total * 2) / 3;
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
        require(
            maliciousEvaluations[requestorAddress][topicName][
                validatorAddress
            ] == false,
            "Cannot evaluate a validation proposal after marking it false"
        );

        validatorHasVoted[requestorAddress][topicName][validatorAddress][
            address(0x0)
        ] = true;
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
        // We will go ahead and use 100% of validators
        uint counterQuorum = (
            scatterProtocolContract.getValidatorsForFederatedJob(
                requestorAddress,
                topicName
            )
        ).length;

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

            // If it is a malicious validation job, then we can just accept the model
            if (
                averageScore >= proposal.validationThreshold ||
                isMaliciousValidationJob(requestorAddress, topicName)
            ) {
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
