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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"ModelAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"ModelRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"validationThreshold\",\"type\":\"uint256\"}],\"name\":\"createValidationProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"}],\"name\":\"getChallengedChallengers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"getModelValidationStatus\",\"outputs\":[{\"internalType\":\"enumValidationStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"nodeToChallenge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"challengerAddress\",\"type\":\"address\"}],\"name\":\"hasChallengedNode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"challengerAddress\",\"type\":\"address\"}],\"name\":\"isChallengeSuccessfulChallenger\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"isChallengeSuccessfulTrainer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"isChallengeSuccessfulValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"}],\"name\":\"isMaliciousValidationJob\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"isRefrainingFromValidation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"refrainFromValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setEvaluationJobTokenContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setScatterProtocolContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"challengedChallengerAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isMalicious\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"challengerAddress\",\"type\":\"address\"}],\"name\":\"submitChallengerChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"submitScoreVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isMalicious\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"challengerAddress\",\"type\":\"address\"}],\"name\":\"submitTrainerChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isMalicious\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"challengerAddress\",\"type\":\"address\"}],\"name\":\"submitValidatorChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506100193361001e565b61006d565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6122878061007a5f395ff3fe608060405234801561000f575f80fd5b506004361061011c575f3560e01c80637c621163116100a9578063ca88ae241161006e578063ca88ae241461025b578063ce83f60f1461026e578063e13181f514610281578063ea27dd6b14610294578063f2fde38b146102a7575f80fd5b80637c621163146101f55780638da5cb5b1461020857806395faa0571461022257806397e12c3614610235578063b12a4c8214610248575f80fd5b806343faf37f116100ef57806343faf37f146101a1578063648c5e97146101b4578063715018a6146101c757806371f65bd7146101cf57806373138ef5146101e2575f80fd5b806308f187c4146101205780631aed680e14610149578063297a43ed1461016957806338222c7b1461017e575b5f80fd5b61013361012e366004611bcb565b6102ba565b6040516101409190611c18565b60405180910390f35b61015c610157366004611c64565b61034c565b6040516101409190611cc4565b61017c610177366004611cea565b610408565b005b61019161018c366004611cea565b610b86565b6040519015158152602001610140565b6101916101af366004611c64565b610c49565b61017c6101c2366004611d68565b610d18565b61017c610f94565b61017c6101dd366004611ded565b610fa7565b61017c6101f0366004611d68565b610fd1565b61017c610203366004611c64565b61115f565b5f546040516001600160a01b039091168152602001610140565b610191610230366004611bcb565b611505565b610191610243366004611c64565b611640565b61017c610256366004611ded565b611697565b610191610269366004611c64565b6116c1565b61017c61027c366004611e08565b61172e565b61019161028f366004611c64565b6117ea565b61017c6102a2366004611d68565b611857565b61017c6102b5366004611ded565b6119e5565b6001600160a01b0382165f90815260136020526040908190209051606091906102e4908490611e7f565b908152604080519182900360209081018320805480830285018301909352828452919083018282801561033e57602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610320575b505050505090505b92915050565b6001600160a01b0383165f908152600860205260408082209051610371908590611e7f565b90815260408051602092819003830190206001600160a01b0385165f908152925290205460ff166103a457506002610401565b6001600160a01b0384165f908152600760205260409081902090516103ca908590611e7f565b90815260408051602092819003830190206001600160a01b0385165f908152925290205460ff16156103fe57506001610401565b505f5b9392505050565b6003546001600160a01b0316331461048d5760405162461bcd60e51b815260206004820152603f60248201527f54686973206d6574686f642063616e206f6e6c792062652063616c6c6564206260448201527f7920746865206576616c756174696f6e20746f6b656e20636f6e74726163740060648201526084015b60405180910390fd5b6001600160a01b0384165f908152600560205260409081902090516104b3908590611e7f565b90815260408051602092819003830190206001600160a01b038581165f908152918452828220908516825290925290205460ff161561056a5760405162461bcd60e51b815260206004820152604760248201527f56616c696461746f722063616e6e6f7420766f746520747769636520666f722060448201527f612073696e676c65206d6f64656c206f6e20612076616c69646174696f6e20706064820152661c9bdc1bdcd85b60ca1b608482015260a401610484565b6001600160a01b0384165f90815260096020526040908190209051610590908590611e7f565b90815260408051602092819003830190206001600160a01b0385165f908152925290205460ff161561062a5760405162461bcd60e51b815260206004820152603c60248201527f43616e6e6f74206576616c7561746520612076616c69646174696f6e2070726f60448201527f706f73616c206166746572206d61726b696e672069742066616c7365000000006064820152608401610484565b6001600160a01b0384165f9081526005602052604090819020905160019190610654908690611e7f565b90815260408051602092819003830181206001600160a01b038781165f9081529185528382208280528552838220805460ff191696151596909617909555938816845260059092529091206001916106ad908690611e7f565b90815260408051602092819003830181206001600160a01b038781165f90815291855283822087821683528552838220805460ff19169615159690961790955593881684526004909252822090610705908690611e7f565b9081526040805191829003602090810183206060840190925281546001600160a01b031683526001820180549184019161073e90611e9a565b80601f016020809104026020016040519081016040528092919081815260200182805461076a90611e9a565b80156107b55780601f1061078c576101008083540402835291602001916107b5565b820191905f5260205f20905b81548152906001019060200180831161079857829003601f168201915b505050505081526020016002820154815250509050600160065f876001600160a01b03166001600160a01b031681526020019081526020015f20856040516107fd9190611e7f565b90815260200160405180910390205f846001600160a01b03166001600160a01b031681526020019081526020015f205f82825461083a9190611ee6565b90915550506001600160a01b0385165f908152600660205260408082209051610864908790611e7f565b9081526040805191829003602090810183206001600160a01b038088165f9081529190925291822054600154637e720f6d60e01b85529094509192911690637e720f6d906108b8908a908a90600401611f24565b5f60405180830381865afa1580156108d2573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526108f99190810190611f4f565b519050808203610b7d5760035460405163d7287ff760e01b81525f916001600160a01b03169063d7287ff790610937908b908b908a90600401611ffc565b602060405180830381865afa158015610952573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109769190612030565b6001600160a01b0389165f908152600860205260409081902090519192506001916109a2908a90611e7f565b90815260408051602092819003830190206001600160a01b0389165f908152925290819020805460ff19169215159290921790915584015181101580610a4757506040516395faa05760e01b815230906395faa05790610a08908b908b90600401611f24565b602060405180830381865afa158015610a23573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a479190612047565b15610ae8577fbc327debbdb9caec6930e265a65fbce81ad18c570f59750876735d5f94142014888887604051610a7f93929190611ffc565b60405180910390a16001600160a01b0388165f9081526007602052604090819020905160019190610ab1908a90611e7f565b90815260408051602092819003830190206001600160a01b0389165f90815292529020805460ff1916911515919091179055610b7b565b7f365b665108d2a8a56d80311e8038ed9d127a08de675009f04c38ba78417403a3888887604051610b1b93929190611ffc565b60405180910390a16001600160a01b0388165f908152600760205260408082209051610b48908a90611e7f565b90815260408051602092819003830190206001600160a01b0389165f90815292529020805460ff19169115159190911790555b505b50505050505050565b6001600160a01b0384165f908152600a60205260408082209051610bab908690611e7f565b90815260408051602092819003830190206001600160a01b038681165f908152918452828220908616825290925290205460ff1680610c4057506001600160a01b0385165f908152600d6020526040908190209051610c0b908690611e7f565b90815260408051602092819003830190206001600160a01b038681165f908152918452828220908616825290925290205460ff165b95945050505050565b6001600160a01b0383165f908152600e60205260408082209051829190610c71908690611e7f565b9081526040805191829003602090810183206001600160a01b038088165f90815291835283822054908a168252600f9092529182209093509091610cb6908790611e7f565b90815260408051602092819003830190206001600160a01b0387165f908152925290205490508115801590610d0157506003610cf3826002612062565b610cfd9190612079565b8210155b8015610d0e575060028110155b9695505050505050565b6001546001600160a01b03163314610d425760405162461bcd60e51b815260040161048490612098565b6001600160a01b0385165f90815260106020526040908190209051610d68908690611e7f565b90815260408051602092819003830190206001600160a01b038681165f908152918452828220908516825290925290205460ff16610f8d576001600160a01b0385165f9081526010602052604090819020905160019190610dca908790611e7f565b90815260408051602092819003830181206001600160a01b038881165f90815291855283822087821683528552838220805460ff1916961515969096179095559389168452601290925290912090610e23908690611e7f565b90815260408051602092819003830190206001600160a01b0386165f90815292528120549003610ead576001600160a01b0385165f90815260136020526040908190209051610e73908690611e7f565b9081526040516020918190038201902080546001810182555f9182529190200180546001600160a01b0319166001600160a01b0385161790555b8115610f20576001600160a01b0385165f9081526011602052604090819020905160019190610edd908790611e7f565b90815260200160405180910390205f856001600160a01b03166001600160a01b031681526020019081526020015f205f828254610f1a9190611ee6565b90915550505b6001600160a01b0385165f9081526012602052604090819020905160019190610f4a908790611e7f565b90815260200160405180910390205f856001600160a01b03166001600160a01b031681526020019081526020015f205f828254610f879190611ee6565b90915550505b5050505050565b610f9c611a5e565b610fa55f611ab7565b565b610faf611a5e565b600380546001600160a01b0319166001600160a01b0392909216919091179055565b6001546001600160a01b03163314610ffb5760405162461bcd60e51b815260040161048490612098565b6001600160a01b0385165f908152600d6020526040908190209051611021908690611e7f565b90815260408051602092819003830190206001600160a01b038681165f908152918452828220908516825290925290205460ff16610f8d576001600160a01b0385165f908152600d602052604090819020905160019190611083908790611e7f565b90815260408051602092819003830190206001600160a01b038781165f90815291845282822090861682529092529020805460ff19169115159190911790558115611135576001600160a01b0385165f908152600e6020526040908190209051600191906110f2908790611e7f565b90815260200160405180910390205f856001600160a01b03166001600160a01b031681526020019081526020015f205f82825461112f9190611ee6565b90915550505b6001600160a01b0385165f908152600f602052604090819020905160019190610f4a908790611e7f565b6001546001600160a01b031633146111895760405162461bcd60e51b815260040161048490612098565b6001600160a01b0383165f908152600560205260409081902090516111af908490611e7f565b90815260408051602092819003830190206001600160a01b0384165f90815290835281812081805290925290205460ff161561126a5760405162461bcd60e51b815260206004820152604e60248201527f56616c696461746f722063616e6e6f74206d61726b2061206a6f62206173206d60448201527f616c6963696f757320616674657220686176696e67207375626d69747465642060648201526d18481cd8dbdc9948199bdc881a5d60921b608482015260a401610484565b6001600160a01b0383165f90815260096020526040908190209051611290908490611e7f565b90815260408051602092819003830190206001600160a01b0384165f908152925290205460ff16611500576001600160a01b0383165f90815260096020526040908190209051600191906112e5908590611e7f565b90815260408051602092819003830181206001600160a01b038681165f90815291909452918220805460ff19169415159490941790935560015463289a278f60e01b8452909291169063289a278f906113449087908790600401611f24565b5f60405180830381865afa15801561135e573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526113859190810190611f4f565b90505f5b8151811015610f8d575f8282815181106113a5576113a56120ff565b602002602001015190505f6001600160a01b0316816001600160a01b0316036113ce5750610f8d565b60035460405163032a1cdd60e01b81526001600160a01b039091169063032a1cdd90611404908990899089908790600401612113565b5f604051808303815f87803b15801561141b575f80fd5b505af115801561142d573d5f803e3d5ffd5b5050506001600160a01b0387165f908152600560205260409081902090516001925061145a908890611e7f565b90815260408051602092819003830181206001600160a01b038981165f90815291855283822087821683528552838220805460ff191696151596909617909555938a16845260069092529091206001916114b5908890611e7f565b90815260200160405180910390205f836001600160a01b03166001600160a01b031681526020019081526020015f205f8282546114f29190611ee6565b909155505050600101611389565b505050565b600154604051637e720f6d60e01b81525f9182916001600160a01b0390911690637e720f6d9061153b9087908790600401611f24565b5f60405180830381865afa158015611555573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261157c9190810190611f4f565b90505f805b825181101561161b575f83828151811061159d5761159d6120ff565b6020026020010151905060095f886001600160a01b03166001600160a01b031681526020019081526020015f20866040516115d89190611e7f565b90815260408051602092819003830190206001600160a01b0384165f908152925290205460ff16156116125761160f600184611ee6565b92505b50600101611581565b5060038251600261162c9190612062565b6116369190612079565b1115949350505050565b6001600160a01b0383165f908152600960205260408082209051611665908590611e7f565b908152604080519182900360209081019092206001600160a01b0385165f908152925290205460ff1690509392505050565b61169f611a5e565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b0383165f908152600b602052604080822090518291906116e9908690611e7f565b9081526040805191829003602090810183206001600160a01b038088165f90815291835283822054908a168252600c9092529182209093509091610cb6908790611e7f565b6001546001600160a01b031633146117585760405162461bcd60e51b815260040161048490612098565b604080516060810182526001600160a01b03851680825260208083018690528284018590525f918252600490528290209151909190611798908590611e7f565b90815260405160209181900382019020825181546001600160a01b0319166001600160a01b039091161781559082015160018201906117d79082612191565b5060408201518160020155905050505050565b6001600160a01b0383165f908152601160205260408082209051829190611812908690611e7f565b9081526040805191829003602090810183206001600160a01b038088165f90815291835283822054908a16825260129092529182209093509091610cb6908790611e7f565b6001546001600160a01b031633146118815760405162461bcd60e51b815260040161048490612098565b6001600160a01b0385165f908152600a60205260409081902090516118a7908690611e7f565b90815260408051602092819003830190206001600160a01b038681165f908152918452828220908516825290925290205460ff16610f8d576001600160a01b0385165f908152600a602052604090819020905160019190611909908790611e7f565b90815260408051602092819003830190206001600160a01b038781165f90815291845282822090861682529092529020805460ff191691151591909117905581156119bb576001600160a01b0385165f908152600b602052604090819020905160019190611978908790611e7f565b90815260200160405180910390205f856001600160a01b03166001600160a01b031681526020019081526020015f205f8282546119b59190611ee6565b90915550505b6001600160a01b0385165f908152600c602052604090819020905160019190610f4a908790611e7f565b6119ed611a5e565b6001600160a01b038116611a525760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610484565b611a5b81611ab7565b50565b5f546001600160a01b03163314610fa55760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610484565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b0381168114611a5b575f80fd5b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611b5757611b57611b1a565b604052919050565b5f82601f830112611b6e575f80fd5b813567ffffffffffffffff811115611b8857611b88611b1a565b611b9b601f8201601f1916602001611b2e565b818152846020838601011115611baf575f80fd5b816020850160208301375f918101602001919091529392505050565b5f8060408385031215611bdc575f80fd5b8235611be781611b06565b9150602083013567ffffffffffffffff811115611c02575f80fd5b611c0e85828601611b5f565b9150509250929050565b602080825282518282018190525f9190848201906040850190845b81811015611c585783516001600160a01b031683529284019291840191600101611c33565b50909695505050505050565b5f805f60608486031215611c76575f80fd5b8335611c8181611b06565b9250602084013567ffffffffffffffff811115611c9c575f80fd5b611ca886828701611b5f565b9250506040840135611cb981611b06565b809150509250925092565b6020810160038310611ce457634e487b7160e01b5f52602160045260245ffd5b91905290565b5f805f8060808587031215611cfd575f80fd5b8435611d0881611b06565b9350602085013567ffffffffffffffff811115611d23575f80fd5b611d2f87828801611b5f565b9350506040850135611d4081611b06565b91506060850135611d5081611b06565b939692955090935050565b8015158114611a5b575f80fd5b5f805f805f60a08688031215611d7c575f80fd5b8535611d8781611b06565b9450602086013567ffffffffffffffff811115611da2575f80fd5b611dae88828901611b5f565b9450506040860135611dbf81611b06565b92506060860135611dcf81611d5b565b91506080860135611ddf81611b06565b809150509295509295909350565b5f60208284031215611dfd575f80fd5b813561040181611b06565b5f805f60608486031215611e1a575f80fd5b8335611e2581611b06565b9250602084013567ffffffffffffffff811115611e40575f80fd5b611e4c86828701611b5f565b925050604084013590509250925092565b5f5b83811015611e77578181015183820152602001611e5f565b50505f910152565b5f8251611e90818460208701611e5d565b9190910192915050565b600181811c90821680611eae57607f821691505b602082108103611ecc57634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561034657610346611ed2565b5f8151808452611f10816020860160208601611e5d565b601f01601f19169290920160200192915050565b6001600160a01b03831681526040602082018190525f90611f4790830184611ef9565b949350505050565b5f6020808385031215611f60575f80fd5b825167ffffffffffffffff80821115611f77575f80fd5b818501915085601f830112611f8a575f80fd5b815181811115611f9c57611f9c611b1a565b8060051b9150611fad848301611b2e565b8181529183018401918481019088841115611fc6575f80fd5b938501935b83851015611ff05784519250611fe083611b06565b8282529385019390850190611fcb565b98975050505050505050565b5f60018060a01b0380861683526060602084015261201d6060840186611ef9565b9150808416604084015250949350505050565b5f60208284031215612040575f80fd5b5051919050565b5f60208284031215612057575f80fd5b815161040181611d5b565b808202811582820484141761034657610346611ed2565b5f8261209357634e487b7160e01b5f52601260045260245ffd5b500490565b60208082526041908201527f54686973206d6574686f642063616e206f6e6c792062652063616c6c6564206260408201527f792074686520747261696e696e67206a6f6220746f6b656e20636f6e747261636060820152601d60fa1b608082015260a00190565b634e487b7160e01b5f52603260045260245ffd5b5f60018060a01b038087168352608060208401526121346080840187611ef9565b9481166040840152929092166060909101525092915050565b601f82111561150057805f5260205f20601f840160051c810160208510156121725750805b601f840160051c820191505b81811015610f8d575f815560010161217e565b815167ffffffffffffffff8111156121ab576121ab611b1a565b6121bf816121b98454611e9a565b8461214d565b602080601f8311600181146121f2575f84156121db5750858301515b5f19600386901b1c1916600185901b178555612249565b5f85815260208120601f198616915b8281101561222057888601518255948401946001909101908401612201565b508582101561223d57878501515f19600388901b60f8161c191681555b505060018460011b0185555b50505050505056fea2646970667358221220a44c627acedb54bbbed0d8a0ff622fa3363b6525e7cb131c13cfb6805276651a64736f6c63430008170033",
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

