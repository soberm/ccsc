// SPDX-License-Identifier: MIT
pragma solidity ^0.7.0;
pragma experimental ABIEncoderV2;

import "./StorageProxy.sol";
import "../ccsc/Protocol.sol";

contract Client {
    StorageProxy private storageProxy;

    bool private stored;
    uint256 private retrievedNumber;

    event Stored(uint256 indexed id);

    constructor(address _storageAddress) {
        storageProxy = StorageProxy(_storageAddress);
    }

    function store(uint256 num) public {
        Protocol.CallOpts memory callOpts =
            Protocol.CallOpts(
                250000,
                1000,
                address(this),
                "storeCallback",
                250000
            );
        storageProxy.store(num, callOpts);
    }

    function storeCallback(
        uint256 id,
        bytes memory result,
        bool success
    ) external {
        stored = success;
        emit Stored(id);
    }

    function retrieve() public {
        Protocol.CallOpts memory callOpts =
            Protocol.CallOpts(
                250000,
                1000,
                address(this),
                "retrieveCallback",
                250000
            );
        storageProxy.retrieve(callOpts);
    }

    function retrieveCallback(
        uint256 id,
        bytes memory result,
        bool success
    ) external {
        retrievedNumber = abi.decode(result, (uint256));
    }

    function isStored() public view returns (bool) {
        return stored;
    }

    function getRetrievedNumber() public view returns (uint256) {
        return retrievedNumber;
    }
}
