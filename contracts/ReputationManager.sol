// SPDX-License-Identifier: MIT

pragma solidity ^0.8.17;

import "./IModelToken.sol";
import "./IScatterProtocol.sol";
import "./IEvaluationJobToken.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

contract ReputationManager {
    IModelToken modelTokenContract;
    IEvaluationJobToken evaluationJobTokenContract;
    IScatterProtocol scatterProtocolContract;

    function setModelTokenContract(address modelTokenContractAddress) external {
        modelTokenContract = IModelToken(modelTokenContractAddress);
    }

    function setScatterProtocolContract(
        address scatterProtocolContractAddress
    ) external {
        scatterProtocolContract = IScatterProtocol(
            scatterProtocolContractAddress
        );
    }

    function setEvaluationJobContract(
        address evaluationJobTokenAddress
    ) external {
        evaluationJobTokenContract = IEvaluationJobToken(
            evaluationJobTokenAddress
        );
    }

    function trainerIsRogue(
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
            }
        }
        return malevolentTrainers;
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
            for (uint j = 0; j < trainerList.length; j++) {
                // We do not care about rogue trainers - skip them entirely
                if (
                    this.trainerIsRogue(
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
                    !evaluationJobTokenContract.isEvaluationScoreSet(
                        requestorAddress,
                        topicName,
                        validatorList[i],
                        trainerList[j]
                    )
                ) {
                    benevolentValidators[numBenevolent] = validatorList[i];
                    numBenevolent += 1;
                }
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
                // We do not care about rogue trainers - skip them entirely
                if (
                    this.trainerIsRogue(
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
                    evaluationJobTokenContract.isEvaluationScoreSet(
                        requestorAddress,
                        topicName,
                        validatorList[i],
                        trainerList[j]
                    )
                ) {
                    malevolentValidators[numMalevolent] = validatorList[i];
                    numMalevolent += 1;
                }
            }
        }
        return malevolentValidators;
    }

    function HelloWorld() public pure {
        return;
    }
}
