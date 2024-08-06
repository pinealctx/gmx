// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package withdrawalvault

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
	_ = abi.ConvertType
)

// WithdrawalvaultMetaData contains all meta data concerning the Withdrawalvault contract.
var WithdrawalvaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractRoleStore\",\"name\":\"_roleStore\",\"type\":\"address\"},{\"internalType\":\"contractDataStore\",\"name\":\"_dataStore\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"EmptyHoldingAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"EmptyTokenTranferGasLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"}],\"name\":\"InvalidNativeTokenSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SelfTransferNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenTransferError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"dataStore\",\"outputs\":[{\"internalType\":\"contractDataStore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"recordTransferIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roleStore\",\"outputs\":[{\"internalType\":\"contractRoleStore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"syncTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"shouldUnwrapNativeToken\",\"type\":\"bool\"}],\"name\":\"transferOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferOutNativeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// WithdrawalvaultABI is the input ABI used to generate the binding from.
// Deprecated: Use WithdrawalvaultMetaData.ABI instead.
var WithdrawalvaultABI = WithdrawalvaultMetaData.ABI

// Withdrawalvault is an auto generated Go binding around an Ethereum contract.
type Withdrawalvault struct {
	WithdrawalvaultCaller     // Read-only binding to the contract
	WithdrawalvaultTransactor // Write-only binding to the contract
	WithdrawalvaultFilterer   // Log filterer for contract events
}

// WithdrawalvaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type WithdrawalvaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawalvaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WithdrawalvaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawalvaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WithdrawalvaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawalvaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WithdrawalvaultSession struct {
	Contract     *Withdrawalvault  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WithdrawalvaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WithdrawalvaultCallerSession struct {
	Contract *WithdrawalvaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// WithdrawalvaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WithdrawalvaultTransactorSession struct {
	Contract     *WithdrawalvaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// WithdrawalvaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type WithdrawalvaultRaw struct {
	Contract *Withdrawalvault // Generic contract binding to access the raw methods on
}

// WithdrawalvaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WithdrawalvaultCallerRaw struct {
	Contract *WithdrawalvaultCaller // Generic read-only contract binding to access the raw methods on
}

// WithdrawalvaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WithdrawalvaultTransactorRaw struct {
	Contract *WithdrawalvaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWithdrawalvault creates a new instance of Withdrawalvault, bound to a specific deployed contract.
func NewWithdrawalvault(address common.Address, backend bind.ContractBackend) (*Withdrawalvault, error) {
	contract, err := bindWithdrawalvault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Withdrawalvault{WithdrawalvaultCaller: WithdrawalvaultCaller{contract: contract}, WithdrawalvaultTransactor: WithdrawalvaultTransactor{contract: contract}, WithdrawalvaultFilterer: WithdrawalvaultFilterer{contract: contract}}, nil
}

// NewWithdrawalvaultCaller creates a new read-only instance of Withdrawalvault, bound to a specific deployed contract.
func NewWithdrawalvaultCaller(address common.Address, caller bind.ContractCaller) (*WithdrawalvaultCaller, error) {
	contract, err := bindWithdrawalvault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawalvaultCaller{contract: contract}, nil
}

// NewWithdrawalvaultTransactor creates a new write-only instance of Withdrawalvault, bound to a specific deployed contract.
func NewWithdrawalvaultTransactor(address common.Address, transactor bind.ContractTransactor) (*WithdrawalvaultTransactor, error) {
	contract, err := bindWithdrawalvault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawalvaultTransactor{contract: contract}, nil
}

// NewWithdrawalvaultFilterer creates a new log filterer instance of Withdrawalvault, bound to a specific deployed contract.
func NewWithdrawalvaultFilterer(address common.Address, filterer bind.ContractFilterer) (*WithdrawalvaultFilterer, error) {
	contract, err := bindWithdrawalvault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WithdrawalvaultFilterer{contract: contract}, nil
}

// bindWithdrawalvault binds a generic wrapper to an already deployed contract.
func bindWithdrawalvault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WithdrawalvaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Withdrawalvault *WithdrawalvaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Withdrawalvault.Contract.WithdrawalvaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Withdrawalvault *WithdrawalvaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Withdrawalvault.Contract.WithdrawalvaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Withdrawalvault *WithdrawalvaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Withdrawalvault.Contract.WithdrawalvaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Withdrawalvault *WithdrawalvaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Withdrawalvault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Withdrawalvault *WithdrawalvaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Withdrawalvault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Withdrawalvault *WithdrawalvaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Withdrawalvault.Contract.contract.Transact(opts, method, params...)
}

