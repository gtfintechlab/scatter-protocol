// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

interface IScatterToken {
    function canBecomeValidator() external returns (bool);

    function requestorLockToken(
        address requestorAddress,
        string memory topicName,
        uint256 pooledReward
    ) external;
}
