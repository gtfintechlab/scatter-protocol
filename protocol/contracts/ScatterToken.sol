// SPDX-License-Identifier: GPL-2.0-or-later

pragma solidity ^0.8.17;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract ScatterToken is ERC20 {
    constructor(uint256 initialSupply) ERC20("ScatterToken", "ST") {
        _mint(msg.sender, initialSupply);
    }
}
