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
)

// ScattertokenMetaData contains all meta data concerning the Scattertoken contract.
var ScattertokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensUnstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addressToStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"canBecomeValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"}],\"name\":\"donateToLottery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwnStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"trainers\",\"type\":\"address[]\"}],\"name\":\"punishRogueTrainers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"punishRogueValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"removeStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"requestorLockToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requiredModelValidatorStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"trainers\",\"type\":\"address[]\"}],\"name\":\"rewardBenevolentTrainers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"rewardBenevolentValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scatterProtocolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"setScatterProtocolAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stakeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"trainerLockToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040526000600b556002600c556064600d556032600e556032600f553480156200002a57600080fd5b5060405162002456380380620024568339810160408190526200004d916200036e565b6200005b6012600a6200049d565b620000679082620004b5565b6040518060400160405280600c81526020016b29b1b0ba3a32b92a37b5b2b760a11b8152506040518060400160405280600281526020016114d560f21b8152508160039081620000b8919062000573565b506004620000c7828262000573565b50505060008111620001205760405162461bcd60e51b815260206004820152601560248201527f45524332304361707065643a206361702069732030000000000000000000000060448201526064015b60405180910390fd5b608052600680546001600160a01b03191633908117909155620001659062000146601290565b6200015390600a6200049d565b6200015f9084620004b5565b6200017f565b506127106005556010805460ff60a01b1916905562000655565b60805181620001986200020f60201b620005561760201c565b620001a491906200063f565b1115620001f45760405162461bcd60e51b815260206004820152601960248201527f45524332304361707065643a2063617020657863656564656400000000000000604482015260640162000117565b6200020b82826200021560201b620010ff1760201c565b5050565b60025490565b608051816200022e6200020f60201b620005561760201c565b6200023a91906200063f565b11156200028a5760405162461bcd60e51b815260206004820152601960248201527f45524332304361707065643a2063617020657863656564656400000000000000604482015260640162000117565b6200020b8282620002a160201b620011841760201c565b6001600160a01b038216620002f95760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640162000117565b80600260008282546200030d91906200063f565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a36200020b600083835b505050565b6000602082840312156200038157600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b600181815b80851115620003df578160001904821115620003c357620003c362000388565b80851615620003d157918102915b93841c9390800290620003a3565b509250929050565b600082620003f85750600162000497565b81620004075750600062000497565b81600181146200042057600281146200042b576200044b565b600191505062000497565b60ff8411156200043f576200043f62000388565b50506001821b62000497565b5060208310610133831016604e8410600b841016171562000470575081810a62000497565b6200047c83836200039e565b806000190482111562000493576200049362000388565b0290505b92915050565b6000620004ae60ff841683620003e7565b9392505050565b808202811582820484141762000497576200049762000388565b634e487b7160e01b600052604160045260246000fd5b600181811c90821680620004fa57607f821691505b6020821081036200051b57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200036957600081815260208120601f850160051c810160208610156200054a5750805b601f850160051c820191505b818110156200056b5782815560010162000556565b505050505050565b81516001600160401b038111156200058f576200058f620004cf565b620005a781620005a08454620004e5565b8462000521565b602080601f831160018114620005df5760008415620005c65750858301515b600019600386901b1c1916600185901b1785556200056b565b600085815260208120601f198616915b828110156200061057888601518255948401946001909101908401620005ef565b50858210156200062f5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b8082018082111562000497576200049762000388565b608051611dd76200067f600039600081816102610152818161110101526116b30152611dd76000f3fe608060405234801561001057600080fd5b50600436106101e55760003560e01c806370a082311161010f578063a1a33f44116100a2578063ccd712d511610071578063ccd712d51461045c578063cda6b8471461046f578063dd62ed3e14610482578063fb7ba5821461049557600080fd5b8063a1a33f4414610403578063a457c2d714610423578063a9059cbb14610436578063bea55db91461044957600080fd5b8063939624ab116100de578063939624ab146103a557806395d89b41146103b8578063985fb65d146103c05780639dacf42d146103d357600080fd5b806370a082311461032b57806379cc679014610354578063824205d1146103675780638da5cb5b1461037a57600080fd5b8063395093511161018757806343f893511161015657806343f89351146102dc578063454cbd3e146102ef57806347dc19c7146102ef5780635e30913f1461030257600080fd5b8063395093511461028e578063408111f5146102a157806342966c68146102b657806342a0f9bf146102c957600080fd5b806323b872dd116101c357806323b872dd1461023d578063313ce56714610250578063355274ea1461025f57806336cb7f801461028557600080fd5b806306fdde03146101ea578063095ea7b31461020857806318160ddd1461022b575b600080fd5b6101f26104aa565b6040516101ff919061175e565b60405180910390f35b61021b6102163660046117ad565b61053c565b60405190151581526020016101ff565b6002545b6040519081526020016101ff565b61021b61024b3660046117d7565b61055c565b604051601281526020016101ff565b7f000000000000000000000000000000000000000000000000000000000000000061022f565b61022f60055481565b61021b61029c3660046117ad565b610580565b6102b46102af366004611951565b6105a2565b005b6102b46102c43660046119c5565b6106f6565b6102b46102d73660046119de565b610703565b6102b46102ea366004611a35565b6107fd565b6102b46102fd366004611951565b610914565b61022f610310366004611a72565b6001600160a01b031660009081526007602052604090205490565b61022f610339366004611a72565b6001600160a01b031660009081526020819052604090205490565b6102b46103623660046117ad565b610943565b6102b4610375366004611a94565b610958565b60065461038d906001600160a01b031681565b6040516001600160a01b0390911681526020016101ff565b6102b46103b33660046119c5565b610a40565b6101f2610cda565b60105461038d906001600160a01b031681565b61021b6103e1366004611a72565b6005546001600160a01b03909116600090815260076020526040902054101590565b61022f610411366004611a72565b60076020526000908152604090205481565b61021b6104313660046117ad565b610ce9565b61021b6104443660046117ad565b610d64565b6102b4610457366004611a72565b610d72565b6102b461046a366004611ae2565b610dfc565b6102b461047d3660046119c5565b610f21565b61022f610490366004611b48565b6110d4565b3360009081526007602052604090205461022f565b6060600380546104b990611b7b565b80601f01602080910402602001604051908101604052809291908181526020018280546104e590611b7b565b80156105325780601f1061050757610100808354040283529160200191610532565b820191906000526020600020905b81548152906001019060200180831161051557829003601f168201915b5050505050905090565b60003361054a818585611243565b60019150505b92915050565b60025490565b60003361056a858285611367565b6105758585856113db565b506001949350505050565b60003361054a81858561059383836110d4565b61059d9190611bcb565b611243565b6010546001600160a01b031633146105d55760405162461bcd60e51b81526004016105cc90611bde565b60405180910390fd5b60005b81518110156106f0576001600160a01b038416600090815260096020526040908190209051610608908590611c3b565b9081526020016040518091039020600083838151811061062a5761062a611c57565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054600b60008282546106649190611bcb565b90915550506001600160a01b038416600090815260096020526040808220905161068f908690611c3b565b908152602001604051809103902060008484815181106106b1576106b1611c57565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000208190555080806106e890611c6d565b9150506105d8565b50505050565b610700338261157f565b50565b6010546001600160a01b0316331461072d5760405162461bcd60e51b81526004016105cc90611bde565b6040516370a0823160e01b81526001600160a01b038416600482015230906370a0823190602401602060405180830381865afa158015610771573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107959190611c86565b8111156107b45760405162461bcd60e51b81526004016105cc90611c9f565b6107be838261157f565b6001600160a01b0383166000908152600a60205260409081902090518291906107e8908590611c3b565b90815260405190819003602001902055505050565b6010546001600160a01b031633146108275760405162461bcd60e51b81526004016105cc90611bde565b60005b81518110156109105760006064600d546007600086868151811061085057610850611c57565b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020546108839190611ce3565b61088d9190611cfa565b905080600b60008282546108a19190611bcb565b9250508190555080600760008585815181106108bf576108bf611c57565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060008282546108f69190611d1c565b90915550829150610908905081611c6d565b91505061082a565b5050565b6010546001600160a01b0316331461093e5760405162461bcd60e51b81526004016105cc90611bde565b505050565b61094e823383611367565b610910828261157f565b6010546001600160a01b031633146109825760405162461bcd60e51b81526004016105cc90611bde565b6001600160a01b0382166000908152600a602052604080822090516109a8908490611c3b565b908152602001604051809103902054905060006064600c54836109cb9190611ce3565b6109d59190611cfa565b905080600b60008282546109e99190611bcb565b90915550506001600160a01b0384166000908152600a6020526040908190209051829190610a18908690611c3b565b90815260200160405180910390206000828254610a359190611d1c565b909155505050505050565b601054600160a01b900460ff1615610a8b5760405162461bcd60e51b815260206004820152600e60248201526d4e6f2072652d656e7472616e637960901b60448201526064016105cc565b6010805460ff60a01b1916600160a01b17905533600090815260086020526040902054421115610b235760405162461bcd60e51b815260206004820152603860248201527f596f75206d7573742077616974206265666f7265206265696e672061626c652060448201527f746f20756e7374616b65205363617474657220546f6b656e000000000000000060648201526084016105cc565b336000908152600760205260409020548110610b905760405162461bcd60e51b815260206004820152602660248201527f416d6f756e74206d757374206265206c657373207468616e207374616b656420604482015265185b5bdd5b9d60d21b60648201526084016105cc565b610b9a33826116b1565b3360009081526007602052604081208054839290610bb9908490611d1c565b9091555050601054604051630a9d3d9b60e11b81523360048201526000916001600160a01b03169063153a7b3690602401602060405180830381865afa158015610c07573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c2b9190611d2f565b90506003816004811115610c4157610c41611d50565b148015610c5e575060055433600090815260076020526040902054105b15610cc9576010546040516302b8e1f360e51b81526001600160a01b039091169063571c3e6090610c96903390600090600401611d66565b600060405180830381600087803b158015610cb057600080fd5b505af1158015610cc4573d6000803e3d6000fd5b505050505b50506010805460ff60a01b19169055565b6060600480546104b990611b7b565b60003381610cf782866110d4565b905083811015610d575760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084016105cc565b6105758286868403611243565b60003361054a8185856113db565b6006546001600160a01b03163314610dda5760405162461bcd60e51b815260206004820152602560248201527f4f6e6c7920746865206f776e65722063616e2063616c6c20746869732066756e60448201526431ba34b7b760d91b60648201526084016105cc565b601080546001600160a01b0319166001600160a01b0392909216919091179055565b6010546001600160a01b03163314610e265760405162461bcd60e51b81526004016105cc90611bde565b6040516370a0823160e01b81526001600160a01b038516600482015230906370a0823190602401602060405180830381865afa158015610e6a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e8e9190611c86565b811115610ead5760405162461bcd60e51b81526004016105cc90611c9f565b610eb7848261157f565b6001600160a01b038316600090815260096020526040908190209051829190610ee1908590611c3b565b90815260200160405180910390206000866001600160a01b03166001600160a01b031681526020019081526020016000206000828254610a359190611bcb565b601054600160a01b900460ff1615610f6c5760405162461bcd60e51b815260206004820152600e60248201526d4e6f2072652d656e7472616e637960901b60448201526064016105cc565b6010805460ff60a01b1916600160a01b1790556040516370a0823160e01b815233600482015230906370a0823190602401602060405180830381865afa158015610fba573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fde9190611c86565b81111561103b5760405162461bcd60e51b815260206004820152602560248201527f43616e6e6f74207374616b65206d6f726520746f6b656e73207468616e20796f6044820152643a9037bbb760d91b60648201526084016105cc565b611045338261157f565b3360009081526007602052604081208054839290611064908490611bcb565b9091555061107790504262278d00611bcb565b336000818152600860209081526040918290209390935580519182529181018390527fb539ca1e5c8d398ddf1c41c30166f33404941683be4683319b57669a93dad4ef910160405180910390a1506010805460ff60a01b19169055565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b7f00000000000000000000000000000000000000000000000000000000000000008161112a60025490565b6111349190611bcb565b111561117e5760405162461bcd60e51b8152602060048201526019602482015278115490cc8c10d85c1c19590e8818d85c08195e18d959591959603a1b60448201526064016105cc565b61091082825b6001600160a01b0382166111da5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016105cc565b80600260008282546111ec9190611bcb565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b6001600160a01b0383166112a55760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084016105cc565b6001600160a01b0382166113065760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016105cc565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b600061137384846110d4565b905060001981146106f057818110156113ce5760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016105cc565b6106f08484848403611243565b6001600160a01b03831661143f5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016105cc565b6001600160a01b0382166114a15760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016105cc565b6001600160a01b038316600090815260208190526040902054818110156115195760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b60648201526084016105cc565b6001600160a01b03848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a36106f0565b6001600160a01b0382166115df5760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b60648201526084016105cc565b6001600160a01b038216600090815260208190526040902054818110156116535760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b60648201526084016105cc565b6001600160a01b0383166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3505050565b7f0000000000000000000000000000000000000000000000000000000000000000816116dc60025490565b6116e69190611bcb565b11156117305760405162461bcd60e51b8152602060048201526019602482015278115490cc8c10d85c1c19590e8818d85c08195e18d959591959603a1b60448201526064016105cc565b61091082826110ff565b60005b8381101561175557818101518382015260200161173d565b50506000910152565b602081526000825180602084015261177d81604085016020870161173a565b601f01601f19169190910160400192915050565b80356001600160a01b03811681146117a857600080fd5b919050565b600080604083850312156117c057600080fd5b6117c983611791565b946020939093013593505050565b6000806000606084860312156117ec57600080fd5b6117f584611791565b925061180360208501611791565b9150604084013590509250925092565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561185257611852611813565b604052919050565b600082601f83011261186b57600080fd5b813567ffffffffffffffff81111561188557611885611813565b611898601f8201601f1916602001611829565b8181528460208386010111156118ad57600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f8301126118db57600080fd5b8135602067ffffffffffffffff8211156118f7576118f7611813565b8160051b611906828201611829565b928352848101820192828101908785111561192057600080fd5b83870192505b848310156119465761193783611791565b82529183019190830190611926565b979650505050505050565b60008060006060848603121561196657600080fd5b61196f84611791565b9250602084013567ffffffffffffffff8082111561198c57600080fd5b6119988783880161185a565b935060408601359150808211156119ae57600080fd5b506119bb868287016118ca565b9150509250925092565b6000602082840312156119d757600080fd5b5035919050565b6000806000606084860312156119f357600080fd5b6119fc84611791565b9250602084013567ffffffffffffffff811115611a1857600080fd5b611a248682870161185a565b925050604084013590509250925092565b600060208284031215611a4757600080fd5b813567ffffffffffffffff811115611a5e57600080fd5b611a6a848285016118ca565b949350505050565b600060208284031215611a8457600080fd5b611a8d82611791565b9392505050565b60008060408385031215611aa757600080fd5b611ab083611791565b9150602083013567ffffffffffffffff811115611acc57600080fd5b611ad88582860161185a565b9150509250929050565b60008060008060808587031215611af857600080fd5b611b0185611791565b9350611b0f60208601611791565b9250604085013567ffffffffffffffff811115611b2b57600080fd5b611b378782880161185a565b949793965093946060013593505050565b60008060408385031215611b5b57600080fd5b611b6483611791565b9150611b7260208401611791565b90509250929050565b600181811c90821680611b8f57607f821691505b602082108103611baf57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b8082018082111561055057610550611bb5565b6020808252603f908201527f54686973206d6574686f642063616e206f6e6c792062652063616c6c6564206260408201527f792074686520736361747465722070726f746f636f6c20636f6e747261637400606082015260800190565b60008251611c4d81846020870161173a565b9190910192915050565b634e487b7160e01b600052603260045260246000fd5b600060018201611c7f57611c7f611bb5565b5060010190565b600060208284031215611c9857600080fd5b5051919050565b60208082526024908201527f43616e6e6f74206c6f636b206d6f726520746f6b656e73207468616e20796f756040820152631037bbb760e11b606082015260800190565b808202811582820484141761055057610550611bb5565b600082611d1757634e487b7160e01b600052601260045260246000fd5b500490565b8181038181111561055057610550611bb5565b600060208284031215611d4157600080fd5b815160058110611a8d57600080fd5b634e487b7160e01b600052602160045260246000fd5b6001600160a01b03831681526040810160058310611d9457634e487b7160e01b600052602160045260246000fd5b826020830152939250505056fea2646970667358221220dbf777c1c327bae98480764f3e0e76908317062a398166e5ab7b279f4208184764736f6c63430008110033",
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
	parsed, err := abi.JSON(strings.NewReader(ScattertokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
