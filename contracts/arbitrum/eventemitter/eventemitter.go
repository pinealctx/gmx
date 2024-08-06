// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eventemitter

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

// EventUtilsAddressArrayKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsAddressArrayKeyValue struct {
	Key   string
	Value []common.Address
}

// EventUtilsAddressItems is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsAddressItems struct {
	Items      []EventUtilsAddressKeyValue
	ArrayItems []EventUtilsAddressArrayKeyValue
}

// EventUtilsAddressKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsAddressKeyValue struct {
	Key   string
	Value common.Address
}

// EventUtilsBoolArrayKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsBoolArrayKeyValue struct {
	Key   string
	Value []bool
}

// EventUtilsBoolItems is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsBoolItems struct {
	Items      []EventUtilsBoolKeyValue
	ArrayItems []EventUtilsBoolArrayKeyValue
}

// EventUtilsBoolKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsBoolKeyValue struct {
	Key   string
	Value bool
}

// EventUtilsBytes32ArrayKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsBytes32ArrayKeyValue struct {
	Key   string
	Value [][32]byte
}

// EventUtilsBytes32Items is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsBytes32Items struct {
	Items      []EventUtilsBytes32KeyValue
	ArrayItems []EventUtilsBytes32ArrayKeyValue
}

// EventUtilsBytes32KeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsBytes32KeyValue struct {
	Key   string
	Value [32]byte
}

// EventUtilsBytesArrayKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsBytesArrayKeyValue struct {
	Key   string
	Value [][]byte
}

// EventUtilsBytesItems is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsBytesItems struct {
	Items      []EventUtilsBytesKeyValue
	ArrayItems []EventUtilsBytesArrayKeyValue
}

// EventUtilsBytesKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsBytesKeyValue struct {
	Key   string
	Value []byte
}

// EventUtilsEventLogData is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsEventLogData struct {
	AddressItems EventUtilsAddressItems
	UintItems    EventUtilsUintItems
	IntItems     EventUtilsIntItems
	BoolItems    EventUtilsBoolItems
	Bytes32Items EventUtilsBytes32Items
	BytesItems   EventUtilsBytesItems
	StringItems  EventUtilsStringItems
}

// EventUtilsIntArrayKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsIntArrayKeyValue struct {
	Key   string
	Value []*big.Int
}

// EventUtilsIntItems is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsIntItems struct {
	Items      []EventUtilsIntKeyValue
	ArrayItems []EventUtilsIntArrayKeyValue
}

// EventUtilsIntKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsIntKeyValue struct {
	Key   string
	Value *big.Int
}

// EventUtilsStringArrayKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsStringArrayKeyValue struct {
	Key   string
	Value []string
}

// EventUtilsStringItems is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsStringItems struct {
	Items      []EventUtilsStringKeyValue
	ArrayItems []EventUtilsStringArrayKeyValue
}

// EventUtilsStringKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsStringKeyValue struct {
	Key   string
	Value string
}

// EventUtilsUintArrayKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsUintArrayKeyValue struct {
	Key   string
	Value []*big.Int
}

// EventUtilsUintItems is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsUintItems struct {
	Items      []EventUtilsUintKeyValue
	ArrayItems []EventUtilsUintArrayKeyValue
}

// EventUtilsUintKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsUintKeyValue struct {
	Key   string
	Value *big.Int
}

