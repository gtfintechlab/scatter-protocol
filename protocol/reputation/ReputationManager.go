// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package reputationmanager

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

// ReputationmanagerMetaData contains all meta data concerning the Reputationmanager contract.
var ReputationmanagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"HelloWorld\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"}],\"name\":\"getBenevolentTrainers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"}],\"name\":\"getBenevolentValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"}],\"name\":\"getMalevolentTrainers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"}],\"name\":\"getMalevolentValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"evaluationJobTokenAddress\",\"type\":\"address\"}],\"name\":\"setEvaluationJobContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"modelTokenContractAddress\",\"type\":\"address\"}],\"name\":\"setModelTokenContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scatterProtocolContractAddress\",\"type\":\"address\"}],\"name\":\"setScatterProtocolContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"trainerIsRogue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506110338061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610090575f3560e01c8063ab966cba11610063578063ab966cba14610101578063b12a4c8214610114578063b85a370514610144578063c74cad0114610174578063d8e1445014610197575f80fd5b806328882e811461009457806340ca9b6e146100bd5780637fffb7bd146100ec57806385f08a60146100ee575b5f80fd5b6100a76100a2366004610c93565b6101aa565b6040516100b49190610ce0565b60405180910390f35b6100ec6100cb366004610d2c565b5f80546001600160a01b0319166001600160a01b0392909216919091179055565b005b6100a76100fc366004610c93565b6104a8565b6100a761010f366004610c93565b610662565b6100ec610122366004610d2c565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b6100ec610152366004610d2c565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b610187610182366004610d4e565b610815565b60405190151581526020016100b4565b6100a76101a5366004610c93565b6108cd565b600254604051637e720f6d60e01b81526060915f916001600160a01b0390911690637e720f6d906101e19087908790600401610dfb565b5f60405180830381865afa1580156101fb573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526102229190810190610e1e565b60025460405163289a278f60e01b81529192505f916001600160a01b039091169063289a278f906102599088908890600401610dfb565b5f60405180830381865afa158015610273573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261029a9190810190610e1e565b90505f825167ffffffffffffffff8111156102b7576102b7610bd5565b6040519080825280602002602001820160405280156102e0578160200160208202803683370190505b5090505f805b845181101561049a575f5b845181101561049157306001600160a01b031663c74cad018a8a88858151811061031d5761031d610ecb565b60200260200101516040518463ffffffff1660e01b815260040161034393929190610edf565b602060405180830381865afa15801561035e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103829190610f13565b6104895760015486516001600160a01b0390911690631e2af3cb908b908b908a90879081106103b3576103b3610ecb565b60200260200101518986815181106103cd576103cd610ecb565b60200260200101516040518563ffffffff1660e01b81526004016103f49493929190610f32565b602060405180830381865afa15801561040f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104339190610f13565b6104895785828151811061044957610449610ecb565b602002602001015184848151811061046357610463610ecb565b6001600160a01b0390921660209283029190910190910152610486600184610f6c565b92505b6001016102f1565b506001016102e6565b509093505050505b92915050565b60025460405163289a278f60e01b81526060915f916001600160a01b039091169063289a278f906104df9087908790600401610dfb565b5f60405180830381865afa1580156104f9573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526105209190810190610e1e565b90505f815167ffffffffffffffff81111561053d5761053d610bd5565b604051908082528060200260200182016040528015610566578160200160208202803683370190505b5090505f805b835181101561065757306001600160a01b031663c74cad01888887858151811061059857610598610ecb565b60200260200101516040518463ffffffff1660e01b81526004016105be93929190610edf565b602060405180830381865afa1580156105d9573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105fd9190610f13565b1561064f5783818151811061061457610614610ecb565b602002602001015183838151811061062e5761062e610ecb565b60200260200101906001600160a01b031690816001600160a01b0316815250505b60010161056c565b509095945050505050565b60025460405163289a278f60e01b81526060915f916001600160a01b039091169063289a278f906106999087908790600401610dfb565b5f60405180830381865afa1580156106b3573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526106da9190810190610e1e565b90505f815167ffffffffffffffff8111156106f7576106f7610bd5565b604051908082528060200260200182016040528015610720578160200160208202803683370190505b5090505f805b835181101561065757306001600160a01b031663c74cad01888887858151811061075257610752610ecb565b60200260200101516040518463ffffffff1660e01b815260040161077893929190610edf565b602060405180830381865afa158015610793573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107b79190610f13565b61080d578381815181106107cd576107cd610ecb565b60200260200101518383815181106107e7576107e7610ecb565b6001600160a01b039092166020928302919091019091015261080a600183610f6c565b91505b600101610726565b5f8054604051632e8a493560e11b81526108c5916001600160a01b031690635d14926a9061084b90889088908890600401610edf565b5f60405180830381865afa158015610865573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261088c9190810190610f8b565b6040805160208082019092525f905281519101207fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4701490565b949350505050565b600254604051637e720f6d60e01b81526060915f916001600160a01b0390911690637e720f6d906109049087908790600401610dfb565b5f60405180830381865afa15801561091e573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526109459190810190610e1e565b60025460405163289a278f60e01b81529192505f916001600160a01b039091169063289a278f9061097c9088908890600401610dfb565b5f60405180830381865afa158015610996573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526109bd9190810190610e1e565b90505f825167ffffffffffffffff8111156109da576109da610bd5565b604051908082528060200260200182016040528015610a03578160200160208202803683370190505b5090505f805b845181101561049a575f5b8451811015610bb557306001600160a01b031663c74cad018a8a888581518110610a4057610a40610ecb565b60200260200101516040518463ffffffff1660e01b8152600401610a6693929190610edf565b602060405180830381865afa158015610a81573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610aa59190610f13565b610bad5760015486516001600160a01b0390911690631e2af3cb908b908b908a9087908110610ad657610ad6610ecb565b6020026020010151898681518110610af057610af0610ecb565b60200260200101516040518563ffffffff1660e01b8152600401610b179493929190610f32565b602060405180830381865afa158015610b32573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b569190610f13565b15610bad57858281518110610b6d57610b6d610ecb565b6020026020010151848481518110610b8757610b87610ecb565b6001600160a01b0390921660209283029190910190910152610baa600184610f6c565b92505b600101610a14565b50600101610a09565b6001600160a01b0381168114610bd2575f80fd5b50565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610c1257610c12610bd5565b604052919050565b5f67ffffffffffffffff821115610c3357610c33610bd5565b50601f01601f191660200190565b5f82601f830112610c50575f80fd5b8135610c63610c5e82610c1a565b610be9565b818152846020838601011115610c77575f80fd5b816020850160208301375f918101602001919091529392505050565b5f8060408385031215610ca4575f80fd5b8235610caf81610bbe565b9150602083013567ffffffffffffffff811115610cca575f80fd5b610cd685828601610c41565b9150509250929050565b602080825282518282018190525f9190848201906040850190845b81811015610d205783516001600160a01b031683529284019291840191600101610cfb565b50909695505050505050565b5f60208284031215610d3c575f80fd5b8135610d4781610bbe565b9392505050565b5f805f60608486031215610d60575f80fd5b8335610d6b81610bbe565b9250602084013567ffffffffffffffff811115610d86575f80fd5b610d9286828701610c41565b9250506040840135610da381610bbe565b809150509250925092565b5f5b83811015610dc8578181015183820152602001610db0565b50505f910152565b5f8151808452610de7816020860160208601610dae565b601f01601f19169290920160200192915050565b6001600160a01b03831681526040602082018190525f906108c590830184610dd0565b5f6020808385031215610e2f575f80fd5b825167ffffffffffffffff80821115610e46575f80fd5b818501915085601f830112610e59575f80fd5b815181811115610e6b57610e6b610bd5565b8060051b9150610e7c848301610be9565b8181529183018401918481019088841115610e95575f80fd5b938501935b83851015610ebf5784519250610eaf83610bbe565b8282529385019390850190610e9a565b98975050505050505050565b634e487b7160e01b5f52603260045260245ffd5b5f60018060a01b03808616835260606020840152610f006060840186610dd0565b9150808416604084015250949350505050565b5f60208284031215610f23575f80fd5b81518015158114610d47575f80fd5b5f60018060a01b03808716835260806020840152610f536080840187610dd0565b9481166040840152929092166060909101525092915050565b808201808211156104a257634e487b7160e01b5f52601160045260245ffd5b5f60208284031215610f9b575f80fd5b815167ffffffffffffffff811115610fb1575f80fd5b8201601f81018413610fc1575f80fd5b8051610fcf610c5e82610c1a565b818152856020838501011115610fe3575f80fd5b610ff4826020830160208601610dae565b9594505050505056fea264697066735822122092f825fa298d1f66821e02bfa1589ccba22f9fe862f5704d4e24073b732ed3ac64736f6c63430008170033",
}

// ReputationmanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ReputationmanagerMetaData.ABI instead.
var ReputationmanagerABI = ReputationmanagerMetaData.ABI

// ReputationmanagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReputationmanagerMetaData.Bin instead.
var ReputationmanagerBin = ReputationmanagerMetaData.Bin

// DeployReputationmanager deploys a new Ethereum contract, binding an instance of Reputationmanager to it.
func DeployReputationmanager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Reputationmanager, error) {
	parsed, err := ReputationmanagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReputationmanagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Reputationmanager{ReputationmanagerCaller: ReputationmanagerCaller{contract: contract}, ReputationmanagerTransactor: ReputationmanagerTransactor{contract: contract}, ReputationmanagerFilterer: ReputationmanagerFilterer{contract: contract}}, nil
}

// Reputationmanager is an auto generated Go binding around an Ethereum contract.
type Reputationmanager struct {
	ReputationmanagerCaller     // Read-only binding to the contract
	ReputationmanagerTransactor // Write-only binding to the contract
	ReputationmanagerFilterer   // Log filterer for contract events
}

// ReputationmanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReputationmanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationmanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReputationmanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationmanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReputationmanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationmanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReputationmanagerSession struct {
	Contract     *Reputationmanager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ReputationmanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReputationmanagerCallerSession struct {
	Contract *ReputationmanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ReputationmanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReputationmanagerTransactorSession struct {
	Contract     *ReputationmanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ReputationmanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReputationmanagerRaw struct {
	Contract *Reputationmanager // Generic contract binding to access the raw methods on
}

// ReputationmanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReputationmanagerCallerRaw struct {
	Contract *ReputationmanagerCaller // Generic read-only contract binding to access the raw methods on
}

// ReputationmanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReputationmanagerTransactorRaw struct {
	Contract *ReputationmanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReputationmanager creates a new instance of Reputationmanager, bound to a specific deployed contract.
func NewReputationmanager(address common.Address, backend bind.ContractBackend) (*Reputationmanager, error) {
	contract, err := bindReputationmanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Reputationmanager{ReputationmanagerCaller: ReputationmanagerCaller{contract: contract}, ReputationmanagerTransactor: ReputationmanagerTransactor{contract: contract}, ReputationmanagerFilterer: ReputationmanagerFilterer{contract: contract}}, nil
}

// NewReputationmanagerCaller creates a new read-only instance of Reputationmanager, bound to a specific deployed contract.
func NewReputationmanagerCaller(address common.Address, caller bind.ContractCaller) (*ReputationmanagerCaller, error) {
	contract, err := bindReputationmanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReputationmanagerCaller{contract: contract}, nil
}

// NewReputationmanagerTransactor creates a new write-only instance of Reputationmanager, bound to a specific deployed contract.
func NewReputationmanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ReputationmanagerTransactor, error) {
	contract, err := bindReputationmanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReputationmanagerTransactor{contract: contract}, nil
}

