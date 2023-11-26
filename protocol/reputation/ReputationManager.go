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
)

// ReputationmanagerMetaData contains all meta data concerning the Reputationmanager contract.
var ReputationmanagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"HelloWorld\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"}],\"name\":\"getBenevolentTrainers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"}],\"name\":\"getBenevolentValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"}],\"name\":\"getMalevolentTrainers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"}],\"name\":\"getMalevolentValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"evaluationJobTokenAddress\",\"type\":\"address\"}],\"name\":\"setEvaluationJobContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"modelTokenContractAddress\",\"type\":\"address\"}],\"name\":\"setModelTokenContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scatterProtocolContractAddress\",\"type\":\"address\"}],\"name\":\"setScatterProtocolContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"trainerIsRogue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506110f8806100206000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063ab966cba11610066578063ab966cba14610106578063b12a4c8214610119578063b85a370514610149578063c74cad0114610179578063d8e144501461019c57600080fd5b806328882e811461009857806340ca9b6e146100c15780637fffb7bd146100f157806385f08a60146100f3575b600080fd5b6100ab6100a6366004610d16565b6101af565b6040516100b89190610d66565b60405180910390f35b6100f16100cf366004610db3565b600080546001600160a01b0319166001600160a01b0392909216919091179055565b005b6100ab610101366004610d16565b6104d2565b6100ab610114366004610d16565b61069f565b6100f1610127366004610db3565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b6100f1610157366004610db3565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b61018c610187366004610dd7565b610865565b60405190151581526020016100b8565b6100ab6101aa366004610d16565b610923565b600254604051637e720f6d60e01b81526060916000916001600160a01b0390911690637e720f6d906101e79087908790600401610e8b565b600060405180830381865afa158015610204573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261022c9190810190610eaf565b60025460405163289a278f60e01b81529192506000916001600160a01b039091169063289a278f906102649088908890600401610e8b565b600060405180830381865afa158015610281573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526102a99190810190610eaf565b90506000825167ffffffffffffffff8111156102c7576102c7610c51565b6040519080825280602002602001820160405280156102f0578160200160208202803683370190505b5090506000805b84518110156104c45760005b84518110156104b157306001600160a01b031663c74cad018a8a88858151811061032f5761032f610f61565b60200260200101516040518463ffffffff1660e01b815260040161035593929190610f77565b602060405180830381865afa158015610372573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103969190610fac565b61049f5760015486516001600160a01b0390911690631e2af3cb908b908b908a90879081106103c7576103c7610f61565b60200260200101518986815181106103e1576103e1610f61565b60200260200101516040518563ffffffff1660e01b81526004016104089493929190610fce565b602060405180830381865afa158015610425573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104499190610fac565b61049f5785828151811061045f5761045f610f61565b602002602001015184848151811061047957610479610f61565b6001600160a01b039092166020928302919091019091015261049c60018461101f565b92505b806104a981611032565b915050610303565b50806104bc81611032565b9150506102f7565b509093505050505b92915050565b60025460405163289a278f60e01b81526060916000916001600160a01b039091169063289a278f9061050a9087908790600401610e8b565b600060405180830381865afa158015610527573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261054f9190810190610eaf565b90506000815167ffffffffffffffff81111561056d5761056d610c51565b604051908082528060200260200182016040528015610596578160200160208202803683370190505b5090506000805b835181101561069457306001600160a01b031663c74cad0188888785815181106105c9576105c9610f61565b60200260200101516040518463ffffffff1660e01b81526004016105ef93929190610f77565b602060405180830381865afa15801561060c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106309190610fac565b156106825783818151811061064757610647610f61565b602002602001015183838151811061066157610661610f61565b60200260200101906001600160a01b031690816001600160a01b0316815250505b8061068c81611032565b91505061059d565b509095945050505050565b60025460405163289a278f60e01b81526060916000916001600160a01b039091169063289a278f906106d79087908790600401610e8b565b600060405180830381865afa1580156106f4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261071c9190810190610eaf565b90506000815167ffffffffffffffff81111561073a5761073a610c51565b604051908082528060200260200182016040528015610763578160200160208202803683370190505b5090506000805b835181101561069457306001600160a01b031663c74cad01888887858151811061079657610796610f61565b60200260200101516040518463ffffffff1660e01b81526004016107bc93929190610f77565b602060405180830381865afa1580156107d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107fd9190610fac565b6108535783818151811061081357610813610f61565b602002602001015183838151811061082d5761082d610f61565b6001600160a01b039092166020928302919091019091015261085060018361101f565b91505b8061085d81611032565b91505061076a565b60008054604051632e8a493560e11b815261091b916001600160a01b031690635d14926a9061089c90889088908890600401610f77565b600060405180830381865afa1580156108b9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526108e1919081019061104b565b6040805160208082019092526000905281519101207fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4701490565b949350505050565b600254604051637e720f6d60e01b81526060916000916001600160a01b0390911690637e720f6d9061095b9087908790600401610e8b565b600060405180830381865afa158015610978573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526109a09190810190610eaf565b60025460405163289a278f60e01b81529192506000916001600160a01b039091169063289a278f906109d89088908890600401610e8b565b600060405180830381865afa1580156109f5573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610a1d9190810190610eaf565b90506000825167ffffffffffffffff811115610a3b57610a3b610c51565b604051908082528060200260200182016040528015610a64578160200160208202803683370190505b5090506000805b84518110156104c45760005b8451811015610c2657306001600160a01b031663c74cad018a8a888581518110610aa357610aa3610f61565b60200260200101516040518463ffffffff1660e01b8152600401610ac993929190610f77565b602060405180830381865afa158015610ae6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b0a9190610fac565b610c145760015486516001600160a01b0390911690631e2af3cb908b908b908a9087908110610b3b57610b3b610f61565b6020026020010151898681518110610b5557610b55610f61565b60200260200101516040518563ffffffff1660e01b8152600401610b7c9493929190610fce565b602060405180830381865afa158015610b99573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bbd9190610fac565b15610c1457858281518110610bd457610bd4610f61565b6020026020010151848481518110610bee57610bee610f61565b6001600160a01b0390921660209283029190910190910152610c1160018461101f565b92505b80610c1e81611032565b915050610a77565b5080610c3181611032565b915050610a6b565b6001600160a01b0381168114610c4e57600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610c9057610c90610c51565b604052919050565b600067ffffffffffffffff821115610cb257610cb2610c51565b50601f01601f191660200190565b600082601f830112610cd157600080fd5b8135610ce4610cdf82610c98565b610c67565b818152846020838601011115610cf957600080fd5b816020850160208301376000918101602001919091529392505050565b60008060408385031215610d2957600080fd5b8235610d3481610c39565b9150602083013567ffffffffffffffff811115610d5057600080fd5b610d5c85828601610cc0565b9150509250929050565b6020808252825182820181905260009190848201906040850190845b81811015610da75783516001600160a01b031683529284019291840191600101610d82565b50909695505050505050565b600060208284031215610dc557600080fd5b8135610dd081610c39565b9392505050565b600080600060608486031215610dec57600080fd5b8335610df781610c39565b9250602084013567ffffffffffffffff811115610e1357600080fd5b610e1f86828701610cc0565b9250506040840135610e3081610c39565b809150509250925092565b60005b83811015610e56578181015183820152602001610e3e565b50506000910152565b60008151808452610e77816020860160208601610e3b565b601f01601f19169290920160200192915050565b6001600160a01b038316815260406020820181905260009061091b90830184610e5f565b60006020808385031215610ec257600080fd5b825167ffffffffffffffff80821115610eda57600080fd5b818501915085601f830112610eee57600080fd5b815181811115610f0057610f00610c51565b8060051b9150610f11848301610c67565b8181529183018401918481019088841115610f2b57600080fd5b938501935b83851015610f555784519250610f4583610c39565b8282529385019390850190610f30565b98975050505050505050565b634e487b7160e01b600052603260045260246000fd5b600060018060a01b03808616835260606020840152610f996060840186610e5f565b9150808416604084015250949350505050565b600060208284031215610fbe57600080fd5b81518015158114610dd057600080fd5b600060018060a01b03808716835260806020840152610ff06080840187610e5f565b9481166040840152929092166060909101525092915050565b634e487b7160e01b600052601160045260246000fd5b808201808211156104cc576104cc611009565b60006001820161104457611044611009565b5060010190565b60006020828403121561105d57600080fd5b815167ffffffffffffffff81111561107457600080fd5b8201601f8101841361108557600080fd5b8051611093610cdf82610c98565b8181528560208385010111156110a857600080fd5b6110b9826020830160208601610e3b565b9594505050505056fea2646970667358221220e8d2767bf162a763ef0030616ab4ecbb51d4574436e12fcbd565063b6779800b64736f6c63430008110033",
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
	parsed, err := abi.JSON(strings.NewReader(ReputationmanagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
