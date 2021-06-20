// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// TokenSaleABI is the input ABI used to generate the binding from.
const TokenSaleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"duration\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxyAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minPayment\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"endTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"walletAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bonusUsed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"availableUnits\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isActiveSale\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unitsSold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownerAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiReceived\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"availableBonus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_walletAddress\",\"type\":\"address\"},{\"name\":\"_proxyAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"PriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SaleDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_startTime\",\"type\":\"uint256\"}],\"name\":\"SaleStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_endTime\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"SaleTerminated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"unitCount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"isBitcoin\",\"type\":\"bool\"}],\"name\":\"TokensPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RefundProcessed\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"startSale\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"terminateSale\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"input\",\"type\":\"address[]\"}],\"name\":\"whitelistAdd\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"input\",\"type\":\"address[]\"}],\"name\":\"whitelistRemove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"updatePrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"transfer\",\"type\":\"uint256\"}],\"name\":\"buyTokensBTC\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"input\",\"type\":\"address\"}],\"name\":\"whitelistCheck\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sold\",\"type\":\"uint256\"}],\"name\":\"getBonusTier\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sought\",\"type\":\"uint256\"},{\"name\":\"sold\",\"type\":\"uint256\"}],\"name\":\"getBonusTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TokenSale is an auto generated Go binding around an Ethereum contract.
type TokenSale struct {
	TokenSaleCaller     // Read-only binding to the contract
	TokenSaleTransactor // Write-only binding to the contract
	TokenSaleFilterer   // Log filterer for contract events
}

// TokenSaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenSaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenSaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenSaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSaleSession struct {
	Contract     *TokenSale        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenSaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenSaleCallerSession struct {
	Contract *TokenSaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TokenSaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenSaleTransactorSession struct {
	Contract     *TokenSaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TokenSaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenSaleRaw struct {
	Contract *TokenSale // Generic contract binding to access the raw methods on
}

// TokenSaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenSaleCallerRaw struct {
	Contract *TokenSaleCaller // Generic read-only contract binding to access the raw methods on
}

// TokenSaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenSaleTransactorRaw struct {
	Contract *TokenSaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenSale creates a new instance of TokenSale, bound to a specific deployed contract.
func NewTokenSale(address common.Address, backend bind.ContractBackend) (*TokenSale, error) {
	contract, err := bindTokenSale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenSale{TokenSaleCaller: TokenSaleCaller{contract: contract}, TokenSaleTransactor: TokenSaleTransactor{contract: contract}, TokenSaleFilterer: TokenSaleFilterer{contract: contract}}, nil
}

// NewTokenSaleCaller creates a new read-only instance of TokenSale, bound to a specific deployed contract.
func NewTokenSaleCaller(address common.Address, caller bind.ContractCaller) (*TokenSaleCaller, error) {
	contract, err := bindTokenSale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenSaleCaller{contract: contract}, nil
}

// NewTokenSaleTransactor creates a new write-only instance of TokenSale, bound to a specific deployed contract.
func NewTokenSaleTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenSaleTransactor, error) {
	contract, err := bindTokenSale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenSaleTransactor{contract: contract}, nil
}

// NewTokenSaleFilterer creates a new log filterer instance of TokenSale, bound to a specific deployed contract.
func NewTokenSaleFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenSaleFilterer, error) {
	contract, err := bindTokenSale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenSaleFilterer{contract: contract}, nil
}

// bindTokenSale binds a generic wrapper to an already deployed contract.
func bindTokenSale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenSaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenSale *TokenSaleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenSale.Contract.TokenSaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenSale *TokenSaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenSale.Contract.TokenSaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenSale *TokenSaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenSale.Contract.TokenSaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenSale *TokenSaleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenSale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenSale *TokenSaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenSale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenSale *TokenSaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenSale.Contract.contract.Transact(opts, method, params...)
}

// AvailableBonus is a free data retrieval call binding the contract method 0xd1b1910a.
//
// Solidity: function availableBonus() constant returns(uint256)
func (_TokenSale *TokenSaleCaller) AvailableBonus(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenSale.contract.Call(opts, out, "availableBonus")
	return *ret0, err
}

// AvailableBonus is a free data retrieval call binding the contract method 0xd1b1910a.
//
// Solidity: function availableBonus() constant returns(uint256)
func (_TokenSale *TokenSaleSession) AvailableBonus() (*big.Int, error) {
	return _TokenSale.Contract.AvailableBonus(&_TokenSale.CallOpts)
}

// AvailableBonus is a free data retrieval call binding the contract method 0xd1b1910a.
//
// Solidity: function availableBonus() constant returns(uint256)
func (_TokenSale *TokenSaleCallerSession) AvailableBonus() (*big.Int, error) {
	return _TokenSale.Contract.AvailableBonus(&_TokenSale.CallOpts)
}

