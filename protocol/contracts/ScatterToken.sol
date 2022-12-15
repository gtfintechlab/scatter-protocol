// SPDX-License-Identifier: GPL-2.0-or-later

pragma solidity ^0.8.17;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Capped.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";

contract ScatterToken is ERC20Capped, ERC20Burnable {
    address payable public owner;
    uint256 public blockReward;
    mapping(address => uint256) internal stakes;

    constructor(uint256 cap, uint256 reward)
        ERC20("ScatterToken", "ST")
        ERC20Capped(cap * (10**decimals()))
    {
        owner = payable(msg.sender);
        _mint(owner, 70000000 * (10**decimals()));
        blockReward = reward * (10**decimals());
    }

    function addStake(uint256 amount) public {
        require(amount > 0, "Stake amount must be positive");
        _burn(msg.sender, amount);
        stakes[msg.sender] += amount;
    }

    function removeStake(uint256 amount) public {
        require(
            amount <= stakes[msg.sender],
            "Amount must be less than staked amount"
        );

        require(amount > 0, "Stake amount must be positive");

        _mint(msg.sender, amount);
        stakes[msg.sender] -= amount;
    }

    function getStake() public view returns (uint256) {
        return stakes[msg.sender];
    }

    function _mint(address account, uint256 amount)
        internal
        virtual
        override(ERC20Capped, ERC20)
    {
        require(
            ERC20.totalSupply() + amount <= cap(),
            "ERC20Capped: cap exceeded"
        );
        super._mint(account, amount);
    }

    function setBlockReward(uint256 reward) public onlyOwner {
        blockReward = reward * (10**decimals());
    }

    function _mintMinerReward() internal {
        _mint(block.coinbase, blockReward);
    }

    function _beforeTokenTransfer(
        address from,
        address to,
        uint256 value
    ) internal virtual override {
        if (
            from != address(0) &&
            to != block.coinbase &&
            block.coinbase != address(0)
        ) {
            _mintMinerReward();
        }
        super._beforeTokenTransfer(from, to, value);
    }

    modifier onlyOwner() {
        require(
            payable(msg.sender) == owner,
            "Only the owner can call this function"
        );
        _;
    }

    function destroy() public onlyOwner {
        selfdestruct(owner);
    }
}
