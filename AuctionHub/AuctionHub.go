// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AuctionHub

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

// AuctionHubAuctionData is an auto generated low-level Go binding around an user-defined struct.
type AuctionHubAuctionData struct {
	StartPrice    *big.Int
	StartAt       *big.Int
	EndAt         *big.Int
	Fee           *big.Int
	HighestBidder common.Address
	ItemId        *big.Int
	ItemContract  common.Address
}

// AuctionHubMetaData contains all meta data concerning the AuctionHub contract.
var AuctionHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_auctionFee\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_auctionPeriod\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPrice\",\"type\":\"uint256\"}],\"name\":\"AuctionCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionFee\",\"type\":\"uint256\"}],\"name\":\"AuctionFeeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionPeriod\",\"type\":\"uint256\"}],\"name\":\"AuctionPeriodChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"startAt\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"endAt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPrice\",\"type\":\"uint256\"}],\"name\":\"AuctionStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"highestBidder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"endPrice\",\"type\":\"uint256\"}],\"name\":\"AuctionStoped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newHighestBidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newHighestBid\",\"type\":\"uint256\"}],\"name\":\"BidIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_bid\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"getAuctionData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"highestBidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"itemContract\",\"type\":\"address\"}],\"internalType\":\"structAuctionHub.AuctionData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuctionFee\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuctionPeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"getBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newSeller\",\"type\":\"address\"}],\"name\":\"move\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_auctionFee\",\"type\":\"uint32\"}],\"name\":\"setAuctionFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_auctionPeriod\",\"type\":\"uint32\"}],\"name\":\"setAuctionPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"stopAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AuctionHubABI is the input ABI used to generate the binding from.
// Deprecated: Use AuctionHubMetaData.ABI instead.
var AuctionHubABI = AuctionHubMetaData.ABI

// AuctionHub is an auto generated Go binding around an Ethereum contract.
type AuctionHub struct {
	AuctionHubCaller     // Read-only binding to the contract
	AuctionHubTransactor // Write-only binding to the contract
	AuctionHubFilterer   // Log filterer for contract events
}

// AuctionHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuctionHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuctionHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuctionHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuctionHubSession struct {
	Contract     *AuctionHub       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuctionHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuctionHubCallerSession struct {
	Contract *AuctionHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AuctionHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuctionHubTransactorSession struct {
	Contract     *AuctionHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AuctionHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuctionHubRaw struct {
	Contract *AuctionHub // Generic contract binding to access the raw methods on
}

// AuctionHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuctionHubCallerRaw struct {
	Contract *AuctionHubCaller // Generic read-only contract binding to access the raw methods on
}

// AuctionHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuctionHubTransactorRaw struct {
	Contract *AuctionHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuctionHub creates a new instance of AuctionHub, bound to a specific deployed contract.
func NewAuctionHub(address common.Address, backend bind.ContractBackend) (*AuctionHub, error) {
	contract, err := bindAuctionHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AuctionHub{AuctionHubCaller: AuctionHubCaller{contract: contract}, AuctionHubTransactor: AuctionHubTransactor{contract: contract}, AuctionHubFilterer: AuctionHubFilterer{contract: contract}}, nil
}

// NewAuctionHubCaller creates a new read-only instance of AuctionHub, bound to a specific deployed contract.
func NewAuctionHubCaller(address common.Address, caller bind.ContractCaller) (*AuctionHubCaller, error) {
	contract, err := bindAuctionHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionHubCaller{contract: contract}, nil
}

// NewAuctionHubTransactor creates a new write-only instance of AuctionHub, bound to a specific deployed contract.
func NewAuctionHubTransactor(address common.Address, transactor bind.ContractTransactor) (*AuctionHubTransactor, error) {
	contract, err := bindAuctionHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionHubTransactor{contract: contract}, nil
}

// NewAuctionHubFilterer creates a new log filterer instance of AuctionHub, bound to a specific deployed contract.
func NewAuctionHubFilterer(address common.Address, filterer bind.ContractFilterer) (*AuctionHubFilterer, error) {
	contract, err := bindAuctionHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuctionHubFilterer{contract: contract}, nil
}

// bindAuctionHub binds a generic wrapper to an already deployed contract.
func bindAuctionHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuctionHubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuctionHub *AuctionHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuctionHub.Contract.AuctionHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuctionHub *AuctionHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionHub.Contract.AuctionHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuctionHub *AuctionHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuctionHub.Contract.AuctionHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuctionHub *AuctionHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuctionHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuctionHub *AuctionHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuctionHub *AuctionHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuctionHub.Contract.contract.Transact(opts, method, params...)
}

// GetAuctionData is a free data retrieval call binding the contract method 0x64bcbc31.
//
// Solidity: function getAuctionData(address seller) view returns((uint256,uint256,uint256,uint256,address,uint256,address))
func (_AuctionHub *AuctionHubCaller) GetAuctionData(opts *bind.CallOpts, seller common.Address) (AuctionHubAuctionData, error) {
	var out []interface{}
	err := _AuctionHub.contract.Call(opts, &out, "getAuctionData", seller)

	if err != nil {
		return *new(AuctionHubAuctionData), err
	}

	out0 := *abi.ConvertType(out[0], new(AuctionHubAuctionData)).(*AuctionHubAuctionData)

	return out0, err

}