// EventemitterMetaData contains all meta data concerning the Eventemitter contract.
var EventemitterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractRoleStore\",\"name\":\"_roleStore\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"eventName\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"eventNameHash\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"internalType\":\"structEventUtils.AddressKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"value\",\"type\":\"address[]\"}],\"internalType\":\"structEventUtils.AddressArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.AddressItems\",\"name\":\"addressItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structEventUtils.UintKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"}],\"internalType\":\"structEventUtils.UintArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.UintItems\",\"name\":\"uintItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"internalType\":\"structEventUtils.IntKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"int256[]\",\"name\":\"value\",\"type\":\"int256[]\"}],\"internalType\":\"structEventUtils.IntArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.IntItems\",\"name\":\"intItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"internalType\":\"structEventUtils.BoolKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bool[]\",\"name\":\"value\",\"type\":\"bool[]\"}],\"internalType\":\"structEventUtils.BoolArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.BoolItems\",\"name\":\"boolItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"internalType\":\"structEventUtils.Bytes32KeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes32[]\",\"name\":\"value\",\"type\":\"bytes32[]\"}],\"internalType\":\"structEventUtils.Bytes32ArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.Bytes32Items\",\"name\":\"bytes32Items\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structEventUtils.BytesKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"value\",\"type\":\"bytes[]\"}],\"internalType\":\"structEventUtils.BytesArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.BytesItems\",\"name\":\"bytesItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structEventUtils.StringKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"value\",\"type\":\"string[]\"}],\"internalType\":\"structEventUtils.StringArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.StringItems\",\"name\":\"stringItems\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structEventUtils.EventLogData\",\"name\":\"eventData\",\"type\":\"tuple\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"eventName\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"eventNameHash\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"topic1\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"internalType\":\"structEventUtils.AddressKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"value\",\"type\":\"address[]\"}],\"internalType\":\"structEventUtils.AddressArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.AddressItems\",\"name\":\"addressItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structEventUtils.UintKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"}],\"internalType\":\"structEventUtils.UintArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.UintItems\",\"name\":\"uintItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"internalType\":\"structEventUtils.IntKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"int256[]\",\"name\":\"value\",\"type\":\"int256[]\"}],\"internalType\":\"structEventUtils.IntArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.IntItems\",\"name\":\"intItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"internalType\":\"structEventUtils.BoolKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bool[]\",\"name\":\"value\",\"type\":\"bool[]\"}],\"internalType\":\"structEventUtils.BoolArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.BoolItems\",\"name\":\"boolItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"internalType\":\"structEventUtils.Bytes32KeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes32[]\",\"name\":\"value\",\"type\":\"bytes32[]\"}],\"internalType\":\"structEventUtils.Bytes32ArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.Bytes32Items\",\"name\":\"bytes32Items\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structEventUtils.BytesKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"value\",\"type\":\"bytes[]\"}],\"internalType\":\"structEventUtils.BytesArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.BytesItems\",\"name\":\"bytesItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structEventUtils.StringKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"value\",\"type\":\"string[]\"}],\"internalType\":\"structEventUtils.StringArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.StringItems\",\"name\":\"stringItems\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structEventUtils.EventLogData\",\"name\":\"eventData\",\"type\":\"tuple\"}],\"name\":\"EventLog1\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"eventName\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"eventNameHash\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"topic1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"topic2\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"internalType\":\"structEventUtils.AddressKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"value\",\"type\":\"address[]\"}],\"internalType\":\"structEventUtils.AddressArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.AddressItems\",\"name\":\"addressItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structEventUtils.UintKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"}],\"internalType\":\"structEventUtils.UintArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.UintItems\",\"name\":\"uintItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"internalType\":\"structEventUtils.IntKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"int256[]\",\"name\":\"value\",\"type\":\"int256[]\"}],\"internalType\":\"structEventUtils.IntArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.IntItems\",\"name\":\"intItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"internalType\":\"structEventUtils.BoolKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bool[]\",\"name\":\"value\",\"type\":\"bool[]\"}],\"internalType\":\"structEventUtils.BoolArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.BoolItems\",\"name\":\"boolItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"internalType\":\"structEventUtils.Bytes32KeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes32[]\",\"name\":\"value\",\"type\":\"bytes32[]\"}],\"internalType\":\"structEventUtils.Bytes32ArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.Bytes32Items\",\"name\":\"bytes32Items\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structEventUtils.BytesKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"value\",\"type\":\"bytes[]\"}],\"internalType\":\"structEventUtils.BytesArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.BytesItems\",\"name\":\"bytesItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structEventUtils.StringKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"value\",\"type\":\"string[]\"}],\"internalType\":\"structEventUtils.StringArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.StringItems\",\"name\":\"stringItems\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structEventUtils.EventLogData\",\"name\":\"eventData\",\"type\":\"tuple\"}],\"name\":\"EventLog2\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"topic1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"emitDataLog1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"topic1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"topic2\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"emitDataLog2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"topic1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"topic2\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"topic3\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"emitDataLog3\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"topic1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"topic2\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"topic3\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"topic4\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"emitDataLog4\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventName\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"internalType\":\"structEventUtils.AddressKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"value\",\"type\":\"address[]\"}],\"internalType\":\"structEventUtils.AddressArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.AddressItems\",\"name\":\"addressItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structEventUtils.UintKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"}],\"internalType\":\"structEventUtils.UintArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.UintItems\",\"name\":\"uintItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"internalType\":\"structEventUtils.IntKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"int256[]\",\"name\":\"value\",\"type\":\"int256[]\"}],\"internalType\":\"structEventUtils.IntArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.IntItems\",\"name\":\"intItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"internalType\":\"structEventUtils.BoolKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bool[]\",\"name\":\"value\",\"type\":\"bool[]\"}],\"internalType\":\"structEventUtils.BoolArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.BoolItems\",\"name\":\"boolItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"internalType\":\"structEventUtils.Bytes32KeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes32[]\",\"name\":\"value\",\"type\":\"bytes32[]\"}],\"internalType\":\"structEventUtils.Bytes32ArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.Bytes32Items\",\"name\":\"bytes32Items\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structEventUtils.BytesKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"value\",\"type\":\"bytes[]\"}],\"internalType\":\"structEventUtils.BytesArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.BytesItems\",\"name\":\"bytesItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structEventUtils.StringKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"value\",\"type\":\"string[]\"}],\"internalType\":\"structEventUtils.StringArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.StringItems\",\"name\":\"stringItems\",\"type\":\"tuple\"}],\"internalType\":\"structEventUtils.EventLogData\",\"name\":\"eventData\",\"type\":\"tuple\"}],\"name\":\"emitEventLog\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventName\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"topic1\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"internalType\":\"structEventUtils.AddressKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"value\",\"type\":\"address[]\"}],\"internalType\":\"structEventUtils.AddressArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.AddressItems\",\"name\":\"addressItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structEventUtils.UintKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"}],\"internalType\":\"structEventUtils.UintArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.UintItems\",\"name\":\"uintItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"internalType\":\"structEventUtils.IntKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"int256[]\",\"name\":\"value\",\"type\":\"int256[]\"}],\"internalType\":\"structEventUtils.IntArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.IntItems\",\"name\":\"intItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"internalType\":\"structEventUtils.BoolKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bool[]\",\"name\":\"value\",\"type\":\"bool[]\"}],\"internalType\":\"structEventUtils.BoolArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.BoolItems\",\"name\":\"boolItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"internalType\":\"structEventUtils.Bytes32KeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes32[]\",\"name\":\"value\",\"type\":\"bytes32[]\"}],\"internalType\":\"structEventUtils.Bytes32ArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.Bytes32Items\",\"name\":\"bytes32Items\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structEventUtils.BytesKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"value\",\"type\":\"bytes[]\"}],\"internalType\":\"structEventUtils.BytesArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.BytesItems\",\"name\":\"bytesItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structEventUtils.StringKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"value\",\"type\":\"string[]\"}],\"internalType\":\"structEventUtils.StringArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.StringItems\",\"name\":\"stringItems\",\"type\":\"tuple\"}],\"internalType\":\"structEventUtils.EventLogData\",\"name\":\"eventData\",\"type\":\"tuple\"}],\"name\":\"emitEventLog1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventName\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"topic1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"topic2\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"internalType\":\"structEventUtils.AddressKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"value\",\"type\":\"address[]\"}],\"internalType\":\"structEventUtils.AddressArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.AddressItems\",\"name\":\"addressItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structEventUtils.UintKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"}],\"internalType\":\"structEventUtils.UintArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.UintItems\",\"name\":\"uintItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"internalType\":\"structEventUtils.IntKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"int256[]\",\"name\":\"value\",\"type\":\"int256[]\"}],\"internalType\":\"structEventUtils.IntArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.IntItems\",\"name\":\"intItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"internalType\":\"structEventUtils.BoolKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bool[]\",\"name\":\"value\",\"type\":\"bool[]\"}],\"internalType\":\"structEventUtils.BoolArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.BoolItems\",\"name\":\"boolItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"internalType\":\"structEventUtils.Bytes32KeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes32[]\",\"name\":\"value\",\"type\":\"bytes32[]\"}],\"internalType\":\"structEventUtils.Bytes32ArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.Bytes32Items\",\"name\":\"bytes32Items\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structEventUtils.BytesKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"value\",\"type\":\"bytes[]\"}],\"internalType\":\"structEventUtils.BytesArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.BytesItems\",\"name\":\"bytesItems\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structEventUtils.StringKeyValue[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"value\",\"type\":\"string[]\"}],\"internalType\":\"structEventUtils.StringArrayKeyValue[]\",\"name\":\"arrayItems\",\"type\":\"tuple[]\"}],\"internalType\":\"structEventUtils.StringItems\",\"name\":\"stringItems\",\"type\":\"tuple\"}],\"internalType\":\"structEventUtils.EventLogData\",\"name\":\"eventData\",\"type\":\"tuple\"}],\"name\":\"emitEventLog2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roleStore\",\"outputs\":[{\"internalType\":\"contractRoleStore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// EventemitterABI is the input ABI used to generate the binding from.
// Deprecated: Use EventemitterMetaData.ABI instead.
var EventemitterABI = EventemitterMetaData.ABI

