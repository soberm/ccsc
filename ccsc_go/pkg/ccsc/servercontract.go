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

// ProtocolProof is an auto generated low-level Go binding around an user-defined struct.
type ProtocolProof struct {
	RlpHeader              []byte
	RlpEncodedTx           []byte
	RlpEncodedReceipt      []byte
	Path                   []byte
	RlpEncodedTxNodes      []byte
	RlpEncodedReceiptNodes []byte
}

// ServerCrossChainCall is an auto generated low-level Go binding around an user-defined struct.
type ServerCrossChainCall struct {
	Id       *big.Int
	Sender   common.Address
	Origin   common.Address
	Status   bool
	GasLimit *big.Int
	CallData []byte
}

// ServerContractABI is the input ABI used to generate the binding from.
const ServerContractABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_relayAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"CrossChainCallExecuted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CONFIRMATIONS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAY_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rlpTransaction\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rlpReceipt\",\"type\":\"bytes\"}],\"name\":\"decodeTransactionReceipt\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structServer.CrossChainCall\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"rlpHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rlpEncodedTx\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rlpEncodedReceipt\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rlpEncodedTxNodes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rlpEncodedReceiptNodes\",\"type\":\"bytes\"}],\"internalType\":\"structProtocol.Proof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"executeCrossChainCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxyAddress\",\"type\":\"address\"}],\"name\":\"setProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ServerContract is an auto generated Go binding around an Ethereum contract.
type ServerContract struct {
	ServerContractCaller     // Read-only binding to the contract
	ServerContractTransactor // Write-only binding to the contract
	ServerContractFilterer   // Log filterer for contract events
}

// ServerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ServerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ServerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ServerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ServerContractSession struct {
	Contract     *ServerContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ServerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ServerContractCallerSession struct {
	Contract *ServerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ServerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ServerContractTransactorSession struct {
	Contract     *ServerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ServerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ServerContractRaw struct {
	Contract *ServerContract // Generic contract binding to access the raw methods on
}

// ServerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ServerContractCallerRaw struct {
	Contract *ServerContractCaller // Generic read-only contract binding to access the raw methods on
}

// ServerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ServerContractTransactorRaw struct {
	Contract *ServerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewServerContract creates a new instance of ServerContract, bound to a specific deployed contract.
func NewServerContract(address common.Address, backend bind.ContractBackend) (*ServerContract, error) {
	contract, err := bindServerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ServerContract{ServerContractCaller: ServerContractCaller{contract: contract}, ServerContractTransactor: ServerContractTransactor{contract: contract}, ServerContractFilterer: ServerContractFilterer{contract: contract}}, nil
}

// NewServerContractCaller creates a new read-only instance of ServerContract, bound to a specific deployed contract.
func NewServerContractCaller(address common.Address, caller bind.ContractCaller) (*ServerContractCaller, error) {
	contract, err := bindServerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ServerContractCaller{contract: contract}, nil
}

// NewServerContractTransactor creates a new write-only instance of ServerContract, bound to a specific deployed contract.
func NewServerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ServerContractTransactor, error) {
	contract, err := bindServerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ServerContractTransactor{contract: contract}, nil
}

// NewServerContractFilterer creates a new log filterer instance of ServerContract, bound to a specific deployed contract.
func NewServerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ServerContractFilterer, error) {
	contract, err := bindServerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ServerContractFilterer{contract: contract}, nil
}

// bindServerContract binds a generic wrapper to an already deployed contract.
func bindServerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ServerContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ServerContract *ServerContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ServerContract.Contract.ServerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ServerContract *ServerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ServerContract.Contract.ServerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ServerContract *ServerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ServerContract.Contract.ServerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ServerContract *ServerContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ServerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ServerContract *ServerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ServerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ServerContract *ServerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ServerContract.Contract.contract.Transact(opts, method, params...)
}

