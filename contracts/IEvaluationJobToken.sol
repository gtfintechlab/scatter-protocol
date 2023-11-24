// SPDX-License-Identifier: MIT

pragma solidity ^0.8.17;

interface IEvaluationJobToken {
    function publishEvaluationJob(
        address recipient,
        string memory tokenURI
    ) external;

    function publishEvaluationData(
        address dataPublisher,
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
}