// NewReputationmanagerFilterer creates a new log filterer instance of Reputationmanager, bound to a specific deployed contract.
func NewReputationmanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ReputationmanagerFilterer, error) {
	contract, err := bindReputationmanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReputationmanagerFilterer{contract: contract}, nil
}

// bindReputationmanager binds a generic wrapper to an already deployed contract.
func bindReputationmanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReputationmanagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reputationmanager *ReputationmanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reputationmanager.Contract.ReputationmanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reputationmanager *ReputationmanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reputationmanager.Contract.ReputationmanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reputationmanager *ReputationmanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reputationmanager.Contract.ReputationmanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reputationmanager *ReputationmanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reputationmanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reputationmanager *ReputationmanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reputationmanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reputationmanager *ReputationmanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reputationmanager.Contract.contract.Transact(opts, method, params...)
}

// HelloWorld is a free data retrieval call binding the contract method 0x7fffb7bd.
//
// Solidity: function HelloWorld() pure returns()
func (_Reputationmanager *ReputationmanagerCaller) HelloWorld(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Reputationmanager.contract.Call(opts, &out, "HelloWorld")

	if err != nil {
		return err
	}

	return err

}

// HelloWorld is a free data retrieval call binding the contract method 0x7fffb7bd.
//
// Solidity: function HelloWorld() pure returns()
func (_Reputationmanager *ReputationmanagerSession) HelloWorld() error {
	return _Reputationmanager.Contract.HelloWorld(&_Reputationmanager.CallOpts)
}

