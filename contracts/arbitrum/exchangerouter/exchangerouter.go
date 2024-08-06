// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package exchangerouter

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

// DepositUtilsCreateDepositParams is an auto generated low-level Go binding around an user-defined struct.
type DepositUtilsCreateDepositParams struct {
	Receiver                common.Address
	CallbackContract        common.Address
	UiFeeReceiver           common.Address
	Market                  common.Address
	InitialLongToken        common.Address
	InitialShortToken       common.Address
	LongTokenSwapPath       []common.Address
	ShortTokenSwapPath      []common.Address
	MinMarketTokens         *big.Int
	ShouldUnwrapNativeToken bool
	ExecutionFee            *big.Int
	CallbackGasLimit        *big.Int
}

// IBaseOrderUtilsCreateOrderParams is an auto generated low-level Go binding around an user-defined struct.
type IBaseOrderUtilsCreateOrderParams struct {
	Addresses                IBaseOrderUtilsCreateOrderParamsAddresses
	Numbers                  IBaseOrderUtilsCreateOrderParamsNumbers
	OrderType                uint8
	DecreasePositionSwapType uint8
	IsLong                   bool
	ShouldUnwrapNativeToken  bool
	AutoCancel               bool
	ReferralCode             [32]byte
}

// IBaseOrderUtilsCreateOrderParamsAddresses is an auto generated low-level Go binding around an user-defined struct.
type IBaseOrderUtilsCreateOrderParamsAddresses struct {
	Receiver               common.Address
	CancellationReceiver   common.Address
	CallbackContract       common.Address
	UiFeeReceiver          common.Address
	Market                 common.Address
	InitialCollateralToken common.Address
	SwapPath               []common.Address
}

// IBaseOrderUtilsCreateOrderParamsNumbers is an auto generated low-level Go binding around an user-defined struct.
type IBaseOrderUtilsCreateOrderParamsNumbers struct {
	SizeDeltaUsd                 *big.Int
	InitialCollateralDeltaAmount *big.Int
	TriggerPrice                 *big.Int
	AcceptablePrice              *big.Int
	ExecutionFee                 *big.Int
	CallbackGasLimit             *big.Int
	MinOutputAmount              *big.Int
}

// OracleUtilsSetPricesParams is an auto generated low-level Go binding around an user-defined struct.
type OracleUtilsSetPricesParams struct {
	Tokens    []common.Address
	Providers []common.Address
	Data      [][]byte
}

// OracleUtilsSimulatePricesParams is an auto generated low-level Go binding around an user-defined struct.
type OracleUtilsSimulatePricesParams struct {
	PrimaryTokens []common.Address
	PrimaryPrices []PriceProps
	MinTimestamp  *big.Int
	MaxTimestamp  *big.Int
}

// PriceProps is an auto generated low-level Go binding around an user-defined struct.
type PriceProps struct {
	Min *big.Int
	Max *big.Int
}

// ShiftUtilsCreateShiftParams is an auto generated low-level Go binding around an user-defined struct.
type ShiftUtilsCreateShiftParams struct {
	Receiver         common.Address
	CallbackContract common.Address
	UiFeeReceiver    common.Address
	FromMarket       common.Address
	ToMarket         common.Address
	MinMarketTokens  *big.Int
	ExecutionFee     *big.Int
	CallbackGasLimit *big.Int
}

// WithdrawalUtilsCreateWithdrawalParams is an auto generated low-level Go binding around an user-defined struct.
type WithdrawalUtilsCreateWithdrawalParams struct {
	Receiver                common.Address
	CallbackContract        common.Address
	UiFeeReceiver           common.Address
	Market                  common.Address
	LongTokenSwapPath       []common.Address
	ShortTokenSwapPath      []common.Address
	MinLongTokenAmount      *big.Int
	MinShortTokenAmount     *big.Int
	ShouldUnwrapNativeToken bool
	ExecutionFee            *big.Int
	CallbackGasLimit        *big.Int
}

