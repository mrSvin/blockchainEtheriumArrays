// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// SimpleStorageWallet is an auto generated low-level Go binding around an user-defined struct.
type SimpleStorageWallet struct {
	WalletName string
	Balance    *big.Int
}

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nameWallet\",\"type\":\"string\"}],\"name\":\"getWallet\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"walletName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structSimpleStorage.Wallet\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nameWalletSender\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"nameWalletRecipient\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"sendMoney\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"senders\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"recipients\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"sendMoneyOneHundredTransactions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"senders\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"recipients\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"sendMoneyOneThousandTransactions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"senders\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"recipients\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"sendMoneyTenThousandTransactions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"senders\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"recipients\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"sendMoneyTenTransactions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nameWallet\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"setBalance\",\"type\":\"uint256\"}],\"name\":\"setWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"wallets\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"walletName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610dbf806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063a4e2df661161005b578063a4e2df66146100f2578063b418cf3214610112578063b6fcfcc114610125578063e2f63b521461013857600080fd5b806358ef40961461008d57806368ea2cde146100a257806372f179f8146100b55780639cc9d3b5146100df575b600080fd5b6100a061009b366004610925565b61014b565b005b6100a06100b0366004610a31565b6101f2565b6100c86100c3366004610a8d565b61031c565b6040516100d6929190610b1a565b60405180910390f35b6100a06100ed366004610a31565b6103cb565b610105610100366004610a8d565b6104f0565b6040516100d69190610b3c565b6100a0610120366004610b6e565b6105d1565b6100a0610133366004610a31565b610625565b6100a0610146366004610a31565b610749565b8060008460405161015c9190610bb3565b9081526020016040518091039020600101546101789190610be5565b6000846040516101889190610bb3565b908152602001604051809103902060010181905550806000836040516101ae9190610bb3565b9081526020016040518091039020600101546101ca9190610bfe565b6000836040516101da9190610bb3565b90815260405190819003602001902060010155505050565b60005b60638110156103165781600085838151811061021357610213610c11565b60200260200101516040516102289190610bb3565b9081526020016040518091039020600101546102449190610be5565b600085838151811061025857610258610c11565b602002602001015160405161026d9190610bb3565b90815260200160405180910390206001018190555081600084838151811061029757610297610c11565b60200260200101516040516102ac9190610bb3565b9081526020016040518091039020600101546102c89190610bfe565b60008483815181106102dc576102dc610c11565b60200260200101516040516102f19190610bb3565b908152604051908190036020019020600101558061030e81610c27565b9150506101f5565b50505050565b805160208183018101805160008252928201919093012091528054819061034290610c40565b80601f016020809104026020016040519081016040528092919081815260200182805461036e90610c40565b80156103bb5780601f10610390576101008083540402835291602001916103bb565b820191906000526020600020905b81548152906001019060200180831161039e57829003601f168201915b5050505050908060010154905082565b60005b6103e7811015610316578160008583815181106103ed576103ed610c11565b60200260200101516040516104029190610bb3565b90815260200160405180910390206001015461041e9190610be5565b600085838151811061043257610432610c11565b60200260200101516040516104479190610bb3565b90815260200160405180910390206001018190555081600084838151811061047157610471610c11565b60200260200101516040516104869190610bb3565b9081526020016040518091039020600101546104a29190610bfe565b60008483815181106104b6576104b6610c11565b60200260200101516040516104cb9190610bb3565b90815260405190819003602001902060010155806104e881610c27565b9150506103ce565b6040805180820190915260608152600060208201526000826040516105159190610bb3565b908152602001604051809103902060405180604001604052908160008201805461053e90610c40565b80601f016020809104026020016040519081016040528092919081815260200182805461056a90610c40565b80156105b75780601f1061058c576101008083540402835291602001916105b7565b820191906000526020600020905b81548152906001019060200180831161059a57829003601f168201915b505050505081526020016001820154815250509050919050565b816000836040516105e29190610bb3565b908152604051908190036020019020906105fc9082610cc9565b508060008360405161060e9190610bb3565b908152604051908190036020019020600101555050565b60005b60098110156103165781600085838151811061064657610646610c11565b602002602001015160405161065b9190610bb3565b9081526020016040518091039020600101546106779190610be5565b600085838151811061068b5761068b610c11565b60200260200101516040516106a09190610bb3565b9081526020016040518091039020600101819055508160008483815181106106ca576106ca610c11565b60200260200101516040516106df9190610bb3565b9081526020016040518091039020600101546106fb9190610bfe565b600084838151811061070f5761070f610c11565b60200260200101516040516107249190610bb3565b908152604051908190036020019020600101558061074181610c27565b915050610628565b60005b61270f8110156103165781600085838151811061076b5761076b610c11565b60200260200101516040516107809190610bb3565b90815260200160405180910390206001015461079c9190610be5565b60008583815181106107b0576107b0610c11565b60200260200101516040516107c59190610bb3565b9081526020016040518091039020600101819055508160008483815181106107ef576107ef610c11565b60200260200101516040516108049190610bb3565b9081526020016040518091039020600101546108209190610bfe565b600084838151811061083457610834610c11565b60200260200101516040516108499190610bb3565b908152604051908190036020019020600101558061086681610c27565b91505061074c565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff811182821017156108ad576108ad61086e565b604052919050565b600082601f8301126108c657600080fd5b813567ffffffffffffffff8111156108e0576108e061086e565b6108f3601f8201601f1916602001610884565b81815284602083860101111561090857600080fd5b816020850160208301376000918101602001919091529392505050565b60008060006060848603121561093a57600080fd5b833567ffffffffffffffff8082111561095257600080fd5b61095e878388016108b5565b9450602086013591508082111561097457600080fd5b50610981868287016108b5565b925050604084013590509250925092565b600082601f8301126109a357600080fd5b8135602067ffffffffffffffff808311156109c0576109c061086e565b8260051b6109cf838201610884565b93845285810183019383810190888611156109e957600080fd5b84880192505b85831015610a2557823584811115610a075760008081fd5b610a158a87838c01016108b5565b83525091840191908401906109ef565b98975050505050505050565b600080600060608486031215610a4657600080fd5b833567ffffffffffffffff80821115610a5e57600080fd5b610a6a87838801610992565b94506020860135915080821115610a8057600080fd5b5061098186828701610992565b600060208284031215610a9f57600080fd5b813567ffffffffffffffff811115610ab657600080fd5b610ac2848285016108b5565b949350505050565b60005b83811015610ae5578181015183820152602001610acd565b50506000910152565b60008151808452610b06816020860160208601610aca565b601f01601f19169290920160200192915050565b604081526000610b2d6040830185610aee565b90508260208301529392505050565b602081526000825160406020840152610b586060840182610aee565b9050602084015160408401528091505092915050565b60008060408385031215610b8157600080fd5b823567ffffffffffffffff811115610b9857600080fd5b610ba4858286016108b5565b95602094909401359450505050565b60008251610bc5818460208701610aca565b9190910192915050565b634e487b7160e01b600052601160045260246000fd5b81810381811115610bf857610bf8610bcf565b92915050565b80820180821115610bf857610bf8610bcf565b634e487b7160e01b600052603260045260246000fd5b600060018201610c3957610c39610bcf565b5060010190565b600181811c90821680610c5457607f821691505b602082108103610c7457634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115610cc457600081815260208120601f850160051c81016020861015610ca15750805b601f850160051c820191505b81811015610cc057828155600101610cad565b5050505b505050565b815167ffffffffffffffff811115610ce357610ce361086e565b610cf781610cf18454610c40565b84610c7a565b602080601f831160018114610d2c5760008415610d145750858301515b600019600386901b1c1916600185901b178555610cc0565b600085815260208120601f198616915b82811015610d5b57888601518255948401946001909101908401610d3c565b5085821015610d795787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea26469706673582212203335da3af7b0a4de3f91e5712ed16439dc981287060809386dcd5b020570673664736f6c63430008110033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// GetWallet is a free data retrieval call binding the contract method 0xa4e2df66.
//
// Solidity: function getWallet(string nameWallet) view returns((string,uint256))
func (_Api *ApiCaller) GetWallet(opts *bind.CallOpts, nameWallet string) (SimpleStorageWallet, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getWallet", nameWallet)

	if err != nil {
		return *new(SimpleStorageWallet), err
	}

	out0 := *abi.ConvertType(out[0], new(SimpleStorageWallet)).(*SimpleStorageWallet)

	return out0, err

}