// Eventemitter is an auto generated Go binding around an Ethereum contract.
type Eventemitter struct {
	EventemitterCaller     // Read-only binding to the contract
	EventemitterTransactor // Write-only binding to the contract
	EventemitterFilterer   // Log filterer for contract events
}

// EventemitterCaller is an auto generated read-only Go binding around an Ethereum contract.
type EventemitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventemitterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EventemitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventemitterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EventemitterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventemitterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EventemitterSession struct {
	Contract     *Eventemitter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventemitterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EventemitterCallerSession struct {
	Contract *EventemitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// EventemitterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EventemitterTransactorSession struct {
	Contract     *EventemitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// EventemitterRaw is an auto generated low-level Go binding around an Ethereum contract.
type EventemitterRaw struct {
	Contract *Eventemitter // Generic contract binding to access the raw methods on
}

// EventemitterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EventemitterCallerRaw struct {
	Contract *EventemitterCaller // Generic read-only contract binding to access the raw methods on
}

// EventemitterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EventemitterTransactorRaw struct {
	Contract *EventemitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEventemitter creates a new instance of Eventemitter, bound to a specific deployed contract.
func NewEventemitter(address common.Address, backend bind.ContractBackend) (*Eventemitter, error) {
	contract, err := bindEventemitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Eventemitter{EventemitterCaller: EventemitterCaller{contract: contract}, EventemitterTransactor: EventemitterTransactor{contract: contract}, EventemitterFilterer: EventemitterFilterer{contract: contract}}, nil
}

// NewEventemitterCaller creates a new read-only instance of Eventemitter, bound to a specific deployed contract.
func NewEventemitterCaller(address common.Address, caller bind.ContractCaller) (*EventemitterCaller, error) {
	contract, err := bindEventemitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventemitterCaller{contract: contract}, nil
}

// NewEventemitterTransactor creates a new write-only instance of Eventemitter, bound to a specific deployed contract.
func NewEventemitterTransactor(address common.Address, transactor bind.ContractTransactor) (*EventemitterTransactor, error) {
	contract, err := bindEventemitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventemitterTransactor{contract: contract}, nil
}

// NewEventemitterFilterer creates a new log filterer instance of Eventemitter, bound to a specific deployed contract.
func NewEventemitterFilterer(address common.Address, filterer bind.ContractFilterer) (*EventemitterFilterer, error) {
	contract, err := bindEventemitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventemitterFilterer{contract: contract}, nil
}

// bindEventemitter binds a generic wrapper to an already deployed contract.
func bindEventemitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EventemitterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eventemitter *EventemitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eventemitter.Contract.EventemitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eventemitter *EventemitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eventemitter.Contract.EventemitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eventemitter *EventemitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eventemitter.Contract.EventemitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eventemitter *EventemitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eventemitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eventemitter *EventemitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eventemitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eventemitter *EventemitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eventemitter.Contract.contract.Transact(opts, method, params...)
}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_Eventemitter *EventemitterCaller) RoleStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Eventemitter.contract.Call(opts, &out, "roleStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_Eventemitter *EventemitterSession) RoleStore() (common.Address, error) {
	return _Eventemitter.Contract.RoleStore(&_Eventemitter.CallOpts)
}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_Eventemitter *EventemitterCallerSession) RoleStore() (common.Address, error) {
	return _Eventemitter.Contract.RoleStore(&_Eventemitter.CallOpts)
}