// ExchangerouterMetaData contains all meta data concerning the Exchangerouter contract.
var ExchangerouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractRouter\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"contractRoleStore\",\"name\":\"_roleStore\",\"type\":\"address\"},{\"internalType\":\"contractDataStore\",\"name\":\"_dataStore\",\"type\":\"address\"},{\"internalType\":\"contractEventEmitter\",\"name\":\"_eventEmitter\",\"type\":\"address\"},{\"internalType\":\"contractIDepositHandler\",\"name\":\"_depositHandler\",\"type\":\"address\"},{\"internalType\":\"contractIWithdrawalHandler\",\"name\":\"_withdrawalHandler\",\"type\":\"address\"},{\"internalType\":\"contractIShiftHandler\",\"name\":\"_shiftHandler\",\"type\":\"address\"},{\"internalType\":\"contractIOrderHandler\",\"name\":\"_orderHandler\",\"type\":\"address\"},{\"internalType\":\"contractIExternalHandler\",\"name\":\"_externalHandler\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"adjustedClaimableAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimedAmount\",\"type\":\"uint256\"}],\"name\":\"CollateralAlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"DisabledFeature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"DisabledMarket\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"EmptyAddressInMarketTokenBalanceValidation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyDeposit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyHoldingAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyMarket\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"EmptyTokenTranferGasLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"marketsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensLength\",\"type\":\"uint256\"}],\"name\":\"InvalidClaimAffiliateRewardsInput\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"marketsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeKeysLength\",\"type\":\"uint256\"}],\"name\":\"InvalidClaimCollateralInput\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"marketsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensLength\",\"type\":\"uint256\"}],\"name\":\"InvalidClaimFundingFeesInput\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"marketsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensLength\",\"type\":\"uint256\"}],\"name\":\"InvalidClaimUiFeesInput\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InvalidClaimableFactor\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinBalance\",\"type\":\"uint256\"}],\"name\":\"InvalidMarketTokenBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimableFundingFeeAmount\",\"type\":\"uint256\"}],\"name\":\"InvalidMarketTokenBalanceForClaimableFunding\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"InvalidMarketTokenBalanceForCollateralAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uiFeeFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxUiFeeFactor\",\"type\":\"uint256\"}],\"name\":\"InvalidUiFeeFactor\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenTransferError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"cancelDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"cancelShift\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"cancelWithdrawal\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"markets\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"claimAffiliateRewards\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"markets\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"timeKeys\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"claimCollateral\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"markets\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"claimFundingFees\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"markets\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"claimUiFees\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callbackContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uiFeeReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialLongToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialShortToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"longTokenSwapPath\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"shortTokenSwapPath\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"minMarketTokens\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"shouldUnwrapNativeToken\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callbackGasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structDepositUtils.CreateDepositParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"createDeposit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cancellationReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callbackContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uiFeeReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialCollateralToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"swapPath\",\"type\":\"address[]\"}],\"internalType\":\"structIBaseOrderUtils.CreateOrderParamsAddresses\",\"name\":\"addresses\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sizeDeltaUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialCollateralDeltaAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"triggerPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callbackGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minOutputAmount\",\"type\":\"uint256\"}],\"internalType\":\"structIBaseOrderUtils.CreateOrderParamsNumbers\",\"name\":\"numbers\",\"type\":\"tuple\"},{\"internalType\":\"enumOrder.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"enumOrder.DecreasePositionSwapType\",\"name\":\"decreasePositionSwapType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldUnwrapNativeToken\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"autoCancel\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"referralCode\",\"type\":\"bytes32\"}],\"internalType\":\"structIBaseOrderUtils.CreateOrderParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"createOrder\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callbackContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uiFeeReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fromMarket\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toMarket\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minMarketTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callbackGasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structShiftUtils.CreateShiftParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"createShift\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callbackContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uiFeeReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"longTokenSwapPath\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"shortTokenSwapPath\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"minLongTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minShortTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"shouldUnwrapNativeToken\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callbackGasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structWithdrawalUtils.CreateWithdrawalParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"createWithdrawal\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataStore\",\"outputs\":[{\"internalType\":\"contractDataStore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositHandler\",\"outputs\":[{\"internalType\":\"contractIDepositHandler\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eventEmitter\",\"outputs\":[{\"internalType\":\"contractEventEmitter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callbackContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uiFeeReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"longTokenSwapPath\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"shortTokenSwapPath\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"minLongTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minShortTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"shouldUnwrapNativeToken\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callbackGasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structWithdrawalUtils.CreateWithdrawalParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"providers\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"internalType\":\"structOracleUtils.SetPricesParams\",\"name\":\"oracleParams\",\"type\":\"tuple\"}],\"name\":\"executeAtomicWithdrawal\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"externalHandler\",\"outputs\":[{\"internalType\":\"contractIExternalHandler\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"externalCallTargets\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"externalCallDataList\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"refundTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"refundReceivers\",\"type\":\"address[]\"}],\"name\":\"makeExternalCalls\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"orderHandler\",\"outputs\":[{\"internalType\":\"contractIOrderHandler\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roleStore\",\"outputs\":[{\"internalType\":\"contractRoleStore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"contractRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendNativeToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendWnt\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callbackContract\",\"type\":\"address\"}],\"name\":\"setSavedCallbackContract\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uiFeeFactor\",\"type\":\"uint256\"}],\"name\":\"setUiFeeFactor\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shiftHandler\",\"outputs\":[{\"internalType\":\"contractIShiftHandler\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"primaryTokens\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props[]\",\"name\":\"primaryPrices\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"minTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structOracleUtils.SimulatePricesParams\",\"name\":\"simulatedOracleParams\",\"type\":\"tuple\"}],\"name\":\"simulateExecuteDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"primaryTokens\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props[]\",\"name\":\"primaryPrices\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"minTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structOracleUtils.SimulatePricesParams\",\"name\":\"simulatedOracleParams\",\"type\":\"tuple\"}],\"name\":\"simulateExecuteOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"primaryTokens\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props[]\",\"name\":\"primaryPrices\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"minTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structOracleUtils.SimulatePricesParams\",\"name\":\"simulatedOracleParams\",\"type\":\"tuple\"}],\"name\":\"simulateExecuteShift\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"primaryTokens\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structPrice.Props[]\",\"name\":\"primaryPrices\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"minTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structOracleUtils.SimulatePricesParams\",\"name\":\"simulatedOracleParams\",\"type\":\"tuple\"},{\"internalType\":\"enumISwapPricingUtils.SwapPricingType\",\"name\":\"swapPricingType\",\"type\":\"uint8\"}],\"name\":\"simulateExecuteWithdrawal\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sizeDeltaUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"triggerPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minOutputAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"autoCancel\",\"type\":\"bool\"}],\"name\":\"updateOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalHandler\",\"outputs\":[{\"internalType\":\"contractIWithdrawalHandler\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ExchangerouterABI is the input ABI used to generate the binding from.
// Deprecated: Use ExchangerouterMetaData.ABI instead.
var ExchangerouterABI = ExchangerouterMetaData.ABI

// Exchangerouter is an auto generated Go binding around an Ethereum contract.
type Exchangerouter struct {
	ExchangerouterCaller     // Read-only binding to the contract
	ExchangerouterTransactor // Write-only binding to the contract
	ExchangerouterFilterer   // Log filterer for contract events
}

// ExchangerouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangerouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangerouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangerouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangerouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangerouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangerouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangerouterSession struct {
	Contract     *Exchangerouter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangerouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangerouterCallerSession struct {
	Contract *ExchangerouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ExchangerouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangerouterTransactorSession struct {
	Contract     *ExchangerouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ExchangerouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangerouterRaw struct {
	Contract *Exchangerouter // Generic contract binding to access the raw methods on
}

// ExchangerouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangerouterCallerRaw struct {
	Contract *ExchangerouterCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangerouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangerouterTransactorRaw struct {
	Contract *ExchangerouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchangerouter creates a new instance of Exchangerouter, bound to a specific deployed contract.
func NewExchangerouter(address common.Address, backend bind.ContractBackend) (*Exchangerouter, error) {
	contract, err := bindExchangerouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Exchangerouter{ExchangerouterCaller: ExchangerouterCaller{contract: contract}, ExchangerouterTransactor: ExchangerouterTransactor{contract: contract}, ExchangerouterFilterer: ExchangerouterFilterer{contract: contract}}, nil
}

// NewExchangerouterCaller creates a new read-only instance of Exchangerouter, bound to a specific deployed contract.
func NewExchangerouterCaller(address common.Address, caller bind.ContractCaller) (*ExchangerouterCaller, error) {
	contract, err := bindExchangerouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangerouterCaller{contract: contract}, nil
}

// NewExchangerouterTransactor creates a new write-only instance of Exchangerouter, bound to a specific deployed contract.
func NewExchangerouterTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangerouterTransactor, error) {
	contract, err := bindExchangerouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangerouterTransactor{contract: contract}, nil
}

// NewExchangerouterFilterer creates a new log filterer instance of Exchangerouter, bound to a specific deployed contract.
func NewExchangerouterFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangerouterFilterer, error) {
	contract, err := bindExchangerouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangerouterFilterer{contract: contract}, nil
}

// bindExchangerouter binds a generic wrapper to an already deployed contract.
func bindExchangerouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExchangerouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchangerouter *ExchangerouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exchangerouter.Contract.ExchangerouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchangerouter *ExchangerouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchangerouter.Contract.ExchangerouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchangerouter *ExchangerouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchangerouter.Contract.ExchangerouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchangerouter *ExchangerouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exchangerouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchangerouter *ExchangerouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchangerouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchangerouter *ExchangerouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchangerouter.Contract.contract.Transact(opts, method, params...)
}