// GetWallet is a free data retrieval call binding the contract method 0xa4e2df66.
//
// Solidity: function getWallet(string nameWallet) view returns((string,uint256))
func (_Api *ApiSession) GetWallet(nameWallet string) (SimpleStorageWallet, error) {
	return _Api.Contract.GetWallet(&_Api.CallOpts, nameWallet)
}

// GetWallet is a free data retrieval call binding the contract method 0xa4e2df66.
//
// Solidity: function getWallet(string nameWallet) view returns((string,uint256))
func (_Api *ApiCallerSession) GetWallet(nameWallet string) (SimpleStorageWallet, error) {
	return _Api.Contract.GetWallet(&_Api.CallOpts, nameWallet)
}

// Wallets is a free data retrieval call binding the contract method 0x72f179f8.
//
// Solidity: function wallets(string ) view returns(string walletName, uint256 balance)
func (_Api *ApiCaller) Wallets(opts *bind.CallOpts, arg0 string) (struct {
	WalletName string
	Balance    *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "wallets", arg0)

	outstruct := new(struct {
		WalletName string
		Balance    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.WalletName = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Balance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Wallets is a free data retrieval call binding the contract method 0x72f179f8.
//
// Solidity: function wallets(string ) view returns(string walletName, uint256 balance)
func (_Api *ApiSession) Wallets(arg0 string) (struct {
	WalletName string
	Balance    *big.Int
}, error) {
	return _Api.Contract.Wallets(&_Api.CallOpts, arg0)
}

// Wallets is a free data retrieval call binding the contract method 0x72f179f8.
//
// Solidity: function wallets(string ) view returns(string walletName, uint256 balance)
func (_Api *ApiCallerSession) Wallets(arg0 string) (struct {
	WalletName string
	Balance    *big.Int
}, error) {
	return _Api.Contract.Wallets(&_Api.CallOpts, arg0)
}

// SendMoney is a paid mutator transaction binding the contract method 0x58ef4096.
//
// Solidity: function sendMoney(string nameWalletSender, string nameWalletRecipient, uint256 money) returns()
func (_Api *ApiTransactor) SendMoney(opts *bind.TransactOpts, nameWalletSender string, nameWalletRecipient string, money *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "sendMoney", nameWalletSender, nameWalletRecipient, money)
}

// SendMoney is a paid mutator transaction binding the contract method 0x58ef4096.
//
// Solidity: function sendMoney(string nameWalletSender, string nameWalletRecipient, uint256 money) returns()
func (_Api *ApiSession) SendMoney(nameWalletSender string, nameWalletRecipient string, money *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SendMoney(&_Api.TransactOpts, nameWalletSender, nameWalletRecipient, money)
}

// SendMoney is a paid mutator transaction binding the contract method 0x58ef4096.
//
// Solidity: function sendMoney(string nameWalletSender, string nameWalletRecipient, uint256 money) returns()
func (_Api *ApiTransactorSession) SendMoney(nameWalletSender string, nameWalletRecipient string, money *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SendMoney(&_Api.TransactOpts, nameWalletSender, nameWalletRecipient, money)
}

// SendMoneyOneHundredTransactions is a paid mutator transaction binding the contract method 0x68ea2cde.
//
// Solidity: function sendMoneyOneHundredTransactions(string[] senders, string[] recipients, uint256 money) returns()
func (_Api *ApiTransactor) SendMoneyOneHundredTransactions(opts *bind.TransactOpts, senders []string, recipients []string, money *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "sendMoneyOneHundredTransactions", senders, recipients, money)
}

// SendMoneyOneHundredTransactions is a paid mutator transaction binding the contract method 0x68ea2cde.
//
// Solidity: function sendMoneyOneHundredTransactions(string[] senders, string[] recipients, uint256 money) returns()
func (_Api *ApiSession) SendMoneyOneHundredTransactions(senders []string, recipients []string, money *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SendMoneyOneHundredTransactions(&_Api.TransactOpts, senders, recipients, money)
}

// SendMoneyOneHundredTransactions is a paid mutator transaction binding the contract method 0x68ea2cde.
//
// Solidity: function sendMoneyOneHundredTransactions(string[] senders, string[] recipients, uint256 money) returns()
func (_Api *ApiTransactorSession) SendMoneyOneHundredTransactions(senders []string, recipients []string, money *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SendMoneyOneHundredTransactions(&_Api.TransactOpts, senders, recipients, money)
}

// SendMoneyOneThousandTransactions is a paid mutator transaction binding the contract method 0x9cc9d3b5.
//
// Solidity: function sendMoneyOneThousandTransactions(string[] senders, string[] recipients, uint256 money) returns()
func (_Api *ApiTransactor) SendMoneyOneThousandTransactions(opts *bind.TransactOpts, senders []string, recipients []string, money *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "sendMoneyOneThousandTransactions", senders, recipients, money)
}

