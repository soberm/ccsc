// SPDX-License-Identifier: MIT
pragma solidity ^0.7.0;
pragma experimental ABIEncoderV2;

import "../ccsc/Proxy.sol";
import "../ccsc/Protocol.sol";

contract StorageProxy is Proxy {
    constructor(address _serverAddress, address _relayAddress)
        Proxy(_serverAddress, _relayAddress)
    {}

    function store(uint256 num, Protocol.CallOpts memory callOpts) external {
        initiateCrossChainCall(
            abi.encodeWithSignature(
                "store(uint256,address,address)",
                num,
                tx.origin,
                msg.sender
            ),
            callOpts
        );
    }

    function retrieve(Protocol.CallOpts memory callOpts) external {
        initiateCrossChainCall(
            abi.encodeWithSignature(
                "retrieve(address,address)",
                tx.origin,
                msg.sender
            ),
            callOpts
        );
    }
}
