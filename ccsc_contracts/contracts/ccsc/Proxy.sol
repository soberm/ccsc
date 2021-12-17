// SPDX-License-Identifier: MIT
pragma solidity ^0.7.0;
pragma experimental ABIEncoderV2;

import "../libs//RLPReader.sol";
import "./Relay.sol";
import "./Server.sol";
import "./Protocol.sol";

contract Proxy {
    using RLPReader for RLPReader.RLPItem;
    using RLPReader for bytes;
    using Protocol for Protocol.CrossChainRequest;

    event Deposited(address indexed sender, uint256 amount);

    event Withdrawn(address indexed sender);

    event CrossChainCallInitiated(
        uint256 callID,
        address origin,
        address callBackTo,
        bytes callData,
        uint256 gas,
        string callBackName,
        uint256 callBackGas,
        uint256 timeout
    );

    event CrossChainCallPrepared(
        uint256 indexed call,
        address origin,
        uint256 gas,
        bytes callData
    );

    event CrossChainCallAcknowledged(uint256 indexed call);

    uint256 private callID;
    mapping(uint256 => bytes32) private calls;
    mapping(address => uint256) private numberOfCalls;
    mapping(address => uint256) private deposits;

    Server private server;
    Relay private relay;

    uint256 public constant MIN_CALL_GAS = 1000000;
    string public constant CALLBACK_PARAMS = "(uint256,bytes,bool)";
    uint8 public constant CONFIRMATIONS = 6;
    uint256 public constant RELAY_FEE = 0;
    uint256 public constant MIN_DEPOSIT_PER_CALL = 1 wei;

    constructor(address serverAddress, address relayAddress) {
        server = Server(serverAddress);
        relay = Relay(relayAddress);
    }

    function deposit() external payable {
        deposits[msg.sender] += msg.value;
        emit Deposited(msg.sender, msg.value);
    }

    function withdraw() external {
        require(numberOfCalls[msg.sender] == 0, "calls pending");
        msg.sender.transfer(deposits[msg.sender]);
        delete deposits[msg.sender];

        emit Withdrawn(msg.sender);
    }

    function initiateCrossChainCall(
        bytes memory callData,
        Protocol.CallOpts memory callOpts
    ) public {
        require(checkDeposit(tx.origin), "deposit too low");

        callID += 1;
        numberOfCalls[tx.origin] += 1;

        calls[callID] = Protocol
            .CrossChainRequest(
            callID,
            tx
                .origin,
            callOpts
                .callBackTo,
            callData,
            callOpts
                .gasLimit,
            callOpts
                .callBackName,
            callOpts
                .callBackGasLimit,
            block.number + callOpts.timeout
        )
            .hash();

        emit CrossChainCallInitiated(
            callID,
            tx.origin,
            callOpts.callBackTo,
            callData,
            callOpts.gasLimit,
            callOpts.callBackName,
            callOpts.callBackGasLimit,
            block.number + callOpts.timeout
        );
    }

    function checkDeposit(address _sender) private view returns (bool) {
        uint256 maxCalls = deposits[_sender] / MIN_DEPOSIT_PER_CALL;
        return numberOfCalls[_sender] < maxCalls;
    }

    function prepareCrossChainCall(Protocol.CrossChainRequest memory call)
        external
    {
        require(calls[call.id] == call.hash(), "call does not exist");

        emit CrossChainCallPrepared(
            call.id,
            call.origin,
            call.gas,
            call.callData
        );
    }

    function acknowledgeCrossChainCall(
        Protocol.Proof calldata proof,
        Protocol.CrossChainRequest memory call
    ) external payable {
        Protocol.CrossChainResponse memory response =
            decodeTransactionReceipt(
                proof.rlpEncodedTx,
                proof.rlpEncodedReceipt
            );

        require(call.id == response.id, "invalid call id");
        require(calls[call.id] == call.hash(), "call does not exist");
        require(response.server == address(server), "invalid sender");
        require(response.status == true, "invalid status");

        if (block.number < call.timeout) {
            require(msg.sender == call.origin, "invalid sender");
        } else {
            deposits[call.origin] -= MIN_DEPOSIT_PER_CALL;
            payable(msg.sender).transfer(MIN_DEPOSIT_PER_CALL);
        }

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

        require(gasleft() >= call.callBackGas, "out of gas");
        (bool success, ) =
            call.callBackTo.call{gas: call.callBackGas}(
                abi.encodeWithSignature(
                    string(abi.encodePacked(call.callBack, CALLBACK_PARAMS)),
                    response.id,
                    response.data,
                    response.success
                )
            );

        numberOfCalls[msg.sender] -= 1;
        delete calls[response.id];
        emit CrossChainCallAcknowledged(response.id);
    }

    function decodeTransactionReceipt(
        bytes memory rlpTransaction,
        bytes memory rlpReceipt
    ) public pure returns (Protocol.CrossChainResponse memory) {
        RLPReader.RLPItem[] memory transaction =
            rlpTransaction.toRlpItem().toList();
        RLPReader.RLPItem[] memory receipt = rlpReceipt.toRlpItem().toList();
        RLPReader.RLPItem[] memory logs = receipt[3].toList();

        RLPReader.RLPItem[] memory log;
        RLPReader.RLPItem[] memory eventTopics;
        bool eventFound = false;
        for (uint256 i = 0; i < logs.length; i++) {
            log = logs[i].toList();
            eventTopics = log[1].toList();
            if (
                eventTopics[0].toUint() ==
                0xd502609f911e36dbff128057b05aee250aad0efda2f06a626f99b6740d193761
            ) {
                eventFound = true;
                break;
            }
        }

        require(eventFound, "event not found");

        bool success;
        bytes memory data;

        (success, data) = abi.decode(log[2].toBytes(), (bool, bytes));

        return
            Protocol.CrossChainResponse(
                eventTopics[1].toUint(),
                receipt[0].toBoolean(),
                success,
                data,
                transaction[3].toAddress()
            );
    }
}
