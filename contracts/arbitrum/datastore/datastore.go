// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package datastore

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

// DatastoreMetaData contains all meta data concerning the Datastore contract.
var DatastoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractRoleStore\",\"name\":\"_roleStore\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"setKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"addAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"setKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"addBytes32\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"setKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"addUint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addressArrayValues\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"addressValues\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"applyBoundedDeltaToUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"applyDeltaToInt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"errorMessage\",\"type\":\"string\"}],\"name\":\"applyDeltaToUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"applyDeltaToUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"boolArrayValues\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"boolValues\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bytes32ArrayValues\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bytes32Values\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"setKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"containsAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"setKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"containsBytes32\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"setKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"containsUint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"decrementInt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"decrementUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getAddressArray\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"setKey\",\"type\":\"bytes32\"}],\"name\":\"getAddressCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"setKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"getAddressValuesAt\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getBool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getBoolArray\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getBytes32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getBytes32Array\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"setKey\",\"type\":\"bytes32\"}],\"name\":\"getBytes32Count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"setKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"getBytes32ValuesAt\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getInt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getIntArray\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"\",\"type\":\"int256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getStringArray\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getUintArray\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"setKey\",\"type\":\"bytes32\"}],\"name\":\"getUintCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"setKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"getUintValuesAt\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"incrementInt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"incrementUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"intArrayValues\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"intValues\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"setKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"removeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"removeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"removeAddressArray\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"removeBool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"removeBoolArray\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"setKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"removeBytes32\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"removeBytes32\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"removeBytes32Array\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"removeInt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"removeIntArray\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"removeString\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"removeStringArray\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"removeUint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"setKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"removeUint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"removeUintArray\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roleStore\",\"outputs\":[{\"internalType\":\"contractRoleStore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"value\",\"type\":\"address[]\"}],\"name\":\"setAddressArray\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setBool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bool[]\",\"name\":\"value\",\"type\":\"bool[]\"}],\"name\":\"setBoolArray\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"setBytes32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"value\",\"type\":\"bytes32[]\"}],\"name\":\"setBytes32Array\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"setInt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"int256[]\",\"name\":\"value\",\"type\":\"int256[]\"}],\"name\":\"setIntArray\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"string[]\",\"name\":\"value\",\"type\":\"string[]\"}],\"name\":\"setStringArray\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"}],\"name\":\"setUintArray\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stringArrayValues\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"stringValues\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uintArrayValues\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"uintValues\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DatastoreABI is the input ABI used to generate the binding from.
// Deprecated: Use DatastoreMetaData.ABI instead.
var DatastoreABI = DatastoreMetaData.ABI

// Datastore is an auto generated Go binding around an Ethereum contract.
type Datastore struct {
	DatastoreCaller     // Read-only binding to the contract
	DatastoreTransactor // Write-only binding to the contract
	DatastoreFilterer   // Log filterer for contract events
}

// DatastoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type DatastoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatastoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DatastoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatastoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DatastoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatastoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DatastoreSession struct {
	Contract     *Datastore        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DatastoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DatastoreCallerSession struct {
	Contract *DatastoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DatastoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DatastoreTransactorSession struct {
	Contract     *DatastoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DatastoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type DatastoreRaw struct {
	Contract *Datastore // Generic contract binding to access the raw methods on
}

// DatastoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DatastoreCallerRaw struct {
	Contract *DatastoreCaller // Generic read-only contract binding to access the raw methods on
}

// DatastoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DatastoreTransactorRaw struct {
	Contract *DatastoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDatastore creates a new instance of Datastore, bound to a specific deployed contract.
func NewDatastore(address common.Address, backend bind.ContractBackend) (*Datastore, error) {
	contract, err := bindDatastore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Datastore{DatastoreCaller: DatastoreCaller{contract: contract}, DatastoreTransactor: DatastoreTransactor{contract: contract}, DatastoreFilterer: DatastoreFilterer{contract: contract}}, nil
}

// NewDatastoreCaller creates a new read-only instance of Datastore, bound to a specific deployed contract.
func NewDatastoreCaller(address common.Address, caller bind.ContractCaller) (*DatastoreCaller, error) {
	contract, err := bindDatastore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DatastoreCaller{contract: contract}, nil
}

// NewDatastoreTransactor creates a new write-only instance of Datastore, bound to a specific deployed contract.
func NewDatastoreTransactor(address common.Address, transactor bind.ContractTransactor) (*DatastoreTransactor, error) {
	contract, err := bindDatastore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DatastoreTransactor{contract: contract}, nil
}

// NewDatastoreFilterer creates a new log filterer instance of Datastore, bound to a specific deployed contract.
func NewDatastoreFilterer(address common.Address, filterer bind.ContractFilterer) (*DatastoreFilterer, error) {
	contract, err := bindDatastore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DatastoreFilterer{contract: contract}, nil
}

// bindDatastore binds a generic wrapper to an already deployed contract.
func bindDatastore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DatastoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Datastore *DatastoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Datastore.Contract.DatastoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Datastore *DatastoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Datastore.Contract.DatastoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Datastore *DatastoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Datastore.Contract.DatastoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Datastore *DatastoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Datastore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Datastore *DatastoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Datastore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Datastore *DatastoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Datastore.Contract.contract.Transact(opts, method, params...)
}

