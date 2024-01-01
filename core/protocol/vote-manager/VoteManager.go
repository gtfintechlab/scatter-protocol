// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package votemanager

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

// VotemanagerMetaData contains all meta data concerning the Votemanager contract.
var VotemanagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"ModelAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"ModelRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"validationThreshold\",\"type\":\"uint256\"}],\"name\":\"createValidationProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"getModelValidationStatus\",\"outputs\":[{\"internalType\":\"enumValidationStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"refrainFromValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setEvaluationJobTokenContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setScatterProtocolContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"submitScoreVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506100193361001e565b61006d565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6111d28061007a5f395ff3fe608060405234801561000f575f80fd5b5060043610610090575f3560e01c80637c621163116100635780637c621163146100ed5780638da5cb5b14610100578063b12a4c821461011a578063ce83f60f1461012d578063f2fde38b14610140575f80fd5b80631aed680e14610094578063297a43ed146100bd578063715018a6146100d257806371f65bd7146100da575b5f80fd5b6100a76100a2366004610cda565b610153565b6040516100b49190610d3a565b60405180910390f35b6100d06100cb366004610d60565b61020f565b005b6100d0610861565b6100d06100e8366004610dd1565b610874565b6100d06100fb366004610cda565b61089e565b5f546040516001600160a01b0390911681526020016100b4565b6100d0610128366004610dd1565b610a0e565b6100d061013b366004610dec565b610a38565b6100d061014e366004610dd1565b610af4565b6001600160a01b0383165f908152600860205260408082209051610178908590610e63565b90815260408051602092819003830190206001600160a01b0385165f908152925290205460ff166101ab57506002610208565b6001600160a01b0384165f908152600760205260409081902090516101d1908590610e63565b90815260408051602092819003830190206001600160a01b0385165f908152925290205460ff161561020557506001610208565b505f5b9392505050565b6003546001600160a01b031633146102945760405162461bcd60e51b815260206004820152603f60248201527f54686973206d6574686f642063616e206f6e6c792062652063616c6c6564206260448201527f7920746865206576616c756174696f6e20746f6b656e20636f6e74726163740060648201526084015b60405180910390fd5b6001600160a01b0384165f908152600560205260409081902090516102ba908590610e63565b90815260408051602092819003830190206001600160a01b038581165f908152918452828220908516825290925290205460ff16156103715760405162461bcd60e51b815260206004820152604760248201527f56616c696461746f722063616e6e6f7420766f746520747769636520666f722060448201527f612073696e676c65206d6f64656c206f6e20612076616c69646174696f6e20706064820152661c9bdc1bdcd85b60ca1b608482015260a40161028b565b6001600160a01b0384165f908152600560205260409081902090516001919061039b908690610e63565b90815260408051602092819003830181206001600160a01b038781165f9081529185528382208280528552838220805460ff191696151596909617909555938816845260059092529091206001916103f4908690610e63565b90815260408051602092819003830181206001600160a01b038781165f90815291855283822087821683528552838220805460ff1916961515969096179095559388168452600490925282209061044c908690610e63565b9081526040805191829003602090810183206060840190925281546001600160a01b031683526001820180549184019161048590610e7e565b80601f01602080910402602001604051908101604052809291908181526020018280546104b190610e7e565b80156104fc5780601f106104d3576101008083540402835291602001916104fc565b820191905f5260205f20905b8154815290600101906020018083116104df57829003601f168201915b505050505081526020016002820154815250509050600160065f876001600160a01b03166001600160a01b031681526020019081526020015f20856040516105449190610e63565b90815260200160405180910390205f846001600160a01b03166001600160a01b031681526020019081526020015f205f8282546105819190610eb6565b90915550506001600160a01b0385165f9081526006602052604080822090516105ab908790610e63565b9081526040805191829003602090810183206001600160a01b038088165f9081529190925291822054600154637e720f6d60e01b85529094509192911690637e720f6d906105ff908a908a90600401610f06565b5f60405180830381865afa158015610619573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526106409190810190610f31565b5190508082036108585760035460405163d7287ff760e01b81525f916001600160a01b03169063d7287ff79061067e908b908b908a90600401610fde565b602060405180830381865afa158015610699573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106bd9190611012565b6001600160a01b0389165f908152600860205260409081902090519192506001916106e9908a90610e63565b90815260408051602092819003830190206001600160a01b0389165f908152925290819020805460ff19169215159290921790915584015181106107c3577fbc327debbdb9caec6930e265a65fbce81ad18c570f59750876735d5f9414201488888760405161075a93929190610fde565b60405180910390a16001600160a01b0388165f908152600760205260409081902090516001919061078c908a90610e63565b90815260408051602092819003830190206001600160a01b0389165f90815292529020805460ff1916911515919091179055610856565b7f365b665108d2a8a56d80311e8038ed9d127a08de675009f04c38ba78417403a38888876040516107f693929190610fde565b60405180910390a16001600160a01b0388165f908152600760205260408082209051610823908a90610e63565b90815260408051602092819003830190206001600160a01b0389165f90815292529020805460ff19169115159190911790555b505b50505050505050565b610869610b6d565b6108725f610bc6565b565b61087c610b6d565b600380546001600160a01b0319166001600160a01b0392909216919091179055565b6001546001600160a01b031633146108c85760405162461bcd60e51b815260040161028b90611029565b6001600160a01b0383165f908152600560205260409081902090516108ee908490610e63565b90815260408051602092819003830190206001600160a01b0384165f90815290835281812081805290925290205460ff16156109a95760405162461bcd60e51b815260206004820152604e60248201527f56616c696461746f722063616e6e6f74206d61726b2061206a6f62206173206d60448201527f616c6963696f757320616674657220686176696e67207375626d69747465642060648201526d18481cd8dbdc9948199bdc881a5d60921b608482015260a40161028b565b6001600160a01b0383165f90815260096020526040908190209051600191906109d3908590610e63565b90815260408051602092819003830190206001600160a01b03949094165f90815293909152909120805460ff19169115159190911790555050565b610a16610b6d565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6001546001600160a01b03163314610a625760405162461bcd60e51b815260040161028b90611029565b604080516060810182526001600160a01b03851680825260208083018690528284018590525f918252600490528290209151909190610aa2908590610e63565b90815260405160209181900382019020825181546001600160a01b0319166001600160a01b03909116178155908201516001820190610ae190826110dc565b5060408201518160020155905050505050565b610afc610b6d565b6001600160a01b038116610b615760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161028b565b610b6a81610bc6565b50565b5f546001600160a01b031633146108725760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161028b565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b0381168114610b6a575f80fd5b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610c6657610c66610c29565b604052919050565b5f82601f830112610c7d575f80fd5b813567ffffffffffffffff811115610c9757610c97610c29565b610caa601f8201601f1916602001610c3d565b818152846020838601011115610cbe575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f60608486031215610cec575f80fd5b8335610cf781610c15565b9250602084013567ffffffffffffffff811115610d12575f80fd5b610d1e86828701610c6e565b9250506040840135610d2f81610c15565b809150509250925092565b6020810160038310610d5a57634e487b7160e01b5f52602160045260245ffd5b91905290565b5f805f8060808587031215610d73575f80fd5b8435610d7e81610c15565b9350602085013567ffffffffffffffff811115610d99575f80fd5b610da587828801610c6e565b9350506040850135610db681610c15565b91506060850135610dc681610c15565b939692955090935050565b5f60208284031215610de1575f80fd5b813561020881610c15565b5f805f60608486031215610dfe575f80fd5b8335610e0981610c15565b9250602084013567ffffffffffffffff811115610e24575f80fd5b610e3086828701610c6e565b925050604084013590509250925092565b5f5b83811015610e5b578181015183820152602001610e43565b50505f910152565b5f8251610e74818460208701610e41565b9190910192915050565b600181811c90821680610e9257607f821691505b602082108103610eb057634e487b7160e01b5f52602260045260245ffd5b50919050565b80820180821115610ed557634e487b7160e01b5f52601160045260245ffd5b92915050565b5f8151808452610ef2816020860160208601610e41565b601f01601f19169290920160200192915050565b6001600160a01b03831681526040602082018190525f90610f2990830184610edb565b949350505050565b5f6020808385031215610f42575f80fd5b825167ffffffffffffffff80821115610f59575f80fd5b818501915085601f830112610f6c575f80fd5b815181811115610f7e57610f7e610c29565b8060051b9150610f8f848301610c3d565b8181529183018401918481019088841115610fa8575f80fd5b938501935b83851015610fd25784519250610fc283610c15565b8282529385019390850190610fad565b98975050505050505050565b5f60018060a01b03808616835260606020840152610fff6060840186610edb565b9150808416604084015250949350505050565b5f60208284031215611022575f80fd5b5051919050565b60208082526041908201527f54686973206d6574686f642063616e206f6e6c792062652063616c6c6564206260408201527f792074686520747261696e696e67206a6f6220746f6b656e20636f6e747261636060820152601d60fa1b608082015260a00190565b601f8211156110d757805f5260205f20601f840160051c810160208510156110b55750805b601f840160051c820191505b818110156110d4575f81556001016110c1565b50505b505050565b815167ffffffffffffffff8111156110f6576110f6610c29565b61110a816111048454610e7e565b84611090565b602080601f83116001811461113d575f84156111265750858301515b5f19600386901b1c1916600185901b178555611194565b5f85815260208120601f198616915b8281101561116b5788860151825594840194600190910190840161114c565b508582101561118857878501515f19600388901b60f8161c191681555b505060018460011b0185555b50505050505056fea26469706673582212208ba0f9f27a265c8130ba460c1bfb30928c0f11da455851414f85052c4cb647d164736f6c63430008170033",
}

// VotemanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use VotemanagerMetaData.ABI instead.
var VotemanagerABI = VotemanagerMetaData.ABI

// VotemanagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VotemanagerMetaData.Bin instead.
var VotemanagerBin = VotemanagerMetaData.Bin

// DeployVotemanager deploys a new Ethereum contract, binding an instance of Votemanager to it.
func DeployVotemanager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Votemanager, error) {
	parsed, err := VotemanagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VotemanagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Votemanager{VotemanagerCaller: VotemanagerCaller{contract: contract}, VotemanagerTransactor: VotemanagerTransactor{contract: contract}, VotemanagerFilterer: VotemanagerFilterer{contract: contract}}, nil
}

// Votemanager is an auto generated Go binding around an Ethereum contract.
type Votemanager struct {
	VotemanagerCaller     // Read-only binding to the contract
	VotemanagerTransactor // Write-only binding to the contract
	VotemanagerFilterer   // Log filterer for contract events
}

// VotemanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotemanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotemanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotemanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotemanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotemanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotemanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotemanagerSession struct {
	Contract     *Votemanager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotemanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotemanagerCallerSession struct {
	Contract *VotemanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// VotemanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotemanagerTransactorSession struct {
	Contract     *VotemanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// VotemanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotemanagerRaw struct {
	Contract *Votemanager // Generic contract binding to access the raw methods on
}

// VotemanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotemanagerCallerRaw struct {
	Contract *VotemanagerCaller // Generic read-only contract binding to access the raw methods on
}

// VotemanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotemanagerTransactorRaw struct {
	Contract *VotemanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVotemanager creates a new instance of Votemanager, bound to a specific deployed contract.
func NewVotemanager(address common.Address, backend bind.ContractBackend) (*Votemanager, error) {
	contract, err := bindVotemanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Votemanager{VotemanagerCaller: VotemanagerCaller{contract: contract}, VotemanagerTransactor: VotemanagerTransactor{contract: contract}, VotemanagerFilterer: VotemanagerFilterer{contract: contract}}, nil
}

// NewVotemanagerCaller creates a new read-only instance of Votemanager, bound to a specific deployed contract.
func NewVotemanagerCaller(address common.Address, caller bind.ContractCaller) (*VotemanagerCaller, error) {
	contract, err := bindVotemanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotemanagerCaller{contract: contract}, nil
}

// NewVotemanagerTransactor creates a new write-only instance of Votemanager, bound to a specific deployed contract.
func NewVotemanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*VotemanagerTransactor, error) {
	contract, err := bindVotemanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotemanagerTransactor{contract: contract}, nil
}

// NewVotemanagerFilterer creates a new log filterer instance of Votemanager, bound to a specific deployed contract.
func NewVotemanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*VotemanagerFilterer, error) {
	contract, err := bindVotemanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotemanagerFilterer{contract: contract}, nil
}

