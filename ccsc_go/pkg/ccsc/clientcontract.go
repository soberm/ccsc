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

// ClientContractABI is the input ABI used to generate the binding from.
const ClientContractABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_storageAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Stored\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getRetrievedNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isStored\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"retrieve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"retrieveCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"storeCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ClientContract is an auto generated Go binding around an Ethereum contract.
type ClientContract struct {
	ClientContractCaller     // Read-only binding to the contract
	ClientContractTransactor // Write-only binding to the contract
	ClientContractFilterer   // Log filterer for contract events
}

// ClientContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClientContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClientContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClientContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClientContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClientContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClientContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClientContractSession struct {
	Contract     *ClientContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClientContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClientContractCallerSession struct {
	Contract *ClientContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ClientContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClientContractTransactorSession struct {
	Contract     *ClientContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ClientContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClientContractRaw struct {
	Contract *ClientContract // Generic contract binding to access the raw methods on
}

// ClientContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClientContractCallerRaw struct {
	Contract *ClientContractCaller // Generic read-only contract binding to access the raw methods on
}

// ClientContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClientContractTransactorRaw struct {
	Contract *ClientContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClientContract creates a new instance of ClientContract, bound to a specific deployed contract.
func NewClientContract(address common.Address, backend bind.ContractBackend) (*ClientContract, error) {
	contract, err := bindClientContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClientContract{ClientContractCaller: ClientContractCaller{contract: contract}, ClientContractTransactor: ClientContractTransactor{contract: contract}, ClientContractFilterer: ClientContractFilterer{contract: contract}}, nil
}

// NewClientContractCaller creates a new read-only instance of ClientContract, bound to a specific deployed contract.
func NewClientContractCaller(address common.Address, caller bind.ContractCaller) (*ClientContractCaller, error) {
	contract, err := bindClientContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClientContractCaller{contract: contract}, nil
}

// NewClientContractTransactor creates a new write-only instance of ClientContract, bound to a specific deployed contract.
func NewClientContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ClientContractTransactor, error) {
	contract, err := bindClientContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClientContractTransactor{contract: contract}, nil
}

// NewClientContractFilterer creates a new log filterer instance of ClientContract, bound to a specific deployed contract.
func NewClientContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ClientContractFilterer, error) {
	contract, err := bindClientContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClientContractFilterer{contract: contract}, nil
}

// bindClientContract binds a generic wrapper to an already deployed contract.
func bindClientContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClientContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClientContract *ClientContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClientContract.Contract.ClientContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClientContract *ClientContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClientContract.Contract.ClientContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClientContract *ClientContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClientContract.Contract.ClientContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClientContract *ClientContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClientContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClientContract *ClientContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClientContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClientContract *ClientContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClientContract.Contract.contract.Transact(opts, method, params...)
}