// DataStore is a free data retrieval call binding the contract method 0x660d0d67.
//
// Solidity: function dataStore() view returns(address)
func (_Withdrawalvault *WithdrawalvaultCaller) DataStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Withdrawalvault.contract.Call(opts, &out, "dataStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataStore is a free data retrieval call binding the contract method 0x660d0d67.
//
// Solidity: function dataStore() view returns(address)
func (_Withdrawalvault *WithdrawalvaultSession) DataStore() (common.Address, error) {
	return _Withdrawalvault.Contract.DataStore(&_Withdrawalvault.CallOpts)
}

// DataStore is a free data retrieval call binding the contract method 0x660d0d67.
//
// Solidity: function dataStore() view returns(address)
func (_Withdrawalvault *WithdrawalvaultCallerSession) DataStore() (common.Address, error) {
	return _Withdrawalvault.Contract.DataStore(&_Withdrawalvault.CallOpts)
}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_Withdrawalvault *WithdrawalvaultCaller) RoleStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Withdrawalvault.contract.Call(opts, &out, "roleStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_Withdrawalvault *WithdrawalvaultSession) RoleStore() (common.Address, error) {
	return _Withdrawalvault.Contract.RoleStore(&_Withdrawalvault.CallOpts)
}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_Withdrawalvault *WithdrawalvaultCallerSession) RoleStore() (common.Address, error) {
	return _Withdrawalvault.Contract.RoleStore(&_Withdrawalvault.CallOpts)
}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_Withdrawalvault *WithdrawalvaultCaller) TokenBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Withdrawalvault.contract.Call(opts, &out, "tokenBalances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_Withdrawalvault *WithdrawalvaultSession) TokenBalances(arg0 common.Address) (*big.Int, error) {
	return _Withdrawalvault.Contract.TokenBalances(&_Withdrawalvault.CallOpts, arg0)
}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_Withdrawalvault *WithdrawalvaultCallerSession) TokenBalances(arg0 common.Address) (*big.Int, error) {
	return _Withdrawalvault.Contract.TokenBalances(&_Withdrawalvault.CallOpts, arg0)
}

// RecordTransferIn is a paid mutator transaction binding the contract method 0x352f9aed.
//
// Solidity: function recordTransferIn(address token) returns(uint256)
func (_Withdrawalvault *WithdrawalvaultTransactor) RecordTransferIn(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Withdrawalvault.contract.Transact(opts, "recordTransferIn", token)
}

// RecordTransferIn is a paid mutator transaction binding the contract method 0x352f9aed.
//
// Solidity: function recordTransferIn(address token) returns(uint256)
func (_Withdrawalvault *WithdrawalvaultSession) RecordTransferIn(token common.Address) (*types.Transaction, error) {
	return _Withdrawalvault.Contract.RecordTransferIn(&_Withdrawalvault.TransactOpts, token)
}

// RecordTransferIn is a paid mutator transaction binding the contract method 0x352f9aed.
//
// Solidity: function recordTransferIn(address token) returns(uint256)
func (_Withdrawalvault *WithdrawalvaultTransactorSession) RecordTransferIn(token common.Address) (*types.Transaction, error) {
	return _Withdrawalvault.Contract.RecordTransferIn(&_Withdrawalvault.TransactOpts, token)
}

// SyncTokenBalance is a paid mutator transaction binding the contract method 0xeb40133f.
//
// Solidity: function syncTokenBalance(address token) returns(uint256)
func (_Withdrawalvault *WithdrawalvaultTransactor) SyncTokenBalance(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Withdrawalvault.contract.Transact(opts, "syncTokenBalance", token)
}

// SyncTokenBalance is a paid mutator transaction binding the contract method 0xeb40133f.
//
// Solidity: function syncTokenBalance(address token) returns(uint256)
func (_Withdrawalvault *WithdrawalvaultSession) SyncTokenBalance(token common.Address) (*types.Transaction, error) {
	return _Withdrawalvault.Contract.SyncTokenBalance(&_Withdrawalvault.TransactOpts, token)
}

// SyncTokenBalance is a paid mutator transaction binding the contract method 0xeb40133f.
//
// Solidity: function syncTokenBalance(address token) returns(uint256)
func (_Withdrawalvault *WithdrawalvaultTransactorSession) SyncTokenBalance(token common.Address) (*types.Transaction, error) {
	return _Withdrawalvault.Contract.SyncTokenBalance(&_Withdrawalvault.TransactOpts, token)
}

