// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package syntheticsrouter

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

// SyntheticsrouterMetaData contains all meta data concerning the Syntheticsrouter contract.
var SyntheticsrouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractRoleStore\",\"name\":\"_roleStore\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"pluginTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roleStore\",\"outputs\":[{\"internalType\":\"contractRoleStore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SyntheticsrouterABI is the input ABI used to generate the binding from.
// Deprecated: Use SyntheticsrouterMetaData.ABI instead.
var SyntheticsrouterABI = SyntheticsrouterMetaData.ABI

// Syntheticsrouter is an auto generated Go binding around an Ethereum contract.
type Syntheticsrouter struct {
	SyntheticsrouterCaller     // Read-only binding to the contract
	SyntheticsrouterTransactor // Write-only binding to the contract
	SyntheticsrouterFilterer   // Log filterer for contract events
}

// SyntheticsrouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type SyntheticsrouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SyntheticsrouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SyntheticsrouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SyntheticsrouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SyntheticsrouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SyntheticsrouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SyntheticsrouterSession struct {
	Contract     *Syntheticsrouter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SyntheticsrouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SyntheticsrouterCallerSession struct {
	Contract *SyntheticsrouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// SyntheticsrouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SyntheticsrouterTransactorSession struct {
	Contract     *SyntheticsrouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SyntheticsrouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type SyntheticsrouterRaw struct {
	Contract *Syntheticsrouter // Generic contract binding to access the raw methods on
}

// SyntheticsrouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SyntheticsrouterCallerRaw struct {
	Contract *SyntheticsrouterCaller // Generic read-only contract binding to access the raw methods on
}

// SyntheticsrouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SyntheticsrouterTransactorRaw struct {
	Contract *SyntheticsrouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSyntheticsrouter creates a new instance of Syntheticsrouter, bound to a specific deployed contract.
func NewSyntheticsrouter(address common.Address, backend bind.ContractBackend) (*Syntheticsrouter, error) {
	contract, err := bindSyntheticsrouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Syntheticsrouter{SyntheticsrouterCaller: SyntheticsrouterCaller{contract: contract}, SyntheticsrouterTransactor: SyntheticsrouterTransactor{contract: contract}, SyntheticsrouterFilterer: SyntheticsrouterFilterer{contract: contract}}, nil
}

// NewSyntheticsrouterCaller creates a new read-only instance of Syntheticsrouter, bound to a specific deployed contract.
func NewSyntheticsrouterCaller(address common.Address, caller bind.ContractCaller) (*SyntheticsrouterCaller, error) {
	contract, err := bindSyntheticsrouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SyntheticsrouterCaller{contract: contract}, nil
}

// NewSyntheticsrouterTransactor creates a new write-only instance of Syntheticsrouter, bound to a specific deployed contract.
func NewSyntheticsrouterTransactor(address common.Address, transactor bind.ContractTransactor) (*SyntheticsrouterTransactor, error) {
	contract, err := bindSyntheticsrouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SyntheticsrouterTransactor{contract: contract}, nil
}

// NewSyntheticsrouterFilterer creates a new log filterer instance of Syntheticsrouter, bound to a specific deployed contract.
func NewSyntheticsrouterFilterer(address common.Address, filterer bind.ContractFilterer) (*SyntheticsrouterFilterer, error) {
	contract, err := bindSyntheticsrouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SyntheticsrouterFilterer{contract: contract}, nil
}

// bindSyntheticsrouter binds a generic wrapper to an already deployed contract.
func bindSyntheticsrouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SyntheticsrouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Syntheticsrouter *SyntheticsrouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Syntheticsrouter.Contract.SyntheticsrouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Syntheticsrouter *SyntheticsrouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Syntheticsrouter.Contract.SyntheticsrouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Syntheticsrouter *SyntheticsrouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Syntheticsrouter.Contract.SyntheticsrouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Syntheticsrouter *SyntheticsrouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Syntheticsrouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Syntheticsrouter *SyntheticsrouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Syntheticsrouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Syntheticsrouter *SyntheticsrouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Syntheticsrouter.Contract.contract.Transact(opts, method, params...)
}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_Syntheticsrouter *SyntheticsrouterCaller) RoleStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Syntheticsrouter.contract.Call(opts, &out, "roleStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_Syntheticsrouter *SyntheticsrouterSession) RoleStore() (common.Address, error) {
	return _Syntheticsrouter.Contract.RoleStore(&_Syntheticsrouter.CallOpts)
}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_Syntheticsrouter *SyntheticsrouterCallerSession) RoleStore() (common.Address, error) {
	return _Syntheticsrouter.Contract.RoleStore(&_Syntheticsrouter.CallOpts)
}

// PluginTransfer is a paid mutator transaction binding the contract method 0x1b827878.
//
// Solidity: function pluginTransfer(address token, address account, address receiver, uint256 amount) returns()
func (_Syntheticsrouter *SyntheticsrouterTransactor) PluginTransfer(opts *bind.TransactOpts, token common.Address, account common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Syntheticsrouter.contract.Transact(opts, "pluginTransfer", token, account, receiver, amount)
}

// PluginTransfer is a paid mutator transaction binding the contract method 0x1b827878.
//
// Solidity: function pluginTransfer(address token, address account, address receiver, uint256 amount) returns()
func (_Syntheticsrouter *SyntheticsrouterSession) PluginTransfer(token common.Address, account common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Syntheticsrouter.Contract.PluginTransfer(&_Syntheticsrouter.TransactOpts, token, account, receiver, amount)
}

// PluginTransfer is a paid mutator transaction binding the contract method 0x1b827878.
//
// Solidity: function pluginTransfer(address token, address account, address receiver, uint256 amount) returns()
func (_Syntheticsrouter *SyntheticsrouterTransactorSession) PluginTransfer(token common.Address, account common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Syntheticsrouter.Contract.PluginTransfer(&_Syntheticsrouter.TransactOpts, token, account, receiver, amount)
}