// EmitDataLog1 is a paid mutator transaction binding the contract method 0xf9d5c0ea.
//
// Solidity: function emitDataLog1(bytes32 topic1, bytes data) returns()
func (_Eventemitter *EventemitterTransactor) EmitDataLog1(opts *bind.TransactOpts, topic1 [32]byte, data []byte) (*types.Transaction, error) {
	return _Eventemitter.contract.Transact(opts, "emitDataLog1", topic1, data)
}

// EmitDataLog1 is a paid mutator transaction binding the contract method 0xf9d5c0ea.
//
// Solidity: function emitDataLog1(bytes32 topic1, bytes data) returns()
func (_Eventemitter *EventemitterSession) EmitDataLog1(topic1 [32]byte, data []byte) (*types.Transaction, error) {
	return _Eventemitter.Contract.EmitDataLog1(&_Eventemitter.TransactOpts, topic1, data)
}

// EmitDataLog1 is a paid mutator transaction binding the contract method 0xf9d5c0ea.
//
// Solidity: function emitDataLog1(bytes32 topic1, bytes data) returns()
func (_Eventemitter *EventemitterTransactorSession) EmitDataLog1(topic1 [32]byte, data []byte) (*types.Transaction, error) {
	return _Eventemitter.Contract.EmitDataLog1(&_Eventemitter.TransactOpts, topic1, data)
}

// EmitDataLog2 is a paid mutator transaction binding the contract method 0xdda0db32.
//
// Solidity: function emitDataLog2(bytes32 topic1, bytes32 topic2, bytes data) returns()
func (_Eventemitter *EventemitterTransactor) EmitDataLog2(opts *bind.TransactOpts, topic1 [32]byte, topic2 [32]byte, data []byte) (*types.Transaction, error) {
	return _Eventemitter.contract.Transact(opts, "emitDataLog2", topic1, topic2, data)
}

// EmitDataLog2 is a paid mutator transaction binding the contract method 0xdda0db32.
//
// Solidity: function emitDataLog2(bytes32 topic1, bytes32 topic2, bytes data) returns()
func (_Eventemitter *EventemitterSession) EmitDataLog2(topic1 [32]byte, topic2 [32]byte, data []byte) (*types.Transaction, error) {
	return _Eventemitter.Contract.EmitDataLog2(&_Eventemitter.TransactOpts, topic1, topic2, data)
}

