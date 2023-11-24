// SPDX-License-Identifier: MIT

pragma solidity ^0.8.17;

interface IModelToken {
    function publishModel(
        string memory tokenURI,
        address recipient,
        address requestorAddress,
        string memory topicName
    ) external;

    function getModelCidForTrainer(
        address requestorAddress,
        string memory topicName,
        address trainer
    ) external view returns (string memory);
}