// GetChallengedChallengers is a free data retrieval call binding the contract method 0x08f187c4.
//
// Solidity: function getChallengedChallengers(address requestorAddress, string topicName) view returns(address[])
func (_Votemanager *VotemanagerCaller) GetChallengedChallengers(opts *bind.CallOpts, requestorAddress common.Address, topicName string) ([]common.Address, error) {
	var out []interface{}
	err := _Votemanager.contract.Call(opts, &out, "getChallengedChallengers", requestorAddress, topicName)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetChallengedChallengers is a free data retrieval call binding the contract method 0x08f187c4.
//
// Solidity: function getChallengedChallengers(address requestorAddress, string topicName) view returns(address[])
func (_Votemanager *VotemanagerSession) GetChallengedChallengers(requestorAddress common.Address, topicName string) ([]common.Address, error) {
	return _Votemanager.Contract.GetChallengedChallengers(&_Votemanager.CallOpts, requestorAddress, topicName)
}

// GetChallengedChallengers is a free data retrieval call binding the contract method 0x08f187c4.
//
// Solidity: function getChallengedChallengers(address requestorAddress, string topicName) view returns(address[])
func (_Votemanager *VotemanagerCallerSession) GetChallengedChallengers(requestorAddress common.Address, topicName string) ([]common.Address, error) {
	return _Votemanager.Contract.GetChallengedChallengers(&_Votemanager.CallOpts, requestorAddress, topicName)
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

// IsChallengeSuccessfulChallenger is a free data retrieval call binding the contract method 0xe13181f5.
//
// Solidity: function isChallengeSuccessfulChallenger(address requestorAddress, string topicName, address challengerAddress) view returns(bool)
func (_Votemanager *VotemanagerCaller) IsChallengeSuccessfulChallenger(opts *bind.CallOpts, requestorAddress common.Address, topicName string, challengerAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Votemanager.contract.Call(opts, &out, "isChallengeSuccessfulChallenger", requestorAddress, topicName, challengerAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsChallengeSuccessfulChallenger is a free data retrieval call binding the contract method 0xe13181f5.
//
// Solidity: function isChallengeSuccessfulChallenger(address requestorAddress, string topicName, address challengerAddress) view returns(bool)
func (_Votemanager *VotemanagerSession) IsChallengeSuccessfulChallenger(requestorAddress common.Address, topicName string, challengerAddress common.Address) (bool, error) {
	return _Votemanager.Contract.IsChallengeSuccessfulChallenger(&_Votemanager.CallOpts, requestorAddress, topicName, challengerAddress)
}

// IsChallengeSuccessfulChallenger is a free data retrieval call binding the contract method 0xe13181f5.
//
// Solidity: function isChallengeSuccessfulChallenger(address requestorAddress, string topicName, address challengerAddress) view returns(bool)
func (_Votemanager *VotemanagerCallerSession) IsChallengeSuccessfulChallenger(requestorAddress common.Address, topicName string, challengerAddress common.Address) (bool, error) {
	return _Votemanager.Contract.IsChallengeSuccessfulChallenger(&_Votemanager.CallOpts, requestorAddress, topicName, challengerAddress)
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

// SubmitChallengerChallenge is a paid mutator transaction binding the contract method 0x648c5e97.
//
// Solidity: function submitChallengerChallenge(address requestorAddress, string topicName, address challengedChallengerAddress, bool isMalicious, address challengerAddress) returns()
func (_Votemanager *VotemanagerTransactor) SubmitChallengerChallenge(opts *bind.TransactOpts, requestorAddress common.Address, topicName string, challengedChallengerAddress common.Address, isMalicious bool, challengerAddress common.Address) (*types.Transaction, error) {
	return _Votemanager.contract.Transact(opts, "submitChallengerChallenge", requestorAddress, topicName, challengedChallengerAddress, isMalicious, challengerAddress)
}

// SubmitChallengerChallenge is a paid mutator transaction binding the contract method 0x648c5e97.
//
// Solidity: function submitChallengerChallenge(address requestorAddress, string topicName, address challengedChallengerAddress, bool isMalicious, address challengerAddress) returns()
func (_Votemanager *VotemanagerSession) SubmitChallengerChallenge(requestorAddress common.Address, topicName string, challengedChallengerAddress common.Address, isMalicious bool, challengerAddress common.Address) (*types.Transaction, error) {
	return _Votemanager.Contract.SubmitChallengerChallenge(&_Votemanager.TransactOpts, requestorAddress, topicName, challengedChallengerAddress, isMalicious, challengerAddress)
}

// SubmitChallengerChallenge is a paid mutator transaction binding the contract method 0x648c5e97.
//
// Solidity: function submitChallengerChallenge(address requestorAddress, string topicName, address challengedChallengerAddress, bool isMalicious, address challengerAddress) returns()
func (_Votemanager *VotemanagerTransactorSession) SubmitChallengerChallenge(requestorAddress common.Address, topicName string, challengedChallengerAddress common.Address, isMalicious bool, challengerAddress common.Address) (*types.Transaction, error) {
	return _Votemanager.Contract.SubmitChallengerChallenge(&_Votemanager.TransactOpts, requestorAddress, topicName, challengedChallengerAddress, isMalicious, challengerAddress)
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