// EmitDataLog2 is a paid mutator transaction binding the contract method 0xdda0db32.
//
// Solidity: function emitDataLog2(bytes32 topic1, bytes32 topic2, bytes data) returns()
func (_Eventemitter *EventemitterTransactorSession) EmitDataLog2(topic1 [32]byte, topic2 [32]byte, data []byte) (*types.Transaction, error) {
	return _Eventemitter.Contract.EmitDataLog2(&_Eventemitter.TransactOpts, topic1, topic2, data)
}

// EmitDataLog3 is a paid mutator transaction binding the contract method 0xb3ac1c38.
//
// Solidity: function emitDataLog3(bytes32 topic1, bytes32 topic2, bytes32 topic3, bytes data) returns()
func (_Eventemitter *EventemitterTransactor) EmitDataLog3(opts *bind.TransactOpts, topic1 [32]byte, topic2 [32]byte, topic3 [32]byte, data []byte) (*types.Transaction, error) {
	return _Eventemitter.contract.Transact(opts, "emitDataLog3", topic1, topic2, topic3, data)
}

// EmitDataLog3 is a paid mutator transaction binding the contract method 0xb3ac1c38.
//
// Solidity: function emitDataLog3(bytes32 topic1, bytes32 topic2, bytes32 topic3, bytes data) returns()
func (_Eventemitter *EventemitterSession) EmitDataLog3(topic1 [32]byte, topic2 [32]byte, topic3 [32]byte, data []byte) (*types.Transaction, error) {
	return _Eventemitter.Contract.EmitDataLog3(&_Eventemitter.TransactOpts, topic1, topic2, topic3, data)
}

// EmitDataLog3 is a paid mutator transaction binding the contract method 0xb3ac1c38.
//
// Solidity: function emitDataLog3(bytes32 topic1, bytes32 topic2, bytes32 topic3, bytes data) returns()
func (_Eventemitter *EventemitterTransactorSession) EmitDataLog3(topic1 [32]byte, topic2 [32]byte, topic3 [32]byte, data []byte) (*types.Transaction, error) {
	return _Eventemitter.Contract.EmitDataLog3(&_Eventemitter.TransactOpts, topic1, topic2, topic3, data)
}

// EmitDataLog4 is a paid mutator transaction binding the contract method 0xee288ce8.
//
// Solidity: function emitDataLog4(bytes32 topic1, bytes32 topic2, bytes32 topic3, bytes32 topic4, bytes data) returns()
func (_Eventemitter *EventemitterTransactor) EmitDataLog4(opts *bind.TransactOpts, topic1 [32]byte, topic2 [32]byte, topic3 [32]byte, topic4 [32]byte, data []byte) (*types.Transaction, error) {
	return _Eventemitter.contract.Transact(opts, "emitDataLog4", topic1, topic2, topic3, topic4, data)
}

// EmitDataLog4 is a paid mutator transaction binding the contract method 0xee288ce8.
//
// Solidity: function emitDataLog4(bytes32 topic1, bytes32 topic2, bytes32 topic3, bytes32 topic4, bytes data) returns()
func (_Eventemitter *EventemitterSession) EmitDataLog4(topic1 [32]byte, topic2 [32]byte, topic3 [32]byte, topic4 [32]byte, data []byte) (*types.Transaction, error) {
	return _Eventemitter.Contract.EmitDataLog4(&_Eventemitter.TransactOpts, topic1, topic2, topic3, topic4, data)
}

// EmitDataLog4 is a paid mutator transaction binding the contract method 0xee288ce8.
//
// Solidity: function emitDataLog4(bytes32 topic1, bytes32 topic2, bytes32 topic3, bytes32 topic4, bytes data) returns()
func (_Eventemitter *EventemitterTransactorSession) EmitDataLog4(topic1 [32]byte, topic2 [32]byte, topic3 [32]byte, topic4 [32]byte, data []byte) (*types.Transaction, error) {
	return _Eventemitter.Contract.EmitDataLog4(&_Eventemitter.TransactOpts, topic1, topic2, topic3, topic4, data)
}

// EmitEventLog is a paid mutator transaction binding the contract method 0x906c49f6.
//
// Solidity: function emitEventLog(string eventName, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData) returns()
func (_Eventemitter *EventemitterTransactor) EmitEventLog(opts *bind.TransactOpts, eventName string, eventData EventUtilsEventLogData) (*types.Transaction, error) {
	return _Eventemitter.contract.Transact(opts, "emitEventLog", eventName, eventData)
}

// EmitEventLog is a paid mutator transaction binding the contract method 0x906c49f6.
//
// Solidity: function emitEventLog(string eventName, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData) returns()
func (_Eventemitter *EventemitterSession) EmitEventLog(eventName string, eventData EventUtilsEventLogData) (*types.Transaction, error) {
	return _Eventemitter.Contract.EmitEventLog(&_Eventemitter.TransactOpts, eventName, eventData)
}

