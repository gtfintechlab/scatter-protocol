// SPDX-License-Identifier: MIT

pragma solidity ^0.8.17;

interface IModelToken {
    function publishModel(string memory tokenURI, address recipient) external;
}