// AddressArrayValues is a free data retrieval call binding the contract method 0x22f87464.
//
// Solidity: function addressArrayValues(bytes32 , uint256 ) view returns(address)
func (_Datastore *DatastoreCaller) AddressArrayValues(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "addressArrayValues", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressArrayValues is a free data retrieval call binding the contract method 0x22f87464.
//
// Solidity: function addressArrayValues(bytes32 , uint256 ) view returns(address)
func (_Datastore *DatastoreSession) AddressArrayValues(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _Datastore.Contract.AddressArrayValues(&_Datastore.CallOpts, arg0, arg1)
}

// AddressArrayValues is a free data retrieval call binding the contract method 0x22f87464.
//
// Solidity: function addressArrayValues(bytes32 , uint256 ) view returns(address)
func (_Datastore *DatastoreCallerSession) AddressArrayValues(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _Datastore.Contract.AddressArrayValues(&_Datastore.CallOpts, arg0, arg1)
}

// AddressValues is a free data retrieval call binding the contract method 0x22538dae.
//
// Solidity: function addressValues(bytes32 ) view returns(address)
func (_Datastore *DatastoreCaller) AddressValues(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "addressValues", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressValues is a free data retrieval call binding the contract method 0x22538dae.
//
// Solidity: function addressValues(bytes32 ) view returns(address)
func (_Datastore *DatastoreSession) AddressValues(arg0 [32]byte) (common.Address, error) {
	return _Datastore.Contract.AddressValues(&_Datastore.CallOpts, arg0)
}

// AddressValues is a free data retrieval call binding the contract method 0x22538dae.
//
// Solidity: function addressValues(bytes32 ) view returns(address)
func (_Datastore *DatastoreCallerSession) AddressValues(arg0 [32]byte) (common.Address, error) {
	return _Datastore.Contract.AddressValues(&_Datastore.CallOpts, arg0)
}

// BoolArrayValues is a free data retrieval call binding the contract method 0x80aacdcd.
//
// Solidity: function boolArrayValues(bytes32 , uint256 ) view returns(bool)
func (_Datastore *DatastoreCaller) BoolArrayValues(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "boolArrayValues", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BoolArrayValues is a free data retrieval call binding the contract method 0x80aacdcd.
//
// Solidity: function boolArrayValues(bytes32 , uint256 ) view returns(bool)
func (_Datastore *DatastoreSession) BoolArrayValues(arg0 [32]byte, arg1 *big.Int) (bool, error) {
	return _Datastore.Contract.BoolArrayValues(&_Datastore.CallOpts, arg0, arg1)
}

// BoolArrayValues is a free data retrieval call binding the contract method 0x80aacdcd.
//
// Solidity: function boolArrayValues(bytes32 , uint256 ) view returns(bool)
func (_Datastore *DatastoreCallerSession) BoolArrayValues(arg0 [32]byte, arg1 *big.Int) (bool, error) {
	return _Datastore.Contract.BoolArrayValues(&_Datastore.CallOpts, arg0, arg1)
}

// BoolValues is a free data retrieval call binding the contract method 0x44a242b1.
//
// Solidity: function boolValues(bytes32 ) view returns(bool)
func (_Datastore *DatastoreCaller) BoolValues(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "boolValues", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BoolValues is a free data retrieval call binding the contract method 0x44a242b1.
//
// Solidity: function boolValues(bytes32 ) view returns(bool)
func (_Datastore *DatastoreSession) BoolValues(arg0 [32]byte) (bool, error) {
	return _Datastore.Contract.BoolValues(&_Datastore.CallOpts, arg0)
}

// BoolValues is a free data retrieval call binding the contract method 0x44a242b1.
//
// Solidity: function boolValues(bytes32 ) view returns(bool)
func (_Datastore *DatastoreCallerSession) BoolValues(arg0 [32]byte) (bool, error) {
	return _Datastore.Contract.BoolValues(&_Datastore.CallOpts, arg0)
}

// Bytes32ArrayValues is a free data retrieval call binding the contract method 0xbf498dd3.
//
// Solidity: function bytes32ArrayValues(bytes32 , uint256 ) view returns(bytes32)
func (_Datastore *DatastoreCaller) Bytes32ArrayValues(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "bytes32ArrayValues", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Bytes32ArrayValues is a free data retrieval call binding the contract method 0xbf498dd3.
//
// Solidity: function bytes32ArrayValues(bytes32 , uint256 ) view returns(bytes32)
func (_Datastore *DatastoreSession) Bytes32ArrayValues(arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	return _Datastore.Contract.Bytes32ArrayValues(&_Datastore.CallOpts, arg0, arg1)
}

// Bytes32ArrayValues is a free data retrieval call binding the contract method 0xbf498dd3.
//
// Solidity: function bytes32ArrayValues(bytes32 , uint256 ) view returns(bytes32)
func (_Datastore *DatastoreCallerSession) Bytes32ArrayValues(arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	return _Datastore.Contract.Bytes32ArrayValues(&_Datastore.CallOpts, arg0, arg1)
}

// Bytes32Values is a free data retrieval call binding the contract method 0xd52852af.
//
// Solidity: function bytes32Values(bytes32 ) view returns(bytes32)
func (_Datastore *DatastoreCaller) Bytes32Values(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "bytes32Values", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Bytes32Values is a free data retrieval call binding the contract method 0xd52852af.
//
// Solidity: function bytes32Values(bytes32 ) view returns(bytes32)
func (_Datastore *DatastoreSession) Bytes32Values(arg0 [32]byte) ([32]byte, error) {
	return _Datastore.Contract.Bytes32Values(&_Datastore.CallOpts, arg0)
}

// Bytes32Values is a free data retrieval call binding the contract method 0xd52852af.
//
// Solidity: function bytes32Values(bytes32 ) view returns(bytes32)
func (_Datastore *DatastoreCallerSession) Bytes32Values(arg0 [32]byte) ([32]byte, error) {
	return _Datastore.Contract.Bytes32Values(&_Datastore.CallOpts, arg0)
}

// ContainsAddress is a free data retrieval call binding the contract method 0xc769d1a1.
//
// Solidity: function containsAddress(bytes32 setKey, address value) view returns(bool)
func (_Datastore *DatastoreCaller) ContainsAddress(opts *bind.CallOpts, setKey [32]byte, value common.Address) (bool, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "containsAddress", setKey, value)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContainsAddress is a free data retrieval call binding the contract method 0xc769d1a1.
//
// Solidity: function containsAddress(bytes32 setKey, address value) view returns(bool)
func (_Datastore *DatastoreSession) ContainsAddress(setKey [32]byte, value common.Address) (bool, error) {
	return _Datastore.Contract.ContainsAddress(&_Datastore.CallOpts, setKey, value)
}

// ContainsAddress is a free data retrieval call binding the contract method 0xc769d1a1.
//
// Solidity: function containsAddress(bytes32 setKey, address value) view returns(bool)
func (_Datastore *DatastoreCallerSession) ContainsAddress(setKey [32]byte, value common.Address) (bool, error) {
	return _Datastore.Contract.ContainsAddress(&_Datastore.CallOpts, setKey, value)
}

// ContainsBytes32 is a free data retrieval call binding the contract method 0x91d4403c.
//
// Solidity: function containsBytes32(bytes32 setKey, bytes32 value) view returns(bool)
func (_Datastore *DatastoreCaller) ContainsBytes32(opts *bind.CallOpts, setKey [32]byte, value [32]byte) (bool, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "containsBytes32", setKey, value)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContainsBytes32 is a free data retrieval call binding the contract method 0x91d4403c.
//
// Solidity: function containsBytes32(bytes32 setKey, bytes32 value) view returns(bool)
func (_Datastore *DatastoreSession) ContainsBytes32(setKey [32]byte, value [32]byte) (bool, error) {
	return _Datastore.Contract.ContainsBytes32(&_Datastore.CallOpts, setKey, value)
}

// ContainsBytes32 is a free data retrieval call binding the contract method 0x91d4403c.
//
// Solidity: function containsBytes32(bytes32 setKey, bytes32 value) view returns(bool)
func (_Datastore *DatastoreCallerSession) ContainsBytes32(setKey [32]byte, value [32]byte) (bool, error) {
	return _Datastore.Contract.ContainsBytes32(&_Datastore.CallOpts, setKey, value)
}

// ContainsUint is a free data retrieval call binding the contract method 0x310b8882.
//
// Solidity: function containsUint(bytes32 setKey, uint256 value) view returns(bool)
func (_Datastore *DatastoreCaller) ContainsUint(opts *bind.CallOpts, setKey [32]byte, value *big.Int) (bool, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "containsUint", setKey, value)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContainsUint is a free data retrieval call binding the contract method 0x310b8882.
//
// Solidity: function containsUint(bytes32 setKey, uint256 value) view returns(bool)
func (_Datastore *DatastoreSession) ContainsUint(setKey [32]byte, value *big.Int) (bool, error) {
	return _Datastore.Contract.ContainsUint(&_Datastore.CallOpts, setKey, value)
}

// ContainsUint is a free data retrieval call binding the contract method 0x310b8882.
//
// Solidity: function containsUint(bytes32 setKey, uint256 value) view returns(bool)
func (_Datastore *DatastoreCallerSession) ContainsUint(setKey [32]byte, value *big.Int) (bool, error) {
	return _Datastore.Contract.ContainsUint(&_Datastore.CallOpts, setKey, value)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 key) view returns(address)
func (_Datastore *DatastoreCaller) GetAddress(opts *bind.CallOpts, key [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "getAddress", key)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 key) view returns(address)
func (_Datastore *DatastoreSession) GetAddress(key [32]byte) (common.Address, error) {
	return _Datastore.Contract.GetAddress(&_Datastore.CallOpts, key)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 key) view returns(address)
func (_Datastore *DatastoreCallerSession) GetAddress(key [32]byte) (common.Address, error) {
	return _Datastore.Contract.GetAddress(&_Datastore.CallOpts, key)
}

// GetAddressArray is a free data retrieval call binding the contract method 0x5948f733.
//
// Solidity: function getAddressArray(bytes32 key) view returns(address[])
func (_Datastore *DatastoreCaller) GetAddressArray(opts *bind.CallOpts, key [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "getAddressArray", key)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAddressArray is a free data retrieval call binding the contract method 0x5948f733.
//
// Solidity: function getAddressArray(bytes32 key) view returns(address[])
func (_Datastore *DatastoreSession) GetAddressArray(key [32]byte) ([]common.Address, error) {
	return _Datastore.Contract.GetAddressArray(&_Datastore.CallOpts, key)
}

// GetAddressArray is a free data retrieval call binding the contract method 0x5948f733.
//
// Solidity: function getAddressArray(bytes32 key) view returns(address[])
func (_Datastore *DatastoreCallerSession) GetAddressArray(key [32]byte) ([]common.Address, error) {
	return _Datastore.Contract.GetAddressArray(&_Datastore.CallOpts, key)
}

// GetAddressCount is a free data retrieval call binding the contract method 0x35ea8059.
//
// Solidity: function getAddressCount(bytes32 setKey) view returns(uint256)
func (_Datastore *DatastoreCaller) GetAddressCount(opts *bind.CallOpts, setKey [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "getAddressCount", setKey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAddressCount is a free data retrieval call binding the contract method 0x35ea8059.
//
// Solidity: function getAddressCount(bytes32 setKey) view returns(uint256)
func (_Datastore *DatastoreSession) GetAddressCount(setKey [32]byte) (*big.Int, error) {
	return _Datastore.Contract.GetAddressCount(&_Datastore.CallOpts, setKey)
}

// GetAddressCount is a free data retrieval call binding the contract method 0x35ea8059.
//
// Solidity: function getAddressCount(bytes32 setKey) view returns(uint256)
func (_Datastore *DatastoreCallerSession) GetAddressCount(setKey [32]byte) (*big.Int, error) {
	return _Datastore.Contract.GetAddressCount(&_Datastore.CallOpts, setKey)
}

// GetAddressValuesAt is a free data retrieval call binding the contract method 0xe7e4148e.
//
// Solidity: function getAddressValuesAt(bytes32 setKey, uint256 start, uint256 end) view returns(address[])
func (_Datastore *DatastoreCaller) GetAddressValuesAt(opts *bind.CallOpts, setKey [32]byte, start *big.Int, end *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "getAddressValuesAt", setKey, start, end)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAddressValuesAt is a free data retrieval call binding the contract method 0xe7e4148e.
//
// Solidity: function getAddressValuesAt(bytes32 setKey, uint256 start, uint256 end) view returns(address[])
func (_Datastore *DatastoreSession) GetAddressValuesAt(setKey [32]byte, start *big.Int, end *big.Int) ([]common.Address, error) {
	return _Datastore.Contract.GetAddressValuesAt(&_Datastore.CallOpts, setKey, start, end)
}

// GetAddressValuesAt is a free data retrieval call binding the contract method 0xe7e4148e.
//
// Solidity: function getAddressValuesAt(bytes32 setKey, uint256 start, uint256 end) view returns(address[])
func (_Datastore *DatastoreCallerSession) GetAddressValuesAt(setKey [32]byte, start *big.Int, end *big.Int) ([]common.Address, error) {
	return _Datastore.Contract.GetAddressValuesAt(&_Datastore.CallOpts, setKey, start, end)
}

// GetBool is a free data retrieval call binding the contract method 0x7ae1cfca.
//
// Solidity: function getBool(bytes32 key) view returns(bool)
func (_Datastore *DatastoreCaller) GetBool(opts *bind.CallOpts, key [32]byte) (bool, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "getBool", key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetBool is a free data retrieval call binding the contract method 0x7ae1cfca.
//
// Solidity: function getBool(bytes32 key) view returns(bool)
func (_Datastore *DatastoreSession) GetBool(key [32]byte) (bool, error) {
	return _Datastore.Contract.GetBool(&_Datastore.CallOpts, key)
}

// GetBool is a free data retrieval call binding the contract method 0x7ae1cfca.
//
// Solidity: function getBool(bytes32 key) view returns(bool)
func (_Datastore *DatastoreCallerSession) GetBool(key [32]byte) (bool, error) {
	return _Datastore.Contract.GetBool(&_Datastore.CallOpts, key)
}

// GetBoolArray is a free data retrieval call binding the contract method 0x116bb929.
//
// Solidity: function getBoolArray(bytes32 key) view returns(bool[])
func (_Datastore *DatastoreCaller) GetBoolArray(opts *bind.CallOpts, key [32]byte) ([]bool, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "getBoolArray", key)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// GetBoolArray is a free data retrieval call binding the contract method 0x116bb929.
//
// Solidity: function getBoolArray(bytes32 key) view returns(bool[])
func (_Datastore *DatastoreSession) GetBoolArray(key [32]byte) ([]bool, error) {
	return _Datastore.Contract.GetBoolArray(&_Datastore.CallOpts, key)
}

// GetBoolArray is a free data retrieval call binding the contract method 0x116bb929.
//
// Solidity: function getBoolArray(bytes32 key) view returns(bool[])
func (_Datastore *DatastoreCallerSession) GetBoolArray(key [32]byte) ([]bool, error) {
	return _Datastore.Contract.GetBoolArray(&_Datastore.CallOpts, key)
}

// GetBytes32 is a free data retrieval call binding the contract method 0xa6ed563e.
//
// Solidity: function getBytes32(bytes32 key) view returns(bytes32)
func (_Datastore *DatastoreCaller) GetBytes32(opts *bind.CallOpts, key [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "getBytes32", key)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBytes32 is a free data retrieval call binding the contract method 0xa6ed563e.
//
// Solidity: function getBytes32(bytes32 key) view returns(bytes32)
func (_Datastore *DatastoreSession) GetBytes32(key [32]byte) ([32]byte, error) {
	return _Datastore.Contract.GetBytes32(&_Datastore.CallOpts, key)
}

// GetBytes32 is a free data retrieval call binding the contract method 0xa6ed563e.
//
// Solidity: function getBytes32(bytes32 key) view returns(bytes32)
func (_Datastore *DatastoreCallerSession) GetBytes32(key [32]byte) ([32]byte, error) {
	return _Datastore.Contract.GetBytes32(&_Datastore.CallOpts, key)
}

// GetBytes32Array is a free data retrieval call binding the contract method 0xdd031997.
//
// Solidity: function getBytes32Array(bytes32 key) view returns(bytes32[])
func (_Datastore *DatastoreCaller) GetBytes32Array(opts *bind.CallOpts, key [32]byte) ([][32]byte, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "getBytes32Array", key)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBytes32Array is a free data retrieval call binding the contract method 0xdd031997.
//
// Solidity: function getBytes32Array(bytes32 key) view returns(bytes32[])
func (_Datastore *DatastoreSession) GetBytes32Array(key [32]byte) ([][32]byte, error) {
	return _Datastore.Contract.GetBytes32Array(&_Datastore.CallOpts, key)
}

// GetBytes32Array is a free data retrieval call binding the contract method 0xdd031997.
//
// Solidity: function getBytes32Array(bytes32 key) view returns(bytes32[])
func (_Datastore *DatastoreCallerSession) GetBytes32Array(key [32]byte) ([][32]byte, error) {
	return _Datastore.Contract.GetBytes32Array(&_Datastore.CallOpts, key)
}

// GetBytes32Count is a free data retrieval call binding the contract method 0xf3903b9f.
//
// Solidity: function getBytes32Count(bytes32 setKey) view returns(uint256)
func (_Datastore *DatastoreCaller) GetBytes32Count(opts *bind.CallOpts, setKey [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "getBytes32Count", setKey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBytes32Count is a free data retrieval call binding the contract method 0xf3903b9f.
//
// Solidity: function getBytes32Count(bytes32 setKey) view returns(uint256)
func (_Datastore *DatastoreSession) GetBytes32Count(setKey [32]byte) (*big.Int, error) {
	return _Datastore.Contract.GetBytes32Count(&_Datastore.CallOpts, setKey)
}

// GetBytes32Count is a free data retrieval call binding the contract method 0xf3903b9f.
//
// Solidity: function getBytes32Count(bytes32 setKey) view returns(uint256)
func (_Datastore *DatastoreCallerSession) GetBytes32Count(setKey [32]byte) (*big.Int, error) {
	return _Datastore.Contract.GetBytes32Count(&_Datastore.CallOpts, setKey)
}

// GetBytes32ValuesAt is a free data retrieval call binding the contract method 0xf069052a.
//
// Solidity: function getBytes32ValuesAt(bytes32 setKey, uint256 start, uint256 end) view returns(bytes32[])
func (_Datastore *DatastoreCaller) GetBytes32ValuesAt(opts *bind.CallOpts, setKey [32]byte, start *big.Int, end *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "getBytes32ValuesAt", setKey, start, end)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBytes32ValuesAt is a free data retrieval call binding the contract method 0xf069052a.
//
// Solidity: function getBytes32ValuesAt(bytes32 setKey, uint256 start, uint256 end) view returns(bytes32[])
func (_Datastore *DatastoreSession) GetBytes32ValuesAt(setKey [32]byte, start *big.Int, end *big.Int) ([][32]byte, error) {
	return _Datastore.Contract.GetBytes32ValuesAt(&_Datastore.CallOpts, setKey, start, end)
}

// GetBytes32ValuesAt is a free data retrieval call binding the contract method 0xf069052a.
//
// Solidity: function getBytes32ValuesAt(bytes32 setKey, uint256 start, uint256 end) view returns(bytes32[])
func (_Datastore *DatastoreCallerSession) GetBytes32ValuesAt(setKey [32]byte, start *big.Int, end *big.Int) ([][32]byte, error) {
	return _Datastore.Contract.GetBytes32ValuesAt(&_Datastore.CallOpts, setKey, start, end)
}

// GetInt is a free data retrieval call binding the contract method 0xdc97d962.
//
// Solidity: function getInt(bytes32 key) view returns(int256)
func (_Datastore *DatastoreCaller) GetInt(opts *bind.CallOpts, key [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "getInt", key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInt is a free data retrieval call binding the contract method 0xdc97d962.
//
// Solidity: function getInt(bytes32 key) view returns(int256)
func (_Datastore *DatastoreSession) GetInt(key [32]byte) (*big.Int, error) {
	return _Datastore.Contract.GetInt(&_Datastore.CallOpts, key)
}

// GetInt is a free data retrieval call binding the contract method 0xdc97d962.
//
// Solidity: function getInt(bytes32 key) view returns(int256)
func (_Datastore *DatastoreCallerSession) GetInt(key [32]byte) (*big.Int, error) {
	return _Datastore.Contract.GetInt(&_Datastore.CallOpts, key)
}

// GetIntArray is a free data retrieval call binding the contract method 0x2d2899b6.
//
// Solidity: function getIntArray(bytes32 key) view returns(int256[])
func (_Datastore *DatastoreCaller) GetIntArray(opts *bind.CallOpts, key [32]byte) ([]*big.Int, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "getIntArray", key)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetIntArray is a free data retrieval call binding the contract method 0x2d2899b6.
//
// Solidity: function getIntArray(bytes32 key) view returns(int256[])
func (_Datastore *DatastoreSession) GetIntArray(key [32]byte) ([]*big.Int, error) {
	return _Datastore.Contract.GetIntArray(&_Datastore.CallOpts, key)
}

// GetIntArray is a free data retrieval call binding the contract method 0x2d2899b6.
//
// Solidity: function getIntArray(bytes32 key) view returns(int256[])
func (_Datastore *DatastoreCallerSession) GetIntArray(key [32]byte) ([]*big.Int, error) {
	return _Datastore.Contract.GetIntArray(&_Datastore.CallOpts, key)
}

// GetString is a free data retrieval call binding the contract method 0x986e791a.
//
// Solidity: function getString(bytes32 key) view returns(string)
func (_Datastore *DatastoreCaller) GetString(opts *bind.CallOpts, key [32]byte) (string, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "getString", key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetString is a free data retrieval call binding the contract method 0x986e791a.
//
// Solidity: function getString(bytes32 key) view returns(string)
func (_Datastore *DatastoreSession) GetString(key [32]byte) (string, error) {
	return _Datastore.Contract.GetString(&_Datastore.CallOpts, key)
}

// GetString is a free data retrieval call binding the contract method 0x986e791a.
//
// Solidity: function getString(bytes32 key) view returns(string)
func (_Datastore *DatastoreCallerSession) GetString(key [32]byte) (string, error) {
	return _Datastore.Contract.GetString(&_Datastore.CallOpts, key)
}

// GetStringArray is a free data retrieval call binding the contract method 0x01677da2.
//
// Solidity: function getStringArray(bytes32 key) view returns(string[])
func (_Datastore *DatastoreCaller) GetStringArray(opts *bind.CallOpts, key [32]byte) ([]string, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "getStringArray", key)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetStringArray is a free data retrieval call binding the contract method 0x01677da2.
//
// Solidity: function getStringArray(bytes32 key) view returns(string[])
func (_Datastore *DatastoreSession) GetStringArray(key [32]byte) ([]string, error) {
	return _Datastore.Contract.GetStringArray(&_Datastore.CallOpts, key)
}

// GetStringArray is a free data retrieval call binding the contract method 0x01677da2.
//
// Solidity: function getStringArray(bytes32 key) view returns(string[])
func (_Datastore *DatastoreCallerSession) GetStringArray(key [32]byte) ([]string, error) {
	return _Datastore.Contract.GetStringArray(&_Datastore.CallOpts, key)
}

// GetUint is a free data retrieval call binding the contract method 0xbd02d0f5.
//
// Solidity: function getUint(bytes32 key) view returns(uint256)
func (_Datastore *DatastoreCaller) GetUint(opts *bind.CallOpts, key [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "getUint", key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUint is a free data retrieval call binding the contract method 0xbd02d0f5.
//
// Solidity: function getUint(bytes32 key) view returns(uint256)
func (_Datastore *DatastoreSession) GetUint(key [32]byte) (*big.Int, error) {
	return _Datastore.Contract.GetUint(&_Datastore.CallOpts, key)
}

// GetUint is a free data retrieval call binding the contract method 0xbd02d0f5.
//
// Solidity: function getUint(bytes32 key) view returns(uint256)
func (_Datastore *DatastoreCallerSession) GetUint(key [32]byte) (*big.Int, error) {
	return _Datastore.Contract.GetUint(&_Datastore.CallOpts, key)
}

// GetUintArray is a free data retrieval call binding the contract method 0x86ac6bdf.
//
// Solidity: function getUintArray(bytes32 key) view returns(uint256[])
func (_Datastore *DatastoreCaller) GetUintArray(opts *bind.CallOpts, key [32]byte) ([]*big.Int, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "getUintArray", key)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUintArray is a free data retrieval call binding the contract method 0x86ac6bdf.
//
// Solidity: function getUintArray(bytes32 key) view returns(uint256[])
func (_Datastore *DatastoreSession) GetUintArray(key [32]byte) ([]*big.Int, error) {
	return _Datastore.Contract.GetUintArray(&_Datastore.CallOpts, key)
}

// GetUintArray is a free data retrieval call binding the contract method 0x86ac6bdf.
//
// Solidity: function getUintArray(bytes32 key) view returns(uint256[])
func (_Datastore *DatastoreCallerSession) GetUintArray(key [32]byte) ([]*big.Int, error) {
	return _Datastore.Contract.GetUintArray(&_Datastore.CallOpts, key)
}

// GetUintCount is a free data retrieval call binding the contract method 0x065f21a7.
//
// Solidity: function getUintCount(bytes32 setKey) view returns(uint256)
func (_Datastore *DatastoreCaller) GetUintCount(opts *bind.CallOpts, setKey [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "getUintCount", setKey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUintCount is a free data retrieval call binding the contract method 0x065f21a7.
//
// Solidity: function getUintCount(bytes32 setKey) view returns(uint256)
func (_Datastore *DatastoreSession) GetUintCount(setKey [32]byte) (*big.Int, error) {
	return _Datastore.Contract.GetUintCount(&_Datastore.CallOpts, setKey)
}

// GetUintCount is a free data retrieval call binding the contract method 0x065f21a7.
//
// Solidity: function getUintCount(bytes32 setKey) view returns(uint256)
func (_Datastore *DatastoreCallerSession) GetUintCount(setKey [32]byte) (*big.Int, error) {
	return _Datastore.Contract.GetUintCount(&_Datastore.CallOpts, setKey)
}

// GetUintValuesAt is a free data retrieval call binding the contract method 0x7026d42c.
//
// Solidity: function getUintValuesAt(bytes32 setKey, uint256 start, uint256 end) view returns(uint256[])
func (_Datastore *DatastoreCaller) GetUintValuesAt(opts *bind.CallOpts, setKey [32]byte, start *big.Int, end *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "getUintValuesAt", setKey, start, end)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUintValuesAt is a free data retrieval call binding the contract method 0x7026d42c.
//
// Solidity: function getUintValuesAt(bytes32 setKey, uint256 start, uint256 end) view returns(uint256[])
func (_Datastore *DatastoreSession) GetUintValuesAt(setKey [32]byte, start *big.Int, end *big.Int) ([]*big.Int, error) {
	return _Datastore.Contract.GetUintValuesAt(&_Datastore.CallOpts, setKey, start, end)
}

// GetUintValuesAt is a free data retrieval call binding the contract method 0x7026d42c.
//
// Solidity: function getUintValuesAt(bytes32 setKey, uint256 start, uint256 end) view returns(uint256[])
func (_Datastore *DatastoreCallerSession) GetUintValuesAt(setKey [32]byte, start *big.Int, end *big.Int) ([]*big.Int, error) {
	return _Datastore.Contract.GetUintValuesAt(&_Datastore.CallOpts, setKey, start, end)
}

// IntArrayValues is a free data retrieval call binding the contract method 0x6339734d.
//
// Solidity: function intArrayValues(bytes32 , uint256 ) view returns(int256)
func (_Datastore *DatastoreCaller) IntArrayValues(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "intArrayValues", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntArrayValues is a free data retrieval call binding the contract method 0x6339734d.
//
// Solidity: function intArrayValues(bytes32 , uint256 ) view returns(int256)
func (_Datastore *DatastoreSession) IntArrayValues(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _Datastore.Contract.IntArrayValues(&_Datastore.CallOpts, arg0, arg1)
}

// IntArrayValues is a free data retrieval call binding the contract method 0x6339734d.
//
// Solidity: function intArrayValues(bytes32 , uint256 ) view returns(int256)
func (_Datastore *DatastoreCallerSession) IntArrayValues(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _Datastore.Contract.IntArrayValues(&_Datastore.CallOpts, arg0, arg1)
}

// IntValues is a free data retrieval call binding the contract method 0x743df325.
//
// Solidity: function intValues(bytes32 ) view returns(int256)
func (_Datastore *DatastoreCaller) IntValues(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "intValues", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntValues is a free data retrieval call binding the contract method 0x743df325.
//
// Solidity: function intValues(bytes32 ) view returns(int256)
func (_Datastore *DatastoreSession) IntValues(arg0 [32]byte) (*big.Int, error) {
	return _Datastore.Contract.IntValues(&_Datastore.CallOpts, arg0)
}

// IntValues is a free data retrieval call binding the contract method 0x743df325.
//
// Solidity: function intValues(bytes32 ) view returns(int256)
func (_Datastore *DatastoreCallerSession) IntValues(arg0 [32]byte) (*big.Int, error) {
	return _Datastore.Contract.IntValues(&_Datastore.CallOpts, arg0)
}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_Datastore *DatastoreCaller) RoleStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "roleStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_Datastore *DatastoreSession) RoleStore() (common.Address, error) {
	return _Datastore.Contract.RoleStore(&_Datastore.CallOpts)
}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_Datastore *DatastoreCallerSession) RoleStore() (common.Address, error) {
	return _Datastore.Contract.RoleStore(&_Datastore.CallOpts)
}

// StringArrayValues is a free data retrieval call binding the contract method 0xb8320a08.
//
// Solidity: function stringArrayValues(bytes32 , uint256 ) view returns(string)
func (_Datastore *DatastoreCaller) StringArrayValues(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "stringArrayValues", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// StringArrayValues is a free data retrieval call binding the contract method 0xb8320a08.
//
// Solidity: function stringArrayValues(bytes32 , uint256 ) view returns(string)
func (_Datastore *DatastoreSession) StringArrayValues(arg0 [32]byte, arg1 *big.Int) (string, error) {
	return _Datastore.Contract.StringArrayValues(&_Datastore.CallOpts, arg0, arg1)
}

// StringArrayValues is a free data retrieval call binding the contract method 0xb8320a08.
//
// Solidity: function stringArrayValues(bytes32 , uint256 ) view returns(string)
func (_Datastore *DatastoreCallerSession) StringArrayValues(arg0 [32]byte, arg1 *big.Int) (string, error) {
	return _Datastore.Contract.StringArrayValues(&_Datastore.CallOpts, arg0, arg1)
}

// StringValues is a free data retrieval call binding the contract method 0xf15caeac.
//
// Solidity: function stringValues(bytes32 ) view returns(string)
func (_Datastore *DatastoreCaller) StringValues(opts *bind.CallOpts, arg0 [32]byte) (string, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "stringValues", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// StringValues is a free data retrieval call binding the contract method 0xf15caeac.
//
// Solidity: function stringValues(bytes32 ) view returns(string)
func (_Datastore *DatastoreSession) StringValues(arg0 [32]byte) (string, error) {
	return _Datastore.Contract.StringValues(&_Datastore.CallOpts, arg0)
}

// StringValues is a free data retrieval call binding the contract method 0xf15caeac.
//
// Solidity: function stringValues(bytes32 ) view returns(string)
func (_Datastore *DatastoreCallerSession) StringValues(arg0 [32]byte) (string, error) {
	return _Datastore.Contract.StringValues(&_Datastore.CallOpts, arg0)
}

// UintArrayValues is a free data retrieval call binding the contract method 0xc4f00fde.
//
// Solidity: function uintArrayValues(bytes32 , uint256 ) view returns(uint256)
func (_Datastore *DatastoreCaller) UintArrayValues(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "uintArrayValues", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UintArrayValues is a free data retrieval call binding the contract method 0xc4f00fde.
//
// Solidity: function uintArrayValues(bytes32 , uint256 ) view returns(uint256)
func (_Datastore *DatastoreSession) UintArrayValues(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _Datastore.Contract.UintArrayValues(&_Datastore.CallOpts, arg0, arg1)
}

// UintArrayValues is a free data retrieval call binding the contract method 0xc4f00fde.
//
// Solidity: function uintArrayValues(bytes32 , uint256 ) view returns(uint256)
func (_Datastore *DatastoreCallerSession) UintArrayValues(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _Datastore.Contract.UintArrayValues(&_Datastore.CallOpts, arg0, arg1)
}

// UintValues is a free data retrieval call binding the contract method 0xd38eebc7.
//
// Solidity: function uintValues(bytes32 ) view returns(uint256)
func (_Datastore *DatastoreCaller) UintValues(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "uintValues", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UintValues is a free data retrieval call binding the contract method 0xd38eebc7.
//
// Solidity: function uintValues(bytes32 ) view returns(uint256)
func (_Datastore *DatastoreSession) UintValues(arg0 [32]byte) (*big.Int, error) {
	return _Datastore.Contract.UintValues(&_Datastore.CallOpts, arg0)
}

// UintValues is a free data retrieval call binding the contract method 0xd38eebc7.
//
// Solidity: function uintValues(bytes32 ) view returns(uint256)
func (_Datastore *DatastoreCallerSession) UintValues(arg0 [32]byte) (*big.Int, error) {
	return _Datastore.Contract.UintValues(&_Datastore.CallOpts, arg0)
}

// AddAddress is a paid mutator transaction binding the contract method 0xb348e639.
//
// Solidity: function addAddress(bytes32 setKey, address value) returns()
func (_Datastore *DatastoreTransactor) AddAddress(opts *bind.TransactOpts, setKey [32]byte, value common.Address) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "addAddress", setKey, value)
}

// AddAddress is a paid mutator transaction binding the contract method 0xb348e639.
//
// Solidity: function addAddress(bytes32 setKey, address value) returns()
func (_Datastore *DatastoreSession) AddAddress(setKey [32]byte, value common.Address) (*types.Transaction, error) {
	return _Datastore.Contract.AddAddress(&_Datastore.TransactOpts, setKey, value)
}

// AddAddress is a paid mutator transaction binding the contract method 0xb348e639.
//
// Solidity: function addAddress(bytes32 setKey, address value) returns()
func (_Datastore *DatastoreTransactorSession) AddAddress(setKey [32]byte, value common.Address) (*types.Transaction, error) {
	return _Datastore.Contract.AddAddress(&_Datastore.TransactOpts, setKey, value)
}

// AddBytes32 is a paid mutator transaction binding the contract method 0xc80f4c62.
//
// Solidity: function addBytes32(bytes32 setKey, bytes32 value) returns()
func (_Datastore *DatastoreTransactor) AddBytes32(opts *bind.TransactOpts, setKey [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "addBytes32", setKey, value)
}

// AddBytes32 is a paid mutator transaction binding the contract method 0xc80f4c62.
//
// Solidity: function addBytes32(bytes32 setKey, bytes32 value) returns()
func (_Datastore *DatastoreSession) AddBytes32(setKey [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.AddBytes32(&_Datastore.TransactOpts, setKey, value)
}

// AddBytes32 is a paid mutator transaction binding the contract method 0xc80f4c62.
//
// Solidity: function addBytes32(bytes32 setKey, bytes32 value) returns()
func (_Datastore *DatastoreTransactorSession) AddBytes32(setKey [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.AddBytes32(&_Datastore.TransactOpts, setKey, value)
}

// AddUint is a paid mutator transaction binding the contract method 0xadb353dc.
//
// Solidity: function addUint(bytes32 setKey, uint256 value) returns()
func (_Datastore *DatastoreTransactor) AddUint(opts *bind.TransactOpts, setKey [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "addUint", setKey, value)
}

// AddUint is a paid mutator transaction binding the contract method 0xadb353dc.
//
// Solidity: function addUint(bytes32 setKey, uint256 value) returns()
func (_Datastore *DatastoreSession) AddUint(setKey [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.Contract.AddUint(&_Datastore.TransactOpts, setKey, value)
}

// AddUint is a paid mutator transaction binding the contract method 0xadb353dc.
//
// Solidity: function addUint(bytes32 setKey, uint256 value) returns()
func (_Datastore *DatastoreTransactorSession) AddUint(setKey [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.Contract.AddUint(&_Datastore.TransactOpts, setKey, value)
}

// ApplyBoundedDeltaToUint is a paid mutator transaction binding the contract method 0x8ca498b0.
//
// Solidity: function applyBoundedDeltaToUint(bytes32 key, int256 value) returns(uint256)
func (_Datastore *DatastoreTransactor) ApplyBoundedDeltaToUint(opts *bind.TransactOpts, key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "applyBoundedDeltaToUint", key, value)
}

// ApplyBoundedDeltaToUint is a paid mutator transaction binding the contract method 0x8ca498b0.
//
// Solidity: function applyBoundedDeltaToUint(bytes32 key, int256 value) returns(uint256)
func (_Datastore *DatastoreSession) ApplyBoundedDeltaToUint(key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.Contract.ApplyBoundedDeltaToUint(&_Datastore.TransactOpts, key, value)
}

// ApplyBoundedDeltaToUint is a paid mutator transaction binding the contract method 0x8ca498b0.
//
// Solidity: function applyBoundedDeltaToUint(bytes32 key, int256 value) returns(uint256)
func (_Datastore *DatastoreTransactorSession) ApplyBoundedDeltaToUint(key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.Contract.ApplyBoundedDeltaToUint(&_Datastore.TransactOpts, key, value)
}

// ApplyDeltaToInt is a paid mutator transaction binding the contract method 0xe4e36c4e.
//
// Solidity: function applyDeltaToInt(bytes32 key, int256 value) returns(int256)
func (_Datastore *DatastoreTransactor) ApplyDeltaToInt(opts *bind.TransactOpts, key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "applyDeltaToInt", key, value)
}

// ApplyDeltaToInt is a paid mutator transaction binding the contract method 0xe4e36c4e.
//
// Solidity: function applyDeltaToInt(bytes32 key, int256 value) returns(int256)
func (_Datastore *DatastoreSession) ApplyDeltaToInt(key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.Contract.ApplyDeltaToInt(&_Datastore.TransactOpts, key, value)
}

// ApplyDeltaToInt is a paid mutator transaction binding the contract method 0xe4e36c4e.
//
// Solidity: function applyDeltaToInt(bytes32 key, int256 value) returns(int256)
func (_Datastore *DatastoreTransactorSession) ApplyDeltaToInt(key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.Contract.ApplyDeltaToInt(&_Datastore.TransactOpts, key, value)
}

// ApplyDeltaToUint is a paid mutator transaction binding the contract method 0x32f85bbf.
//
// Solidity: function applyDeltaToUint(bytes32 key, int256 value, string errorMessage) returns(uint256)
func (_Datastore *DatastoreTransactor) ApplyDeltaToUint(opts *bind.TransactOpts, key [32]byte, value *big.Int, errorMessage string) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "applyDeltaToUint", key, value, errorMessage)
}

// ApplyDeltaToUint is a paid mutator transaction binding the contract method 0x32f85bbf.
//
// Solidity: function applyDeltaToUint(bytes32 key, int256 value, string errorMessage) returns(uint256)
func (_Datastore *DatastoreSession) ApplyDeltaToUint(key [32]byte, value *big.Int, errorMessage string) (*types.Transaction, error) {
	return _Datastore.Contract.ApplyDeltaToUint(&_Datastore.TransactOpts, key, value, errorMessage)
}

// ApplyDeltaToUint is a paid mutator transaction binding the contract method 0x32f85bbf.
//
// Solidity: function applyDeltaToUint(bytes32 key, int256 value, string errorMessage) returns(uint256)
func (_Datastore *DatastoreTransactorSession) ApplyDeltaToUint(key [32]byte, value *big.Int, errorMessage string) (*types.Transaction, error) {
	return _Datastore.Contract.ApplyDeltaToUint(&_Datastore.TransactOpts, key, value, errorMessage)
}

// ApplyDeltaToUint0 is a paid mutator transaction binding the contract method 0x3dbacd1a.
//
// Solidity: function applyDeltaToUint(bytes32 key, uint256 value) returns(uint256)
func (_Datastore *DatastoreTransactor) ApplyDeltaToUint0(opts *bind.TransactOpts, key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "applyDeltaToUint0", key, value)
}

// ApplyDeltaToUint0 is a paid mutator transaction binding the contract method 0x3dbacd1a.
//
// Solidity: function applyDeltaToUint(bytes32 key, uint256 value) returns(uint256)
func (_Datastore *DatastoreSession) ApplyDeltaToUint0(key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.Contract.ApplyDeltaToUint0(&_Datastore.TransactOpts, key, value)
}

// ApplyDeltaToUint0 is a paid mutator transaction binding the contract method 0x3dbacd1a.
//
// Solidity: function applyDeltaToUint(bytes32 key, uint256 value) returns(uint256)
func (_Datastore *DatastoreTransactorSession) ApplyDeltaToUint0(key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.Contract.ApplyDeltaToUint0(&_Datastore.TransactOpts, key, value)
}

// DecrementInt is a paid mutator transaction binding the contract method 0x6fae54f0.
//
// Solidity: function decrementInt(bytes32 key, int256 value) returns(int256)
func (_Datastore *DatastoreTransactor) DecrementInt(opts *bind.TransactOpts, key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "decrementInt", key, value)
}

// DecrementInt is a paid mutator transaction binding the contract method 0x6fae54f0.
//
// Solidity: function decrementInt(bytes32 key, int256 value) returns(int256)
func (_Datastore *DatastoreSession) DecrementInt(key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.Contract.DecrementInt(&_Datastore.TransactOpts, key, value)
}

// DecrementInt is a paid mutator transaction binding the contract method 0x6fae54f0.
//
// Solidity: function decrementInt(bytes32 key, int256 value) returns(int256)
func (_Datastore *DatastoreTransactorSession) DecrementInt(key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.Contract.DecrementInt(&_Datastore.TransactOpts, key, value)
}

// DecrementUint is a paid mutator transaction binding the contract method 0xe98aabc1.
//
// Solidity: function decrementUint(bytes32 key, uint256 value) returns(uint256)
func (_Datastore *DatastoreTransactor) DecrementUint(opts *bind.TransactOpts, key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "decrementUint", key, value)
}

// DecrementUint is a paid mutator transaction binding the contract method 0xe98aabc1.
//
// Solidity: function decrementUint(bytes32 key, uint256 value) returns(uint256)
func (_Datastore *DatastoreSession) DecrementUint(key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.Contract.DecrementUint(&_Datastore.TransactOpts, key, value)
}

// DecrementUint is a paid mutator transaction binding the contract method 0xe98aabc1.
//
// Solidity: function decrementUint(bytes32 key, uint256 value) returns(uint256)
func (_Datastore *DatastoreTransactorSession) DecrementUint(key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.Contract.DecrementUint(&_Datastore.TransactOpts, key, value)
}

// IncrementInt is a paid mutator transaction binding the contract method 0xcbb093dd.
//
// Solidity: function incrementInt(bytes32 key, int256 value) returns(int256)
func (_Datastore *DatastoreTransactor) IncrementInt(opts *bind.TransactOpts, key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "incrementInt", key, value)
}

// IncrementInt is a paid mutator transaction binding the contract method 0xcbb093dd.
//
// Solidity: function incrementInt(bytes32 key, int256 value) returns(int256)
func (_Datastore *DatastoreSession) IncrementInt(key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.Contract.IncrementInt(&_Datastore.TransactOpts, key, value)
}

// IncrementInt is a paid mutator transaction binding the contract method 0xcbb093dd.
//
// Solidity: function incrementInt(bytes32 key, int256 value) returns(int256)
func (_Datastore *DatastoreTransactorSession) IncrementInt(key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.Contract.IncrementInt(&_Datastore.TransactOpts, key, value)
}

// IncrementUint is a paid mutator transaction binding the contract method 0x340dbab3.
//
// Solidity: function incrementUint(bytes32 key, uint256 value) returns(uint256)
func (_Datastore *DatastoreTransactor) IncrementUint(opts *bind.TransactOpts, key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "incrementUint", key, value)
}

// IncrementUint is a paid mutator transaction binding the contract method 0x340dbab3.
//
// Solidity: function incrementUint(bytes32 key, uint256 value) returns(uint256)
func (_Datastore *DatastoreSession) IncrementUint(key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.Contract.IncrementUint(&_Datastore.TransactOpts, key, value)
}

// IncrementUint is a paid mutator transaction binding the contract method 0x340dbab3.
//
// Solidity: function incrementUint(bytes32 key, uint256 value) returns(uint256)
func (_Datastore *DatastoreTransactorSession) IncrementUint(key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.Contract.IncrementUint(&_Datastore.TransactOpts, key, value)
}

// RemoveAddress is a paid mutator transaction binding the contract method 0x69721d41.
//
// Solidity: function removeAddress(bytes32 setKey, address value) returns()
func (_Datastore *DatastoreTransactor) RemoveAddress(opts *bind.TransactOpts, setKey [32]byte, value common.Address) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "removeAddress", setKey, value)
}

// RemoveAddress is a paid mutator transaction binding the contract method 0x69721d41.
//
// Solidity: function removeAddress(bytes32 setKey, address value) returns()
func (_Datastore *DatastoreSession) RemoveAddress(setKey [32]byte, value common.Address) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveAddress(&_Datastore.TransactOpts, setKey, value)
}

// RemoveAddress is a paid mutator transaction binding the contract method 0x69721d41.
//
// Solidity: function removeAddress(bytes32 setKey, address value) returns()
func (_Datastore *DatastoreTransactorSession) RemoveAddress(setKey [32]byte, value common.Address) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveAddress(&_Datastore.TransactOpts, setKey, value)
}

// RemoveAddress0 is a paid mutator transaction binding the contract method 0x9faf6fb6.
//
// Solidity: function removeAddress(bytes32 key) returns()
func (_Datastore *DatastoreTransactor) RemoveAddress0(opts *bind.TransactOpts, key [32]byte) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "removeAddress0", key)
}

// RemoveAddress0 is a paid mutator transaction binding the contract method 0x9faf6fb6.
//
// Solidity: function removeAddress(bytes32 key) returns()
func (_Datastore *DatastoreSession) RemoveAddress0(key [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveAddress0(&_Datastore.TransactOpts, key)
}

// RemoveAddress0 is a paid mutator transaction binding the contract method 0x9faf6fb6.
//
// Solidity: function removeAddress(bytes32 key) returns()
func (_Datastore *DatastoreTransactorSession) RemoveAddress0(key [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveAddress0(&_Datastore.TransactOpts, key)
}

// RemoveAddressArray is a paid mutator transaction binding the contract method 0xc1dc9182.
//
// Solidity: function removeAddressArray(bytes32 key) returns()
func (_Datastore *DatastoreTransactor) RemoveAddressArray(opts *bind.TransactOpts, key [32]byte) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "removeAddressArray", key)
}

// RemoveAddressArray is a paid mutator transaction binding the contract method 0xc1dc9182.
//
// Solidity: function removeAddressArray(bytes32 key) returns()
func (_Datastore *DatastoreSession) RemoveAddressArray(key [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveAddressArray(&_Datastore.TransactOpts, key)
}

// RemoveAddressArray is a paid mutator transaction binding the contract method 0xc1dc9182.
//
// Solidity: function removeAddressArray(bytes32 key) returns()
func (_Datastore *DatastoreTransactorSession) RemoveAddressArray(key [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveAddressArray(&_Datastore.TransactOpts, key)
}

// RemoveBool is a paid mutator transaction binding the contract method 0x9fe7ac12.
//
// Solidity: function removeBool(bytes32 key) returns()
func (_Datastore *DatastoreTransactor) RemoveBool(opts *bind.TransactOpts, key [32]byte) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "removeBool", key)
}

// RemoveBool is a paid mutator transaction binding the contract method 0x9fe7ac12.
//
// Solidity: function removeBool(bytes32 key) returns()
func (_Datastore *DatastoreSession) RemoveBool(key [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveBool(&_Datastore.TransactOpts, key)
}

// RemoveBool is a paid mutator transaction binding the contract method 0x9fe7ac12.
//
// Solidity: function removeBool(bytes32 key) returns()
func (_Datastore *DatastoreTransactorSession) RemoveBool(key [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveBool(&_Datastore.TransactOpts, key)
}

// RemoveBoolArray is a paid mutator transaction binding the contract method 0xf51fc0d9.
//
// Solidity: function removeBoolArray(bytes32 key) returns()
func (_Datastore *DatastoreTransactor) RemoveBoolArray(opts *bind.TransactOpts, key [32]byte) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "removeBoolArray", key)
}

// RemoveBoolArray is a paid mutator transaction binding the contract method 0xf51fc0d9.
//
// Solidity: function removeBoolArray(bytes32 key) returns()
func (_Datastore *DatastoreSession) RemoveBoolArray(key [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveBoolArray(&_Datastore.TransactOpts, key)
}

// RemoveBoolArray is a paid mutator transaction binding the contract method 0xf51fc0d9.
//
// Solidity: function removeBoolArray(bytes32 key) returns()
func (_Datastore *DatastoreTransactorSession) RemoveBoolArray(key [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveBoolArray(&_Datastore.TransactOpts, key)
}

// RemoveBytes32 is a paid mutator transaction binding the contract method 0x9921c3cc.
//
// Solidity: function removeBytes32(bytes32 setKey, bytes32 value) returns()
func (_Datastore *DatastoreTransactor) RemoveBytes32(opts *bind.TransactOpts, setKey [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "removeBytes32", setKey, value)
}

// RemoveBytes32 is a paid mutator transaction binding the contract method 0x9921c3cc.
//
// Solidity: function removeBytes32(bytes32 setKey, bytes32 value) returns()
func (_Datastore *DatastoreSession) RemoveBytes32(setKey [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveBytes32(&_Datastore.TransactOpts, setKey, value)
}

// RemoveBytes32 is a paid mutator transaction binding the contract method 0x9921c3cc.
//
// Solidity: function removeBytes32(bytes32 setKey, bytes32 value) returns()
func (_Datastore *DatastoreTransactorSession) RemoveBytes32(setKey [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveBytes32(&_Datastore.TransactOpts, setKey, value)
}

// RemoveBytes320 is a paid mutator transaction binding the contract method 0xcf6a8722.
//
// Solidity: function removeBytes32(bytes32 key) returns()
func (_Datastore *DatastoreTransactor) RemoveBytes320(opts *bind.TransactOpts, key [32]byte) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "removeBytes320", key)
}

// RemoveBytes320 is a paid mutator transaction binding the contract method 0xcf6a8722.
//
// Solidity: function removeBytes32(bytes32 key) returns()
func (_Datastore *DatastoreSession) RemoveBytes320(key [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveBytes320(&_Datastore.TransactOpts, key)
}

// RemoveBytes320 is a paid mutator transaction binding the contract method 0xcf6a8722.
//
// Solidity: function removeBytes32(bytes32 key) returns()
func (_Datastore *DatastoreTransactorSession) RemoveBytes320(key [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveBytes320(&_Datastore.TransactOpts, key)
}

// RemoveBytes32Array is a paid mutator transaction binding the contract method 0xbf7f035a.
//
// Solidity: function removeBytes32Array(bytes32 key) returns()
func (_Datastore *DatastoreTransactor) RemoveBytes32Array(opts *bind.TransactOpts, key [32]byte) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "removeBytes32Array", key)
}

// RemoveBytes32Array is a paid mutator transaction binding the contract method 0xbf7f035a.
//
// Solidity: function removeBytes32Array(bytes32 key) returns()
func (_Datastore *DatastoreSession) RemoveBytes32Array(key [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveBytes32Array(&_Datastore.TransactOpts, key)
}

// RemoveBytes32Array is a paid mutator transaction binding the contract method 0xbf7f035a.
//
// Solidity: function removeBytes32Array(bytes32 key) returns()
func (_Datastore *DatastoreTransactorSession) RemoveBytes32Array(key [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveBytes32Array(&_Datastore.TransactOpts, key)
}

// RemoveInt is a paid mutator transaction binding the contract method 0xe62461ce.
//
// Solidity: function removeInt(bytes32 key) returns()
func (_Datastore *DatastoreTransactor) RemoveInt(opts *bind.TransactOpts, key [32]byte) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "removeInt", key)
}

// RemoveInt is a paid mutator transaction binding the contract method 0xe62461ce.
//
// Solidity: function removeInt(bytes32 key) returns()
func (_Datastore *DatastoreSession) RemoveInt(key [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveInt(&_Datastore.TransactOpts, key)
}

// RemoveInt is a paid mutator transaction binding the contract method 0xe62461ce.
//
// Solidity: function removeInt(bytes32 key) returns()
func (_Datastore *DatastoreTransactorSession) RemoveInt(key [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveInt(&_Datastore.TransactOpts, key)
}

// RemoveIntArray is a paid mutator transaction binding the contract method 0x499ea50e.
//
// Solidity: function removeIntArray(bytes32 key) returns()
func (_Datastore *DatastoreTransactor) RemoveIntArray(opts *bind.TransactOpts, key [32]byte) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "removeIntArray", key)
}

// RemoveIntArray is a paid mutator transaction binding the contract method 0x499ea50e.
//
// Solidity: function removeIntArray(bytes32 key) returns()
func (_Datastore *DatastoreSession) RemoveIntArray(key [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveIntArray(&_Datastore.TransactOpts, key)
}

// RemoveIntArray is a paid mutator transaction binding the contract method 0x499ea50e.
//
// Solidity: function removeIntArray(bytes32 key) returns()
func (_Datastore *DatastoreTransactorSession) RemoveIntArray(key [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveIntArray(&_Datastore.TransactOpts, key)
}

// RemoveString is a paid mutator transaction binding the contract method 0xcc50eadd.
//
// Solidity: function removeString(bytes32 key) returns()
func (_Datastore *DatastoreTransactor) RemoveString(opts *bind.TransactOpts, key [32]byte) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "removeString", key)
}

// RemoveString is a paid mutator transaction binding the contract method 0xcc50eadd.
//
// Solidity: function removeString(bytes32 key) returns()
func (_Datastore *DatastoreSession) RemoveString(key [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveString(&_Datastore.TransactOpts, key)
}

// RemoveString is a paid mutator transaction binding the contract method 0xcc50eadd.
//
// Solidity: function removeString(bytes32 key) returns()
func (_Datastore *DatastoreTransactorSession) RemoveString(key [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveString(&_Datastore.TransactOpts, key)
}

// RemoveStringArray is a paid mutator transaction binding the contract method 0xe208a70d.
//
// Solidity: function removeStringArray(bytes32 key) returns()
func (_Datastore *DatastoreTransactor) RemoveStringArray(opts *bind.TransactOpts, key [32]byte) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "removeStringArray", key)
}

// RemoveStringArray is a paid mutator transaction binding the contract method 0xe208a70d.
//
// Solidity: function removeStringArray(bytes32 key) returns()
func (_Datastore *DatastoreSession) RemoveStringArray(key [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveStringArray(&_Datastore.TransactOpts, key)
}

// RemoveStringArray is a paid mutator transaction binding the contract method 0xe208a70d.
//
// Solidity: function removeStringArray(bytes32 key) returns()
func (_Datastore *DatastoreTransactorSession) RemoveStringArray(key [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveStringArray(&_Datastore.TransactOpts, key)
}

// RemoveUint is a paid mutator transaction binding the contract method 0x42c3bd96.
//
// Solidity: function removeUint(bytes32 key) returns()
func (_Datastore *DatastoreTransactor) RemoveUint(opts *bind.TransactOpts, key [32]byte) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "removeUint", key)
}

// RemoveUint is a paid mutator transaction binding the contract method 0x42c3bd96.
//
// Solidity: function removeUint(bytes32 key) returns()
func (_Datastore *DatastoreSession) RemoveUint(key [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveUint(&_Datastore.TransactOpts, key)
}

// RemoveUint is a paid mutator transaction binding the contract method 0x42c3bd96.
//
// Solidity: function removeUint(bytes32 key) returns()
func (_Datastore *DatastoreTransactorSession) RemoveUint(key [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveUint(&_Datastore.TransactOpts, key)
}

// RemoveUint0 is a paid mutator transaction binding the contract method 0x93266f9a.
//
// Solidity: function removeUint(bytes32 setKey, uint256 value) returns()
func (_Datastore *DatastoreTransactor) RemoveUint0(opts *bind.TransactOpts, setKey [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "removeUint0", setKey, value)
}

// RemoveUint0 is a paid mutator transaction binding the contract method 0x93266f9a.
//
// Solidity: function removeUint(bytes32 setKey, uint256 value) returns()
func (_Datastore *DatastoreSession) RemoveUint0(setKey [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveUint0(&_Datastore.TransactOpts, setKey, value)
}

// RemoveUint0 is a paid mutator transaction binding the contract method 0x93266f9a.
//
// Solidity: function removeUint(bytes32 setKey, uint256 value) returns()
func (_Datastore *DatastoreTransactorSession) RemoveUint0(setKey [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveUint0(&_Datastore.TransactOpts, setKey, value)
}

// RemoveUintArray is a paid mutator transaction binding the contract method 0xbe43caa3.
//
// Solidity: function removeUintArray(bytes32 key) returns()
func (_Datastore *DatastoreTransactor) RemoveUintArray(opts *bind.TransactOpts, key [32]byte) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "removeUintArray", key)
}

// RemoveUintArray is a paid mutator transaction binding the contract method 0xbe43caa3.
//
// Solidity: function removeUintArray(bytes32 key) returns()
func (_Datastore *DatastoreSession) RemoveUintArray(key [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveUintArray(&_Datastore.TransactOpts, key)
}

// RemoveUintArray is a paid mutator transaction binding the contract method 0xbe43caa3.
//
// Solidity: function removeUintArray(bytes32 key) returns()
func (_Datastore *DatastoreTransactorSession) RemoveUintArray(key [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.RemoveUintArray(&_Datastore.TransactOpts, key)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 key, address value) returns(address)
func (_Datastore *DatastoreTransactor) SetAddress(opts *bind.TransactOpts, key [32]byte, value common.Address) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "setAddress", key, value)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 key, address value) returns(address)
func (_Datastore *DatastoreSession) SetAddress(key [32]byte, value common.Address) (*types.Transaction, error) {
	return _Datastore.Contract.SetAddress(&_Datastore.TransactOpts, key, value)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 key, address value) returns(address)
func (_Datastore *DatastoreTransactorSession) SetAddress(key [32]byte, value common.Address) (*types.Transaction, error) {
	return _Datastore.Contract.SetAddress(&_Datastore.TransactOpts, key, value)
}

// SetAddressArray is a paid mutator transaction binding the contract method 0xec672cf6.
//
// Solidity: function setAddressArray(bytes32 key, address[] value) returns()
func (_Datastore *DatastoreTransactor) SetAddressArray(opts *bind.TransactOpts, key [32]byte, value []common.Address) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "setAddressArray", key, value)
}

// SetAddressArray is a paid mutator transaction binding the contract method 0xec672cf6.
//
// Solidity: function setAddressArray(bytes32 key, address[] value) returns()
func (_Datastore *DatastoreSession) SetAddressArray(key [32]byte, value []common.Address) (*types.Transaction, error) {
	return _Datastore.Contract.SetAddressArray(&_Datastore.TransactOpts, key, value)
}

// SetAddressArray is a paid mutator transaction binding the contract method 0xec672cf6.
//
// Solidity: function setAddressArray(bytes32 key, address[] value) returns()
func (_Datastore *DatastoreTransactorSession) SetAddressArray(key [32]byte, value []common.Address) (*types.Transaction, error) {
	return _Datastore.Contract.SetAddressArray(&_Datastore.TransactOpts, key, value)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(bytes32 key, bool value) returns(bool)
func (_Datastore *DatastoreTransactor) SetBool(opts *bind.TransactOpts, key [32]byte, value bool) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "setBool", key, value)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(bytes32 key, bool value) returns(bool)
func (_Datastore *DatastoreSession) SetBool(key [32]byte, value bool) (*types.Transaction, error) {
	return _Datastore.Contract.SetBool(&_Datastore.TransactOpts, key, value)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(bytes32 key, bool value) returns(bool)
func (_Datastore *DatastoreTransactorSession) SetBool(key [32]byte, value bool) (*types.Transaction, error) {
	return _Datastore.Contract.SetBool(&_Datastore.TransactOpts, key, value)
}

// SetBoolArray is a paid mutator transaction binding the contract method 0x35d4d407.
//
// Solidity: function setBoolArray(bytes32 key, bool[] value) returns()
func (_Datastore *DatastoreTransactor) SetBoolArray(opts *bind.TransactOpts, key [32]byte, value []bool) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "setBoolArray", key, value)
}

// SetBoolArray is a paid mutator transaction binding the contract method 0x35d4d407.
//
// Solidity: function setBoolArray(bytes32 key, bool[] value) returns()
func (_Datastore *DatastoreSession) SetBoolArray(key [32]byte, value []bool) (*types.Transaction, error) {
	return _Datastore.Contract.SetBoolArray(&_Datastore.TransactOpts, key, value)
}

// SetBoolArray is a paid mutator transaction binding the contract method 0x35d4d407.
//
// Solidity: function setBoolArray(bytes32 key, bool[] value) returns()
func (_Datastore *DatastoreTransactorSession) SetBoolArray(key [32]byte, value []bool) (*types.Transaction, error) {
	return _Datastore.Contract.SetBoolArray(&_Datastore.TransactOpts, key, value)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0x4e91db08.
//
// Solidity: function setBytes32(bytes32 key, bytes32 value) returns(bytes32)
func (_Datastore *DatastoreTransactor) SetBytes32(opts *bind.TransactOpts, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "setBytes32", key, value)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0x4e91db08.
//
// Solidity: function setBytes32(bytes32 key, bytes32 value) returns(bytes32)
func (_Datastore *DatastoreSession) SetBytes32(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.SetBytes32(&_Datastore.TransactOpts, key, value)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0x4e91db08.
//
// Solidity: function setBytes32(bytes32 key, bytes32 value) returns(bytes32)
func (_Datastore *DatastoreTransactorSession) SetBytes32(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.SetBytes32(&_Datastore.TransactOpts, key, value)
}

// SetBytes32Array is a paid mutator transaction binding the contract method 0x26004846.
//
// Solidity: function setBytes32Array(bytes32 key, bytes32[] value) returns()
func (_Datastore *DatastoreTransactor) SetBytes32Array(opts *bind.TransactOpts, key [32]byte, value [][32]byte) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "setBytes32Array", key, value)
}

// SetBytes32Array is a paid mutator transaction binding the contract method 0x26004846.
//
// Solidity: function setBytes32Array(bytes32 key, bytes32[] value) returns()
func (_Datastore *DatastoreSession) SetBytes32Array(key [32]byte, value [][32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.SetBytes32Array(&_Datastore.TransactOpts, key, value)
}

// SetBytes32Array is a paid mutator transaction binding the contract method 0x26004846.
//
// Solidity: function setBytes32Array(bytes32 key, bytes32[] value) returns()
func (_Datastore *DatastoreTransactorSession) SetBytes32Array(key [32]byte, value [][32]byte) (*types.Transaction, error) {
	return _Datastore.Contract.SetBytes32Array(&_Datastore.TransactOpts, key, value)
}

// SetInt is a paid mutator transaction binding the contract method 0x3e49bed0.
//
// Solidity: function setInt(bytes32 key, int256 value) returns(int256)
func (_Datastore *DatastoreTransactor) SetInt(opts *bind.TransactOpts, key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "setInt", key, value)
}

// SetInt is a paid mutator transaction binding the contract method 0x3e49bed0.
//
// Solidity: function setInt(bytes32 key, int256 value) returns(int256)
func (_Datastore *DatastoreSession) SetInt(key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.Contract.SetInt(&_Datastore.TransactOpts, key, value)
}

// SetInt is a paid mutator transaction binding the contract method 0x3e49bed0.
//
// Solidity: function setInt(bytes32 key, int256 value) returns(int256)
func (_Datastore *DatastoreTransactorSession) SetInt(key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.Contract.SetInt(&_Datastore.TransactOpts, key, value)
}

// SetIntArray is a paid mutator transaction binding the contract method 0xa9fcf76b.
//
// Solidity: function setIntArray(bytes32 key, int256[] value) returns()
func (_Datastore *DatastoreTransactor) SetIntArray(opts *bind.TransactOpts, key [32]byte, value []*big.Int) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "setIntArray", key, value)
}

// SetIntArray is a paid mutator transaction binding the contract method 0xa9fcf76b.
//
// Solidity: function setIntArray(bytes32 key, int256[] value) returns()
func (_Datastore *DatastoreSession) SetIntArray(key [32]byte, value []*big.Int) (*types.Transaction, error) {
	return _Datastore.Contract.SetIntArray(&_Datastore.TransactOpts, key, value)
}

// SetIntArray is a paid mutator transaction binding the contract method 0xa9fcf76b.
//
// Solidity: function setIntArray(bytes32 key, int256[] value) returns()
func (_Datastore *DatastoreTransactorSession) SetIntArray(key [32]byte, value []*big.Int) (*types.Transaction, error) {
	return _Datastore.Contract.SetIntArray(&_Datastore.TransactOpts, key, value)
}

// SetString is a paid mutator transaction binding the contract method 0x6e899550.
//
// Solidity: function setString(bytes32 key, string value) returns(string)
func (_Datastore *DatastoreTransactor) SetString(opts *bind.TransactOpts, key [32]byte, value string) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "setString", key, value)
}

// SetString is a paid mutator transaction binding the contract method 0x6e899550.
//
// Solidity: function setString(bytes32 key, string value) returns(string)
func (_Datastore *DatastoreSession) SetString(key [32]byte, value string) (*types.Transaction, error) {
	return _Datastore.Contract.SetString(&_Datastore.TransactOpts, key, value)
}

// SetString is a paid mutator transaction binding the contract method 0x6e899550.
//
// Solidity: function setString(bytes32 key, string value) returns(string)
func (_Datastore *DatastoreTransactorSession) SetString(key [32]byte, value string) (*types.Transaction, error) {
	return _Datastore.Contract.SetString(&_Datastore.TransactOpts, key, value)
}

// SetStringArray is a paid mutator transaction binding the contract method 0x88021a72.
//
// Solidity: function setStringArray(bytes32 key, string[] value) returns()
func (_Datastore *DatastoreTransactor) SetStringArray(opts *bind.TransactOpts, key [32]byte, value []string) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "setStringArray", key, value)
}

// SetStringArray is a paid mutator transaction binding the contract method 0x88021a72.
//
// Solidity: function setStringArray(bytes32 key, string[] value) returns()
func (_Datastore *DatastoreSession) SetStringArray(key [32]byte, value []string) (*types.Transaction, error) {
	return _Datastore.Contract.SetStringArray(&_Datastore.TransactOpts, key, value)
}

// SetStringArray is a paid mutator transaction binding the contract method 0x88021a72.
//
// Solidity: function setStringArray(bytes32 key, string[] value) returns()
func (_Datastore *DatastoreTransactorSession) SetStringArray(key [32]byte, value []string) (*types.Transaction, error) {
	return _Datastore.Contract.SetStringArray(&_Datastore.TransactOpts, key, value)
}

// SetUint is a paid mutator transaction binding the contract method 0xe2a4853a.
//
// Solidity: function setUint(bytes32 key, uint256 value) returns(uint256)
func (_Datastore *DatastoreTransactor) SetUint(opts *bind.TransactOpts, key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "setUint", key, value)
}

// SetUint is a paid mutator transaction binding the contract method 0xe2a4853a.
//
// Solidity: function setUint(bytes32 key, uint256 value) returns(uint256)
func (_Datastore *DatastoreSession) SetUint(key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.Contract.SetUint(&_Datastore.TransactOpts, key, value)
}

// SetUint is a paid mutator transaction binding the contract method 0xe2a4853a.
//
// Solidity: function setUint(bytes32 key, uint256 value) returns(uint256)
func (_Datastore *DatastoreTransactorSession) SetUint(key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Datastore.Contract.SetUint(&_Datastore.TransactOpts, key, value)
}

// SetUintArray is a paid mutator transaction binding the contract method 0x5eb07dbd.
//
// Solidity: function setUintArray(bytes32 key, uint256[] value) returns()
func (_Datastore *DatastoreTransactor) SetUintArray(opts *bind.TransactOpts, key [32]byte, value []*big.Int) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "setUintArray", key, value)
}

// SetUintArray is a paid mutator transaction binding the contract method 0x5eb07dbd.
//
// Solidity: function setUintArray(bytes32 key, uint256[] value) returns()
func (_Datastore *DatastoreSession) SetUintArray(key [32]byte, value []*big.Int) (*types.Transaction, error) {
	return _Datastore.Contract.SetUintArray(&_Datastore.TransactOpts, key, value)
}

// SetUintArray is a paid mutator transaction binding the contract method 0x5eb07dbd.
//
// Solidity: function setUintArray(bytes32 key, uint256[] value) returns()
func (_Datastore *DatastoreTransactorSession) SetUintArray(key [32]byte, value []*big.Int) (*types.Transaction, error) {
	return _Datastore.Contract.SetUintArray(&_Datastore.TransactOpts, key, value)
}
