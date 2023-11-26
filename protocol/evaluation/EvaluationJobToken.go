// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package evaluationtoken

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

// EvaluationtokenMetaData contains all meta data concerning the Evaluationtoken contract.
var EvaluationtokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getTokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"isEvaluationScoreSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolDeployer\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dataPublisher\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"jobURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"evaluationDataURI\",\"type\":\"string\"}],\"name\":\"publishEvaluationData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"publishEvaluationJob\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scatterProtocolContract\",\"outputs\":[{\"internalType\":\"contractIScatterProtocol\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setScatterContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"score\",\"type\":\"uint256\"}],\"name\":\"submitEvaluationScore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060408051808201825260208082527f536361747465722050726f746f636f6c204576616c756174696f6e204a6f6273818301528251808401909352600483526329a822a560e11b908301529060006200006c8382620001ab565b5060016200007b8282620001ab565b5050506200009862000092620000b060201b60201c565b620000b4565b600c80546001600160a01b0319163317905562000277565b3390565b600780546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806200013157607f821691505b6020821081036200015257634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620001a657600081815260208120601f850160051c81016020861015620001815750805b601f850160051c820191505b81811015620001a2578281556001016200018d565b5050505b505050565b81516001600160401b03811115620001c757620001c762000106565b620001df81620001d884546200011c565b8462000158565b602080601f831160018114620002175760008415620001fe5750858301515b600019600386901b1c1916600185901b178555620001a2565b600085815260208120601f198616915b82811015620002485788860151825594840194600190910190840162000227565b5085821015620002675787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6121f880620002876000396000f3fe608060405234801561001057600080fd5b50600436106101735760003560e01c806351bd2421116100de5780638da5cb5b11610097578063b88d4fde11610071578063b88d4fde1461032e578063c87b56dd14610341578063e985e9c514610354578063f2fde38b1461039057600080fd5b80638da5cb5b1461030257806395d89b4114610313578063a22cb4651461031b57600080fd5b806351bd24211461029b57806355f804b3146102ae578063572d107b146102c15780636352211e146102d457806370a08231146102e7578063715018a6146102fa57600080fd5b806323b872dd1161013057806323b872dd1461021b578063330c093c1461022e5780633bb3a24d146102415780633c7645b71461025457806342842e0e1461026757806349db9dc31461027a57600080fd5b806301ffc9a71461017857806306fdde03146101a0578063081812fc146101b5578063095ea7b3146101e0578063126d5b6a146101f55780631e2af3cb14610208575b600080fd5b61018b610186366004611919565b6103a3565b60405190151581526020015b60405180910390f35b6101a86103ce565b6040516101979190611986565b6101c86101c3366004611999565b610460565b6040516001600160a01b039091168152602001610197565b6101f36101ee3660046119ce565b610487565b005b6101f36102033660046119f8565b6105a1565b61018b610216366004611abf565b6105cb565b6101f3610229366004611b2e565b610631565b600c546101c8906001600160a01b031681565b6101a861024f366004611999565b610662565b6101f3610262366004611b6a565b61066d565b6101f3610275366004611b2e565b610826565b61028d610288366004611bde565b610841565b604051908152602001610197565b6101f36102a9366004611c2c565b6108e5565b6101f36102bc366004611ca3565b610b99565b600d546101c8906001600160a01b031681565b6101c86102e2366004611999565b610c11565b61028d6102f53660046119f8565b610c71565b6101f3610cf7565b6007546001600160a01b03166101c8565b6101a8610d0b565b6101f3610329366004611ce6565b610d1a565b6101f361033c366004611d1d565b610d25565b6101a861034f366004611999565b610d57565b61018b610362366004611d99565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b6101f361039e3660046119f8565b610e52565b60006001600160e01b03198216632483248360e11b14806103c857506103c882610ecb565b92915050565b6060600080546103dd90611dcc565b80601f016020809104026020016040519081016040528092919081815260200182805461040990611dcc565b80156104565780601f1061042b57610100808354040283529160200191610456565b820191906000526020600020905b81548152906001019060200180831161043957829003601f168201915b5050505050905090565b600061046b82610f1b565b506000908152600460205260409020546001600160a01b031690565b600061049282610c11565b9050806001600160a01b0316836001600160a01b0316036105045760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b60648201526084015b60405180910390fd5b336001600160a01b038216148061052057506105208133610362565b6105925760405162461bcd60e51b815260206004820152603d60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206f7220617070726f76656420666f7220616c6c00000060648201526084016104fb565b61059c8383610f7a565b505050565b6105a9610fe8565b600d80546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b03841660009081526010602052604080822090516105f1908690611e06565b90815260408051602092819003830190206001600160a01b038681166000908152918452828220908616825290925290205460ff1690505b949350505050565b61063b3382611042565b6106575760405162461bcd60e51b81526004016104fb90611e22565b61059c8383836110c0565b60606103c882610d57565b600d546001600160a01b031633146106975760405162461bcd60e51b81526004016104fb90611e6f565b826001600160a01b03166009836040516106b19190611e06565b908152604051908190036020019020546001600160a01b03161461073d5760405162461bcd60e51b815260206004820152603960248201527f54686973206576616c756174696f6e206a6f6220646f6573206e6f742062656c60448201527f6f6e6720746f207468652064617461207075626c69736865720000000000000060648201526084016104fb565b6040805160208101825260009052517fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090600a9061077c908590611e06565b90815260405190819003602001812061079491611ecc565b6040518091039020146107f55760405162461bcd60e51b8152602060048201526024808201527f54686973206a6f6220616c7265616479206861732061206d617070696e6720746044820152631bc81a5d60e21b60648201526084016104fb565b80600a836040516108069190611e06565b908152602001604051809103902090816108209190611f90565b50505050565b61059c83838360405180602001604052806000815250610d25565b600d546000906001600160a01b0316331461086e5760405162461bcd60e51b81526004016104fb90611e6f565b61087c600e80546001019055565b6000610887600e5490565b90506108938482611224565b61089d818461123e565b836009846040516108ae9190611e06565b90815260405190819003602001902080546001600160a01b03929092166001600160a01b0319909216919091179055905092915050565b600d546001600160a01b0316331461090f5760405162461bcd60e51b81526004016104fb90611e6f565b6001600160a01b038516600090815260106020526040908190209051610936908690611e06565b90815260408051602092819003830190206001600160a01b038581166000908152918452828220908716825290925290205460ff16156109ef5760405162461bcd60e51b815260206004820152604860248201527f43616e6e6f742072657375626d697420612073636f726520666f72206120747260448201527f61696e65722074686174206861732070726576696f75736c79206265656e20736064820152671d589b5a5d1d195960c21b608482015260a4016104fb565b600d546040516303a120eb60e51b81526001600160a01b03909116906374241d6090610a2390889088908790600401612050565b602060405180830381865afa158015610a40573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a649190612085565b610ad65760405162461bcd60e51b815260206004820152603760248201527f596f7520617265206e6f7420616e2061737369676e65642076616c696461746f60448201527f7220666f72207468697320747261696e696e67206a6f6200000000000000000060648201526084016104fb565b6001600160a01b0385166000908152600f6020526040908190209051829190610b00908790611e06565b9081526040805191829003602090810183206001600160a01b0380881660009081529183528382208982168352835283822095909555938916845260109052909120600191610b50908790611e06565b90815260408051602092819003830190206001600160a01b03958616600090815290835281812096909516855294905292909120805460ff191692151592909217909155505050565b600c546001600160a01b03163314610c015760405162461bcd60e51b815260206004820152602560248201527f4f6e6c7920746865206f776e65722063616e206163636573732074686973206d604482015264195d1a1bd960da1b60648201526084016104fb565b600b610c0d8282611f90565b5050565b6000818152600260205260408120546001600160a01b0316806103c85760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b60448201526064016104fb565b60006001600160a01b038216610cdb5760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b60648201526084016104fb565b506001600160a01b031660009081526003602052604090205490565b610cff610fe8565b610d096000611248565b565b6060600180546103dd90611dcc565b610c0d33838361129a565b610d2f3383611042565b610d4b5760405162461bcd60e51b81526004016104fb90611e22565b61082084848484611368565b6060610d6282610f1b565b60008281526006602052604081208054610d7b90611dcc565b80601f0160208091040260200160405190810160405280929190818152602001828054610da790611dcc565b8015610df45780601f10610dc957610100808354040283529160200191610df4565b820191906000526020600020905b815481529060010190602001808311610dd757829003601f168201915b505050505090506000610e0561139b565b90508051600003610e17575092915050565b815115610e49578082604051602001610e319291906120a2565b60405160208183030381529060405292505050919050565b610629846113aa565b610e5a610fe8565b6001600160a01b038116610ebf5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016104fb565b610ec881611248565b50565b60006001600160e01b031982166380ac58cd60e01b1480610efc57506001600160e01b03198216635b5e139f60e01b145b806103c857506301ffc9a760e01b6001600160e01b03198316146103c8565b6000818152600260205260409020546001600160a01b0316610ec85760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b60448201526064016104fb565b600081815260046020526040902080546001600160a01b0319166001600160a01b0384169081179091558190610faf82610c11565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6007546001600160a01b03163314610d095760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104fb565b60008061104e83610c11565b9050806001600160a01b0316846001600160a01b0316148061109557506001600160a01b0380821660009081526005602090815260408083209388168352929052205460ff165b806106295750836001600160a01b03166110ae84610460565b6001600160a01b031614949350505050565b826001600160a01b03166110d382610c11565b6001600160a01b0316146110f95760405162461bcd60e51b81526004016104fb906120d1565b6001600160a01b03821661115b5760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b60648201526084016104fb565b826001600160a01b031661116e82610c11565b6001600160a01b0316146111945760405162461bcd60e51b81526004016104fb906120d1565b600081815260046020908152604080832080546001600160a01b03199081169091556001600160a01b0387811680865260038552838620805460001901905590871680865283862080546001019055868652600290945282852080549092168417909155905184937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b610c0d828260405180602001604052806000815250611411565b610c0d8282611444565b600780546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b816001600160a01b0316836001600160a01b0316036112fb5760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c65720000000000000060448201526064016104fb565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6113738484846110c0565b61137f8484848461150f565b6108205760405162461bcd60e51b81526004016104fb90612116565b6060600b80546103dd90611dcc565b60606113b582610f1b565b60006113bf61139b565b905060008151116113df576040518060200160405280600081525061140a565b806113e98461160d565b6040516020016113fa9291906120a2565b6040516020818303038152906040525b9392505050565b61141b83836116a0565b611428600084848461150f565b61059c5760405162461bcd60e51b81526004016104fb90612116565b6000828152600260205260409020546001600160a01b03166114bf5760405162461bcd60e51b815260206004820152602e60248201527f45524337323155524953746f726167653a2055524920736574206f66206e6f6e60448201526d32bc34b9ba32b73a103a37b5b2b760911b60648201526084016104fb565b60008281526006602052604090206114d78282611f90565b506040518281527ff8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce79060200160405180910390a15050565b60006001600160a01b0384163b1561160557604051630a85bd0160e11b81526001600160a01b0385169063150b7a0290611553903390899088908890600401612168565b6020604051808303816000875af192505050801561158e575060408051601f3d908101601f1916820190925261158b918101906121a5565b60015b6115eb573d8080156115bc576040519150601f19603f3d011682016040523d82523d6000602084013e6115c1565b606091505b5080516000036115e35760405162461bcd60e51b81526004016104fb90612116565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050610629565b506001610629565b6060600061161a8361182b565b600101905060008167ffffffffffffffff81111561163a5761163a611a13565b6040519080825280601f01601f191660200182016040528015611664576020820181803683370190505b5090508181016020015b600019016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a850494508461166e57509392505050565b6001600160a01b0382166116f65760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f206164647265737360448201526064016104fb565b6000818152600260205260409020546001600160a01b03161561175b5760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e7465640000000060448201526064016104fb565b6000818152600260205260409020546001600160a01b0316156117c05760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e7465640000000060448201526064016104fb565b6001600160a01b038216600081815260036020908152604080832080546001019055848352600290915280822080546001600160a01b0319168417905551839291907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b60008072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b831061186a5772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef81000000008310611896576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc1000083106118b457662386f26fc10000830492506010015b6305f5e10083106118cc576305f5e100830492506008015b61271083106118e057612710830492506004015b606483106118f2576064830492506002015b600a83106103c85760010192915050565b6001600160e01b031981168114610ec857600080fd5b60006020828403121561192b57600080fd5b813561140a81611903565b60005b83811015611951578181015183820152602001611939565b50506000910152565b60008151808452611972816020860160208601611936565b601f01601f19169290920160200192915050565b60208152600061140a602083018461195a565b6000602082840312156119ab57600080fd5b5035919050565b80356001600160a01b03811681146119c957600080fd5b919050565b600080604083850312156119e157600080fd5b6119ea836119b2565b946020939093013593505050565b600060208284031215611a0a57600080fd5b61140a826119b2565b634e487b7160e01b600052604160045260246000fd5b600067ffffffffffffffff80841115611a4457611a44611a13565b604051601f8501601f19908116603f01168101908282118183101715611a6c57611a6c611a13565b81604052809350858152868686011115611a8557600080fd5b858560208301376000602087830101525050509392505050565b600082601f830112611ab057600080fd5b61140a83833560208501611a29565b60008060008060808587031215611ad557600080fd5b611ade856119b2565b9350602085013567ffffffffffffffff811115611afa57600080fd5b611b0687828801611a9f565b935050611b15604086016119b2565b9150611b23606086016119b2565b905092959194509250565b600080600060608486031215611b4357600080fd5b611b4c846119b2565b9250611b5a602085016119b2565b9150604084013590509250925092565b600080600060608486031215611b7f57600080fd5b611b88846119b2565b9250602084013567ffffffffffffffff80821115611ba557600080fd5b611bb187838801611a9f565b93506040860135915080821115611bc757600080fd5b50611bd486828701611a9f565b9150509250925092565b60008060408385031215611bf157600080fd5b611bfa836119b2565b9150602083013567ffffffffffffffff811115611c1657600080fd5b611c2285828601611a9f565b9150509250929050565b600080600080600060a08688031215611c4457600080fd5b611c4d866119b2565b9450602086013567ffffffffffffffff811115611c6957600080fd5b611c7588828901611a9f565b945050611c84604087016119b2565b9250611c92606087016119b2565b949793965091946080013592915050565b600060208284031215611cb557600080fd5b813567ffffffffffffffff811115611ccc57600080fd5b61062984828501611a9f565b8015158114610ec857600080fd5b60008060408385031215611cf957600080fd5b611d02836119b2565b91506020830135611d1281611cd8565b809150509250929050565b60008060008060808587031215611d3357600080fd5b611d3c856119b2565b9350611d4a602086016119b2565b925060408501359150606085013567ffffffffffffffff811115611d6d57600080fd5b8501601f81018713611d7e57600080fd5b611d8d87823560208401611a29565b91505092959194509250565b60008060408385031215611dac57600080fd5b611db5836119b2565b9150611dc3602084016119b2565b90509250929050565b600181811c90821680611de057607f821691505b602082108103611e0057634e487b7160e01b600052602260045260246000fd5b50919050565b60008251611e18818460208701611936565b9190910192915050565b6020808252602d908201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560408201526c1c881bdc88185c1c1c9bdd9959609a1b606082015260800190565b6020808252603f908201527f54686973206d6574686f642063616e206f6e6c792062652063616c6c6564206260408201527f792074686520736361747465722070726f746f636f6c20636f6e747261637400606082015260800190565b6000808354611eda81611dcc565b60018281168015611ef25760018114611f0757611f36565b60ff1984168752821515830287019450611f36565b8760005260208060002060005b85811015611f2d5781548a820152908401908201611f14565b50505082870194505b50929695505050505050565b601f82111561059c57600081815260208120601f850160051c81016020861015611f695750805b601f850160051c820191505b81811015611f8857828155600101611f75565b505050505050565b815167ffffffffffffffff811115611faa57611faa611a13565b611fbe81611fb88454611dcc565b84611f42565b602080601f831160018114611ff35760008415611fdb5750858301515b600019600386901b1c1916600185901b178555611f88565b600085815260208120601f198616915b8281101561202257888601518255948401946001909101908401612003565b50858210156120405787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b600060018060a01b03808616835260606020840152612072606084018661195a565b9150808416604084015250949350505050565b60006020828403121561209757600080fd5b815161140a81611cd8565b600083516120b4818460208801611936565b8351908301906120c8818360208801611936565b01949350505050565b60208082526025908201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060408201526437bbb732b960d91b606082015260800190565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b6001600160a01b038581168252841660208201526040810183905260806060820181905260009061219b9083018461195a565b9695505050505050565b6000602082840312156121b757600080fd5b815161140a8161190356fea2646970667358221220c2def1d92ffa21f2be3f1a4837bcc9131e6f39cf9948ea52961518e9640cd3c864736f6c63430008110033",
}

