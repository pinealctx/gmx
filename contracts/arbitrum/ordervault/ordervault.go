// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ordervault

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

// OrdervaultMetaData contains all meta data concerning the Ordervault contract.
var OrdervaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractRoleStore\",\"name\":\"_roleStore\",\"type\":\"address\"},{\"internalType\":\"contractDataStore\",\"name\":\"_dataStore\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"EmptyHoldingAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"EmptyTokenTranferGasLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"}],\"name\":\"InvalidNativeTokenSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SelfTransferNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenTransferError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"dataStore\",\"outputs\":[{\"internalType\":\"contractDataStore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"recordTransferIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roleStore\",\"outputs\":[{\"internalType\":\"contractRoleStore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"syncTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"shouldUnwrapNativeToken\",\"type\":\"bool\"}],\"name\":\"transferOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferOutNativeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// OrdervaultABI is the input ABI used to generate the binding from.
// Deprecated: Use OrdervaultMetaData.ABI instead.
var OrdervaultABI = OrdervaultMetaData.ABI

// Ordervault is an auto generated Go binding around an Ethereum contract.
type Ordervault struct {
	OrdervaultCaller     // Read-only binding to the contract
	OrdervaultTransactor // Write-only binding to the contract
	OrdervaultFilterer   // Log filterer for contract events
}

// OrdervaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrdervaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrdervaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrdervaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrdervaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrdervaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrdervaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrdervaultSession struct {
	Contract     *Ordervault       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrdervaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrdervaultCallerSession struct {
	Contract *OrdervaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// OrdervaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrdervaultTransactorSession struct {
	Contract     *OrdervaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OrdervaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrdervaultRaw struct {
	Contract *Ordervault // Generic contract binding to access the raw methods on
}

// OrdervaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrdervaultCallerRaw struct {
	Contract *OrdervaultCaller // Generic read-only contract binding to access the raw methods on
}

// OrdervaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrdervaultTransactorRaw struct {
	Contract *OrdervaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrdervault creates a new instance of Ordervault, bound to a specific deployed contract.
func NewOrdervault(address common.Address, backend bind.ContractBackend) (*Ordervault, error) {
	contract, err := bindOrdervault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ordervault{OrdervaultCaller: OrdervaultCaller{contract: contract}, OrdervaultTransactor: OrdervaultTransactor{contract: contract}, OrdervaultFilterer: OrdervaultFilterer{contract: contract}}, nil
}

// NewOrdervaultCaller creates a new read-only instance of Ordervault, bound to a specific deployed contract.
func NewOrdervaultCaller(address common.Address, caller bind.ContractCaller) (*OrdervaultCaller, error) {
	contract, err := bindOrdervault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrdervaultCaller{contract: contract}, nil
}

// NewOrdervaultTransactor creates a new write-only instance of Ordervault, bound to a specific deployed contract.
func NewOrdervaultTransactor(address common.Address, transactor bind.ContractTransactor) (*OrdervaultTransactor, error) {
	contract, err := bindOrdervault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrdervaultTransactor{contract: contract}, nil
}

// NewOrdervaultFilterer creates a new log filterer instance of Ordervault, bound to a specific deployed contract.
func NewOrdervaultFilterer(address common.Address, filterer bind.ContractFilterer) (*OrdervaultFilterer, error) {
	contract, err := bindOrdervault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrdervaultFilterer{contract: contract}, nil
}

// bindOrdervault binds a generic wrapper to an already deployed contract.
func bindOrdervault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OrdervaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ordervault *OrdervaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ordervault.Contract.OrdervaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ordervault *OrdervaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ordervault.Contract.OrdervaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ordervault *OrdervaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ordervault.Contract.OrdervaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ordervault *OrdervaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ordervault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ordervault *OrdervaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ordervault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ordervault *OrdervaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ordervault.Contract.contract.Transact(opts, method, params...)
}

