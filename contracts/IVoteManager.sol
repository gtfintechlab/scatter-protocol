// SPDX-License-Identifier: MIT

pragma solidity >=0.8.17;

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
        address trainerAddress,
        uint256 score
    ) external;
}
