// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

interface IScatterToken {
    function canBecomeValidator(address account) external returns (bool);

    function requestorLockToken(
        address requestorAddress,
        string memory topicName,
        uint256 pooledReward
    ) external;

    function trainerLockToken(
        address trainerAddress,
        address requestorAddress,
        string memory topicName,
        uint256 amount
    ) external;

    function donateToLottery(
        address requestorAddress,
        string memory topicName
    ) external;

    function punishRogueTrainers(
        address requestorAddress,
        string memory topicName,
        address[] memory trainers
    ) external;

    function punishRogueValidators(address[] memory validators) external;

    function rewardBenevolentTrainers(
        address requestorAddress,
        string memory topicName,
        address[] memory trainers
    ) external;

    function rewardBenevolentValidators(
        address requestorAddress,
        string memory topicName,
        address[] memory validators
    ) external;
}
