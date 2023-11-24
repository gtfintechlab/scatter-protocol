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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"}],\"name\":\"getBenevolentTrainers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"}],\"name\":\"getBenevolentValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"}],\"name\":\"getMalevolentTrainers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"}],\"name\":\"getMalevolentValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"evaluationJobTokenAddress\",\"type\":\"address\"}],\"name\":\"setEvaluationJobContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"modelTokenContractAddress\",\"type\":\"address\"}],\"name\":\"setModelTokenContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scatterProtocolContractAddress\",\"type\":\"address\"}],\"name\":\"setScatterProtocolContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"trainerIsRogue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506117f6806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063b12a4c821161005b578063b12a4c8214610139578063b85a370514610155578063c74cad0114610171578063d8e14450146101a157610088565b806328882e811461008d57806340ca9b6e146100bd57806385f08a60146100d9578063ab966cba14610109575b600080fd5b6100a760048036038101906100a2919061114e565b6101d1565b6040516100b49190611268565b60405180910390f35b6100d760048036038101906100d2919061128a565b6105ac565b005b6100f360048036038101906100ee919061114e565b6105ef565b6040516101009190611268565b60405180910390f35b610123600480360381019061011e919061114e565b610819565b6040516101309190611268565b60405180910390f35b610153600480360381019061014e919061128a565b610a51565b005b61016f600480360381019061016a919061128a565b610a95565b005b61018b600480360381019061018691906112b7565b610ad9565b6040516101989190611341565b60405180910390f35b6101bb60048036038101906101b6919061114e565b610b9f565b6040516101c89190611268565b60405180910390f35b60606000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637e720f6d85856040518363ffffffff1660e01b81526004016102329291906113ea565b600060405180830381865afa15801561024f573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061027891906114f7565b90506000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663289a278f86866040518363ffffffff1660e01b81526004016102d99291906113ea565b600060405180830381865afa1580156102f6573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061031f91906114f7565b90506000825167ffffffffffffffff81111561033e5761033d611023565b5b60405190808252806020026020018201604052801561036c5781602001602082028036833780820191505090505b5090506000805b845181101561059e5760005b845181101561058a573073ffffffffffffffffffffffffffffffffffffffff1663c74cad018a8a8885815181106103b9576103b8611540565b5b60200260200101516040518463ffffffff1660e01b81526004016103df9392919061156f565b602060405180830381865afa1580156103fc573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061042091906115d9565b61057757600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631e2af3cb8a8a89868151811061047757610476611540565b5b602002602001015189868151811061049257610491611540565b5b60200260200101516040518563ffffffff1660e01b81526004016104b99493929190611606565b602060405180830381865afa1580156104d6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104fa91906115d9565b6105765785828151811061051157610510611540565b5b602002602001015184848151811061052c5761052b611540565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050600183610573919061168b565b92505b5b8080610582906116bf565b91505061037f565b508080610596906116bf565b915050610373565b508194505050505092915050565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60606000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663289a278f85856040518363ffffffff1660e01b81526004016106509291906113ea565b600060405180830381865afa15801561066d573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061069691906114f7565b90506000815167ffffffffffffffff8111156106b5576106b4611023565b5b6040519080825280602002602001820160405280156106e35781602001602082028036833780820191505090505b5090506000805b835181101561080c573073ffffffffffffffffffffffffffffffffffffffff1663c74cad01888887858151811061072457610723611540565b5b60200260200101516040518463ffffffff1660e01b815260040161074a9392919061156f565b602060405180830381865afa158015610767573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061078b91906115d9565b156107f9578381815181106107a3576107a2611540565b5b60200260200101518383815181106107be576107bd611540565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250505b8080610804906116bf565b9150506106ea565b5081935050505092915050565b60606000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663289a278f85856040518363ffffffff1660e01b815260040161087a9291906113ea565b600060405180830381865afa158015610897573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906108c091906114f7565b90506000815167ffffffffffffffff8111156108df576108de611023565b5b60405190808252806020026020018201604052801561090d5781602001602082028036833780820191505090505b5090506000805b8351811015610a44573073ffffffffffffffffffffffffffffffffffffffff1663c74cad01888887858151811061094e5761094d611540565b5b60200260200101516040518463ffffffff1660e01b81526004016109749392919061156f565b602060405180830381865afa158015610991573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109b591906115d9565b610a31578381815181106109cc576109cb611540565b5b60200260200101518383815181106109e7576109e6611540565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050600182610a2e919061168b565b91505b8080610a3c906116bf565b915050610914565b5081935050505092915050565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000610b9660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635d14926a8686866040518463ffffffff1660e01b8152600401610b3b9392919061156f565b600060405180830381865afa158015610b58573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610b819190611777565b60405180602001604052806000815250610f7b565b90509392505050565b60606000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637e720f6d85856040518363ffffffff1660e01b8152600401610c009291906113ea565b600060405180830381865afa158015610c1d573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610c4691906114f7565b90506000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663289a278f86866040518363ffffffff1660e01b8152600401610ca79291906113ea565b600060405180830381865afa158015610cc4573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610ced91906114f7565b90506000825167ffffffffffffffff811115610d0c57610d0b611023565b5b604051908082528060200260200182016040528015610d3a5781602001602082028036833780820191505090505b5090506000805b8451811015610f6d5760005b8451811015610f59573073ffffffffffffffffffffffffffffffffffffffff1663c74cad018a8a888581518110610d8757610d86611540565b5b60200260200101516040518463ffffffff1660e01b8152600401610dad9392919061156f565b602060405180830381865afa158015610dca573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dee91906115d9565b610f4657600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631e2af3cb8a8a898681518110610e4557610e44611540565b5b6020026020010151898681518110610e6057610e5f611540565b5b60200260200101516040518563ffffffff1660e01b8152600401610e879493929190611606565b602060405180830381865afa158015610ea4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ec891906115d9565b15610f4557858281518110610ee057610edf611540565b5b6020026020010151848481518110610efb57610efa611540565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050600183610f42919061168b565b92505b5b8080610f51906116bf565b915050610d4d565b508080610f65906116bf565b915050610d41565b508194505050505092915050565b60008180519060200120838051906020012014905092915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610fd582610faa565b9050919050565b610fe581610fca565b8114610ff057600080fd5b50565b60008135905061100281610fdc565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61105b82611012565b810181811067ffffffffffffffff8211171561107a57611079611023565b5b80604052505050565b600061108d610f96565b90506110998282611052565b919050565b600067ffffffffffffffff8211156110b9576110b8611023565b5b6110c282611012565b9050602081019050919050565b82818337600083830152505050565b60006110f16110ec8461109e565b611083565b90508281526020810184848401111561110d5761110c61100d565b5b6111188482856110cf565b509392505050565b600082601f83011261113557611134611008565b5b81356111458482602086016110de565b91505092915050565b6000806040838503121561116557611164610fa0565b5b600061117385828601610ff3565b925050602083013567ffffffffffffffff81111561119457611193610fa5565b5b6111a085828601611120565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6111df81610fca565b82525050565b60006111f183836111d6565b60208301905092915050565b6000602082019050919050565b6000611215826111aa565b61121f81856111b5565b935061122a836111c6565b8060005b8381101561125b57815161124288826111e5565b975061124d836111fd565b92505060018101905061122e565b5085935050505092915050565b60006020820190508181036000830152611282818461120a565b905092915050565b6000602082840312156112a05761129f610fa0565b5b60006112ae84828501610ff3565b91505092915050565b6000806000606084860312156112d0576112cf610fa0565b5b60006112de86828701610ff3565b935050602084013567ffffffffffffffff8111156112ff576112fe610fa5565b5b61130b86828701611120565b925050604061131c86828701610ff3565b9150509250925092565b60008115159050919050565b61133b81611326565b82525050565b60006020820190506113566000830184611332565b92915050565b61136581610fca565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b838110156113a557808201518184015260208101905061138a565b60008484015250505050565b60006113bc8261136b565b6113c68185611376565b93506113d6818560208601611387565b6113df81611012565b840191505092915050565b60006040820190506113ff600083018561135c565b818103602083015261141181846113b1565b90509392505050565b600067ffffffffffffffff82111561143557611434611023565b5b602082029050602081019050919050565b600080fd5b60008151905061145a81610fdc565b92915050565b600061147361146e8461141a565b611083565b9050808382526020820190506020840283018581111561149657611495611446565b5b835b818110156114bf57806114ab888261144b565b845260208401935050602081019050611498565b5050509392505050565b600082601f8301126114de576114dd611008565b5b81516114ee848260208601611460565b91505092915050565b60006020828403121561150d5761150c610fa0565b5b600082015167ffffffffffffffff81111561152b5761152a610fa5565b5b611537848285016114c9565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000606082019050611584600083018661135c565b818103602083015261159681856113b1565b90506115a5604083018461135c565b949350505050565b6115b681611326565b81146115c157600080fd5b50565b6000815190506115d3816115ad565b92915050565b6000602082840312156115ef576115ee610fa0565b5b60006115fd848285016115c4565b91505092915050565b600060808201905061161b600083018761135c565b818103602083015261162d81866113b1565b905061163c604083018561135c565b611649606083018461135c565b95945050505050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061169682611652565b91506116a183611652565b92508282019050808211156116b9576116b861165c565b5b92915050565b60006116ca82611652565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036116fc576116fb61165c565b5b600182019050919050565b600061171a6117158461109e565b611083565b9050828152602081018484840111156117365761173561100d565b5b611741848285611387565b509392505050565b600082601f83011261175e5761175d611008565b5b815161176e848260208601611707565b91505092915050565b60006020828403121561178d5761178c610fa0565b5b600082015167ffffffffffffffff8111156117ab576117aa610fa5565b5b6117b784828501611749565b9150509291505056fea264697066735822122059518c4ed830021925517db213f3d13262aa0fee2275bfae18959c712f990ffc64736f6c63430008110033",
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
