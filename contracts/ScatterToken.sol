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
    mapping(address => uint256) public addressToStake;
    mapping(address => uint256) internal addressToStakeTime;

    mapping(address => mapping(string => uint256)) requestorLockedTokenForTrainingJob;

    address public scatterProtocolAddress;

    event TokensStaked(address from, uint256 amount);
    event TokensUnstaked(address to, uint256 amount);
    bool reentrancyLock;

    constructor(
        uint256 cap
    ) ERC20("ScatterToken", "ST") ERC20Capped(cap * (10 ** decimals())) {
        owner = payable(msg.sender);
        _mint(owner, cap * (10 ** decimals()));

        requiredModelValidatorStake = 10000;
        reentrancyLock = false;
    }

    modifier noReentrant() {
        require(!reentrancyLock, "No re-entrancy");
        reentrancyLock = true;
        _;
        reentrancyLock = false;
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

    function requestorLockToken(
        address requestorAddress,
        string memory topicName,
        uint256 amount
    ) external onlyScatterProtocolContract {
        require(
            amount <= this.balanceOf(requestorAddress),
            "Cannot lock more tokens than you own"
        );
        _burn(requestorAddress, amount);
        requestorLockedTokenForTrainingJob[requestorAddress][
            topicName
        ] = amount;
    }

    function stakeToken(uint256 amount) public noReentrant {
        require(
            amount <= this.balanceOf(msg.sender),
            "Cannot stake more tokens than you own"
        );
        _burn(msg.sender, amount);
        addressToStake[msg.sender] += amount;

        // Anytime more stake is added, stake timer increases
        addressToStakeTime[msg.sender] = block.timestamp + 30 days;
        emit TokensStaked(msg.sender, amount);
    }

    function removeStake(uint256 amount) public noReentrant {
        require(
            addressToStakeTime[msg.sender] >= block.timestamp,
            "You must wait before being able to unstake Scatter Token"
        );
        require(
            amount < addressToStake[msg.sender],
            "Amount must be less than staked amount"
        );

        _mint(msg.sender, amount);
        addressToStake[msg.sender] -= amount;

        roles messageSenderRole = IScatterProtocol(scatterProtocolAddress)
            .getEnumRoleByAddress(msg.sender);
        if (
            messageSenderRole == roles.Validator &&
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

    function canBecomeValidator(address account) external view returns (bool) {
        return addressToStake[account] >= requiredModelValidatorStake;
    }

    modifier onlyScatterProtocolContract() {
        require(
            msg.sender == scatterProtocolAddress,
            "This method can only be called by the scatter protocol contract"
        );
        _;
    }

    modifier onlyOwner() {
        require(
            payable(msg.sender) == owner,
            "Only the owner can call this function"
        );
        _;
    }
}
