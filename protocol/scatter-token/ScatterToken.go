// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package scattertoken

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

// ScattertokenMetaData contains all meta data concerning the Scattertoken contract.
var ScattertokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensUnstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addressToStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"canBecomeValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"}],\"name\":\"donateToLottery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwnStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"trainers\",\"type\":\"address[]\"}],\"name\":\"punishRogueTrainers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"punishRogueValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"removeStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"requestorLockToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requiredModelValidatorStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"trainers\",\"type\":\"address[]\"}],\"name\":\"rewardBenevolentTrainers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"rewardBenevolentValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scatterProtocolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"setScatterProtocolAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stakeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"trainerLockToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040525f600b556002600c556064600d556032600e556032600f5534801562000028575f80fd5b5060405162002334380380620023348339810160408190526200004b9162000330565b620000596012600a62000457565b6200006590826200046e565b6040518060400160405280600c81526020016b29b1b0ba3a32b92a37b5b2b760a11b8152506040518060400160405280600281526020016114d560f21b8152508160039081620000b6919062000525565b506004620000c5828262000525565b5050505f81116200011d5760405162461bcd60e51b815260206004820152601560248201527f45524332304361707065643a206361702069732030000000000000000000000060448201526064015b60405180910390fd5b608052600680546001600160a01b03191633908117909155620001629062000143601290565b6200015090600a62000457565b6200015c90846200046e565b6200017c565b506127106005556010805460ff60a01b1916905562000607565b608051816200018a60025490565b620001969190620005f1565b1115620001e65760405162461bcd60e51b815260206004820152601960248201527f45524332304361707065643a2063617020657863656564656400000000000000604482015260640162000114565b620001f28282620001f6565b5050565b608051816200020460025490565b620002109190620005f1565b1115620002605760405162461bcd60e51b815260206004820152601960248201527f45524332304361707065643a2063617020657863656564656400000000000000604482015260640162000114565b620001f282826001600160a01b038216620002be5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640162000114565b8060025f828254620002d19190620005f1565b90915550506001600160a01b0382165f81815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3620001f25f83835b505050565b5f6020828403121562000341575f80fd5b5051919050565b634e487b7160e01b5f52601160045260245ffd5b600181815b808511156200039c57815f190482111562000380576200038062000348565b808516156200038e57918102915b93841c939080029062000361565b509250929050565b5f82620003b45750600162000451565b81620003c257505f62000451565b8160018114620003db5760028114620003e65762000406565b600191505062000451565b60ff841115620003fa57620003fa62000348565b50506001821b62000451565b5060208310610133831016604e8410600b84101617156200042b575081810a62000451565b6200043783836200035c565b805f19048211156200044d576200044d62000348565b0290505b92915050565b5f6200046760ff841683620003a4565b9392505050565b808202811582820484141762000451576200045162000348565b634e487b7160e01b5f52604160045260245ffd5b600181811c90821680620004b157607f821691505b602082108103620004d057634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156200032b57805f5260205f20601f840160051c81016020851015620004fd5750805b601f840160051c820191505b818110156200051e575f815560010162000509565b5050505050565b81516001600160401b0381111562000541576200054162000488565b62000559816200055284546200049c565b84620004d6565b602080601f8311600181146200058f575f8415620005775750858301515b5f19600386901b1c1916600185901b178555620005e9565b5f85815260208120601f198616915b82811015620005bf578886015182559484019460019091019084016200059e565b5085821015620005dd57878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b8082018082111562000451576200045162000348565b608051611d066200062e5f395f8181610257015281816114fd01526115810152611d065ff3fe608060405234801561000f575f80fd5b50600436106101dc575f3560e01c806370a0823111610109578063a1a33f441161009e578063ccd712d51161006e578063ccd712d51461044e578063cda6b84714610461578063dd62ed3e14610474578063fb7ba58214610487575f80fd5b8063a1a33f44146103f6578063a457c2d714610415578063a9059cbb14610428578063bea55db91461043b575f80fd5b8063939624ab116100d9578063939624ab1461039957806395d89b41146103ac578063985fb65d146103b45780639dacf42d146103c7575f80fd5b806370a082311461032057806379cc679014610348578063824205d11461035b5780638da5cb5b1461036e575f80fd5b8063395093511161017f57806343f893511161014f57806343f89351146102d2578063454cbd3e146102e557806347dc19c7146102e55780635e30913f146102f8575f80fd5b80633950935114610284578063408111f51461029757806342966c68146102ac57806342a0f9bf146102bf575f80fd5b806323b872dd116101ba57806323b872dd14610233578063313ce56714610246578063355274ea1461025557806336cb7f801461027b575f80fd5b806306fdde03146101e0578063095ea7b3146101fe57806318160ddd14610221575b5f80fd5b6101e861049b565b6040516101f591906116e2565b60405180910390f35b61021161020c36600461172f565b61052b565b60405190151581526020016101f5565b6002545b6040519081526020016101f5565b610211610241366004611757565b610544565b604051601281526020016101f5565b7f0000000000000000000000000000000000000000000000000000000000000000610225565b61022560055481565b61021161029236600461172f565b610567565b6102aa6102a53660046118c5565b610588565b005b6102aa6102ba366004611934565b6106c1565b6102aa6102cd36600461194b565b6106ce565b6102aa6102e036600461199e565b6107c5565b6102aa6102f33660046118c5565b6108c8565b6102256103063660046119d8565b6001600160a01b03165f9081526007602052604090205490565b61022561032e3660046119d8565b6001600160a01b03165f9081526020819052604090205490565b6102aa61035636600461172f565b6108f7565b6102aa6103693660046119f8565b61090c565b600654610381906001600160a01b031681565b6040516001600160a01b0390911681526020016101f5565b6102aa6103a7366004611934565b6109ef565b6101e8610c7c565b601054610381906001600160a01b031681565b6102116103d53660046119d8565b6005546001600160a01b039091165f90815260076020526040902054101590565b6102256104043660046119d8565b60076020525f908152604090205481565b61021161042336600461172f565b610c8b565b61021161043636600461172f565b610d05565b6102aa6104493660046119d8565b610d12565b6102aa61045c366004611a43565b610d9c565b6102aa61046f366004611934565b610ebb565b610225610482366004611aa5565b61106a565b335f90815260076020526040902054610225565b6060600380546104aa90611ad6565b80601f01602080910402602001604051908101604052809291908181526020018280546104d690611ad6565b80156105215780601f106104f857610100808354040283529160200191610521565b820191905f5260205f20905b81548152906001019060200180831161050457829003601f168201915b5050505050905090565b5f33610538818585611094565b60019150505b92915050565b5f336105518582856111b7565b61055c858585611229565b506001949350505050565b5f33610538818585610579838361106a565b6105839190611b22565b611094565b6010546001600160a01b031633146105bb5760405162461bcd60e51b81526004016105b290611b35565b60405180910390fd5b5f5b81518110156106bb576001600160a01b0384165f908152600960205260409081902090516105ec908590611b92565b90815260200160405180910390205f83838151811061060d5761060d611bad565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f2054600b5f8282546106459190611b22565b90915550506001600160a01b0384165f90815260096020526040808220905161066f908690611b92565b90815260200160405180910390205f84848151811061069057610690611bad565b6020908102919091018101516001600160a01b031682528101919091526040015f20556001016105bd565b50505050565b6106cb33826113cb565b50565b6010546001600160a01b031633146106f85760405162461bcd60e51b81526004016105b290611b35565b6040516370a0823160e01b81526001600160a01b038416600482015230906370a0823190602401602060405180830381865afa15801561073a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061075e9190611bc1565b81111561077d5760405162461bcd60e51b81526004016105b290611bd8565b61078783826113cb565b6001600160a01b0383165f908152600a60205260409081902090518291906107b0908590611b92565b90815260405190819003602001902055505050565b6010546001600160a01b031633146107ef5760405162461bcd60e51b81526004016105b290611b35565b5f5b81518110156108c4575f6064600d5460075f86868151811061081557610815611bad565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f20546108479190611c1c565b6108519190611c33565b905080600b5f8282546108649190611b22565b925050819055508060075f85858151811061088157610881611bad565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f205f8282546108b69190611c52565b9091555050506001016107f1565b5050565b6010546001600160a01b031633146108f25760405162461bcd60e51b81526004016105b290611b35565b505050565b6109028233836111b7565b6108c482826113cb565b6010546001600160a01b031633146109365760405162461bcd60e51b81526004016105b290611b35565b6001600160a01b0382165f908152600a6020526040808220905161095b908490611b92565b90815260200160405180910390205490505f6064600c548361097d9190611c1c565b6109879190611c33565b905080600b5f82825461099a9190611b22565b90915550506001600160a01b0384165f908152600a60205260409081902090518291906109c8908690611b92565b90815260200160405180910390205f8282546109e49190611c52565b909155505050505050565b601054600160a01b900460ff1615610a3a5760405162461bcd60e51b815260206004820152600e60248201526d4e6f2072652d656e7472616e637960901b60448201526064016105b2565b6010805460ff60a01b1916600160a01b179055335f90815260086020526040902054421115610ad15760405162461bcd60e51b815260206004820152603860248201527f596f75206d7573742077616974206265666f7265206265696e672061626c652060448201527f746f20756e7374616b65205363617474657220546f6b656e000000000000000060648201526084016105b2565b335f908152600760205260409020548110610b3d5760405162461bcd60e51b815260206004820152602660248201527f416d6f756e74206d757374206265206c657373207468616e207374616b656420604482015265185b5bdd5b9d60d21b60648201526084016105b2565b610b4733826114fb565b335f9081526007602052604081208054839290610b65908490611c52565b9091555050601054604051630a9d3d9b60e11b81523360048201525f916001600160a01b03169063153a7b3690602401602060405180830381865afa158015610bb0573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bd49190611c65565b90506003816004811115610bea57610bea611c83565b148015610c065750600554335f90815260076020526040902054105b15610c6b576010546040516302b8e1f360e51b81526001600160a01b039091169063571c3e6090610c3d9033905f90600401611c97565b5f604051808303815f87803b158015610c54575f80fd5b505af1158015610c66573d5f803e3d5ffd5b505050505b50506010805460ff60a01b19169055565b6060600480546104aa90611ad6565b5f3381610c98828661106a565b905083811015610cf85760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084016105b2565b61055c8286868403611094565b5f33610538818585611229565b6006546001600160a01b03163314610d7a5760405162461bcd60e51b815260206004820152602560248201527f4f6e6c7920746865206f776e65722063616e2063616c6c20746869732066756e60448201526431ba34b7b760d91b60648201526084016105b2565b601080546001600160a01b0319166001600160a01b0392909216919091179055565b6010546001600160a01b03163314610dc65760405162461bcd60e51b81526004016105b290611b35565b6040516370a0823160e01b81526001600160a01b038516600482015230906370a0823190602401602060405180830381865afa158015610e08573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e2c9190611bc1565b811115610e4b5760405162461bcd60e51b81526004016105b290611bd8565b610e5584826113cb565b6001600160a01b0383165f90815260096020526040908190209051829190610e7e908590611b92565b90815260200160405180910390205f866001600160a01b03166001600160a01b031681526020019081526020015f205f8282546109e49190611b22565b601054600160a01b900460ff1615610f065760405162461bcd60e51b815260206004820152600e60248201526d4e6f2072652d656e7472616e637960901b60448201526064016105b2565b6010805460ff60a01b1916600160a01b1790556040516370a0823160e01b815233600482015230906370a0823190602401602060405180830381865afa158015610f52573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f769190611bc1565b811115610fd35760405162461bcd60e51b815260206004820152602560248201527f43616e6e6f74207374616b65206d6f726520746f6b656e73207468616e20796f6044820152643a9037bbb760d91b60648201526084016105b2565b610fdd33826113cb565b335f9081526007602052604081208054839290610ffb908490611b22565b9091555061100e90504262278d00611b22565b335f818152600860209081526040918290209390935580519182529181018390527fb539ca1e5c8d398ddf1c41c30166f33404941683be4683319b57669a93dad4ef910160405180910390a1506010805460ff60a01b19169055565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b6001600160a01b0383166110f65760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084016105b2565b6001600160a01b0382166111575760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016105b2565b6001600160a01b038381165f8181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b5f6111c2848461106a565b90505f1981146106bb578181101561121c5760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016105b2565b6106bb8484848403611094565b6001600160a01b03831661128d5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016105b2565b6001600160a01b0382166112ef5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016105b2565b6001600160a01b0383165f90815260208190526040902054818110156113665760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b60648201526084016105b2565b6001600160a01b038481165f81815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a36106bb565b6001600160a01b03821661142b5760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b60648201526084016105b2565b6001600160a01b0382165f908152602081905260409020548181101561149e5760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b60648201526084016105b2565b6001600160a01b0383165f818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3505050565b7f00000000000000000000000000000000000000000000000000000000000000008161152660025490565b6115309190611b22565b111561157a5760405162461bcd60e51b8152602060048201526019602482015278115490cc8c10d85c1c19590e8818d85c08195e18d959591959603a1b60448201526064016105b2565b6108c482827f0000000000000000000000000000000000000000000000000000000000000000816115aa60025490565b6115b49190611b22565b11156115fe5760405162461bcd60e51b8152602060048201526019602482015278115490cc8c10d85c1c19590e8818d85c08195e18d959591959603a1b60448201526064016105b2565b6108c482826001600160a01b0382166116595760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016105b2565b8060025f82825461166a9190611b22565b90915550506001600160a01b0382165f81815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b5f5b838110156116da5781810151838201526020016116c2565b50505f910152565b602081525f82518060208401526117008160408501602087016116c0565b601f01601f19169190910160400192915050565b80356001600160a01b038116811461172a575f80fd5b919050565b5f8060408385031215611740575f80fd5b61174983611714565b946020939093013593505050565b5f805f60608486031215611769575f80fd5b61177284611714565b925061178060208501611714565b9150604084013590509250925092565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff811182821017156117cd576117cd611790565b604052919050565b5f82601f8301126117e4575f80fd5b813567ffffffffffffffff8111156117fe576117fe611790565b611811601f8201601f19166020016117a4565b818152846020838601011115611825575f80fd5b816020850160208301375f918101602001919091529392505050565b5f82601f830112611850575f80fd5b8135602067ffffffffffffffff82111561186c5761186c611790565b8160051b61187b8282016117a4565b9283528481018201928281019087851115611894575f80fd5b83870192505b848310156118ba576118ab83611714565b8252918301919083019061189a565b979650505050505050565b5f805f606084860312156118d7575f80fd5b6118e084611714565b9250602084013567ffffffffffffffff808211156118fc575f80fd5b611908878388016117d5565b9350604086013591508082111561191d575f80fd5b5061192a86828701611841565b9150509250925092565b5f60208284031215611944575f80fd5b5035919050565b5f805f6060848603121561195d575f80fd5b61196684611714565b9250602084013567ffffffffffffffff811115611981575f80fd5b61198d868287016117d5565b925050604084013590509250925092565b5f602082840312156119ae575f80fd5b813567ffffffffffffffff8111156119c4575f80fd5b6119d084828501611841565b949350505050565b5f602082840312156119e8575f80fd5b6119f182611714565b9392505050565b5f8060408385031215611a09575f80fd5b611a1283611714565b9150602083013567ffffffffffffffff811115611a2d575f80fd5b611a39858286016117d5565b9150509250929050565b5f805f8060808587031215611a56575f80fd5b611a5f85611714565b9350611a6d60208601611714565b9250604085013567ffffffffffffffff811115611a88575f80fd5b611a94878288016117d5565b949793965093946060013593505050565b5f8060408385031215611ab6575f80fd5b611abf83611714565b9150611acd60208401611714565b90509250929050565b600181811c90821680611aea57607f821691505b602082108103611b0857634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561053e5761053e611b0e565b6020808252603f908201527f54686973206d6574686f642063616e206f6e6c792062652063616c6c6564206260408201527f792074686520736361747465722070726f746f636f6c20636f6e747261637400606082015260800190565b5f8251611ba38184602087016116c0565b9190910192915050565b634e487b7160e01b5f52603260045260245ffd5b5f60208284031215611bd1575f80fd5b5051919050565b60208082526024908201527f43616e6e6f74206c6f636b206d6f726520746f6b656e73207468616e20796f756040820152631037bbb760e11b606082015260800190565b808202811582820484141761053e5761053e611b0e565b5f82611c4d57634e487b7160e01b5f52601260045260245ffd5b500490565b8181038181111561053e5761053e611b0e565b5f60208284031215611c75575f80fd5b8151600581106119f1575f80fd5b634e487b7160e01b5f52602160045260245ffd5b6001600160a01b03831681526040810160058310611cc357634e487b7160e01b5f52602160045260245ffd5b826020830152939250505056fea2646970667358221220aeaf1497881934b33b0b5fa2258cd2ca0d2453f7959242df08b7b5c749a8745664736f6c63430008170033",
}