// TransferOut is a paid mutator transaction binding the contract method 0x078d3b79.
//
// Solidity: function transferOut(address token, address receiver, uint256 amount) returns()
func (_Withdrawalvault *WithdrawalvaultTransactor) TransferOut(opts *bind.TransactOpts, token common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Withdrawalvault.contract.Transact(opts, "transferOut", token, receiver, amount)
}

// TransferOut is a paid mutator transaction binding the contract method 0x078d3b79.
//
// Solidity: function transferOut(address token, address receiver, uint256 amount) returns()
func (_Withdrawalvault *WithdrawalvaultSession) TransferOut(token common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Withdrawalvault.Contract.TransferOut(&_Withdrawalvault.TransactOpts, token, receiver, amount)
}

// TransferOut is a paid mutator transaction binding the contract method 0x078d3b79.
//
// Solidity: function transferOut(address token, address receiver, uint256 amount) returns()
func (_Withdrawalvault *WithdrawalvaultTransactorSession) TransferOut(token common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Withdrawalvault.Contract.TransferOut(&_Withdrawalvault.TransactOpts, token, receiver, amount)
}

// TransferOut0 is a paid mutator transaction binding the contract method 0x2fb12605.
//
// Solidity: function transferOut(address token, address receiver, uint256 amount, bool shouldUnwrapNativeToken) returns()
func (_Withdrawalvault *WithdrawalvaultTransactor) TransferOut0(opts *bind.TransactOpts, token common.Address, receiver common.Address, amount *big.Int, shouldUnwrapNativeToken bool) (*types.Transaction, error) {
	return _Withdrawalvault.contract.Transact(opts, "transferOut0", token, receiver, amount, shouldUnwrapNativeToken)
}

// TransferOut0 is a paid mutator transaction binding the contract method 0x2fb12605.
//
// Solidity: function transferOut(address token, address receiver, uint256 amount, bool shouldUnwrapNativeToken) returns()
func (_Withdrawalvault *WithdrawalvaultSession) TransferOut0(token common.Address, receiver common.Address, amount *big.Int, shouldUnwrapNativeToken bool) (*types.Transaction, error) {
	return _Withdrawalvault.Contract.TransferOut0(&_Withdrawalvault.TransactOpts, token, receiver, amount, shouldUnwrapNativeToken)
}

// TransferOut0 is a paid mutator transaction binding the contract method 0x2fb12605.
//
// Solidity: function transferOut(address token, address receiver, uint256 amount, bool shouldUnwrapNativeToken) returns()
func (_Withdrawalvault *WithdrawalvaultTransactorSession) TransferOut0(token common.Address, receiver common.Address, amount *big.Int, shouldUnwrapNativeToken bool) (*types.Transaction, error) {
	return _Withdrawalvault.Contract.TransferOut0(&_Withdrawalvault.TransactOpts, token, receiver, amount, shouldUnwrapNativeToken)
}

// TransferOutNativeToken is a paid mutator transaction binding the contract method 0xd443ca94.
//
// Solidity: function transferOutNativeToken(address receiver, uint256 amount) returns()
func (_Withdrawalvault *WithdrawalvaultTransactor) TransferOutNativeToken(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Withdrawalvault.contract.Transact(opts, "transferOutNativeToken", receiver, amount)
}

// TransferOutNativeToken is a paid mutator transaction binding the contract method 0xd443ca94.
//
// Solidity: function transferOutNativeToken(address receiver, uint256 amount) returns()
func (_Withdrawalvault *WithdrawalvaultSession) TransferOutNativeToken(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Withdrawalvault.Contract.TransferOutNativeToken(&_Withdrawalvault.TransactOpts, receiver, amount)
}

// TransferOutNativeToken is a paid mutator transaction binding the contract method 0xd443ca94.
//
// Solidity: function transferOutNativeToken(address receiver, uint256 amount) returns()
func (_Withdrawalvault *WithdrawalvaultTransactorSession) TransferOutNativeToken(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Withdrawalvault.Contract.TransferOutNativeToken(&_Withdrawalvault.TransactOpts, receiver, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Withdrawalvault *WithdrawalvaultTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Withdrawalvault.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Withdrawalvault *WithdrawalvaultSession) Receive() (*types.Transaction, error) {
	return _Withdrawalvault.Contract.Receive(&_Withdrawalvault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Withdrawalvault *WithdrawalvaultTransactorSession) Receive() (*types.Transaction, error) {
	return _Withdrawalvault.Contract.Receive(&_Withdrawalvault.TransactOpts)
}
