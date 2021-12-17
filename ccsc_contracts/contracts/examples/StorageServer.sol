// SPDX-License-Identifier: MIT
pragma solidity ^0.7.0;
pragma experimental ABIEncoderV2;

import "hardhat/console.sol";
import "../ccsc/Server.sol";
import "./Storage.sol";

contract StorageServer is Server {
    Storage private storageContract;

    constructor(address _storageAddress, address _relayAddress)
        Server(_relayAddress)
    {
        storageContract = Storage(_storageAddress);
    }

    function store(
        uint256 num,
        address origin,
        address sender
    ) public {
        storageContract.store(num);
    }

    function retrieve(address origin, address sender)
        public
        view
        returns (uint256)
    {
        return storageContract.retrieve();
    }
}