// bindVotemanager binds a generic wrapper to an already deployed contract.
func bindVotemanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VotemanagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Votemanager *VotemanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Votemanager.Contract.VotemanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Votemanager *VotemanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Votemanager.Contract.VotemanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Votemanager *VotemanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Votemanager.Contract.VotemanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Votemanager *VotemanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Votemanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Votemanager *VotemanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Votemanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Votemanager *VotemanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Votemanager.Contract.contract.Transact(opts, method, params...)
}

// GetModelValidationStatus is a free data retrieval call binding the contract method 0x1aed680e.
//
// Solidity: function getModelValidationStatus(address requestorAddress, string topicName, address trainerAddress) view returns(uint8)
func (_Votemanager *VotemanagerCaller) GetModelValidationStatus(opts *bind.CallOpts, requestorAddress common.Address, topicName string, trainerAddress common.Address) (uint8, error) {
	var out []interface{}
	err := _Votemanager.contract.Call(opts, &out, "getModelValidationStatus", requestorAddress, topicName, trainerAddress)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetModelValidationStatus is a free data retrieval call binding the contract method 0x1aed680e.
//
// Solidity: function getModelValidationStatus(address requestorAddress, string topicName, address trainerAddress) view returns(uint8)
func (_Votemanager *VotemanagerSession) GetModelValidationStatus(requestorAddress common.Address, topicName string, trainerAddress common.Address) (uint8, error) {
	return _Votemanager.Contract.GetModelValidationStatus(&_Votemanager.CallOpts, requestorAddress, topicName, trainerAddress)
}

// GetModelValidationStatus is a free data retrieval call binding the contract method 0x1aed680e.
//
// Solidity: function getModelValidationStatus(address requestorAddress, string topicName, address trainerAddress) view returns(uint8)
func (_Votemanager *VotemanagerCallerSession) GetModelValidationStatus(requestorAddress common.Address, topicName string, trainerAddress common.Address) (uint8, error) {
	return _Votemanager.Contract.GetModelValidationStatus(&_Votemanager.CallOpts, requestorAddress, topicName, trainerAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Votemanager *VotemanagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Votemanager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Votemanager *VotemanagerSession) Owner() (common.Address, error) {
	return _Votemanager.Contract.Owner(&_Votemanager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Votemanager *VotemanagerCallerSession) Owner() (common.Address, error) {
	return _Votemanager.Contract.Owner(&_Votemanager.CallOpts)
}

// CreateValidationProposal is a paid mutator transaction binding the contract method 0xce83f60f.
//
// Solidity: function createValidationProposal(address requestorAddress, string topicName, uint256 validationThreshold) returns()
func (_Votemanager *VotemanagerTransactor) CreateValidationProposal(opts *bind.TransactOpts, requestorAddress common.Address, topicName string, validationThreshold *big.Int) (*types.Transaction, error) {
	return _Votemanager.contract.Transact(opts, "createValidationProposal", requestorAddress, topicName, validationThreshold)
}

// CreateValidationProposal is a paid mutator transaction binding the contract method 0xce83f60f.
//
// Solidity: function createValidationProposal(address requestorAddress, string topicName, uint256 validationThreshold) returns()
func (_Votemanager *VotemanagerSession) CreateValidationProposal(requestorAddress common.Address, topicName string, validationThreshold *big.Int) (*types.Transaction, error) {
	return _Votemanager.Contract.CreateValidationProposal(&_Votemanager.TransactOpts, requestorAddress, topicName, validationThreshold)
}

// CreateValidationProposal is a paid mutator transaction binding the contract method 0xce83f60f.
//
// Solidity: function createValidationProposal(address requestorAddress, string topicName, uint256 validationThreshold) returns()
func (_Votemanager *VotemanagerTransactorSession) CreateValidationProposal(requestorAddress common.Address, topicName string, validationThreshold *big.Int) (*types.Transaction, error) {
	return _Votemanager.Contract.CreateValidationProposal(&_Votemanager.TransactOpts, requestorAddress, topicName, validationThreshold)
}

// RefrainFromValidation is a paid mutator transaction binding the contract method 0x7c621163.
//
// Solidity: function refrainFromValidation(address requestorAddress, string topicName, address validatorAddress) returns()
func (_Votemanager *VotemanagerTransactor) RefrainFromValidation(opts *bind.TransactOpts, requestorAddress common.Address, topicName string, validatorAddress common.Address) (*types.Transaction, error) {
	return _Votemanager.contract.Transact(opts, "refrainFromValidation", requestorAddress, topicName, validatorAddress)
}

// RefrainFromValidation is a paid mutator transaction binding the contract method 0x7c621163.
//
// Solidity: function refrainFromValidation(address requestorAddress, string topicName, address validatorAddress) returns()
func (_Votemanager *VotemanagerSession) RefrainFromValidation(requestorAddress common.Address, topicName string, validatorAddress common.Address) (*types.Transaction, error) {
	return _Votemanager.Contract.RefrainFromValidation(&_Votemanager.TransactOpts, requestorAddress, topicName, validatorAddress)
}

// RefrainFromValidation is a paid mutator transaction binding the contract method 0x7c621163.
//
// Solidity: function refrainFromValidation(address requestorAddress, string topicName, address validatorAddress) returns()
func (_Votemanager *VotemanagerTransactorSession) RefrainFromValidation(requestorAddress common.Address, topicName string, validatorAddress common.Address) (*types.Transaction, error) {
	return _Votemanager.Contract.RefrainFromValidation(&_Votemanager.TransactOpts, requestorAddress, topicName, validatorAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Votemanager *VotemanagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Votemanager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Votemanager *VotemanagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Votemanager.Contract.RenounceOwnership(&_Votemanager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Votemanager *VotemanagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Votemanager.Contract.RenounceOwnership(&_Votemanager.TransactOpts)
}

// SetEvaluationJobTokenContract is a paid mutator transaction binding the contract method 0x71f65bd7.
//
// Solidity: function setEvaluationJobTokenContract(address contractAddress) returns()
func (_Votemanager *VotemanagerTransactor) SetEvaluationJobTokenContract(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _Votemanager.contract.Transact(opts, "setEvaluationJobTokenContract", contractAddress)
}

// SetEvaluationJobTokenContract is a paid mutator transaction binding the contract method 0x71f65bd7.
//
// Solidity: function setEvaluationJobTokenContract(address contractAddress) returns()
func (_Votemanager *VotemanagerSession) SetEvaluationJobTokenContract(contractAddress common.Address) (*types.Transaction, error) {
	return _Votemanager.Contract.SetEvaluationJobTokenContract(&_Votemanager.TransactOpts, contractAddress)
}

// SetEvaluationJobTokenContract is a paid mutator transaction binding the contract method 0x71f65bd7.
//
// Solidity: function setEvaluationJobTokenContract(address contractAddress) returns()
func (_Votemanager *VotemanagerTransactorSession) SetEvaluationJobTokenContract(contractAddress common.Address) (*types.Transaction, error) {
	return _Votemanager.Contract.SetEvaluationJobTokenContract(&_Votemanager.TransactOpts, contractAddress)
}

// SetScatterProtocolContract is a paid mutator transaction binding the contract method 0xb12a4c82.
//
// Solidity: function setScatterProtocolContract(address contractAddress) returns()
func (_Votemanager *VotemanagerTransactor) SetScatterProtocolContract(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _Votemanager.contract.Transact(opts, "setScatterProtocolContract", contractAddress)
}

// SetScatterProtocolContract is a paid mutator transaction binding the contract method 0xb12a4c82.
//
// Solidity: function setScatterProtocolContract(address contractAddress) returns()
func (_Votemanager *VotemanagerSession) SetScatterProtocolContract(contractAddress common.Address) (*types.Transaction, error) {
	return _Votemanager.Contract.SetScatterProtocolContract(&_Votemanager.TransactOpts, contractAddress)
}

// SetScatterProtocolContract is a paid mutator transaction binding the contract method 0xb12a4c82.
//
// Solidity: function setScatterProtocolContract(address contractAddress) returns()
func (_Votemanager *VotemanagerTransactorSession) SetScatterProtocolContract(contractAddress common.Address) (*types.Transaction, error) {
	return _Votemanager.Contract.SetScatterProtocolContract(&_Votemanager.TransactOpts, contractAddress)
}

// SubmitScoreVote is a paid mutator transaction binding the contract method 0x297a43ed.
//
// Solidity: function submitScoreVote(address requestorAddress, string topicName, address validatorAddress, address trainerAddress) returns()
func (_Votemanager *VotemanagerTransactor) SubmitScoreVote(opts *bind.TransactOpts, requestorAddress common.Address, topicName string, validatorAddress common.Address, trainerAddress common.Address) (*types.Transaction, error) {
	return _Votemanager.contract.Transact(opts, "submitScoreVote", requestorAddress, topicName, validatorAddress, trainerAddress)
}

// SubmitScoreVote is a paid mutator transaction binding the contract method 0x297a43ed.
//
// Solidity: function submitScoreVote(address requestorAddress, string topicName, address validatorAddress, address trainerAddress) returns()
func (_Votemanager *VotemanagerSession) SubmitScoreVote(requestorAddress common.Address, topicName string, validatorAddress common.Address, trainerAddress common.Address) (*types.Transaction, error) {
	return _Votemanager.Contract.SubmitScoreVote(&_Votemanager.TransactOpts, requestorAddress, topicName, validatorAddress, trainerAddress)
}

// SubmitScoreVote is a paid mutator transaction binding the contract method 0x297a43ed.
//
// Solidity: function submitScoreVote(address requestorAddress, string topicName, address validatorAddress, address trainerAddress) returns()
func (_Votemanager *VotemanagerTransactorSession) SubmitScoreVote(requestorAddress common.Address, topicName string, validatorAddress common.Address, trainerAddress common.Address) (*types.Transaction, error) {
	return _Votemanager.Contract.SubmitScoreVote(&_Votemanager.TransactOpts, requestorAddress, topicName, validatorAddress, trainerAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Votemanager *VotemanagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Votemanager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Votemanager *VotemanagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Votemanager.Contract.TransferOwnership(&_Votemanager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Votemanager *VotemanagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Votemanager.Contract.TransferOwnership(&_Votemanager.TransactOpts, newOwner)
}

// VotemanagerModelAcceptedIterator is returned from FilterModelAccepted and is used to iterate over the raw logs and unpacked data for ModelAccepted events raised by the Votemanager contract.
type VotemanagerModelAcceptedIterator struct {
	Event *VotemanagerModelAccepted // Event containing the contract specifics and raw log

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
func (it *VotemanagerModelAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotemanagerModelAccepted)
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
		it.Event = new(VotemanagerModelAccepted)
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
func (it *VotemanagerModelAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotemanagerModelAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotemanagerModelAccepted represents a ModelAccepted event raised by the Votemanager contract.
type VotemanagerModelAccepted struct {
	RequestorAddress common.Address
	TopicName        string
	TrainerAddress   common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterModelAccepted is a free log retrieval operation binding the contract event 0xbc327debbdb9caec6930e265a65fbce81ad18c570f59750876735d5f94142014.
//
// Solidity: event ModelAccepted(address requestorAddress, string topicName, address trainerAddress)
func (_Votemanager *VotemanagerFilterer) FilterModelAccepted(opts *bind.FilterOpts) (*VotemanagerModelAcceptedIterator, error) {

	logs, sub, err := _Votemanager.contract.FilterLogs(opts, "ModelAccepted")
	if err != nil {
		return nil, err
	}
	return &VotemanagerModelAcceptedIterator{contract: _Votemanager.contract, event: "ModelAccepted", logs: logs, sub: sub}, nil
}

// WatchModelAccepted is a free log subscription operation binding the contract event 0xbc327debbdb9caec6930e265a65fbce81ad18c570f59750876735d5f94142014.
//
// Solidity: event ModelAccepted(address requestorAddress, string topicName, address trainerAddress)
func (_Votemanager *VotemanagerFilterer) WatchModelAccepted(opts *bind.WatchOpts, sink chan<- *VotemanagerModelAccepted) (event.Subscription, error) {

	logs, sub, err := _Votemanager.contract.WatchLogs(opts, "ModelAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotemanagerModelAccepted)
				if err := _Votemanager.contract.UnpackLog(event, "ModelAccepted", log); err != nil {
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

// ParseModelAccepted is a log parse operation binding the contract event 0xbc327debbdb9caec6930e265a65fbce81ad18c570f59750876735d5f94142014.
//
// Solidity: event ModelAccepted(address requestorAddress, string topicName, address trainerAddress)
func (_Votemanager *VotemanagerFilterer) ParseModelAccepted(log types.Log) (*VotemanagerModelAccepted, error) {
	event := new(VotemanagerModelAccepted)
	if err := _Votemanager.contract.UnpackLog(event, "ModelAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotemanagerModelRejectedIterator is returned from FilterModelRejected and is used to iterate over the raw logs and unpacked data for ModelRejected events raised by the Votemanager contract.
type VotemanagerModelRejectedIterator struct {
	Event *VotemanagerModelRejected // Event containing the contract specifics and raw log

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
func (it *VotemanagerModelRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotemanagerModelRejected)
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
		it.Event = new(VotemanagerModelRejected)
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
func (it *VotemanagerModelRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotemanagerModelRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotemanagerModelRejected represents a ModelRejected event raised by the Votemanager contract.
type VotemanagerModelRejected struct {
	RequestorAddress common.Address
	TopicName        string
	TrainerAddress   common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterModelRejected is a free log retrieval operation binding the contract event 0x365b665108d2a8a56d80311e8038ed9d127a08de675009f04c38ba78417403a3.
//
// Solidity: event ModelRejected(address requestorAddress, string topicName, address trainerAddress)
func (_Votemanager *VotemanagerFilterer) FilterModelRejected(opts *bind.FilterOpts) (*VotemanagerModelRejectedIterator, error) {

	logs, sub, err := _Votemanager.contract.FilterLogs(opts, "ModelRejected")
	if err != nil {
		return nil, err
	}
	return &VotemanagerModelRejectedIterator{contract: _Votemanager.contract, event: "ModelRejected", logs: logs, sub: sub}, nil
}

// WatchModelRejected is a free log subscription operation binding the contract event 0x365b665108d2a8a56d80311e8038ed9d127a08de675009f04c38ba78417403a3.
//
// Solidity: event ModelRejected(address requestorAddress, string topicName, address trainerAddress)
func (_Votemanager *VotemanagerFilterer) WatchModelRejected(opts *bind.WatchOpts, sink chan<- *VotemanagerModelRejected) (event.Subscription, error) {

	logs, sub, err := _Votemanager.contract.WatchLogs(opts, "ModelRejected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotemanagerModelRejected)
				if err := _Votemanager.contract.UnpackLog(event, "ModelRejected", log); err != nil {
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

// ParseModelRejected is a log parse operation binding the contract event 0x365b665108d2a8a56d80311e8038ed9d127a08de675009f04c38ba78417403a3.
//
// Solidity: event ModelRejected(address requestorAddress, string topicName, address trainerAddress)
func (_Votemanager *VotemanagerFilterer) ParseModelRejected(log types.Log) (*VotemanagerModelRejected, error) {
	event := new(VotemanagerModelRejected)
	if err := _Votemanager.contract.UnpackLog(event, "ModelRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotemanagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Votemanager contract.
type VotemanagerOwnershipTransferredIterator struct {
	Event *VotemanagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VotemanagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotemanagerOwnershipTransferred)
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
		it.Event = new(VotemanagerOwnershipTransferred)
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
func (it *VotemanagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotemanagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotemanagerOwnershipTransferred represents a OwnershipTransferred event raised by the Votemanager contract.
type VotemanagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Votemanager *VotemanagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VotemanagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Votemanager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VotemanagerOwnershipTransferredIterator{contract: _Votemanager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Votemanager *VotemanagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VotemanagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Votemanager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotemanagerOwnershipTransferred)
				if err := _Votemanager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Votemanager *VotemanagerFilterer) ParseOwnershipTransferred(log types.Log) (*VotemanagerOwnershipTransferred, error) {
	event := new(VotemanagerOwnershipTransferred)
	if err := _Votemanager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