// GetRetrievedNumber is a free data retrieval call binding the contract method 0xe8cc1b6a.
//
// Solidity: function getRetrievedNumber() view returns(uint256)
func (_ClientContract *ClientContractCaller) GetRetrievedNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ClientContract.contract.Call(opts, &out, "getRetrievedNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRetrievedNumber is a free data retrieval call binding the contract method 0xe8cc1b6a.
//
// Solidity: function getRetrievedNumber() view returns(uint256)
func (_ClientContract *ClientContractSession) GetRetrievedNumber() (*big.Int, error) {
	return _ClientContract.Contract.GetRetrievedNumber(&_ClientContract.CallOpts)
}

// GetRetrievedNumber is a free data retrieval call binding the contract method 0xe8cc1b6a.
//
// Solidity: function getRetrievedNumber() view returns(uint256)
func (_ClientContract *ClientContractCallerSession) GetRetrievedNumber() (*big.Int, error) {
	return _ClientContract.Contract.GetRetrievedNumber(&_ClientContract.CallOpts)
}

// IsStored is a free data retrieval call binding the contract method 0xee460c64.
//
// Solidity: function isStored() view returns(bool)
func (_ClientContract *ClientContractCaller) IsStored(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ClientContract.contract.Call(opts, &out, "isStored")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStored is a free data retrieval call binding the contract method 0xee460c64.
//
// Solidity: function isStored() view returns(bool)
func (_ClientContract *ClientContractSession) IsStored() (bool, error) {
	return _ClientContract.Contract.IsStored(&_ClientContract.CallOpts)
}

// IsStored is a free data retrieval call binding the contract method 0xee460c64.
//
// Solidity: function isStored() view returns(bool)
func (_ClientContract *ClientContractCallerSession) IsStored() (bool, error) {
	return _ClientContract.Contract.IsStored(&_ClientContract.CallOpts)
}

// Retrieve is a paid mutator transaction binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() returns()
func (_ClientContract *ClientContractTransactor) Retrieve(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClientContract.contract.Transact(opts, "retrieve")
}

// Retrieve is a paid mutator transaction binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() returns()
func (_ClientContract *ClientContractSession) Retrieve() (*types.Transaction, error) {
	return _ClientContract.Contract.Retrieve(&_ClientContract.TransactOpts)
}

// Retrieve is a paid mutator transaction binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() returns()
func (_ClientContract *ClientContractTransactorSession) Retrieve() (*types.Transaction, error) {
	return _ClientContract.Contract.Retrieve(&_ClientContract.TransactOpts)
}

// RetrieveCallback is a paid mutator transaction binding the contract method 0x0138d899.
//
// Solidity: function retrieveCallback(uint256 id, bytes result, bool success) returns()
func (_ClientContract *ClientContractTransactor) RetrieveCallback(opts *bind.TransactOpts, id *big.Int, result []byte, success bool) (*types.Transaction, error) {
	return _ClientContract.contract.Transact(opts, "retrieveCallback", id, result, success)
}

// RetrieveCallback is a paid mutator transaction binding the contract method 0x0138d899.
//
// Solidity: function retrieveCallback(uint256 id, bytes result, bool success) returns()
func (_ClientContract *ClientContractSession) RetrieveCallback(id *big.Int, result []byte, success bool) (*types.Transaction, error) {
	return _ClientContract.Contract.RetrieveCallback(&_ClientContract.TransactOpts, id, result, success)
}

// RetrieveCallback is a paid mutator transaction binding the contract method 0x0138d899.
//
// Solidity: function retrieveCallback(uint256 id, bytes result, bool success) returns()
func (_ClientContract *ClientContractTransactorSession) RetrieveCallback(id *big.Int, result []byte, success bool) (*types.Transaction, error) {
	return _ClientContract.Contract.RetrieveCallback(&_ClientContract.TransactOpts, id, result, success)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_ClientContract *ClientContractTransactor) Store(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _ClientContract.contract.Transact(opts, "store", num)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_ClientContract *ClientContractSession) Store(num *big.Int) (*types.Transaction, error) {
	return _ClientContract.Contract.Store(&_ClientContract.TransactOpts, num)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_ClientContract *ClientContractTransactorSession) Store(num *big.Int) (*types.Transaction, error) {
	return _ClientContract.Contract.Store(&_ClientContract.TransactOpts, num)
}

// StoreCallback is a paid mutator transaction binding the contract method 0xc88c6969.
//
// Solidity: function storeCallback(uint256 id, bytes result, bool success) returns()
func (_ClientContract *ClientContractTransactor) StoreCallback(opts *bind.TransactOpts, id *big.Int, result []byte, success bool) (*types.Transaction, error) {
	return _ClientContract.contract.Transact(opts, "storeCallback", id, result, success)
}

// StoreCallback is a paid mutator transaction binding the contract method 0xc88c6969.
//
// Solidity: function storeCallback(uint256 id, bytes result, bool success) returns()
func (_ClientContract *ClientContractSession) StoreCallback(id *big.Int, result []byte, success bool) (*types.Transaction, error) {
	return _ClientContract.Contract.StoreCallback(&_ClientContract.TransactOpts, id, result, success)
}

// StoreCallback is a paid mutator transaction binding the contract method 0xc88c6969.
//
// Solidity: function storeCallback(uint256 id, bytes result, bool success) returns()
func (_ClientContract *ClientContractTransactorSession) StoreCallback(id *big.Int, result []byte, success bool) (*types.Transaction, error) {
	return _ClientContract.Contract.StoreCallback(&_ClientContract.TransactOpts, id, result, success)
}

// ClientContractStoredIterator is returned from FilterStored and is used to iterate over the raw logs and unpacked data for Stored events raised by the ClientContract contract.
type ClientContractStoredIterator struct {
	Event *ClientContractStored // Event containing the contract specifics and raw log

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
func (it *ClientContractStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClientContractStored)
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
		it.Event = new(ClientContractStored)
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
func (it *ClientContractStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClientContractStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClientContractStored represents a Stored event raised by the ClientContract contract.
type ClientContractStored struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStored is a free log retrieval operation binding the contract event 0xc6d8c0af6d21f291e7c359603aa97e0ed500f04db6e983b9fce75a91c6b8da6b.
//
// Solidity: event Stored(uint256 indexed id)
func (_ClientContract *ClientContractFilterer) FilterStored(opts *bind.FilterOpts, id []*big.Int) (*ClientContractStoredIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ClientContract.contract.FilterLogs(opts, "Stored", idRule)
	if err != nil {
		return nil, err
	}
	return &ClientContractStoredIterator{contract: _ClientContract.contract, event: "Stored", logs: logs, sub: sub}, nil
}

// WatchStored is a free log subscription operation binding the contract event 0xc6d8c0af6d21f291e7c359603aa97e0ed500f04db6e983b9fce75a91c6b8da6b.
//
// Solidity: event Stored(uint256 indexed id)
func (_ClientContract *ClientContractFilterer) WatchStored(opts *bind.WatchOpts, sink chan<- *ClientContractStored, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ClientContract.contract.WatchLogs(opts, "Stored", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClientContractStored)
				if err := _ClientContract.contract.UnpackLog(event, "Stored", log); err != nil {
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

// ParseStored is a log parse operation binding the contract event 0xc6d8c0af6d21f291e7c359603aa97e0ed500f04db6e983b9fce75a91c6b8da6b.
//
// Solidity: event Stored(uint256 indexed id)
func (_ClientContract *ClientContractFilterer) ParseStored(log types.Log) (*ClientContractStored, error) {
	event := new(ClientContractStored)
	if err := _ClientContract.contract.UnpackLog(event, "Stored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
