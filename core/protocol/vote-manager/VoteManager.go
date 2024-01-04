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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"ModelAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"ModelRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"validationThreshold\",\"type\":\"uint256\"}],\"name\":\"createValidationProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"getModelValidationStatus\",\"outputs\":[{\"internalType\":\"enumValidationStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"nodeToChallenge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"challengerAddress\",\"type\":\"address\"}],\"name\":\"hasChallengedNode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"isChallengeSuccessfulTrainer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"isChallengeSuccessfulValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"}],\"name\":\"isMaliciousValidationJob\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"isRefrainingFromValidation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"refrainFromValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setEvaluationJobTokenContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setScatterProtocolContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"submitScoreVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isMalicious\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"challengerAddress\",\"type\":\"address\"}],\"name\":\"submitTrainerChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isMalicious\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"challengerAddress\",\"type\":\"address\"}],\"name\":\"submitValidatorChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506100193361001e565b61006d565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b611e988061007a5f395ff3fe608060405234801561000f575f80fd5b50600436106100fb575f3560e01c80638da5cb5b11610093578063ca88ae2411610063578063ca88ae2414610207578063ce83f60f1461021a578063ea27dd6b1461022d578063f2fde38b14610240575f80fd5b80638da5cb5b146101b457806395faa057146101ce57806397e12c36146101e1578063b12a4c82146101f4575f80fd5b8063715018a6116100ce578063715018a61461017357806371f65bd71461017b57806373138ef51461018e5780637c621163146101a1575f80fd5b80631aed680e146100ff578063297a43ed1461012857806338222c7b1461013d57806343faf37f14610160575b5f80fd5b61011261010d366004611828565b610253565b60405161011f9190611888565b60405180910390f35b61013b6101363660046118ae565b61030f565b005b61015061014b3660046118ae565b610a8d565b604051901515815260200161011f565b61015061016e366004611828565b610b50565b61013b610c12565b61013b61018936600461191f565b610c25565b61013b61019c366004611947565b610c4f565b61013b6101af366004611828565b610e27565b5f546040516001600160a01b03909116815260200161011f565b6101506101dc3660046119cc565b6111cd565b6101506101ef366004611828565b61130a565b61013b61020236600461191f565b611361565b610150610215366004611828565b61138b565b61013b610228366004611a19565b6113f8565b61013b61023b366004611947565b6114b4565b61013b61024e36600461191f565b611642565b6001600160a01b0383165f908152600860205260408082209051610278908590611a90565b90815260408051602092819003830190206001600160a01b0385165f908152925290205460ff166102ab57506002610308565b6001600160a01b0384165f908152600760205260409081902090516102d1908590611a90565b90815260408051602092819003830190206001600160a01b0385165f908152925290205460ff161561030557506001610308565b505f5b9392505050565b6003546001600160a01b031633146103945760405162461bcd60e51b815260206004820152603f60248201527f54686973206d6574686f642063616e206f6e6c792062652063616c6c6564206260448201527f7920746865206576616c756174696f6e20746f6b656e20636f6e74726163740060648201526084015b60405180910390fd5b6001600160a01b0384165f908152600560205260409081902090516103ba908590611a90565b90815260408051602092819003830190206001600160a01b038581165f908152918452828220908516825290925290205460ff16156104715760405162461bcd60e51b815260206004820152604760248201527f56616c696461746f722063616e6e6f7420766f746520747769636520666f722060448201527f612073696e676c65206d6f64656c206f6e20612076616c69646174696f6e20706064820152661c9bdc1bdcd85b60ca1b608482015260a40161038b565b6001600160a01b0384165f90815260096020526040908190209051610497908590611a90565b90815260408051602092819003830190206001600160a01b0385165f908152925290205460ff16156105315760405162461bcd60e51b815260206004820152603c60248201527f43616e6e6f74206576616c7561746520612076616c69646174696f6e2070726f60448201527f706f73616c206166746572206d61726b696e672069742066616c736500000000606482015260840161038b565b6001600160a01b0384165f908152600560205260409081902090516001919061055b908690611a90565b90815260408051602092819003830181206001600160a01b038781165f9081529185528382208280528552838220805460ff191696151596909617909555938816845260059092529091206001916105b4908690611a90565b90815260408051602092819003830181206001600160a01b038781165f90815291855283822087821683528552838220805460ff1916961515969096179095559388168452600490925282209061060c908690611a90565b9081526040805191829003602090810183206060840190925281546001600160a01b031683526001820180549184019161064590611aab565b80601f016020809104026020016040519081016040528092919081815260200182805461067190611aab565b80156106bc5780601f10610693576101008083540402835291602001916106bc565b820191905f5260205f20905b81548152906001019060200180831161069f57829003601f168201915b505050505081526020016002820154815250509050600160065f876001600160a01b03166001600160a01b031681526020019081526020015f20856040516107049190611a90565b90815260200160405180910390205f846001600160a01b03166001600160a01b031681526020019081526020015f205f8282546107419190611af7565b90915550506001600160a01b0385165f90815260066020526040808220905161076b908790611a90565b9081526040805191829003602090810183206001600160a01b038088165f9081529190925291822054600154637e720f6d60e01b85529094509192911690637e720f6d906107bf908a908a90600401611b35565b5f60405180830381865afa1580156107d9573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526108009190810190611b60565b519050808203610a845760035460405163d7287ff760e01b81525f916001600160a01b03169063d7287ff79061083e908b908b908a90600401611c0d565b602060405180830381865afa158015610859573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061087d9190611c41565b6001600160a01b0389165f908152600860205260409081902090519192506001916108a9908a90611a90565b90815260408051602092819003830190206001600160a01b0389165f908152925290819020805460ff1916921515929092179091558401518110158061094e57506040516395faa05760e01b815230906395faa0579061090f908b908b90600401611b35565b602060405180830381865afa15801561092a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061094e9190611c58565b156109ef577fbc327debbdb9caec6930e265a65fbce81ad18c570f59750876735d5f9414201488888760405161098693929190611c0d565b60405180910390a16001600160a01b0388165f90815260076020526040908190209051600191906109b8908a90611a90565b90815260408051602092819003830190206001600160a01b0389165f90815292529020805460ff1916911515919091179055610a82565b7f365b665108d2a8a56d80311e8038ed9d127a08de675009f04c38ba78417403a3888887604051610a2293929190611c0d565b60405180910390a16001600160a01b0388165f908152600760205260408082209051610a4f908a90611a90565b90815260408051602092819003830190206001600160a01b0389165f90815292529020805460ff19169115159190911790555b505b50505050505050565b6001600160a01b0384165f908152600a60205260408082209051610ab2908690611a90565b90815260408051602092819003830190206001600160a01b038681165f908152918452828220908616825290925290205460ff1680610b4757506001600160a01b0385165f908152600d6020526040908190209051610b12908690611a90565b90815260408051602092819003830190206001600160a01b038681165f908152918452828220908616825290925290205460ff165b95945050505050565b6001600160a01b0383165f908152600e60205260408082209051829190610b78908690611a90565b9081526040805191829003602090810183206001600160a01b038088165f90815291835283822054908a168252600f9092529182209093509091610bbd908790611a90565b90815260408051602092819003830190206001600160a01b0387165f908152925290205490508115801590610c0857506003610bfa826002611c73565b610c049190611c8a565b8210155b9695505050505050565b610c1a6116bb565b610c235f611714565b565b610c2d6116bb565b600380546001600160a01b0319166001600160a01b0392909216919091179055565b6001546001600160a01b03163314610c795760405162461bcd60e51b815260040161038b90611ca9565b6001600160a01b0385165f908152600d6020526040908190209051610c9f908690611a90565b90815260408051602092819003830190206001600160a01b038681165f908152918452828220908516825290925290205460ff16610e20576001600160a01b0385165f908152600d602052604090819020905160019190610d01908790611a90565b90815260408051602092819003830190206001600160a01b038781165f90815291845282822090861682529092529020805460ff19169115159190911790558115610db3576001600160a01b0385165f908152600e602052604090819020905160019190610d70908790611a90565b90815260200160405180910390205f856001600160a01b03166001600160a01b031681526020019081526020015f205f828254610dad9190611af7565b90915550505b6001600160a01b0385165f908152600f602052604090819020905160019190610ddd908790611a90565b90815260200160405180910390205f856001600160a01b03166001600160a01b031681526020019081526020015f205f828254610e1a9190611af7565b90915550505b5050505050565b6001546001600160a01b03163314610e515760405162461bcd60e51b815260040161038b90611ca9565b6001600160a01b0383165f90815260056020526040908190209051610e77908490611a90565b90815260408051602092819003830190206001600160a01b0384165f90815290835281812081805290925290205460ff1615610f325760405162461bcd60e51b815260206004820152604e60248201527f56616c696461746f722063616e6e6f74206d61726b2061206a6f62206173206d60448201527f616c6963696f757320616674657220686176696e67207375626d69747465642060648201526d18481cd8dbdc9948199bdc881a5d60921b608482015260a40161038b565b6001600160a01b0383165f90815260096020526040908190209051610f58908490611a90565b90815260408051602092819003830190206001600160a01b0384165f908152925290205460ff166111c8576001600160a01b0383165f9081526009602052604090819020905160019190610fad908590611a90565b90815260408051602092819003830181206001600160a01b038681165f90815291909452918220805460ff19169415159490941790935560015463289a278f60e01b8452909291169063289a278f9061100c9087908790600401611b35565b5f60405180830381865afa158015611026573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261104d9190810190611b60565b90505f5b8151811015610e20575f82828151811061106d5761106d611d10565b602002602001015190505f6001600160a01b0316816001600160a01b0316036110965750610e20565b60035460405163032a1cdd60e01b81526001600160a01b039091169063032a1cdd906110cc908990899089908790600401611d24565b5f604051808303815f87803b1580156110e3575f80fd5b505af11580156110f5573d5f803e3d5ffd5b5050506001600160a01b0387165f9081526005602052604090819020905160019250611122908890611a90565b90815260408051602092819003830181206001600160a01b038981165f90815291855283822087821683528552838220805460ff191696151596909617909555938a168452600690925290912060019161117d908890611a90565b90815260200160405180910390205f836001600160a01b03166001600160a01b031681526020019081526020015f205f8282546111ba9190611af7565b909155505050600101611051565b505050565b600154604051637e720f6d60e01b81525f9182916001600160a01b0390911690637e720f6d906112039087908790600401611b35565b5f60405180830381865afa15801561121d573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526112449190810190611b60565b90505f805b82518110156112e3575f83828151811061126557611265611d10565b6020026020010151905060095f886001600160a01b03166001600160a01b031681526020019081526020015f20866040516112a09190611a90565b90815260408051602092819003830190206001600160a01b0384165f908152925290205460ff16156112da576112d7600184611af7565b92505b50600101611249565b506003825160026112f49190611c73565b6112fe9190611c8a565b11159150505b92915050565b6001600160a01b0383165f90815260096020526040808220905161132f908590611a90565b908152604080519182900360209081019092206001600160a01b0385165f908152925290205460ff1690509392505050565b6113696116bb565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b0383165f908152600b602052604080822090518291906113b3908690611a90565b9081526040805191829003602090810183206001600160a01b038088165f90815291835283822054908a168252600c9092529182209093509091610bbd908790611a90565b6001546001600160a01b031633146114225760405162461bcd60e51b815260040161038b90611ca9565b604080516060810182526001600160a01b03851680825260208083018690528284018590525f918252600490528290209151909190611462908590611a90565b90815260405160209181900382019020825181546001600160a01b0319166001600160a01b039091161781559082015160018201906114a19082611da2565b5060408201518160020155905050505050565b6001546001600160a01b031633146114de5760405162461bcd60e51b815260040161038b90611ca9565b6001600160a01b0385165f908152600a6020526040908190209051611504908690611a90565b90815260408051602092819003830190206001600160a01b038681165f908152918452828220908516825290925290205460ff16610e20576001600160a01b0385165f908152600a602052604090819020905160019190611566908790611a90565b90815260408051602092819003830190206001600160a01b038781165f90815291845282822090861682529092529020805460ff19169115159190911790558115611618576001600160a01b0385165f908152600b6020526040908190209051600191906115d5908790611a90565b90815260200160405180910390205f856001600160a01b03166001600160a01b031681526020019081526020015f205f8282546116129190611af7565b90915550505b6001600160a01b0385165f908152600c602052604090819020905160019190610ddd908790611a90565b61164a6116bb565b6001600160a01b0381166116af5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161038b565b6116b881611714565b50565b5f546001600160a01b03163314610c235760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161038b565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03811681146116b8575f80fd5b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff811182821017156117b4576117b4611777565b604052919050565b5f82601f8301126117cb575f80fd5b813567ffffffffffffffff8111156117e5576117e5611777565b6117f8601f8201601f191660200161178b565b81815284602083860101111561180c575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f6060848603121561183a575f80fd5b833561184581611763565b9250602084013567ffffffffffffffff811115611860575f80fd5b61186c868287016117bc565b925050604084013561187d81611763565b809150509250925092565b60208101600383106118a857634e487b7160e01b5f52602160045260245ffd5b91905290565b5f805f80608085870312156118c1575f80fd5b84356118cc81611763565b9350602085013567ffffffffffffffff8111156118e7575f80fd5b6118f3878288016117bc565b935050604085013561190481611763565b9150606085013561191481611763565b939692955090935050565b5f6020828403121561192f575f80fd5b813561030881611763565b80151581146116b8575f80fd5b5f805f805f60a0868803121561195b575f80fd5b853561196681611763565b9450602086013567ffffffffffffffff811115611981575f80fd5b61198d888289016117bc565b945050604086013561199e81611763565b925060608601356119ae8161193a565b915060808601356119be81611763565b809150509295509295909350565b5f80604083850312156119dd575f80fd5b82356119e881611763565b9150602083013567ffffffffffffffff811115611a03575f80fd5b611a0f858286016117bc565b9150509250929050565b5f805f60608486031215611a2b575f80fd5b8335611a3681611763565b9250602084013567ffffffffffffffff811115611a51575f80fd5b611a5d868287016117bc565b925050604084013590509250925092565b5f5b83811015611a88578181015183820152602001611a70565b50505f910152565b5f8251611aa1818460208701611a6e565b9190910192915050565b600181811c90821680611abf57607f821691505b602082108103611add57634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561130457611304611ae3565b5f8151808452611b21816020860160208601611a6e565b601f01601f19169290920160200192915050565b6001600160a01b03831681526040602082018190525f90611b5890830184611b0a565b949350505050565b5f6020808385031215611b71575f80fd5b825167ffffffffffffffff80821115611b88575f80fd5b818501915085601f830112611b9b575f80fd5b815181811115611bad57611bad611777565b8060051b9150611bbe84830161178b565b8181529183018401918481019088841115611bd7575f80fd5b938501935b83851015611c015784519250611bf183611763565b8282529385019390850190611bdc565b98975050505050505050565b5f60018060a01b03808616835260606020840152611c2e6060840186611b0a565b9150808416604084015250949350505050565b5f60208284031215611c51575f80fd5b5051919050565b5f60208284031215611c68575f80fd5b81516103088161193a565b808202811582820484141761130457611304611ae3565b5f82611ca457634e487b7160e01b5f52601260045260245ffd5b500490565b60208082526041908201527f54686973206d6574686f642063616e206f6e6c792062652063616c6c6564206260408201527f792074686520747261696e696e67206a6f6220746f6b656e20636f6e747261636060820152601d60fa1b608082015260a00190565b634e487b7160e01b5f52603260045260245ffd5b5f60018060a01b03808716835260806020840152611d456080840187611b0a565b9481166040840152929092166060909101525092915050565b601f8211156111c857805f5260205f20601f840160051c81016020851015611d835750805b601f840160051c820191505b81811015610e20575f8155600101611d8f565b815167ffffffffffffffff811115611dbc57611dbc611777565b611dd081611dca8454611aab565b84611d5e565b602080601f831160018114611e03575f8415611dec5750858301515b5f19600386901b1c1916600185901b178555611e5a565b5f85815260208120601f198616915b82811015611e3157888601518255948401946001909101908401611e12565b5085821015611e4e57878501515f19600388901b60f8161c191681555b505060018460011b0185555b50505050505056fea2646970667358221220c80686eeed4c13c0dcd0c867d1ce3d0492fad3c44c268ecf1822d9ab5756ef7764736f6c63430008170033",
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

