// SPDX-License-Identifier: MIT
pragma solidity ^0.7.0;
pragma experimental ABIEncoderV2;

contract RelayMock {
    uint8 private verifyTransactionResult = 0;
    uint8 private verifyReceiptResult = 0;

    function setVerifyTransactionResult(uint8 _value) public {
        verifyTransactionResult = _value;
    }

    function setVerifyReceiptResult(uint8 _value) public {
        verifyReceiptResult = _value;
    }

    /* solhint-disable no-unused-vars */
    function verifyTransaction(
        uint256 feeInWei,
        bytes calldata rlpHeader,
        uint8 noOfConfirmations,
        bytes calldata rlpEncodedTx,
        bytes calldata path,
        bytes calldata rlpEncodedNodes
    ) external payable returns (uint8) {
        return verifyTransactionResult;
    }

    function verifyReceipt(
        uint256 feeInWei,
        bytes calldata rlpHeader,
        uint8 noOfConfirmations,
        bytes calldata rlpEncodedReceipt,
        bytes calldata path,
        bytes calldata rlpEncodedNodes
    ) external payable returns (uint8) {
        return verifyReceiptResult;
    }
}
