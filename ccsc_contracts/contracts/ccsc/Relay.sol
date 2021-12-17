// SPDX-License-Identifier: MIT
pragma solidity ^0.7.0;
pragma experimental ABIEncoderV2;

interface Relay {
    function verifyTransaction(
        uint256 feeInWei,
        bytes calldata rlpHeader,
        uint8 noOfConfirmations,
        bytes calldata rlpEncodedTx,
        bytes calldata path,
        bytes calldata rlpEncodedNodes
    ) external payable returns (uint8);

    function verifyReceipt(
        uint256 feeInWei,
        bytes calldata rlpHeader,
        uint8 noOfConfirmations,
        bytes calldata rlpEncodedReceipt,
        bytes calldata path,
        bytes calldata rlpEncodedNodes
    ) external payable returns (uint8);
}