// GetAuctionData is a free data retrieval call binding the contract method 0x64bcbc31.
//
// Solidity: function getAuctionData(address seller) view returns((uint256,uint256,uint256,uint256,address,uint256,address))
func (_AuctionHub *AuctionHubSession) GetAuctionData(seller common.Address) (AuctionHubAuctionData, error) {
	return _AuctionHub.Contract.GetAuctionData(&_AuctionHub.CallOpts, seller)
}

// GetAuctionData is a free data retrieval call binding the contract method 0x64bcbc31.
//
// Solidity: function getAuctionData(address seller) view returns((uint256,uint256,uint256,uint256,address,uint256,address))
func (_AuctionHub *AuctionHubCallerSession) GetAuctionData(seller common.Address) (AuctionHubAuctionData, error) {
	return _AuctionHub.Contract.GetAuctionData(&_AuctionHub.CallOpts, seller)
}

// GetAuctionFee is a free data retrieval call binding the contract method 0x1031ca44.
//
// Solidity: function getAuctionFee() view returns(uint32)
func (_AuctionHub *AuctionHubCaller) GetAuctionFee(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AuctionHub.contract.Call(opts, &out, "getAuctionFee")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetAuctionFee is a free data retrieval call binding the contract method 0x1031ca44.
//
// Solidity: function getAuctionFee() view returns(uint32)
func (_AuctionHub *AuctionHubSession) GetAuctionFee() (uint32, error) {
	return _AuctionHub.Contract.GetAuctionFee(&_AuctionHub.CallOpts)
}

// GetAuctionFee is a free data retrieval call binding the contract method 0x1031ca44.
//
// Solidity: function getAuctionFee() view returns(uint32)
func (_AuctionHub *AuctionHubCallerSession) GetAuctionFee() (uint32, error) {
	return _AuctionHub.Contract.GetAuctionFee(&_AuctionHub.CallOpts)
}

// GetAuctionPeriod is a free data retrieval call binding the contract method 0xe20ba48b.
//
// Solidity: function getAuctionPeriod() view returns(uint32)
func (_AuctionHub *AuctionHubCaller) GetAuctionPeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AuctionHub.contract.Call(opts, &out, "getAuctionPeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetAuctionPeriod is a free data retrieval call binding the contract method 0xe20ba48b.
//
// Solidity: function getAuctionPeriod() view returns(uint32)
func (_AuctionHub *AuctionHubSession) GetAuctionPeriod() (uint32, error) {
	return _AuctionHub.Contract.GetAuctionPeriod(&_AuctionHub.CallOpts)
}

// GetAuctionPeriod is a free data retrieval call binding the contract method 0xe20ba48b.
//
// Solidity: function getAuctionPeriod() view returns(uint32)
func (_AuctionHub *AuctionHubCallerSession) GetAuctionPeriod() (uint32, error) {
	return _AuctionHub.Contract.GetAuctionPeriod(&_AuctionHub.CallOpts)
}

// GetBid is a free data retrieval call binding the contract method 0xc8b342ab.
//
// Solidity: function getBid(address seller) view returns(uint256)
func (_AuctionHub *AuctionHubCaller) GetBid(opts *bind.CallOpts, seller common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AuctionHub.contract.Call(opts, &out, "getBid", seller)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBid is a free data retrieval call binding the contract method 0xc8b342ab.
//
// Solidity: function getBid(address seller) view returns(uint256)
func (_AuctionHub *AuctionHubSession) GetBid(seller common.Address) (*big.Int, error) {
	return _AuctionHub.Contract.GetBid(&_AuctionHub.CallOpts, seller)
}

// GetBid is a free data retrieval call binding the contract method 0xc8b342ab.
//
// Solidity: function getBid(address seller) view returns(uint256)
func (_AuctionHub *AuctionHubCallerSession) GetBid(seller common.Address) (*big.Int, error) {
	return _AuctionHub.Contract.GetBid(&_AuctionHub.CallOpts, seller)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuctionHub *AuctionHubCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuctionHub.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuctionHub *AuctionHubSession) Owner() (common.Address, error) {
	return _AuctionHub.Contract.Owner(&_AuctionHub.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuctionHub *AuctionHubCallerSession) Owner() (common.Address, error) {
	return _AuctionHub.Contract.Owner(&_AuctionHub.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AuctionHub *AuctionHubCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AuctionHub.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AuctionHub *AuctionHubSession) Paused() (bool, error) {
	return _AuctionHub.Contract.Paused(&_AuctionHub.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AuctionHub *AuctionHubCallerSession) Paused() (bool, error) {
	return _AuctionHub.Contract.Paused(&_AuctionHub.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x59d667a5.
//
// Solidity: function bid(address _seller, uint256 _bid) returns()
func (_AuctionHub *AuctionHubTransactor) Bid(opts *bind.TransactOpts, _seller common.Address, _bid *big.Int) (*types.Transaction, error) {
	return _AuctionHub.contract.Transact(opts, "bid", _seller, _bid)
}

// Bid is a paid mutator transaction binding the contract method 0x59d667a5.
//
// Solidity: function bid(address _seller, uint256 _bid) returns()
func (_AuctionHub *AuctionHubSession) Bid(_seller common.Address, _bid *big.Int) (*types.Transaction, error) {
	return _AuctionHub.Contract.Bid(&_AuctionHub.TransactOpts, _seller, _bid)
}

// Bid is a paid mutator transaction binding the contract method 0x59d667a5.
//
// Solidity: function bid(address _seller, uint256 _bid) returns()
func (_AuctionHub *AuctionHubTransactorSession) Bid(_seller common.Address, _bid *big.Int) (*types.Transaction, error) {
	return _AuctionHub.Contract.Bid(&_AuctionHub.TransactOpts, _seller, _bid)
}

// Move is a paid mutator transaction binding the contract method 0x322265a8.
//
// Solidity: function move(address seller, address newSeller) returns()
func (_AuctionHub *AuctionHubTransactor) Move(opts *bind.TransactOpts, seller common.Address, newSeller common.Address) (*types.Transaction, error) {
	return _AuctionHub.contract.Transact(opts, "move", seller, newSeller)
}

// Move is a paid mutator transaction binding the contract method 0x322265a8.
//
// Solidity: function move(address seller, address newSeller) returns()
func (_AuctionHub *AuctionHubSession) Move(seller common.Address, newSeller common.Address) (*types.Transaction, error) {
	return _AuctionHub.Contract.Move(&_AuctionHub.TransactOpts, seller, newSeller)
}

// Move is a paid mutator transaction binding the contract method 0x322265a8.
//
// Solidity: function move(address seller, address newSeller) returns()
func (_AuctionHub *AuctionHubTransactorSession) Move(seller common.Address, newSeller common.Address) (*types.Transaction, error) {
	return _AuctionHub.Contract.Move(&_AuctionHub.TransactOpts, seller, newSeller)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_AuctionHub *AuctionHubTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _AuctionHub.contract.Transact(opts, "onERC721Received", operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_AuctionHub *AuctionHubSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _AuctionHub.Contract.OnERC721Received(&_AuctionHub.TransactOpts, operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_AuctionHub *AuctionHubTransactorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _AuctionHub.Contract.OnERC721Received(&_AuctionHub.TransactOpts, operator, from, tokenId, data)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AuctionHub *AuctionHubTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionHub.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AuctionHub *AuctionHubSession) Pause() (*types.Transaction, error) {
	return _AuctionHub.Contract.Pause(&_AuctionHub.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AuctionHub *AuctionHubTransactorSession) Pause() (*types.Transaction, error) {
	return _AuctionHub.Contract.Pause(&_AuctionHub.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuctionHub *AuctionHubTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionHub.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuctionHub *AuctionHubSession) RenounceOwnership() (*types.Transaction, error) {
	return _AuctionHub.Contract.RenounceOwnership(&_AuctionHub.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuctionHub *AuctionHubTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AuctionHub.Contract.RenounceOwnership(&_AuctionHub.TransactOpts)
}

// SetAuctionFee is a paid mutator transaction binding the contract method 0x7c602468.
//
// Solidity: function setAuctionFee(uint32 _auctionFee) returns()
func (_AuctionHub *AuctionHubTransactor) SetAuctionFee(opts *bind.TransactOpts, _auctionFee uint32) (*types.Transaction, error) {
	return _AuctionHub.contract.Transact(opts, "setAuctionFee", _auctionFee)
}

// SetAuctionFee is a paid mutator transaction binding the contract method 0x7c602468.
//
// Solidity: function setAuctionFee(uint32 _auctionFee) returns()
func (_AuctionHub *AuctionHubSession) SetAuctionFee(_auctionFee uint32) (*types.Transaction, error) {
	return _AuctionHub.Contract.SetAuctionFee(&_AuctionHub.TransactOpts, _auctionFee)
}

// SetAuctionFee is a paid mutator transaction binding the contract method 0x7c602468.
//
// Solidity: function setAuctionFee(uint32 _auctionFee) returns()
func (_AuctionHub *AuctionHubTransactorSession) SetAuctionFee(_auctionFee uint32) (*types.Transaction, error) {
	return _AuctionHub.Contract.SetAuctionFee(&_AuctionHub.TransactOpts, _auctionFee)
}

// SetAuctionPeriod is a paid mutator transaction binding the contract method 0xa21a3d70.
//
// Solidity: function setAuctionPeriod(uint32 _auctionPeriod) returns()
func (_AuctionHub *AuctionHubTransactor) SetAuctionPeriod(opts *bind.TransactOpts, _auctionPeriod uint32) (*types.Transaction, error) {
	return _AuctionHub.contract.Transact(opts, "setAuctionPeriod", _auctionPeriod)
}

// SetAuctionPeriod is a paid mutator transaction binding the contract method 0xa21a3d70.
//
// Solidity: function setAuctionPeriod(uint32 _auctionPeriod) returns()
func (_AuctionHub *AuctionHubSession) SetAuctionPeriod(_auctionPeriod uint32) (*types.Transaction, error) {
	return _AuctionHub.Contract.SetAuctionPeriod(&_AuctionHub.TransactOpts, _auctionPeriod)
}

// SetAuctionPeriod is a paid mutator transaction binding the contract method 0xa21a3d70.
//
// Solidity: function setAuctionPeriod(uint32 _auctionPeriod) returns()
func (_AuctionHub *AuctionHubTransactorSession) SetAuctionPeriod(_auctionPeriod uint32) (*types.Transaction, error) {
	return _AuctionHub.Contract.SetAuctionPeriod(&_AuctionHub.TransactOpts, _auctionPeriod)
}

// StopAuction is a paid mutator transaction binding the contract method 0x189bd1ae.
//
// Solidity: function stopAuction(address seller) returns()
func (_AuctionHub *AuctionHubTransactor) StopAuction(opts *bind.TransactOpts, seller common.Address) (*types.Transaction, error) {
	return _AuctionHub.contract.Transact(opts, "stopAuction", seller)
}

// StopAuction is a paid mutator transaction binding the contract method 0x189bd1ae.
//
// Solidity: function stopAuction(address seller) returns()
func (_AuctionHub *AuctionHubSession) StopAuction(seller common.Address) (*types.Transaction, error) {
	return _AuctionHub.Contract.StopAuction(&_AuctionHub.TransactOpts, seller)
}

// StopAuction is a paid mutator transaction binding the contract method 0x189bd1ae.
//
// Solidity: function stopAuction(address seller) returns()
func (_AuctionHub *AuctionHubTransactorSession) StopAuction(seller common.Address) (*types.Transaction, error) {
	return _AuctionHub.Contract.StopAuction(&_AuctionHub.TransactOpts, seller)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuctionHub *AuctionHubTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AuctionHub.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuctionHub *AuctionHubSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AuctionHub.Contract.TransferOwnership(&_AuctionHub.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuctionHub *AuctionHubTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AuctionHub.Contract.TransferOwnership(&_AuctionHub.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AuctionHub *AuctionHubTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionHub.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AuctionHub *AuctionHubSession) Unpause() (*types.Transaction, error) {
	return _AuctionHub.Contract.Unpause(&_AuctionHub.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AuctionHub *AuctionHubTransactorSession) Unpause() (*types.Transaction, error) {
	return _AuctionHub.Contract.Unpause(&_AuctionHub.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address seller) returns()
func (_AuctionHub *AuctionHubTransactor) Withdraw(opts *bind.TransactOpts, seller common.Address) (*types.Transaction, error) {
	return _AuctionHub.contract.Transact(opts, "withdraw", seller)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address seller) returns()
func (_AuctionHub *AuctionHubSession) Withdraw(seller common.Address) (*types.Transaction, error) {
	return _AuctionHub.Contract.Withdraw(&_AuctionHub.TransactOpts, seller)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address seller) returns()
func (_AuctionHub *AuctionHubTransactorSession) Withdraw(seller common.Address) (*types.Transaction, error) {
	return _AuctionHub.Contract.Withdraw(&_AuctionHub.TransactOpts, seller)
}

// AuctionHubAuctionCanceledIterator is returned from FilterAuctionCanceled and is used to iterate over the raw logs and unpacked data for AuctionCanceled events raised by the AuctionHub contract.
type AuctionHubAuctionCanceledIterator struct {
	Event *AuctionHubAuctionCanceled // Event containing the contract specifics and raw log

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
func (it *AuctionHubAuctionCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionHubAuctionCanceled)
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
		it.Event = new(AuctionHubAuctionCanceled)
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
func (it *AuctionHubAuctionCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionHubAuctionCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionHubAuctionCanceled represents a AuctionCanceled event raised by the AuctionHub contract.
type AuctionHubAuctionCanceled struct {
	Seller     common.Address
	ItemId     *big.Int
	StartPrice *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionCanceled is a free log retrieval operation binding the contract event 0x4c776500ca4dafe4510855da5de7f14e5c0685abd4c67935f5ad0e5acb40f16a.
//
// Solidity: event AuctionCanceled(address seller, uint256 itemId, uint256 startPrice)
func (_AuctionHub *AuctionHubFilterer) FilterAuctionCanceled(opts *bind.FilterOpts) (*AuctionHubAuctionCanceledIterator, error) {

	logs, sub, err := _AuctionHub.contract.FilterLogs(opts, "AuctionCanceled")
	if err != nil {
		return nil, err
	}
	return &AuctionHubAuctionCanceledIterator{contract: _AuctionHub.contract, event: "AuctionCanceled", logs: logs, sub: sub}, nil
}

// WatchAuctionCanceled is a free log subscription operation binding the contract event 0x4c776500ca4dafe4510855da5de7f14e5c0685abd4c67935f5ad0e5acb40f16a.
//
// Solidity: event AuctionCanceled(address seller, uint256 itemId, uint256 startPrice)
func (_AuctionHub *AuctionHubFilterer) WatchAuctionCanceled(opts *bind.WatchOpts, sink chan<- *AuctionHubAuctionCanceled) (event.Subscription, error) {

	logs, sub, err := _AuctionHub.contract.WatchLogs(opts, "AuctionCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionHubAuctionCanceled)
				if err := _AuctionHub.contract.UnpackLog(event, "AuctionCanceled", log); err != nil {
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

// ParseAuctionCanceled is a log parse operation binding the contract event 0x4c776500ca4dafe4510855da5de7f14e5c0685abd4c67935f5ad0e5acb40f16a.
//
// Solidity: event AuctionCanceled(address seller, uint256 itemId, uint256 startPrice)
func (_AuctionHub *AuctionHubFilterer) ParseAuctionCanceled(log types.Log) (*AuctionHubAuctionCanceled, error) {
	event := new(AuctionHubAuctionCanceled)
	if err := _AuctionHub.contract.UnpackLog(event, "AuctionCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionHubAuctionFeeChangedIterator is returned from FilterAuctionFeeChanged and is used to iterate over the raw logs and unpacked data for AuctionFeeChanged events raised by the AuctionHub contract.
type AuctionHubAuctionFeeChangedIterator struct {
	Event *AuctionHubAuctionFeeChanged // Event containing the contract specifics and raw log

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
func (it *AuctionHubAuctionFeeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionHubAuctionFeeChanged)
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
		it.Event = new(AuctionHubAuctionFeeChanged)
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
func (it *AuctionHubAuctionFeeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionHubAuctionFeeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionHubAuctionFeeChanged represents a AuctionFeeChanged event raised by the AuctionHub contract.
type AuctionHubAuctionFeeChanged struct {
	AuctionFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionFeeChanged is a free log retrieval operation binding the contract event 0xec7394e8821e6f8bc7f15d16250718803f808e06e4627ebba8b48fa6413aec47.
//
// Solidity: event AuctionFeeChanged(uint256 indexed auctionFee)
func (_AuctionHub *AuctionHubFilterer) FilterAuctionFeeChanged(opts *bind.FilterOpts, auctionFee []*big.Int) (*AuctionHubAuctionFeeChangedIterator, error) {

	var auctionFeeRule []interface{}
	for _, auctionFeeItem := range auctionFee {
		auctionFeeRule = append(auctionFeeRule, auctionFeeItem)
	}

	logs, sub, err := _AuctionHub.contract.FilterLogs(opts, "AuctionFeeChanged", auctionFeeRule)
	if err != nil {
		return nil, err
	}
	return &AuctionHubAuctionFeeChangedIterator{contract: _AuctionHub.contract, event: "AuctionFeeChanged", logs: logs, sub: sub}, nil
}

// WatchAuctionFeeChanged is a free log subscription operation binding the contract event 0xec7394e8821e6f8bc7f15d16250718803f808e06e4627ebba8b48fa6413aec47.
//
// Solidity: event AuctionFeeChanged(uint256 indexed auctionFee)
func (_AuctionHub *AuctionHubFilterer) WatchAuctionFeeChanged(opts *bind.WatchOpts, sink chan<- *AuctionHubAuctionFeeChanged, auctionFee []*big.Int) (event.Subscription, error) {

	var auctionFeeRule []interface{}
	for _, auctionFeeItem := range auctionFee {
		auctionFeeRule = append(auctionFeeRule, auctionFeeItem)
	}

	logs, sub, err := _AuctionHub.contract.WatchLogs(opts, "AuctionFeeChanged", auctionFeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionHubAuctionFeeChanged)
				if err := _AuctionHub.contract.UnpackLog(event, "AuctionFeeChanged", log); err != nil {
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

// ParseAuctionFeeChanged is a log parse operation binding the contract event 0xec7394e8821e6f8bc7f15d16250718803f808e06e4627ebba8b48fa6413aec47.
//
// Solidity: event AuctionFeeChanged(uint256 indexed auctionFee)
func (_AuctionHub *AuctionHubFilterer) ParseAuctionFeeChanged(log types.Log) (*AuctionHubAuctionFeeChanged, error) {
	event := new(AuctionHubAuctionFeeChanged)
	if err := _AuctionHub.contract.UnpackLog(event, "AuctionFeeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionHubAuctionPeriodChangedIterator is returned from FilterAuctionPeriodChanged and is used to iterate over the raw logs and unpacked data for AuctionPeriodChanged events raised by the AuctionHub contract.
type AuctionHubAuctionPeriodChangedIterator struct {
	Event *AuctionHubAuctionPeriodChanged // Event containing the contract specifics and raw log

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
func (it *AuctionHubAuctionPeriodChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionHubAuctionPeriodChanged)
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
		it.Event = new(AuctionHubAuctionPeriodChanged)
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
func (it *AuctionHubAuctionPeriodChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionHubAuctionPeriodChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionHubAuctionPeriodChanged represents a AuctionPeriodChanged event raised by the AuctionHub contract.
type AuctionHubAuctionPeriodChanged struct {
	AuctionPeriod *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAuctionPeriodChanged is a free log retrieval operation binding the contract event 0xd7399c076335610df8c3a154c053c06fdea4595034fca437d7b860ffec081192.
//
// Solidity: event AuctionPeriodChanged(uint256 indexed auctionPeriod)
func (_AuctionHub *AuctionHubFilterer) FilterAuctionPeriodChanged(opts *bind.FilterOpts, auctionPeriod []*big.Int) (*AuctionHubAuctionPeriodChangedIterator, error) {

	var auctionPeriodRule []interface{}
	for _, auctionPeriodItem := range auctionPeriod {
		auctionPeriodRule = append(auctionPeriodRule, auctionPeriodItem)
	}

	logs, sub, err := _AuctionHub.contract.FilterLogs(opts, "AuctionPeriodChanged", auctionPeriodRule)
	if err != nil {
		return nil, err
	}
	return &AuctionHubAuctionPeriodChangedIterator{contract: _AuctionHub.contract, event: "AuctionPeriodChanged", logs: logs, sub: sub}, nil
}

// WatchAuctionPeriodChanged is a free log subscription operation binding the contract event 0xd7399c076335610df8c3a154c053c06fdea4595034fca437d7b860ffec081192.
//
// Solidity: event AuctionPeriodChanged(uint256 indexed auctionPeriod)
func (_AuctionHub *AuctionHubFilterer) WatchAuctionPeriodChanged(opts *bind.WatchOpts, sink chan<- *AuctionHubAuctionPeriodChanged, auctionPeriod []*big.Int) (event.Subscription, error) {

	var auctionPeriodRule []interface{}
	for _, auctionPeriodItem := range auctionPeriod {
		auctionPeriodRule = append(auctionPeriodRule, auctionPeriodItem)
	}

	logs, sub, err := _AuctionHub.contract.WatchLogs(opts, "AuctionPeriodChanged", auctionPeriodRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionHubAuctionPeriodChanged)
				if err := _AuctionHub.contract.UnpackLog(event, "AuctionPeriodChanged", log); err != nil {
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

// ParseAuctionPeriodChanged is a log parse operation binding the contract event 0xd7399c076335610df8c3a154c053c06fdea4595034fca437d7b860ffec081192.
//
// Solidity: event AuctionPeriodChanged(uint256 indexed auctionPeriod)
func (_AuctionHub *AuctionHubFilterer) ParseAuctionPeriodChanged(log types.Log) (*AuctionHubAuctionPeriodChanged, error) {
	event := new(AuctionHubAuctionPeriodChanged)
	if err := _AuctionHub.contract.UnpackLog(event, "AuctionPeriodChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionHubAuctionStartedIterator is returned from FilterAuctionStarted and is used to iterate over the raw logs and unpacked data for AuctionStarted events raised by the AuctionHub contract.
type AuctionHubAuctionStartedIterator struct {
	Event *AuctionHubAuctionStarted // Event containing the contract specifics and raw log

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
func (it *AuctionHubAuctionStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionHubAuctionStarted)
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
		it.Event = new(AuctionHubAuctionStarted)
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
func (it *AuctionHubAuctionStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionHubAuctionStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionHubAuctionStarted represents a AuctionStarted event raised by the AuctionHub contract.
type AuctionHubAuctionStarted struct {
	Seller     common.Address
	StartAt    *big.Int
	EndAt      *big.Int
	StartPrice *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionStarted is a free log retrieval operation binding the contract event 0xc033a93229c068dace48c5b4afbbb99230dba9c9234b9153116b5354e361fc44.
//
// Solidity: event AuctionStarted(address indexed seller, uint256 indexed startAt, uint256 indexed endAt, uint256 startPrice)
func (_AuctionHub *AuctionHubFilterer) FilterAuctionStarted(opts *bind.FilterOpts, seller []common.Address, startAt []*big.Int, endAt []*big.Int) (*AuctionHubAuctionStartedIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var startAtRule []interface{}
	for _, startAtItem := range startAt {
		startAtRule = append(startAtRule, startAtItem)
	}
	var endAtRule []interface{}
	for _, endAtItem := range endAt {
		endAtRule = append(endAtRule, endAtItem)
	}

	logs, sub, err := _AuctionHub.contract.FilterLogs(opts, "AuctionStarted", sellerRule, startAtRule, endAtRule)
	if err != nil {
		return nil, err
	}
	return &AuctionHubAuctionStartedIterator{contract: _AuctionHub.contract, event: "AuctionStarted", logs: logs, sub: sub}, nil
}

// WatchAuctionStarted is a free log subscription operation binding the contract event 0xc033a93229c068dace48c5b4afbbb99230dba9c9234b9153116b5354e361fc44.
//
// Solidity: event AuctionStarted(address indexed seller, uint256 indexed startAt, uint256 indexed endAt, uint256 startPrice)
func (_AuctionHub *AuctionHubFilterer) WatchAuctionStarted(opts *bind.WatchOpts, sink chan<- *AuctionHubAuctionStarted, seller []common.Address, startAt []*big.Int, endAt []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var startAtRule []interface{}
	for _, startAtItem := range startAt {
		startAtRule = append(startAtRule, startAtItem)
	}
	var endAtRule []interface{}
	for _, endAtItem := range endAt {
		endAtRule = append(endAtRule, endAtItem)
	}

	logs, sub, err := _AuctionHub.contract.WatchLogs(opts, "AuctionStarted", sellerRule, startAtRule, endAtRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionHubAuctionStarted)
				if err := _AuctionHub.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
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

// ParseAuctionStarted is a log parse operation binding the contract event 0xc033a93229c068dace48c5b4afbbb99230dba9c9234b9153116b5354e361fc44.
//
// Solidity: event AuctionStarted(address indexed seller, uint256 indexed startAt, uint256 indexed endAt, uint256 startPrice)
func (_AuctionHub *AuctionHubFilterer) ParseAuctionStarted(log types.Log) (*AuctionHubAuctionStarted, error) {
	event := new(AuctionHubAuctionStarted)
	if err := _AuctionHub.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionHubAuctionStopedIterator is returned from FilterAuctionStoped and is used to iterate over the raw logs and unpacked data for AuctionStoped events raised by the AuctionHub contract.
type AuctionHubAuctionStopedIterator struct {
	Event *AuctionHubAuctionStoped // Event containing the contract specifics and raw log

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
func (it *AuctionHubAuctionStopedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionHubAuctionStoped)
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
		it.Event = new(AuctionHubAuctionStoped)
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
func (it *AuctionHubAuctionStopedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionHubAuctionStopedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionHubAuctionStoped represents a AuctionStoped event raised by the AuctionHub contract.
type AuctionHubAuctionStoped struct {
	Seller        common.Address
	HighestBidder common.Address
	EndPrice      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAuctionStoped is a free log retrieval operation binding the contract event 0xcb33adaaf4c555df0b2365731547dcd9571a559f932a7e30f90b0a8682e968db.
//
// Solidity: event AuctionStoped(address indexed seller, address indexed highestBidder, uint256 indexed endPrice)
func (_AuctionHub *AuctionHubFilterer) FilterAuctionStoped(opts *bind.FilterOpts, seller []common.Address, highestBidder []common.Address, endPrice []*big.Int) (*AuctionHubAuctionStopedIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var highestBidderRule []interface{}
	for _, highestBidderItem := range highestBidder {
		highestBidderRule = append(highestBidderRule, highestBidderItem)
	}
	var endPriceRule []interface{}
	for _, endPriceItem := range endPrice {
		endPriceRule = append(endPriceRule, endPriceItem)
	}

	logs, sub, err := _AuctionHub.contract.FilterLogs(opts, "AuctionStoped", sellerRule, highestBidderRule, endPriceRule)
	if err != nil {
		return nil, err
	}
	return &AuctionHubAuctionStopedIterator{contract: _AuctionHub.contract, event: "AuctionStoped", logs: logs, sub: sub}, nil
}

// WatchAuctionStoped is a free log subscription operation binding the contract event 0xcb33adaaf4c555df0b2365731547dcd9571a559f932a7e30f90b0a8682e968db.
//
// Solidity: event AuctionStoped(address indexed seller, address indexed highestBidder, uint256 indexed endPrice)
func (_AuctionHub *AuctionHubFilterer) WatchAuctionStoped(opts *bind.WatchOpts, sink chan<- *AuctionHubAuctionStoped, seller []common.Address, highestBidder []common.Address, endPrice []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var highestBidderRule []interface{}
	for _, highestBidderItem := range highestBidder {
		highestBidderRule = append(highestBidderRule, highestBidderItem)
	}
	var endPriceRule []interface{}
	for _, endPriceItem := range endPrice {
		endPriceRule = append(endPriceRule, endPriceItem)
	}

	logs, sub, err := _AuctionHub.contract.WatchLogs(opts, "AuctionStoped", sellerRule, highestBidderRule, endPriceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionHubAuctionStoped)
				if err := _AuctionHub.contract.UnpackLog(event, "AuctionStoped", log); err != nil {
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

// ParseAuctionStoped is a log parse operation binding the contract event 0xcb33adaaf4c555df0b2365731547dcd9571a559f932a7e30f90b0a8682e968db.
//
// Solidity: event AuctionStoped(address indexed seller, address indexed highestBidder, uint256 indexed endPrice)
func (_AuctionHub *AuctionHubFilterer) ParseAuctionStoped(log types.Log) (*AuctionHubAuctionStoped, error) {
	event := new(AuctionHubAuctionStoped)
	if err := _AuctionHub.contract.UnpackLog(event, "AuctionStoped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionHubBidIncreasedIterator is returned from FilterBidIncreased and is used to iterate over the raw logs and unpacked data for BidIncreased events raised by the AuctionHub contract.
type AuctionHubBidIncreasedIterator struct {
	Event *AuctionHubBidIncreased // Event containing the contract specifics and raw log

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
func (it *AuctionHubBidIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionHubBidIncreased)
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
		it.Event = new(AuctionHubBidIncreased)
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
func (it *AuctionHubBidIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionHubBidIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionHubBidIncreased represents a BidIncreased event raised by the AuctionHub contract.
type AuctionHubBidIncreased struct {
	NewHighestBidder common.Address
	NewHighestBid    *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBidIncreased is a free log retrieval operation binding the contract event 0xcf84a46a93294358c23e7c87e557feb461890c72f3547fb25b455167dcea9efb.
//
// Solidity: event BidIncreased(address newHighestBidder, uint256 newHighestBid)
func (_AuctionHub *AuctionHubFilterer) FilterBidIncreased(opts *bind.FilterOpts) (*AuctionHubBidIncreasedIterator, error) {

	logs, sub, err := _AuctionHub.contract.FilterLogs(opts, "BidIncreased")
	if err != nil {
		return nil, err
	}
	return &AuctionHubBidIncreasedIterator{contract: _AuctionHub.contract, event: "BidIncreased", logs: logs, sub: sub}, nil
}

// WatchBidIncreased is a free log subscription operation binding the contract event 0xcf84a46a93294358c23e7c87e557feb461890c72f3547fb25b455167dcea9efb.
//
// Solidity: event BidIncreased(address newHighestBidder, uint256 newHighestBid)
func (_AuctionHub *AuctionHubFilterer) WatchBidIncreased(opts *bind.WatchOpts, sink chan<- *AuctionHubBidIncreased) (event.Subscription, error) {

	logs, sub, err := _AuctionHub.contract.WatchLogs(opts, "BidIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionHubBidIncreased)
				if err := _AuctionHub.contract.UnpackLog(event, "BidIncreased", log); err != nil {
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

// ParseBidIncreased is a log parse operation binding the contract event 0xcf84a46a93294358c23e7c87e557feb461890c72f3547fb25b455167dcea9efb.
//
// Solidity: event BidIncreased(address newHighestBidder, uint256 newHighestBid)
func (_AuctionHub *AuctionHubFilterer) ParseBidIncreased(log types.Log) (*AuctionHubBidIncreased, error) {
	event := new(AuctionHubBidIncreased)
	if err := _AuctionHub.contract.UnpackLog(event, "BidIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionHubOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AuctionHub contract.
type AuctionHubOwnershipTransferredIterator struct {
	Event *AuctionHubOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AuctionHubOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionHubOwnershipTransferred)
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
		it.Event = new(AuctionHubOwnershipTransferred)
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
func (it *AuctionHubOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionHubOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionHubOwnershipTransferred represents a OwnershipTransferred event raised by the AuctionHub contract.
type AuctionHubOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuctionHub *AuctionHubFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AuctionHubOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AuctionHub.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AuctionHubOwnershipTransferredIterator{contract: _AuctionHub.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuctionHub *AuctionHubFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AuctionHubOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AuctionHub.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionHubOwnershipTransferred)
				if err := _AuctionHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuctionHub *AuctionHubFilterer) ParseOwnershipTransferred(log types.Log) (*AuctionHubOwnershipTransferred, error) {
	event := new(AuctionHubOwnershipTransferred)
	if err := _AuctionHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionHubPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AuctionHub contract.
type AuctionHubPausedIterator struct {
	Event *AuctionHubPaused // Event containing the contract specifics and raw log

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
func (it *AuctionHubPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionHubPaused)
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
		it.Event = new(AuctionHubPaused)
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
func (it *AuctionHubPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionHubPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionHubPaused represents a Paused event raised by the AuctionHub contract.
type AuctionHubPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AuctionHub *AuctionHubFilterer) FilterPaused(opts *bind.FilterOpts) (*AuctionHubPausedIterator, error) {

	logs, sub, err := _AuctionHub.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AuctionHubPausedIterator{contract: _AuctionHub.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AuctionHub *AuctionHubFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AuctionHubPaused) (event.Subscription, error) {

	logs, sub, err := _AuctionHub.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionHubPaused)
				if err := _AuctionHub.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AuctionHub *AuctionHubFilterer) ParsePaused(log types.Log) (*AuctionHubPaused, error) {
	event := new(AuctionHubPaused)
	if err := _AuctionHub.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionHubUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AuctionHub contract.
type AuctionHubUnpausedIterator struct {
	Event *AuctionHubUnpaused // Event containing the contract specifics and raw log

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
func (it *AuctionHubUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionHubUnpaused)
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
		it.Event = new(AuctionHubUnpaused)
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
func (it *AuctionHubUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionHubUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionHubUnpaused represents a Unpaused event raised by the AuctionHub contract.
type AuctionHubUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AuctionHub *AuctionHubFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AuctionHubUnpausedIterator, error) {

	logs, sub, err := _AuctionHub.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AuctionHubUnpausedIterator{contract: _AuctionHub.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AuctionHub *AuctionHubFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AuctionHubUnpaused) (event.Subscription, error) {

	logs, sub, err := _AuctionHub.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionHubUnpaused)
				if err := _AuctionHub.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AuctionHub *AuctionHubFilterer) ParseUnpaused(log types.Log) (*AuctionHubUnpaused, error) {
	event := new(AuctionHubUnpaused)
	if err := _AuctionHub.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
