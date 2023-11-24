// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;
import "./Shared.sol";

interface IScatterProtocol {
    function getEnumRoleByAddress(
        address addressToView
    ) external view returns (roles);

    function setRole(address addressToUpdate, roles newRole) external;

    function getTrainersForFederatedJob(
        address requestorAddress,
        string memory topicName
    ) external view returns (address[] memory);

    function isValidatorForTrainingJob(
        address requestorAddress,
        string memory topicName,
        address validatorAddress
    ) external view returns (bool);

    function getValidatorsForFederatedJob(
        address requestorAddress,
        string memory topicName
    ) external view returns (address[] memory);
}
