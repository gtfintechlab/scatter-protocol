// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package modeltoken

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

// ModeltokenMetaData contains all meta data concerning the Modeltoken contract.
var ModeltokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"trainer\",\"type\":\"address\"}],\"name\":\"getModelCidForTrainer\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getTokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"modelLogger\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolDeployer\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"requestorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topicName\",\"type\":\"string\"}],\"name\":\"publishModel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scatterContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setScatterContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b506040518060400160405280601781526020017f536361747465722050726f746f636f6c204d6f64656c730000000000000000008152506040518060400160405280600381526020016253504d60e81b815250815f9081620000739190620001aa565b506001620000828282620001aa565b5050506200009f62000099620000b760201b60201c565b620000bb565b600a80546001600160a01b0319163317905562000276565b3390565b600780546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b634e487b7160e01b5f52604160045260245ffd5b600181811c908216806200013557607f821691505b6020821081036200015457634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115620001a557805f5260205f20601f840160051c81016020851015620001815750805b601f840160051c820191505b81811015620001a2575f81556001016200018d565b50505b505050565b81516001600160401b03811115620001c657620001c66200010c565b620001de81620001d7845462000120565b846200015a565b602080601f83116001811462000214575f8415620001fc5750858301515b5f19600386901b1c1916600185901b1785556200026e565b5f85815260208120601f198616915b82811015620002445788860151825594840194600190910190840162000223565b50858210156200026257878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b611c3880620002845f395ff3fe608060405234801561000f575f80fd5b5060043610610153575f3560e01c80635d14926a116100bf57806395d89b411161007957806395d89b41146102df578063a22cb465146102e7578063b88d4fde146102fa578063c87b56dd1461030d578063e985e9c514610320578063f2fde38b1461035b575f80fd5b80635d14926a1461027a5780636352211e1461028d57806370a08231146102a0578063715018a6146102b35780638672c212146102bb5780638da5cb5b146102ce575f80fd5b8063330c093c11610110578063330c093c146101fa57806333f6b0061461020d5780633bb3a24d1461022e57806342842e0e1461024157806355f804b3146102545780635c11037e14610267575f80fd5b806301ffc9a71461015757806306fdde031461017f578063081812fc14610194578063095ea7b3146101bf578063126d5b6a146101d457806323b872dd146101e7575b5f80fd5b61016a61016536600461158f565b61036e565b60405190151581526020015b60405180910390f35b610187610398565b60405161017691906115f7565b6101a76101a2366004611609565b610427565b6040516001600160a01b039091168152602001610176565b6101d26101cd36600461163b565b61044c565b005b6101d26101e2366004611663565b610565565b6101d26101f536600461167c565b61058f565b600a546101a7906001600160a01b031681565b61022061021b36600461175a565b6105c0565b604051908152602001610176565b61018761023c366004611609565b6106d3565b6101d261024f36600461167c565b6106de565b6101d26102623660046117da565b6106f8565b61018761027536600461180c565b610710565b61018761028836600461180c565b6107c9565b6101a761029b366004611609565b6108a6565b6102206102ae366004611663565b610905565b6101d2610989565b600b546101a7906001600160a01b031681565b6007546001600160a01b03166101a7565b61018761099c565b6101d26102f5366004611866565b6109ab565b6101d261030836600461189f565b6109b6565b61018761031b366004611609565b6109ee565b61016a61032e36600461190a565b6001600160a01b039182165f90815260056020908152604080832093909416825291909152205460ff1690565b6101d2610369366004611663565b610ae4565b5f6001600160e01b03198216632483248360e11b1480610392575061039282610b5d565b92915050565b60605f80546103a69061193b565b80601f01602080910402602001604051908101604052809291908181526020018280546103d29061193b565b801561041d5780601f106103f45761010080835404028352916020019161041d565b820191905f5260205f20905b81548152906001019060200180831161040057829003601f168201915b5050505050905090565b5f61043182610bac565b505f908152600460205260409020546001600160a01b031690565b5f610456826108a6565b9050806001600160a01b0316836001600160a01b0316036104c85760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b60648201526084015b60405180910390fd5b336001600160a01b03821614806104e457506104e4813361032e565b6105565760405162461bcd60e51b815260206004820152603d60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206f7220617070726f76656420666f7220616c6c00000060648201526084016104bf565b6105608383610c0a565b505050565b61056d610c77565b600b80546001600160a01b0319166001600160a01b0392909216919091179055565b6105993382610cd1565b6105b55760405162461bcd60e51b81526004016104bf90611973565b610560838383610d4d565b600b545f906001600160a01b031633146106425760405162461bcd60e51b815260206004820152603f60248201527f54686973206d6574686f642063616e206f6e6c792062652063616c6c6564206260448201527f792074686520736361747465722070726f746f636f6c20636f6e74726163740060648201526084016104bf565b610650600c80546001019055565b5f61065a600c5490565b90506106668582610eaf565b6106708187610ec8565b6001600160a01b0384165f908152600d60205260409081902090518791906106999086906119c0565b90815260408051602092819003830190206001600160a01b0389165f90815292529020906106c79082611a26565b5090505b949350505050565b6060610392826109ee565b61056083838360405180602001604052805f8152506109b6565b610700610c77565b600961070c8282611a26565b5050565b600d60209081525f938452604080852084518086018401805192815290840195840195909520945292905282529020805461074a9061193b565b80601f01602080910402602001604051908101604052809291908181526020018280546107769061193b565b80156107c15780601f10610798576101008083540402835291602001916107c1565b820191905f5260205f20905b8154815290600101906020018083116107a457829003601f168201915b505050505081565b6001600160a01b0383165f908152600d6020526040908190209051606091906107f39085906119c0565b90815260408051602092819003830190206001600160a01b0385165f9081529252902080546108219061193b565b80601f016020809104026020016040519081016040528092919081815260200182805461084d9061193b565b80156108985780601f1061086f57610100808354040283529160200191610898565b820191905f5260205f20905b81548152906001019060200180831161087b57829003601f168201915b505050505090509392505050565b5f818152600260205260408120546001600160a01b0316806103925760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b60448201526064016104bf565b5f6001600160a01b03821661096e5760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b60648201526084016104bf565b506001600160a01b03165f9081526003602052604090205490565b610991610c77565b61099a5f610ed2565b565b6060600180546103a69061193b565b61070c338383610f23565b6109c03383610cd1565b6109dc5760405162461bcd60e51b81526004016104bf90611973565b6109e884848484610ff0565b50505050565b60606109f982610bac565b5f8281526006602052604081208054610a119061193b565b80601f0160208091040260200160405190810160405280929190818152602001828054610a3d9061193b565b8015610a885780601f10610a5f57610100808354040283529160200191610a88565b820191905f5260205f20905b815481529060010190602001808311610a6b57829003601f168201915b505050505090505f610a98611023565b905080515f03610aa9575092915050565b815115610adb578082604051602001610ac3929190611ae6565b60405160208183030381529060405292505050919050565b6106cb84611032565b610aec610c77565b6001600160a01b038116610b515760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016104bf565b610b5a81610ed2565b50565b5f6001600160e01b031982166380ac58cd60e01b1480610b8d57506001600160e01b03198216635b5e139f60e01b145b8061039257506301ffc9a760e01b6001600160e01b0319831614610392565b5f818152600260205260409020546001600160a01b0316610b5a5760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b60448201526064016104bf565b5f81815260046020526040902080546001600160a01b0319166001600160a01b0384169081179091558190610c3e826108a6565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6007546001600160a01b0316331461099a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104bf565b5f80610cdc836108a6565b9050806001600160a01b0316846001600160a01b03161480610d2257506001600160a01b038082165f9081526005602090815260408083209388168352929052205460ff165b806106cb5750836001600160a01b0316610d3b84610427565b6001600160a01b031614949350505050565b826001600160a01b0316610d60826108a6565b6001600160a01b031614610d865760405162461bcd60e51b81526004016104bf90611b14565b6001600160a01b038216610de85760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b60648201526084016104bf565b826001600160a01b0316610dfb826108a6565b6001600160a01b031614610e215760405162461bcd60e51b81526004016104bf90611b14565b5f81815260046020908152604080832080546001600160a01b03199081169091556001600160a01b038781168086526003855283862080545f1901905590871680865283862080546001019055868652600290945282852080549092168417909155905184937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b61070c828260405180602001604052805f815250611096565b61070c82826110c8565b600780546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b816001600160a01b0316836001600160a01b031603610f845760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c65720000000000000060448201526064016104bf565b6001600160a01b038381165f81815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b610ffb848484610d4d565b61100784848484611191565b6109e85760405162461bcd60e51b81526004016104bf90611b59565b6060600980546103a69061193b565b606061103d82610bac565b5f611046611023565b90505f8151116110645760405180602001604052805f81525061108f565b8061106e8461128b565b60405160200161107f929190611ae6565b6040516020818303038152906040525b9392505050565b6110a0838361131b565b6110ac5f848484611191565b6105605760405162461bcd60e51b81526004016104bf90611b59565b5f828152600260205260409020546001600160a01b03166111425760405162461bcd60e51b815260206004820152602e60248201527f45524337323155524953746f726167653a2055524920736574206f66206e6f6e60448201526d32bc34b9ba32b73a103a37b5b2b760911b60648201526084016104bf565b5f8281526006602052604090206111598282611a26565b506040518281527ff8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce79060200160405180910390a15050565b5f6001600160a01b0384163b1561128357604051630a85bd0160e11b81526001600160a01b0385169063150b7a02906111d4903390899088908890600401611bab565b6020604051808303815f875af192505050801561120e575060408051601f3d908101601f1916820190925261120b91810190611be7565b60015b611269573d80801561123b576040519150601f19603f3d011682016040523d82523d5f602084013e611240565b606091505b5080515f036112615760405162461bcd60e51b81526004016104bf90611b59565b805181602001fd5b6001600160e01b031916630a85bd0160e11b1490506106cb565b5060016106cb565b60605f611297836114a3565b60010190505f8167ffffffffffffffff8111156112b6576112b66116b5565b6040519080825280601f01601f1916602001820160405280156112e0576020820181803683370190505b5090508181016020015b5f19016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a85049450846112ea57509392505050565b6001600160a01b0382166113715760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f206164647265737360448201526064016104bf565b5f818152600260205260409020546001600160a01b0316156113d55760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e7465640000000060448201526064016104bf565b5f818152600260205260409020546001600160a01b0316156114395760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e7465640000000060448201526064016104bf565b6001600160a01b0382165f81815260036020908152604080832080546001019055848352600290915280822080546001600160a01b0319168417905551839291907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b5f8072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b83106114e15772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef8100000000831061150d576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061152b57662386f26fc10000830492506010015b6305f5e1008310611543576305f5e100830492506008015b612710831061155757612710830492506004015b60648310611569576064830492506002015b600a83106103925760010192915050565b6001600160e01b031981168114610b5a575f80fd5b5f6020828403121561159f575f80fd5b813561108f8161157a565b5f5b838110156115c45781810151838201526020016115ac565b50505f910152565b5f81518084526115e38160208601602086016115aa565b601f01601f19169290920160200192915050565b602081525f61108f60208301846115cc565b5f60208284031215611619575f80fd5b5035919050565b80356001600160a01b0381168114611636575f80fd5b919050565b5f806040838503121561164c575f80fd5b61165583611620565b946020939093013593505050565b5f60208284031215611673575f80fd5b61108f82611620565b5f805f6060848603121561168e575f80fd5b61169784611620565b92506116a560208501611620565b9150604084013590509250925092565b634e487b7160e01b5f52604160045260245ffd5b5f67ffffffffffffffff808411156116e3576116e36116b5565b604051601f8501601f19908116603f0116810190828211818310171561170b5761170b6116b5565b81604052809350858152868686011115611723575f80fd5b858560208301375f602087830101525050509392505050565b5f82601f83011261174b575f80fd5b61108f838335602085016116c9565b5f805f806080858703121561176d575f80fd5b843567ffffffffffffffff80821115611784575f80fd5b6117908883890161173c565b955061179e60208801611620565b94506117ac60408801611620565b935060608701359150808211156117c1575f80fd5b506117ce8782880161173c565b91505092959194509250565b5f602082840312156117ea575f80fd5b813567ffffffffffffffff811115611800575f80fd5b6106cb8482850161173c565b5f805f6060848603121561181e575f80fd5b61182784611620565b9250602084013567ffffffffffffffff811115611842575f80fd5b61184e8682870161173c565b92505061185d60408501611620565b90509250925092565b5f8060408385031215611877575f80fd5b61188083611620565b915060208301358015158114611894575f80fd5b809150509250929050565b5f805f80608085870312156118b2575f80fd5b6118bb85611620565b93506118c960208601611620565b925060408501359150606085013567ffffffffffffffff8111156118eb575f80fd5b8501601f810187136118fb575f80fd5b6117ce878235602084016116c9565b5f806040838503121561191b575f80fd5b61192483611620565b915061193260208401611620565b90509250929050565b600181811c9082168061194f57607f821691505b60208210810361196d57634e487b7160e01b5f52602260045260245ffd5b50919050565b6020808252602d908201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560408201526c1c881bdc88185c1c1c9bdd9959609a1b606082015260800190565b5f82516119d18184602087016115aa565b9190910192915050565b601f82111561056057805f5260205f20601f840160051c81016020851015611a005750805b601f840160051c820191505b81811015611a1f575f8155600101611a0c565b5050505050565b815167ffffffffffffffff811115611a4057611a406116b5565b611a5481611a4e845461193b565b846119db565b602080601f831160018114611a87575f8415611a705750858301515b5f19600386901b1c1916600185901b178555611ade565b5f85815260208120601f198616915b82811015611ab557888601518255948401946001909101908401611a96565b5085821015611ad257878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b5f8351611af78184602088016115aa565b835190830190611b0b8183602088016115aa565b01949350505050565b60208082526025908201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060408201526437bbb732b960d91b606082015260800190565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b6001600160a01b03858116825284166020820152604081018390526080606082018190525f90611bdd908301846115cc565b9695505050505050565b5f60208284031215611bf7575f80fd5b815161108f8161157a56fea26469706673582212200114358f718deeda6c7ad696876fa1412842b3c359f6128e3e290a09da5aeb7964736f6c63430008170033",
}

