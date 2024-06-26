// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./contracts/token/ERC20/ERC20.sol";

contract MockERC20 is ERC20 {
    constructor(string memory name_, string memory symbol_)
    ERC20(name_, symbol_)
    {}

    function mint(address account, uint256 amount) public {
        _mint(account, amount);
    }
}