/*
 * @Author: pengxb master@pengxb.com
 * @Date: 2023-10-16 10:00:17
 * @LastEditors: pengxb master@pengxb.com
 * @LastEditTime: 2023-10-20 14:26:23
 * @FilePath: /test-solidity/migrations/2_deploy_contracts.js
 * @Description: 迁移文件来部署合约
 * Copyright (c) 2023 by pengxb email: master@pengxb.com, All Rights Reserved.
 */
const Dtop = artifacts.require("Dtop");

module.exports = function (deployer) {
  deployer.deploy(Dtop);
};
