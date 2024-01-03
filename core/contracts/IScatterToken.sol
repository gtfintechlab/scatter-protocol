// SPDX-License-Identifier: MIT

pragma solidity >=0.8.17;

interface IScatterToken {
    function canBecomeValidator(address account) external returns (bool);

    function canBecomeChallenger(address account) external returns (bool);

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
        string memory topicName
    ) external;

    function punishRogueValidators(
        address requestorAddress,
        string memory topicName
    ) external;

    function rewardBenevolentTrainers(
        address requestorAddress,
        string memory topicName
    ) external;

    function rewardBenevolentValidators(
        address requestorAddress,
        string memory topicName
    ) external;

    function returnTokensToTrainers(
        address requestorAddress,
        string memory topicName
    ) external;

    function rewardChallengers(
        address requestorAddress,
        string memory topicName
    ) external;

    function getLotteryPoolExternal() external view returns (uint256);
}
