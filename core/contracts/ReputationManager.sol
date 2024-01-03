// SPDX-License-Identifier: MIT

pragma solidity >=0.8.17;

import "./IModelToken.sol";
import "./IScatterProtocol.sol";
import "./IEvaluationJobToken.sol";
import "./IReputationManager.sol";
import "./IVoteManager.sol";
import "./Shared.sol";
import "hardhat/console.sol";

import "@openzeppelin/contracts/utils/Strings.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract ReputationManager is IReputationManager, Ownable {
    IModelToken modelTokenContract;
    IEvaluationJobToken evaluationJobTokenContract;
    IScatterProtocol scatterProtocolContract;
    IVoteManager voteManagerContract;

    function setModelTokenContract(address modelTokenContractAddress) external {
        modelTokenContract = IModelToken(modelTokenContractAddress);
    }

    function setScatterProtocolContract(
        address scatterProtocolContractAddress
    ) external onlyOwner {
        scatterProtocolContract = IScatterProtocol(
            scatterProtocolContractAddress
        );
    }

    function setEvaluationJobContract(
        address evaluationJobTokenAddress
    ) external onlyOwner {
        evaluationJobTokenContract = IEvaluationJobToken(
            evaluationJobTokenAddress
        );
    }

    function setVoteManagerContract(
        address voteManagerAddress
    ) external onlyOwner {
        voteManagerContract = IVoteManager(voteManagerAddress);
    }

    function trainerIsRogue(
        address requestorAddress,
        string memory topicName,
        address trainerAddress
    ) external view returns (bool) {
        bool modelIsInvalid = voteManagerContract.getModelValidationStatus(
            requestorAddress,
            topicName,
            trainerAddress
        ) == ValidationStatus.InvalidModel;

        bool modelNotSubmitted = this.trainerModelNotSubmitted(
            requestorAddress,
            topicName,
            trainerAddress
        );

        bool isChallenged = voteManagerContract.isChallengeSuccessfulTrainer(
            requestorAddress,
            topicName,
            trainerAddress
        );
        return modelIsInvalid || modelNotSubmitted || isChallenged;
    }

    function trainerModelNotSubmitted(
        address requestorAddress,
        string memory topicName,
        address trainerAddress
    ) external view returns (bool) {
        return
            Strings.equal(
                modelTokenContract.getModelCidForTrainer(
                    requestorAddress,
                    topicName,
                    trainerAddress
                ),
                ""
            );
    }

    function getBenevolentTrainers(
        address requestorAddress,
        string memory topicName
    ) external view returns (address[] memory) {
        address[] memory allTrainers = scatterProtocolContract
            .getTrainersForFederatedJob(requestorAddress, topicName);
        address[] memory benevolentTrainers = new address[](allTrainers.length);
        uint numBenevolent = 0;
        for (uint i = 0; i < allTrainers.length; i++) {
            if (
                !this.trainerIsRogue(
                    requestorAddress,
                    topicName,
                    allTrainers[i]
                )
            ) {
                benevolentTrainers[numBenevolent] = allTrainers[i];
                numBenevolent += 1;
            }
        }
        return benevolentTrainers;
    }

    function getMalevolentTrainers(
        address requestorAddress,
        string memory topicName
    ) external view returns (address[] memory) {
        address[] memory allTrainers = scatterProtocolContract
            .getTrainersForFederatedJob(requestorAddress, topicName);
        address[] memory malevolentTrainers = new address[](allTrainers.length);
        uint numMalevolent = 0;
        for (uint i = 0; i < allTrainers.length; i++) {
            if (
                this.trainerIsRogue(requestorAddress, topicName, allTrainers[i])
            ) {
                malevolentTrainers[numMalevolent] = allTrainers[i];
                numMalevolent += 1;
            }
        }
        return malevolentTrainers;
    }

    function validatorIsRogue(
        address requestorAddress,
        string memory topicName,
        address validatorAddress,
        address trainerAddress
    ) external view returns (bool) {
        bool evalScoreSubmit = evaluationJobTokenContract.isEvaluationScoreSet(
            requestorAddress,
            topicName,
            validatorAddress,
            trainerAddress
        );

        bool isChallenged = voteManagerContract.isChallengeSuccessfulValidator(
            requestorAddress,
            topicName,
            validatorAddress
        );

        return !evalScoreSubmit || isChallenged;
    }

    function getBenevolentValidators(
        address requestorAddress,
        string memory topicName
    ) external view returns (address[] memory) {
        address[] memory validatorList = scatterProtocolContract
            .getValidatorsForFederatedJob(requestorAddress, topicName);
        address[] memory trainerList = scatterProtocolContract
            .getTrainersForFederatedJob(requestorAddress, topicName);
        address[] memory benevolentValidators = new address[](
            validatorList.length
        );
        uint numBenevolent = 0;

        for (uint i = 0; i < validatorList.length; i++) {
            bool allTrainersEvaluated = true;
            for (uint j = 0; j < trainerList.length; j++) {
                // We do not care about trainers who did not submit a model - skip them entirely
                if (
                    this.trainerModelNotSubmitted(
                        requestorAddress,
                        topicName,
                        trainerList[j]
                    )
                ) {
                    continue;
                }

                // For a non-rogue trainer, check if their evaluation scores have been set
                // Mark them as rogue if not
                allTrainersEvaluated =
                    allTrainersEvaluated &&
                    !this.validatorIsRogue(
                        requestorAddress,
                        topicName,
                        validatorList[i],
                        trainerList[j]
                    );
                if (!allTrainersEvaluated) {
                    break;
                }
            }
            if (allTrainersEvaluated) {
                benevolentValidators[numBenevolent] = validatorList[i];
                numBenevolent += 1;
            }
        }

        return benevolentValidators;
    }

    function getMalevolentValidators(
        address requestorAddress,
        string memory topicName
    ) external view returns (address[] memory) {
        address[] memory validatorList = scatterProtocolContract
            .getValidatorsForFederatedJob(requestorAddress, topicName);
        address[] memory trainerList = scatterProtocolContract
            .getTrainersForFederatedJob(requestorAddress, topicName);
        address[] memory malevolentValidators = new address[](
            validatorList.length
        );
        uint numMalevolent = 0;

        for (uint i = 0; i < validatorList.length; i++) {
            for (uint j = 0; j < trainerList.length; j++) {
                // We do not care about trainers who did not submit a model - skip them entirely
                if (
                    this.trainerModelNotSubmitted(
                        requestorAddress,
                        topicName,
                        trainerList[j]
                    )
                ) {
                    continue;
                }

                // For a non-rogue trainer, check if their evaluation scores have been set
                // Mark them as rogue if not
                if (
                    this.validatorIsRogue(
                        requestorAddress,
                        topicName,
                        validatorList[i],
                        trainerList[j]
                    )
                ) {
                    malevolentValidators[numMalevolent] = validatorList[i];
                    numMalevolent += 1;
                    break;
                }
            }
        }
        return malevolentValidators;
    }

    modifier onlyScatterProtocolContract() {
        require(
            msg.sender == address(scatterProtocolContract),
            "This method can only be called by the scatter protocol contract"
        );
        _;
    }
}
