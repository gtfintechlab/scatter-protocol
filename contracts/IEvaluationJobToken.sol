// SPDX-License-Identifier: MIT

pragma solidity ^0.8.17;

interface IEvaluationJobToken {
    function publishEvaluationJob(
        address recipient,
        string memory tokenURI
    ) external;
}
