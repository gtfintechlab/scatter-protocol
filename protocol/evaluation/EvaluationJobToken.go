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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"getAverageScoreForTrainerForJob\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getTokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"}],\"name\":\"isEvaluationScoreSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolDeployer\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dataPublisher\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"jobURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"evaluationDataURI\",\"type\":\"string\"}],\"name\":\"publishEvaluationData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"publishEvaluationJob\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scatterProtocolContract\",\"outputs\":[{\"internalType\":\"contractIScatterProtocol\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setScatterContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setVoteManagerContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"trainerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"score\",\"type\":\"uint256\"}],\"name\":\"submitEvaluationScore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteManagerContract\",\"outputs\":[{\"internalType\":\"contractIVoteManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b5060408051808201825260208082527f536361747465722050726f746f636f6c204576616c756174696f6e204a6f6273818301528251808401909352600483526329a822a560e11b90830152905f6200006a8382620001a1565b506001620000798282620001a1565b5050506200009662000090620000ae60201b60201c565b620000b2565b600c80546001600160a01b031916331790556200026d565b3390565b600780546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b634e487b7160e01b5f52604160045260245ffd5b600181811c908216806200012c57607f821691505b6020821081036200014b57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156200019c57805f5260205f20601f840160051c81016020851015620001785750805b601f840160051c820191505b8181101562000199575f815560010162000184565b50505b505050565b81516001600160401b03811115620001bd57620001bd62000103565b620001d581620001ce845462000117565b8462000151565b602080601f8311600181146200020b575f8415620001f35750858301515b5f19600386901b1c1916600185901b17855562000265565b5f85815260208120601f198616915b828110156200023b578886015182559484019460019091019084016200021a565b50858210156200025957878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b612451806200027b5f395ff3fe608060405234801561000f575f80fd5b50600436106101bb575f3560e01c806351bd2421116100f357806395d89b4111610093578063c87b56dd1161006e578063c87b56dd146103ae578063d7287ff7146103c1578063e985e9c5146103d4578063f2fde38b1461040f575f80fd5b806395d89b4114610380578063a22cb46514610388578063b88d4fde1461039b575f80fd5b80636352211e116100ce5780636352211e1461034157806370a0823114610354578063715018a6146103675780638da5cb5b1461036f575f80fd5b806351bd24211461030857806355f804b31461031b578063572d107b1461032e575f80fd5b8063330c093c1161015e5780633c7645b7116101395780633c7645b7146102ae57806342842e0e146102c157806342d098e5146102d457806349db9dc3146102e7575f80fd5b8063330c093c1461027557806338fcb3b2146102885780633bb3a24d1461029b575f80fd5b8063095ea7b311610199578063095ea7b314610227578063126d5b6a1461023c5780631e2af3cb1461024f57806323b872dd14610262575f80fd5b806301ffc9a7146101bf57806306fdde03146101e7578063081812fc146101fc575b5f80fd5b6101d26101cd366004611b16565b610422565b60405190151581526020015b60405180910390f35b6101ef61044c565b6040516101de9190611b7e565b61020f61020a366004611b90565b6104db565b6040516001600160a01b0390911681526020016101de565b61023a610235366004611bc2565b610500565b005b61023a61024a366004611bea565b610619565b6101d261025d366004611ca8565b610643565b61023a610270366004611d13565b6106a7565b600c5461020f906001600160a01b031681565b600e5461020f906001600160a01b031681565b6101ef6102a9366004611b90565b6106d8565b61023a6102bc366004611d4c565b6106e3565b61023a6102cf366004611d13565b61089b565b61023a6102e2366004611bea565b6108b5565b6102fa6102f5366004611dbb565b6108df565b6040519081526020016101de565b61023a610316366004611e06565b610981565b61023a610329366004611e78565b610cd6565b600d5461020f906001600160a01b031681565b61020f61034f366004611b90565b610d4e565b6102fa610362366004611bea565b610dad565b61023a610e31565b6007546001600160a01b031661020f565b6101ef610e44565b61023a610396366004611eb7565b610e53565b61023a6103a9366004611eec565b610e5e565b6101ef6103bc366004611b90565b610e90565b6102fa6103cf366004611f63565b610f86565b6101d26103e2366004611fbd565b6001600160a01b039182165f90815260056020908152604080832093909416825291909152205460ff1690565b61023a61041d366004611bea565b61106b565b5f6001600160e01b03198216632483248360e11b14806104465750610446826110e4565b92915050565b60605f805461045a90611fee565b80601f016020809104026020016040519081016040528092919081815260200182805461048690611fee565b80156104d15780601f106104a8576101008083540402835291602001916104d1565b820191905f5260205f20905b8154815290600101906020018083116104b457829003601f168201915b5050505050905090565b5f6104e582611133565b505f908152600460205260409020546001600160a01b031690565b5f61050a82610d4e565b9050806001600160a01b0316836001600160a01b03160361057c5760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b60648201526084015b60405180910390fd5b336001600160a01b0382161480610598575061059881336103e2565b61060a5760405162461bcd60e51b815260206004820152603d60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206f7220617070726f76656420666f7220616c6c0000006064820152608401610573565b6106148383611191565b505050565b6106216111fe565b600d80546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b0384165f908152601160205260408082209051610668908690612026565b90815260408051602092819003830190206001600160a01b038681165f908152918452828220908616825290925290205460ff1690505b949350505050565b6106b13382611258565b6106cd5760405162461bcd60e51b815260040161057390612041565b6106148383836112d4565b606061044682610e90565b600d546001600160a01b0316331461070d5760405162461bcd60e51b81526004016105739061208e565b826001600160a01b03166009836040516107279190612026565b908152604051908190036020019020546001600160a01b0316146107b35760405162461bcd60e51b815260206004820152603960248201527f54686973206576616c756174696f6e206a6f6220646f6573206e6f742062656c60448201527f6f6e6720746f207468652064617461207075626c6973686572000000000000006064820152608401610573565b604080516020810182525f9052517fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090600a906107f1908590612026565b908152604051908190036020018120610809916120eb565b60405180910390201461086a5760405162461bcd60e51b8152602060048201526024808201527f54686973206a6f6220616c7265616479206861732061206d617070696e6720746044820152631bc81a5d60e21b6064820152608401610573565b80600a8360405161087b9190612026565b9081526020016040518091039020908161089591906121a8565b50505050565b61061483838360405180602001604052805f815250610e5e565b6108bd6111fe565b600e80546001600160a01b0319166001600160a01b0392909216919091179055565b600d545f906001600160a01b0316331461090b5760405162461bcd60e51b81526004016105739061208e565b610919600f80546001019055565b5f610923600f5490565b905061092f8482611436565b610939818461144f565b8360098460405161094a9190612026565b90815260405190819003602001902080546001600160a01b03929092166001600160a01b0319909216919091179055905092915050565b600d546001600160a01b031633146109ab5760405162461bcd60e51b81526004016105739061208e565b6001600160a01b0385165f908152601160205260409081902090516109d1908690612026565b90815260408051602092819003830190206001600160a01b038581165f908152918452828220908716825290925290205460ff1615610a895760405162461bcd60e51b815260206004820152604860248201527f43616e6e6f742072657375626d697420612073636f726520666f72206120747260448201527f61696e65722074686174206861732070726576696f75736c79206265656e20736064820152671d589b5a5d1d195960c21b608482015260a401610573565b6064811115610ada5760405162461bcd60e51b815260206004820152601f60248201527f53636f7265206d757374206265206265747765656e203020616e6420313030006044820152606401610573565b600d546040516303a120eb60e51b81526001600160a01b03909116906374241d6090610b0e90889088908790600401612268565b602060405180830381865afa158015610b29573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b4d919061229c565b610bbf5760405162461bcd60e51b815260206004820152603760248201527f596f7520617265206e6f7420616e2061737369676e65642076616c696461746f60448201527f7220666f72207468697320747261696e696e67206a6f620000000000000000006064820152608401610573565b6001600160a01b0385165f90815260106020526040908190209051829190610be8908790612026565b9081526040805191829003602090810183206001600160a01b038088165f9081529183528382208982168352835283822095909555938916845260129052822090610c34908790612026565b9081526040805191829003602090810183206001600160a01b038089165f908152918352838220805460018181018355828552858520909101899055918c16835260119093529290209093509091610c8d908890612026565b90815260408051602092819003830190206001600160a01b039687165f90815290835281812097909616865295905293909220805460ff19169315159390931790925550505050565b600c546001600160a01b03163314610d3e5760405162461bcd60e51b815260206004820152602560248201527f4f6e6c7920746865206f776e65722063616e206163636573732074686973206d604482015264195d1a1bd960da1b6064820152608401610573565b600b610d4a82826121a8565b5050565b5f818152600260205260408120546001600160a01b0316806104465760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b6044820152606401610573565b5f6001600160a01b038216610e165760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b6064820152608401610573565b506001600160a01b03165f9081526003602052604090205490565b610e396111fe565b610e425f611459565b565b60606001805461045a90611fee565b610d4a3383836114aa565b610e683383611258565b610e845760405162461bcd60e51b815260040161057390612041565b61089584848484611577565b6060610e9b82611133565b5f8281526006602052604081208054610eb390611fee565b80601f0160208091040260200160405190810160405280929190818152602001828054610edf90611fee565b8015610f2a5780601f10610f0157610100808354040283529160200191610f2a565b820191905f5260205f20905b815481529060010190602001808311610f0d57829003601f168201915b505050505090505f610f3a6115aa565b905080515f03610f4b575092915050565b815115610f7d578082604051602001610f659291906122b7565b60405160208183030381529060405292505050919050565b61069f846115b9565b6001600160a01b0383165f908152601260205260408082209051829190610fae908690612026565b9081526040805191829003602090810183206001600160a01b0387165f908152908252829020805480830285018301909352828452919083018282801561101257602002820191905f5260205f20905b815481526020019060010190808311610ffe575b505050505090505f805b825181101561105457828181518110611037576110376122e5565b60200260200101518261104a91906122f9565b915060010161101c565b5081516110619082612318565b9695505050505050565b6110736111fe565b6001600160a01b0381166110d85760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610573565b6110e181611459565b50565b5f6001600160e01b031982166380ac58cd60e01b148061111457506001600160e01b03198216635b5e139f60e01b145b8061044657506301ffc9a760e01b6001600160e01b0319831614610446565b5f818152600260205260409020546001600160a01b03166110e15760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b6044820152606401610573565b5f81815260046020526040902080546001600160a01b0319166001600160a01b03841690811790915581906111c582610d4e565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6007546001600160a01b03163314610e425760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610573565b5f8061126383610d4e565b9050806001600160a01b0316846001600160a01b031614806112a957506001600160a01b038082165f9081526005602090815260408083209388168352929052205460ff165b8061069f5750836001600160a01b03166112c2846104db565b6001600160a01b031614949350505050565b826001600160a01b03166112e782610d4e565b6001600160a01b03161461130d5760405162461bcd60e51b815260040161057390612337565b6001600160a01b03821661136f5760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b6064820152608401610573565b826001600160a01b031661138282610d4e565b6001600160a01b0316146113a85760405162461bcd60e51b815260040161057390612337565b5f81815260046020908152604080832080546001600160a01b03199081169091556001600160a01b038781168086526003855283862080545f1901905590871680865283862080546001019055868652600290945282852080549092168417909155905184937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b610d4a828260405180602001604052805f81525061161d565b610d4a828261164f565b600780546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b816001600160a01b0316836001600160a01b03160361150b5760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c6572000000000000006044820152606401610573565b6001600160a01b038381165f81815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6115828484846112d4565b61158e84848484611718565b6108955760405162461bcd60e51b81526004016105739061237c565b6060600b805461045a90611fee565b60606115c482611133565b5f6115cd6115aa565b90505f8151116115eb5760405180602001604052805f815250611616565b806115f584611812565b6040516020016116069291906122b7565b6040516020818303038152906040525b9392505050565b61162783836118a2565b6116335f848484611718565b6106145760405162461bcd60e51b81526004016105739061237c565b5f828152600260205260409020546001600160a01b03166116c95760405162461bcd60e51b815260206004820152602e60248201527f45524337323155524953746f726167653a2055524920736574206f66206e6f6e60448201526d32bc34b9ba32b73a103a37b5b2b760911b6064820152608401610573565b5f8281526006602052604090206116e082826121a8565b506040518281527ff8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce79060200160405180910390a15050565b5f6001600160a01b0384163b1561180a57604051630a85bd0160e11b81526001600160a01b0385169063150b7a029061175b9033908990889088906004016123ce565b6020604051808303815f875af1925050508015611795575060408051601f3d908101601f1916820190925261179291810190612400565b60015b6117f0573d8080156117c2576040519150601f19603f3d011682016040523d82523d5f602084013e6117c7565b606091505b5080515f036117e85760405162461bcd60e51b81526004016105739061237c565b805181602001fd5b6001600160e01b031916630a85bd0160e11b14905061069f565b50600161069f565b60605f61181e83611a2a565b60010190505f8167ffffffffffffffff81111561183d5761183d611c03565b6040519080825280601f01601f191660200182016040528015611867576020820181803683370190505b5090508181016020015b5f19016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a850494508461187157509392505050565b6001600160a01b0382166118f85760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f20616464726573736044820152606401610573565b5f818152600260205260409020546001600160a01b03161561195c5760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610573565b5f818152600260205260409020546001600160a01b0316156119c05760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610573565b6001600160a01b0382165f81815260036020908152604080832080546001019055848352600290915280822080546001600160a01b0319168417905551839291907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b5f8072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b8310611a685772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef81000000008310611a94576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc100008310611ab257662386f26fc10000830492506010015b6305f5e1008310611aca576305f5e100830492506008015b6127108310611ade57612710830492506004015b60648310611af0576064830492506002015b600a83106104465760010192915050565b6001600160e01b0319811681146110e1575f80fd5b5f60208284031215611b26575f80fd5b813561161681611b01565b5f5b83811015611b4b578181015183820152602001611b33565b50505f910152565b5f8151808452611b6a816020860160208601611b31565b601f01601f19169290920160200192915050565b602081525f6116166020830184611b53565b5f60208284031215611ba0575f80fd5b5035919050565b80356001600160a01b0381168114611bbd575f80fd5b919050565b5f8060408385031215611bd3575f80fd5b611bdc83611ba7565b946020939093013593505050565b5f60208284031215611bfa575f80fd5b61161682611ba7565b634e487b7160e01b5f52604160045260245ffd5b5f67ffffffffffffffff80841115611c3157611c31611c03565b604051601f8501601f19908116603f01168101908282118183101715611c5957611c59611c03565b81604052809350858152868686011115611c71575f80fd5b858560208301375f602087830101525050509392505050565b5f82601f830112611c99575f80fd5b61161683833560208501611c17565b5f805f8060808587031215611cbb575f80fd5b611cc485611ba7565b9350602085013567ffffffffffffffff811115611cdf575f80fd5b611ceb87828801611c8a565b935050611cfa60408601611ba7565b9150611d0860608601611ba7565b905092959194509250565b5f805f60608486031215611d25575f80fd5b611d2e84611ba7565b9250611d3c60208501611ba7565b9150604084013590509250925092565b5f805f60608486031215611d5e575f80fd5b611d6784611ba7565b9250602084013567ffffffffffffffff80821115611d83575f80fd5b611d8f87838801611c8a565b93506040860135915080821115611da4575f80fd5b50611db186828701611c8a565b9150509250925092565b5f8060408385031215611dcc575f80fd5b611dd583611ba7565b9150602083013567ffffffffffffffff811115611df0575f80fd5b611dfc85828601611c8a565b9150509250929050565b5f805f805f60a08688031215611e1a575f80fd5b611e2386611ba7565b9450602086013567ffffffffffffffff811115611e3e575f80fd5b611e4a88828901611c8a565b945050611e5960408701611ba7565b9250611e6760608701611ba7565b949793965091946080013592915050565b5f60208284031215611e88575f80fd5b813567ffffffffffffffff811115611e9e575f80fd5b61069f84828501611c8a565b80151581146110e1575f80fd5b5f8060408385031215611ec8575f80fd5b611ed183611ba7565b91506020830135611ee181611eaa565b809150509250929050565b5f805f8060808587031215611eff575f80fd5b611f0885611ba7565b9350611f1660208601611ba7565b925060408501359150606085013567ffffffffffffffff811115611f38575f80fd5b8501601f81018713611f48575f80fd5b611f5787823560208401611c17565b91505092959194509250565b5f805f60608486031215611f75575f80fd5b611f7e84611ba7565b9250602084013567ffffffffffffffff811115611f99575f80fd5b611fa586828701611c8a565b925050611fb460408501611ba7565b90509250925092565b5f8060408385031215611fce575f80fd5b611fd783611ba7565b9150611fe560208401611ba7565b90509250929050565b600181811c9082168061200257607f821691505b60208210810361202057634e487b7160e01b5f52602260045260245ffd5b50919050565b5f8251612037818460208701611b31565b9190910192915050565b6020808252602d908201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560408201526c1c881bdc88185c1c1c9bdd9959609a1b606082015260800190565b6020808252603f908201527f54686973206d6574686f642063616e206f6e6c792062652063616c6c6564206260408201527f792074686520736361747465722070726f746f636f6c20636f6e747261637400606082015260800190565b5f8083546120f881611fee565b60018281168015612110576001811461212557612151565b60ff1984168752821515830287019450612151565b875f526020805f205f5b858110156121485781548a82015290840190820161212f565b50505082870194505b50929695505050505050565b601f82111561061457805f5260205f20601f840160051c810160208510156121825750805b601f840160051c820191505b818110156121a1575f815560010161218e565b5050505050565b815167ffffffffffffffff8111156121c2576121c2611c03565b6121d6816121d08454611fee565b8461215d565b602080601f831160018114612209575f84156121f25750858301515b5f19600386901b1c1916600185901b178555612260565b5f85815260208120601f198616915b8281101561223757888601518255948401946001909101908401612218565b508582101561225457878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b5f60018060a01b038086168352606060208401526122896060840186611b53565b9150808416604084015250949350505050565b5f602082840312156122ac575f80fd5b815161161681611eaa565b5f83516122c8818460208801611b31565b8351908301906122dc818360208801611b31565b01949350505050565b634e487b7160e01b5f52603260045260245ffd5b8082018082111561044657634e487b7160e01b5f52601160045260245ffd5b5f8261233257634e487b7160e01b5f52601260045260245ffd5b500490565b60208082526025908201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060408201526437bbb732b960d91b606082015260800190565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b6001600160a01b03858116825284166020820152604081018390526080606082018190525f9061106190830184611b53565b5f60208284031215612410575f80fd5b815161161681611b0156fea2646970667358221220d303416102794a494e1c1feb7f242140a0a695f0ee2ab9d29f40b54411aa5a8264736f6c63430008170033",
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