// HelloWorld is a free data retrieval call binding the contract method 0x7fffb7bd.
//
// Solidity: function HelloWorld() pure returns()
func (_Reputationmanager *ReputationmanagerCallerSession) HelloWorld() error {
	return _Reputationmanager.Contract.HelloWorld(&_Reputationmanager.CallOpts)
}

// GetBenevolentTrainers is a free data retrieval call binding the contract method 0xab966cba.
//
// Solidity: function getBenevolentTrainers(address requestorAddress, string topicName) view returns(address[])
func (_Reputationmanager *ReputationmanagerCaller) GetBenevolentTrainers(opts *bind.CallOpts, requestorAddress common.Address, topicName string) ([]common.Address, error) {
	var out []interface{}
	err := _Reputationmanager.contract.Call(opts, &out, "getBenevolentTrainers", requestorAddress, topicName)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBenevolentTrainers is a free data retrieval call binding the contract method 0xab966cba.
//
// Solidity: function getBenevolentTrainers(address requestorAddress, string topicName) view returns(address[])
func (_Reputationmanager *ReputationmanagerSession) GetBenevolentTrainers(requestorAddress common.Address, topicName string) ([]common.Address, error) {
	return _Reputationmanager.Contract.GetBenevolentTrainers(&_Reputationmanager.CallOpts, requestorAddress, topicName)
}

// GetBenevolentTrainers is a free data retrieval call binding the contract method 0xab966cba.
//
// Solidity: function getBenevolentTrainers(address requestorAddress, string topicName) view returns(address[])
func (_Reputationmanager *ReputationmanagerCallerSession) GetBenevolentTrainers(requestorAddress common.Address, topicName string) ([]common.Address, error) {
	return _Reputationmanager.Contract.GetBenevolentTrainers(&_Reputationmanager.CallOpts, requestorAddress, topicName)
}

// GetBenevolentValidators is a free data retrieval call binding the contract method 0x28882e81.
//
// Solidity: function getBenevolentValidators(address requestorAddress, string topicName) view returns(address[])
func (_Reputationmanager *ReputationmanagerCaller) GetBenevolentValidators(opts *bind.CallOpts, requestorAddress common.Address, topicName string) ([]common.Address, error) {
	var out []interface{}
	err := _Reputationmanager.contract.Call(opts, &out, "getBenevolentValidators", requestorAddress, topicName)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBenevolentValidators is a free data retrieval call binding the contract method 0x28882e81.
//
// Solidity: function getBenevolentValidators(address requestorAddress, string topicName) view returns(address[])
func (_Reputationmanager *ReputationmanagerSession) GetBenevolentValidators(requestorAddress common.Address, topicName string) ([]common.Address, error) {
	return _Reputationmanager.Contract.GetBenevolentValidators(&_Reputationmanager.CallOpts, requestorAddress, topicName)
}

// GetBenevolentValidators is a free data retrieval call binding the contract method 0x28882e81.
//
// Solidity: function getBenevolentValidators(address requestorAddress, string topicName) view returns(address[])
func (_Reputationmanager *ReputationmanagerCallerSession) GetBenevolentValidators(requestorAddress common.Address, topicName string) ([]common.Address, error) {
	return _Reputationmanager.Contract.GetBenevolentValidators(&_Reputationmanager.CallOpts, requestorAddress, topicName)
}

// GetMalevolentTrainers is a free data retrieval call binding the contract method 0x85f08a60.
//
// Solidity: function getMalevolentTrainers(address requestorAddress, string topicName) view returns(address[])
func (_Reputationmanager *ReputationmanagerCaller) GetMalevolentTrainers(opts *bind.CallOpts, requestorAddress common.Address, topicName string) ([]common.Address, error) {
	var out []interface{}
	err := _Reputationmanager.contract.Call(opts, &out, "getMalevolentTrainers", requestorAddress, topicName)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMalevolentTrainers is a free data retrieval call binding the contract method 0x85f08a60.
//
// Solidity: function getMalevolentTrainers(address requestorAddress, string topicName) view returns(address[])
func (_Reputationmanager *ReputationmanagerSession) GetMalevolentTrainers(requestorAddress common.Address, topicName string) ([]common.Address, error) {
	return _Reputationmanager.Contract.GetMalevolentTrainers(&_Reputationmanager.CallOpts, requestorAddress, topicName)
}

// GetMalevolentTrainers is a free data retrieval call binding the contract method 0x85f08a60.
//
// Solidity: function getMalevolentTrainers(address requestorAddress, string topicName) view returns(address[])
func (_Reputationmanager *ReputationmanagerCallerSession) GetMalevolentTrainers(requestorAddress common.Address, topicName string) ([]common.Address, error) {
	return _Reputationmanager.Contract.GetMalevolentTrainers(&_Reputationmanager.CallOpts, requestorAddress, topicName)
}

// GetMalevolentValidators is a free data retrieval call binding the contract method 0xd8e14450.
//
// Solidity: function getMalevolentValidators(address requestorAddress, string topicName) view returns(address[])
func (_Reputationmanager *ReputationmanagerCaller) GetMalevolentValidators(opts *bind.CallOpts, requestorAddress common.Address, topicName string) ([]common.Address, error) {
	var out []interface{}
	err := _Reputationmanager.contract.Call(opts, &out, "getMalevolentValidators", requestorAddress, topicName)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMalevolentValidators is a free data retrieval call binding the contract method 0xd8e14450.
//
// Solidity: function getMalevolentValidators(address requestorAddress, string topicName) view returns(address[])
func (_Reputationmanager *ReputationmanagerSession) GetMalevolentValidators(requestorAddress common.Address, topicName string) ([]common.Address, error) {
	return _Reputationmanager.Contract.GetMalevolentValidators(&_Reputationmanager.CallOpts, requestorAddress, topicName)
}

// GetMalevolentValidators is a free data retrieval call binding the contract method 0xd8e14450.
//
// Solidity: function getMalevolentValidators(address requestorAddress, string topicName) view returns(address[])
func (_Reputationmanager *ReputationmanagerCallerSession) GetMalevolentValidators(requestorAddress common.Address, topicName string) ([]common.Address, error) {
	return _Reputationmanager.Contract.GetMalevolentValidators(&_Reputationmanager.CallOpts, requestorAddress, topicName)
}

// TrainerIsRogue is a free data retrieval call binding the contract method 0xc74cad01.
//
// Solidity: function trainerIsRogue(address requestorAddress, string topicName, address trainerAddress) view returns(bool)
func (_Reputationmanager *ReputationmanagerCaller) TrainerIsRogue(opts *bind.CallOpts, requestorAddress common.Address, topicName string, trainerAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Reputationmanager.contract.Call(opts, &out, "trainerIsRogue", requestorAddress, topicName, trainerAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TrainerIsRogue is a free data retrieval call binding the contract method 0xc74cad01.
//
// Solidity: function trainerIsRogue(address requestorAddress, string topicName, address trainerAddress) view returns(bool)
func (_Reputationmanager *ReputationmanagerSession) TrainerIsRogue(requestorAddress common.Address, topicName string, trainerAddress common.Address) (bool, error) {
	return _Reputationmanager.Contract.TrainerIsRogue(&_Reputationmanager.CallOpts, requestorAddress, topicName, trainerAddress)
}

// TrainerIsRogue is a free data retrieval call binding the contract method 0xc74cad01.
//
// Solidity: function trainerIsRogue(address requestorAddress, string topicName, address trainerAddress) view returns(bool)
func (_Reputationmanager *ReputationmanagerCallerSession) TrainerIsRogue(requestorAddress common.Address, topicName string, trainerAddress common.Address) (bool, error) {
	return _Reputationmanager.Contract.TrainerIsRogue(&_Reputationmanager.CallOpts, requestorAddress, topicName, trainerAddress)
}

// SetEvaluationJobContract is a paid mutator transaction binding the contract method 0xb85a3705.
//
// Solidity: function setEvaluationJobContract(address evaluationJobTokenAddress) returns()
func (_Reputationmanager *ReputationmanagerTransactor) SetEvaluationJobContract(opts *bind.TransactOpts, evaluationJobTokenAddress common.Address) (*types.Transaction, error) {
	return _Reputationmanager.contract.Transact(opts, "setEvaluationJobContract", evaluationJobTokenAddress)
}

// SetEvaluationJobContract is a paid mutator transaction binding the contract method 0xb85a3705.
//
// Solidity: function setEvaluationJobContract(address evaluationJobTokenAddress) returns()
func (_Reputationmanager *ReputationmanagerSession) SetEvaluationJobContract(evaluationJobTokenAddress common.Address) (*types.Transaction, error) {
	return _Reputationmanager.Contract.SetEvaluationJobContract(&_Reputationmanager.TransactOpts, evaluationJobTokenAddress)
}

// SetEvaluationJobContract is a paid mutator transaction binding the contract method 0xb85a3705.
//
// Solidity: function setEvaluationJobContract(address evaluationJobTokenAddress) returns()
func (_Reputationmanager *ReputationmanagerTransactorSession) SetEvaluationJobContract(evaluationJobTokenAddress common.Address) (*types.Transaction, error) {
	return _Reputationmanager.Contract.SetEvaluationJobContract(&_Reputationmanager.TransactOpts, evaluationJobTokenAddress)
}

// SetModelTokenContract is a paid mutator transaction binding the contract method 0x40ca9b6e.
//
// Solidity: function setModelTokenContract(address modelTokenContractAddress) returns()
func (_Reputationmanager *ReputationmanagerTransactor) SetModelTokenContract(opts *bind.TransactOpts, modelTokenContractAddress common.Address) (*types.Transaction, error) {
	return _Reputationmanager.contract.Transact(opts, "setModelTokenContract", modelTokenContractAddress)
}

// SetModelTokenContract is a paid mutator transaction binding the contract method 0x40ca9b6e.
//
// Solidity: function setModelTokenContract(address modelTokenContractAddress) returns()
func (_Reputationmanager *ReputationmanagerSession) SetModelTokenContract(modelTokenContractAddress common.Address) (*types.Transaction, error) {
	return _Reputationmanager.Contract.SetModelTokenContract(&_Reputationmanager.TransactOpts, modelTokenContractAddress)
}

// SetModelTokenContract is a paid mutator transaction binding the contract method 0x40ca9b6e.
//
// Solidity: function setModelTokenContract(address modelTokenContractAddress) returns()
func (_Reputationmanager *ReputationmanagerTransactorSession) SetModelTokenContract(modelTokenContractAddress common.Address) (*types.Transaction, error) {
	return _Reputationmanager.Contract.SetModelTokenContract(&_Reputationmanager.TransactOpts, modelTokenContractAddress)
}

// SetScatterProtocolContract is a paid mutator transaction binding the contract method 0xb12a4c82.
//
// Solidity: function setScatterProtocolContract(address scatterProtocolContractAddress) returns()
func (_Reputationmanager *ReputationmanagerTransactor) SetScatterProtocolContract(opts *bind.TransactOpts, scatterProtocolContractAddress common.Address) (*types.Transaction, error) {
	return _Reputationmanager.contract.Transact(opts, "setScatterProtocolContract", scatterProtocolContractAddress)
}

// SetScatterProtocolContract is a paid mutator transaction binding the contract method 0xb12a4c82.
//
// Solidity: function setScatterProtocolContract(address scatterProtocolContractAddress) returns()
func (_Reputationmanager *ReputationmanagerSession) SetScatterProtocolContract(scatterProtocolContractAddress common.Address) (*types.Transaction, error) {
	return _Reputationmanager.Contract.SetScatterProtocolContract(&_Reputationmanager.TransactOpts, scatterProtocolContractAddress)
}

// SetScatterProtocolContract is a paid mutator transaction binding the contract method 0xb12a4c82.
//
// Solidity: function setScatterProtocolContract(address scatterProtocolContractAddress) returns()
func (_Reputationmanager *ReputationmanagerTransactorSession) SetScatterProtocolContract(scatterProtocolContractAddress common.Address) (*types.Transaction, error) {
	return _Reputationmanager.Contract.SetScatterProtocolContract(&_Reputationmanager.TransactOpts, scatterProtocolContractAddress)
}
