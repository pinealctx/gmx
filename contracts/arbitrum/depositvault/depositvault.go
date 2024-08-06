// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package depositvault

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

// DepositvaultMetaData contains all meta data concerning the Depositvault contract.
var DepositvaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractRoleStore\",\"name\":\"_roleStore\",\"type\":\"address\"},{\"internalType\":\"contractDataStore\",\"name\":\"_dataStore\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"EmptyHoldingAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"EmptyTokenTranferGasLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"}],\"name\":\"InvalidNativeTokenSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SelfTransferNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenTransferError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"dataStore\",\"outputs\":[{\"internalType\":\"contractDataStore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"recordTransferIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roleStore\",\"outputs\":[{\"internalType\":\"contractRoleStore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"syncTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"shouldUnwrapNativeToken\",\"type\":\"bool\"}],\"name\":\"transferOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferOutNativeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// DepositvaultABI is the input ABI used to generate the binding from.
// Deprecated: Use DepositvaultMetaData.ABI instead.
var DepositvaultABI = DepositvaultMetaData.ABI

// Depositvault is an auto generated Go binding around an Ethereum contract.
type Depositvault struct {
	DepositvaultCaller     // Read-only binding to the contract
	DepositvaultTransactor // Write-only binding to the contract
	DepositvaultFilterer   // Log filterer for contract events
}

// DepositvaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type DepositvaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositvaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DepositvaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositvaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DepositvaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositvaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DepositvaultSession struct {
	Contract     *Depositvault     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DepositvaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DepositvaultCallerSession struct {
	Contract *DepositvaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DepositvaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DepositvaultTransactorSession struct {
	Contract     *DepositvaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DepositvaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type DepositvaultRaw struct {
	Contract *Depositvault // Generic contract binding to access the raw methods on
}

// DepositvaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DepositvaultCallerRaw struct {
	Contract *DepositvaultCaller // Generic read-only contract binding to access the raw methods on
}

// DepositvaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DepositvaultTransactorRaw struct {
	Contract *DepositvaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDepositvault creates a new instance of Depositvault, bound to a specific deployed contract.
func NewDepositvault(address common.Address, backend bind.ContractBackend) (*Depositvault, error) {
	contract, err := bindDepositvault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Depositvault{DepositvaultCaller: DepositvaultCaller{contract: contract}, DepositvaultTransactor: DepositvaultTransactor{contract: contract}, DepositvaultFilterer: DepositvaultFilterer{contract: contract}}, nil
}

// NewDepositvaultCaller creates a new read-only instance of Depositvault, bound to a specific deployed contract.
func NewDepositvaultCaller(address common.Address, caller bind.ContractCaller) (*DepositvaultCaller, error) {
	contract, err := bindDepositvault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositvaultCaller{contract: contract}, nil
}

// NewDepositvaultTransactor creates a new write-only instance of Depositvault, bound to a specific deployed contract.
func NewDepositvaultTransactor(address common.Address, transactor bind.ContractTransactor) (*DepositvaultTransactor, error) {
	contract, err := bindDepositvault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositvaultTransactor{contract: contract}, nil
}

// NewDepositvaultFilterer creates a new log filterer instance of Depositvault, bound to a specific deployed contract.
func NewDepositvaultFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositvaultFilterer, error) {
	contract, err := bindDepositvault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositvaultFilterer{contract: contract}, nil
}

// bindDepositvault binds a generic wrapper to an already deployed contract.
func bindDepositvault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DepositvaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Depositvault *DepositvaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Depositvault.Contract.DepositvaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Depositvault *DepositvaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositvault.Contract.DepositvaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Depositvault *DepositvaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Depositvault.Contract.DepositvaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Depositvault *DepositvaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Depositvault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Depositvault *DepositvaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositvault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Depositvault *DepositvaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Depositvault.Contract.contract.Transact(opts, method, params...)
}

