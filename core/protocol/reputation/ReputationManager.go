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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"}],\"name\":\"getBenevolentTrainers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"}],\"name\":\"getBenevolentValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"}],\"name\":\"getMalevolentChallengers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"}],\"name\":\"getMalevolentTrainers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"}],\"name\":\"getMalevolentValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"evaluationJobTokenAddress\",\"type\":\"address\"}],\"name\":\"setEvaluationJobContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"modelTokenContractAddress\",\"type\":\"address\"}],\"name\":\"setModelTokenContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scatterProtocolContractAddress\",\"type\":\"address\"}],\"name\":\"setScatterProtocolContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voteManagerAddress\",\"type\":\"address\"}],\"name\":\"setVoteManagerContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"trainerIsRogue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"trainerModelNotSubmitted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"validatorIsRogue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506100193361001e565b61006d565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6117898061007a5f395ff3fe608060405234801561000f575f80fd5b50600436106100f0575f3560e01c8063ab966cba11610093578063c74cad0111610063578063c74cad0114610206578063d1609f7414610219578063d8e144501461022c578063f2fde38b1461023f575f80fd5b8063ab966cba146101ba578063b12a4c82146101cd578063b85a3705146101e0578063c6c0ee91146101f3575f80fd5b8063715018a6116100ce578063715018a614610162578063789b07621461016a57806385f08a601461018d5780638da5cb5b146101a0575f80fd5b806328882e81146100f457806340ca9b6e1461011d57806342d098e51461014f575b5f80fd5b610107610102366004611346565b610252565b6040516101149190611393565b60405180910390f35b61014d61012b3660046113df565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b005b61014d61015d3660046113df565b61055e565b61014d610588565b61017d610178366004611401565b61059b565b6040519015158152602001610114565b61010761019b366004611346565b610657565b5f546040516001600160a01b039091168152602001610114565b6101076101c8366004611346565b610816565b61014d6101db3660046113df565b6109c9565b61014d6101ee3660046113df565b6109f3565b610107610201366004611346565b610a1d565b61017d610214366004611401565b610bd8565b61017d610227366004611461565b610d5f565b61010761023a366004611346565b610e63565b61014d61024d3660046113df565b61114e565b600354604051637e720f6d60e01b81526060915f916001600160a01b0390911690637e720f6d90610289908790879060040161151f565b5f60405180830381865afa1580156102a3573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526102ca9190810190611542565b60035460405163289a278f60e01b81529192505f916001600160a01b039091169063289a278f90610301908890889060040161151f565b5f60405180830381865afa15801561031b573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526103429190810190611542565b90505f825167ffffffffffffffff81111561035f5761035f611288565b604051908082528060200260200182016040528015610388578160200160208202803683370190505b5090505f805b84518110156105505760015f5b85518110156104ee57306001600160a01b031663789b07628b8b8985815181106103c7576103c76115ef565b60200260200101516040518463ffffffff1660e01b81526004016103ed93929190611603565b602060405180830381865afa158015610408573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061042c9190611637565b6104e6578180156104dd5750306001600160a01b031663d1609f748b8b8a878151811061045b5761045b6115ef565b60200260200101518a8681518110610475576104756115ef565b60200260200101516040518563ffffffff1660e01b815260040161049c9493929190611656565b602060405180830381865afa1580156104b7573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104db9190611637565b155b915081156104ee575b60010161039b565b50801561054757858281518110610507576105076115ef565b6020026020010151848481518110610521576105216115ef565b6001600160a01b0390921660209283029190910190910152610544600184611690565b92505b5060010161038e565b509093505050505b92915050565b6105666111cc565b600480546001600160a01b0319166001600160a01b0392909216919091179055565b6105906111cc565b6105995f611225565b565b600154604051632e8a493560e11b81525f9161064f916001600160a01b0390911690635d14926a906105d590889088908890600401611603565b5f60405180830381865afa1580156105ef573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261061691908101906116af565b6040805160208082019092525f905281519101207fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4701490565b949350505050565b60035460405163289a278f60e01b81526060915f916001600160a01b039091169063289a278f9061068e908790879060040161151f565b5f60405180830381865afa1580156106a8573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526106cf9190810190611542565b90505f815167ffffffffffffffff8111156106ec576106ec611288565b604051908082528060200260200182016040528015610715578160200160208202803683370190505b5090505f805b835181101561080b57306001600160a01b031663c74cad018888878581518110610747576107476115ef565b60200260200101516040518463ffffffff1660e01b815260040161076d93929190611603565b602060405180830381865afa158015610788573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107ac9190611637565b15610803578381815181106107c3576107c36115ef565b60200260200101518383815181106107dd576107dd6115ef565b6001600160a01b0390921660209283029190910190910152610800600183611690565b91505b60010161071b565b509095945050505050565b60035460405163289a278f60e01b81526060915f916001600160a01b039091169063289a278f9061084d908790879060040161151f565b5f60405180830381865afa158015610867573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261088e9190810190611542565b90505f815167ffffffffffffffff8111156108ab576108ab611288565b6040519080825280602002602001820160405280156108d4578160200160208202803683370190505b5090505f805b835181101561080b57306001600160a01b031663c74cad018888878581518110610906576109066115ef565b60200260200101516040518463ffffffff1660e01b815260040161092c93929190611603565b602060405180830381865afa158015610947573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061096b9190611637565b6109c157838181518110610981576109816115ef565b602002602001015183838151811061099b5761099b6115ef565b6001600160a01b03909216602092830291909101909101526109be600183611690565b91505b6001016108da565b6109d16111cc565b600380546001600160a01b0319166001600160a01b0392909216919091179055565b6109fb6111cc565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b6004805460405163023c61f160e21b81526060925f926001600160a01b0316916308f187c491610a5191889188910161151f565b5f60405180830381865afa158015610a6b573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610a929190810190611542565b90505f815167ffffffffffffffff811115610aaf57610aaf611288565b604051908082528060200260200182016040528015610ad8578160200160208202803683370190505b5090505f805b835181101561080b5760045484516001600160a01b039091169063e13181f59089908990889086908110610b1457610b146115ef565b60200260200101516040518463ffffffff1660e01b8152600401610b3a93929190611603565b602060405180830381865afa158015610b55573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b799190611637565b15610bd057838181518110610b9057610b906115ef565b6020026020010151838381518110610baa57610baa6115ef565b6001600160a01b0390921660209283029190910190910152610bcd600183611690565b91505b600101610ade565b5f808060048054604051630d76b40760e11b81526001600160a01b0390911691631aed680e91610c0e918a918a918a9101611603565b602060405180830381865afa158015610c29573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c4d9190611735565b6002811115610c5e57610c5e611721565b604051633c4d83b160e11b8152911491505f90309063789b076290610c8b90899089908990600401611603565b602060405180830381865afa158015610ca6573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610cca9190611637565b600480546040516343faf37f60e01b81529293505f926001600160a01b03909116916343faf37f91610d02918b918b918b9101611603565b602060405180830381865afa158015610d1d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d419190611637565b90508280610d4c5750815b80610d545750805b979650505050505050565b600254604051631e2af3cb60e01b81525f9182916001600160a01b0390911690631e2af3cb90610d99908990899089908990600401611656565b602060405180830381865afa158015610db4573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610dd89190611637565b600480546040516332a22b8960e21b81529293505f926001600160a01b039091169163ca88ae2491610e10918b918b918b9101611603565b602060405180830381865afa158015610e2b573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e4f9190611637565b9050811580610d5457509695505050505050565b600354604051637e720f6d60e01b81526060915f916001600160a01b0390911690637e720f6d90610e9a908790879060040161151f565b5f60405180830381865afa158015610eb4573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610edb9190810190611542565b60035460405163289a278f60e01b81529192505f916001600160a01b039091169063289a278f90610f12908890889060040161151f565b5f60405180830381865afa158015610f2c573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610f539190810190611542565b90505f825167ffffffffffffffff811115610f7057610f70611288565b604051908082528060200260200182016040528015610f99578160200160208202803683370190505b5090505f805b8451811015610550575f5b845181101561114557306001600160a01b031663789b07628a8a888581518110610fd657610fd66115ef565b60200260200101516040518463ffffffff1660e01b8152600401610ffc93929190611603565b602060405180830381865afa158015611017573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061103b9190611637565b61113d57306001600160a01b031663d1609f748a8a898681518110611062576110626115ef565b602002602001015189868151811061107c5761107c6115ef565b60200260200101516040518563ffffffff1660e01b81526004016110a39493929190611656565b602060405180830381865afa1580156110be573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110e29190611637565b1561113d578582815181106110f9576110f96115ef565b6020026020010151848481518110611113576111136115ef565b6001600160a01b0390921660209283029190910190910152611136600184611690565b9250611145565b600101610faa565b50600101610f9f565b6111566111cc565b6001600160a01b0381166111c05760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b6111c981611225565b50565b5f546001600160a01b031633146105995760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016111b7565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03811681146111c9575f80fd5b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff811182821017156112c5576112c5611288565b604052919050565b5f67ffffffffffffffff8211156112e6576112e6611288565b50601f01601f191660200190565b5f82601f830112611303575f80fd5b8135611316611311826112cd565b61129c565b81815284602083860101111561132a575f80fd5b816020850160208301375f918101602001919091529392505050565b5f8060408385031215611357575f80fd5b823561136281611274565b9150602083013567ffffffffffffffff81111561137d575f80fd5b611389858286016112f4565b9150509250929050565b602080825282518282018190525f9190848201906040850190845b818110156113d35783516001600160a01b0316835292840192918401916001016113ae565b50909695505050505050565b5f602082840312156113ef575f80fd5b81356113fa81611274565b9392505050565b5f805f60608486031215611413575f80fd5b833561141e81611274565b9250602084013567ffffffffffffffff811115611439575f80fd5b611445868287016112f4565b925050604084013561145681611274565b809150509250925092565b5f805f8060808587031215611474575f80fd5b843561147f81611274565b9350602085013567ffffffffffffffff81111561149a575f80fd5b6114a6878288016112f4565b93505060408501356114b781611274565b915060608501356114c781611274565b939692955090935050565b5f5b838110156114ec5781810151838201526020016114d4565b50505f910152565b5f815180845261150b8160208601602086016114d2565b601f01601f19169290920160200192915050565b6001600160a01b03831681526040602082018190525f9061064f908301846114f4565b5f6020808385031215611553575f80fd5b825167ffffffffffffffff8082111561156a575f80fd5b818501915085601f83011261157d575f80fd5b81518181111561158f5761158f611288565b8060051b91506115a084830161129c565b81815291830184019184810190888411156115b9575f80fd5b938501935b838510156115e357845192506115d383611274565b82825293850193908501906115be565b98975050505050505050565b634e487b7160e01b5f52603260045260245ffd5b5f60018060a01b0380861683526060602084015261162460608401866114f4565b9150808416604084015250949350505050565b5f60208284031215611647575f80fd5b815180151581146113fa575f80fd5b5f60018060a01b0380871683526080602084015261167760808401876114f4565b9481166040840152929092166060909101525092915050565b8082018082111561055857634e487b7160e01b5f52601160045260245ffd5b5f602082840312156116bf575f80fd5b815167ffffffffffffffff8111156116d5575f80fd5b8201601f810184136116e5575f80fd5b80516116f3611311826112cd565b818152856020838501011115611707575f80fd5b6117188260208301602086016114d2565b95945050505050565b634e487b7160e01b5f52602160045260245ffd5b5f60208284031215611745575f80fd5b8151600381106113fa575f80fdfea2646970667358221220f2ba42450cabd2c3285b4f156e23fdecf0c7c8306c02145c5b0f373808af141d64736f6c63430008170033",
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