// ScattertokenABI is the input ABI used to generate the binding from.
// Deprecated: Use ScattertokenMetaData.ABI instead.
var ScattertokenABI = ScattertokenMetaData.ABI

// ScattertokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ScattertokenMetaData.Bin instead.
var ScattertokenBin = ScattertokenMetaData.Bin

// DeployScattertoken deploys a new Ethereum contract, binding an instance of Scattertoken to it.
func DeployScattertoken(auth *bind.TransactOpts, backend bind.ContractBackend, cap *big.Int) (common.Address, *types.Transaction, *Scattertoken, error) {
	parsed, err := ScattertokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ScattertokenBin), backend, cap)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Scattertoken{ScattertokenCaller: ScattertokenCaller{contract: contract}, ScattertokenTransactor: ScattertokenTransactor{contract: contract}, ScattertokenFilterer: ScattertokenFilterer{contract: contract}}, nil
}

// Scattertoken is an auto generated Go binding around an Ethereum contract.
type Scattertoken struct {
	ScattertokenCaller     // Read-only binding to the contract
	ScattertokenTransactor // Write-only binding to the contract
	ScattertokenFilterer   // Log filterer for contract events
}

// ScattertokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ScattertokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScattertokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ScattertokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScattertokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ScattertokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScattertokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ScattertokenSession struct {
	Contract     *Scattertoken     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ScattertokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ScattertokenCallerSession struct {
	Contract *ScattertokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ScattertokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ScattertokenTransactorSession struct {
	Contract     *ScattertokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ScattertokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ScattertokenRaw struct {
	Contract *Scattertoken // Generic contract binding to access the raw methods on
}

// ScattertokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ScattertokenCallerRaw struct {
	Contract *ScattertokenCaller // Generic read-only contract binding to access the raw methods on
}

// ScattertokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ScattertokenTransactorRaw struct {
	Contract *ScattertokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewScattertoken creates a new instance of Scattertoken, bound to a specific deployed contract.
func NewScattertoken(address common.Address, backend bind.ContractBackend) (*Scattertoken, error) {
	contract, err := bindScattertoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Scattertoken{ScattertokenCaller: ScattertokenCaller{contract: contract}, ScattertokenTransactor: ScattertokenTransactor{contract: contract}, ScattertokenFilterer: ScattertokenFilterer{contract: contract}}, nil
}

// NewScattertokenCaller creates a new read-only instance of Scattertoken, bound to a specific deployed contract.
func NewScattertokenCaller(address common.Address, caller bind.ContractCaller) (*ScattertokenCaller, error) {
	contract, err := bindScattertoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ScattertokenCaller{contract: contract}, nil
}

// NewScattertokenTransactor creates a new write-only instance of Scattertoken, bound to a specific deployed contract.
func NewScattertokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ScattertokenTransactor, error) {
	contract, err := bindScattertoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ScattertokenTransactor{contract: contract}, nil
}

