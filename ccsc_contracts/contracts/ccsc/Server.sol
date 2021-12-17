// SPDX-License-Identifier: MIT
pragma solidity ^0.7.0;
pragma experimental ABIEncoderV2;

import "../libs/RLPReader.sol";
import "../libs/BytesLib.sol";
import "./Relay.sol";
import "./Proxy.sol";
import "./Protocol.sol";

contract Server {
    using RLPReader for RLPReader.RLPItem;
    using RLPReader for bytes;

    struct CrossChainCall {
        uint256 id;
        address sender;
        address origin;
        bool status;
        uint256 gasLimit;
        bytes callData;
    }

    event CrossChainCallExecuted(uint256 indexed id, bool success, bytes data);

    Relay private relay;
    Proxy private proxy;

    mapping(uint256 => bool) private executedCalls;

    uint8 public constant CONFIRMATIONS = 6;
    uint256 public constant RELAY_FEE = 1 wei;

    constructor(address _relayAddress) {
        relay = Relay(_relayAddress);
    }

    function setProxy(address proxyAddress) public {
        proxy = Proxy(proxyAddress);
    }

    function executeCrossChainCall(Protocol.Proof calldata proof)
        public
        payable
    {
        CrossChainCall memory crossChainCall =
            decodeTransactionReceipt(
                proof.rlpEncodedTx,
                proof.rlpEncodedReceipt
            );

        require(crossChainCall.sender == address(proxy), "invalid sender");
        require(crossChainCall.status == true, "invalid status");

        require(executedCalls[crossChainCall.id] == false, "already executed");
        executedCalls[crossChainCall.id] = true;

        uint8 transactionVerified =
            relay.verifyTransaction{value: RELAY_FEE}(
                RELAY_FEE,
                proof.rlpHeader,
                CONFIRMATIONS,
                proof.rlpEncodedTx,
                proof.path,
                proof.rlpEncodedTxNodes
            );
        require(transactionVerified == 0, "transaction not verified");

        uint8 receiptVerified =
            relay.verifyReceipt{value: RELAY_FEE}(
                RELAY_FEE,
                proof.rlpHeader,
                CONFIRMATIONS,
                proof.rlpEncodedReceipt,
                proof.path,
                proof.rlpEncodedReceiptNodes
            );
        require(receiptVerified == 0, "transaction receipt not verified");

        require(gasleft() >= crossChainCall.gasLimit, "out of gas");
        (bool success, bytes memory data) =
            address(this).call{gas: crossChainCall.gasLimit}(
                crossChainCall.callData
            );

        emit CrossChainCallExecuted(crossChainCall.id, success, data);
    }

    function decodeTransactionReceipt(
        bytes memory rlpTransaction,
        bytes memory rlpReceipt
    ) public pure returns (CrossChainCall memory) {
        RLPReader.RLPItem memory transactionEncoded =
            rlpTransaction.toRlpItem();

        byte txType;

        RLPReader.RLPItem[] memory transaction;
        if (!transactionEncoded.isList()) {
            bytes memory transactionStripped =
                BytesLib.slice(
                    transactionEncoded.toBytes(),
                    1,
                    transactionEncoded.toBytes().length - 1
                );
            transaction = transactionStripped.toRlpItem().toList();

            txType = transactionEncoded.toBytes()[0];
        } else {
            transaction = transactionEncoded.toList();
        }

        RLPReader.RLPItem memory receiptEncoded = rlpReceipt.toRlpItem();

        RLPReader.RLPItem[] memory receipt;
        if (!receiptEncoded.isList()) {
            bytes memory receiptStripped =
                BytesLib.slice(
                    receiptEncoded.toBytes(),
                    1,
                    receiptEncoded.toBytes().length - 1
                );
            receipt = receiptStripped.toRlpItem().toList();
        } else {
            receipt = receiptEncoded.toList();
        }

        RLPReader.RLPItem[] memory logs = receipt[3].toList();

        RLPReader.RLPItem[] memory log;
        RLPReader.RLPItem[] memory eventTopics;
        bool eventFound = false;
        for (uint256 i = 0; i < logs.length; i++) {
            log = logs[i].toList();
            eventTopics = log[1].toList();
            if (
                eventTopics[0].toUint() ==
                0x27c4a4683a3d10ba190b00c77780aeb028bf0e761d7e9d47a469004983ade7e6
            ) {
                eventFound = true;
                break;
            }
        }

        require(eventFound, "event not found");

        address origin;
        uint256 gasLimit;
        bytes memory callData;

        (origin, gasLimit, callData) = abi.decode(
            log[2].toBytes(),
            (address, uint256, bytes)
        );

        address to;
        if (txType == 0x00) {
            to = transaction[3].toAddress();
        } else if (txType == 0x01) {
            to = transaction[4].toAddress();
        } else if (txType == 0x02) {
            to = transaction[5].toAddress();
        } else {
            revert("unsupported tx type");
        }

        return
            CrossChainCall(
                eventTopics[1].toUint(),
                to,
                origin,
                receipt[0].toBoolean(),
                gasLimit,
                callData
            );
    }
}
