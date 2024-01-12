// SPDX-License-Identifier: MIT

pragma solidity >=0.8.17;

interface IReputationManager {
    function getBenevolentTrainers(
        address requestorAddress,
        string memory topicName
    ) external view returns (address[] memory);

    function getMalevolentTrainers(
        address requestorAddress,
        string memory topicName
    ) external view returns (address[] memory);

    function getBenevolentValidators(
        address requestorAddress,
        string memory topicName
    ) external view returns (address[] memory);

    function getMalevolentValidators(
        address requestorAddress,
        string memory topicName
    ) external view returns (address[] memory);

    function getMalevolentChallengers(
        address requestorAddress,
        string memory topicName
    ) external view returns (address[] memory);

    function trainerIsRogue(
        address requestorAddress,
        string memory topicName,
        address trainerAddress
    ) external returns (bool);
}
