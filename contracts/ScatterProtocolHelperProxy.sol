// SPDX-License-Identifier: MIT

pragma solidity ^0.8.17;

import "./Shared.sol";

contract ScatterProtocolHelperProxy {
    address payable public owner;
    address public scatterContractAddress;

    constructor() {
        owner = payable(msg.sender);
    }

    function setScatterContractAddress(
        address contractAddress
    ) public onlyOwner {
        scatterContractAddress = contractAddress;
    }

    modifier onlyOwner() {
        require(
            payable(msg.sender) == owner,
            "Only the owner can call this function"
        );
        _;
    }

    modifier onlyScatterProtocolContract() {
        require(
            msg.sender == scatterContractAddress,
            "This method can only be called by the scatter protocol contract"
        );
        _;
    }
}
