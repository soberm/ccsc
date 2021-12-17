// SPDX-License-Identifier: MIT
pragma solidity ^0.7.0;

library Protocol {
    struct Proof {
        bytes rlpHeader;
        bytes rlpEncodedTx;
        bytes rlpEncodedReceipt;
        bytes path;
        bytes rlpEncodedTxNodes;
        bytes rlpEncodedReceiptNodes;
    }

    struct CallOpts {
        uint256 gasLimit;
        uint256 timeout;
        address callBackTo;
        string callBackName;
        uint256 callBackGasLimit;
    }

    struct CrossChainRequest {
        uint256 id;
        address origin;
        address callBackTo;
        bytes callData;
        uint256 gas;
        string callBack;
        uint256 callBackGas;
        uint256 timeout;
    }

    struct CrossChainResponse {
        uint256 id;
        bool status;
        bool success;
        bytes data;
        address server;
    }

    function hash(CrossChainRequest memory call)
        internal
        pure
        returns (bytes32)
    {
        return
            keccak256(
                abi.encodePacked(
                    call.id,
                    call.origin,
                    call.callBackTo,
                    call.callData,
                    call.gas,
                    call.callBack,
                    call.callBackGas,
                    call.timeout
                )
            );
    }
}