// SendMoneyOneThousandTransactions is a paid mutator transaction binding the contract method 0x9cc9d3b5.
//
// Solidity: function sendMoneyOneThousandTransactions(string[] senders, string[] recipients, uint256 money) returns()
func (_Api *ApiSession) SendMoneyOneThousandTransactions(senders []string, recipients []string, money *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SendMoneyOneThousandTransactions(&_Api.TransactOpts, senders, recipients, money)
}

// SendMoneyOneThousandTransactions is a paid mutator transaction binding the contract method 0x9cc9d3b5.
//
// Solidity: function sendMoneyOneThousandTransactions(string[] senders, string[] recipients, uint256 money) returns()
func (_Api *ApiTransactorSession) SendMoneyOneThousandTransactions(senders []string, recipients []string, money *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SendMoneyOneThousandTransactions(&_Api.TransactOpts, senders, recipients, money)
}

// SendMoneyTenThousandTransactions is a paid mutator transaction binding the contract method 0xe2f63b52.
//
// Solidity: function sendMoneyTenThousandTransactions(string[] senders, string[] recipients, uint256 money) returns()
func (_Api *ApiTransactor) SendMoneyTenThousandTransactions(opts *bind.TransactOpts, senders []string, recipients []string, money *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "sendMoneyTenThousandTransactions", senders, recipients, money)
}