// DataStore is a free data retrieval call binding the contract method 0x660d0d67.
//
// Solidity: function dataStore() view returns(address)
func (_Ordervault *OrdervaultCaller) DataStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ordervault.contract.Call(opts, &out, "dataStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataStore is a free data retrieval call binding the contract method 0x660d0d67.
//
// Solidity: function dataStore() view returns(address)
func (_Ordervault *OrdervaultSession) DataStore() (common.Address, error) {
	return _Ordervault.Contract.DataStore(&_Ordervault.CallOpts)
}

// DataStore is a free data retrieval call binding the contract method 0x660d0d67.
//
// Solidity: function dataStore() view returns(address)
func (_Ordervault *OrdervaultCallerSession) DataStore() (common.Address, error) {
	return _Ordervault.Contract.DataStore(&_Ordervault.CallOpts)
}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_Ordervault *OrdervaultCaller) RoleStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ordervault.contract.Call(opts, &out, "roleStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_Ordervault *OrdervaultSession) RoleStore() (common.Address, error) {
	return _Ordervault.Contract.RoleStore(&_Ordervault.CallOpts)
}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_Ordervault *OrdervaultCallerSession) RoleStore() (common.Address, error) {
	return _Ordervault.Contract.RoleStore(&_Ordervault.CallOpts)
}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_Ordervault *OrdervaultCaller) TokenBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ordervault.contract.Call(opts, &out, "tokenBalances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_Ordervault *OrdervaultSession) TokenBalances(arg0 common.Address) (*big.Int, error) {
	return _Ordervault.Contract.TokenBalances(&_Ordervault.CallOpts, arg0)
}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_Ordervault *OrdervaultCallerSession) TokenBalances(arg0 common.Address) (*big.Int, error) {
	return _Ordervault.Contract.TokenBalances(&_Ordervault.CallOpts, arg0)
}

// RecordTransferIn is a paid mutator transaction binding the contract method 0x352f9aed.
//
// Solidity: function recordTransferIn(address token) returns(uint256)
func (_Ordervault *OrdervaultTransactor) RecordTransferIn(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Ordervault.contract.Transact(opts, "recordTransferIn", token)
}

// RecordTransferIn is a paid mutator transaction binding the contract method 0x352f9aed.
//
// Solidity: function recordTransferIn(address token) returns(uint256)
func (_Ordervault *OrdervaultSession) RecordTransferIn(token common.Address) (*types.Transaction, error) {
	return _Ordervault.Contract.RecordTransferIn(&_Ordervault.TransactOpts, token)
}

// RecordTransferIn is a paid mutator transaction binding the contract method 0x352f9aed.
//
// Solidity: function recordTransferIn(address token) returns(uint256)
func (_Ordervault *OrdervaultTransactorSession) RecordTransferIn(token common.Address) (*types.Transaction, error) {
	return _Ordervault.Contract.RecordTransferIn(&_Ordervault.TransactOpts, token)
}

// SyncTokenBalance is a paid mutator transaction binding the contract method 0xeb40133f.
//
// Solidity: function syncTokenBalance(address token) returns(uint256)
func (_Ordervault *OrdervaultTransactor) SyncTokenBalance(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Ordervault.contract.Transact(opts, "syncTokenBalance", token)
}

// SyncTokenBalance is a paid mutator transaction binding the contract method 0xeb40133f.
//
// Solidity: function syncTokenBalance(address token) returns(uint256)
func (_Ordervault *OrdervaultSession) SyncTokenBalance(token common.Address) (*types.Transaction, error) {
	return _Ordervault.Contract.SyncTokenBalance(&_Ordervault.TransactOpts, token)
}

// SyncTokenBalance is a paid mutator transaction binding the contract method 0xeb40133f.
//
// Solidity: function syncTokenBalance(address token) returns(uint256)
func (_Ordervault *OrdervaultTransactorSession) SyncTokenBalance(token common.Address) (*types.Transaction, error) {
	return _Ordervault.Contract.SyncTokenBalance(&_Ordervault.TransactOpts, token)
}

