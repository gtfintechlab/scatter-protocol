// SPDX-License-Identifier: MIT

pragma solidity ^0.8.17;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Capped.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";

import "./Shared.sol";
import "./IScatterProtocol.sol";

contract ScatterToken is ERC20Capped, ERC20Burnable {
    uint256 public requiredModelValidatorStake;
    address payable public owner;
    mapping(address => uint256) internal addressToStake;
    mapping(address => uint256) internal addressToStakeTime;

    address public scatterProtocolAddress;

    constructor(
        uint256 cap
    ) ERC20("ScatterToken", "ST") ERC20Capped(cap * (10 ** decimals())) {
        owner = payable(msg.sender);
        _mint(owner, (cap / 100) * 70 * (10 ** decimals()));

        requiredModelValidatorStake = 500000;
    }

    function setScatterProtocolAddress(address newAddress) public onlyOwner {
        scatterProtocolAddress = newAddress;
    }

    function _mint(
        address account,
        uint256 amount
    ) internal virtual override(ERC20Capped, ERC20) {
        require(
            ERC20.totalSupply() + amount <= cap(),
            "ERC20Capped: cap exceeded"
        );
        super._mint(account, amount);
    }

    function addStake(uint256 amount) public {
        require(amount > 0, "Stake amount must be positive");
        _burn(msg.sender, amount);
        addressToStake[msg.sender] += amount;
        // Anytime more stake is added, stake timer increases
        addressToStakeTime[msg.sender] = block.timestamp + 30 days;
    }

    function removeStake(uint256 amount) public {
        require(
            addressToStakeTime[msg.sender] >= block.timestamp,
            "You must wait 30 days before being able to unstake Scatter Token"
        );
        require(
            amount <= addressToStake[msg.sender],
            "Amount must be less than staked amount"
        );
        require(amount > 0, "Stake amount must be positive");

        _mint(msg.sender, amount);
        addressToStake[msg.sender] -= amount;
        roles messageSenderRole = IScatterProtocol(scatterProtocolAddress)
            .getEnumRoleByAddress(msg.sender);
        if (
            messageSenderRole == roles.ModelValidator &&
            addressToStake[msg.sender] < requiredModelValidatorStake
        ) {
            IScatterProtocol(scatterProtocolAddress).setRole(
                msg.sender,
                roles.NoRole
            );
        }
    }

    function getOwnStake() public view returns (uint256) {
        return addressToStake[msg.sender];
    }

    function getAccountStake(address account) public view returns (uint256) {
        return addressToStake[account];
    }

    modifier onlyOwner() {
        require(
            payable(msg.sender) == owner,
            "Only the owner can call this function"
        );
        _;
    }
}