// CONFIRMATIONS is a free data retrieval call binding the contract method 0x0438b0a4.
//
// Solidity: function CONFIRMATIONS() view returns(uint8)
func (_ServerContract *ServerContractCaller) CONFIRMATIONS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ServerContract.contract.Call(opts, &out, "CONFIRMATIONS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CONFIRMATIONS is a free data retrieval call binding the contract method 0x0438b0a4.
//
// Solidity: function CONFIRMATIONS() view returns(uint8)
func (_ServerContract *ServerContractSession) CONFIRMATIONS() (uint8, error) {
	return _ServerContract.Contract.CONFIRMATIONS(&_ServerContract.CallOpts)
}

// CONFIRMATIONS is a free data retrieval call binding the contract method 0x0438b0a4.
//
// Solidity: function CONFIRMATIONS() view returns(uint8)
func (_ServerContract *ServerContractCallerSession) CONFIRMATIONS() (uint8, error) {
	return _ServerContract.Contract.CONFIRMATIONS(&_ServerContract.CallOpts)
}

// RELAYFEE is a free data retrieval call binding the contract method 0x6fbe391c.
//
// Solidity: function RELAY_FEE() view returns(uint256)
func (_ServerContract *ServerContractCaller) RELAYFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ServerContract.contract.Call(opts, &out, "RELAY_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RELAYFEE is a free data retrieval call binding the contract method 0x6fbe391c.
//
// Solidity: function RELAY_FEE() view returns(uint256)
func (_ServerContract *ServerContractSession) RELAYFEE() (*big.Int, error) {
	return _ServerContract.Contract.RELAYFEE(&_ServerContract.CallOpts)
}

// RELAYFEE is a free data retrieval call binding the contract method 0x6fbe391c.
//
// Solidity: function RELAY_FEE() view returns(uint256)
func (_ServerContract *ServerContractCallerSession) RELAYFEE() (*big.Int, error) {
	return _ServerContract.Contract.RELAYFEE(&_ServerContract.CallOpts)
}

// DecodeTransactionReceipt is a free data retrieval call binding the contract method 0x1eec6e41.
//
// Solidity: function decodeTransactionReceipt(bytes rlpTransaction, bytes rlpReceipt) pure returns((uint256,address,address,bool,uint256,bytes))
func (_ServerContract *ServerContractCaller) DecodeTransactionReceipt(opts *bind.CallOpts, rlpTransaction []byte, rlpReceipt []byte) (ServerCrossChainCall, error) {
	var out []interface{}
	err := _ServerContract.contract.Call(opts, &out, "decodeTransactionReceipt", rlpTransaction, rlpReceipt)

	if err != nil {
		return *new(ServerCrossChainCall), err
	}

	out0 := *abi.ConvertType(out[0], new(ServerCrossChainCall)).(*ServerCrossChainCall)

	return out0, err

}

// DecodeTransactionReceipt is a free data retrieval call binding the contract method 0x1eec6e41.
//
// Solidity: function decodeTransactionReceipt(bytes rlpTransaction, bytes rlpReceipt) pure returns((uint256,address,address,bool,uint256,bytes))
func (_ServerContract *ServerContractSession) DecodeTransactionReceipt(rlpTransaction []byte, rlpReceipt []byte) (ServerCrossChainCall, error) {
	return _ServerContract.Contract.DecodeTransactionReceipt(&_ServerContract.CallOpts, rlpTransaction, rlpReceipt)
}

// DecodeTransactionReceipt is a free data retrieval call binding the contract method 0x1eec6e41.
//
// Solidity: function decodeTransactionReceipt(bytes rlpTransaction, bytes rlpReceipt) pure returns((uint256,address,address,bool,uint256,bytes))
func (_ServerContract *ServerContractCallerSession) DecodeTransactionReceipt(rlpTransaction []byte, rlpReceipt []byte) (ServerCrossChainCall, error) {
	return _ServerContract.Contract.DecodeTransactionReceipt(&_ServerContract.CallOpts, rlpTransaction, rlpReceipt)
}

// ExecuteCrossChainCall is a paid mutator transaction binding the contract method 0x3998c8c4.
//
// Solidity: function executeCrossChainCall((bytes,bytes,bytes,bytes,bytes,bytes) proof) payable returns()
func (_ServerContract *ServerContractTransactor) ExecuteCrossChainCall(opts *bind.TransactOpts, proof ProtocolProof) (*types.Transaction, error) {
	return _ServerContract.contract.Transact(opts, "executeCrossChainCall", proof)
}

// ExecuteCrossChainCall is a paid mutator transaction binding the contract method 0x3998c8c4.
//
// Solidity: function executeCrossChainCall((bytes,bytes,bytes,bytes,bytes,bytes) proof) payable returns()
func (_ServerContract *ServerContractSession) ExecuteCrossChainCall(proof ProtocolProof) (*types.Transaction, error) {
	return _ServerContract.Contract.ExecuteCrossChainCall(&_ServerContract.TransactOpts, proof)
}

// ExecuteCrossChainCall is a paid mutator transaction binding the contract method 0x3998c8c4.
//
// Solidity: function executeCrossChainCall((bytes,bytes,bytes,bytes,bytes,bytes) proof) payable returns()
func (_ServerContract *ServerContractTransactorSession) ExecuteCrossChainCall(proof ProtocolProof) (*types.Transaction, error) {
	return _ServerContract.Contract.ExecuteCrossChainCall(&_ServerContract.TransactOpts, proof)
}

// SetProxy is a paid mutator transaction binding the contract method 0x97107d6d.
//
// Solidity: function setProxy(address proxyAddress) returns()
func (_ServerContract *ServerContractTransactor) SetProxy(opts *bind.TransactOpts, proxyAddress common.Address) (*types.Transaction, error) {
	return _ServerContract.contract.Transact(opts, "setProxy", proxyAddress)
}

// SetProxy is a paid mutator transaction binding the contract method 0x97107d6d.
//
// Solidity: function setProxy(address proxyAddress) returns()
func (_ServerContract *ServerContractSession) SetProxy(proxyAddress common.Address) (*types.Transaction, error) {
	return _ServerContract.Contract.SetProxy(&_ServerContract.TransactOpts, proxyAddress)
}

// SetProxy is a paid mutator transaction binding the contract method 0x97107d6d.
//
// Solidity: function setProxy(address proxyAddress) returns()
func (_ServerContract *ServerContractTransactorSession) SetProxy(proxyAddress common.Address) (*types.Transaction, error) {
	return _ServerContract.Contract.SetProxy(&_ServerContract.TransactOpts, proxyAddress)
}

// ServerContractCrossChainCallExecutedIterator is returned from FilterCrossChainCallExecuted and is used to iterate over the raw logs and unpacked data for CrossChainCallExecuted events raised by the ServerContract contract.
type ServerContractCrossChainCallExecutedIterator struct {
	Event *ServerContractCrossChainCallExecuted // Event containing the contract specifics and raw log

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
func (it *ServerContractCrossChainCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServerContractCrossChainCallExecuted)
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
		it.Event = new(ServerContractCrossChainCallExecuted)
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
func (it *ServerContractCrossChainCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServerContractCrossChainCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServerContractCrossChainCallExecuted represents a CrossChainCallExecuted event raised by the ServerContract contract.
type ServerContractCrossChainCallExecuted struct {
	Id      *big.Int
	Success bool
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCrossChainCallExecuted is a free log retrieval operation binding the contract event 0xd502609f911e36dbff128057b05aee250aad0efda2f06a626f99b6740d193761.
//
// Solidity: event CrossChainCallExecuted(uint256 indexed id, bool success, bytes data)
func (_ServerContract *ServerContractFilterer) FilterCrossChainCallExecuted(opts *bind.FilterOpts, id []*big.Int) (*ServerContractCrossChainCallExecutedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ServerContract.contract.FilterLogs(opts, "CrossChainCallExecuted", idRule)
	if err != nil {
		return nil, err
	}
	return &ServerContractCrossChainCallExecutedIterator{contract: _ServerContract.contract, event: "CrossChainCallExecuted", logs: logs, sub: sub}, nil
}

// WatchCrossChainCallExecuted is a free log subscription operation binding the contract event 0xd502609f911e36dbff128057b05aee250aad0efda2f06a626f99b6740d193761.
//
// Solidity: event CrossChainCallExecuted(uint256 indexed id, bool success, bytes data)
func (_ServerContract *ServerContractFilterer) WatchCrossChainCallExecuted(opts *bind.WatchOpts, sink chan<- *ServerContractCrossChainCallExecuted, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ServerContract.contract.WatchLogs(opts, "CrossChainCallExecuted", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServerContractCrossChainCallExecuted)
				if err := _ServerContract.contract.UnpackLog(event, "CrossChainCallExecuted", log); err != nil {
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

// ParseCrossChainCallExecuted is a log parse operation binding the contract event 0xd502609f911e36dbff128057b05aee250aad0efda2f06a626f99b6740d193761.
//
// Solidity: event CrossChainCallExecuted(uint256 indexed id, bool success, bytes data)
func (_ServerContract *ServerContractFilterer) ParseCrossChainCallExecuted(log types.Log) (*ServerContractCrossChainCallExecuted, error) {
	event := new(ServerContractCrossChainCallExecuted)
	if err := _ServerContract.contract.UnpackLog(event, "CrossChainCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
