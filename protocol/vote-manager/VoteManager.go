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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"ModelAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"ModelRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"validationThreshold\",\"type\":\"uint256\"}],\"name\":\"createValidationProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setEvaluationJobTokenContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setScatterProtocolContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"submitScoreVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526001600455348015610014575f80fd5b5061001e33610023565b610072565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b610a998061007f5f395ff3fe608060405234801561000f575f80fd5b506004361061007a575f3560e01c80638da5cb5b116100585780638da5cb5b146100ae578063b12a4c82146100cc578063ce83f60f146100df578063f2fde38b146100f2575f80fd5b8063297a43ed1461007e578063715018a61461009357806371f65bd71461009b575b5f80fd5b61009161008c36600461088b565b610105565b005b6100916105bc565b6100916100a93660046108f6565b6105cf565b5f54604080516001600160a01b039092168252519081900360200190f35b6100916100da3660046108f6565b6105f9565b6100916100ed366004610916565b610623565b6100916101003660046108f6565b6106b2565b6001600160a01b0384165f9081526006602052604090819020905161012b90859061098b565b90815260408051602092819003830190206001600160a01b038581165f908152918452828220908516825290925290205460ff16156101e75760405162461bcd60e51b815260206004820152604760248201527f56616c696461746f722063616e6e6f7420766f746520747769636520666f722060448201527f612073696e676c65206d6f64656c206f6e20612076616c69646174696f6e20706064820152661c9bdc1bdcd85b60ca1b608482015260a4015b60405180910390fd5b6001600160a01b0384165f908152600660205260409081902090516001919061021190869061098b565b90815260408051602092819003830181206001600160a01b038781165f90815291855283822087821683528552838220805460ff1916961515969096179095559388168452600590925282209061026990869061098b565b9081526040805191829003602090810183206060840190925281546001600160a01b03168352600182018054918401916102a2906109a6565b80601f01602080910402602001604051908101604052809291908181526020018280546102ce906109a6565b80156103195780601f106102f057610100808354040283529160200191610319565b820191905f5260205f20905b8154815290600101906020018083116102fc57829003601f168201915b505050505081526020016002820154815250509050600160075f876001600160a01b03166001600160a01b031681526020019081526020015f2085604051610361919061098b565b90815260200160405180910390205f846001600160a01b03166001600160a01b031681526020019081526020015f205f82825461039e91906109de565b90915550506001600160a01b0385165f9081526007602052604080822090516103c890879061098b565b90815260200160405180910390205f846001600160a01b03166001600160a01b031681526020019081526020015f2054905060045481036105b45760035460405163d7287ff760e01b81525f916001600160a01b03169063d7287ff790610437908a908a908990600401610a03565b602060405180830381865afa158015610452573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104769190610a4c565b90508260400151811061051f577fbc327debbdb9caec6930e265a65fbce81ad18c570f59750876735d5f941420148787866040516104b693929190610a03565b60405180910390a16001600160a01b0387165f90815260086020526040908190209051600191906104e890899061098b565b90815260408051602092819003830190206001600160a01b0388165f90815292529020805460ff19169115159190911790556105b2565b7f365b665108d2a8a56d80311e8038ed9d127a08de675009f04c38ba78417403a387878660405161055293929190610a03565b60405180910390a16001600160a01b0387165f90815260086020526040808220905161057f90899061098b565b90815260408051602092819003830190206001600160a01b0388165f90815292529020805460ff19169115159190911790555b505b505050505050565b6105c461072b565b6105cd5f610784565b565b6105d761072b565b600380546001600160a01b0319166001600160a01b0392909216919091179055565b61060161072b565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6001546001600160a01b031633146106ad5760405162461bcd60e51b815260206004820152604160248201527f54686973206d6574686f642063616e206f6e6c792062652063616c6c6564206260448201527f792074686520747261696e696e67206a6f6220746f6b656e20636f6e747261636064820152601d60fa1b608482015260a4016101de565b505050565b6106ba61072b565b6001600160a01b03811661071f5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016101de565b61072881610784565b50565b5f546001600160a01b031633146105cd5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016101de565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80356001600160a01b03811681146107e9575f80fd5b919050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112610811575f80fd5b813567ffffffffffffffff8082111561082c5761082c6107ee565b604051601f8301601f19908116603f01168101908282118183101715610854576108546107ee565b8160405283815286602085880101111561086c575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f805f806080858703121561089e575f80fd5b6108a7856107d3565b9350602085013567ffffffffffffffff8111156108c2575f80fd5b6108ce87828801610802565b9350506108dd604086016107d3565b91506108eb606086016107d3565b905092959194509250565b5f60208284031215610906575f80fd5b61090f826107d3565b9392505050565b5f805f60608486031215610928575f80fd5b610931846107d3565b9250602084013567ffffffffffffffff81111561094c575f80fd5b61095886828701610802565b925050604084013590509250925092565b5f5b8381101561098357818101518382015260200161096b565b50505f910152565b5f825161099c818460208701610969565b9190910192915050565b600181811c908216806109ba57607f821691505b6020821081036109d857634e487b7160e01b5f52602260045260245ffd5b50919050565b808201808211156109fd57634e487b7160e01b5f52601160045260245ffd5b92915050565b5f60018060a01b038086168352606060208401528451806060850152610a30816080860160208901610969565b9316604083015250601f91909101601f19160160800192915050565b5f60208284031215610a5c575f80fd5b505191905056fea26469706673582212205ff8e579394b43a9f46d0b3d4caec2bdf8ac45ae5c7f66147445f5117eda11a264736f6c63430008170033",
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
