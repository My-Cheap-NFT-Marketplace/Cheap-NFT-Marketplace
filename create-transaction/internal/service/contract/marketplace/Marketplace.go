// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Marketplace

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

// MarketplaceAuctionData is an auto generated low-level Go binding around an user-defined struct.
type MarketplaceAuctionData struct {
	CollectionAddress common.Address
	Erc20Address      common.Address
	TokenId           *big.Int
	Bid               *big.Int
}

// MarketplaceMetaData contains all meta data concerning the Marketplace contract.
var MarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"collectionAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"internalType\":\"structMarketplace.AuctionData\",\"name\":\"auctionData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"bidderSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"ownerApprovedSig\",\"type\":\"bytes\"}],\"name\":\"finishAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610c48806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80630f96837b14610030575b600080fd5b61004a60048036038101906100459190610754565b61004c565b005b600083600001518460200151856040015186606001516040516020016100759493929190610848565b60405160208183030381529060405280519060200120905060006100aa8461009c846101ef565b61021f90919063ffffffff16565b905060008480519060200120905060006100d5856100c7846101ef565b61021f90919063ffffffff16565b9050866020015173ffffffffffffffffffffffffffffffffffffffff166323b872dd84838a606001516040518463ffffffff1660e01b815260040161011c939291906108b4565b602060405180830381600087803b15801561013657600080fd5b505af115801561014a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061016e9190610923565b50866000015173ffffffffffffffffffffffffffffffffffffffff166342842e0e82858a604001516040518463ffffffff1660e01b81526004016101b4939291906108b4565b600060405180830381600087803b1580156101ce57600080fd5b505af11580156101e2573d6000803e3d6000fd5b5050505050505050505050565b60008160405160200161020291906109d2565b604051602081830303815290604052805190602001209050919050565b600080600061022e8585610246565b9150915061023b81610298565b819250505092915050565b6000806041835114156102885760008060006020860151925060408601519150606086015160001a905061027c87828585610406565b94509450505050610291565b60006002915091505b9250929050565b600060048111156102ac576102ab6109f8565b5b8160048111156102bf576102be6109f8565b5b14156102ca57610403565b600160048111156102de576102dd6109f8565b5b8160048111156102f1576102f06109f8565b5b1415610332576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161032990610a84565b60405180910390fd5b60026004811115610346576103456109f8565b5b816004811115610359576103586109f8565b5b141561039a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039190610af0565b60405180910390fd5b600360048111156103ae576103ad6109f8565b5b8160048111156103c1576103c06109f8565b5b1415610402576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103f990610b82565b60405180910390fd5b5b50565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08360001c11156104415760006003915091506104e0565b6000600187878787604051600081526020016040526040516104669493929190610bcd565b6020604051602081039080840390855afa158015610488573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156104d7576000600192509250506104e0565b80600092509250505b94509492505050565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61054b82610502565b810181811067ffffffffffffffff8211171561056a57610569610513565b5b80604052505050565b600061057d6104e9565b90506105898282610542565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006105b98261058e565b9050919050565b6105c9816105ae565b81146105d457600080fd5b50565b6000813590506105e6816105c0565b92915050565b6000819050919050565b6105ff816105ec565b811461060a57600080fd5b50565b60008135905061061c816105f6565b92915050565b600060808284031215610638576106376104fd565b5b6106426080610573565b90506000610652848285016105d7565b6000830152506020610666848285016105d7565b602083015250604061067a8482850161060d565b604083015250606061068e8482850161060d565b60608301525092915050565b600080fd5b600080fd5b600067ffffffffffffffff8211156106bf576106be610513565b5b6106c882610502565b9050602081019050919050565b82818337600083830152505050565b60006106f76106f2846106a4565b610573565b9050828152602081018484840111156107135761071261069f565b5b61071e8482856106d5565b509392505050565b600082601f83011261073b5761073a61069a565b5b813561074b8482602086016106e4565b91505092915050565b600080600060c0848603121561076d5761076c6104f3565b5b600061077b86828701610622565b935050608084013567ffffffffffffffff81111561079c5761079b6104f8565b5b6107a886828701610726565b92505060a084013567ffffffffffffffff8111156107c9576107c86104f8565b5b6107d586828701610726565b9150509250925092565b60008160601b9050919050565b60006107f7826107df565b9050919050565b6000610809826107ec565b9050919050565b61082161081c826105ae565b6107fe565b82525050565b6000819050919050565b61084261083d826105ec565b610827565b82525050565b60006108548287610810565b6014820191506108648286610810565b6014820191506108748285610831565b6020820191506108848284610831565b60208201915081905095945050505050565b61089f816105ae565b82525050565b6108ae816105ec565b82525050565b60006060820190506108c96000830186610896565b6108d66020830185610896565b6108e360408301846108a5565b949350505050565b60008115159050919050565b610900816108eb565b811461090b57600080fd5b50565b60008151905061091d816108f7565b92915050565b600060208284031215610939576109386104f3565b5b60006109478482850161090e565b91505092915050565b600081905092915050565b7f19457468657265756d205369676e6564204d6573736167653a0a333200000000600082015250565b6000610991601c83610950565b915061099c8261095b565b601c82019050919050565b6000819050919050565b6000819050919050565b6109cc6109c7826109a7565b6109b1565b82525050565b60006109dd82610984565b91506109e982846109bb565b60208201915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600082825260208201905092915050565b7f45434453413a20696e76616c6964207369676e61747572650000000000000000600082015250565b6000610a6e601883610a27565b9150610a7982610a38565b602082019050919050565b60006020820190508181036000830152610a9d81610a61565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265206c656e67746800600082015250565b6000610ada601f83610a27565b9150610ae582610aa4565b602082019050919050565b60006020820190508181036000830152610b0981610acd565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265202773272076616c60008201527f7565000000000000000000000000000000000000000000000000000000000000602082015250565b6000610b6c602283610a27565b9150610b7782610b10565b604082019050919050565b60006020820190508181036000830152610b9b81610b5f565b9050919050565b610bab816109a7565b82525050565b600060ff82169050919050565b610bc781610bb1565b82525050565b6000608082019050610be26000830187610ba2565b610bef6020830186610bbe565b610bfc6040830185610ba2565b610c096060830184610ba2565b9594505050505056fea2646970667358221220bf72431ac45fb1f6f367c510ea238fe3ee03e834f4ed96830b346cfd77f0cdd364736f6c63430008090033",
}

// MarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketplaceMetaData.ABI instead.
var MarketplaceABI = MarketplaceMetaData.ABI

// MarketplaceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MarketplaceMetaData.Bin instead.
var MarketplaceBin = MarketplaceMetaData.Bin

// DeployMarketplace deploys a new Ethereum contract, binding an instance of Marketplace to it.
func DeployMarketplace(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Marketplace, error) {
	parsed, err := MarketplaceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MarketplaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Marketplace{MarketplaceCaller: MarketplaceCaller{contract: contract}, MarketplaceTransactor: MarketplaceTransactor{contract: contract}, MarketplaceFilterer: MarketplaceFilterer{contract: contract}}, nil
}

// Marketplace is an auto generated Go binding around an Ethereum contract.
type Marketplace struct {
	MarketplaceCaller     // Read-only binding to the contract
	MarketplaceTransactor // Write-only binding to the contract
	MarketplaceFilterer   // Log filterer for contract events
}

// MarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketplaceSession struct {
	Contract     *Marketplace      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketplaceCallerSession struct {
	Contract *MarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketplaceTransactorSession struct {
	Contract     *MarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketplaceRaw struct {
	Contract *Marketplace // Generic contract binding to access the raw methods on
}

// MarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketplaceCallerRaw struct {
	Contract *MarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// MarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketplaceTransactorRaw struct {
	Contract *MarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketplace creates a new instance of Marketplace, bound to a specific deployed contract.
func NewMarketplace(address common.Address, backend bind.ContractBackend) (*Marketplace, error) {
	contract, err := bindMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Marketplace{MarketplaceCaller: MarketplaceCaller{contract: contract}, MarketplaceTransactor: MarketplaceTransactor{contract: contract}, MarketplaceFilterer: MarketplaceFilterer{contract: contract}}, nil
}

// NewMarketplaceCaller creates a new read-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*MarketplaceCaller, error) {
	contract, err := bindMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCaller{contract: contract}, nil
}

// NewMarketplaceTransactor creates a new write-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketplaceTransactor, error) {
	contract, err := bindMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTransactor{contract: contract}, nil
}

// NewMarketplaceFilterer creates a new log filterer instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketplaceFilterer, error) {
	contract, err := bindMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketplaceFilterer{contract: contract}, nil
}

// bindMarketplace binds a generic wrapper to an already deployed contract.
func bindMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MarketplaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.MarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transact(opts, method, params...)
}

// FinishAuction is a paid mutator transaction binding the contract method 0x0f96837b.
//
// Solidity: function finishAuction((address,address,uint256,uint256) auctionData, bytes bidderSig, bytes ownerApprovedSig) returns()
func (_Marketplace *MarketplaceTransactor) FinishAuction(opts *bind.TransactOpts, auctionData MarketplaceAuctionData, bidderSig []byte, ownerApprovedSig []byte) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "finishAuction", auctionData, bidderSig, ownerApprovedSig)
}

// FinishAuction is a paid mutator transaction binding the contract method 0x0f96837b.
//
// Solidity: function finishAuction((address,address,uint256,uint256) auctionData, bytes bidderSig, bytes ownerApprovedSig) returns()
func (_Marketplace *MarketplaceSession) FinishAuction(auctionData MarketplaceAuctionData, bidderSig []byte, ownerApprovedSig []byte) (*types.Transaction, error) {
	return _Marketplace.Contract.FinishAuction(&_Marketplace.TransactOpts, auctionData, bidderSig, ownerApprovedSig)
}

// FinishAuction is a paid mutator transaction binding the contract method 0x0f96837b.
//
// Solidity: function finishAuction((address,address,uint256,uint256) auctionData, bytes bidderSig, bytes ownerApprovedSig) returns()
func (_Marketplace *MarketplaceTransactorSession) FinishAuction(auctionData MarketplaceAuctionData, bidderSig []byte, ownerApprovedSig []byte) (*types.Transaction, error) {
	return _Marketplace.Contract.FinishAuction(&_Marketplace.TransactOpts, auctionData, bidderSig, ownerApprovedSig)
}
