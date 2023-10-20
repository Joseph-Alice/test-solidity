// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol";

contract Dtop is ERC20, ERC20Permit {
    // 定义一个名为"LogConstructorCalled"的事件，该事件将记录合约的部署者地址
    event LogConstructorCalled(address indexed deployer);

    constructor() ERC20("Dtop", "HELLO") ERC20Permit("Dtop") {
        // 触发事件并发出日志
        emit LogConstructorCalled(msg.sender);
        _mint(msg.sender, 100000000 * 10 ** uint256(decimals()));
    }
}