// ModeltokenABI is the input ABI used to generate the binding from.
// Deprecated: Use ModeltokenMetaData.ABI instead.
var ModeltokenABI = ModeltokenMetaData.ABI

// ModeltokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ModeltokenMetaData.Bin instead.
var ModeltokenBin = ModeltokenMetaData.Bin

// DeployModeltoken deploys a new Ethereum contract, binding an instance of Modeltoken to it.
func DeployModeltoken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Modeltoken, error) {
	parsed, err := ModeltokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ModeltokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Modeltoken{ModeltokenCaller: ModeltokenCaller{contract: contract}, ModeltokenTransactor: ModeltokenTransactor{contract: contract}, ModeltokenFilterer: ModeltokenFilterer{contract: contract}}, nil
}

// Modeltoken is an auto generated Go binding around an Ethereum contract.
type Modeltoken struct {
	ModeltokenCaller     // Read-only binding to the contract
	ModeltokenTransactor // Write-only binding to the contract
	ModeltokenFilterer   // Log filterer for contract events
}

// ModeltokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ModeltokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModeltokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ModeltokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModeltokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ModeltokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModeltokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ModeltokenSession struct {
	Contract     *Modeltoken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModeltokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ModeltokenCallerSession struct {
	Contract *ModeltokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ModeltokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ModeltokenTransactorSession struct {
	Contract     *ModeltokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ModeltokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ModeltokenRaw struct {
	Contract *Modeltoken // Generic contract binding to access the raw methods on
}

// ModeltokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ModeltokenCallerRaw struct {
	Contract *ModeltokenCaller // Generic read-only contract binding to access the raw methods on
}

// ModeltokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ModeltokenTransactorRaw struct {
	Contract *ModeltokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewModeltoken creates a new instance of Modeltoken, bound to a specific deployed contract.
func NewModeltoken(address common.Address, backend bind.ContractBackend) (*Modeltoken, error) {
	contract, err := bindModeltoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Modeltoken{ModeltokenCaller: ModeltokenCaller{contract: contract}, ModeltokenTransactor: ModeltokenTransactor{contract: contract}, ModeltokenFilterer: ModeltokenFilterer{contract: contract}}, nil
}

// NewModeltokenCaller creates a new read-only instance of Modeltoken, bound to a specific deployed contract.
func NewModeltokenCaller(address common.Address, caller bind.ContractCaller) (*ModeltokenCaller, error) {
	contract, err := bindModeltoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ModeltokenCaller{contract: contract}, nil
}

// NewModeltokenTransactor creates a new write-only instance of Modeltoken, bound to a specific deployed contract.
func NewModeltokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ModeltokenTransactor, error) {
	contract, err := bindModeltoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ModeltokenTransactor{contract: contract}, nil
}