// TransferOut is a paid mutator transaction binding the contract method 0x078d3b79.
//
// Solidity: function transferOut(address token, address receiver, uint256 amount) returns()
func (_Ordervault *OrdervaultTransactor) TransferOut(opts *bind.TransactOpts, token common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ordervault.contract.Transact(opts, "transferOut", token, receiver, amount)
}

// TransferOut is a paid mutator transaction binding the contract method 0x078d3b79.
//
// Solidity: function transferOut(address token, address receiver, uint256 amount) returns()
func (_Ordervault *OrdervaultSession) TransferOut(token common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ordervault.Contract.TransferOut(&_Ordervault.TransactOpts, token, receiver, amount)
}

// TransferOut is a paid mutator transaction binding the contract method 0x078d3b79.
//
// Solidity: function transferOut(address token, address receiver, uint256 amount) returns()
func (_Ordervault *OrdervaultTransactorSession) TransferOut(token common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ordervault.Contract.TransferOut(&_Ordervault.TransactOpts, token, receiver, amount)
}

// TransferOut0 is a paid mutator transaction binding the contract method 0x2fb12605.
//
// Solidity: function transferOut(address token, address receiver, uint256 amount, bool shouldUnwrapNativeToken) returns()
func (_Ordervault *OrdervaultTransactor) TransferOut0(opts *bind.TransactOpts, token common.Address, receiver common.Address, amount *big.Int, shouldUnwrapNativeToken bool) (*types.Transaction, error) {
	return _Ordervault.contract.Transact(opts, "transferOut0", token, receiver, amount, shouldUnwrapNativeToken)
}

// TransferOut0 is a paid mutator transaction binding the contract method 0x2fb12605.
//
// Solidity: function transferOut(address token, address receiver, uint256 amount, bool shouldUnwrapNativeToken) returns()
func (_Ordervault *OrdervaultSession) TransferOut0(token common.Address, receiver common.Address, amount *big.Int, shouldUnwrapNativeToken bool) (*types.Transaction, error) {
	return _Ordervault.Contract.TransferOut0(&_Ordervault.TransactOpts, token, receiver, amount, shouldUnwrapNativeToken)
}

// TransferOut0 is a paid mutator transaction binding the contract method 0x2fb12605.
//
// Solidity: function transferOut(address token, address receiver, uint256 amount, bool shouldUnwrapNativeToken) returns()
func (_Ordervault *OrdervaultTransactorSession) TransferOut0(token common.Address, receiver common.Address, amount *big.Int, shouldUnwrapNativeToken bool) (*types.Transaction, error) {
	return _Ordervault.Contract.TransferOut0(&_Ordervault.TransactOpts, token, receiver, amount, shouldUnwrapNativeToken)
}

// TransferOutNativeToken is a paid mutator transaction binding the contract method 0xd443ca94.
//
// Solidity: function transferOutNativeToken(address receiver, uint256 amount) returns()
func (_Ordervault *OrdervaultTransactor) TransferOutNativeToken(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ordervault.contract.Transact(opts, "transferOutNativeToken", receiver, amount)
}

// TransferOutNativeToken is a paid mutator transaction binding the contract method 0xd443ca94.
//
// Solidity: function transferOutNativeToken(address receiver, uint256 amount) returns()
func (_Ordervault *OrdervaultSession) TransferOutNativeToken(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ordervault.Contract.TransferOutNativeToken(&_Ordervault.TransactOpts, receiver, amount)
}

// TransferOutNativeToken is a paid mutator transaction binding the contract method 0xd443ca94.
//
// Solidity: function transferOutNativeToken(address receiver, uint256 amount) returns()
func (_Ordervault *OrdervaultTransactorSession) TransferOutNativeToken(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ordervault.Contract.TransferOutNativeToken(&_Ordervault.TransactOpts, receiver, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ordervault *OrdervaultTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ordervault.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ordervault *OrdervaultSession) Receive() (*types.Transaction, error) {
	return _Ordervault.Contract.Receive(&_Ordervault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ordervault *OrdervaultTransactorSession) Receive() (*types.Transaction, error) {
	return _Ordervault.Contract.Receive(&_Ordervault.TransactOpts)
}