// EmitEventLog is a paid mutator transaction binding the contract method 0x906c49f6.
//
// Solidity: function emitEventLog(string eventName, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData) returns()
func (_Eventemitter *EventemitterTransactorSession) EmitEventLog(eventName string, eventData EventUtilsEventLogData) (*types.Transaction, error) {
	return _Eventemitter.Contract.EmitEventLog(&_Eventemitter.TransactOpts, eventName, eventData)
}

// EmitEventLog1 is a paid mutator transaction binding the contract method 0x24de01e4.
//
// Solidity: function emitEventLog1(string eventName, bytes32 topic1, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData) returns()
func (_Eventemitter *EventemitterTransactor) EmitEventLog1(opts *bind.TransactOpts, eventName string, topic1 [32]byte, eventData EventUtilsEventLogData) (*types.Transaction, error) {
	return _Eventemitter.contract.Transact(opts, "emitEventLog1", eventName, topic1, eventData)
}

// EmitEventLog1 is a paid mutator transaction binding the contract method 0x24de01e4.
//
// Solidity: function emitEventLog1(string eventName, bytes32 topic1, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData) returns()
func (_Eventemitter *EventemitterSession) EmitEventLog1(eventName string, topic1 [32]byte, eventData EventUtilsEventLogData) (*types.Transaction, error) {
	return _Eventemitter.Contract.EmitEventLog1(&_Eventemitter.TransactOpts, eventName, topic1, eventData)
}

// EmitEventLog1 is a paid mutator transaction binding the contract method 0x24de01e4.
//
// Solidity: function emitEventLog1(string eventName, bytes32 topic1, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData) returns()
func (_Eventemitter *EventemitterTransactorSession) EmitEventLog1(eventName string, topic1 [32]byte, eventData EventUtilsEventLogData) (*types.Transaction, error) {
	return _Eventemitter.Contract.EmitEventLog1(&_Eventemitter.TransactOpts, eventName, topic1, eventData)
}

// EmitEventLog2 is a paid mutator transaction binding the contract method 0x63d16363.
//
// Solidity: function emitEventLog2(string eventName, bytes32 topic1, bytes32 topic2, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData) returns()
func (_Eventemitter *EventemitterTransactor) EmitEventLog2(opts *bind.TransactOpts, eventName string, topic1 [32]byte, topic2 [32]byte, eventData EventUtilsEventLogData) (*types.Transaction, error) {
	return _Eventemitter.contract.Transact(opts, "emitEventLog2", eventName, topic1, topic2, eventData)
}

// EmitEventLog2 is a paid mutator transaction binding the contract method 0x63d16363.
//
// Solidity: function emitEventLog2(string eventName, bytes32 topic1, bytes32 topic2, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData) returns()
func (_Eventemitter *EventemitterSession) EmitEventLog2(eventName string, topic1 [32]byte, topic2 [32]byte, eventData EventUtilsEventLogData) (*types.Transaction, error) {
	return _Eventemitter.Contract.EmitEventLog2(&_Eventemitter.TransactOpts, eventName, topic1, topic2, eventData)
}

// EmitEventLog2 is a paid mutator transaction binding the contract method 0x63d16363.
//
// Solidity: function emitEventLog2(string eventName, bytes32 topic1, bytes32 topic2, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData) returns()
func (_Eventemitter *EventemitterTransactorSession) EmitEventLog2(eventName string, topic1 [32]byte, topic2 [32]byte, eventData EventUtilsEventLogData) (*types.Transaction, error) {
	return _Eventemitter.Contract.EmitEventLog2(&_Eventemitter.TransactOpts, eventName, topic1, topic2, eventData)
}