// HasChallengedNode is a free data retrieval call binding the contract method 0x38222c7b.
//
// Solidity: function hasChallengedNode(address requestorAddress, string topicName, address nodeToChallenge, address challengerAddress) view returns(bool)
func (_Votemanager *VotemanagerCaller) HasChallengedNode(opts *bind.CallOpts, requestorAddress common.Address, topicName string, nodeToChallenge common.Address, challengerAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Votemanager.contract.Call(opts, &out, "hasChallengedNode", requestorAddress, topicName, nodeToChallenge, challengerAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasChallengedNode is a free data retrieval call binding the contract method 0x38222c7b.
//
// Solidity: function hasChallengedNode(address requestorAddress, string topicName, address nodeToChallenge, address challengerAddress) view returns(bool)
func (_Votemanager *VotemanagerSession) HasChallengedNode(requestorAddress common.Address, topicName string, nodeToChallenge common.Address, challengerAddress common.Address) (bool, error) {
	return _Votemanager.Contract.HasChallengedNode(&_Votemanager.CallOpts, requestorAddress, topicName, nodeToChallenge, challengerAddress)
}

// HasChallengedNode is a free data retrieval call binding the contract method 0x38222c7b.
//
// Solidity: function hasChallengedNode(address requestorAddress, string topicName, address nodeToChallenge, address challengerAddress) view returns(bool)
func (_Votemanager *VotemanagerCallerSession) HasChallengedNode(requestorAddress common.Address, topicName string, nodeToChallenge common.Address, challengerAddress common.Address) (bool, error) {
	return _Votemanager.Contract.HasChallengedNode(&_Votemanager.CallOpts, requestorAddress, topicName, nodeToChallenge, challengerAddress)
}

// IsChallengeSuccessfulTrainer is a free data retrieval call binding the contract method 0x43faf37f.
//
// Solidity: function isChallengeSuccessfulTrainer(address requestorAddress, string topicName, address trainerAddress) view returns(bool)
func (_Votemanager *VotemanagerCaller) IsChallengeSuccessfulTrainer(opts *bind.CallOpts, requestorAddress common.Address, topicName string, trainerAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Votemanager.contract.Call(opts, &out, "isChallengeSuccessfulTrainer", requestorAddress, topicName, trainerAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsChallengeSuccessfulTrainer is a free data retrieval call binding the contract method 0x43faf37f.
//
// Solidity: function isChallengeSuccessfulTrainer(address requestorAddress, string topicName, address trainerAddress) view returns(bool)
func (_Votemanager *VotemanagerSession) IsChallengeSuccessfulTrainer(requestorAddress common.Address, topicName string, trainerAddress common.Address) (bool, error) {
	return _Votemanager.Contract.IsChallengeSuccessfulTrainer(&_Votemanager.CallOpts, requestorAddress, topicName, trainerAddress)
}

// IsChallengeSuccessfulTrainer is a free data retrieval call binding the contract method 0x43faf37f.
//
// Solidity: function isChallengeSuccessfulTrainer(address requestorAddress, string topicName, address trainerAddress) view returns(bool)
func (_Votemanager *VotemanagerCallerSession) IsChallengeSuccessfulTrainer(requestorAddress common.Address, topicName string, trainerAddress common.Address) (bool, error) {
	return _Votemanager.Contract.IsChallengeSuccessfulTrainer(&_Votemanager.CallOpts, requestorAddress, topicName, trainerAddress)
}

// IsChallengeSuccessfulValidator is a free data retrieval call binding the contract method 0xca88ae24.
//
// Solidity: function isChallengeSuccessfulValidator(address requestorAddress, string topicName, address validatorAddress) view returns(bool)
func (_Votemanager *VotemanagerCaller) IsChallengeSuccessfulValidator(opts *bind.CallOpts, requestorAddress common.Address, topicName string, validatorAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Votemanager.contract.Call(opts, &out, "isChallengeSuccessfulValidator", requestorAddress, topicName, validatorAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsChallengeSuccessfulValidator is a free data retrieval call binding the contract method 0xca88ae24.
//
// Solidity: function isChallengeSuccessfulValidator(address requestorAddress, string topicName, address validatorAddress) view returns(bool)
func (_Votemanager *VotemanagerSession) IsChallengeSuccessfulValidator(requestorAddress common.Address, topicName string, validatorAddress common.Address) (bool, error) {
	return _Votemanager.Contract.IsChallengeSuccessfulValidator(&_Votemanager.CallOpts, requestorAddress, topicName, validatorAddress)
}

// IsChallengeSuccessfulValidator is a free data retrieval call binding the contract method 0xca88ae24.
//
// Solidity: function isChallengeSuccessfulValidator(address requestorAddress, string topicName, address validatorAddress) view returns(bool)
func (_Votemanager *VotemanagerCallerSession) IsChallengeSuccessfulValidator(requestorAddress common.Address, topicName string, validatorAddress common.Address) (bool, error) {
	return _Votemanager.Contract.IsChallengeSuccessfulValidator(&_Votemanager.CallOpts, requestorAddress, topicName, validatorAddress)
}

// IsMaliciousValidationJob is a free data retrieval call binding the contract method 0x95faa057.
//
// Solidity: function isMaliciousValidationJob(address requestorAddress, string topicName) view returns(bool)
func (_Votemanager *VotemanagerCaller) IsMaliciousValidationJob(opts *bind.CallOpts, requestorAddress common.Address, topicName string) (bool, error) {
	var out []interface{}
	err := _Votemanager.contract.Call(opts, &out, "isMaliciousValidationJob", requestorAddress, topicName)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMaliciousValidationJob is a free data retrieval call binding the contract method 0x95faa057.
//
// Solidity: function isMaliciousValidationJob(address requestorAddress, string topicName) view returns(bool)
func (_Votemanager *VotemanagerSession) IsMaliciousValidationJob(requestorAddress common.Address, topicName string) (bool, error) {
	return _Votemanager.Contract.IsMaliciousValidationJob(&_Votemanager.CallOpts, requestorAddress, topicName)
}

// IsMaliciousValidationJob is a free data retrieval call binding the contract method 0x95faa057.
//
// Solidity: function isMaliciousValidationJob(address requestorAddress, string topicName) view returns(bool)
func (_Votemanager *VotemanagerCallerSession) IsMaliciousValidationJob(requestorAddress common.Address, topicName string) (bool, error) {
	return _Votemanager.Contract.IsMaliciousValidationJob(&_Votemanager.CallOpts, requestorAddress, topicName)
}

// IsRefrainingFromValidation is a free data retrieval call binding the contract method 0x97e12c36.
//
// Solidity: function isRefrainingFromValidation(address requestorAddress, string topicName, address validatorAddress) view returns(bool)
func (_Votemanager *VotemanagerCaller) IsRefrainingFromValidation(opts *bind.CallOpts, requestorAddress common.Address, topicName string, validatorAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Votemanager.contract.Call(opts, &out, "isRefrainingFromValidation", requestorAddress, topicName, validatorAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRefrainingFromValidation is a free data retrieval call binding the contract method 0x97e12c36.
//
// Solidity: function isRefrainingFromValidation(address requestorAddress, string topicName, address validatorAddress) view returns(bool)
func (_Votemanager *VotemanagerSession) IsRefrainingFromValidation(requestorAddress common.Address, topicName string, validatorAddress common.Address) (bool, error) {
	return _Votemanager.Contract.IsRefrainingFromValidation(&_Votemanager.CallOpts, requestorAddress, topicName, validatorAddress)
}

// IsRefrainingFromValidation is a free data retrieval call binding the contract method 0x97e12c36.
//
// Solidity: function isRefrainingFromValidation(address requestorAddress, string topicName, address validatorAddress) view returns(bool)
func (_Votemanager *VotemanagerCallerSession) IsRefrainingFromValidation(requestorAddress common.Address, topicName string, validatorAddress common.Address) (bool, error) {
	return _Votemanager.Contract.IsRefrainingFromValidation(&_Votemanager.CallOpts, requestorAddress, topicName, validatorAddress)
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

// SubmitTrainerChallenge is a paid mutator transaction binding the contract method 0x73138ef5.
//
// Solidity: function submitTrainerChallenge(address requestorAddress, string topicName, address trainerAddress, bool isMalicious, address challengerAddress) returns()
func (_Votemanager *VotemanagerTransactor) SubmitTrainerChallenge(opts *bind.TransactOpts, requestorAddress common.Address, topicName string, trainerAddress common.Address, isMalicious bool, challengerAddress common.Address) (*types.Transaction, error) {
	return _Votemanager.contract.Transact(opts, "submitTrainerChallenge", requestorAddress, topicName, trainerAddress, isMalicious, challengerAddress)
}

// SubmitTrainerChallenge is a paid mutator transaction binding the contract method 0x73138ef5.
//
// Solidity: function submitTrainerChallenge(address requestorAddress, string topicName, address trainerAddress, bool isMalicious, address challengerAddress) returns()
func (_Votemanager *VotemanagerSession) SubmitTrainerChallenge(requestorAddress common.Address, topicName string, trainerAddress common.Address, isMalicious bool, challengerAddress common.Address) (*types.Transaction, error) {
	return _Votemanager.Contract.SubmitTrainerChallenge(&_Votemanager.TransactOpts, requestorAddress, topicName, trainerAddress, isMalicious, challengerAddress)
}

// SubmitTrainerChallenge is a paid mutator transaction binding the contract method 0x73138ef5.
//
// Solidity: function submitTrainerChallenge(address requestorAddress, string topicName, address trainerAddress, bool isMalicious, address challengerAddress) returns()
func (_Votemanager *VotemanagerTransactorSession) SubmitTrainerChallenge(requestorAddress common.Address, topicName string, trainerAddress common.Address, isMalicious bool, challengerAddress common.Address) (*types.Transaction, error) {
	return _Votemanager.Contract.SubmitTrainerChallenge(&_Votemanager.TransactOpts, requestorAddress, topicName, trainerAddress, isMalicious, challengerAddress)
}

// SubmitValidatorChallenge is a paid mutator transaction binding the contract method 0xea27dd6b.
//
// Solidity: function submitValidatorChallenge(address requestorAddress, string topicName, address validatorAddress, bool isMalicious, address challengerAddress) returns()
func (_Votemanager *VotemanagerTransactor) SubmitValidatorChallenge(opts *bind.TransactOpts, requestorAddress common.Address, topicName string, validatorAddress common.Address, isMalicious bool, challengerAddress common.Address) (*types.Transaction, error) {
	return _Votemanager.contract.Transact(opts, "submitValidatorChallenge", requestorAddress, topicName, validatorAddress, isMalicious, challengerAddress)
}

// SubmitValidatorChallenge is a paid mutator transaction binding the contract method 0xea27dd6b.
//
// Solidity: function submitValidatorChallenge(address requestorAddress, string topicName, address validatorAddress, bool isMalicious, address challengerAddress) returns()
func (_Votemanager *VotemanagerSession) SubmitValidatorChallenge(requestorAddress common.Address, topicName string, validatorAddress common.Address, isMalicious bool, challengerAddress common.Address) (*types.Transaction, error) {
	return _Votemanager.Contract.SubmitValidatorChallenge(&_Votemanager.TransactOpts, requestorAddress, topicName, validatorAddress, isMalicious, challengerAddress)
}

// SubmitValidatorChallenge is a paid mutator transaction binding the contract method 0xea27dd6b.
//
// Solidity: function submitValidatorChallenge(address requestorAddress, string topicName, address validatorAddress, bool isMalicious, address challengerAddress) returns()
func (_Votemanager *VotemanagerTransactorSession) SubmitValidatorChallenge(requestorAddress common.Address, topicName string, validatorAddress common.Address, isMalicious bool, challengerAddress common.Address) (*types.Transaction, error) {
	return _Votemanager.Contract.SubmitValidatorChallenge(&_Votemanager.TransactOpts, requestorAddress, topicName, validatorAddress, isMalicious, challengerAddress)
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