// GetMalevolentChallengers is a free data retrieval call binding the contract method 0xc6c0ee91.
//
// Solidity: function getMalevolentChallengers(address requestorAddress, string topicName) view returns(address[])
func (_Reputationmanager *ReputationmanagerCaller) GetMalevolentChallengers(opts *bind.CallOpts, requestorAddress common.Address, topicName string) ([]common.Address, error) {
	var out []interface{}
	err := _Reputationmanager.contract.Call(opts, &out, "getMalevolentChallengers", requestorAddress, topicName)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMalevolentChallengers is a free data retrieval call binding the contract method 0xc6c0ee91.
//
// Solidity: function getMalevolentChallengers(address requestorAddress, string topicName) view returns(address[])
func (_Reputationmanager *ReputationmanagerSession) GetMalevolentChallengers(requestorAddress common.Address, topicName string) ([]common.Address, error) {
	return _Reputationmanager.Contract.GetMalevolentChallengers(&_Reputationmanager.CallOpts, requestorAddress, topicName)
}

// GetMalevolentChallengers is a free data retrieval call binding the contract method 0xc6c0ee91.
//
// Solidity: function getMalevolentChallengers(address requestorAddress, string topicName) view returns(address[])
func (_Reputationmanager *ReputationmanagerCallerSession) GetMalevolentChallengers(requestorAddress common.Address, topicName string) ([]common.Address, error) {
	return _Reputationmanager.Contract.GetMalevolentChallengers(&_Reputationmanager.CallOpts, requestorAddress, topicName)
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

// ValidatorIsRogue is a free data retrieval call binding the contract method 0xd1609f74.
//
// Solidity: function validatorIsRogue(address requestorAddress, string topicName, address validatorAddress, address trainerAddress) view returns(bool)
func (_Reputationmanager *ReputationmanagerCaller) ValidatorIsRogue(opts *bind.CallOpts, requestorAddress common.Address, topicName string, validatorAddress common.Address, trainerAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Reputationmanager.contract.Call(opts, &out, "validatorIsRogue", requestorAddress, topicName, validatorAddress, trainerAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidatorIsRogue is a free data retrieval call binding the contract method 0xd1609f74.
//
// Solidity: function validatorIsRogue(address requestorAddress, string topicName, address validatorAddress, address trainerAddress) view returns(bool)
func (_Reputationmanager *ReputationmanagerSession) ValidatorIsRogue(requestorAddress common.Address, topicName string, validatorAddress common.Address, trainerAddress common.Address) (bool, error) {
	return _Reputationmanager.Contract.ValidatorIsRogue(&_Reputationmanager.CallOpts, requestorAddress, topicName, validatorAddress, trainerAddress)
}

// ValidatorIsRogue is a free data retrieval call binding the contract method 0xd1609f74.
//
// Solidity: function validatorIsRogue(address requestorAddress, string topicName, address validatorAddress, address trainerAddress) view returns(bool)
func (_Reputationmanager *ReputationmanagerCallerSession) ValidatorIsRogue(requestorAddress common.Address, topicName string, validatorAddress common.Address, trainerAddress common.Address) (bool, error) {
	return _Reputationmanager.Contract.ValidatorIsRogue(&_Reputationmanager.CallOpts, requestorAddress, topicName, validatorAddress, trainerAddress)
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