// EvaluationtokenABI is the input ABI used to generate the binding from.
// Deprecated: Use EvaluationtokenMetaData.ABI instead.
var EvaluationtokenABI = EvaluationtokenMetaData.ABI

// EvaluationtokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EvaluationtokenMetaData.Bin instead.
var EvaluationtokenBin = EvaluationtokenMetaData.Bin

// DeployEvaluationtoken deploys a new Ethereum contract, binding an instance of Evaluationtoken to it.
func DeployEvaluationtoken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Evaluationtoken, error) {
	parsed, err := EvaluationtokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EvaluationtokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Evaluationtoken{EvaluationtokenCaller: EvaluationtokenCaller{contract: contract}, EvaluationtokenTransactor: EvaluationtokenTransactor{contract: contract}, EvaluationtokenFilterer: EvaluationtokenFilterer{contract: contract}}, nil
}

// Evaluationtoken is an auto generated Go binding around an Ethereum contract.
type Evaluationtoken struct {
	EvaluationtokenCaller     // Read-only binding to the contract
	EvaluationtokenTransactor // Write-only binding to the contract
	EvaluationtokenFilterer   // Log filterer for contract events
}

// EvaluationtokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type EvaluationtokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EvaluationtokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EvaluationtokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EvaluationtokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EvaluationtokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EvaluationtokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EvaluationtokenSession struct {
	Contract     *Evaluationtoken  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EvaluationtokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EvaluationtokenCallerSession struct {
	Contract *EvaluationtokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// EvaluationtokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EvaluationtokenTransactorSession struct {
	Contract     *EvaluationtokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// EvaluationtokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type EvaluationtokenRaw struct {
	Contract *Evaluationtoken // Generic contract binding to access the raw methods on
}

// EvaluationtokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EvaluationtokenCallerRaw struct {
	Contract *EvaluationtokenCaller // Generic read-only contract binding to access the raw methods on
}

// EvaluationtokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EvaluationtokenTransactorRaw struct {
	Contract *EvaluationtokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEvaluationtoken creates a new instance of Evaluationtoken, bound to a specific deployed contract.
func NewEvaluationtoken(address common.Address, backend bind.ContractBackend) (*Evaluationtoken, error) {
	contract, err := bindEvaluationtoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Evaluationtoken{EvaluationtokenCaller: EvaluationtokenCaller{contract: contract}, EvaluationtokenTransactor: EvaluationtokenTransactor{contract: contract}, EvaluationtokenFilterer: EvaluationtokenFilterer{contract: contract}}, nil
}

// NewEvaluationtokenCaller creates a new read-only instance of Evaluationtoken, bound to a specific deployed contract.
func NewEvaluationtokenCaller(address common.Address, caller bind.ContractCaller) (*EvaluationtokenCaller, error) {
	contract, err := bindEvaluationtoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EvaluationtokenCaller{contract: contract}, nil
}

// NewEvaluationtokenTransactor creates a new write-only instance of Evaluationtoken, bound to a specific deployed contract.
func NewEvaluationtokenTransactor(address common.Address, transactor bind.ContractTransactor) (*EvaluationtokenTransactor, error) {
	contract, err := bindEvaluationtoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EvaluationtokenTransactor{contract: contract}, nil
}