// AvailableUnits is a free data retrieval call binding the contract method 0x8324bb51.
//
// Solidity: function availableUnits() constant returns(uint256)
func (_TokenSale *TokenSaleCaller) AvailableUnits(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenSale.contract.Call(opts, out, "availableUnits")
	return *ret0, err
}

// AvailableUnits is a free data retrieval call binding the contract method 0x8324bb51.
//
// Solidity: function availableUnits() constant returns(uint256)
func (_TokenSale *TokenSaleSession) AvailableUnits() (*big.Int, error) {
	return _TokenSale.Contract.AvailableUnits(&_TokenSale.CallOpts)
}

// AvailableUnits is a free data retrieval call binding the contract method 0x8324bb51.
//
// Solidity: function availableUnits() constant returns(uint256)
func (_TokenSale *TokenSaleCallerSession) AvailableUnits() (*big.Int, error) {
	return _TokenSale.Contract.AvailableUnits(&_TokenSale.CallOpts)
}

// BonusUsed is a free data retrieval call binding the contract method 0x7974fb87.
//
// Solidity: function bonusUsed() constant returns(uint256)
func (_TokenSale *TokenSaleCaller) BonusUsed(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenSale.contract.Call(opts, out, "bonusUsed")
	return *ret0, err
}

// BonusUsed is a free data retrieval call binding the contract method 0x7974fb87.
//
// Solidity: function bonusUsed() constant returns(uint256)
func (_TokenSale *TokenSaleSession) BonusUsed() (*big.Int, error) {
	return _TokenSale.Contract.BonusUsed(&_TokenSale.CallOpts)
}

