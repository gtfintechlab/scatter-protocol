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
	_ = abi.ConvertType
)

// EvaluationtokenMetaData contains all meta data concerning the Evaluationtoken contract.
var EvaluationtokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"getAverageScoreForTrainerForJob\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getTokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"isEvaluationScoreSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolDeployer\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"jobURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"evaluationDataURI\",\"type\":\"string\"}],\"name\":\"publishEvaluationData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"publishEvaluationJob\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scatterProtocolContract\",\"outputs\":[{\"internalType\":\"contractIScatterProtocol\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setScatterContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setVoteManagerContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"score\",\"type\":\"uint256\"}],\"name\":\"submitEvaluationScore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteManagerContract\",\"outputs\":[{\"internalType\":\"contractIVoteManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b5060408051808201825260208082527f536361747465722050726f746f636f6c204576616c756174696f6e204a6f6273818301528251808401909352600483526329a822a560e11b90830152905f6200006a8382620001a1565b506001620000798282620001a1565b5050506200009662000090620000ae60201b60201c565b620000b2565b600c80546001600160a01b031916331790556200026d565b3390565b600780546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b634e487b7160e01b5f52604160045260245ffd5b600181811c908216806200012c57607f821691505b6020821081036200014b57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156200019c57805f5260205f20601f840160051c81016020851015620001785750805b601f840160051c820191505b8181101562000199575f815560010162000184565b50505b505050565b81516001600160401b03811115620001bd57620001bd62000103565b620001d581620001ce845462000117565b8462000151565b602080601f8311600181146200020b575f8415620001f35750858301515b5f19600386901b1c1916600185901b17855562000265565b5f85815260208120601f198616915b828110156200023b578886015182559484019460019091019084016200021a565b50858210156200025957878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b61256d806200027b5f395ff3fe608060405234801561000f575f80fd5b50600436106101bb575f3560e01c806355f804b3116100f3578063a22cb46511610093578063d7287ff71161006e578063d7287ff7146103ae578063e54c23e8146103c1578063e985e9c5146103d4578063f2fde38b146103e7575f80fd5b8063a22cb46514610375578063b88d4fde14610388578063c87b56dd1461039b575f80fd5b806370a08231116100ce57806370a0823114610333578063715018a6146103545780638da5cb5b1461035c57806395d89b411461036d575f80fd5b806355f804b3146102fa578063572d107b1461030d5780636352211e14610320575f80fd5b806323b872dd1161015e5780633bb3a24d116101395780633bb3a24d146102ae57806342842e0e146102c157806342d098e5146102d457806351bd2421146102e7575f80fd5b806323b872dd14610275578063330c093c1461028857806338fcb3b21461029b575f80fd5b806309092301116101995780630909230114610227578063095ea7b31461023c578063126d5b6a1461024f5780631e2af3cb14610262575f80fd5b806301ffc9a7146101bf57806306fdde03146101e7578063081812fc146101fc575b5f80fd5b6101d26101cd366004611bbd565b6103fa565b60405190151581526020015b60405180910390f35b6101ef610424565b6040516101de9190611c25565b61020f61020a366004611c37565b6104b3565b6040516001600160a01b0390911681526020016101de565b61023a610235366004611d0e565b6104d8565b005b61023a61024a366004611da1565b610761565b61023a61025d366004611dc9565b610875565b6101d2610270366004611de2565b61089f565b61023a610283366004611e4d565b610903565b600c5461020f906001600160a01b031681565b600e5461020f906001600160a01b031681565b6101ef6102bc366004611c37565b610934565b61023a6102cf366004611e4d565b61093f565b61023a6102e2366004611dc9565b610959565b61023a6102f5366004611e86565b610983565b61023a610308366004611ef8565b610cb0565b600d5461020f906001600160a01b031681565b61020f61032e366004611c37565b610d28565b610346610341366004611dc9565b610d87565b6040519081526020016101de565b61023a610e0b565b6007546001600160a01b031661020f565b6101ef610e1e565b61023a610383366004611f37565b610e2d565b61023a610396366004611f6c565b610e38565b6101ef6103a9366004611c37565b610e70565b6103466103bc366004611fd7565b610f66565b6103466103cf366004612031565b61105f565b6101d26103e23660046120a0565b611105565b61023a6103f5366004611dc9565b611132565b5f6001600160e01b03198216632483248360e11b148061041e575061041e826111ab565b92915050565b60605f8054610432906120d1565b80601f016020809104026020016040519081016040528092919081815260200182805461045e906120d1565b80156104a95780601f10610480576101008083540402835291602001916104a9565b820191905f5260205f20905b81548152906001019060200180831161048c57829003601f168201915b5050505050905090565b5f6104bd826111fa565b505f908152600460205260409020546001600160a01b031690565b600d546001600160a01b0316331461050b5760405162461bcd60e51b815260040161050290612109565b60405180910390fd5b6001600160a01b0384165f908152600960205260408082209051610530908690612166565b90815260200160405180910390208054610549906120d1565b80601f0160208091040260200160405190810160405280929190818152602001828054610575906120d1565b80156105c05780601f10610597576101008083540402835291602001916105c0565b820191905f5260205f20905b8154815290600101906020018083116105a357829003601f168201915b50505050509050828051906020012081805190602001201461064a5760405162461bcd60e51b815260206004820152603960248201527f54686973206576616c756174696f6e206a6f6220646f6573206e6f742062656c60448201527f6f6e6720746f207468652064617461207075626c6973686572000000000000006064820152608401610502565b60408051602080820183525f918290526001600160a01b0388168252600a905281902090517fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470919061069d908790612166565b9081526040519081900360200181206106b591612181565b6040518091039020146107165760405162461bcd60e51b8152602060048201526024808201527f54686973206a6f6220616c7265616479206861732061206d617070696e6720746044820152631bc81a5d60e21b6064820152608401610502565b6001600160a01b0385165f908152600a602052604090819020905183919061073f908790612166565b908152602001604051809103902090816107599190612237565b505050505050565b5f61076b82610d28565b9050806001600160a01b0316836001600160a01b0316036107d85760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b6064820152608401610502565b336001600160a01b03821614806107f457506107f48133611105565b6108665760405162461bcd60e51b815260206004820152603d60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206f7220617070726f76656420666f7220616c6c0000006064820152608401610502565b6108708383611258565b505050565b61087d6112c5565b600d80546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b0384165f9081526011602052604080822090516108c4908690612166565b90815260408051602092819003830190206001600160a01b038681165f908152918452828220908616825290925290205460ff1690505b949350505050565b61090d338261131f565b6109295760405162461bcd60e51b8152600401610502906122f3565b61087083838361137c565b606061041e82610e70565b61087083838360405180602001604052805f815250610e38565b6109616112c5565b600e80546001600160a01b0319166001600160a01b0392909216919091179055565b600d546001600160a01b031633146109ad5760405162461bcd60e51b815260040161050290612109565b6001600160a01b0385165f908152601160205260409081902090516109d3908690612166565b90815260408051602092819003830190206001600160a01b038581165f908152918452828220908716825290925290205460ff16610ca9576064811115610a5c5760405162461bcd60e51b815260206004820152601f60248201527f53636f7265206d757374206265206265747765656e203020616e6420313030006044820152606401610502565b600d546040516303a120eb60e51b81526001600160a01b03909116906374241d6090610a9090889088908790600401612340565b602060405180830381865afa158015610aab573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610acf9190612374565b610b415760405162461bcd60e51b815260206004820152603760248201527f596f7520617265206e6f7420616e2061737369676e65642076616c696461746f60448201527f7220666f72207468697320747261696e696e67206a6f620000000000000000006064820152608401610502565b6001600160a01b0385165f90815260106020526040908190209051829190610b6a908790612166565b9081526040805191829003602090810183206001600160a01b038088165f9081529183528382208982168352835283822095909555938916845260129052822090610bb6908790612166565b9081526040805191829003602090810183206001600160a01b038089165f908152918352838220805460018181018355828552858520909101899055918c16835260119093529290209093509091610c0f908890612166565b90815260408051602092819003830181206001600160a01b038881165f9081529185528382208a8216835290945291909120805460ff191693151593909317909255600e5463297a43ed60e01b8352169063297a43ed90610c7a908990899088908a9060040161238f565b5f604051808303815f87803b158015610c91575f80fd5b505af1158015610ca3573d5f803e3d5ffd5b50505050505b5050505050565b600c546001600160a01b03163314610d185760405162461bcd60e51b815260206004820152602560248201527f4f6e6c7920746865206f776e65722063616e206163636573732074686973206d604482015264195d1a1bd960da1b6064820152608401610502565b600b610d248282612237565b5050565b5f818152600260205260408120546001600160a01b03168061041e5760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b6044820152606401610502565b5f6001600160a01b038216610df05760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b6064820152608401610502565b506001600160a01b03165f9081526003602052604090205490565b610e136112c5565b610e1c5f6114de565b565b606060018054610432906120d1565b610d2433838361152f565b610e42338361131f565b610e5e5760405162461bcd60e51b8152600401610502906122f3565b610e6a848484846115fc565b50505050565b6060610e7b826111fa565b5f8281526006602052604081208054610e93906120d1565b80601f0160208091040260200160405190810160405280929190818152602001828054610ebf906120d1565b8015610f0a5780601f10610ee157610100808354040283529160200191610f0a565b820191905f5260205f20905b815481529060010190602001808311610eed57829003601f168201915b505050505090505f610f1a61162f565b905080515f03610f2b575092915050565b815115610f5d578082604051602001610f459291906123c9565b60405160208183030381529060405292505050919050565b6108fb8461163e565b6001600160a01b0383165f908152601260205260408082209051829190610f8e908690612166565b9081526040805191829003602090810183206001600160a01b0387165f9081529082528290208054808302850183019093528284529190830182828015610ff257602002820191905f5260205f20905b815481526020019060010190808311610fde575b505050505090505f805b825181101561103457828181518110611017576110176123f7565b60200260200101518261102a919061240b565b9150600101610ffc565b5081515f03611047575f92505050611058565b8151611053908261242a565b925050505b9392505050565b600d545f906001600160a01b0316331461108b5760405162461bcd60e51b815260040161050290612109565b611099600f80546001019055565b5f6110a3600f5490565b90506110af85826116a1565b6110b981846116ba565b6001600160a01b0385165f908152600960205260409081902090518491906110e2908790612166565b908152602001604051809103902090816110fc9190612237565b50949350505050565b6001600160a01b039182165f90815260056020908152604080832093909416825291909152205460ff1690565b61113a6112c5565b6001600160a01b03811661119f5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610502565b6111a8816114de565b50565b5f6001600160e01b031982166380ac58cd60e01b14806111db57506001600160e01b03198216635b5e139f60e01b145b8061041e57506301ffc9a760e01b6001600160e01b031983161461041e565b5f818152600260205260409020546001600160a01b03166111a85760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b6044820152606401610502565b5f81815260046020526040902080546001600160a01b0319166001600160a01b038416908117909155819061128c82610d28565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6007546001600160a01b03163314610e1c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610502565b5f8061132a83610d28565b9050806001600160a01b0316846001600160a01b0316148061135157506113518185611105565b806108fb5750836001600160a01b031661136a846104b3565b6001600160a01b031614949350505050565b826001600160a01b031661138f82610d28565b6001600160a01b0316146113b55760405162461bcd60e51b815260040161050290612449565b6001600160a01b0382166114175760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b6064820152608401610502565b826001600160a01b031661142a82610d28565b6001600160a01b0316146114505760405162461bcd60e51b815260040161050290612449565b5f81815260046020908152604080832080546001600160a01b03199081169091556001600160a01b038781168086526003855283862080545f1901905590871680865283862080546001019055868652600290945282852080549092168417909155905184937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b600780546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b816001600160a01b0316836001600160a01b0316036115905760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c6572000000000000006044820152606401610502565b6001600160a01b038381165f81815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b61160784848461137c565b611613848484846116c4565b610e6a5760405162461bcd60e51b81526004016105029061248e565b6060600b8054610432906120d1565b6060611649826111fa565b5f61165261162f565b90505f8151116116705760405180602001604052805f815250611058565b8061167a846117be565b60405160200161168b9291906123c9565b6040516020818303038152906040529392505050565b610d24828260405180602001604052805f81525061184e565b610d248282611880565b5f6001600160a01b0384163b156117b657604051630a85bd0160e11b81526001600160a01b0385169063150b7a02906117079033908990889088906004016124e0565b6020604051808303815f875af1925050508015611741575060408051601f3d908101601f1916820190925261173e9181019061251c565b60015b61179c573d80801561176e576040519150601f19603f3d011682016040523d82523d5f602084013e611773565b606091505b5080515f036117945760405162461bcd60e51b81526004016105029061248e565b805181602001fd5b6001600160e01b031916630a85bd0160e11b1490506108fb565b5060016108fb565b60605f6117ca83611949565b60010190505f8167ffffffffffffffff8111156117e9576117e9611c69565b6040519080825280601f01601f191660200182016040528015611813576020820181803683370190505b5090508181016020015b5f19016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a850494508461181d57509392505050565b6118588383611a20565b6118645f8484846116c4565b6108705760405162461bcd60e51b81526004016105029061248e565b5f828152600260205260409020546001600160a01b03166118fa5760405162461bcd60e51b815260206004820152602e60248201527f45524337323155524953746f726167653a2055524920736574206f66206e6f6e60448201526d32bc34b9ba32b73a103a37b5b2b760911b6064820152608401610502565b5f8281526006602052604090206119118282612237565b506040518281527ff8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce79060200160405180910390a15050565b5f8072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b83106119875772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef810000000083106119b3576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc1000083106119d157662386f26fc10000830492506010015b6305f5e10083106119e9576305f5e100830492506008015b61271083106119fd57612710830492506004015b60648310611a0f576064830492506002015b600a831061041e5760010192915050565b6001600160a01b038216611a765760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f20616464726573736044820152606401610502565b5f818152600260205260409020546001600160a01b031615611ada5760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610502565b5f818152600260205260409020546001600160a01b031615611b3e5760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610502565b6001600160a01b0382165f81815260036020908152604080832080546001019055848352600290915280822080546001600160a01b0319168417905551839291907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b6001600160e01b0319811681146111a8575f80fd5b5f60208284031215611bcd575f80fd5b813561105881611ba8565b5f5b83811015611bf2578181015183820152602001611bda565b50505f910152565b5f8151808452611c11816020860160208601611bd8565b601f01601f19169290920160200192915050565b602081525f6110586020830184611bfa565b5f60208284031215611c47575f80fd5b5035919050565b80356001600160a01b0381168114611c64575f80fd5b919050565b634e487b7160e01b5f52604160045260245ffd5b5f67ffffffffffffffff80841115611c9757611c97611c69565b604051601f8501601f19908116603f01168101908282118183101715611cbf57611cbf611c69565b81604052809350858152868686011115611cd7575f80fd5b858560208301375f602087830101525050509392505050565b5f82601f830112611cff575f80fd5b61105883833560208501611c7d565b5f805f8060808587031215611d21575f80fd5b611d2a85611c4e565b9350602085013567ffffffffffffffff80821115611d46575f80fd5b611d5288838901611cf0565b94506040870135915080821115611d67575f80fd5b611d7388838901611cf0565b93506060870135915080821115611d88575f80fd5b50611d9587828801611cf0565b91505092959194509250565b5f8060408385031215611db2575f80fd5b611dbb83611c4e565b946020939093013593505050565b5f60208284031215611dd9575f80fd5b61105882611c4e565b5f805f8060808587031215611df5575f80fd5b611dfe85611c4e565b9350602085013567ffffffffffffffff811115611e19575f80fd5b611e2587828801611cf0565b935050611e3460408601611c4e565b9150611e4260608601611c4e565b905092959194509250565b5f805f60608486031215611e5f575f80fd5b611e6884611c4e565b9250611e7660208501611c4e565b9150604084013590509250925092565b5f805f805f60a08688031215611e9a575f80fd5b611ea386611c4e565b9450602086013567ffffffffffffffff811115611ebe575f80fd5b611eca88828901611cf0565b945050611ed960408701611c4e565b9250611ee760608701611c4e565b949793965091946080013592915050565b5f60208284031215611f08575f80fd5b813567ffffffffffffffff811115611f1e575f80fd5b6108fb84828501611cf0565b80151581146111a8575f80fd5b5f8060408385031215611f48575f80fd5b611f5183611c4e565b91506020830135611f6181611f2a565b809150509250929050565b5f805f8060808587031215611f7f575f80fd5b611f8885611c4e565b9350611f9660208601611c4e565b925060408501359150606085013567ffffffffffffffff811115611fb8575f80fd5b8501601f81018713611fc8575f80fd5b611d9587823560208401611c7d565b5f805f60608486031215611fe9575f80fd5b611ff284611c4e565b9250602084013567ffffffffffffffff81111561200d575f80fd5b61201986828701611cf0565b92505061202860408501611c4e565b90509250925092565b5f805f60608486031215612043575f80fd5b61204c84611c4e565b9250602084013567ffffffffffffffff80821115612068575f80fd5b61207487838801611cf0565b93506040860135915080821115612089575f80fd5b5061209686828701611cf0565b9150509250925092565b5f80604083850312156120b1575f80fd5b6120ba83611c4e565b91506120c860208401611c4e565b90509250929050565b600181811c908216806120e557607f821691505b60208210810361210357634e487b7160e01b5f52602260045260245ffd5b50919050565b6020808252603f908201527f54686973206d6574686f642063616e206f6e6c792062652063616c6c6564206260408201527f792074686520736361747465722070726f746f636f6c20636f6e747261637400606082015260800190565b5f8251612177818460208701611bd8565b9190910192915050565b5f80835461218e816120d1565b600182811680156121a657600181146121bb576121e7565b60ff19841687528215158302870194506121e7565b875f526020805f205f5b858110156121de5781548a8201529084019082016121c5565b50505082870194505b50929695505050505050565b601f82111561087057805f5260205f20601f840160051c810160208510156122185750805b601f840160051c820191505b81811015610ca9575f8155600101612224565b815167ffffffffffffffff81111561225157612251611c69565b6122658161225f84546120d1565b846121f3565b602080601f831160018114612298575f84156122815750858301515b5f19600386901b1c1916600185901b178555610759565b5f85815260208120601f198616915b828110156122c6578886015182559484019460019091019084016122a7565b50858210156122e357878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b6020808252602d908201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560408201526c1c881bdc88185c1c1c9bdd9959609a1b606082015260800190565b5f60018060a01b038086168352606060208401526123616060840186611bfa565b9150808416604084015250949350505050565b5f60208284031215612384575f80fd5b815161105881611f2a565b5f60018060a01b038087168352608060208401526123b06080840187611bfa565b9481166040840152929092166060909101525092915050565b5f83516123da818460208801611bd8565b8351908301906123ee818360208801611bd8565b01949350505050565b634e487b7160e01b5f52603260045260245ffd5b8082018082111561041e57634e487b7160e01b5f52601160045260245ffd5b5f8261244457634e487b7160e01b5f52601260045260245ffd5b500490565b60208082526025908201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060408201526437bbb732b960d91b606082015260800190565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b6001600160a01b03858116825284166020820152604081018390526080606082018190525f9061251290830184611bfa565b9695505050505050565b5f6020828403121561252c575f80fd5b815161105881611ba856fea264697066735822122070a9bb26e8a472f7edd8673c24dee3b575265225a40534dcfd2c81bd94a9861a64736f6c63430008170033",
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
	parsed, err := EvaluationtokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// GetAverageScoreForTrainerForJob is a free data retrieval call binding the contract method 0xd7287ff7.