// DataStore is a free data retrieval call binding the contract method 0x660d0d67.
//
// Solidity: function dataStore() view returns(address)
func (_Depositvault *DepositvaultCaller) DataStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Depositvault.contract.Call(opts, &out, "dataStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataStore is a free data retrieval call binding the contract method 0x660d0d67.
//
// Solidity: function dataStore() view returns(address)
func (_Depositvault *DepositvaultSession) DataStore() (common.Address, error) {
	return _Depositvault.Contract.DataStore(&_Depositvault.CallOpts)
}

// DataStore is a free data retrieval call binding the contract method 0x660d0d67.
//
// Solidity: function dataStore() view returns(address)
func (_Depositvault *DepositvaultCallerSession) DataStore() (common.Address, error) {
	return _Depositvault.Contract.DataStore(&_Depositvault.CallOpts)
}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_Depositvault *DepositvaultCaller) RoleStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Depositvault.contract.Call(opts, &out, "roleStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_Depositvault *DepositvaultSession) RoleStore() (common.Address, error) {
	return _Depositvault.Contract.RoleStore(&_Depositvault.CallOpts)
}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_Depositvault *DepositvaultCallerSession) RoleStore() (common.Address, error) {
	return _Depositvault.Contract.RoleStore(&_Depositvault.CallOpts)
}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_Depositvault *DepositvaultCaller) TokenBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Depositvault.contract.Call(opts, &out, "tokenBalances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_Depositvault *DepositvaultSession) TokenBalances(arg0 common.Address) (*big.Int, error) {
	return _Depositvault.Contract.TokenBalances(&_Depositvault.CallOpts, arg0)
}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_Depositvault *DepositvaultCallerSession) TokenBalances(arg0 common.Address) (*big.Int, error) {
	return _Depositvault.Contract.TokenBalances(&_Depositvault.CallOpts, arg0)
}

// RecordTransferIn is a paid mutator transaction binding the contract method 0x352f9aed.
//
// Solidity: function recordTransferIn(address token) returns(uint256)
func (_Depositvault *DepositvaultTransactor) RecordTransferIn(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Depositvault.contract.Transact(opts, "recordTransferIn", token)
}

// RecordTransferIn is a paid mutator transaction binding the contract method 0x352f9aed.
//
// Solidity: function recordTransferIn(address token) returns(uint256)
func (_Depositvault *DepositvaultSession) RecordTransferIn(token common.Address) (*types.Transaction, error) {
	return _Depositvault.Contract.RecordTransferIn(&_Depositvault.TransactOpts, token)
}

// RecordTransferIn is a paid mutator transaction binding the contract method 0x352f9aed.
//
// Solidity: function recordTransferIn(address token) returns(uint256)
func (_Depositvault *DepositvaultTransactorSession) RecordTransferIn(token common.Address) (*types.Transaction, error) {
	return _Depositvault.Contract.RecordTransferIn(&_Depositvault.TransactOpts, token)
}

// SyncTokenBalance is a paid mutator transaction binding the contract method 0xeb40133f.
//
// Solidity: function syncTokenBalance(address token) returns(uint256)
func (_Depositvault *DepositvaultTransactor) SyncTokenBalance(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Depositvault.contract.Transact(opts, "syncTokenBalance", token)
}

// SyncTokenBalance is a paid mutator transaction binding the contract method 0xeb40133f.
//
// Solidity: function syncTokenBalance(address token) returns(uint256)
func (_Depositvault *DepositvaultSession) SyncTokenBalance(token common.Address) (*types.Transaction, error) {
	return _Depositvault.Contract.SyncTokenBalance(&_Depositvault.TransactOpts, token)
}

// SyncTokenBalance is a paid mutator transaction binding the contract method 0xeb40133f.
//
// Solidity: function syncTokenBalance(address token) returns(uint256)
func (_Depositvault *DepositvaultTransactorSession) SyncTokenBalance(token common.Address) (*types.Transaction, error) {
	return _Depositvault.Contract.SyncTokenBalance(&_Depositvault.TransactOpts, token)
}

