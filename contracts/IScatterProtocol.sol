// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;
import "./Shared.sol";

interface IScatterProtocol {
    function getEnumRoleByAddress(
        address addressToView
    ) external view returns (roles);

    function setRole(address addressToUpdate, roles newRole) external;
}
