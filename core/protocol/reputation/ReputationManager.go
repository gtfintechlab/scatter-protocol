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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"}],\"name\":\"getBenevolentTrainers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"}],\"name\":\"getBenevolentValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"}],\"name\":\"getMalevolentTrainers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"}],\"name\":\"getMalevolentValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"evaluationJobTokenAddress\",\"type\":\"address\"}],\"name\":\"setEvaluationJobContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"modelTokenContractAddress\",\"type\":\"address\"}],\"name\":\"setModelTokenContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scatterProtocolContractAddress\",\"type\":\"address\"}],\"name\":\"setScatterProtocolContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voteManagerAddress\",\"type\":\"address\"}],\"name\":\"setVoteManagerContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"trainerIsRogue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"trainerModelNotSubmitted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506100193361001e565b61006d565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b61139a8061007a5f395ff3fe608060405234801561000f575f80fd5b50600436106100cb575f3560e01c80638da5cb5b11610088578063b85a370511610063578063b85a3705146101bb578063c74cad01146101ce578063d8e14450146101e1578063f2fde38b146101f4575f80fd5b80638da5cb5b1461017b578063ab966cba14610195578063b12a4c82146101a8575f80fd5b806328882e81146100cf57806340ca9b6e146100f857806342d098e51461012a578063715018a61461013d578063789b07621461014557806385f08a6014610168575b5f80fd5b6100e26100dd366004610fc8565b610207565b6040516100ef9190611015565b60405180910390f35b610128610106366004611061565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b005b610128610138366004611061565b61051b565b610128610545565b610158610153366004611083565b610558565b60405190151581526020016100ef565b6100e2610176366004610fc8565b610614565b5f546040516001600160a01b0390911681526020016100ef565b6100e26101a3366004610fc8565b6107ce565b6101286101b6366004611061565b610981565b6101286101c9366004611061565b6109ab565b6101586101dc366004611083565b6109d5565b6100e26101ef366004610fc8565b610adc565b610128610202366004611061565b610dd0565b600354604051637e720f6d60e01b81526060915f916001600160a01b0390911690637e720f6d9061023e9087908790600401611130565b5f60405180830381865afa158015610258573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261027f9190810190611153565b60035460405163289a278f60e01b81529192505f916001600160a01b039091169063289a278f906102b69088908890600401611130565b5f60405180830381865afa1580156102d0573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526102f79190810190611153565b90505f825167ffffffffffffffff81111561031457610314610f0a565b60405190808252806020026020018201604052801561033d578160200160208202803683370190505b5090505f805b845181101561050d5760015f5b85518110156104ab57306001600160a01b031663789b07628b8b89858151811061037c5761037c611200565b60200260200101516040518463ffffffff1660e01b81526004016103a293929190611214565b602060405180830381865afa1580156103bd573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103e19190611248565b6104a35781801561049a575060025487516001600160a01b0390911690631e2af3cb908c908c908b908890811061041a5761041a611200565b60200260200101518a868151811061043457610434611200565b60200260200101516040518563ffffffff1660e01b815260040161045b9493929190611267565b602060405180830381865afa158015610476573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061049a9190611248565b915081156104ab575b600101610350565b508015610504578582815181106104c4576104c4611200565b60200260200101518484815181106104de576104de611200565b6001600160a01b03909216602092830291909101909101526105016001846112a1565b92505b50600101610343565b509093505050505b92915050565b610523610e4e565b600480546001600160a01b0319166001600160a01b0392909216919091179055565b61054d610e4e565b6105565f610ea7565b565b600154604051632e8a493560e11b81525f9161060c916001600160a01b0390911690635d14926a9061059290889088908890600401611214565b5f60405180830381865afa1580156105ac573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526105d391908101906112c0565b6040805160208082019092525f905281519101207fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4701490565b949350505050565b60035460405163289a278f60e01b81526060915f916001600160a01b039091169063289a278f9061064b9087908790600401611130565b5f60405180830381865afa158015610665573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261068c9190810190611153565b90505f815167ffffffffffffffff8111156106a9576106a9610f0a565b6040519080825280602002602001820160405280156106d2578160200160208202803683370190505b5090505f805b83518110156107c357306001600160a01b031663c74cad01888887858151811061070457610704611200565b60200260200101516040518463ffffffff1660e01b815260040161072a93929190611214565b602060405180830381865afa158015610745573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107699190611248565b156107bb5783818151811061078057610780611200565b602002602001015183838151811061079a5761079a611200565b60200260200101906001600160a01b031690816001600160a01b0316815250505b6001016106d8565b509095945050505050565b60035460405163289a278f60e01b81526060915f916001600160a01b039091169063289a278f906108059087908790600401611130565b5f60405180830381865afa15801561081f573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526108469190810190611153565b90505f815167ffffffffffffffff81111561086357610863610f0a565b60405190808252806020026020018201604052801561088c578160200160208202803683370190505b5090505f805b83518110156107c357306001600160a01b031663c74cad0188888785815181106108be576108be611200565b60200260200101516040518463ffffffff1660e01b81526004016108e493929190611214565b602060405180830381865afa1580156108ff573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109239190611248565b6109795783818151811061093957610939611200565b602002602001015183838151811061095357610953611200565b6001600160a01b03909216602092830291909101909101526109766001836112a1565b91505b600101610892565b610989610e4e565b600380546001600160a01b0319166001600160a01b0392909216919091179055565b6109b3610e4e565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b5f808060048054604051630d76b40760e11b81526001600160a01b0390911691631aed680e91610a0b918a918a918a9101611214565b602060405180830381865afa158015610a26573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a4a9190611346565b6002811115610a5b57610a5b611332565b604051633c4d83b160e11b8152911491505f90309063789b076290610a8890899089908990600401611214565b602060405180830381865afa158015610aa3573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ac79190611248565b90508180610ad25750805b9695505050505050565b600354604051637e720f6d60e01b81526060915f916001600160a01b0390911690637e720f6d90610b139087908790600401611130565b5f60405180830381865afa158015610b2d573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610b549190810190611153565b60035460405163289a278f60e01b81529192505f916001600160a01b039091169063289a278f90610b8b9088908890600401611130565b5f60405180830381865afa158015610ba5573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610bcc9190810190611153565b90505f825167ffffffffffffffff811115610be957610be9610f0a565b604051908082528060200260200182016040528015610c12578160200160208202803683370190505b5090505f805b845181101561050d575f5b8451811015610dc757306001600160a01b031663789b07628a8a888581518110610c4f57610c4f611200565b60200260200101516040518463ffffffff1660e01b8152600401610c7593929190611214565b602060405180830381865afa158015610c90573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610cb49190611248565b610dbf5760025486516001600160a01b0390911690631e2af3cb908b908b908a9087908110610ce557610ce5611200565b6020026020010151898681518110610cff57610cff611200565b60200260200101516040518563ffffffff1660e01b8152600401610d269493929190611267565b602060405180830381865afa158015610d41573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d659190611248565b610dbf57858281518110610d7b57610d7b611200565b6020026020010151848481518110610d9557610d95611200565b6001600160a01b0390921660209283029190910190910152610db86001846112a1565b9250610dc7565b600101610c23565b50600101610c18565b610dd8610e4e565b6001600160a01b038116610e425760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b610e4b81610ea7565b50565b5f546001600160a01b031633146105565760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610e39565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b0381168114610e4b575f80fd5b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610f4757610f47610f0a565b604052919050565b5f67ffffffffffffffff821115610f6857610f68610f0a565b50601f01601f191660200190565b5f82601f830112610f85575f80fd5b8135610f98610f9382610f4f565b610f1e565b818152846020838601011115610fac575f80fd5b816020850160208301375f918101602001919091529392505050565b5f8060408385031215610fd9575f80fd5b8235610fe481610ef6565b9150602083013567ffffffffffffffff811115610fff575f80fd5b61100b85828601610f76565b9150509250929050565b602080825282518282018190525f9190848201906040850190845b818110156110555783516001600160a01b031683529284019291840191600101611030565b50909695505050505050565b5f60208284031215611071575f80fd5b813561107c81610ef6565b9392505050565b5f805f60608486031215611095575f80fd5b83356110a081610ef6565b9250602084013567ffffffffffffffff8111156110bb575f80fd5b6110c786828701610f76565b92505060408401356110d881610ef6565b809150509250925092565b5f5b838110156110fd5781810151838201526020016110e5565b50505f910152565b5f815180845261111c8160208601602086016110e3565b601f01601f19169290920160200192915050565b6001600160a01b03831681526040602082018190525f9061060c90830184611105565b5f6020808385031215611164575f80fd5b825167ffffffffffffffff8082111561117b575f80fd5b818501915085601f83011261118e575f80fd5b8151818111156111a0576111a0610f0a565b8060051b91506111b1848301610f1e565b81815291830184019184810190888411156111ca575f80fd5b938501935b838510156111f457845192506111e483610ef6565b82825293850193908501906111cf565b98975050505050505050565b634e487b7160e01b5f52603260045260245ffd5b5f60018060a01b038086168352606060208401526112356060840186611105565b9150808416604084015250949350505050565b5f60208284031215611258575f80fd5b8151801515811461107c575f80fd5b5f60018060a01b038087168352608060208401526112886080840187611105565b9481166040840152929092166060909101525092915050565b8082018082111561051557634e487b7160e01b5f52601160045260245ffd5b5f602082840312156112d0575f80fd5b815167ffffffffffffffff8111156112e6575f80fd5b8201601f810184136112f6575f80fd5b8051611304610f9382610f4f565b818152856020838501011115611318575f80fd5b6113298260208301602086016110e3565b95945050505050565b634e487b7160e01b5f52602160045260245ffd5b5f60208284031215611356575f80fd5b81516003811061107c575f80fdfea26469706673582212202e08b2b4ca4e3519febb9593de99c9ac215bb35beb1bbc6c7619279e5182d30264736f6c63430008170033",
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Reputationmanager *ReputationmanagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Reputationmanager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Reputationmanager *ReputationmanagerSession) Owner() (common.Address, error) {
	return _Reputationmanager.Contract.Owner(&_Reputationmanager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Reputationmanager *ReputationmanagerCallerSession) Owner() (common.Address, error) {
	return _Reputationmanager.Contract.Owner(&_Reputationmanager.CallOpts)
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

// TrainerModelNotSubmitted is a free data retrieval call binding the contract method 0x789b0762.
//
// Solidity: function trainerModelNotSubmitted(address requestorAddress, string topicName, address trainerAddress) view returns(bool)
func (_Reputationmanager *ReputationmanagerCaller) TrainerModelNotSubmitted(opts *bind.CallOpts, requestorAddress common.Address, topicName string, trainerAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Reputationmanager.contract.Call(opts, &out, "trainerModelNotSubmitted", requestorAddress, topicName, trainerAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TrainerModelNotSubmitted is a free data retrieval call binding the contract method 0x789b0762.
//
// Solidity: function trainerModelNotSubmitted(address requestorAddress, string topicName, address trainerAddress) view returns(bool)
func (_Reputationmanager *ReputationmanagerSession) TrainerModelNotSubmitted(requestorAddress common.Address, topicName string, trainerAddress common.Address) (bool, error) {
	return _Reputationmanager.Contract.TrainerModelNotSubmitted(&_Reputationmanager.CallOpts, requestorAddress, topicName, trainerAddress)
}

// TrainerModelNotSubmitted is a free data retrieval call binding the contract method 0x789b0762.
//
// Solidity: function trainerModelNotSubmitted(address requestorAddress, string topicName, address trainerAddress) view returns(bool)
func (_Reputationmanager *ReputationmanagerCallerSession) TrainerModelNotSubmitted(requestorAddress common.Address, topicName string, trainerAddress common.Address) (bool, error) {
	return _Reputationmanager.Contract.TrainerModelNotSubmitted(&_Reputationmanager.CallOpts, requestorAddress, topicName, trainerAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Reputationmanager *ReputationmanagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reputationmanager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Reputationmanager *ReputationmanagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Reputationmanager.Contract.RenounceOwnership(&_Reputationmanager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Reputationmanager *ReputationmanagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Reputationmanager.Contract.RenounceOwnership(&_Reputationmanager.TransactOpts)
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

// SetVoteManagerContract is a paid mutator transaction binding the contract method 0x42d098e5.
//
// Solidity: function setVoteManagerContract(address voteManagerAddress) returns()
func (_Reputationmanager *ReputationmanagerTransactor) SetVoteManagerContract(opts *bind.TransactOpts, voteManagerAddress common.Address) (*types.Transaction, error) {
	return _Reputationmanager.contract.Transact(opts, "setVoteManagerContract", voteManagerAddress)
}

// SetVoteManagerContract is a paid mutator transaction binding the contract method 0x42d098e5.
//
// Solidity: function setVoteManagerContract(address voteManagerAddress) returns()
func (_Reputationmanager *ReputationmanagerSession) SetVoteManagerContract(voteManagerAddress common.Address) (*types.Transaction, error) {
	return _Reputationmanager.Contract.SetVoteManagerContract(&_Reputationmanager.TransactOpts, voteManagerAddress)
}

// SetVoteManagerContract is a paid mutator transaction binding the contract method 0x42d098e5.
//
// Solidity: function setVoteManagerContract(address voteManagerAddress) returns()
func (_Reputationmanager *ReputationmanagerTransactorSession) SetVoteManagerContract(voteManagerAddress common.Address) (*types.Transaction, error) {
	return _Reputationmanager.Contract.SetVoteManagerContract(&_Reputationmanager.TransactOpts, voteManagerAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Reputationmanager *ReputationmanagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Reputationmanager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Reputationmanager *ReputationmanagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Reputationmanager.Contract.TransferOwnership(&_Reputationmanager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Reputationmanager *ReputationmanagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Reputationmanager.Contract.TransferOwnership(&_Reputationmanager.TransactOpts, newOwner)
}

// ReputationmanagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Reputationmanager contract.
type ReputationmanagerOwnershipTransferredIterator struct {
	Event *ReputationmanagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ReputationmanagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReputationmanagerOwnershipTransferred)
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
		it.Event = new(ReputationmanagerOwnershipTransferred)
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
func (it *ReputationmanagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReputationmanagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReputationmanagerOwnershipTransferred represents a OwnershipTransferred event raised by the Reputationmanager contract.
type ReputationmanagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Reputationmanager *ReputationmanagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ReputationmanagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Reputationmanager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ReputationmanagerOwnershipTransferredIterator{contract: _Reputationmanager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Reputationmanager *ReputationmanagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ReputationmanagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Reputationmanager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReputationmanagerOwnershipTransferred)
				if err := _Reputationmanager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Reputationmanager *ReputationmanagerFilterer) ParseOwnershipTransferred(log types.Log) (*ReputationmanagerOwnershipTransferred, error) {
	event := new(ReputationmanagerOwnershipTransferred)
	if err := _Reputationmanager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