// TransferOut is a paid mutator transaction binding the contract method 0x078d3b79.
//
// Solidity: function transferOut(address token, address receiver, uint256 amount) returns()
func (_Depositvault *DepositvaultTransactor) TransferOut(opts *bind.TransactOpts, token common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Depositvault.contract.Transact(opts, "transferOut", token, receiver, amount)
}

// TransferOut is a paid mutator transaction binding the contract method 0x078d3b79.
//
// Solidity: function transferOut(address token, address receiver, uint256 amount) returns()
func (_Depositvault *DepositvaultSession) TransferOut(token common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Depositvault.Contract.TransferOut(&_Depositvault.TransactOpts, token, receiver, amount)
}

// TransferOut is a paid mutator transaction binding the contract method 0x078d3b79.
//
// Solidity: function transferOut(address token, address receiver, uint256 amount) returns()
func (_Depositvault *DepositvaultTransactorSession) TransferOut(token common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Depositvault.Contract.TransferOut(&_Depositvault.TransactOpts, token, receiver, amount)
}

// TransferOut0 is a paid mutator transaction binding the contract method 0x2fb12605.
//
// Solidity: function transferOut(address token, address receiver, uint256 amount, bool shouldUnwrapNativeToken) returns()
func (_Depositvault *DepositvaultTransactor) TransferOut0(opts *bind.TransactOpts, token common.Address, receiver common.Address, amount *big.Int, shouldUnwrapNativeToken bool) (*types.Transaction, error) {
	return _Depositvault.contract.Transact(opts, "transferOut0", token, receiver, amount, shouldUnwrapNativeToken)
}

// TransferOut0 is a paid mutator transaction binding the contract method 0x2fb12605.
//
// Solidity: function transferOut(address token, address receiver, uint256 amount, bool shouldUnwrapNativeToken) returns()
func (_Depositvault *DepositvaultSession) TransferOut0(token common.Address, receiver common.Address, amount *big.Int, shouldUnwrapNativeToken bool) (*types.Transaction, error) {
	return _Depositvault.Contract.TransferOut0(&_Depositvault.TransactOpts, token, receiver, amount, shouldUnwrapNativeToken)
}

// TransferOut0 is a paid mutator transaction binding the contract method 0x2fb12605.
//
// Solidity: function transferOut(address token, address receiver, uint256 amount, bool shouldUnwrapNativeToken) returns()
func (_Depositvault *DepositvaultTransactorSession) TransferOut0(token common.Address, receiver common.Address, amount *big.Int, shouldUnwrapNativeToken bool) (*types.Transaction, error) {
	return _Depositvault.Contract.TransferOut0(&_Depositvault.TransactOpts, token, receiver, amount, shouldUnwrapNativeToken)
}

// TransferOutNativeToken is a paid mutator transaction binding the contract method 0xd443ca94.
//
// Solidity: function transferOutNativeToken(address receiver, uint256 amount) returns()
func (_Depositvault *DepositvaultTransactor) TransferOutNativeToken(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Depositvault.contract.Transact(opts, "transferOutNativeToken", receiver, amount)
}

// TransferOutNativeToken is a paid mutator transaction binding the contract method 0xd443ca94.
//
// Solidity: function transferOutNativeToken(address receiver, uint256 amount) returns()
func (_Depositvault *DepositvaultSession) TransferOutNativeToken(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Depositvault.Contract.TransferOutNativeToken(&_Depositvault.TransactOpts, receiver, amount)
}

// TransferOutNativeToken is a paid mutator transaction binding the contract method 0xd443ca94.
//
// Solidity: function transferOutNativeToken(address receiver, uint256 amount) returns()
func (_Depositvault *DepositvaultTransactorSession) TransferOutNativeToken(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Depositvault.Contract.TransferOutNativeToken(&_Depositvault.TransactOpts, receiver, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Depositvault *DepositvaultTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositvault.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Depositvault *DepositvaultSession) Receive() (*types.Transaction, error) {
	return _Depositvault.Contract.Receive(&_Depositvault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Depositvault *DepositvaultTransactorSession) Receive() (*types.Transaction, error) {
	return _Depositvault.Contract.Receive(&_Depositvault.TransactOpts)
}