// BonusUsed is a free data retrieval call binding the contract method 0x7974fb87.
//
// Solidity: function bonusUsed() constant returns(uint256)
func (_TokenSale *TokenSaleCallerSession) BonusUsed() (*big.Int, error) {
	return _TokenSale.Contract.BonusUsed(&_TokenSale.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() constant returns(uint256)
func (_TokenSale *TokenSaleCaller) Duration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenSale.contract.Call(opts, out, "duration")
	return *ret0, err
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() constant returns(uint256)
func (_TokenSale *TokenSaleSession) Duration() (*big.Int, error) {
	return _TokenSale.Contract.Duration(&_TokenSale.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() constant returns(uint256)
func (_TokenSale *TokenSaleCallerSession) Duration() (*big.Int, error) {
	return _TokenSale.Contract.Duration(&_TokenSale.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() constant returns(uint256)
func (_TokenSale *TokenSaleCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenSale.contract.Call(opts, out, "endTime")
	return *ret0, err
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() constant returns(uint256)
func (_TokenSale *TokenSaleSession) EndTime() (*big.Int, error) {
	return _TokenSale.Contract.EndTime(&_TokenSale.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() constant returns(uint256)
func (_TokenSale *TokenSaleCallerSession) EndTime() (*big.Int, error) {
	return _TokenSale.Contract.EndTime(&_TokenSale.CallOpts)
}

// GetBonusTier is a free data retrieval call binding the contract method 0x935e85b2.
//
// Solidity: function getBonusTier(sold uint256) constant returns(uint8)
func (_TokenSale *TokenSaleCaller) GetBonusTier(opts *bind.CallOpts, sold *big.Int) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _TokenSale.contract.Call(opts, out, "getBonusTier", sold)
	return *ret0, err
}

// GetBonusTier is a free data retrieval call binding the contract method 0x935e85b2.
//
// Solidity: function getBonusTier(sold uint256) constant returns(uint8)
func (_TokenSale *TokenSaleSession) GetBonusTier(sold *big.Int) (uint8, error) {
	return _TokenSale.Contract.GetBonusTier(&_TokenSale.CallOpts, sold)
}

// GetBonusTier is a free data retrieval call binding the contract method 0x935e85b2.
//
// Solidity: function getBonusTier(sold uint256) constant returns(uint8)
func (_TokenSale *TokenSaleCallerSession) GetBonusTier(sold *big.Int) (uint8, error) {
	return _TokenSale.Contract.GetBonusTier(&_TokenSale.CallOpts, sold)
}

// GetBonusTokens is a free data retrieval call binding the contract method 0x1153e79c.
//
// Solidity: function getBonusTokens(sought uint256, sold uint256) constant returns(uint256)
func (_TokenSale *TokenSaleCaller) GetBonusTokens(opts *bind.CallOpts, sought *big.Int, sold *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenSale.contract.Call(opts, out, "getBonusTokens", sought, sold)
	return *ret0, err
}

// GetBonusTokens is a free data retrieval call binding the contract method 0x1153e79c.
//
// Solidity: function getBonusTokens(sought uint256, sold uint256) constant returns(uint256)
func (_TokenSale *TokenSaleSession) GetBonusTokens(sought *big.Int, sold *big.Int) (*big.Int, error) {
	return _TokenSale.Contract.GetBonusTokens(&_TokenSale.CallOpts, sought, sold)
}

// GetBonusTokens is a free data retrieval call binding the contract method 0x1153e79c.
//
// Solidity: function getBonusTokens(sought uint256, sold uint256) constant returns(uint256)
func (_TokenSale *TokenSaleCallerSession) GetBonusTokens(sought *big.Int, sold *big.Int) (*big.Int, error) {
	return _TokenSale.Contract.GetBonusTokens(&_TokenSale.CallOpts, sought, sold)
}

// IsActiveSale is a free data retrieval call binding the contract method 0x89a44a44.
//
// Solidity: function isActiveSale() constant returns(bool)
func (_TokenSale *TokenSaleCaller) IsActiveSale(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenSale.contract.Call(opts, out, "isActiveSale")
	return *ret0, err
}

// IsActiveSale is a free data retrieval call binding the contract method 0x89a44a44.
//
// Solidity: function isActiveSale() constant returns(bool)
func (_TokenSale *TokenSaleSession) IsActiveSale() (bool, error) {
	return _TokenSale.Contract.IsActiveSale(&_TokenSale.CallOpts)
}

// IsActiveSale is a free data retrieval call binding the contract method 0x89a44a44.
//
// Solidity: function isActiveSale() constant returns(bool)
func (_TokenSale *TokenSaleCallerSession) IsActiveSale() (bool, error) {
	return _TokenSale.Contract.IsActiveSale(&_TokenSale.CallOpts)
}

// MinPayment is a free data retrieval call binding the contract method 0x2e276499.
//
// Solidity: function minPayment() constant returns(uint256)
func (_TokenSale *TokenSaleCaller) MinPayment(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenSale.contract.Call(opts, out, "minPayment")
	return *ret0, err
}

// MinPayment is a free data retrieval call binding the contract method 0x2e276499.
//
// Solidity: function minPayment() constant returns(uint256)
func (_TokenSale *TokenSaleSession) MinPayment() (*big.Int, error) {
	return _TokenSale.Contract.MinPayment(&_TokenSale.CallOpts)
}

// MinPayment is a free data retrieval call binding the contract method 0x2e276499.
//
// Solidity: function minPayment() constant returns(uint256)
func (_TokenSale *TokenSaleCallerSession) MinPayment() (*big.Int, error) {
	return _TokenSale.Contract.MinPayment(&_TokenSale.CallOpts)
}

// OwnerAddress is a free data retrieval call binding the contract method 0x8f84aa09.
//
// Solidity: function ownerAddress() constant returns(address)
func (_TokenSale *TokenSaleCaller) OwnerAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenSale.contract.Call(opts, out, "ownerAddress")
	return *ret0, err
}

// OwnerAddress is a free data retrieval call binding the contract method 0x8f84aa09.
//
// Solidity: function ownerAddress() constant returns(address)
func (_TokenSale *TokenSaleSession) OwnerAddress() (common.Address, error) {
	return _TokenSale.Contract.OwnerAddress(&_TokenSale.CallOpts)
}

// OwnerAddress is a free data retrieval call binding the contract method 0x8f84aa09.
//
// Solidity: function ownerAddress() constant returns(address)
func (_TokenSale *TokenSaleCallerSession) OwnerAddress() (common.Address, error) {
	return _TokenSale.Contract.OwnerAddress(&_TokenSale.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_TokenSale *TokenSaleCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenSale.contract.Call(opts, out, "price")
	return *ret0, err
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_TokenSale *TokenSaleSession) Price() (*big.Int, error) {
	return _TokenSale.Contract.Price(&_TokenSale.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_TokenSale *TokenSaleCallerSession) Price() (*big.Int, error) {
	return _TokenSale.Contract.Price(&_TokenSale.CallOpts)
}

// ProxyAddress is a free data retrieval call binding the contract method 0x23f5c02d.
//
// Solidity: function proxyAddress() constant returns(address)
func (_TokenSale *TokenSaleCaller) ProxyAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenSale.contract.Call(opts, out, "proxyAddress")
	return *ret0, err
}

// ProxyAddress is a free data retrieval call binding the contract method 0x23f5c02d.
//
// Solidity: function proxyAddress() constant returns(address)
func (_TokenSale *TokenSaleSession) ProxyAddress() (common.Address, error) {
	return _TokenSale.Contract.ProxyAddress(&_TokenSale.CallOpts)
}

// ProxyAddress is a free data retrieval call binding the contract method 0x23f5c02d.
//
// Solidity: function proxyAddress() constant returns(address)
func (_TokenSale *TokenSaleCallerSession) ProxyAddress() (common.Address, error) {
	return _TokenSale.Contract.ProxyAddress(&_TokenSale.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_TokenSale *TokenSaleCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenSale.contract.Call(opts, out, "startTime")
	return *ret0, err
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_TokenSale *TokenSaleSession) StartTime() (*big.Int, error) {
	return _TokenSale.Contract.StartTime(&_TokenSale.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_TokenSale *TokenSaleCallerSession) StartTime() (*big.Int, error) {
	return _TokenSale.Contract.StartTime(&_TokenSale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_TokenSale *TokenSaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenSale.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_TokenSale *TokenSaleSession) Token() (common.Address, error) {
	return _TokenSale.Contract.Token(&_TokenSale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_TokenSale *TokenSaleCallerSession) Token() (common.Address, error) {
	return _TokenSale.Contract.Token(&_TokenSale.CallOpts)
}

// UnitsSold is a free data retrieval call binding the contract method 0x8e26532b.
//
// Solidity: function unitsSold() constant returns(uint256)
func (_TokenSale *TokenSaleCaller) UnitsSold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenSale.contract.Call(opts, out, "unitsSold")
	return *ret0, err
}

// UnitsSold is a free data retrieval call binding the contract method 0x8e26532b.
//
// Solidity: function unitsSold() constant returns(uint256)
func (_TokenSale *TokenSaleSession) UnitsSold() (*big.Int, error) {
	return _TokenSale.Contract.UnitsSold(&_TokenSale.CallOpts)
}

// UnitsSold is a free data retrieval call binding the contract method 0x8e26532b.
//
// Solidity: function unitsSold() constant returns(uint256)
func (_TokenSale *TokenSaleCallerSession) UnitsSold() (*big.Int, error) {
	return _TokenSale.Contract.UnitsSold(&_TokenSale.CallOpts)
}

// WalletAddress is a free data retrieval call binding the contract method 0x6ad5b3ea.
//
// Solidity: function walletAddress() constant returns(address)
func (_TokenSale *TokenSaleCaller) WalletAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenSale.contract.Call(opts, out, "walletAddress")
	return *ret0, err
}

// WalletAddress is a free data retrieval call binding the contract method 0x6ad5b3ea.
//
// Solidity: function walletAddress() constant returns(address)
func (_TokenSale *TokenSaleSession) WalletAddress() (common.Address, error) {
	return _TokenSale.Contract.WalletAddress(&_TokenSale.CallOpts)
}

// WalletAddress is a free data retrieval call binding the contract method 0x6ad5b3ea.
//
// Solidity: function walletAddress() constant returns(address)
func (_TokenSale *TokenSaleCallerSession) WalletAddress() (common.Address, error) {
	return _TokenSale.Contract.WalletAddress(&_TokenSale.CallOpts)
}

// WeiReceived is a free data retrieval call binding the contract method 0xa000aeb7.
//
// Solidity: function weiReceived() constant returns(uint256)
func (_TokenSale *TokenSaleCaller) WeiReceived(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenSale.contract.Call(opts, out, "weiReceived")
	return *ret0, err
}

// WeiReceived is a free data retrieval call binding the contract method 0xa000aeb7.
//
// Solidity: function weiReceived() constant returns(uint256)
func (_TokenSale *TokenSaleSession) WeiReceived() (*big.Int, error) {
	return _TokenSale.Contract.WeiReceived(&_TokenSale.CallOpts)
}

// WeiReceived is a free data retrieval call binding the contract method 0xa000aeb7.
//
// Solidity: function weiReceived() constant returns(uint256)
func (_TokenSale *TokenSaleCallerSession) WeiReceived() (*big.Int, error) {
	return _TokenSale.Contract.WeiReceived(&_TokenSale.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_TokenSale *TokenSaleCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenSale.contract.Call(opts, out, "whitelist", arg0)
	return *ret0, err
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_TokenSale *TokenSaleSession) Whitelist(arg0 common.Address) (bool, error) {
	return _TokenSale.Contract.Whitelist(&_TokenSale.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_TokenSale *TokenSaleCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _TokenSale.Contract.Whitelist(&_TokenSale.CallOpts, arg0)
}

// WhitelistCheck is a free data retrieval call binding the contract method 0x882eb672.
//
// Solidity: function whitelistCheck(input address) constant returns(bool)
func (_TokenSale *TokenSaleCaller) WhitelistCheck(opts *bind.CallOpts, input common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenSale.contract.Call(opts, out, "whitelistCheck", input)
	return *ret0, err
}

// WhitelistCheck is a free data retrieval call binding the contract method 0x882eb672.
//
// Solidity: function whitelistCheck(input address) constant returns(bool)
func (_TokenSale *TokenSaleSession) WhitelistCheck(input common.Address) (bool, error) {
	return _TokenSale.Contract.WhitelistCheck(&_TokenSale.CallOpts, input)
}

// WhitelistCheck is a free data retrieval call binding the contract method 0x882eb672.
//
// Solidity: function whitelistCheck(input address) constant returns(bool)
func (_TokenSale *TokenSaleCallerSession) WhitelistCheck(input common.Address) (bool, error) {
	return _TokenSale.Contract.WhitelistCheck(&_TokenSale.CallOpts, input)
}

// BuyTokensBTC is a paid mutator transaction binding the contract method 0x844197a4.
//
// Solidity: function buyTokensBTC(sender address, transfer uint256) returns()
func (_TokenSale *TokenSaleTransactor) BuyTokensBTC(opts *bind.TransactOpts, sender common.Address, transfer *big.Int) (*types.Transaction, error) {
	return _TokenSale.contract.Transact(opts, "buyTokensBTC", sender, transfer)
}

// BuyTokensBTC is a paid mutator transaction binding the contract method 0x844197a4.
//
// Solidity: function buyTokensBTC(sender address, transfer uint256) returns()
func (_TokenSale *TokenSaleSession) BuyTokensBTC(sender common.Address, transfer *big.Int) (*types.Transaction, error) {
	return _TokenSale.Contract.BuyTokensBTC(&_TokenSale.TransactOpts, sender, transfer)
}

// BuyTokensBTC is a paid mutator transaction binding the contract method 0x844197a4.
//
// Solidity: function buyTokensBTC(sender address, transfer uint256) returns()
func (_TokenSale *TokenSaleTransactorSession) BuyTokensBTC(sender common.Address, transfer *big.Int) (*types.Transaction, error) {
	return _TokenSale.Contract.BuyTokensBTC(&_TokenSale.TransactOpts, sender, transfer)
}

// StartSale is a paid mutator transaction binding the contract method 0x3f2916d9.
//
// Solidity: function startSale(_tokenAddress address) returns()
func (_TokenSale *TokenSaleTransactor) StartSale(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _TokenSale.contract.Transact(opts, "startSale", _tokenAddress)
}

// StartSale is a paid mutator transaction binding the contract method 0x3f2916d9.
//
// Solidity: function startSale(_tokenAddress address) returns()
func (_TokenSale *TokenSaleSession) StartSale(_tokenAddress common.Address) (*types.Transaction, error) {
	return _TokenSale.Contract.StartSale(&_TokenSale.TransactOpts, _tokenAddress)
}

// StartSale is a paid mutator transaction binding the contract method 0x3f2916d9.
//
// Solidity: function startSale(_tokenAddress address) returns()
func (_TokenSale *TokenSaleTransactorSession) StartSale(_tokenAddress common.Address) (*types.Transaction, error) {
	return _TokenSale.Contract.StartSale(&_TokenSale.TransactOpts, _tokenAddress)
}

// TerminateSale is a paid mutator transaction binding the contract method 0x80787f2c.
//
// Solidity: function terminateSale() returns()
func (_TokenSale *TokenSaleTransactor) TerminateSale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenSale.contract.Transact(opts, "terminateSale")
}

// TerminateSale is a paid mutator transaction binding the contract method 0x80787f2c.
//
// Solidity: function terminateSale() returns()
func (_TokenSale *TokenSaleSession) TerminateSale() (*types.Transaction, error) {
	return _TokenSale.Contract.TerminateSale(&_TokenSale.TransactOpts)
}

// TerminateSale is a paid mutator transaction binding the contract method 0x80787f2c.
//
// Solidity: function terminateSale() returns()
func (_TokenSale *TokenSaleTransactorSession) TerminateSale() (*types.Transaction, error) {
	return _TokenSale.Contract.TerminateSale(&_TokenSale.TransactOpts)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8d6cc56d.
//
// Solidity: function updatePrice(value uint256) returns()
func (_TokenSale *TokenSaleTransactor) UpdatePrice(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _TokenSale.contract.Transact(opts, "updatePrice", value)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8d6cc56d.
//
// Solidity: function updatePrice(value uint256) returns()
func (_TokenSale *TokenSaleSession) UpdatePrice(value *big.Int) (*types.Transaction, error) {
	return _TokenSale.Contract.UpdatePrice(&_TokenSale.TransactOpts, value)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8d6cc56d.
//
// Solidity: function updatePrice(value uint256) returns()
func (_TokenSale *TokenSaleTransactorSession) UpdatePrice(value *big.Int) (*types.Transaction, error) {
	return _TokenSale.Contract.UpdatePrice(&_TokenSale.TransactOpts, value)
}

// WhitelistAdd is a paid mutator transaction binding the contract method 0x9cf94943.
//
// Solidity: function whitelistAdd(input address[]) returns()
func (_TokenSale *TokenSaleTransactor) WhitelistAdd(opts *bind.TransactOpts, input []common.Address) (*types.Transaction, error) {
	return _TokenSale.contract.Transact(opts, "whitelistAdd", input)
}

// WhitelistAdd is a paid mutator transaction binding the contract method 0x9cf94943.
//
// Solidity: function whitelistAdd(input address[]) returns()
func (_TokenSale *TokenSaleSession) WhitelistAdd(input []common.Address) (*types.Transaction, error) {
	return _TokenSale.Contract.WhitelistAdd(&_TokenSale.TransactOpts, input)
}

// WhitelistAdd is a paid mutator transaction binding the contract method 0x9cf94943.
//
// Solidity: function whitelistAdd(input address[]) returns()
func (_TokenSale *TokenSaleTransactorSession) WhitelistAdd(input []common.Address) (*types.Transaction, error) {
	return _TokenSale.Contract.WhitelistAdd(&_TokenSale.TransactOpts, input)
}

// WhitelistRemove is a paid mutator transaction binding the contract method 0x0c82b942.
//
// Solidity: function whitelistRemove(input address[]) returns()
func (_TokenSale *TokenSaleTransactor) WhitelistRemove(opts *bind.TransactOpts, input []common.Address) (*types.Transaction, error) {
	return _TokenSale.contract.Transact(opts, "whitelistRemove", input)
}

// WhitelistRemove is a paid mutator transaction binding the contract method 0x0c82b942.
//
// Solidity: function whitelistRemove(input address[]) returns()
func (_TokenSale *TokenSaleSession) WhitelistRemove(input []common.Address) (*types.Transaction, error) {
	return _TokenSale.Contract.WhitelistRemove(&_TokenSale.TransactOpts, input)
}

// WhitelistRemove is a paid mutator transaction binding the contract method 0x0c82b942.
//
// Solidity: function whitelistRemove(input address[]) returns()
func (_TokenSale *TokenSaleTransactorSession) WhitelistRemove(input []common.Address) (*types.Transaction, error) {
	return _TokenSale.Contract.WhitelistRemove(&_TokenSale.TransactOpts, input)
}

// TokenSalePriceUpdatedIterator is returned from FilterPriceUpdated and is used to iterate over the raw logs and unpacked data for PriceUpdated events raised by the TokenSale contract.
type TokenSalePriceUpdatedIterator struct {
	Event *TokenSalePriceUpdated // Event containing the contract specifics and raw log

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
func (it *TokenSalePriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSalePriceUpdated)
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
		it.Event = new(TokenSalePriceUpdated)
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
func (it *TokenSalePriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSalePriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSalePriceUpdated represents a PriceUpdated event raised by the TokenSale contract.
type TokenSalePriceUpdated struct {
	Time  *big.Int
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPriceUpdated is a free log retrieval operation binding the contract event 0x945c1c4e99aa89f648fbfe3df471b916f719e16d960fcec0737d4d56bd696838.
//
// Solidity: event PriceUpdated(time uint256, value uint256)
func (_TokenSale *TokenSaleFilterer) FilterPriceUpdated(opts *bind.FilterOpts) (*TokenSalePriceUpdatedIterator, error) {

	logs, sub, err := _TokenSale.contract.FilterLogs(opts, "PriceUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenSalePriceUpdatedIterator{contract: _TokenSale.contract, event: "PriceUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceUpdated is a free log subscription operation binding the contract event 0x945c1c4e99aa89f648fbfe3df471b916f719e16d960fcec0737d4d56bd696838.
//
// Solidity: event PriceUpdated(time uint256, value uint256)
func (_TokenSale *TokenSaleFilterer) WatchPriceUpdated(opts *bind.WatchOpts, sink chan<- *TokenSalePriceUpdated) (event.Subscription, error) {

	logs, sub, err := _TokenSale.contract.WatchLogs(opts, "PriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSalePriceUpdated)
				if err := _TokenSale.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
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

// TokenSaleRefundProcessedIterator is returned from FilterRefundProcessed and is used to iterate over the raw logs and unpacked data for RefundProcessed events raised by the TokenSale contract.
type TokenSaleRefundProcessedIterator struct {
	Event *TokenSaleRefundProcessed // Event containing the contract specifics and raw log

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
func (it *TokenSaleRefundProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSaleRefundProcessed)
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
		it.Event = new(TokenSaleRefundProcessed)
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
func (it *TokenSaleRefundProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSaleRefundProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSaleRefundProcessed represents a RefundProcessed event raised by the TokenSale contract.
type TokenSaleRefundProcessed struct {
	Beneficiary common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRefundProcessed is a free log retrieval operation binding the contract event 0x3367befd2b2f39615cd79917c2153263c4af1d3945ec003e5d5bfc13a8d85833.
//
// Solidity: event RefundProcessed(_beneficiary indexed address, amount uint256)
func (_TokenSale *TokenSaleFilterer) FilterRefundProcessed(opts *bind.FilterOpts, _beneficiary []common.Address) (*TokenSaleRefundProcessedIterator, error) {

	var _beneficiaryRule []interface{}
	for _, _beneficiaryItem := range _beneficiary {
		_beneficiaryRule = append(_beneficiaryRule, _beneficiaryItem)
	}

	logs, sub, err := _TokenSale.contract.FilterLogs(opts, "RefundProcessed", _beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &TokenSaleRefundProcessedIterator{contract: _TokenSale.contract, event: "RefundProcessed", logs: logs, sub: sub}, nil
}

// WatchRefundProcessed is a free log subscription operation binding the contract event 0x3367befd2b2f39615cd79917c2153263c4af1d3945ec003e5d5bfc13a8d85833.
//
// Solidity: event RefundProcessed(_beneficiary indexed address, amount uint256)
func (_TokenSale *TokenSaleFilterer) WatchRefundProcessed(opts *bind.WatchOpts, sink chan<- *TokenSaleRefundProcessed, _beneficiary []common.Address) (event.Subscription, error) {

	var _beneficiaryRule []interface{}
	for _, _beneficiaryItem := range _beneficiary {
		_beneficiaryRule = append(_beneficiaryRule, _beneficiaryItem)
	}

	logs, sub, err := _TokenSale.contract.WatchLogs(opts, "RefundProcessed", _beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSaleRefundProcessed)
				if err := _TokenSale.contract.UnpackLog(event, "RefundProcessed", log); err != nil {
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

// TokenSaleSaleDeployedIterator is returned from FilterSaleDeployed and is used to iterate over the raw logs and unpacked data for SaleDeployed events raised by the TokenSale contract.
type TokenSaleSaleDeployedIterator struct {
	Event *TokenSaleSaleDeployed // Event containing the contract specifics and raw log

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
func (it *TokenSaleSaleDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSaleSaleDeployed)
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
		it.Event = new(TokenSaleSaleDeployed)
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
func (it *TokenSaleSaleDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSaleSaleDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSaleSaleDeployed represents a SaleDeployed event raised by the TokenSale contract.
type TokenSaleSaleDeployed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSaleDeployed is a free log retrieval operation binding the contract event 0xc513e7dd56e3ecfbde0e741f1308e1c0d3c9736a314d624fae79b0beec647b6c.
//
// Solidity: event SaleDeployed()
func (_TokenSale *TokenSaleFilterer) FilterSaleDeployed(opts *bind.FilterOpts) (*TokenSaleSaleDeployedIterator, error) {

	logs, sub, err := _TokenSale.contract.FilterLogs(opts, "SaleDeployed")
	if err != nil {
		return nil, err
	}
	return &TokenSaleSaleDeployedIterator{contract: _TokenSale.contract, event: "SaleDeployed", logs: logs, sub: sub}, nil
}

// WatchSaleDeployed is a free log subscription operation binding the contract event 0xc513e7dd56e3ecfbde0e741f1308e1c0d3c9736a314d624fae79b0beec647b6c.
//
// Solidity: event SaleDeployed()
func (_TokenSale *TokenSaleFilterer) WatchSaleDeployed(opts *bind.WatchOpts, sink chan<- *TokenSaleSaleDeployed) (event.Subscription, error) {

	logs, sub, err := _TokenSale.contract.WatchLogs(opts, "SaleDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSaleSaleDeployed)
				if err := _TokenSale.contract.UnpackLog(event, "SaleDeployed", log); err != nil {
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

// TokenSaleSaleStartedIterator is returned from FilterSaleStarted and is used to iterate over the raw logs and unpacked data for SaleStarted events raised by the TokenSale contract.
type TokenSaleSaleStartedIterator struct {
	Event *TokenSaleSaleStarted // Event containing the contract specifics and raw log

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
func (it *TokenSaleSaleStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSaleSaleStarted)
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
		it.Event = new(TokenSaleSaleStarted)
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
func (it *TokenSaleSaleStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSaleSaleStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSaleSaleStarted represents a SaleStarted event raised by the TokenSale contract.
type TokenSaleSaleStarted struct {
	StartTime *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSaleStarted is a free log retrieval operation binding the contract event 0xa78c547613f6306e7a70d1bd161c18a496cae1eeb8d4f9e58b60d69ad72ddf58.
//
// Solidity: event SaleStarted(_startTime uint256)
func (_TokenSale *TokenSaleFilterer) FilterSaleStarted(opts *bind.FilterOpts) (*TokenSaleSaleStartedIterator, error) {

	logs, sub, err := _TokenSale.contract.FilterLogs(opts, "SaleStarted")
	if err != nil {
		return nil, err
	}
	return &TokenSaleSaleStartedIterator{contract: _TokenSale.contract, event: "SaleStarted", logs: logs, sub: sub}, nil
}

// WatchSaleStarted is a free log subscription operation binding the contract event 0xa78c547613f6306e7a70d1bd161c18a496cae1eeb8d4f9e58b60d69ad72ddf58.
//
// Solidity: event SaleStarted(_startTime uint256)
func (_TokenSale *TokenSaleFilterer) WatchSaleStarted(opts *bind.WatchOpts, sink chan<- *TokenSaleSaleStarted) (event.Subscription, error) {

	logs, sub, err := _TokenSale.contract.WatchLogs(opts, "SaleStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSaleSaleStarted)
				if err := _TokenSale.contract.UnpackLog(event, "SaleStarted", log); err != nil {
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

// TokenSaleSaleTerminatedIterator is returned from FilterSaleTerminated and is used to iterate over the raw logs and unpacked data for SaleTerminated events raised by the TokenSale contract.
type TokenSaleSaleTerminatedIterator struct {
	Event *TokenSaleSaleTerminated // Event containing the contract specifics and raw log

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
func (it *TokenSaleSaleTerminatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSaleSaleTerminated)
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
		it.Event = new(TokenSaleSaleTerminated)
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
func (it *TokenSaleSaleTerminatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSaleSaleTerminatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSaleSaleTerminated represents a SaleTerminated event raised by the TokenSale contract.
type TokenSaleSaleTerminated struct {
	EndTime *big.Int
	Reason  uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSaleTerminated is a free log retrieval operation binding the contract event 0xdd7de22ead9f1d549ffaf3932971732af4a50599b882abc7e091ab84ce9865c9.
//
// Solidity: event SaleTerminated(_endTime uint256, reason uint8)
func (_TokenSale *TokenSaleFilterer) FilterSaleTerminated(opts *bind.FilterOpts) (*TokenSaleSaleTerminatedIterator, error) {

	logs, sub, err := _TokenSale.contract.FilterLogs(opts, "SaleTerminated")
	if err != nil {
		return nil, err
	}
	return &TokenSaleSaleTerminatedIterator{contract: _TokenSale.contract, event: "SaleTerminated", logs: logs, sub: sub}, nil
}

// WatchSaleTerminated is a free log subscription operation binding the contract event 0xdd7de22ead9f1d549ffaf3932971732af4a50599b882abc7e091ab84ce9865c9.
//
// Solidity: event SaleTerminated(_endTime uint256, reason uint8)
func (_TokenSale *TokenSaleFilterer) WatchSaleTerminated(opts *bind.WatchOpts, sink chan<- *TokenSaleSaleTerminated) (event.Subscription, error) {

	logs, sub, err := _TokenSale.contract.WatchLogs(opts, "SaleTerminated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSaleSaleTerminated)
				if err := _TokenSale.contract.UnpackLog(event, "SaleTerminated", log); err != nil {
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

// TokenSaleTokensPurchasedIterator is returned from FilterTokensPurchased and is used to iterate over the raw logs and unpacked data for TokensPurchased events raised by the TokenSale contract.
type TokenSaleTokensPurchasedIterator struct {
	Event *TokenSaleTokensPurchased // Event containing the contract specifics and raw log

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
func (it *TokenSaleTokensPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSaleTokensPurchased)
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
		it.Event = new(TokenSaleTokensPurchased)
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
func (it *TokenSaleTokensPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSaleTokensPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSaleTokensPurchased represents a TokensPurchased event raised by the TokenSale contract.
type TokenSaleTokensPurchased struct {
	Beneficiary common.Address
	UnitCount   *big.Int
	IsBitcoin   bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensPurchased is a free log retrieval operation binding the contract event 0x0ea576a332c77589cbeda817b213b2523a657301fb817b66aaab967a646d8b10.
//
// Solidity: event TokensPurchased(_beneficiary indexed address, unitCount uint256, isBitcoin bool)
func (_TokenSale *TokenSaleFilterer) FilterTokensPurchased(opts *bind.FilterOpts, _beneficiary []common.Address) (*TokenSaleTokensPurchasedIterator, error) {

	var _beneficiaryRule []interface{}
	for _, _beneficiaryItem := range _beneficiary {
		_beneficiaryRule = append(_beneficiaryRule, _beneficiaryItem)
	}

	logs, sub, err := _TokenSale.contract.FilterLogs(opts, "TokensPurchased", _beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &TokenSaleTokensPurchasedIterator{contract: _TokenSale.contract, event: "TokensPurchased", logs: logs, sub: sub}, nil
}

// WatchTokensPurchased is a free log subscription operation binding the contract event 0x0ea576a332c77589cbeda817b213b2523a657301fb817b66aaab967a646d8b10.
//
// Solidity: event TokensPurchased(_beneficiary indexed address, unitCount uint256, isBitcoin bool)
func (_TokenSale *TokenSaleFilterer) WatchTokensPurchased(opts *bind.WatchOpts, sink chan<- *TokenSaleTokensPurchased, _beneficiary []common.Address) (event.Subscription, error) {

	var _beneficiaryRule []interface{}
	for _, _beneficiaryItem := range _beneficiary {
		_beneficiaryRule = append(_beneficiaryRule, _beneficiaryItem)
	}

	logs, sub, err := _TokenSale.contract.WatchLogs(opts, "TokensPurchased", _beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSaleTokensPurchased)
				if err := _TokenSale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
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