// DataStore is a free data retrieval call binding the contract method 0x660d0d67.
//
// Solidity: function dataStore() view returns(address)
func (_Exchangerouter *ExchangerouterCaller) DataStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchangerouter.contract.Call(opts, &out, "dataStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataStore is a free data retrieval call binding the contract method 0x660d0d67.
//
// Solidity: function dataStore() view returns(address)
func (_Exchangerouter *ExchangerouterSession) DataStore() (common.Address, error) {
	return _Exchangerouter.Contract.DataStore(&_Exchangerouter.CallOpts)
}

// DataStore is a free data retrieval call binding the contract method 0x660d0d67.
//
// Solidity: function dataStore() view returns(address)
func (_Exchangerouter *ExchangerouterCallerSession) DataStore() (common.Address, error) {
	return _Exchangerouter.Contract.DataStore(&_Exchangerouter.CallOpts)
}

// DepositHandler is a free data retrieval call binding the contract method 0x9c8b2cfb.
//
// Solidity: function depositHandler() view returns(address)
func (_Exchangerouter *ExchangerouterCaller) DepositHandler(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchangerouter.contract.Call(opts, &out, "depositHandler")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositHandler is a free data retrieval call binding the contract method 0x9c8b2cfb.
//
// Solidity: function depositHandler() view returns(address)
func (_Exchangerouter *ExchangerouterSession) DepositHandler() (common.Address, error) {
	return _Exchangerouter.Contract.DepositHandler(&_Exchangerouter.CallOpts)
}

// DepositHandler is a free data retrieval call binding the contract method 0x9c8b2cfb.
//
// Solidity: function depositHandler() view returns(address)
func (_Exchangerouter *ExchangerouterCallerSession) DepositHandler() (common.Address, error) {
	return _Exchangerouter.Contract.DepositHandler(&_Exchangerouter.CallOpts)
}

// EventEmitter is a free data retrieval call binding the contract method 0x9ff78c30.
//
// Solidity: function eventEmitter() view returns(address)
func (_Exchangerouter *ExchangerouterCaller) EventEmitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchangerouter.contract.Call(opts, &out, "eventEmitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EventEmitter is a free data retrieval call binding the contract method 0x9ff78c30.
//
// Solidity: function eventEmitter() view returns(address)
func (_Exchangerouter *ExchangerouterSession) EventEmitter() (common.Address, error) {
	return _Exchangerouter.Contract.EventEmitter(&_Exchangerouter.CallOpts)
}

// EventEmitter is a free data retrieval call binding the contract method 0x9ff78c30.
//
// Solidity: function eventEmitter() view returns(address)
func (_Exchangerouter *ExchangerouterCallerSession) EventEmitter() (common.Address, error) {
	return _Exchangerouter.Contract.EventEmitter(&_Exchangerouter.CallOpts)
}

// ExternalHandler is a free data retrieval call binding the contract method 0x2e944bd6.
//
// Solidity: function externalHandler() view returns(address)
func (_Exchangerouter *ExchangerouterCaller) ExternalHandler(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchangerouter.contract.Call(opts, &out, "externalHandler")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExternalHandler is a free data retrieval call binding the contract method 0x2e944bd6.
//
// Solidity: function externalHandler() view returns(address)
func (_Exchangerouter *ExchangerouterSession) ExternalHandler() (common.Address, error) {
	return _Exchangerouter.Contract.ExternalHandler(&_Exchangerouter.CallOpts)
}

// ExternalHandler is a free data retrieval call binding the contract method 0x2e944bd6.
//
// Solidity: function externalHandler() view returns(address)
func (_Exchangerouter *ExchangerouterCallerSession) ExternalHandler() (common.Address, error) {
	return _Exchangerouter.Contract.ExternalHandler(&_Exchangerouter.CallOpts)
}

// OrderHandler is a free data retrieval call binding the contract method 0xb5848305.
//
// Solidity: function orderHandler() view returns(address)
func (_Exchangerouter *ExchangerouterCaller) OrderHandler(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchangerouter.contract.Call(opts, &out, "orderHandler")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OrderHandler is a free data retrieval call binding the contract method 0xb5848305.
//
// Solidity: function orderHandler() view returns(address)
func (_Exchangerouter *ExchangerouterSession) OrderHandler() (common.Address, error) {
	return _Exchangerouter.Contract.OrderHandler(&_Exchangerouter.CallOpts)
}

// OrderHandler is a free data retrieval call binding the contract method 0xb5848305.
//
// Solidity: function orderHandler() view returns(address)
func (_Exchangerouter *ExchangerouterCallerSession) OrderHandler() (common.Address, error) {
	return _Exchangerouter.Contract.OrderHandler(&_Exchangerouter.CallOpts)
}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_Exchangerouter *ExchangerouterCaller) RoleStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchangerouter.contract.Call(opts, &out, "roleStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_Exchangerouter *ExchangerouterSession) RoleStore() (common.Address, error) {
	return _Exchangerouter.Contract.RoleStore(&_Exchangerouter.CallOpts)
}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_Exchangerouter *ExchangerouterCallerSession) RoleStore() (common.Address, error) {
	return _Exchangerouter.Contract.RoleStore(&_Exchangerouter.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Exchangerouter *ExchangerouterCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchangerouter.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Exchangerouter *ExchangerouterSession) Router() (common.Address, error) {
	return _Exchangerouter.Contract.Router(&_Exchangerouter.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Exchangerouter *ExchangerouterCallerSession) Router() (common.Address, error) {
	return _Exchangerouter.Contract.Router(&_Exchangerouter.CallOpts)
}

// ShiftHandler is a free data retrieval call binding the contract method 0xe65c9ae1.
//
// Solidity: function shiftHandler() view returns(address)
func (_Exchangerouter *ExchangerouterCaller) ShiftHandler(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchangerouter.contract.Call(opts, &out, "shiftHandler")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ShiftHandler is a free data retrieval call binding the contract method 0xe65c9ae1.
//
// Solidity: function shiftHandler() view returns(address)
func (_Exchangerouter *ExchangerouterSession) ShiftHandler() (common.Address, error) {
	return _Exchangerouter.Contract.ShiftHandler(&_Exchangerouter.CallOpts)
}

// ShiftHandler is a free data retrieval call binding the contract method 0xe65c9ae1.
//
// Solidity: function shiftHandler() view returns(address)
func (_Exchangerouter *ExchangerouterCallerSession) ShiftHandler() (common.Address, error) {
	return _Exchangerouter.Contract.ShiftHandler(&_Exchangerouter.CallOpts)
}

// WithdrawalHandler is a free data retrieval call binding the contract method 0x2c2f3c07.
//
// Solidity: function withdrawalHandler() view returns(address)
func (_Exchangerouter *ExchangerouterCaller) WithdrawalHandler(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchangerouter.contract.Call(opts, &out, "withdrawalHandler")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawalHandler is a free data retrieval call binding the contract method 0x2c2f3c07.
//
// Solidity: function withdrawalHandler() view returns(address)
func (_Exchangerouter *ExchangerouterSession) WithdrawalHandler() (common.Address, error) {
	return _Exchangerouter.Contract.WithdrawalHandler(&_Exchangerouter.CallOpts)
}

// WithdrawalHandler is a free data retrieval call binding the contract method 0x2c2f3c07.
//
// Solidity: function withdrawalHandler() view returns(address)
func (_Exchangerouter *ExchangerouterCallerSession) WithdrawalHandler() (common.Address, error) {
	return _Exchangerouter.Contract.WithdrawalHandler(&_Exchangerouter.CallOpts)
}

// CancelDeposit is a paid mutator transaction binding the contract method 0x31404484.
//
// Solidity: function cancelDeposit(bytes32 key) payable returns()
func (_Exchangerouter *ExchangerouterTransactor) CancelDeposit(opts *bind.TransactOpts, key [32]byte) (*types.Transaction, error) {
	return _Exchangerouter.contract.Transact(opts, "cancelDeposit", key)
}

// CancelDeposit is a paid mutator transaction binding the contract method 0x31404484.
//
// Solidity: function cancelDeposit(bytes32 key) payable returns()
func (_Exchangerouter *ExchangerouterSession) CancelDeposit(key [32]byte) (*types.Transaction, error) {
	return _Exchangerouter.Contract.CancelDeposit(&_Exchangerouter.TransactOpts, key)
}

// CancelDeposit is a paid mutator transaction binding the contract method 0x31404484.
//
// Solidity: function cancelDeposit(bytes32 key) payable returns()
func (_Exchangerouter *ExchangerouterTransactorSession) CancelDeposit(key [32]byte) (*types.Transaction, error) {
	return _Exchangerouter.Contract.CancelDeposit(&_Exchangerouter.TransactOpts, key)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x7489ec23.
//
// Solidity: function cancelOrder(bytes32 key) payable returns()
func (_Exchangerouter *ExchangerouterTransactor) CancelOrder(opts *bind.TransactOpts, key [32]byte) (*types.Transaction, error) {
	return _Exchangerouter.contract.Transact(opts, "cancelOrder", key)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x7489ec23.
//
// Solidity: function cancelOrder(bytes32 key) payable returns()
func (_Exchangerouter *ExchangerouterSession) CancelOrder(key [32]byte) (*types.Transaction, error) {
	return _Exchangerouter.Contract.CancelOrder(&_Exchangerouter.TransactOpts, key)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x7489ec23.
//
// Solidity: function cancelOrder(bytes32 key) payable returns()
func (_Exchangerouter *ExchangerouterTransactorSession) CancelOrder(key [32]byte) (*types.Transaction, error) {
	return _Exchangerouter.Contract.CancelOrder(&_Exchangerouter.TransactOpts, key)
}

// CancelShift is a paid mutator transaction binding the contract method 0x96be2898.
//
// Solidity: function cancelShift(bytes32 key) payable returns()
func (_Exchangerouter *ExchangerouterTransactor) CancelShift(opts *bind.TransactOpts, key [32]byte) (*types.Transaction, error) {
	return _Exchangerouter.contract.Transact(opts, "cancelShift", key)
}

// CancelShift is a paid mutator transaction binding the contract method 0x96be2898.
//
// Solidity: function cancelShift(bytes32 key) payable returns()
func (_Exchangerouter *ExchangerouterSession) CancelShift(key [32]byte) (*types.Transaction, error) {
	return _Exchangerouter.Contract.CancelShift(&_Exchangerouter.TransactOpts, key)
}

// CancelShift is a paid mutator transaction binding the contract method 0x96be2898.
//
// Solidity: function cancelShift(bytes32 key) payable returns()
func (_Exchangerouter *ExchangerouterTransactorSession) CancelShift(key [32]byte) (*types.Transaction, error) {
	return _Exchangerouter.Contract.CancelShift(&_Exchangerouter.TransactOpts, key)
}

// CancelWithdrawal is a paid mutator transaction binding the contract method 0x7213c5a0.
//
// Solidity: function cancelWithdrawal(bytes32 key) payable returns()
func (_Exchangerouter *ExchangerouterTransactor) CancelWithdrawal(opts *bind.TransactOpts, key [32]byte) (*types.Transaction, error) {
	return _Exchangerouter.contract.Transact(opts, "cancelWithdrawal", key)
}

// CancelWithdrawal is a paid mutator transaction binding the contract method 0x7213c5a0.
//
// Solidity: function cancelWithdrawal(bytes32 key) payable returns()
func (_Exchangerouter *ExchangerouterSession) CancelWithdrawal(key [32]byte) (*types.Transaction, error) {
	return _Exchangerouter.Contract.CancelWithdrawal(&_Exchangerouter.TransactOpts, key)
}

// CancelWithdrawal is a paid mutator transaction binding the contract method 0x7213c5a0.
//
// Solidity: function cancelWithdrawal(bytes32 key) payable returns()
func (_Exchangerouter *ExchangerouterTransactorSession) CancelWithdrawal(key [32]byte) (*types.Transaction, error) {
	return _Exchangerouter.Contract.CancelWithdrawal(&_Exchangerouter.TransactOpts, key)
}

// ClaimAffiliateRewards is a paid mutator transaction binding the contract method 0x49287a22.
//
// Solidity: function claimAffiliateRewards(address[] markets, address[] tokens, address receiver) payable returns(uint256[])
func (_Exchangerouter *ExchangerouterTransactor) ClaimAffiliateRewards(opts *bind.TransactOpts, markets []common.Address, tokens []common.Address, receiver common.Address) (*types.Transaction, error) {
	return _Exchangerouter.contract.Transact(opts, "claimAffiliateRewards", markets, tokens, receiver)
}

// ClaimAffiliateRewards is a paid mutator transaction binding the contract method 0x49287a22.
//
// Solidity: function claimAffiliateRewards(address[] markets, address[] tokens, address receiver) payable returns(uint256[])
func (_Exchangerouter *ExchangerouterSession) ClaimAffiliateRewards(markets []common.Address, tokens []common.Address, receiver common.Address) (*types.Transaction, error) {
	return _Exchangerouter.Contract.ClaimAffiliateRewards(&_Exchangerouter.TransactOpts, markets, tokens, receiver)
}

// ClaimAffiliateRewards is a paid mutator transaction binding the contract method 0x49287a22.
//
// Solidity: function claimAffiliateRewards(address[] markets, address[] tokens, address receiver) payable returns(uint256[])
func (_Exchangerouter *ExchangerouterTransactorSession) ClaimAffiliateRewards(markets []common.Address, tokens []common.Address, receiver common.Address) (*types.Transaction, error) {
	return _Exchangerouter.Contract.ClaimAffiliateRewards(&_Exchangerouter.TransactOpts, markets, tokens, receiver)
}

// ClaimCollateral is a paid mutator transaction binding the contract method 0xe9249b57.
//
// Solidity: function claimCollateral(address[] markets, address[] tokens, uint256[] timeKeys, address receiver) payable returns(uint256[])
func (_Exchangerouter *ExchangerouterTransactor) ClaimCollateral(opts *bind.TransactOpts, markets []common.Address, tokens []common.Address, timeKeys []*big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Exchangerouter.contract.Transact(opts, "claimCollateral", markets, tokens, timeKeys, receiver)
}

// ClaimCollateral is a paid mutator transaction binding the contract method 0xe9249b57.
//
// Solidity: function claimCollateral(address[] markets, address[] tokens, uint256[] timeKeys, address receiver) payable returns(uint256[])
func (_Exchangerouter *ExchangerouterSession) ClaimCollateral(markets []common.Address, tokens []common.Address, timeKeys []*big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Exchangerouter.Contract.ClaimCollateral(&_Exchangerouter.TransactOpts, markets, tokens, timeKeys, receiver)
}

// ClaimCollateral is a paid mutator transaction binding the contract method 0xe9249b57.
//
// Solidity: function claimCollateral(address[] markets, address[] tokens, uint256[] timeKeys, address receiver) payable returns(uint256[])
func (_Exchangerouter *ExchangerouterTransactorSession) ClaimCollateral(markets []common.Address, tokens []common.Address, timeKeys []*big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Exchangerouter.Contract.ClaimCollateral(&_Exchangerouter.TransactOpts, markets, tokens, timeKeys, receiver)
}

// ClaimFundingFees is a paid mutator transaction binding the contract method 0xc41b1ab3.
//
// Solidity: function claimFundingFees(address[] markets, address[] tokens, address receiver) payable returns(uint256[])
func (_Exchangerouter *ExchangerouterTransactor) ClaimFundingFees(opts *bind.TransactOpts, markets []common.Address, tokens []common.Address, receiver common.Address) (*types.Transaction, error) {
	return _Exchangerouter.contract.Transact(opts, "claimFundingFees", markets, tokens, receiver)
}

// ClaimFundingFees is a paid mutator transaction binding the contract method 0xc41b1ab3.
//
// Solidity: function claimFundingFees(address[] markets, address[] tokens, address receiver) payable returns(uint256[])
func (_Exchangerouter *ExchangerouterSession) ClaimFundingFees(markets []common.Address, tokens []common.Address, receiver common.Address) (*types.Transaction, error) {
	return _Exchangerouter.Contract.ClaimFundingFees(&_Exchangerouter.TransactOpts, markets, tokens, receiver)
}

// ClaimFundingFees is a paid mutator transaction binding the contract method 0xc41b1ab3.
//
// Solidity: function claimFundingFees(address[] markets, address[] tokens, address receiver) payable returns(uint256[])
func (_Exchangerouter *ExchangerouterTransactorSession) ClaimFundingFees(markets []common.Address, tokens []common.Address, receiver common.Address) (*types.Transaction, error) {
	return _Exchangerouter.Contract.ClaimFundingFees(&_Exchangerouter.TransactOpts, markets, tokens, receiver)
}

// ClaimUiFees is a paid mutator transaction binding the contract method 0x01a9cbb2.
//
// Solidity: function claimUiFees(address[] markets, address[] tokens, address receiver) payable returns(uint256[])
func (_Exchangerouter *ExchangerouterTransactor) ClaimUiFees(opts *bind.TransactOpts, markets []common.Address, tokens []common.Address, receiver common.Address) (*types.Transaction, error) {
	return _Exchangerouter.contract.Transact(opts, "claimUiFees", markets, tokens, receiver)
}

// ClaimUiFees is a paid mutator transaction binding the contract method 0x01a9cbb2.
//
// Solidity: function claimUiFees(address[] markets, address[] tokens, address receiver) payable returns(uint256[])
func (_Exchangerouter *ExchangerouterSession) ClaimUiFees(markets []common.Address, tokens []common.Address, receiver common.Address) (*types.Transaction, error) {
	return _Exchangerouter.Contract.ClaimUiFees(&_Exchangerouter.TransactOpts, markets, tokens, receiver)
}

// ClaimUiFees is a paid mutator transaction binding the contract method 0x01a9cbb2.
//
// Solidity: function claimUiFees(address[] markets, address[] tokens, address receiver) payable returns(uint256[])
func (_Exchangerouter *ExchangerouterTransactorSession) ClaimUiFees(markets []common.Address, tokens []common.Address, receiver common.Address) (*types.Transaction, error) {
	return _Exchangerouter.Contract.ClaimUiFees(&_Exchangerouter.TransactOpts, markets, tokens, receiver)
}

// CreateDeposit is a paid mutator transaction binding the contract method 0x5b4e9561.
//
// Solidity: function createDeposit((address,address,address,address,address,address,address[],address[],uint256,bool,uint256,uint256) params) payable returns(bytes32)
func (_Exchangerouter *ExchangerouterTransactor) CreateDeposit(opts *bind.TransactOpts, params DepositUtilsCreateDepositParams) (*types.Transaction, error) {
	return _Exchangerouter.contract.Transact(opts, "createDeposit", params)
}

// CreateDeposit is a paid mutator transaction binding the contract method 0x5b4e9561.
//
// Solidity: function createDeposit((address,address,address,address,address,address,address[],address[],uint256,bool,uint256,uint256) params) payable returns(bytes32)
func (_Exchangerouter *ExchangerouterSession) CreateDeposit(params DepositUtilsCreateDepositParams) (*types.Transaction, error) {
	return _Exchangerouter.Contract.CreateDeposit(&_Exchangerouter.TransactOpts, params)
}

// CreateDeposit is a paid mutator transaction binding the contract method 0x5b4e9561.
//
// Solidity: function createDeposit((address,address,address,address,address,address,address[],address[],uint256,bool,uint256,uint256) params) payable returns(bytes32)
func (_Exchangerouter *ExchangerouterTransactorSession) CreateDeposit(params DepositUtilsCreateDepositParams) (*types.Transaction, error) {
	return _Exchangerouter.Contract.CreateDeposit(&_Exchangerouter.TransactOpts, params)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x083cfcee.
//
// Solidity: function createOrder(((address,address,address,address,address,address,address[]),(uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint8,uint8,bool,bool,bool,bytes32) params) payable returns(bytes32)
func (_Exchangerouter *ExchangerouterTransactor) CreateOrder(opts *bind.TransactOpts, params IBaseOrderUtilsCreateOrderParams) (*types.Transaction, error) {
	return _Exchangerouter.contract.Transact(opts, "createOrder", params)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x083cfcee.
//
// Solidity: function createOrder(((address,address,address,address,address,address,address[]),(uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint8,uint8,bool,bool,bool,bytes32) params) payable returns(bytes32)
func (_Exchangerouter *ExchangerouterSession) CreateOrder(params IBaseOrderUtilsCreateOrderParams) (*types.Transaction, error) {
	return _Exchangerouter.Contract.CreateOrder(&_Exchangerouter.TransactOpts, params)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x083cfcee.
//
// Solidity: function createOrder(((address,address,address,address,address,address,address[]),(uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint8,uint8,bool,bool,bool,bytes32) params) payable returns(bytes32)
func (_Exchangerouter *ExchangerouterTransactorSession) CreateOrder(params IBaseOrderUtilsCreateOrderParams) (*types.Transaction, error) {
	return _Exchangerouter.Contract.CreateOrder(&_Exchangerouter.TransactOpts, params)
}

// CreateShift is a paid mutator transaction binding the contract method 0xb1f906b9.
//
// Solidity: function createShift((address,address,address,address,address,uint256,uint256,uint256) params) payable returns(bytes32)
func (_Exchangerouter *ExchangerouterTransactor) CreateShift(opts *bind.TransactOpts, params ShiftUtilsCreateShiftParams) (*types.Transaction, error) {
	return _Exchangerouter.contract.Transact(opts, "createShift", params)
}

// CreateShift is a paid mutator transaction binding the contract method 0xb1f906b9.
//
// Solidity: function createShift((address,address,address,address,address,uint256,uint256,uint256) params) payable returns(bytes32)
func (_Exchangerouter *ExchangerouterSession) CreateShift(params ShiftUtilsCreateShiftParams) (*types.Transaction, error) {
	return _Exchangerouter.Contract.CreateShift(&_Exchangerouter.TransactOpts, params)
}

// CreateShift is a paid mutator transaction binding the contract method 0xb1f906b9.
//
// Solidity: function createShift((address,address,address,address,address,uint256,uint256,uint256) params) payable returns(bytes32)
func (_Exchangerouter *ExchangerouterTransactorSession) CreateShift(params ShiftUtilsCreateShiftParams) (*types.Transaction, error) {
	return _Exchangerouter.Contract.CreateShift(&_Exchangerouter.TransactOpts, params)
}

// CreateWithdrawal is a paid mutator transaction binding the contract method 0xad23c5a1.
//
// Solidity: function createWithdrawal((address,address,address,address,address[],address[],uint256,uint256,bool,uint256,uint256) params) payable returns(bytes32)
func (_Exchangerouter *ExchangerouterTransactor) CreateWithdrawal(opts *bind.TransactOpts, params WithdrawalUtilsCreateWithdrawalParams) (*types.Transaction, error) {
	return _Exchangerouter.contract.Transact(opts, "createWithdrawal", params)
}

// CreateWithdrawal is a paid mutator transaction binding the contract method 0xad23c5a1.
//
// Solidity: function createWithdrawal((address,address,address,address,address[],address[],uint256,uint256,bool,uint256,uint256) params) payable returns(bytes32)
func (_Exchangerouter *ExchangerouterSession) CreateWithdrawal(params WithdrawalUtilsCreateWithdrawalParams) (*types.Transaction, error) {
	return _Exchangerouter.Contract.CreateWithdrawal(&_Exchangerouter.TransactOpts, params)
}

// CreateWithdrawal is a paid mutator transaction binding the contract method 0xad23c5a1.
//
// Solidity: function createWithdrawal((address,address,address,address,address[],address[],uint256,uint256,bool,uint256,uint256) params) payable returns(bytes32)
func (_Exchangerouter *ExchangerouterTransactorSession) CreateWithdrawal(params WithdrawalUtilsCreateWithdrawalParams) (*types.Transaction, error) {
	return _Exchangerouter.Contract.CreateWithdrawal(&_Exchangerouter.TransactOpts, params)
}

// ExecuteAtomicWithdrawal is a paid mutator transaction binding the contract method 0x6b5341a1.
//
// Solidity: function executeAtomicWithdrawal((address,address,address,address,address[],address[],uint256,uint256,bool,uint256,uint256) params, (address[],address[],bytes[]) oracleParams) payable returns()
func (_Exchangerouter *ExchangerouterTransactor) ExecuteAtomicWithdrawal(opts *bind.TransactOpts, params WithdrawalUtilsCreateWithdrawalParams, oracleParams OracleUtilsSetPricesParams) (*types.Transaction, error) {
	return _Exchangerouter.contract.Transact(opts, "executeAtomicWithdrawal", params, oracleParams)
}

// ExecuteAtomicWithdrawal is a paid mutator transaction binding the contract method 0x6b5341a1.
//
// Solidity: function executeAtomicWithdrawal((address,address,address,address,address[],address[],uint256,uint256,bool,uint256,uint256) params, (address[],address[],bytes[]) oracleParams) payable returns()
func (_Exchangerouter *ExchangerouterSession) ExecuteAtomicWithdrawal(params WithdrawalUtilsCreateWithdrawalParams, oracleParams OracleUtilsSetPricesParams) (*types.Transaction, error) {
	return _Exchangerouter.Contract.ExecuteAtomicWithdrawal(&_Exchangerouter.TransactOpts, params, oracleParams)
}

// ExecuteAtomicWithdrawal is a paid mutator transaction binding the contract method 0x6b5341a1.
//
// Solidity: function executeAtomicWithdrawal((address,address,address,address,address[],address[],uint256,uint256,bool,uint256,uint256) params, (address[],address[],bytes[]) oracleParams) payable returns()
func (_Exchangerouter *ExchangerouterTransactorSession) ExecuteAtomicWithdrawal(params WithdrawalUtilsCreateWithdrawalParams, oracleParams OracleUtilsSetPricesParams) (*types.Transaction, error) {
	return _Exchangerouter.Contract.ExecuteAtomicWithdrawal(&_Exchangerouter.TransactOpts, params, oracleParams)
}

// MakeExternalCalls is a paid mutator transaction binding the contract method 0xd59922b0.
//
// Solidity: function makeExternalCalls(address[] externalCallTargets, bytes[] externalCallDataList, address[] refundTokens, address[] refundReceivers) returns()
func (_Exchangerouter *ExchangerouterTransactor) MakeExternalCalls(opts *bind.TransactOpts, externalCallTargets []common.Address, externalCallDataList [][]byte, refundTokens []common.Address, refundReceivers []common.Address) (*types.Transaction, error) {
	return _Exchangerouter.contract.Transact(opts, "makeExternalCalls", externalCallTargets, externalCallDataList, refundTokens, refundReceivers)
}

// MakeExternalCalls is a paid mutator transaction binding the contract method 0xd59922b0.
//
// Solidity: function makeExternalCalls(address[] externalCallTargets, bytes[] externalCallDataList, address[] refundTokens, address[] refundReceivers) returns()
func (_Exchangerouter *ExchangerouterSession) MakeExternalCalls(externalCallTargets []common.Address, externalCallDataList [][]byte, refundTokens []common.Address, refundReceivers []common.Address) (*types.Transaction, error) {
	return _Exchangerouter.Contract.MakeExternalCalls(&_Exchangerouter.TransactOpts, externalCallTargets, externalCallDataList, refundTokens, refundReceivers)
}

// MakeExternalCalls is a paid mutator transaction binding the contract method 0xd59922b0.
//
// Solidity: function makeExternalCalls(address[] externalCallTargets, bytes[] externalCallDataList, address[] refundTokens, address[] refundReceivers) returns()
func (_Exchangerouter *ExchangerouterTransactorSession) MakeExternalCalls(externalCallTargets []common.Address, externalCallDataList [][]byte, refundTokens []common.Address, refundReceivers []common.Address) (*types.Transaction, error) {
	return _Exchangerouter.Contract.MakeExternalCalls(&_Exchangerouter.TransactOpts, externalCallTargets, externalCallDataList, refundTokens, refundReceivers)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Exchangerouter *ExchangerouterTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Exchangerouter.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Exchangerouter *ExchangerouterSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Exchangerouter.Contract.Multicall(&_Exchangerouter.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Exchangerouter *ExchangerouterTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Exchangerouter.Contract.Multicall(&_Exchangerouter.TransactOpts, data)
}

// SendNativeToken is a paid mutator transaction binding the contract method 0x53ead2d3.
//
// Solidity: function sendNativeToken(address receiver, uint256 amount) payable returns()
func (_Exchangerouter *ExchangerouterTransactor) SendNativeToken(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Exchangerouter.contract.Transact(opts, "sendNativeToken", receiver, amount)
}

// SendNativeToken is a paid mutator transaction binding the contract method 0x53ead2d3.
//
// Solidity: function sendNativeToken(address receiver, uint256 amount) payable returns()
func (_Exchangerouter *ExchangerouterSession) SendNativeToken(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Exchangerouter.Contract.SendNativeToken(&_Exchangerouter.TransactOpts, receiver, amount)
}

// SendNativeToken is a paid mutator transaction binding the contract method 0x53ead2d3.
//
// Solidity: function sendNativeToken(address receiver, uint256 amount) payable returns()
func (_Exchangerouter *ExchangerouterTransactorSession) SendNativeToken(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Exchangerouter.Contract.SendNativeToken(&_Exchangerouter.TransactOpts, receiver, amount)
}

// SendTokens is a paid mutator transaction binding the contract method 0xe6d66ac8.
//
// Solidity: function sendTokens(address token, address receiver, uint256 amount) payable returns()
func (_Exchangerouter *ExchangerouterTransactor) SendTokens(opts *bind.TransactOpts, token common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Exchangerouter.contract.Transact(opts, "sendTokens", token, receiver, amount)
}

// SendTokens is a paid mutator transaction binding the contract method 0xe6d66ac8.
//
// Solidity: function sendTokens(address token, address receiver, uint256 amount) payable returns()
func (_Exchangerouter *ExchangerouterSession) SendTokens(token common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Exchangerouter.Contract.SendTokens(&_Exchangerouter.TransactOpts, token, receiver, amount)
}

// SendTokens is a paid mutator transaction binding the contract method 0xe6d66ac8.
//
// Solidity: function sendTokens(address token, address receiver, uint256 amount) payable returns()
func (_Exchangerouter *ExchangerouterTransactorSession) SendTokens(token common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Exchangerouter.Contract.SendTokens(&_Exchangerouter.TransactOpts, token, receiver, amount)
}

// SendWnt is a paid mutator transaction binding the contract method 0x7d39aaf1.
//
// Solidity: function sendWnt(address receiver, uint256 amount) payable returns()
func (_Exchangerouter *ExchangerouterTransactor) SendWnt(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Exchangerouter.contract.Transact(opts, "sendWnt", receiver, amount)
}

// SendWnt is a paid mutator transaction binding the contract method 0x7d39aaf1.
//
// Solidity: function sendWnt(address receiver, uint256 amount) payable returns()
func (_Exchangerouter *ExchangerouterSession) SendWnt(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Exchangerouter.Contract.SendWnt(&_Exchangerouter.TransactOpts, receiver, amount)
}

// SendWnt is a paid mutator transaction binding the contract method 0x7d39aaf1.
//
// Solidity: function sendWnt(address receiver, uint256 amount) payable returns()
func (_Exchangerouter *ExchangerouterTransactorSession) SendWnt(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Exchangerouter.Contract.SendWnt(&_Exchangerouter.TransactOpts, receiver, amount)
}

// SetSavedCallbackContract is a paid mutator transaction binding the contract method 0x073fb09e.
//
// Solidity: function setSavedCallbackContract(address market, address callbackContract) payable returns()
func (_Exchangerouter *ExchangerouterTransactor) SetSavedCallbackContract(opts *bind.TransactOpts, market common.Address, callbackContract common.Address) (*types.Transaction, error) {
	return _Exchangerouter.contract.Transact(opts, "setSavedCallbackContract", market, callbackContract)
}

// SetSavedCallbackContract is a paid mutator transaction binding the contract method 0x073fb09e.
//
// Solidity: function setSavedCallbackContract(address market, address callbackContract) payable returns()
func (_Exchangerouter *ExchangerouterSession) SetSavedCallbackContract(market common.Address, callbackContract common.Address) (*types.Transaction, error) {
	return _Exchangerouter.Contract.SetSavedCallbackContract(&_Exchangerouter.TransactOpts, market, callbackContract)
}

// SetSavedCallbackContract is a paid mutator transaction binding the contract method 0x073fb09e.
//
// Solidity: function setSavedCallbackContract(address market, address callbackContract) payable returns()
func (_Exchangerouter *ExchangerouterTransactorSession) SetSavedCallbackContract(market common.Address, callbackContract common.Address) (*types.Transaction, error) {
	return _Exchangerouter.Contract.SetSavedCallbackContract(&_Exchangerouter.TransactOpts, market, callbackContract)
}

// SetUiFeeFactor is a paid mutator transaction binding the contract method 0x5a03cd94.
//
// Solidity: function setUiFeeFactor(uint256 uiFeeFactor) payable returns()
func (_Exchangerouter *ExchangerouterTransactor) SetUiFeeFactor(opts *bind.TransactOpts, uiFeeFactor *big.Int) (*types.Transaction, error) {
	return _Exchangerouter.contract.Transact(opts, "setUiFeeFactor", uiFeeFactor)
}

// SetUiFeeFactor is a paid mutator transaction binding the contract method 0x5a03cd94.
//
// Solidity: function setUiFeeFactor(uint256 uiFeeFactor) payable returns()
func (_Exchangerouter *ExchangerouterSession) SetUiFeeFactor(uiFeeFactor *big.Int) (*types.Transaction, error) {
	return _Exchangerouter.Contract.SetUiFeeFactor(&_Exchangerouter.TransactOpts, uiFeeFactor)
}

// SetUiFeeFactor is a paid mutator transaction binding the contract method 0x5a03cd94.
//
// Solidity: function setUiFeeFactor(uint256 uiFeeFactor) payable returns()
func (_Exchangerouter *ExchangerouterTransactorSession) SetUiFeeFactor(uiFeeFactor *big.Int) (*types.Transaction, error) {
	return _Exchangerouter.Contract.SetUiFeeFactor(&_Exchangerouter.TransactOpts, uiFeeFactor)
}

// SimulateExecuteDeposit is a paid mutator transaction binding the contract method 0x2a3db3eb.
//
// Solidity: function simulateExecuteDeposit(bytes32 key, (address[],(uint256,uint256)[],uint256,uint256) simulatedOracleParams) payable returns()
func (_Exchangerouter *ExchangerouterTransactor) SimulateExecuteDeposit(opts *bind.TransactOpts, key [32]byte, simulatedOracleParams OracleUtilsSimulatePricesParams) (*types.Transaction, error) {
	return _Exchangerouter.contract.Transact(opts, "simulateExecuteDeposit", key, simulatedOracleParams)
}

// SimulateExecuteDeposit is a paid mutator transaction binding the contract method 0x2a3db3eb.
//
// Solidity: function simulateExecuteDeposit(bytes32 key, (address[],(uint256,uint256)[],uint256,uint256) simulatedOracleParams) payable returns()
func (_Exchangerouter *ExchangerouterSession) SimulateExecuteDeposit(key [32]byte, simulatedOracleParams OracleUtilsSimulatePricesParams) (*types.Transaction, error) {
	return _Exchangerouter.Contract.SimulateExecuteDeposit(&_Exchangerouter.TransactOpts, key, simulatedOracleParams)
}

// SimulateExecuteDeposit is a paid mutator transaction binding the contract method 0x2a3db3eb.
//
// Solidity: function simulateExecuteDeposit(bytes32 key, (address[],(uint256,uint256)[],uint256,uint256) simulatedOracleParams) payable returns()
func (_Exchangerouter *ExchangerouterTransactorSession) SimulateExecuteDeposit(key [32]byte, simulatedOracleParams OracleUtilsSimulatePricesParams) (*types.Transaction, error) {
	return _Exchangerouter.Contract.SimulateExecuteDeposit(&_Exchangerouter.TransactOpts, key, simulatedOracleParams)
}

// SimulateExecuteOrder is a paid mutator transaction binding the contract method 0xa7115fa8.
//
// Solidity: function simulateExecuteOrder(bytes32 key, (address[],(uint256,uint256)[],uint256,uint256) simulatedOracleParams) payable returns()
func (_Exchangerouter *ExchangerouterTransactor) SimulateExecuteOrder(opts *bind.TransactOpts, key [32]byte, simulatedOracleParams OracleUtilsSimulatePricesParams) (*types.Transaction, error) {
	return _Exchangerouter.contract.Transact(opts, "simulateExecuteOrder", key, simulatedOracleParams)
}

// SimulateExecuteOrder is a paid mutator transaction binding the contract method 0xa7115fa8.
//
// Solidity: function simulateExecuteOrder(bytes32 key, (address[],(uint256,uint256)[],uint256,uint256) simulatedOracleParams) payable returns()
func (_Exchangerouter *ExchangerouterSession) SimulateExecuteOrder(key [32]byte, simulatedOracleParams OracleUtilsSimulatePricesParams) (*types.Transaction, error) {
	return _Exchangerouter.Contract.SimulateExecuteOrder(&_Exchangerouter.TransactOpts, key, simulatedOracleParams)
}

// SimulateExecuteOrder is a paid mutator transaction binding the contract method 0xa7115fa8.
//
// Solidity: function simulateExecuteOrder(bytes32 key, (address[],(uint256,uint256)[],uint256,uint256) simulatedOracleParams) payable returns()
func (_Exchangerouter *ExchangerouterTransactorSession) SimulateExecuteOrder(key [32]byte, simulatedOracleParams OracleUtilsSimulatePricesParams) (*types.Transaction, error) {
	return _Exchangerouter.Contract.SimulateExecuteOrder(&_Exchangerouter.TransactOpts, key, simulatedOracleParams)
}

// SimulateExecuteShift is a paid mutator transaction binding the contract method 0x49e58799.
//
// Solidity: function simulateExecuteShift(bytes32 key, (address[],(uint256,uint256)[],uint256,uint256) simulatedOracleParams) payable returns()
func (_Exchangerouter *ExchangerouterTransactor) SimulateExecuteShift(opts *bind.TransactOpts, key [32]byte, simulatedOracleParams OracleUtilsSimulatePricesParams) (*types.Transaction, error) {
	return _Exchangerouter.contract.Transact(opts, "simulateExecuteShift", key, simulatedOracleParams)
}

// SimulateExecuteShift is a paid mutator transaction binding the contract method 0x49e58799.
//
// Solidity: function simulateExecuteShift(bytes32 key, (address[],(uint256,uint256)[],uint256,uint256) simulatedOracleParams) payable returns()
func (_Exchangerouter *ExchangerouterSession) SimulateExecuteShift(key [32]byte, simulatedOracleParams OracleUtilsSimulatePricesParams) (*types.Transaction, error) {
	return _Exchangerouter.Contract.SimulateExecuteShift(&_Exchangerouter.TransactOpts, key, simulatedOracleParams)
}

// SimulateExecuteShift is a paid mutator transaction binding the contract method 0x49e58799.
//
// Solidity: function simulateExecuteShift(bytes32 key, (address[],(uint256,uint256)[],uint256,uint256) simulatedOracleParams) payable returns()
func (_Exchangerouter *ExchangerouterTransactorSession) SimulateExecuteShift(key [32]byte, simulatedOracleParams OracleUtilsSimulatePricesParams) (*types.Transaction, error) {
	return _Exchangerouter.Contract.SimulateExecuteShift(&_Exchangerouter.TransactOpts, key, simulatedOracleParams)
}

// SimulateExecuteWithdrawal is a paid mutator transaction binding the contract method 0xbc190c14.
//
// Solidity: function simulateExecuteWithdrawal(bytes32 key, (address[],(uint256,uint256)[],uint256,uint256) simulatedOracleParams, uint8 swapPricingType) payable returns()
func (_Exchangerouter *ExchangerouterTransactor) SimulateExecuteWithdrawal(opts *bind.TransactOpts, key [32]byte, simulatedOracleParams OracleUtilsSimulatePricesParams, swapPricingType uint8) (*types.Transaction, error) {
	return _Exchangerouter.contract.Transact(opts, "simulateExecuteWithdrawal", key, simulatedOracleParams, swapPricingType)
}

// SimulateExecuteWithdrawal is a paid mutator transaction binding the contract method 0xbc190c14.
//
// Solidity: function simulateExecuteWithdrawal(bytes32 key, (address[],(uint256,uint256)[],uint256,uint256) simulatedOracleParams, uint8 swapPricingType) payable returns()
func (_Exchangerouter *ExchangerouterSession) SimulateExecuteWithdrawal(key [32]byte, simulatedOracleParams OracleUtilsSimulatePricesParams, swapPricingType uint8) (*types.Transaction, error) {
	return _Exchangerouter.Contract.SimulateExecuteWithdrawal(&_Exchangerouter.TransactOpts, key, simulatedOracleParams, swapPricingType)
}

// SimulateExecuteWithdrawal is a paid mutator transaction binding the contract method 0xbc190c14.
//
// Solidity: function simulateExecuteWithdrawal(bytes32 key, (address[],(uint256,uint256)[],uint256,uint256) simulatedOracleParams, uint8 swapPricingType) payable returns()
func (_Exchangerouter *ExchangerouterTransactorSession) SimulateExecuteWithdrawal(key [32]byte, simulatedOracleParams OracleUtilsSimulatePricesParams, swapPricingType uint8) (*types.Transaction, error) {
	return _Exchangerouter.Contract.SimulateExecuteWithdrawal(&_Exchangerouter.TransactOpts, key, simulatedOracleParams, swapPricingType)
}

// UpdateOrder is a paid mutator transaction binding the contract method 0xf82a2272.
//
// Solidity: function updateOrder(bytes32 key, uint256 sizeDeltaUsd, uint256 acceptablePrice, uint256 triggerPrice, uint256 minOutputAmount, bool autoCancel) payable returns()
func (_Exchangerouter *ExchangerouterTransactor) UpdateOrder(opts *bind.TransactOpts, key [32]byte, sizeDeltaUsd *big.Int, acceptablePrice *big.Int, triggerPrice *big.Int, minOutputAmount *big.Int, autoCancel bool) (*types.Transaction, error) {
	return _Exchangerouter.contract.Transact(opts, "updateOrder", key, sizeDeltaUsd, acceptablePrice, triggerPrice, minOutputAmount, autoCancel)
}

// UpdateOrder is a paid mutator transaction binding the contract method 0xf82a2272.
//
// Solidity: function updateOrder(bytes32 key, uint256 sizeDeltaUsd, uint256 acceptablePrice, uint256 triggerPrice, uint256 minOutputAmount, bool autoCancel) payable returns()
func (_Exchangerouter *ExchangerouterSession) UpdateOrder(key [32]byte, sizeDeltaUsd *big.Int, acceptablePrice *big.Int, triggerPrice *big.Int, minOutputAmount *big.Int, autoCancel bool) (*types.Transaction, error) {
	return _Exchangerouter.Contract.UpdateOrder(&_Exchangerouter.TransactOpts, key, sizeDeltaUsd, acceptablePrice, triggerPrice, minOutputAmount, autoCancel)
}

// UpdateOrder is a paid mutator transaction binding the contract method 0xf82a2272.
//
// Solidity: function updateOrder(bytes32 key, uint256 sizeDeltaUsd, uint256 acceptablePrice, uint256 triggerPrice, uint256 minOutputAmount, bool autoCancel) payable returns()
func (_Exchangerouter *ExchangerouterTransactorSession) UpdateOrder(key [32]byte, sizeDeltaUsd *big.Int, acceptablePrice *big.Int, triggerPrice *big.Int, minOutputAmount *big.Int, autoCancel bool) (*types.Transaction, error) {
	return _Exchangerouter.Contract.UpdateOrder(&_Exchangerouter.TransactOpts, key, sizeDeltaUsd, acceptablePrice, triggerPrice, minOutputAmount, autoCancel)
}