// NewEvaluationtokenFilterer creates a new log filterer instance of Evaluationtoken, bound to a specific deployed contract.
func NewEvaluationtokenFilterer(address common.Address, filterer bind.ContractFilterer) (*EvaluationtokenFilterer, error) {
	contract, err := bindEvaluationtoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EvaluationtokenFilterer{contract: contract}, nil
}

// bindEvaluationtoken binds a generic wrapper to an already deployed contract.
func bindEvaluationtoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EvaluationtokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Evaluationtoken *EvaluationtokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Evaluationtoken.Contract.EvaluationtokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Evaluationtoken *EvaluationtokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.EvaluationtokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Evaluationtoken *EvaluationtokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.EvaluationtokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Evaluationtoken *EvaluationtokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Evaluationtoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Evaluationtoken *EvaluationtokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Evaluationtoken *EvaluationtokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Evaluationtoken *EvaluationtokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Evaluationtoken.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Evaluationtoken *EvaluationtokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Evaluationtoken.Contract.BalanceOf(&_Evaluationtoken.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Evaluationtoken *EvaluationtokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Evaluationtoken.Contract.BalanceOf(&_Evaluationtoken.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Evaluationtoken *EvaluationtokenCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Evaluationtoken.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Evaluationtoken *EvaluationtokenSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Evaluationtoken.Contract.GetApproved(&_Evaluationtoken.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Evaluationtoken *EvaluationtokenCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Evaluationtoken.Contract.GetApproved(&_Evaluationtoken.CallOpts, tokenId)
}

// GetTokenURI is a free data retrieval call binding the contract method 0x3bb3a24d.
//
// Solidity: function getTokenURI(uint256 tokenId) view returns(string)
func (_Evaluationtoken *EvaluationtokenCaller) GetTokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Evaluationtoken.contract.Call(opts, &out, "getTokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetTokenURI is a free data retrieval call binding the contract method 0x3bb3a24d.
//
// Solidity: function getTokenURI(uint256 tokenId) view returns(string)
func (_Evaluationtoken *EvaluationtokenSession) GetTokenURI(tokenId *big.Int) (string, error) {
	return _Evaluationtoken.Contract.GetTokenURI(&_Evaluationtoken.CallOpts, tokenId)
}

// GetTokenURI is a free data retrieval call binding the contract method 0x3bb3a24d.
//
// Solidity: function getTokenURI(uint256 tokenId) view returns(string)
func (_Evaluationtoken *EvaluationtokenCallerSession) GetTokenURI(tokenId *big.Int) (string, error) {
	return _Evaluationtoken.Contract.GetTokenURI(&_Evaluationtoken.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Evaluationtoken *EvaluationtokenCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Evaluationtoken.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Evaluationtoken *EvaluationtokenSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Evaluationtoken.Contract.IsApprovedForAll(&_Evaluationtoken.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Evaluationtoken *EvaluationtokenCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Evaluationtoken.Contract.IsApprovedForAll(&_Evaluationtoken.CallOpts, owner, operator)
}

// IsEvaluationScoreSet is a free data retrieval call binding the contract method 0x1e2af3cb.
//
// Solidity: function isEvaluationScoreSet(address requestorAddress, string topicName, address validatorAddress, address trainerAddress) view returns(bool)
func (_Evaluationtoken *EvaluationtokenCaller) IsEvaluationScoreSet(opts *bind.CallOpts, requestorAddress common.Address, topicName string, validatorAddress common.Address, trainerAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Evaluationtoken.contract.Call(opts, &out, "isEvaluationScoreSet", requestorAddress, topicName, validatorAddress, trainerAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEvaluationScoreSet is a free data retrieval call binding the contract method 0x1e2af3cb.
//
// Solidity: function isEvaluationScoreSet(address requestorAddress, string topicName, address validatorAddress, address trainerAddress) view returns(bool)
func (_Evaluationtoken *EvaluationtokenSession) IsEvaluationScoreSet(requestorAddress common.Address, topicName string, validatorAddress common.Address, trainerAddress common.Address) (bool, error) {
	return _Evaluationtoken.Contract.IsEvaluationScoreSet(&_Evaluationtoken.CallOpts, requestorAddress, topicName, validatorAddress, trainerAddress)
}

// IsEvaluationScoreSet is a free data retrieval call binding the contract method 0x1e2af3cb.
//
// Solidity: function isEvaluationScoreSet(address requestorAddress, string topicName, address validatorAddress, address trainerAddress) view returns(bool)
func (_Evaluationtoken *EvaluationtokenCallerSession) IsEvaluationScoreSet(requestorAddress common.Address, topicName string, validatorAddress common.Address, trainerAddress common.Address) (bool, error) {
	return _Evaluationtoken.Contract.IsEvaluationScoreSet(&_Evaluationtoken.CallOpts, requestorAddress, topicName, validatorAddress, trainerAddress)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Evaluationtoken *EvaluationtokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Evaluationtoken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Evaluationtoken *EvaluationtokenSession) Name() (string, error) {
	return _Evaluationtoken.Contract.Name(&_Evaluationtoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Evaluationtoken *EvaluationtokenCallerSession) Name() (string, error) {
	return _Evaluationtoken.Contract.Name(&_Evaluationtoken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Evaluationtoken *EvaluationtokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Evaluationtoken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Evaluationtoken *EvaluationtokenSession) Owner() (common.Address, error) {
	return _Evaluationtoken.Contract.Owner(&_Evaluationtoken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Evaluationtoken *EvaluationtokenCallerSession) Owner() (common.Address, error) {
	return _Evaluationtoken.Contract.Owner(&_Evaluationtoken.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Evaluationtoken *EvaluationtokenCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Evaluationtoken.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Evaluationtoken *EvaluationtokenSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Evaluationtoken.Contract.OwnerOf(&_Evaluationtoken.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Evaluationtoken *EvaluationtokenCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Evaluationtoken.Contract.OwnerOf(&_Evaluationtoken.CallOpts, tokenId)
}

// ProtocolDeployer is a free data retrieval call binding the contract method 0x330c093c.
//
// Solidity: function protocolDeployer() view returns(address)
func (_Evaluationtoken *EvaluationtokenCaller) ProtocolDeployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Evaluationtoken.contract.Call(opts, &out, "protocolDeployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolDeployer is a free data retrieval call binding the contract method 0x330c093c.
//
// Solidity: function protocolDeployer() view returns(address)
func (_Evaluationtoken *EvaluationtokenSession) ProtocolDeployer() (common.Address, error) {
	return _Evaluationtoken.Contract.ProtocolDeployer(&_Evaluationtoken.CallOpts)
}

// ProtocolDeployer is a free data retrieval call binding the contract method 0x330c093c.
//
// Solidity: function protocolDeployer() view returns(address)
func (_Evaluationtoken *EvaluationtokenCallerSession) ProtocolDeployer() (common.Address, error) {
	return _Evaluationtoken.Contract.ProtocolDeployer(&_Evaluationtoken.CallOpts)
}

// ScatterProtocolContract is a free data retrieval call binding the contract method 0x572d107b.
//
// Solidity: function scatterProtocolContract() view returns(address)
func (_Evaluationtoken *EvaluationtokenCaller) ScatterProtocolContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Evaluationtoken.contract.Call(opts, &out, "scatterProtocolContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ScatterProtocolContract is a free data retrieval call binding the contract method 0x572d107b.
//
// Solidity: function scatterProtocolContract() view returns(address)
func (_Evaluationtoken *EvaluationtokenSession) ScatterProtocolContract() (common.Address, error) {
	return _Evaluationtoken.Contract.ScatterProtocolContract(&_Evaluationtoken.CallOpts)
}

// ScatterProtocolContract is a free data retrieval call binding the contract method 0x572d107b.
//
// Solidity: function scatterProtocolContract() view returns(address)
func (_Evaluationtoken *EvaluationtokenCallerSession) ScatterProtocolContract() (common.Address, error) {
	return _Evaluationtoken.Contract.ScatterProtocolContract(&_Evaluationtoken.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Evaluationtoken *EvaluationtokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Evaluationtoken.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Evaluationtoken *EvaluationtokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Evaluationtoken.Contract.SupportsInterface(&_Evaluationtoken.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Evaluationtoken *EvaluationtokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Evaluationtoken.Contract.SupportsInterface(&_Evaluationtoken.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Evaluationtoken *EvaluationtokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Evaluationtoken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Evaluationtoken *EvaluationtokenSession) Symbol() (string, error) {
	return _Evaluationtoken.Contract.Symbol(&_Evaluationtoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Evaluationtoken *EvaluationtokenCallerSession) Symbol() (string, error) {
	return _Evaluationtoken.Contract.Symbol(&_Evaluationtoken.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Evaluationtoken *EvaluationtokenCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Evaluationtoken.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Evaluationtoken *EvaluationtokenSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Evaluationtoken.Contract.TokenURI(&_Evaluationtoken.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Evaluationtoken *EvaluationtokenCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Evaluationtoken.Contract.TokenURI(&_Evaluationtoken.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Evaluationtoken *EvaluationtokenTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Evaluationtoken.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Evaluationtoken *EvaluationtokenSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.Approve(&_Evaluationtoken.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Evaluationtoken *EvaluationtokenTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.Approve(&_Evaluationtoken.TransactOpts, to, tokenId)
}

// PublishEvaluationData is a paid mutator transaction binding the contract method 0x3c7645b7.
//
// Solidity: function publishEvaluationData(address dataPublisher, string jobURI, string evaluationDataURI) returns()
func (_Evaluationtoken *EvaluationtokenTransactor) PublishEvaluationData(opts *bind.TransactOpts, dataPublisher common.Address, jobURI string, evaluationDataURI string) (*types.Transaction, error) {
	return _Evaluationtoken.contract.Transact(opts, "publishEvaluationData", dataPublisher, jobURI, evaluationDataURI)
}

// PublishEvaluationData is a paid mutator transaction binding the contract method 0x3c7645b7.
//
// Solidity: function publishEvaluationData(address dataPublisher, string jobURI, string evaluationDataURI) returns()
func (_Evaluationtoken *EvaluationtokenSession) PublishEvaluationData(dataPublisher common.Address, jobURI string, evaluationDataURI string) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.PublishEvaluationData(&_Evaluationtoken.TransactOpts, dataPublisher, jobURI, evaluationDataURI)
}

// PublishEvaluationData is a paid mutator transaction binding the contract method 0x3c7645b7.
//
// Solidity: function publishEvaluationData(address dataPublisher, string jobURI, string evaluationDataURI) returns()
func (_Evaluationtoken *EvaluationtokenTransactorSession) PublishEvaluationData(dataPublisher common.Address, jobURI string, evaluationDataURI string) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.PublishEvaluationData(&_Evaluationtoken.TransactOpts, dataPublisher, jobURI, evaluationDataURI)
}

// PublishEvaluationJob is a paid mutator transaction binding the contract method 0x49db9dc3.
//
// Solidity: function publishEvaluationJob(address recipient, string tokenURI) returns(uint256)
func (_Evaluationtoken *EvaluationtokenTransactor) PublishEvaluationJob(opts *bind.TransactOpts, recipient common.Address, tokenURI string) (*types.Transaction, error) {
	return _Evaluationtoken.contract.Transact(opts, "publishEvaluationJob", recipient, tokenURI)
}

// PublishEvaluationJob is a paid mutator transaction binding the contract method 0x49db9dc3.
//
// Solidity: function publishEvaluationJob(address recipient, string tokenURI) returns(uint256)
func (_Evaluationtoken *EvaluationtokenSession) PublishEvaluationJob(recipient common.Address, tokenURI string) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.PublishEvaluationJob(&_Evaluationtoken.TransactOpts, recipient, tokenURI)
}

// PublishEvaluationJob is a paid mutator transaction binding the contract method 0x49db9dc3.
//
// Solidity: function publishEvaluationJob(address recipient, string tokenURI) returns(uint256)
func (_Evaluationtoken *EvaluationtokenTransactorSession) PublishEvaluationJob(recipient common.Address, tokenURI string) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.PublishEvaluationJob(&_Evaluationtoken.TransactOpts, recipient, tokenURI)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Evaluationtoken *EvaluationtokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Evaluationtoken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Evaluationtoken *EvaluationtokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _Evaluationtoken.Contract.RenounceOwnership(&_Evaluationtoken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Evaluationtoken *EvaluationtokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Evaluationtoken.Contract.RenounceOwnership(&_Evaluationtoken.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Evaluationtoken *EvaluationtokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Evaluationtoken.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Evaluationtoken *EvaluationtokenSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.SafeTransferFrom(&_Evaluationtoken.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Evaluationtoken *EvaluationtokenTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.SafeTransferFrom(&_Evaluationtoken.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Evaluationtoken *EvaluationtokenTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Evaluationtoken.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Evaluationtoken *EvaluationtokenSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.SafeTransferFrom0(&_Evaluationtoken.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Evaluationtoken *EvaluationtokenTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.SafeTransferFrom0(&_Evaluationtoken.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Evaluationtoken *EvaluationtokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Evaluationtoken.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Evaluationtoken *EvaluationtokenSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.SetApprovalForAll(&_Evaluationtoken.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Evaluationtoken *EvaluationtokenTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.SetApprovalForAll(&_Evaluationtoken.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI_) returns()
func (_Evaluationtoken *EvaluationtokenTransactor) SetBaseURI(opts *bind.TransactOpts, baseURI_ string) (*types.Transaction, error) {
	return _Evaluationtoken.contract.Transact(opts, "setBaseURI", baseURI_)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI_) returns()
func (_Evaluationtoken *EvaluationtokenSession) SetBaseURI(baseURI_ string) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.SetBaseURI(&_Evaluationtoken.TransactOpts, baseURI_)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI_) returns()
func (_Evaluationtoken *EvaluationtokenTransactorSession) SetBaseURI(baseURI_ string) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.SetBaseURI(&_Evaluationtoken.TransactOpts, baseURI_)
}

// SetScatterContractAddress is a paid mutator transaction binding the contract method 0x126d5b6a.
//
// Solidity: function setScatterContractAddress(address contractAddress) returns()
func (_Evaluationtoken *EvaluationtokenTransactor) SetScatterContractAddress(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _Evaluationtoken.contract.Transact(opts, "setScatterContractAddress", contractAddress)
}

// SetScatterContractAddress is a paid mutator transaction binding the contract method 0x126d5b6a.
//
// Solidity: function setScatterContractAddress(address contractAddress) returns()
func (_Evaluationtoken *EvaluationtokenSession) SetScatterContractAddress(contractAddress common.Address) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.SetScatterContractAddress(&_Evaluationtoken.TransactOpts, contractAddress)
}

// SetScatterContractAddress is a paid mutator transaction binding the contract method 0x126d5b6a.
//
// Solidity: function setScatterContractAddress(address contractAddress) returns()
func (_Evaluationtoken *EvaluationtokenTransactorSession) SetScatterContractAddress(contractAddress common.Address) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.SetScatterContractAddress(&_Evaluationtoken.TransactOpts, contractAddress)
}

// SubmitEvaluationScore is a paid mutator transaction binding the contract method 0x51bd2421.
//
// Solidity: function submitEvaluationScore(address requestorAddress, string topicName, address trainerAddress, address validatorAddress, uint256 score) returns()
func (_Evaluationtoken *EvaluationtokenTransactor) SubmitEvaluationScore(opts *bind.TransactOpts, requestorAddress common.Address, topicName string, trainerAddress common.Address, validatorAddress common.Address, score *big.Int) (*types.Transaction, error) {
	return _Evaluationtoken.contract.Transact(opts, "submitEvaluationScore", requestorAddress, topicName, trainerAddress, validatorAddress, score)
}

// SubmitEvaluationScore is a paid mutator transaction binding the contract method 0x51bd2421.
//
// Solidity: function submitEvaluationScore(address requestorAddress, string topicName, address trainerAddress, address validatorAddress, uint256 score) returns()
func (_Evaluationtoken *EvaluationtokenSession) SubmitEvaluationScore(requestorAddress common.Address, topicName string, trainerAddress common.Address, validatorAddress common.Address, score *big.Int) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.SubmitEvaluationScore(&_Evaluationtoken.TransactOpts, requestorAddress, topicName, trainerAddress, validatorAddress, score)
}

// SubmitEvaluationScore is a paid mutator transaction binding the contract method 0x51bd2421.
//
// Solidity: function submitEvaluationScore(address requestorAddress, string topicName, address trainerAddress, address validatorAddress, uint256 score) returns()
func (_Evaluationtoken *EvaluationtokenTransactorSession) SubmitEvaluationScore(requestorAddress common.Address, topicName string, trainerAddress common.Address, validatorAddress common.Address, score *big.Int) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.SubmitEvaluationScore(&_Evaluationtoken.TransactOpts, requestorAddress, topicName, trainerAddress, validatorAddress, score)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Evaluationtoken *EvaluationtokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Evaluationtoken.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Evaluationtoken *EvaluationtokenSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.TransferFrom(&_Evaluationtoken.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Evaluationtoken *EvaluationtokenTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.TransferFrom(&_Evaluationtoken.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Evaluationtoken *EvaluationtokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Evaluationtoken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Evaluationtoken *EvaluationtokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.TransferOwnership(&_Evaluationtoken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Evaluationtoken *EvaluationtokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.TransferOwnership(&_Evaluationtoken.TransactOpts, newOwner)
}

// EvaluationtokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Evaluationtoken contract.
type EvaluationtokenApprovalIterator struct {
	Event *EvaluationtokenApproval // Event containing the contract specifics and raw log

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
func (it *EvaluationtokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EvaluationtokenApproval)
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
		it.Event = new(EvaluationtokenApproval)
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
func (it *EvaluationtokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EvaluationtokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EvaluationtokenApproval represents a Approval event raised by the Evaluationtoken contract.
type EvaluationtokenApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Evaluationtoken *EvaluationtokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*EvaluationtokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Evaluationtoken.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &EvaluationtokenApprovalIterator{contract: _Evaluationtoken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Evaluationtoken *EvaluationtokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EvaluationtokenApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Evaluationtoken.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EvaluationtokenApproval)
				if err := _Evaluationtoken.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Evaluationtoken *EvaluationtokenFilterer) ParseApproval(log types.Log) (*EvaluationtokenApproval, error) {
	event := new(EvaluationtokenApproval)
	if err := _Evaluationtoken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EvaluationtokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Evaluationtoken contract.
type EvaluationtokenApprovalForAllIterator struct {
	Event *EvaluationtokenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *EvaluationtokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EvaluationtokenApprovalForAll)
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
		it.Event = new(EvaluationtokenApprovalForAll)
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
func (it *EvaluationtokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EvaluationtokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EvaluationtokenApprovalForAll represents a ApprovalForAll event raised by the Evaluationtoken contract.
type EvaluationtokenApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Evaluationtoken *EvaluationtokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*EvaluationtokenApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Evaluationtoken.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &EvaluationtokenApprovalForAllIterator{contract: _Evaluationtoken.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Evaluationtoken *EvaluationtokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *EvaluationtokenApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Evaluationtoken.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EvaluationtokenApprovalForAll)
				if err := _Evaluationtoken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Evaluationtoken *EvaluationtokenFilterer) ParseApprovalForAll(log types.Log) (*EvaluationtokenApprovalForAll, error) {
	event := new(EvaluationtokenApprovalForAll)
	if err := _Evaluationtoken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EvaluationtokenBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the Evaluationtoken contract.
type EvaluationtokenBatchMetadataUpdateIterator struct {
	Event *EvaluationtokenBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *EvaluationtokenBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EvaluationtokenBatchMetadataUpdate)
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
		it.Event = new(EvaluationtokenBatchMetadataUpdate)
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
func (it *EvaluationtokenBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EvaluationtokenBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EvaluationtokenBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the Evaluationtoken contract.
type EvaluationtokenBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Evaluationtoken *EvaluationtokenFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*EvaluationtokenBatchMetadataUpdateIterator, error) {

	logs, sub, err := _Evaluationtoken.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &EvaluationtokenBatchMetadataUpdateIterator{contract: _Evaluationtoken.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Evaluationtoken *EvaluationtokenFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *EvaluationtokenBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Evaluationtoken.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EvaluationtokenBatchMetadataUpdate)
				if err := _Evaluationtoken.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Evaluationtoken *EvaluationtokenFilterer) ParseBatchMetadataUpdate(log types.Log) (*EvaluationtokenBatchMetadataUpdate, error) {
	event := new(EvaluationtokenBatchMetadataUpdate)
	if err := _Evaluationtoken.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EvaluationtokenMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the Evaluationtoken contract.
type EvaluationtokenMetadataUpdateIterator struct {
	Event *EvaluationtokenMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *EvaluationtokenMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EvaluationtokenMetadataUpdate)
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
		it.Event = new(EvaluationtokenMetadataUpdate)
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
func (it *EvaluationtokenMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EvaluationtokenMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EvaluationtokenMetadataUpdate represents a MetadataUpdate event raised by the Evaluationtoken contract.
type EvaluationtokenMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Evaluationtoken *EvaluationtokenFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*EvaluationtokenMetadataUpdateIterator, error) {

	logs, sub, err := _Evaluationtoken.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &EvaluationtokenMetadataUpdateIterator{contract: _Evaluationtoken.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Evaluationtoken *EvaluationtokenFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *EvaluationtokenMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Evaluationtoken.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EvaluationtokenMetadataUpdate)
				if err := _Evaluationtoken.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Evaluationtoken *EvaluationtokenFilterer) ParseMetadataUpdate(log types.Log) (*EvaluationtokenMetadataUpdate, error) {
	event := new(EvaluationtokenMetadataUpdate)
	if err := _Evaluationtoken.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EvaluationtokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Evaluationtoken contract.
type EvaluationtokenOwnershipTransferredIterator struct {
	Event *EvaluationtokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EvaluationtokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EvaluationtokenOwnershipTransferred)
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
		it.Event = new(EvaluationtokenOwnershipTransferred)
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
func (it *EvaluationtokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EvaluationtokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EvaluationtokenOwnershipTransferred represents a OwnershipTransferred event raised by the Evaluationtoken contract.
type EvaluationtokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Evaluationtoken *EvaluationtokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EvaluationtokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Evaluationtoken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EvaluationtokenOwnershipTransferredIterator{contract: _Evaluationtoken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Evaluationtoken *EvaluationtokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EvaluationtokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Evaluationtoken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EvaluationtokenOwnershipTransferred)
				if err := _Evaluationtoken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Evaluationtoken *EvaluationtokenFilterer) ParseOwnershipTransferred(log types.Log) (*EvaluationtokenOwnershipTransferred, error) {
	event := new(EvaluationtokenOwnershipTransferred)
	if err := _Evaluationtoken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EvaluationtokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Evaluationtoken contract.
type EvaluationtokenTransferIterator struct {
	Event *EvaluationtokenTransfer // Event containing the contract specifics and raw log

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
func (it *EvaluationtokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EvaluationtokenTransfer)
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
		it.Event = new(EvaluationtokenTransfer)
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
func (it *EvaluationtokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EvaluationtokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EvaluationtokenTransfer represents a Transfer event raised by the Evaluationtoken contract.
type EvaluationtokenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Evaluationtoken *EvaluationtokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*EvaluationtokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Evaluationtoken.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &EvaluationtokenTransferIterator{contract: _Evaluationtoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Evaluationtoken *EvaluationtokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EvaluationtokenTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Evaluationtoken.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EvaluationtokenTransfer)
				if err := _Evaluationtoken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Evaluationtoken *EvaluationtokenFilterer) ParseTransfer(log types.Log) (*EvaluationtokenTransfer, error) {
	event := new(EvaluationtokenTransfer)
	if err := _Evaluationtoken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
