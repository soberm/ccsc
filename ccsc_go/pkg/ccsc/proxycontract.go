// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ccsc

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ProtocolCallOpts is an auto generated low-level Go binding around an user-defined struct.
type ProtocolCallOpts struct {
	GasLimit         *big.Int
	Timeout          *big.Int
	CallBackTo       common.Address
	CallBackName     string
	CallBackGasLimit *big.Int
}

// ProtocolCrossChainRequest is an auto generated low-level Go binding around an user-defined struct.
type ProtocolCrossChainRequest struct {
	Id          *big.Int
	Origin      common.Address
	CallBackTo  common.Address
	CallData    []byte
	Gas         *big.Int
	CallBack    string
	CallBackGas *big.Int
	Timeout     *big.Int
}

// ProtocolCrossChainResponse is an auto generated low-level Go binding around an user-defined struct.
type ProtocolCrossChainResponse struct {
	Id      *big.Int
	Status  bool
	Success bool
	Data    []byte
	Server  common.Address
}

// ProxyContractABI is the input ABI used to generate the binding from.
const ProxyContractABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"serverAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"relayAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"call\",\"type\":\"uint256\"}],\"name\":\"CrossChainCallAcknowledged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"callID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"callBackTo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"callBackName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"callBackGas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"}],\"name\":\"CrossChainCallInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"call\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"name\":\"CrossChainCallPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CALLBACK_PARAMS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONFIRMATIONS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_CALL_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_DEPOSIT_PER_CALL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAY_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"rlpHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rlpEncodedTx\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rlpEncodedReceipt\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rlpEncodedTxNodes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rlpEncodedReceiptNodes\",\"type\":\"bytes\"}],\"internalType\":\"structProtocol.Proof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callBackTo\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"callBack\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"callBackGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"}],\"internalType\":\"structProtocol.CrossChainRequest\",\"name\":\"call\",\"type\":\"tuple\"}],\"name\":\"acknowledgeCrossChainCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rlpTransaction\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rlpReceipt\",\"type\":\"bytes\"}],\"name\":\"decodeTransactionReceipt\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"server\",\"type\":\"address\"}],\"internalType\":\"structProtocol.CrossChainResponse\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"callBackTo\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"callBackName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"callBackGasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structProtocol.CallOpts\",\"name\":\"callOpts\",\"type\":\"tuple\"}],\"name\":\"initiateCrossChainCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callBackTo\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"callBack\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"callBackGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"}],\"internalType\":\"structProtocol.CrossChainRequest\",\"name\":\"call\",\"type\":\"tuple\"}],\"name\":\"prepareCrossChainCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ProxyContract is an auto generated Go binding around an Ethereum contract.
type ProxyContract struct {
	ProxyContractCaller     // Read-only binding to the contract
	ProxyContractTransactor // Write-only binding to the contract
	ProxyContractFilterer   // Log filterer for contract events
}

// ProxyContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxyContractSession struct {
	Contract     *ProxyContract    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyContractCallerSession struct {
	Contract *ProxyContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ProxyContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyContractTransactorSession struct {
	Contract     *ProxyContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ProxyContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyContractRaw struct {
	Contract *ProxyContract // Generic contract binding to access the raw methods on
}

// ProxyContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyContractCallerRaw struct {
	Contract *ProxyContractCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyContractTransactorRaw struct {
	Contract *ProxyContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxyContract creates a new instance of ProxyContract, bound to a specific deployed contract.
func NewProxyContract(address common.Address, backend bind.ContractBackend) (*ProxyContract, error) {
	contract, err := bindProxyContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProxyContract{ProxyContractCaller: ProxyContractCaller{contract: contract}, ProxyContractTransactor: ProxyContractTransactor{contract: contract}, ProxyContractFilterer: ProxyContractFilterer{contract: contract}}, nil
}

// NewProxyContractCaller creates a new read-only instance of ProxyContract, bound to a specific deployed contract.
func NewProxyContractCaller(address common.Address, caller bind.ContractCaller) (*ProxyContractCaller, error) {
	contract, err := bindProxyContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyContractCaller{contract: contract}, nil
}

// NewProxyContractTransactor creates a new write-only instance of ProxyContract, bound to a specific deployed contract.
func NewProxyContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyContractTransactor, error) {
	contract, err := bindProxyContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyContractTransactor{contract: contract}, nil
}

// NewProxyContractFilterer creates a new log filterer instance of ProxyContract, bound to a specific deployed contract.
func NewProxyContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyContractFilterer, error) {
	contract, err := bindProxyContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyContractFilterer{contract: contract}, nil
}

// bindProxyContract binds a generic wrapper to an already deployed contract.
func bindProxyContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyContract *ProxyContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProxyContract.Contract.ProxyContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyContract *ProxyContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyContract.Contract.ProxyContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyContract *ProxyContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyContract.Contract.ProxyContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyContract *ProxyContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProxyContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyContract *ProxyContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyContract *ProxyContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyContract.Contract.contract.Transact(opts, method, params...)
}