// NewModeltokenFilterer creates a new log filterer instance of Modeltoken, bound to a specific deployed contract.
func NewModeltokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ModeltokenFilterer, error) {
	contract, err := bindModeltoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ModeltokenFilterer{contract: contract}, nil
}

// bindModeltoken binds a generic wrapper to an already deployed contract.
func bindModeltoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ModeltokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Modeltoken *ModeltokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Modeltoken.Contract.ModeltokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Modeltoken *ModeltokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Modeltoken.Contract.ModeltokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Modeltoken *ModeltokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Modeltoken.Contract.ModeltokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Modeltoken *ModeltokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Modeltoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Modeltoken *ModeltokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Modeltoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Modeltoken *ModeltokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Modeltoken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Modeltoken *ModeltokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Modeltoken.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Modeltoken *ModeltokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Modeltoken.Contract.BalanceOf(&_Modeltoken.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Modeltoken *ModeltokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Modeltoken.Contract.BalanceOf(&_Modeltoken.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Modeltoken *ModeltokenCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Modeltoken.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Modeltoken *ModeltokenSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Modeltoken.Contract.GetApproved(&_Modeltoken.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Modeltoken *ModeltokenCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Modeltoken.Contract.GetApproved(&_Modeltoken.CallOpts, tokenId)
}

// GetModelCidForTrainer is a free data retrieval call binding the contract method 0x5d14926a.
//
// Solidity: function getModelCidForTrainer(address requestorAddress, string topicName, address trainer) view returns(string)
func (_Modeltoken *ModeltokenCaller) GetModelCidForTrainer(opts *bind.CallOpts, requestorAddress common.Address, topicName string, trainer common.Address) (string, error) {
	var out []interface{}
	err := _Modeltoken.contract.Call(opts, &out, "getModelCidForTrainer", requestorAddress, topicName, trainer)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetModelCidForTrainer is a free data retrieval call binding the contract method 0x5d14926a.
//
// Solidity: function getModelCidForTrainer(address requestorAddress, string topicName, address trainer) view returns(string)
func (_Modeltoken *ModeltokenSession) GetModelCidForTrainer(requestorAddress common.Address, topicName string, trainer common.Address) (string, error) {
	return _Modeltoken.Contract.GetModelCidForTrainer(&_Modeltoken.CallOpts, requestorAddress, topicName, trainer)
}

// GetModelCidForTrainer is a free data retrieval call binding the contract method 0x5d14926a.
//
// Solidity: function getModelCidForTrainer(address requestorAddress, string topicName, address trainer) view returns(string)
func (_Modeltoken *ModeltokenCallerSession) GetModelCidForTrainer(requestorAddress common.Address, topicName string, trainer common.Address) (string, error) {
	return _Modeltoken.Contract.GetModelCidForTrainer(&_Modeltoken.CallOpts, requestorAddress, topicName, trainer)
}

// GetTokenURI is a free data retrieval call binding the contract method 0x3bb3a24d.
//
// Solidity: function getTokenURI(uint256 tokenId) view returns(string)
func (_Modeltoken *ModeltokenCaller) GetTokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Modeltoken.contract.Call(opts, &out, "getTokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetTokenURI is a free data retrieval call binding the contract method 0x3bb3a24d.
//
// Solidity: function getTokenURI(uint256 tokenId) view returns(string)
func (_Modeltoken *ModeltokenSession) GetTokenURI(tokenId *big.Int) (string, error) {
	return _Modeltoken.Contract.GetTokenURI(&_Modeltoken.CallOpts, tokenId)
}

// GetTokenURI is a free data retrieval call binding the contract method 0x3bb3a24d.
//
// Solidity: function getTokenURI(uint256 tokenId) view returns(string)
func (_Modeltoken *ModeltokenCallerSession) GetTokenURI(tokenId *big.Int) (string, error) {
	return _Modeltoken.Contract.GetTokenURI(&_Modeltoken.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Modeltoken *ModeltokenCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Modeltoken.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Modeltoken *ModeltokenSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Modeltoken.Contract.IsApprovedForAll(&_Modeltoken.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Modeltoken *ModeltokenCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Modeltoken.Contract.IsApprovedForAll(&_Modeltoken.CallOpts, owner, operator)
}

// ModelLogger is a free data retrieval call binding the contract method 0x5c11037e.
//
// Solidity: function modelLogger(address , string , address ) view returns(string)
func (_Modeltoken *ModeltokenCaller) ModelLogger(opts *bind.CallOpts, arg0 common.Address, arg1 string, arg2 common.Address) (string, error) {
	var out []interface{}
	err := _Modeltoken.contract.Call(opts, &out, "modelLogger", arg0, arg1, arg2)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ModelLogger is a free data retrieval call binding the contract method 0x5c11037e.
//
// Solidity: function modelLogger(address , string , address ) view returns(string)
func (_Modeltoken *ModeltokenSession) ModelLogger(arg0 common.Address, arg1 string, arg2 common.Address) (string, error) {
	return _Modeltoken.Contract.ModelLogger(&_Modeltoken.CallOpts, arg0, arg1, arg2)
}

// ModelLogger is a free data retrieval call binding the contract method 0x5c11037e.
//
// Solidity: function modelLogger(address , string , address ) view returns(string)
func (_Modeltoken *ModeltokenCallerSession) ModelLogger(arg0 common.Address, arg1 string, arg2 common.Address) (string, error) {
	return _Modeltoken.Contract.ModelLogger(&_Modeltoken.CallOpts, arg0, arg1, arg2)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Modeltoken *ModeltokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Modeltoken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Modeltoken *ModeltokenSession) Name() (string, error) {
	return _Modeltoken.Contract.Name(&_Modeltoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Modeltoken *ModeltokenCallerSession) Name() (string, error) {
	return _Modeltoken.Contract.Name(&_Modeltoken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Modeltoken *ModeltokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Modeltoken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Modeltoken *ModeltokenSession) Owner() (common.Address, error) {
	return _Modeltoken.Contract.Owner(&_Modeltoken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Modeltoken *ModeltokenCallerSession) Owner() (common.Address, error) {
	return _Modeltoken.Contract.Owner(&_Modeltoken.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Modeltoken *ModeltokenCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Modeltoken.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Modeltoken *ModeltokenSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Modeltoken.Contract.OwnerOf(&_Modeltoken.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Modeltoken *ModeltokenCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Modeltoken.Contract.OwnerOf(&_Modeltoken.CallOpts, tokenId)
}

// ProtocolDeployer is a free data retrieval call binding the contract method 0x330c093c.
//
// Solidity: function protocolDeployer() view returns(address)
func (_Modeltoken *ModeltokenCaller) ProtocolDeployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Modeltoken.contract.Call(opts, &out, "protocolDeployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolDeployer is a free data retrieval call binding the contract method 0x330c093c.
//
// Solidity: function protocolDeployer() view returns(address)
func (_Modeltoken *ModeltokenSession) ProtocolDeployer() (common.Address, error) {
	return _Modeltoken.Contract.ProtocolDeployer(&_Modeltoken.CallOpts)
}

// ProtocolDeployer is a free data retrieval call binding the contract method 0x330c093c.
//
// Solidity: function protocolDeployer() view returns(address)
func (_Modeltoken *ModeltokenCallerSession) ProtocolDeployer() (common.Address, error) {
	return _Modeltoken.Contract.ProtocolDeployer(&_Modeltoken.CallOpts)
}

// ScatterContractAddress is a free data retrieval call binding the contract method 0x8672c212.
//
// Solidity: function scatterContractAddress() view returns(address)
func (_Modeltoken *ModeltokenCaller) ScatterContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Modeltoken.contract.Call(opts, &out, "scatterContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ScatterContractAddress is a free data retrieval call binding the contract method 0x8672c212.
//
// Solidity: function scatterContractAddress() view returns(address)
func (_Modeltoken *ModeltokenSession) ScatterContractAddress() (common.Address, error) {
	return _Modeltoken.Contract.ScatterContractAddress(&_Modeltoken.CallOpts)
}

// ScatterContractAddress is a free data retrieval call binding the contract method 0x8672c212.
//
// Solidity: function scatterContractAddress() view returns(address)
func (_Modeltoken *ModeltokenCallerSession) ScatterContractAddress() (common.Address, error) {
	return _Modeltoken.Contract.ScatterContractAddress(&_Modeltoken.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Modeltoken *ModeltokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Modeltoken.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Modeltoken *ModeltokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Modeltoken.Contract.SupportsInterface(&_Modeltoken.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Modeltoken *ModeltokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Modeltoken.Contract.SupportsInterface(&_Modeltoken.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Modeltoken *ModeltokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Modeltoken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Modeltoken *ModeltokenSession) Symbol() (string, error) {
	return _Modeltoken.Contract.Symbol(&_Modeltoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Modeltoken *ModeltokenCallerSession) Symbol() (string, error) {
	return _Modeltoken.Contract.Symbol(&_Modeltoken.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Modeltoken *ModeltokenCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Modeltoken.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Modeltoken *ModeltokenSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Modeltoken.Contract.TokenURI(&_Modeltoken.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Modeltoken *ModeltokenCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Modeltoken.Contract.TokenURI(&_Modeltoken.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Modeltoken *ModeltokenTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Modeltoken.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Modeltoken *ModeltokenSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Modeltoken.Contract.Approve(&_Modeltoken.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Modeltoken *ModeltokenTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Modeltoken.Contract.Approve(&_Modeltoken.TransactOpts, to, tokenId)
}

// PublishModel is a paid mutator transaction binding the contract method 0x33f6b006.
//
// Solidity: function publishModel(string tokenURI, address recipient, address requestorAddress, string topicName) returns(uint256)
func (_Modeltoken *ModeltokenTransactor) PublishModel(opts *bind.TransactOpts, tokenURI string, recipient common.Address, requestorAddress common.Address, topicName string) (*types.Transaction, error) {
	return _Modeltoken.contract.Transact(opts, "publishModel", tokenURI, recipient, requestorAddress, topicName)
}

// PublishModel is a paid mutator transaction binding the contract method 0x33f6b006.
//
// Solidity: function publishModel(string tokenURI, address recipient, address requestorAddress, string topicName) returns(uint256)
func (_Modeltoken *ModeltokenSession) PublishModel(tokenURI string, recipient common.Address, requestorAddress common.Address, topicName string) (*types.Transaction, error) {
	return _Modeltoken.Contract.PublishModel(&_Modeltoken.TransactOpts, tokenURI, recipient, requestorAddress, topicName)
}

// PublishModel is a paid mutator transaction binding the contract method 0x33f6b006.
//
// Solidity: function publishModel(string tokenURI, address recipient, address requestorAddress, string topicName) returns(uint256)
func (_Modeltoken *ModeltokenTransactorSession) PublishModel(tokenURI string, recipient common.Address, requestorAddress common.Address, topicName string) (*types.Transaction, error) {
	return _Modeltoken.Contract.PublishModel(&_Modeltoken.TransactOpts, tokenURI, recipient, requestorAddress, topicName)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Modeltoken *ModeltokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Modeltoken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Modeltoken *ModeltokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _Modeltoken.Contract.RenounceOwnership(&_Modeltoken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Modeltoken *ModeltokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Modeltoken.Contract.RenounceOwnership(&_Modeltoken.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Modeltoken *ModeltokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Modeltoken.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Modeltoken *ModeltokenSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Modeltoken.Contract.SafeTransferFrom(&_Modeltoken.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Modeltoken *ModeltokenTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Modeltoken.Contract.SafeTransferFrom(&_Modeltoken.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Modeltoken *ModeltokenTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Modeltoken.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Modeltoken *ModeltokenSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Modeltoken.Contract.SafeTransferFrom0(&_Modeltoken.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Modeltoken *ModeltokenTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Modeltoken.Contract.SafeTransferFrom0(&_Modeltoken.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Modeltoken *ModeltokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Modeltoken.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Modeltoken *ModeltokenSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Modeltoken.Contract.SetApprovalForAll(&_Modeltoken.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Modeltoken *ModeltokenTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Modeltoken.Contract.SetApprovalForAll(&_Modeltoken.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI_) returns()
func (_Modeltoken *ModeltokenTransactor) SetBaseURI(opts *bind.TransactOpts, baseURI_ string) (*types.Transaction, error) {
	return _Modeltoken.contract.Transact(opts, "setBaseURI", baseURI_)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI_) returns()
func (_Modeltoken *ModeltokenSession) SetBaseURI(baseURI_ string) (*types.Transaction, error) {
	return _Modeltoken.Contract.SetBaseURI(&_Modeltoken.TransactOpts, baseURI_)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI_) returns()
func (_Modeltoken *ModeltokenTransactorSession) SetBaseURI(baseURI_ string) (*types.Transaction, error) {
	return _Modeltoken.Contract.SetBaseURI(&_Modeltoken.TransactOpts, baseURI_)
}

// SetScatterContractAddress is a paid mutator transaction binding the contract method 0x126d5b6a.
//
// Solidity: function setScatterContractAddress(address contractAddress) returns()
func (_Modeltoken *ModeltokenTransactor) SetScatterContractAddress(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _Modeltoken.contract.Transact(opts, "setScatterContractAddress", contractAddress)
}

// SetScatterContractAddress is a paid mutator transaction binding the contract method 0x126d5b6a.
//
// Solidity: function setScatterContractAddress(address contractAddress) returns()
func (_Modeltoken *ModeltokenSession) SetScatterContractAddress(contractAddress common.Address) (*types.Transaction, error) {
	return _Modeltoken.Contract.SetScatterContractAddress(&_Modeltoken.TransactOpts, contractAddress)
}

// SetScatterContractAddress is a paid mutator transaction binding the contract method 0x126d5b6a.
//
// Solidity: function setScatterContractAddress(address contractAddress) returns()
func (_Modeltoken *ModeltokenTransactorSession) SetScatterContractAddress(contractAddress common.Address) (*types.Transaction, error) {
	return _Modeltoken.Contract.SetScatterContractAddress(&_Modeltoken.TransactOpts, contractAddress)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Modeltoken *ModeltokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Modeltoken.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Modeltoken *ModeltokenSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Modeltoken.Contract.TransferFrom(&_Modeltoken.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Modeltoken *ModeltokenTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Modeltoken.Contract.TransferFrom(&_Modeltoken.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Modeltoken *ModeltokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Modeltoken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Modeltoken *ModeltokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Modeltoken.Contract.TransferOwnership(&_Modeltoken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Modeltoken *ModeltokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Modeltoken.Contract.TransferOwnership(&_Modeltoken.TransactOpts, newOwner)
}

// ModeltokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Modeltoken contract.
type ModeltokenApprovalIterator struct {
	Event *ModeltokenApproval // Event containing the contract specifics and raw log

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
func (it *ModeltokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModeltokenApproval)
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
		it.Event = new(ModeltokenApproval)
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
func (it *ModeltokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModeltokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModeltokenApproval represents a Approval event raised by the Modeltoken contract.
type ModeltokenApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Modeltoken *ModeltokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ModeltokenApprovalIterator, error) {

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

	logs, sub, err := _Modeltoken.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ModeltokenApprovalIterator{contract: _Modeltoken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Modeltoken *ModeltokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ModeltokenApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Modeltoken.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModeltokenApproval)
				if err := _Modeltoken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Modeltoken *ModeltokenFilterer) ParseApproval(log types.Log) (*ModeltokenApproval, error) {
	event := new(ModeltokenApproval)
	if err := _Modeltoken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModeltokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Modeltoken contract.
type ModeltokenApprovalForAllIterator struct {
	Event *ModeltokenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ModeltokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModeltokenApprovalForAll)
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
		it.Event = new(ModeltokenApprovalForAll)
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
func (it *ModeltokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModeltokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModeltokenApprovalForAll represents a ApprovalForAll event raised by the Modeltoken contract.
type ModeltokenApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Modeltoken *ModeltokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ModeltokenApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Modeltoken.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ModeltokenApprovalForAllIterator{contract: _Modeltoken.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Modeltoken *ModeltokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ModeltokenApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Modeltoken.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModeltokenApprovalForAll)
				if err := _Modeltoken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Modeltoken *ModeltokenFilterer) ParseApprovalForAll(log types.Log) (*ModeltokenApprovalForAll, error) {
	event := new(ModeltokenApprovalForAll)
	if err := _Modeltoken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModeltokenBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the Modeltoken contract.
type ModeltokenBatchMetadataUpdateIterator struct {
	Event *ModeltokenBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *ModeltokenBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModeltokenBatchMetadataUpdate)
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
		it.Event = new(ModeltokenBatchMetadataUpdate)
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
func (it *ModeltokenBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModeltokenBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModeltokenBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the Modeltoken contract.
type ModeltokenBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Modeltoken *ModeltokenFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*ModeltokenBatchMetadataUpdateIterator, error) {

	logs, sub, err := _Modeltoken.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &ModeltokenBatchMetadataUpdateIterator{contract: _Modeltoken.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Modeltoken *ModeltokenFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *ModeltokenBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Modeltoken.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModeltokenBatchMetadataUpdate)
				if err := _Modeltoken.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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
func (_Modeltoken *ModeltokenFilterer) ParseBatchMetadataUpdate(log types.Log) (*ModeltokenBatchMetadataUpdate, error) {
	event := new(ModeltokenBatchMetadataUpdate)
	if err := _Modeltoken.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModeltokenMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the Modeltoken contract.
type ModeltokenMetadataUpdateIterator struct {
	Event *ModeltokenMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *ModeltokenMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModeltokenMetadataUpdate)
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
		it.Event = new(ModeltokenMetadataUpdate)
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
func (it *ModeltokenMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModeltokenMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModeltokenMetadataUpdate represents a MetadataUpdate event raised by the Modeltoken contract.
type ModeltokenMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Modeltoken *ModeltokenFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*ModeltokenMetadataUpdateIterator, error) {

	logs, sub, err := _Modeltoken.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &ModeltokenMetadataUpdateIterator{contract: _Modeltoken.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Modeltoken *ModeltokenFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *ModeltokenMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Modeltoken.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModeltokenMetadataUpdate)
				if err := _Modeltoken.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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
func (_Modeltoken *ModeltokenFilterer) ParseMetadataUpdate(log types.Log) (*ModeltokenMetadataUpdate, error) {
	event := new(ModeltokenMetadataUpdate)
	if err := _Modeltoken.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModeltokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Modeltoken contract.
type ModeltokenOwnershipTransferredIterator struct {
	Event *ModeltokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ModeltokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModeltokenOwnershipTransferred)
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
		it.Event = new(ModeltokenOwnershipTransferred)
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
func (it *ModeltokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModeltokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModeltokenOwnershipTransferred represents a OwnershipTransferred event raised by the Modeltoken contract.
type ModeltokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Modeltoken *ModeltokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ModeltokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Modeltoken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ModeltokenOwnershipTransferredIterator{contract: _Modeltoken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Modeltoken *ModeltokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ModeltokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Modeltoken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModeltokenOwnershipTransferred)
				if err := _Modeltoken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Modeltoken *ModeltokenFilterer) ParseOwnershipTransferred(log types.Log) (*ModeltokenOwnershipTransferred, error) {
	event := new(ModeltokenOwnershipTransferred)
	if err := _Modeltoken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModeltokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Modeltoken contract.
type ModeltokenTransferIterator struct {
	Event *ModeltokenTransfer // Event containing the contract specifics and raw log

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
func (it *ModeltokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModeltokenTransfer)
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
		it.Event = new(ModeltokenTransfer)
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
func (it *ModeltokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModeltokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModeltokenTransfer represents a Transfer event raised by the Modeltoken contract.
type ModeltokenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Modeltoken *ModeltokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ModeltokenTransferIterator, error) {

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

	logs, sub, err := _Modeltoken.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ModeltokenTransferIterator{contract: _Modeltoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Modeltoken *ModeltokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ModeltokenTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Modeltoken.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModeltokenTransfer)
				if err := _Modeltoken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Modeltoken *ModeltokenFilterer) ParseTransfer(log types.Log) (*ModeltokenTransfer, error) {
	event := new(ModeltokenTransfer)
	if err := _Modeltoken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
