// SPDX-License-Identifier: MIT

pragma solidity ^0.8.17;

interface ITrainingJobToken {
    function publishTrainingJob(
        string memory tokenURI,
        string memory topicName,
        address recipient,
        uint256 pooledReward
    ) external;
}