// CALLBACKPARAMS is a free data retrieval call binding the contract method 0x9d1f1132.
//
// Solidity: function CALLBACK_PARAMS() view returns(string)
func (_ProxyContract *ProxyContractCaller) CALLBACKPARAMS(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ProxyContract.contract.Call(opts, &out, "CALLBACK_PARAMS")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CALLBACKPARAMS is a free data retrieval call binding the contract method 0x9d1f1132.
//
// Solidity: function CALLBACK_PARAMS() view returns(string)
func (_ProxyContract *ProxyContractSession) CALLBACKPARAMS() (string, error) {
	return _ProxyContract.Contract.CALLBACKPARAMS(&_ProxyContract.CallOpts)
}

// CALLBACKPARAMS is a free data retrieval call binding the contract method 0x9d1f1132.
//
// Solidity: function CALLBACK_PARAMS() view returns(string)
func (_ProxyContract *ProxyContractCallerSession) CALLBACKPARAMS() (string, error) {
	return _ProxyContract.Contract.CALLBACKPARAMS(&_ProxyContract.CallOpts)
}

// CONFIRMATIONS is a free data retrieval call binding the contract method 0x0438b0a4.
//
// Solidity: function CONFIRMATIONS() view returns(uint8)
func (_ProxyContract *ProxyContractCaller) CONFIRMATIONS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ProxyContract.contract.Call(opts, &out, "CONFIRMATIONS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CONFIRMATIONS is a free data retrieval call binding the contract method 0x0438b0a4.
//
// Solidity: function CONFIRMATIONS() view returns(uint8)
func (_ProxyContract *ProxyContractSession) CONFIRMATIONS() (uint8, error) {
	return _ProxyContract.Contract.CONFIRMATIONS(&_ProxyContract.CallOpts)
}

// CONFIRMATIONS is a free data retrieval call binding the contract method 0x0438b0a4.
//
// Solidity: function CONFIRMATIONS() view returns(uint8)
func (_ProxyContract *ProxyContractCallerSession) CONFIRMATIONS() (uint8, error) {
	return _ProxyContract.Contract.CONFIRMATIONS(&_ProxyContract.CallOpts)
}

// MINCALLGAS is a free data retrieval call binding the contract method 0x952c9e32.
//
// Solidity: function MIN_CALL_GAS() view returns(uint256)
func (_ProxyContract *ProxyContractCaller) MINCALLGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProxyContract.contract.Call(opts, &out, "MIN_CALL_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINCALLGAS is a free data retrieval call binding the contract method 0x952c9e32.
//
// Solidity: function MIN_CALL_GAS() view returns(uint256)
func (_ProxyContract *ProxyContractSession) MINCALLGAS() (*big.Int, error) {
	return _ProxyContract.Contract.MINCALLGAS(&_ProxyContract.CallOpts)
}

// MINCALLGAS is a free data retrieval call binding the contract method 0x952c9e32.
//
// Solidity: function MIN_CALL_GAS() view returns(uint256)
func (_ProxyContract *ProxyContractCallerSession) MINCALLGAS() (*big.Int, error) {
	return _ProxyContract.Contract.MINCALLGAS(&_ProxyContract.CallOpts)
}

// MINDEPOSITPERCALL is a free data retrieval call binding the contract method 0x03cb3cde.
//
// Solidity: function MIN_DEPOSIT_PER_CALL() view returns(uint256)
func (_ProxyContract *ProxyContractCaller) MINDEPOSITPERCALL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProxyContract.contract.Call(opts, &out, "MIN_DEPOSIT_PER_CALL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINDEPOSITPERCALL is a free data retrieval call binding the contract method 0x03cb3cde.
//
// Solidity: function MIN_DEPOSIT_PER_CALL() view returns(uint256)
func (_ProxyContract *ProxyContractSession) MINDEPOSITPERCALL() (*big.Int, error) {
	return _ProxyContract.Contract.MINDEPOSITPERCALL(&_ProxyContract.CallOpts)
}

// MINDEPOSITPERCALL is a free data retrieval call binding the contract method 0x03cb3cde.
//
// Solidity: function MIN_DEPOSIT_PER_CALL() view returns(uint256)
func (_ProxyContract *ProxyContractCallerSession) MINDEPOSITPERCALL() (*big.Int, error) {
	return _ProxyContract.Contract.MINDEPOSITPERCALL(&_ProxyContract.CallOpts)
}

// RELAYFEE is a free data retrieval call binding the contract method 0x6fbe391c.
//
// Solidity: function RELAY_FEE() view returns(uint256)
func (_ProxyContract *ProxyContractCaller) RELAYFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProxyContract.contract.Call(opts, &out, "RELAY_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RELAYFEE is a free data retrieval call binding the contract method 0x6fbe391c.
//
// Solidity: function RELAY_FEE() view returns(uint256)
func (_ProxyContract *ProxyContractSession) RELAYFEE() (*big.Int, error) {
	return _ProxyContract.Contract.RELAYFEE(&_ProxyContract.CallOpts)
}

// RELAYFEE is a free data retrieval call binding the contract method 0x6fbe391c.
//
// Solidity: function RELAY_FEE() view returns(uint256)
func (_ProxyContract *ProxyContractCallerSession) RELAYFEE() (*big.Int, error) {
	return _ProxyContract.Contract.RELAYFEE(&_ProxyContract.CallOpts)
}

// DecodeTransactionReceipt is a free data retrieval call binding the contract method 0x1eec6e41.
//
// Solidity: function decodeTransactionReceipt(bytes rlpTransaction, bytes rlpReceipt) pure returns((uint256,bool,bool,bytes,address))
func (_ProxyContract *ProxyContractCaller) DecodeTransactionReceipt(opts *bind.CallOpts, rlpTransaction []byte, rlpReceipt []byte) (ProtocolCrossChainResponse, error) {
	var out []interface{}
	err := _ProxyContract.contract.Call(opts, &out, "decodeTransactionReceipt", rlpTransaction, rlpReceipt)

	if err != nil {
		return *new(ProtocolCrossChainResponse), err
	}

	out0 := *abi.ConvertType(out[0], new(ProtocolCrossChainResponse)).(*ProtocolCrossChainResponse)

	return out0, err

}

// DecodeTransactionReceipt is a free data retrieval call binding the contract method 0x1eec6e41.
//
// Solidity: function decodeTransactionReceipt(bytes rlpTransaction, bytes rlpReceipt) pure returns((uint256,bool,bool,bytes,address))
func (_ProxyContract *ProxyContractSession) DecodeTransactionReceipt(rlpTransaction []byte, rlpReceipt []byte) (ProtocolCrossChainResponse, error) {
	return _ProxyContract.Contract.DecodeTransactionReceipt(&_ProxyContract.CallOpts, rlpTransaction, rlpReceipt)
}

// DecodeTransactionReceipt is a free data retrieval call binding the contract method 0x1eec6e41.
//
// Solidity: function decodeTransactionReceipt(bytes rlpTransaction, bytes rlpReceipt) pure returns((uint256,bool,bool,bytes,address))
func (_ProxyContract *ProxyContractCallerSession) DecodeTransactionReceipt(rlpTransaction []byte, rlpReceipt []byte) (ProtocolCrossChainResponse, error) {
	return _ProxyContract.Contract.DecodeTransactionReceipt(&_ProxyContract.CallOpts, rlpTransaction, rlpReceipt)
}

// AcknowledgeCrossChainCall is a paid mutator transaction binding the contract method 0x82e60531.
//
// Solidity: function acknowledgeCrossChainCall((bytes,bytes,bytes,bytes,bytes,bytes) proof, (uint256,address,address,bytes,uint256,string,uint256,uint256) call) payable returns()
func (_ProxyContract *ProxyContractTransactor) AcknowledgeCrossChainCall(opts *bind.TransactOpts, proof ProtocolProof, call ProtocolCrossChainRequest) (*types.Transaction, error) {
	return _ProxyContract.contract.Transact(opts, "acknowledgeCrossChainCall", proof, call)
}

// AcknowledgeCrossChainCall is a paid mutator transaction binding the contract method 0x82e60531.
//
// Solidity: function acknowledgeCrossChainCall((bytes,bytes,bytes,bytes,bytes,bytes) proof, (uint256,address,address,bytes,uint256,string,uint256,uint256) call) payable returns()
func (_ProxyContract *ProxyContractSession) AcknowledgeCrossChainCall(proof ProtocolProof, call ProtocolCrossChainRequest) (*types.Transaction, error) {
	return _ProxyContract.Contract.AcknowledgeCrossChainCall(&_ProxyContract.TransactOpts, proof, call)
}

// AcknowledgeCrossChainCall is a paid mutator transaction binding the contract method 0x82e60531.
//
// Solidity: function acknowledgeCrossChainCall((bytes,bytes,bytes,bytes,bytes,bytes) proof, (uint256,address,address,bytes,uint256,string,uint256,uint256) call) payable returns()
func (_ProxyContract *ProxyContractTransactorSession) AcknowledgeCrossChainCall(proof ProtocolProof, call ProtocolCrossChainRequest) (*types.Transaction, error) {
	return _ProxyContract.Contract.AcknowledgeCrossChainCall(&_ProxyContract.TransactOpts, proof, call)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_ProxyContract *ProxyContractTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyContract.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_ProxyContract *ProxyContractSession) Deposit() (*types.Transaction, error) {
	return _ProxyContract.Contract.Deposit(&_ProxyContract.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_ProxyContract *ProxyContractTransactorSession) Deposit() (*types.Transaction, error) {
	return _ProxyContract.Contract.Deposit(&_ProxyContract.TransactOpts)
}

// InitiateCrossChainCall is a paid mutator transaction binding the contract method 0xb1967def.
//
// Solidity: function initiateCrossChainCall(bytes callData, (uint256,uint256,address,string,uint256) callOpts) returns()
func (_ProxyContract *ProxyContractTransactor) InitiateCrossChainCall(opts *bind.TransactOpts, callData []byte, callOpts ProtocolCallOpts) (*types.Transaction, error) {
	return _ProxyContract.contract.Transact(opts, "initiateCrossChainCall", callData, callOpts)
}

// InitiateCrossChainCall is a paid mutator transaction binding the contract method 0xb1967def.
//
// Solidity: function initiateCrossChainCall(bytes callData, (uint256,uint256,address,string,uint256) callOpts) returns()
func (_ProxyContract *ProxyContractSession) InitiateCrossChainCall(callData []byte, callOpts ProtocolCallOpts) (*types.Transaction, error) {
	return _ProxyContract.Contract.InitiateCrossChainCall(&_ProxyContract.TransactOpts, callData, callOpts)
}

// InitiateCrossChainCall is a paid mutator transaction binding the contract method 0xb1967def.
//
// Solidity: function initiateCrossChainCall(bytes callData, (uint256,uint256,address,string,uint256) callOpts) returns()
func (_ProxyContract *ProxyContractTransactorSession) InitiateCrossChainCall(callData []byte, callOpts ProtocolCallOpts) (*types.Transaction, error) {
	return _ProxyContract.Contract.InitiateCrossChainCall(&_ProxyContract.TransactOpts, callData, callOpts)
}

// PrepareCrossChainCall is a paid mutator transaction binding the contract method 0xad9127af.
//
// Solidity: function prepareCrossChainCall((uint256,address,address,bytes,uint256,string,uint256,uint256) call) returns()
func (_ProxyContract *ProxyContractTransactor) PrepareCrossChainCall(opts *bind.TransactOpts, call ProtocolCrossChainRequest) (*types.Transaction, error) {
	return _ProxyContract.contract.Transact(opts, "prepareCrossChainCall", call)
}

// PrepareCrossChainCall is a paid mutator transaction binding the contract method 0xad9127af.
//
// Solidity: function prepareCrossChainCall((uint256,address,address,bytes,uint256,string,uint256,uint256) call) returns()
func (_ProxyContract *ProxyContractSession) PrepareCrossChainCall(call ProtocolCrossChainRequest) (*types.Transaction, error) {
	return _ProxyContract.Contract.PrepareCrossChainCall(&_ProxyContract.TransactOpts, call)
}

// PrepareCrossChainCall is a paid mutator transaction binding the contract method 0xad9127af.
//
// Solidity: function prepareCrossChainCall((uint256,address,address,bytes,uint256,string,uint256,uint256) call) returns()
func (_ProxyContract *ProxyContractTransactorSession) PrepareCrossChainCall(call ProtocolCrossChainRequest) (*types.Transaction, error) {
	return _ProxyContract.Contract.PrepareCrossChainCall(&_ProxyContract.TransactOpts, call)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ProxyContract *ProxyContractTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyContract.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ProxyContract *ProxyContractSession) Withdraw() (*types.Transaction, error) {
	return _ProxyContract.Contract.Withdraw(&_ProxyContract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ProxyContract *ProxyContractTransactorSession) Withdraw() (*types.Transaction, error) {
	return _ProxyContract.Contract.Withdraw(&_ProxyContract.TransactOpts)
}

// ProxyContractCrossChainCallAcknowledgedIterator is returned from FilterCrossChainCallAcknowledged and is used to iterate over the raw logs and unpacked data for CrossChainCallAcknowledged events raised by the ProxyContract contract.
type ProxyContractCrossChainCallAcknowledgedIterator struct {
	Event *ProxyContractCrossChainCallAcknowledged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ProxyContractCrossChainCallAcknowledgedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyContractCrossChainCallAcknowledged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ProxyContractCrossChainCallAcknowledged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ProxyContractCrossChainCallAcknowledgedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyContractCrossChainCallAcknowledgedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyContractCrossChainCallAcknowledged represents a CrossChainCallAcknowledged event raised by the ProxyContract contract.
type ProxyContractCrossChainCallAcknowledged struct {
	Call *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCrossChainCallAcknowledged is a free log retrieval operation binding the contract event 0xeab542fbec26b2357b3073c05a524121bd6d06bfe01d97a1d2b7fcb3c906f3fa.
//
// Solidity: event CrossChainCallAcknowledged(uint256 indexed call)
func (_ProxyContract *ProxyContractFilterer) FilterCrossChainCallAcknowledged(opts *bind.FilterOpts, call []*big.Int) (*ProxyContractCrossChainCallAcknowledgedIterator, error) {

	var callRule []interface{}
	for _, callItem := range call {
		callRule = append(callRule, callItem)
	}

	logs, sub, err := _ProxyContract.contract.FilterLogs(opts, "CrossChainCallAcknowledged", callRule)
	if err != nil {
		return nil, err
	}
	return &ProxyContractCrossChainCallAcknowledgedIterator{contract: _ProxyContract.contract, event: "CrossChainCallAcknowledged", logs: logs, sub: sub}, nil
}

// WatchCrossChainCallAcknowledged is a free log subscription operation binding the contract event 0xeab542fbec26b2357b3073c05a524121bd6d06bfe01d97a1d2b7fcb3c906f3fa.
//
// Solidity: event CrossChainCallAcknowledged(uint256 indexed call)
func (_ProxyContract *ProxyContractFilterer) WatchCrossChainCallAcknowledged(opts *bind.WatchOpts, sink chan<- *ProxyContractCrossChainCallAcknowledged, call []*big.Int) (event.Subscription, error) {

	var callRule []interface{}
	for _, callItem := range call {
		callRule = append(callRule, callItem)
	}

	logs, sub, err := _ProxyContract.contract.WatchLogs(opts, "CrossChainCallAcknowledged", callRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyContractCrossChainCallAcknowledged)
				if err := _ProxyContract.contract.UnpackLog(event, "CrossChainCallAcknowledged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCrossChainCallAcknowledged is a log parse operation binding the contract event 0xeab542fbec26b2357b3073c05a524121bd6d06bfe01d97a1d2b7fcb3c906f3fa.
//
// Solidity: event CrossChainCallAcknowledged(uint256 indexed call)
func (_ProxyContract *ProxyContractFilterer) ParseCrossChainCallAcknowledged(log types.Log) (*ProxyContractCrossChainCallAcknowledged, error) {
	event := new(ProxyContractCrossChainCallAcknowledged)
	if err := _ProxyContract.contract.UnpackLog(event, "CrossChainCallAcknowledged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyContractCrossChainCallInitiatedIterator is returned from FilterCrossChainCallInitiated and is used to iterate over the raw logs and unpacked data for CrossChainCallInitiated events raised by the ProxyContract contract.
type ProxyContractCrossChainCallInitiatedIterator struct {
	Event *ProxyContractCrossChainCallInitiated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ProxyContractCrossChainCallInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyContractCrossChainCallInitiated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ProxyContractCrossChainCallInitiated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ProxyContractCrossChainCallInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyContractCrossChainCallInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyContractCrossChainCallInitiated represents a CrossChainCallInitiated event raised by the ProxyContract contract.
type ProxyContractCrossChainCallInitiated struct {
	CallID       *big.Int
	Origin       common.Address
	CallBackTo   common.Address
	CallData     []byte
	Gas          *big.Int
	CallBackName string
	CallBackGas  *big.Int
	Timeout      *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCrossChainCallInitiated is a free log retrieval operation binding the contract event 0x3d7eb8b1e5bd8ce1ed17bd333f7033672dfb01f82d35aed6e217c1431503893c.
//
// Solidity: event CrossChainCallInitiated(uint256 callID, address origin, address callBackTo, bytes callData, uint256 gas, string callBackName, uint256 callBackGas, uint256 timeout)
func (_ProxyContract *ProxyContractFilterer) FilterCrossChainCallInitiated(opts *bind.FilterOpts) (*ProxyContractCrossChainCallInitiatedIterator, error) {

	logs, sub, err := _ProxyContract.contract.FilterLogs(opts, "CrossChainCallInitiated")
	if err != nil {
		return nil, err
	}
	return &ProxyContractCrossChainCallInitiatedIterator{contract: _ProxyContract.contract, event: "CrossChainCallInitiated", logs: logs, sub: sub}, nil
}

// WatchCrossChainCallInitiated is a free log subscription operation binding the contract event 0x3d7eb8b1e5bd8ce1ed17bd333f7033672dfb01f82d35aed6e217c1431503893c.
//
// Solidity: event CrossChainCallInitiated(uint256 callID, address origin, address callBackTo, bytes callData, uint256 gas, string callBackName, uint256 callBackGas, uint256 timeout)
func (_ProxyContract *ProxyContractFilterer) WatchCrossChainCallInitiated(opts *bind.WatchOpts, sink chan<- *ProxyContractCrossChainCallInitiated) (event.Subscription, error) {

	logs, sub, err := _ProxyContract.contract.WatchLogs(opts, "CrossChainCallInitiated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyContractCrossChainCallInitiated)
				if err := _ProxyContract.contract.UnpackLog(event, "CrossChainCallInitiated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCrossChainCallInitiated is a log parse operation binding the contract event 0x3d7eb8b1e5bd8ce1ed17bd333f7033672dfb01f82d35aed6e217c1431503893c.
//
// Solidity: event CrossChainCallInitiated(uint256 callID, address origin, address callBackTo, bytes callData, uint256 gas, string callBackName, uint256 callBackGas, uint256 timeout)
func (_ProxyContract *ProxyContractFilterer) ParseCrossChainCallInitiated(log types.Log) (*ProxyContractCrossChainCallInitiated, error) {
	event := new(ProxyContractCrossChainCallInitiated)
	if err := _ProxyContract.contract.UnpackLog(event, "CrossChainCallInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyContractCrossChainCallPreparedIterator is returned from FilterCrossChainCallPrepared and is used to iterate over the raw logs and unpacked data for CrossChainCallPrepared events raised by the ProxyContract contract.
type ProxyContractCrossChainCallPreparedIterator struct {
	Event *ProxyContractCrossChainCallPrepared // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ProxyContractCrossChainCallPreparedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyContractCrossChainCallPrepared)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ProxyContractCrossChainCallPrepared)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ProxyContractCrossChainCallPreparedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyContractCrossChainCallPreparedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyContractCrossChainCallPrepared represents a CrossChainCallPrepared event raised by the ProxyContract contract.
type ProxyContractCrossChainCallPrepared struct {
	Call     *big.Int
	Origin   common.Address
	Gas      *big.Int
	CallData []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCrossChainCallPrepared is a free log retrieval operation binding the contract event 0x27c4a4683a3d10ba190b00c77780aeb028bf0e761d7e9d47a469004983ade7e6.
//
// Solidity: event CrossChainCallPrepared(uint256 indexed call, address origin, uint256 gas, bytes callData)
func (_ProxyContract *ProxyContractFilterer) FilterCrossChainCallPrepared(opts *bind.FilterOpts, call []*big.Int) (*ProxyContractCrossChainCallPreparedIterator, error) {

	var callRule []interface{}
	for _, callItem := range call {
		callRule = append(callRule, callItem)
	}

	logs, sub, err := _ProxyContract.contract.FilterLogs(opts, "CrossChainCallPrepared", callRule)
	if err != nil {
		return nil, err
	}
	return &ProxyContractCrossChainCallPreparedIterator{contract: _ProxyContract.contract, event: "CrossChainCallPrepared", logs: logs, sub: sub}, nil
}

// WatchCrossChainCallPrepared is a free log subscription operation binding the contract event 0x27c4a4683a3d10ba190b00c77780aeb028bf0e761d7e9d47a469004983ade7e6.
//
// Solidity: event CrossChainCallPrepared(uint256 indexed call, address origin, uint256 gas, bytes callData)
func (_ProxyContract *ProxyContractFilterer) WatchCrossChainCallPrepared(opts *bind.WatchOpts, sink chan<- *ProxyContractCrossChainCallPrepared, call []*big.Int) (event.Subscription, error) {

	var callRule []interface{}
	for _, callItem := range call {
		callRule = append(callRule, callItem)
	}

	logs, sub, err := _ProxyContract.contract.WatchLogs(opts, "CrossChainCallPrepared", callRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyContractCrossChainCallPrepared)
				if err := _ProxyContract.contract.UnpackLog(event, "CrossChainCallPrepared", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCrossChainCallPrepared is a log parse operation binding the contract event 0x27c4a4683a3d10ba190b00c77780aeb028bf0e761d7e9d47a469004983ade7e6.
//
// Solidity: event CrossChainCallPrepared(uint256 indexed call, address origin, uint256 gas, bytes callData)
func (_ProxyContract *ProxyContractFilterer) ParseCrossChainCallPrepared(log types.Log) (*ProxyContractCrossChainCallPrepared, error) {
	event := new(ProxyContractCrossChainCallPrepared)
	if err := _ProxyContract.contract.UnpackLog(event, "CrossChainCallPrepared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyContractDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the ProxyContract contract.
type ProxyContractDepositedIterator struct {
	Event *ProxyContractDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ProxyContractDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyContractDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ProxyContractDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ProxyContractDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyContractDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyContractDeposited represents a Deposited event raised by the ProxyContract contract.
type ProxyContractDeposited struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed sender, uint256 amount)
func (_ProxyContract *ProxyContractFilterer) FilterDeposited(opts *bind.FilterOpts, sender []common.Address) (*ProxyContractDepositedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ProxyContract.contract.FilterLogs(opts, "Deposited", senderRule)
	if err != nil {
		return nil, err
	}
	return &ProxyContractDepositedIterator{contract: _ProxyContract.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed sender, uint256 amount)
func (_ProxyContract *ProxyContractFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *ProxyContractDeposited, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ProxyContract.contract.WatchLogs(opts, "Deposited", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyContractDeposited)
				if err := _ProxyContract.contract.UnpackLog(event, "Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed sender, uint256 amount)
func (_ProxyContract *ProxyContractFilterer) ParseDeposited(log types.Log) (*ProxyContractDeposited, error) {
	event := new(ProxyContractDeposited)
	if err := _ProxyContract.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyContractWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the ProxyContract contract.
type ProxyContractWithdrawnIterator struct {
	Event *ProxyContractWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ProxyContractWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyContractWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ProxyContractWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ProxyContractWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyContractWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyContractWithdrawn represents a Withdrawn event raised by the ProxyContract contract.
type ProxyContractWithdrawn struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xf45a04d08a70caa7eb4b747571305559ad9fdf4a093afd41506b35c8a306fa94.
//
// Solidity: event Withdrawn(address indexed sender)
func (_ProxyContract *ProxyContractFilterer) FilterWithdrawn(opts *bind.FilterOpts, sender []common.Address) (*ProxyContractWithdrawnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ProxyContract.contract.FilterLogs(opts, "Withdrawn", senderRule)
	if err != nil {
		return nil, err
	}
	return &ProxyContractWithdrawnIterator{contract: _ProxyContract.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xf45a04d08a70caa7eb4b747571305559ad9fdf4a093afd41506b35c8a306fa94.
//
// Solidity: event Withdrawn(address indexed sender)
func (_ProxyContract *ProxyContractFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *ProxyContractWithdrawn, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ProxyContract.contract.WatchLogs(opts, "Withdrawn", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyContractWithdrawn)
				if err := _ProxyContract.contract.UnpackLog(event, "Withdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawn is a log parse operation binding the contract event 0xf45a04d08a70caa7eb4b747571305559ad9fdf4a093afd41506b35c8a306fa94.
//
// Solidity: event Withdrawn(address indexed sender)
func (_ProxyContract *ProxyContractFilterer) ParseWithdrawn(log types.Log) (*ProxyContractWithdrawn, error) {
	event := new(ProxyContractWithdrawn)
	if err := _ProxyContract.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