//
// Solidity: function getAverageScoreForTrainerForJob(address requestorAddress, string topicName, address trainerAddress) view returns(uint256)
func (_Evaluationtoken *EvaluationtokenCaller) GetAverageScoreForTrainerForJob(opts *bind.CallOpts, requestorAddress common.Address, topicName string, trainerAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Evaluationtoken.contract.Call(opts, &out, "getAverageScoreForTrainerForJob", requestorAddress, topicName, trainerAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAverageScoreForTrainerForJob is a free data retrieval call binding the contract method 0xd7287ff7.
//
// Solidity: function getAverageScoreForTrainerForJob(address requestorAddress, string topicName, address trainerAddress) view returns(uint256)
func (_Evaluationtoken *EvaluationtokenSession) GetAverageScoreForTrainerForJob(requestorAddress common.Address, topicName string, trainerAddress common.Address) (*big.Int, error) {
	return _Evaluationtoken.Contract.GetAverageScoreForTrainerForJob(&_Evaluationtoken.CallOpts, requestorAddress, topicName, trainerAddress)
}

// GetAverageScoreForTrainerForJob is a free data retrieval call binding the contract method 0xd7287ff7.
//
// Solidity: function getAverageScoreForTrainerForJob(address requestorAddress, string topicName, address trainerAddress) view returns(uint256)
func (_Evaluationtoken *EvaluationtokenCallerSession) GetAverageScoreForTrainerForJob(requestorAddress common.Address, topicName string, trainerAddress common.Address) (*big.Int, error) {
	return _Evaluationtoken.Contract.GetAverageScoreForTrainerForJob(&_Evaluationtoken.CallOpts, requestorAddress, topicName, trainerAddress)
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

// VoteManagerContract is a free data retrieval call binding the contract method 0x38fcb3b2.
//
// Solidity: function voteManagerContract() view returns(address)
func (_Evaluationtoken *EvaluationtokenCaller) VoteManagerContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Evaluationtoken.contract.Call(opts, &out, "voteManagerContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoteManagerContract is a free data retrieval call binding the contract method 0x38fcb3b2.
//
// Solidity: function voteManagerContract() view returns(address)
func (_Evaluationtoken *EvaluationtokenSession) VoteManagerContract() (common.Address, error) {
	return _Evaluationtoken.Contract.VoteManagerContract(&_Evaluationtoken.CallOpts)
}

// VoteManagerContract is a free data retrieval call binding the contract method 0x38fcb3b2.
//
// Solidity: function voteManagerContract() view returns(address)
func (_Evaluationtoken *EvaluationtokenCallerSession) VoteManagerContract() (common.Address, error) {
	return _Evaluationtoken.Contract.VoteManagerContract(&_Evaluationtoken.CallOpts)
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

// PublishEvaluationData is a paid mutator transaction binding the contract method 0x09092301.
//
// Solidity: function publishEvaluationData(address requestorAddress, string topicName, string jobURI, string evaluationDataURI) returns()
func (_Evaluationtoken *EvaluationtokenTransactor) PublishEvaluationData(opts *bind.TransactOpts, requestorAddress common.Address, topicName string, jobURI string, evaluationDataURI string) (*types.Transaction, error) {
	return _Evaluationtoken.contract.Transact(opts, "publishEvaluationData", requestorAddress, topicName, jobURI, evaluationDataURI)
}

// PublishEvaluationData is a paid mutator transaction binding the contract method 0x09092301.
//
// Solidity: function publishEvaluationData(address requestorAddress, string topicName, string jobURI, string evaluationDataURI) returns()
func (_Evaluationtoken *EvaluationtokenSession) PublishEvaluationData(requestorAddress common.Address, topicName string, jobURI string, evaluationDataURI string) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.PublishEvaluationData(&_Evaluationtoken.TransactOpts, requestorAddress, topicName, jobURI, evaluationDataURI)
}

// PublishEvaluationData is a paid mutator transaction binding the contract method 0x09092301.
//
// Solidity: function publishEvaluationData(address requestorAddress, string topicName, string jobURI, string evaluationDataURI) returns()
func (_Evaluationtoken *EvaluationtokenTransactorSession) PublishEvaluationData(requestorAddress common.Address, topicName string, jobURI string, evaluationDataURI string) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.PublishEvaluationData(&_Evaluationtoken.TransactOpts, requestorAddress, topicName, jobURI, evaluationDataURI)
}

// PublishEvaluationJob is a paid mutator transaction binding the contract method 0xe54c23e8.
//
// Solidity: function publishEvaluationJob(address requestorAddress, string topicName, string tokenURI) returns(uint256)
func (_Evaluationtoken *EvaluationtokenTransactor) PublishEvaluationJob(opts *bind.TransactOpts, requestorAddress common.Address, topicName string, tokenURI string) (*types.Transaction, error) {
	return _Evaluationtoken.contract.Transact(opts, "publishEvaluationJob", requestorAddress, topicName, tokenURI)
}

// PublishEvaluationJob is a paid mutator transaction binding the contract method 0xe54c23e8.
//
// Solidity: function publishEvaluationJob(address requestorAddress, string topicName, string tokenURI) returns(uint256)
func (_Evaluationtoken *EvaluationtokenSession) PublishEvaluationJob(requestorAddress common.Address, topicName string, tokenURI string) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.PublishEvaluationJob(&_Evaluationtoken.TransactOpts, requestorAddress, topicName, tokenURI)
}

// PublishEvaluationJob is a paid mutator transaction binding the contract method 0xe54c23e8.
//
// Solidity: function publishEvaluationJob(address requestorAddress, string topicName, string tokenURI) returns(uint256)
func (_Evaluationtoken *EvaluationtokenTransactorSession) PublishEvaluationJob(requestorAddress common.Address, topicName string, tokenURI string) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.PublishEvaluationJob(&_Evaluationtoken.TransactOpts, requestorAddress, topicName, tokenURI)
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

// SetVoteManagerContract is a paid mutator transaction binding the contract method 0x42d098e5.
//
// Solidity: function setVoteManagerContract(address contractAddress) returns()
func (_Evaluationtoken *EvaluationtokenTransactor) SetVoteManagerContract(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _Evaluationtoken.contract.Transact(opts, "setVoteManagerContract", contractAddress)
}

// SetVoteManagerContract is a paid mutator transaction binding the contract method 0x42d098e5.
//
// Solidity: function setVoteManagerContract(address contractAddress) returns()
func (_Evaluationtoken *EvaluationtokenSession) SetVoteManagerContract(contractAddress common.Address) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.SetVoteManagerContract(&_Evaluationtoken.TransactOpts, contractAddress)
}

// SetVoteManagerContract is a paid mutator transaction binding the contract method 0x42d098e5.
//
// Solidity: function setVoteManagerContract(address contractAddress) returns()
func (_Evaluationtoken *EvaluationtokenTransactorSession) SetVoteManagerContract(contractAddress common.Address) (*types.Transaction, error) {
	return _Evaluationtoken.Contract.SetVoteManagerContract(&_Evaluationtoken.TransactOpts, contractAddress)
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