// EventemitterEventLogIterator is returned from FilterEventLog and is used to iterate over the raw logs and unpacked data for EventLog events raised by the Eventemitter contract.
type EventemitterEventLogIterator struct {
	Event *EventemitterEventLog // Event containing the contract specifics and raw log

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
func (it *EventemitterEventLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventemitterEventLog)
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
		it.Event = new(EventemitterEventLog)
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
func (it *EventemitterEventLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventemitterEventLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventemitterEventLog represents a EventLog event raised by the Eventemitter contract.
type EventemitterEventLog struct {
	MsgSender     common.Address
	EventName     string
	EventNameHash common.Hash
	EventData     EventUtilsEventLogData
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEventLog is a free log retrieval operation binding the contract event 0x7e3bde2ba7aca4a8499608ca57f3b0c1c1c93ace63ffd3741a9fab204146fc9a.
//
// Solidity: event EventLog(address msgSender, string eventName, string indexed eventNameHash, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData)
func (_Eventemitter *EventemitterFilterer) FilterEventLog(opts *bind.FilterOpts, eventNameHash []string) (*EventemitterEventLogIterator, error) {

	var eventNameHashRule []interface{}
	for _, eventNameHashItem := range eventNameHash {
		eventNameHashRule = append(eventNameHashRule, eventNameHashItem)
	}

	logs, sub, err := _Eventemitter.contract.FilterLogs(opts, "EventLog", eventNameHashRule)
	if err != nil {
		return nil, err
	}
	return &EventemitterEventLogIterator{contract: _Eventemitter.contract, event: "EventLog", logs: logs, sub: sub}, nil
}

// WatchEventLog is a free log subscription operation binding the contract event 0x7e3bde2ba7aca4a8499608ca57f3b0c1c1c93ace63ffd3741a9fab204146fc9a.
//
// Solidity: event EventLog(address msgSender, string eventName, string indexed eventNameHash, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData)
func (_Eventemitter *EventemitterFilterer) WatchEventLog(opts *bind.WatchOpts, sink chan<- *EventemitterEventLog, eventNameHash []string) (event.Subscription, error) {

	var eventNameHashRule []interface{}
	for _, eventNameHashItem := range eventNameHash {
		eventNameHashRule = append(eventNameHashRule, eventNameHashItem)
	}

	logs, sub, err := _Eventemitter.contract.WatchLogs(opts, "EventLog", eventNameHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventemitterEventLog)
				if err := _Eventemitter.contract.UnpackLog(event, "EventLog", log); err != nil {
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

// ParseEventLog is a log parse operation binding the contract event 0x7e3bde2ba7aca4a8499608ca57f3b0c1c1c93ace63ffd3741a9fab204146fc9a.
//
// Solidity: event EventLog(address msgSender, string eventName, string indexed eventNameHash, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData)
func (_Eventemitter *EventemitterFilterer) ParseEventLog(log types.Log) (*EventemitterEventLog, error) {
	event := new(EventemitterEventLog)
	if err := _Eventemitter.contract.UnpackLog(event, "EventLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventemitterEventLog1Iterator is returned from FilterEventLog1 and is used to iterate over the raw logs and unpacked data for EventLog1 events raised by the Eventemitter contract.
type EventemitterEventLog1Iterator struct {
	Event *EventemitterEventLog1 // Event containing the contract specifics and raw log

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
func (it *EventemitterEventLog1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventemitterEventLog1)
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
		it.Event = new(EventemitterEventLog1)
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
func (it *EventemitterEventLog1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventemitterEventLog1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventemitterEventLog1 represents a EventLog1 event raised by the Eventemitter contract.
type EventemitterEventLog1 struct {
	MsgSender     common.Address
	EventName     string
	EventNameHash common.Hash
	Topic1        [32]byte
	EventData     EventUtilsEventLogData
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEventLog1 is a free log retrieval operation binding the contract event 0x137a44067c8961cd7e1d876f4754a5a3a75989b4552f1843fc69c3b372def160.
//
// Solidity: event EventLog1(address msgSender, string eventName, string indexed eventNameHash, bytes32 indexed topic1, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData)
func (_Eventemitter *EventemitterFilterer) FilterEventLog1(opts *bind.FilterOpts, eventNameHash []string, topic1 [][32]byte) (*EventemitterEventLog1Iterator, error) {

	var eventNameHashRule []interface{}
	for _, eventNameHashItem := range eventNameHash {
		eventNameHashRule = append(eventNameHashRule, eventNameHashItem)
	}
	var topic1Rule []interface{}
	for _, topic1Item := range topic1 {
		topic1Rule = append(topic1Rule, topic1Item)
	}

	logs, sub, err := _Eventemitter.contract.FilterLogs(opts, "EventLog1", eventNameHashRule, topic1Rule)
	if err != nil {
		return nil, err
	}
	return &EventemitterEventLog1Iterator{contract: _Eventemitter.contract, event: "EventLog1", logs: logs, sub: sub}, nil
}

// WatchEventLog1 is a free log subscription operation binding the contract event 0x137a44067c8961cd7e1d876f4754a5a3a75989b4552f1843fc69c3b372def160.
//
// Solidity: event EventLog1(address msgSender, string eventName, string indexed eventNameHash, bytes32 indexed topic1, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData)
func (_Eventemitter *EventemitterFilterer) WatchEventLog1(opts *bind.WatchOpts, sink chan<- *EventemitterEventLog1, eventNameHash []string, topic1 [][32]byte) (event.Subscription, error) {

	var eventNameHashRule []interface{}
	for _, eventNameHashItem := range eventNameHash {
		eventNameHashRule = append(eventNameHashRule, eventNameHashItem)
	}
	var topic1Rule []interface{}
	for _, topic1Item := range topic1 {
		topic1Rule = append(topic1Rule, topic1Item)
	}

	logs, sub, err := _Eventemitter.contract.WatchLogs(opts, "EventLog1", eventNameHashRule, topic1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventemitterEventLog1)
				if err := _Eventemitter.contract.UnpackLog(event, "EventLog1", log); err != nil {
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

// ParseEventLog1 is a log parse operation binding the contract event 0x137a44067c8961cd7e1d876f4754a5a3a75989b4552f1843fc69c3b372def160.
//
// Solidity: event EventLog1(address msgSender, string eventName, string indexed eventNameHash, bytes32 indexed topic1, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData)
func (_Eventemitter *EventemitterFilterer) ParseEventLog1(log types.Log) (*EventemitterEventLog1, error) {
	event := new(EventemitterEventLog1)
	if err := _Eventemitter.contract.UnpackLog(event, "EventLog1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventemitterEventLog2Iterator is returned from FilterEventLog2 and is used to iterate over the raw logs and unpacked data for EventLog2 events raised by the Eventemitter contract.
type EventemitterEventLog2Iterator struct {
	Event *EventemitterEventLog2 // Event containing the contract specifics and raw log

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
func (it *EventemitterEventLog2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventemitterEventLog2)
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
		it.Event = new(EventemitterEventLog2)
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
func (it *EventemitterEventLog2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventemitterEventLog2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventemitterEventLog2 represents a EventLog2 event raised by the Eventemitter contract.
type EventemitterEventLog2 struct {
	MsgSender     common.Address
	EventName     string
	EventNameHash common.Hash
	Topic1        [32]byte
	Topic2        [32]byte
	EventData     EventUtilsEventLogData
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEventLog2 is a free log retrieval operation binding the contract event 0x468a25a7ba624ceea6e540ad6f49171b52495b648417ae91bca21676d8a24dc5.
//
// Solidity: event EventLog2(address msgSender, string eventName, string indexed eventNameHash, bytes32 indexed topic1, bytes32 indexed topic2, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData)
func (_Eventemitter *EventemitterFilterer) FilterEventLog2(opts *bind.FilterOpts, eventNameHash []string, topic1 [][32]byte, topic2 [][32]byte) (*EventemitterEventLog2Iterator, error) {

	var eventNameHashRule []interface{}
	for _, eventNameHashItem := range eventNameHash {
		eventNameHashRule = append(eventNameHashRule, eventNameHashItem)
	}
	var topic1Rule []interface{}
	for _, topic1Item := range topic1 {
		topic1Rule = append(topic1Rule, topic1Item)
	}
	var topic2Rule []interface{}
	for _, topic2Item := range topic2 {
		topic2Rule = append(topic2Rule, topic2Item)
	}

	logs, sub, err := _Eventemitter.contract.FilterLogs(opts, "EventLog2", eventNameHashRule, topic1Rule, topic2Rule)
	if err != nil {
		return nil, err
	}
	return &EventemitterEventLog2Iterator{contract: _Eventemitter.contract, event: "EventLog2", logs: logs, sub: sub}, nil
}

// WatchEventLog2 is a free log subscription operation binding the contract event 0x468a25a7ba624ceea6e540ad6f49171b52495b648417ae91bca21676d8a24dc5.
//
// Solidity: event EventLog2(address msgSender, string eventName, string indexed eventNameHash, bytes32 indexed topic1, bytes32 indexed topic2, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData)
func (_Eventemitter *EventemitterFilterer) WatchEventLog2(opts *bind.WatchOpts, sink chan<- *EventemitterEventLog2, eventNameHash []string, topic1 [][32]byte, topic2 [][32]byte) (event.Subscription, error) {

	var eventNameHashRule []interface{}
	for _, eventNameHashItem := range eventNameHash {
		eventNameHashRule = append(eventNameHashRule, eventNameHashItem)
	}
	var topic1Rule []interface{}
	for _, topic1Item := range topic1 {
		topic1Rule = append(topic1Rule, topic1Item)
	}
	var topic2Rule []interface{}
	for _, topic2Item := range topic2 {
		topic2Rule = append(topic2Rule, topic2Item)
	}

	logs, sub, err := _Eventemitter.contract.WatchLogs(opts, "EventLog2", eventNameHashRule, topic1Rule, topic2Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventemitterEventLog2)
				if err := _Eventemitter.contract.UnpackLog(event, "EventLog2", log); err != nil {
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

// ParseEventLog2 is a log parse operation binding the contract event 0x468a25a7ba624ceea6e540ad6f49171b52495b648417ae91bca21676d8a24dc5.
//
// Solidity: event EventLog2(address msgSender, string eventName, string indexed eventNameHash, bytes32 indexed topic1, bytes32 indexed topic2, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData)
func (_Eventemitter *EventemitterFilterer) ParseEventLog2(log types.Log) (*EventemitterEventLog2, error) {
	event := new(EventemitterEventLog2)
	if err := _Eventemitter.contract.UnpackLog(event, "EventLog2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
