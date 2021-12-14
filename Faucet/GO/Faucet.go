// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package GO

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ApinameMetaData contains all meta data concerning the Apiname contract.
var ApinameMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"withdraw_ETH\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"Empty_Limit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Get_Contract_Amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Limit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"NewOne\",\"type\":\"address\"}],\"name\":\"Transfer_Other\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ApinameABI is the input ABI used to generate the binding from.
// Deprecated: Use ApinameMetaData.ABI instead.
var ApinameABI = ApinameMetaData.ABI

// Apiname is an auto generated Go binding around an Ethereum contract.
type Apiname struct {
	ApinameCaller     // Read-only binding to the contract
	ApinameTransactor // Write-only binding to the contract
	ApinameFilterer   // Log filterer for contract events
}

// ApinameCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApinameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApinameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApinameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApinameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApinameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApinameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApinameSession struct {
	Contract     *Apiname          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApinameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApinameCallerSession struct {
	Contract *ApinameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ApinameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApinameTransactorSession struct {
	Contract     *ApinameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ApinameRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApinameRaw struct {
	Contract *Apiname // Generic contract binding to access the raw methods on
}

// ApinameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApinameCallerRaw struct {
	Contract *ApinameCaller // Generic read-only contract binding to access the raw methods on
}

// ApinameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApinameTransactorRaw struct {
	Contract *ApinameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApiname creates a new instance of Apiname, bound to a specific deployed contract.
func NewApiname(address common.Address, backend bind.ContractBackend) (*Apiname, error) {
	contract, err := bindApiname(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Apiname{ApinameCaller: ApinameCaller{contract: contract}, ApinameTransactor: ApinameTransactor{contract: contract}, ApinameFilterer: ApinameFilterer{contract: contract}}, nil
}

// NewApinameCaller creates a new read-only instance of Apiname, bound to a specific deployed contract.
func NewApinameCaller(address common.Address, caller bind.ContractCaller) (*ApinameCaller, error) {
	contract, err := bindApiname(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApinameCaller{contract: contract}, nil
}

// NewApinameTransactor creates a new write-only instance of Apiname, bound to a specific deployed contract.
func NewApinameTransactor(address common.Address, transactor bind.ContractTransactor) (*ApinameTransactor, error) {
	contract, err := bindApiname(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApinameTransactor{contract: contract}, nil
}

// NewApinameFilterer creates a new log filterer instance of Apiname, bound to a specific deployed contract.
func NewApinameFilterer(address common.Address, filterer bind.ContractFilterer) (*ApinameFilterer, error) {
	contract, err := bindApiname(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApinameFilterer{contract: contract}, nil
}

// bindApiname binds a generic wrapper to an already deployed contract.
func bindApiname(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApinameABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Apiname *ApinameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Apiname.Contract.ApinameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Apiname *ApinameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Apiname.Contract.ApinameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Apiname *ApinameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Apiname.Contract.ApinameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Apiname *ApinameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Apiname.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Apiname *ApinameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Apiname.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Apiname *ApinameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Apiname.Contract.contract.Transact(opts, method, params...)
}

// GetContractAmount is a free data retrieval call binding the contract method 0xa642559b.
//
// Solidity: function Get_Contract_Amount() view returns(uint256)
func (_Apiname *ApinameCaller) GetContractAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Apiname.contract.Call(opts, &out, "Get_Contract_Amount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractAmount is a free data retrieval call binding the contract method 0xa642559b.
//
// Solidity: function Get_Contract_Amount() view returns(uint256)
func (_Apiname *ApinameSession) GetContractAmount() (*big.Int, error) {
	return _Apiname.Contract.GetContractAmount(&_Apiname.CallOpts)
}

// GetContractAmount is a free data retrieval call binding the contract method 0xa642559b.
//
// Solidity: function Get_Contract_Amount() view returns(uint256)
func (_Apiname *ApinameCallerSession) GetContractAmount() (*big.Int, error) {
	return _Apiname.Contract.GetContractAmount(&_Apiname.CallOpts)
}

// Limit is a free data retrieval call binding the contract method 0xe012563a.
//
// Solidity: function Limit(address ) view returns(uint256)
func (_Apiname *ApinameCaller) Limit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Apiname.contract.Call(opts, &out, "Limit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Limit is a free data retrieval call binding the contract method 0xe012563a.
//
// Solidity: function Limit(address ) view returns(uint256)
func (_Apiname *ApinameSession) Limit(arg0 common.Address) (*big.Int, error) {
	return _Apiname.Contract.Limit(&_Apiname.CallOpts, arg0)
}

// Limit is a free data retrieval call binding the contract method 0xe012563a.
//
// Solidity: function Limit(address ) view returns(uint256)
func (_Apiname *ApinameCallerSession) Limit(arg0 common.Address) (*big.Int, error) {
	return _Apiname.Contract.Limit(&_Apiname.CallOpts, arg0)
}

// EmptyLimit is a paid mutator transaction binding the contract method 0xcfb4c43a.
//
// Solidity: function Empty_Limit(address receiver) returns()
func (_Apiname *ApinameTransactor) EmptyLimit(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _Apiname.contract.Transact(opts, "Empty_Limit", receiver)
}

// EmptyLimit is a paid mutator transaction binding the contract method 0xcfb4c43a.
//
// Solidity: function Empty_Limit(address receiver) returns()
func (_Apiname *ApinameSession) EmptyLimit(receiver common.Address) (*types.Transaction, error) {
	return _Apiname.Contract.EmptyLimit(&_Apiname.TransactOpts, receiver)
}

// EmptyLimit is a paid mutator transaction binding the contract method 0xcfb4c43a.
//
// Solidity: function Empty_Limit(address receiver) returns()
func (_Apiname *ApinameTransactorSession) EmptyLimit(receiver common.Address) (*types.Transaction, error) {
	return _Apiname.Contract.EmptyLimit(&_Apiname.TransactOpts, receiver)
}

// TransferOther is a paid mutator transaction binding the contract method 0x49128d82.
//
// Solidity: function Transfer_Other(address NewOne) returns()
func (_Apiname *ApinameTransactor) TransferOther(opts *bind.TransactOpts, NewOne common.Address) (*types.Transaction, error) {
	return _Apiname.contract.Transact(opts, "Transfer_Other", NewOne)
}

// TransferOther is a paid mutator transaction binding the contract method 0x49128d82.
//
// Solidity: function Transfer_Other(address NewOne) returns()
func (_Apiname *ApinameSession) TransferOther(NewOne common.Address) (*types.Transaction, error) {
	return _Apiname.Contract.TransferOther(&_Apiname.TransactOpts, NewOne)
}

// TransferOther is a paid mutator transaction binding the contract method 0x49128d82.
//
// Solidity: function Transfer_Other(address NewOne) returns()
func (_Apiname *ApinameTransactorSession) TransferOther(NewOne common.Address) (*types.Transaction, error) {
	return _Apiname.Contract.TransferOther(&_Apiname.TransactOpts, NewOne)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 value, address receiver) returns()
func (_Apiname *ApinameTransactor) Withdraw(opts *bind.TransactOpts, value *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Apiname.contract.Transact(opts, "withdraw", value, receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 value, address receiver) returns()
func (_Apiname *ApinameSession) Withdraw(value *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Apiname.Contract.Withdraw(&_Apiname.TransactOpts, value, receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 value, address receiver) returns()
func (_Apiname *ApinameTransactorSession) Withdraw(value *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Apiname.Contract.Withdraw(&_Apiname.TransactOpts, value, receiver)
}

// ApinameWithdrawETHIterator is returned from FilterWithdrawETH and is used to iterate over the raw logs and unpacked data for WithdrawETH events raised by the Apiname contract.
type ApinameWithdrawETHIterator struct {
	Event *ApinameWithdrawETH // Event containing the contract specifics and raw log

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
func (it *ApinameWithdrawETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApinameWithdrawETH)
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
		it.Event = new(ApinameWithdrawETH)
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
func (it *ApinameWithdrawETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApinameWithdrawETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApinameWithdrawETH represents a WithdrawETH event raised by the Apiname contract.
type ApinameWithdrawETH struct {
	Receiver common.Address
	Value    *big.Int
	Time     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawETH is a free log retrieval operation binding the contract event 0x4f6eea60fe38a20b37423537036ac5b0683e3652052cf31c42858fa8ad408181.
//
// Solidity: event withdraw_ETH(address receiver, uint256 value, uint256 time)
func (_Apiname *ApinameFilterer) FilterWithdrawETH(opts *bind.FilterOpts) (*ApinameWithdrawETHIterator, error) {

	logs, sub, err := _Apiname.contract.FilterLogs(opts, "withdraw_ETH")
	if err != nil {
		return nil, err
	}
	return &ApinameWithdrawETHIterator{contract: _Apiname.contract, event: "withdraw_ETH", logs: logs, sub: sub}, nil
}

// WatchWithdrawETH is a free log subscription operation binding the contract event 0x4f6eea60fe38a20b37423537036ac5b0683e3652052cf31c42858fa8ad408181.
//
// Solidity: event withdraw_ETH(address receiver, uint256 value, uint256 time)
func (_Apiname *ApinameFilterer) WatchWithdrawETH(opts *bind.WatchOpts, sink chan<- *ApinameWithdrawETH) (event.Subscription, error) {

	logs, sub, err := _Apiname.contract.WatchLogs(opts, "withdraw_ETH")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApinameWithdrawETH)
				if err := _Apiname.contract.UnpackLog(event, "withdraw_ETH", log); err != nil {
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

// ParseWithdrawETH is a log parse operation binding the contract event 0x4f6eea60fe38a20b37423537036ac5b0683e3652052cf31c42858fa8ad408181.
//
// Solidity: event withdraw_ETH(address receiver, uint256 value, uint256 time)
func (_Apiname *ApinameFilterer) ParseWithdrawETH(log types.Log) (*ApinameWithdrawETH, error) {
	event := new(ApinameWithdrawETH)
	if err := _Apiname.contract.UnpackLog(event, "withdraw_ETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
