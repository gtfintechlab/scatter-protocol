// SPDX-License-Identifier: MIT

pragma solidity >=0.8.17;

interface IEvaluationJobToken {
    function publishEvaluationJob(
        address recipient,
        string memory topicName,
        string memory tokenURI
    ) external returns (uint256);

    function publishEvaluationData(
        address dataPublisher,
        string memory topicName,
        string memory jobURI,
        string memory evaluationDataURI
    ) external;

    function submitEvaluationScore(
        address requestorAddress,
        string memory topicName,
        address trainerAddress,
        address validatorAddress,
        uint256 score
    ) external;

    function isEvaluationScoreSet(
        address requestorAddress,
        string memory topicName,
        address validatorAddress,
        address trainerAddress
    ) external view returns (bool);

    function getAverageScoreForTrainerForJob(
        address requestorAddress,
        string memory topicName,
        address trainerAddress
    ) external view returns (uint);
}