// NewScattertokenFilterer creates a new log filterer instance of Scattertoken, bound to a specific deployed contract.
func NewScattertokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ScattertokenFilterer, error) {
	contract, err := bindScattertoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ScattertokenFilterer{contract: contract}, nil
}

// bindScattertoken binds a generic wrapper to an already deployed contract.
func bindScattertoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ScattertokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Scattertoken *ScattertokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Scattertoken.Contract.ScattertokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Scattertoken *ScattertokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Scattertoken.Contract.ScattertokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Scattertoken *ScattertokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Scattertoken.Contract.ScattertokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Scattertoken *ScattertokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Scattertoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Scattertoken *ScattertokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Scattertoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Scattertoken *ScattertokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Scattertoken.Contract.contract.Transact(opts, method, params...)
}

// AddressToStake is a free data retrieval call binding the contract method 0xa1a33f44.
//
// Solidity: function addressToStake(address ) view returns(uint256)
func (_Scattertoken *ScattertokenCaller) AddressToStake(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Scattertoken.contract.Call(opts, &out, "addressToStake", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressToStake is a free data retrieval call binding the contract method 0xa1a33f44.
//
// Solidity: function addressToStake(address ) view returns(uint256)
func (_Scattertoken *ScattertokenSession) AddressToStake(arg0 common.Address) (*big.Int, error) {
	return _Scattertoken.Contract.AddressToStake(&_Scattertoken.CallOpts, arg0)
}

// AddressToStake is a free data retrieval call binding the contract method 0xa1a33f44.
//
// Solidity: function addressToStake(address ) view returns(uint256)
func (_Scattertoken *ScattertokenCallerSession) AddressToStake(arg0 common.Address) (*big.Int, error) {
	return _Scattertoken.Contract.AddressToStake(&_Scattertoken.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Scattertoken *ScattertokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Scattertoken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Scattertoken *ScattertokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Scattertoken.Contract.Allowance(&_Scattertoken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Scattertoken *ScattertokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Scattertoken.Contract.Allowance(&_Scattertoken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Scattertoken *ScattertokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Scattertoken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Scattertoken *ScattertokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Scattertoken.Contract.BalanceOf(&_Scattertoken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Scattertoken *ScattertokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Scattertoken.Contract.BalanceOf(&_Scattertoken.CallOpts, account)
}

// CanBecomeValidator is a free data retrieval call binding the contract method 0x9dacf42d.
//
// Solidity: function canBecomeValidator(address account) view returns(bool)
func (_Scattertoken *ScattertokenCaller) CanBecomeValidator(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Scattertoken.contract.Call(opts, &out, "canBecomeValidator", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanBecomeValidator is a free data retrieval call binding the contract method 0x9dacf42d.
//
// Solidity: function canBecomeValidator(address account) view returns(bool)
func (_Scattertoken *ScattertokenSession) CanBecomeValidator(account common.Address) (bool, error) {
	return _Scattertoken.Contract.CanBecomeValidator(&_Scattertoken.CallOpts, account)
}

// CanBecomeValidator is a free data retrieval call binding the contract method 0x9dacf42d.
//
// Solidity: function canBecomeValidator(address account) view returns(bool)
func (_Scattertoken *ScattertokenCallerSession) CanBecomeValidator(account common.Address) (bool, error) {
	return _Scattertoken.Contract.CanBecomeValidator(&_Scattertoken.CallOpts, account)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_Scattertoken *ScattertokenCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Scattertoken.contract.Call(opts, &out, "cap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_Scattertoken *ScattertokenSession) Cap() (*big.Int, error) {
	return _Scattertoken.Contract.Cap(&_Scattertoken.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_Scattertoken *ScattertokenCallerSession) Cap() (*big.Int, error) {
	return _Scattertoken.Contract.Cap(&_Scattertoken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Scattertoken *ScattertokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Scattertoken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Scattertoken *ScattertokenSession) Decimals() (uint8, error) {
	return _Scattertoken.Contract.Decimals(&_Scattertoken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Scattertoken *ScattertokenCallerSession) Decimals() (uint8, error) {
	return _Scattertoken.Contract.Decimals(&_Scattertoken.CallOpts)
}

// GetAccountStake is a free data retrieval call binding the contract method 0x5e30913f.
//
// Solidity: function getAccountStake(address account) view returns(uint256)
func (_Scattertoken *ScattertokenCaller) GetAccountStake(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Scattertoken.contract.Call(opts, &out, "getAccountStake", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountStake is a free data retrieval call binding the contract method 0x5e30913f.
//
// Solidity: function getAccountStake(address account) view returns(uint256)
func (_Scattertoken *ScattertokenSession) GetAccountStake(account common.Address) (*big.Int, error) {
	return _Scattertoken.Contract.GetAccountStake(&_Scattertoken.CallOpts, account)
}

// GetAccountStake is a free data retrieval call binding the contract method 0x5e30913f.
//
// Solidity: function getAccountStake(address account) view returns(uint256)
func (_Scattertoken *ScattertokenCallerSession) GetAccountStake(account common.Address) (*big.Int, error) {
	return _Scattertoken.Contract.GetAccountStake(&_Scattertoken.CallOpts, account)
}

// GetOwnStake is a free data retrieval call binding the contract method 0xfb7ba582.
//
// Solidity: function getOwnStake() view returns(uint256)
func (_Scattertoken *ScattertokenCaller) GetOwnStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Scattertoken.contract.Call(opts, &out, "getOwnStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOwnStake is a free data retrieval call binding the contract method 0xfb7ba582.
//
// Solidity: function getOwnStake() view returns(uint256)
func (_Scattertoken *ScattertokenSession) GetOwnStake() (*big.Int, error) {
	return _Scattertoken.Contract.GetOwnStake(&_Scattertoken.CallOpts)
}

// GetOwnStake is a free data retrieval call binding the contract method 0xfb7ba582.
//
// Solidity: function getOwnStake() view returns(uint256)
func (_Scattertoken *ScattertokenCallerSession) GetOwnStake() (*big.Int, error) {
	return _Scattertoken.Contract.GetOwnStake(&_Scattertoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Scattertoken *ScattertokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Scattertoken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Scattertoken *ScattertokenSession) Name() (string, error) {
	return _Scattertoken.Contract.Name(&_Scattertoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Scattertoken *ScattertokenCallerSession) Name() (string, error) {
	return _Scattertoken.Contract.Name(&_Scattertoken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Scattertoken *ScattertokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Scattertoken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Scattertoken *ScattertokenSession) Owner() (common.Address, error) {
	return _Scattertoken.Contract.Owner(&_Scattertoken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Scattertoken *ScattertokenCallerSession) Owner() (common.Address, error) {
	return _Scattertoken.Contract.Owner(&_Scattertoken.CallOpts)
}

// RequiredModelValidatorStake is a free data retrieval call binding the contract method 0x36cb7f80.
//
// Solidity: function requiredModelValidatorStake() view returns(uint256)
func (_Scattertoken *ScattertokenCaller) RequiredModelValidatorStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Scattertoken.contract.Call(opts, &out, "requiredModelValidatorStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredModelValidatorStake is a free data retrieval call binding the contract method 0x36cb7f80.
//
// Solidity: function requiredModelValidatorStake() view returns(uint256)
func (_Scattertoken *ScattertokenSession) RequiredModelValidatorStake() (*big.Int, error) {
	return _Scattertoken.Contract.RequiredModelValidatorStake(&_Scattertoken.CallOpts)
}

// RequiredModelValidatorStake is a free data retrieval call binding the contract method 0x36cb7f80.
//
// Solidity: function requiredModelValidatorStake() view returns(uint256)
func (_Scattertoken *ScattertokenCallerSession) RequiredModelValidatorStake() (*big.Int, error) {
	return _Scattertoken.Contract.RequiredModelValidatorStake(&_Scattertoken.CallOpts)
}

// ScatterProtocolAddress is a free data retrieval call binding the contract method 0x985fb65d.
//
// Solidity: function scatterProtocolAddress() view returns(address)
func (_Scattertoken *ScattertokenCaller) ScatterProtocolAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Scattertoken.contract.Call(opts, &out, "scatterProtocolAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ScatterProtocolAddress is a free data retrieval call binding the contract method 0x985fb65d.
//
// Solidity: function scatterProtocolAddress() view returns(address)
func (_Scattertoken *ScattertokenSession) ScatterProtocolAddress() (common.Address, error) {
	return _Scattertoken.Contract.ScatterProtocolAddress(&_Scattertoken.CallOpts)
}

// ScatterProtocolAddress is a free data retrieval call binding the contract method 0x985fb65d.
//
// Solidity: function scatterProtocolAddress() view returns(address)
func (_Scattertoken *ScattertokenCallerSession) ScatterProtocolAddress() (common.Address, error) {
	return _Scattertoken.Contract.ScatterProtocolAddress(&_Scattertoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Scattertoken *ScattertokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Scattertoken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Scattertoken *ScattertokenSession) Symbol() (string, error) {
	return _Scattertoken.Contract.Symbol(&_Scattertoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Scattertoken *ScattertokenCallerSession) Symbol() (string, error) {
	return _Scattertoken.Contract.Symbol(&_Scattertoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Scattertoken *ScattertokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Scattertoken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Scattertoken *ScattertokenSession) TotalSupply() (*big.Int, error) {
	return _Scattertoken.Contract.TotalSupply(&_Scattertoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Scattertoken *ScattertokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Scattertoken.Contract.TotalSupply(&_Scattertoken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Scattertoken *ScattertokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scattertoken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Scattertoken *ScattertokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scattertoken.Contract.Approve(&_Scattertoken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Scattertoken *ScattertokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scattertoken.Contract.Approve(&_Scattertoken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Scattertoken *ScattertokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Scattertoken.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Scattertoken *ScattertokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Scattertoken.Contract.Burn(&_Scattertoken.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Scattertoken *ScattertokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Scattertoken.Contract.Burn(&_Scattertoken.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Scattertoken *ScattertokenTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scattertoken.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Scattertoken *ScattertokenSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scattertoken.Contract.BurnFrom(&_Scattertoken.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Scattertoken *ScattertokenTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scattertoken.Contract.BurnFrom(&_Scattertoken.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Scattertoken *ScattertokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Scattertoken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Scattertoken *ScattertokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Scattertoken.Contract.DecreaseAllowance(&_Scattertoken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Scattertoken *ScattertokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Scattertoken.Contract.DecreaseAllowance(&_Scattertoken.TransactOpts, spender, subtractedValue)
}

// DonateToLottery is a paid mutator transaction binding the contract method 0x824205d1.
//
// Solidity: function donateToLottery(address requestorAddress, string topicName) returns()
func (_Scattertoken *ScattertokenTransactor) DonateToLottery(opts *bind.TransactOpts, requestorAddress common.Address, topicName string) (*types.Transaction, error) {
	return _Scattertoken.contract.Transact(opts, "donateToLottery", requestorAddress, topicName)
}

// DonateToLottery is a paid mutator transaction binding the contract method 0x824205d1.
//
// Solidity: function donateToLottery(address requestorAddress, string topicName) returns()
func (_Scattertoken *ScattertokenSession) DonateToLottery(requestorAddress common.Address, topicName string) (*types.Transaction, error) {
	return _Scattertoken.Contract.DonateToLottery(&_Scattertoken.TransactOpts, requestorAddress, topicName)
}

// DonateToLottery is a paid mutator transaction binding the contract method 0x824205d1.
//
// Solidity: function donateToLottery(address requestorAddress, string topicName) returns()
func (_Scattertoken *ScattertokenTransactorSession) DonateToLottery(requestorAddress common.Address, topicName string) (*types.Transaction, error) {
	return _Scattertoken.Contract.DonateToLottery(&_Scattertoken.TransactOpts, requestorAddress, topicName)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Scattertoken *ScattertokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Scattertoken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Scattertoken *ScattertokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Scattertoken.Contract.IncreaseAllowance(&_Scattertoken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Scattertoken *ScattertokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Scattertoken.Contract.IncreaseAllowance(&_Scattertoken.TransactOpts, spender, addedValue)
}

// PunishRogueTrainers is a paid mutator transaction binding the contract method 0x408111f5.
//
// Solidity: function punishRogueTrainers(address requestorAddress, string topicName, address[] trainers) returns()
func (_Scattertoken *ScattertokenTransactor) PunishRogueTrainers(opts *bind.TransactOpts, requestorAddress common.Address, topicName string, trainers []common.Address) (*types.Transaction, error) {
	return _Scattertoken.contract.Transact(opts, "punishRogueTrainers", requestorAddress, topicName, trainers)
}

// PunishRogueTrainers is a paid mutator transaction binding the contract method 0x408111f5.
//
// Solidity: function punishRogueTrainers(address requestorAddress, string topicName, address[] trainers) returns()
func (_Scattertoken *ScattertokenSession) PunishRogueTrainers(requestorAddress common.Address, topicName string, trainers []common.Address) (*types.Transaction, error) {
	return _Scattertoken.Contract.PunishRogueTrainers(&_Scattertoken.TransactOpts, requestorAddress, topicName, trainers)
}

// PunishRogueTrainers is a paid mutator transaction binding the contract method 0x408111f5.
//
// Solidity: function punishRogueTrainers(address requestorAddress, string topicName, address[] trainers) returns()
func (_Scattertoken *ScattertokenTransactorSession) PunishRogueTrainers(requestorAddress common.Address, topicName string, trainers []common.Address) (*types.Transaction, error) {
	return _Scattertoken.Contract.PunishRogueTrainers(&_Scattertoken.TransactOpts, requestorAddress, topicName, trainers)
}

// PunishRogueValidators is a paid mutator transaction binding the contract method 0x43f89351.
//
// Solidity: function punishRogueValidators(address[] validators) returns()
func (_Scattertoken *ScattertokenTransactor) PunishRogueValidators(opts *bind.TransactOpts, validators []common.Address) (*types.Transaction, error) {
	return _Scattertoken.contract.Transact(opts, "punishRogueValidators", validators)
}

// PunishRogueValidators is a paid mutator transaction binding the contract method 0x43f89351.
//
// Solidity: function punishRogueValidators(address[] validators) returns()
func (_Scattertoken *ScattertokenSession) PunishRogueValidators(validators []common.Address) (*types.Transaction, error) {
	return _Scattertoken.Contract.PunishRogueValidators(&_Scattertoken.TransactOpts, validators)
}

// PunishRogueValidators is a paid mutator transaction binding the contract method 0x43f89351.
//
// Solidity: function punishRogueValidators(address[] validators) returns()
func (_Scattertoken *ScattertokenTransactorSession) PunishRogueValidators(validators []common.Address) (*types.Transaction, error) {
	return _Scattertoken.Contract.PunishRogueValidators(&_Scattertoken.TransactOpts, validators)
}

// RemoveStake is a paid mutator transaction binding the contract method 0x939624ab.
//
// Solidity: function removeStake(uint256 amount) returns()
func (_Scattertoken *ScattertokenTransactor) RemoveStake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Scattertoken.contract.Transact(opts, "removeStake", amount)
}

// RemoveStake is a paid mutator transaction binding the contract method 0x939624ab.
//
// Solidity: function removeStake(uint256 amount) returns()
func (_Scattertoken *ScattertokenSession) RemoveStake(amount *big.Int) (*types.Transaction, error) {
	return _Scattertoken.Contract.RemoveStake(&_Scattertoken.TransactOpts, amount)
}

// RemoveStake is a paid mutator transaction binding the contract method 0x939624ab.
//
// Solidity: function removeStake(uint256 amount) returns()
func (_Scattertoken *ScattertokenTransactorSession) RemoveStake(amount *big.Int) (*types.Transaction, error) {
	return _Scattertoken.Contract.RemoveStake(&_Scattertoken.TransactOpts, amount)
}

// RequestorLockToken is a paid mutator transaction binding the contract method 0x42a0f9bf.
//
// Solidity: function requestorLockToken(address requestorAddress, string topicName, uint256 amount) returns()
func (_Scattertoken *ScattertokenTransactor) RequestorLockToken(opts *bind.TransactOpts, requestorAddress common.Address, topicName string, amount *big.Int) (*types.Transaction, error) {
	return _Scattertoken.contract.Transact(opts, "requestorLockToken", requestorAddress, topicName, amount)
}

// RequestorLockToken is a paid mutator transaction binding the contract method 0x42a0f9bf.
//
// Solidity: function requestorLockToken(address requestorAddress, string topicName, uint256 amount) returns()
func (_Scattertoken *ScattertokenSession) RequestorLockToken(requestorAddress common.Address, topicName string, amount *big.Int) (*types.Transaction, error) {
	return _Scattertoken.Contract.RequestorLockToken(&_Scattertoken.TransactOpts, requestorAddress, topicName, amount)
}

// RequestorLockToken is a paid mutator transaction binding the contract method 0x42a0f9bf.
//
// Solidity: function requestorLockToken(address requestorAddress, string topicName, uint256 amount) returns()
func (_Scattertoken *ScattertokenTransactorSession) RequestorLockToken(requestorAddress common.Address, topicName string, amount *big.Int) (*types.Transaction, error) {
	return _Scattertoken.Contract.RequestorLockToken(&_Scattertoken.TransactOpts, requestorAddress, topicName, amount)
}

// RewardBenevolentTrainers is a paid mutator transaction binding the contract method 0x47dc19c7.
//
// Solidity: function rewardBenevolentTrainers(address requestorAddress, string topicName, address[] trainers) returns()
func (_Scattertoken *ScattertokenTransactor) RewardBenevolentTrainers(opts *bind.TransactOpts, requestorAddress common.Address, topicName string, trainers []common.Address) (*types.Transaction, error) {
	return _Scattertoken.contract.Transact(opts, "rewardBenevolentTrainers", requestorAddress, topicName, trainers)
}

// RewardBenevolentTrainers is a paid mutator transaction binding the contract method 0x47dc19c7.
//
// Solidity: function rewardBenevolentTrainers(address requestorAddress, string topicName, address[] trainers) returns()
func (_Scattertoken *ScattertokenSession) RewardBenevolentTrainers(requestorAddress common.Address, topicName string, trainers []common.Address) (*types.Transaction, error) {
	return _Scattertoken.Contract.RewardBenevolentTrainers(&_Scattertoken.TransactOpts, requestorAddress, topicName, trainers)
}

// RewardBenevolentTrainers is a paid mutator transaction binding the contract method 0x47dc19c7.
//
// Solidity: function rewardBenevolentTrainers(address requestorAddress, string topicName, address[] trainers) returns()
func (_Scattertoken *ScattertokenTransactorSession) RewardBenevolentTrainers(requestorAddress common.Address, topicName string, trainers []common.Address) (*types.Transaction, error) {
	return _Scattertoken.Contract.RewardBenevolentTrainers(&_Scattertoken.TransactOpts, requestorAddress, topicName, trainers)
}

// RewardBenevolentValidators is a paid mutator transaction binding the contract method 0x454cbd3e.
//
// Solidity: function rewardBenevolentValidators(address requestorAddress, string topicName, address[] validators) returns()
func (_Scattertoken *ScattertokenTransactor) RewardBenevolentValidators(opts *bind.TransactOpts, requestorAddress common.Address, topicName string, validators []common.Address) (*types.Transaction, error) {
	return _Scattertoken.contract.Transact(opts, "rewardBenevolentValidators", requestorAddress, topicName, validators)
}

// RewardBenevolentValidators is a paid mutator transaction binding the contract method 0x454cbd3e.
//
// Solidity: function rewardBenevolentValidators(address requestorAddress, string topicName, address[] validators) returns()
func (_Scattertoken *ScattertokenSession) RewardBenevolentValidators(requestorAddress common.Address, topicName string, validators []common.Address) (*types.Transaction, error) {
	return _Scattertoken.Contract.RewardBenevolentValidators(&_Scattertoken.TransactOpts, requestorAddress, topicName, validators)
}

// RewardBenevolentValidators is a paid mutator transaction binding the contract method 0x454cbd3e.
//
// Solidity: function rewardBenevolentValidators(address requestorAddress, string topicName, address[] validators) returns()
func (_Scattertoken *ScattertokenTransactorSession) RewardBenevolentValidators(requestorAddress common.Address, topicName string, validators []common.Address) (*types.Transaction, error) {
	return _Scattertoken.Contract.RewardBenevolentValidators(&_Scattertoken.TransactOpts, requestorAddress, topicName, validators)
}

// SetScatterProtocolAddress is a paid mutator transaction binding the contract method 0xbea55db9.
//
// Solidity: function setScatterProtocolAddress(address newAddress) returns()
func (_Scattertoken *ScattertokenTransactor) SetScatterProtocolAddress(opts *bind.TransactOpts, newAddress common.Address) (*types.Transaction, error) {
	return _Scattertoken.contract.Transact(opts, "setScatterProtocolAddress", newAddress)
}

// SetScatterProtocolAddress is a paid mutator transaction binding the contract method 0xbea55db9.
//
// Solidity: function setScatterProtocolAddress(address newAddress) returns()
func (_Scattertoken *ScattertokenSession) SetScatterProtocolAddress(newAddress common.Address) (*types.Transaction, error) {
	return _Scattertoken.Contract.SetScatterProtocolAddress(&_Scattertoken.TransactOpts, newAddress)
}

// SetScatterProtocolAddress is a paid mutator transaction binding the contract method 0xbea55db9.
//
// Solidity: function setScatterProtocolAddress(address newAddress) returns()
func (_Scattertoken *ScattertokenTransactorSession) SetScatterProtocolAddress(newAddress common.Address) (*types.Transaction, error) {
	return _Scattertoken.Contract.SetScatterProtocolAddress(&_Scattertoken.TransactOpts, newAddress)
}

// StakeToken is a paid mutator transaction binding the contract method 0xcda6b847.
//
// Solidity: function stakeToken(uint256 amount) returns()
func (_Scattertoken *ScattertokenTransactor) StakeToken(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Scattertoken.contract.Transact(opts, "stakeToken", amount)
}

// StakeToken is a paid mutator transaction binding the contract method 0xcda6b847.
//
// Solidity: function stakeToken(uint256 amount) returns()
func (_Scattertoken *ScattertokenSession) StakeToken(amount *big.Int) (*types.Transaction, error) {
	return _Scattertoken.Contract.StakeToken(&_Scattertoken.TransactOpts, amount)
}

// StakeToken is a paid mutator transaction binding the contract method 0xcda6b847.
//
// Solidity: function stakeToken(uint256 amount) returns()
func (_Scattertoken *ScattertokenTransactorSession) StakeToken(amount *big.Int) (*types.Transaction, error) {
	return _Scattertoken.Contract.StakeToken(&_Scattertoken.TransactOpts, amount)
}

// TrainerLockToken is a paid mutator transaction binding the contract method 0xccd712d5.
//
// Solidity: function trainerLockToken(address trainerAddress, address requestorAddress, string topicName, uint256 amount) returns()
func (_Scattertoken *ScattertokenTransactor) TrainerLockToken(opts *bind.TransactOpts, trainerAddress common.Address, requestorAddress common.Address, topicName string, amount *big.Int) (*types.Transaction, error) {
	return _Scattertoken.contract.Transact(opts, "trainerLockToken", trainerAddress, requestorAddress, topicName, amount)
}

// TrainerLockToken is a paid mutator transaction binding the contract method 0xccd712d5.
//
// Solidity: function trainerLockToken(address trainerAddress, address requestorAddress, string topicName, uint256 amount) returns()
func (_Scattertoken *ScattertokenSession) TrainerLockToken(trainerAddress common.Address, requestorAddress common.Address, topicName string, amount *big.Int) (*types.Transaction, error) {
	return _Scattertoken.Contract.TrainerLockToken(&_Scattertoken.TransactOpts, trainerAddress, requestorAddress, topicName, amount)
}

// TrainerLockToken is a paid mutator transaction binding the contract method 0xccd712d5.
//
// Solidity: function trainerLockToken(address trainerAddress, address requestorAddress, string topicName, uint256 amount) returns()
func (_Scattertoken *ScattertokenTransactorSession) TrainerLockToken(trainerAddress common.Address, requestorAddress common.Address, topicName string, amount *big.Int) (*types.Transaction, error) {
	return _Scattertoken.Contract.TrainerLockToken(&_Scattertoken.TransactOpts, trainerAddress, requestorAddress, topicName, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Scattertoken *ScattertokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scattertoken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Scattertoken *ScattertokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scattertoken.Contract.Transfer(&_Scattertoken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Scattertoken *ScattertokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scattertoken.Contract.Transfer(&_Scattertoken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Scattertoken *ScattertokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scattertoken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Scattertoken *ScattertokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scattertoken.Contract.TransferFrom(&_Scattertoken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Scattertoken *ScattertokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scattertoken.Contract.TransferFrom(&_Scattertoken.TransactOpts, from, to, amount)
}

// ScattertokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Scattertoken contract.
type ScattertokenApprovalIterator struct {
	Event *ScattertokenApproval // Event containing the contract specifics and raw log

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
func (it *ScattertokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScattertokenApproval)
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
		it.Event = new(ScattertokenApproval)
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
func (it *ScattertokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScattertokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScattertokenApproval represents a Approval event raised by the Scattertoken contract.
type ScattertokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Scattertoken *ScattertokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ScattertokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Scattertoken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ScattertokenApprovalIterator{contract: _Scattertoken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Scattertoken *ScattertokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ScattertokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Scattertoken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScattertokenApproval)
				if err := _Scattertoken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Scattertoken *ScattertokenFilterer) ParseApproval(log types.Log) (*ScattertokenApproval, error) {
	event := new(ScattertokenApproval)
	if err := _Scattertoken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScattertokenTokensStakedIterator is returned from FilterTokensStaked and is used to iterate over the raw logs and unpacked data for TokensStaked events raised by the Scattertoken contract.
type ScattertokenTokensStakedIterator struct {
	Event *ScattertokenTokensStaked // Event containing the contract specifics and raw log

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
func (it *ScattertokenTokensStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScattertokenTokensStaked)
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
		it.Event = new(ScattertokenTokensStaked)
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
func (it *ScattertokenTokensStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScattertokenTokensStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScattertokenTokensStaked represents a TokensStaked event raised by the Scattertoken contract.
type ScattertokenTokensStaked struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensStaked is a free log retrieval operation binding the contract event 0xb539ca1e5c8d398ddf1c41c30166f33404941683be4683319b57669a93dad4ef.
//
// Solidity: event TokensStaked(address from, uint256 amount)
func (_Scattertoken *ScattertokenFilterer) FilterTokensStaked(opts *bind.FilterOpts) (*ScattertokenTokensStakedIterator, error) {

	logs, sub, err := _Scattertoken.contract.FilterLogs(opts, "TokensStaked")
	if err != nil {
		return nil, err
	}
	return &ScattertokenTokensStakedIterator{contract: _Scattertoken.contract, event: "TokensStaked", logs: logs, sub: sub}, nil
}

// WatchTokensStaked is a free log subscription operation binding the contract event 0xb539ca1e5c8d398ddf1c41c30166f33404941683be4683319b57669a93dad4ef.
//
// Solidity: event TokensStaked(address from, uint256 amount)
func (_Scattertoken *ScattertokenFilterer) WatchTokensStaked(opts *bind.WatchOpts, sink chan<- *ScattertokenTokensStaked) (event.Subscription, error) {

	logs, sub, err := _Scattertoken.contract.WatchLogs(opts, "TokensStaked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScattertokenTokensStaked)
				if err := _Scattertoken.contract.UnpackLog(event, "TokensStaked", log); err != nil {
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

// ParseTokensStaked is a log parse operation binding the contract event 0xb539ca1e5c8d398ddf1c41c30166f33404941683be4683319b57669a93dad4ef.
//
// Solidity: event TokensStaked(address from, uint256 amount)
func (_Scattertoken *ScattertokenFilterer) ParseTokensStaked(log types.Log) (*ScattertokenTokensStaked, error) {
	event := new(ScattertokenTokensStaked)
	if err := _Scattertoken.contract.UnpackLog(event, "TokensStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScattertokenTokensUnstakedIterator is returned from FilterTokensUnstaked and is used to iterate over the raw logs and unpacked data for TokensUnstaked events raised by the Scattertoken contract.
type ScattertokenTokensUnstakedIterator struct {
	Event *ScattertokenTokensUnstaked // Event containing the contract specifics and raw log

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
func (it *ScattertokenTokensUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScattertokenTokensUnstaked)
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
		it.Event = new(ScattertokenTokensUnstaked)
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
func (it *ScattertokenTokensUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScattertokenTokensUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScattertokenTokensUnstaked represents a TokensUnstaked event raised by the Scattertoken contract.
type ScattertokenTokensUnstaked struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensUnstaked is a free log retrieval operation binding the contract event 0x9845e367b683334e5c0b12d7b81721ac518e649376fa65e3d68324e8f34f2679.
//
// Solidity: event TokensUnstaked(address to, uint256 amount)
func (_Scattertoken *ScattertokenFilterer) FilterTokensUnstaked(opts *bind.FilterOpts) (*ScattertokenTokensUnstakedIterator, error) {

	logs, sub, err := _Scattertoken.contract.FilterLogs(opts, "TokensUnstaked")
	if err != nil {
		return nil, err
	}
	return &ScattertokenTokensUnstakedIterator{contract: _Scattertoken.contract, event: "TokensUnstaked", logs: logs, sub: sub}, nil
}

// WatchTokensUnstaked is a free log subscription operation binding the contract event 0x9845e367b683334e5c0b12d7b81721ac518e649376fa65e3d68324e8f34f2679.
//
// Solidity: event TokensUnstaked(address to, uint256 amount)
func (_Scattertoken *ScattertokenFilterer) WatchTokensUnstaked(opts *bind.WatchOpts, sink chan<- *ScattertokenTokensUnstaked) (event.Subscription, error) {

	logs, sub, err := _Scattertoken.contract.WatchLogs(opts, "TokensUnstaked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScattertokenTokensUnstaked)
				if err := _Scattertoken.contract.UnpackLog(event, "TokensUnstaked", log); err != nil {
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

// ParseTokensUnstaked is a log parse operation binding the contract event 0x9845e367b683334e5c0b12d7b81721ac518e649376fa65e3d68324e8f34f2679.
//
// Solidity: event TokensUnstaked(address to, uint256 amount)
func (_Scattertoken *ScattertokenFilterer) ParseTokensUnstaked(log types.Log) (*ScattertokenTokensUnstaked, error) {
	event := new(ScattertokenTokensUnstaked)
	if err := _Scattertoken.contract.UnpackLog(event, "TokensUnstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScattertokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Scattertoken contract.
type ScattertokenTransferIterator struct {
	Event *ScattertokenTransfer // Event containing the contract specifics and raw log

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
func (it *ScattertokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScattertokenTransfer)
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
		it.Event = new(ScattertokenTransfer)
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
func (it *ScattertokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScattertokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScattertokenTransfer represents a Transfer event raised by the Scattertoken contract.
type ScattertokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Scattertoken *ScattertokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ScattertokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Scattertoken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ScattertokenTransferIterator{contract: _Scattertoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Scattertoken *ScattertokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ScattertokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Scattertoken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScattertokenTransfer)
				if err := _Scattertoken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Scattertoken *ScattertokenFilterer) ParseTransfer(log types.Log) (*ScattertokenTransfer, error) {
	event := new(ScattertokenTransfer)
	if err := _Scattertoken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