// SendMoneyTenThousandTransactions is a paid mutator transaction binding the contract method 0xe2f63b52.
//
// Solidity: function sendMoneyTenThousandTransactions(string[] senders, string[] recipients, uint256 money) returns()
func (_Api *ApiSession) SendMoneyTenThousandTransactions(senders []string, recipients []string, money *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SendMoneyTenThousandTransactions(&_Api.TransactOpts, senders, recipients, money)
}

// SendMoneyTenThousandTransactions is a paid mutator transaction binding the contract method 0xe2f63b52.
//
// Solidity: function sendMoneyTenThousandTransactions(string[] senders, string[] recipients, uint256 money) returns()
func (_Api *ApiTransactorSession) SendMoneyTenThousandTransactions(senders []string, recipients []string, money *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SendMoneyTenThousandTransactions(&_Api.TransactOpts, senders, recipients, money)
}

// SendMoneyTenTransactions is a paid mutator transaction binding the contract method 0xb6fcfcc1.
//
// Solidity: function sendMoneyTenTransactions(string[] senders, string[] recipients, uint256 money) returns()
func (_Api *ApiTransactor) SendMoneyTenTransactions(opts *bind.TransactOpts, senders []string, recipients []string, money *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "sendMoneyTenTransactions", senders, recipients, money)
}

// SendMoneyTenTransactions is a paid mutator transaction binding the contract method 0xb6fcfcc1.
//
// Solidity: function sendMoneyTenTransactions(string[] senders, string[] recipients, uint256 money) returns()
func (_Api *ApiSession) SendMoneyTenTransactions(senders []string, recipients []string, money *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SendMoneyTenTransactions(&_Api.TransactOpts, senders, recipients, money)
}

// SendMoneyTenTransactions is a paid mutator transaction binding the contract method 0xb6fcfcc1.
//
// Solidity: function sendMoneyTenTransactions(string[] senders, string[] recipients, uint256 money) returns()
func (_Api *ApiTransactorSession) SendMoneyTenTransactions(senders []string, recipients []string, money *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SendMoneyTenTransactions(&_Api.TransactOpts, senders, recipients, money)
}

// SetWallet is a paid mutator transaction binding the contract method 0xb418cf32.
//
// Solidity: function setWallet(string nameWallet, uint256 setBalance) returns()
func (_Api *ApiTransactor) SetWallet(opts *bind.TransactOpts, nameWallet string, setBalance *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setWallet", nameWallet, setBalance)
}

// SetWallet is a paid mutator transaction binding the contract method 0xb418cf32.
//
// Solidity: function setWallet(string nameWallet, uint256 setBalance) returns()
func (_Api *ApiSession) SetWallet(nameWallet string, setBalance *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetWallet(&_Api.TransactOpts, nameWallet, setBalance)
}

// SetWallet is a paid mutator transaction binding the contract method 0xb418cf32.
//
// Solidity: function setWallet(string nameWallet, uint256 setBalance) returns()
func (_Api *ApiTransactorSession) SetWallet(nameWallet string, setBalance *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetWallet(&_Api.TransactOpts, nameWallet, setBalance)
}
