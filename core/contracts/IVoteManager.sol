// SPDX-License-Identifier: MIT

pragma solidity >=0.8.17;

import "./Shared.sol";

interface IVoteManager {
    function createValidationProposal(
        address requestorAddress,
        string memory topicName,
        uint256 validationThreshold
    ) external;

    function submitScoreVote(
        address requestorAddress,
        string memory topicName,
        address validatorAddress,
        address trainerAddress
    ) external;

    function getModelValidationStatus(
        address requestorAddress,
        string memory topicName,
        address trainerAddress
    ) external view returns (ValidationStatus);

    function submitValidatorChallenge(
        address requestorAddress,
        string memory topicName,
        address validatorAddress,
        bool isMalicious,
        address challengerAddress
    ) external;

    function submitTrainerChallenge(
        address requestorAddress,
        string memory topicName,
        address trainerAddress,
        bool isMalicious,
        address challengerAddress
    ) external;

    function submitChallengerChallenge(
        address requestorAddress,
        string memory topicName,
        address challengedChallengerAddress,
        bool isMalicious,
        address challengerAddress
    ) external;

    function isChallengeSuccessfulValidator(
        address requestorAddress,
        string memory topicName,
        address validatorAddress
    ) external view returns (bool);

    function getChallengedChallengers(
        address requestorAddress,
        string memory topicName
    ) external view returns (address[] memory);

    function isChallengeSuccessfulTrainer(
        address requestorAddress,
        string memory topicName,
        address validatorAddress
    ) external view returns (bool);

    function isChallengeSuccessfulChallenger(
        address requestorAddress,
        string memory topicName,
        address validatorAddress
    ) external view returns (bool);

    function hasChallengedNode(
        address requestorAddress,
        string memory topicName,
        address nodeToChallenge,
        address challengerAddress
    ) external view returns (bool);

    function refrainFromValidation(
        address requestorAddress,
        string memory topicName,
        address validatorAddress
    ) external;

    function isMaliciousValidationJob(
        address requestorAddress,
        string memory topicName
    ) external view returns (bool);

    function isRefrainingFromValidation(
        address requestorAddress,
        string memory topicName,
        address validatorAddress
    ) external view returns (bool);
}
