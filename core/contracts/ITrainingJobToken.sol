// SPDX-License-Identifier: MIT

pragma solidity >=0.8.17;

interface ITrainingJobToken {
    function publishTrainingJob(
        string memory tokenURI,
        address recipient
    ) external returns (uint256);
}
