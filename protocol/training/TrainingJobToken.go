// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package trainingtoken

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

// TrainingtokenMetaData contains all meta data concerning the Trainingtoken contract.
var TrainingtokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getTokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolDeployer\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"publishTrainingJob\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scatterContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setScatterContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b506040518060400160405280601e81526020017f536361747465722050726f746f636f6c20547261696e696e67204a6f627300008152506040518060400160405280600481526020016329a82a2560e11b815250815f9081620000749190620001ab565b506001620000838282620001ab565b505050620000a06200009a620000b860201b60201c565b620000bc565b600a80546001600160a01b0319163317905562000277565b3390565b600780546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b634e487b7160e01b5f52604160045260245ffd5b600181811c908216806200013657607f821691505b6020821081036200015557634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115620001a657805f5260205f20601f840160051c81016020851015620001825750805b601f840160051c820191505b81811015620001a3575f81556001016200018e565b50505b505050565b81516001600160401b03811115620001c757620001c76200010d565b620001df81620001d8845462000121565b846200015b565b602080601f83116001811462000215575f8415620001fd5750858301515b5f19600386901b1c1916600185901b1785556200026f565b5f85815260208120601f198616915b82811015620002455788860151825594840194600190910190840162000224565b50858210156200026357878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b61196d80620002855f395ff3fe608060405234801561000f575f80fd5b506004361061013d575f3560e01c80636838a4f6116100b457806395d89b411161007957806395d89b41146102a3578063a22cb465146102ab578063b88d4fde146102be578063c87b56dd146102d1578063e985e9c5146102e4578063f2fde38b1461031f575f80fd5b80636838a4f61461024357806370a0823114610264578063715018a6146102775780638672c2121461027f5780638da5cb5b14610292575f80fd5b806323b872dd1161010557806323b872dd146101d1578063330c093c146101e45780633bb3a24d146101f757806342842e0e1461020a57806355f804b31461021d5780636352211e14610230575f80fd5b806301ffc9a71461014157806306fdde0314610169578063081812fc1461017e578063095ea7b3146101a9578063126d5b6a146101be575b5f80fd5b61015461014f36600461136b565b610332565b60405190151581526020015b60405180910390f35b61017161035c565b60405161016091906113d3565b61019161018c3660046113e5565b6103eb565b6040516001600160a01b039091168152602001610160565b6101bc6101b7366004611417565b610410565b005b6101bc6101cc36600461143f565b610529565b6101bc6101df366004611458565b610553565b600a54610191906001600160a01b031681565b6101716102053660046113e5565b610584565b6101bc610218366004611458565b61058f565b6101bc61022b366004611536565b6105a9565b61019161023e3660046113e5565b6105c1565b610256610251366004611568565b610620565b604051908152602001610160565b61025661027236600461143f565b6106d7565b6101bc61075b565b600b54610191906001600160a01b031681565b6007546001600160a01b0316610191565b61017161076e565b6101bc6102b93660046115b3565b61077d565b6101bc6102cc3660046115ec565b610788565b6101716102df3660046113e5565b6107c0565b6101546102f2366004611663565b6001600160a01b039182165f90815260056020908152604080832093909416825291909152205460ff1690565b6101bc61032d36600461143f565b6108be565b5f6001600160e01b03198216632483248360e11b1480610356575061035682610937565b92915050565b60605f805461036a9061168b565b80601f01602080910402602001604051908101604052809291908181526020018280546103969061168b565b80156103e15780601f106103b8576101008083540402835291602001916103e1565b820191905f5260205f20905b8154815290600101906020018083116103c457829003601f168201915b5050505050905090565b5f6103f582610986565b505f908152600460205260409020546001600160a01b031690565b5f61041a826105c1565b9050806001600160a01b0316836001600160a01b03160361048c5760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b60648201526084015b60405180910390fd5b336001600160a01b03821614806104a857506104a881336102f2565b61051a5760405162461bcd60e51b815260206004820152603d60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206f7220617070726f76656420666f7220616c6c0000006064820152608401610483565b61052483836109e4565b505050565b610531610a51565b600b80546001600160a01b0319166001600160a01b0392909216919091179055565b61055d3382610aab565b6105795760405162461bcd60e51b8152600401610483906116c3565b610524838383610b27565b6060610356826107c0565b61052483838360405180602001604052805f815250610788565b6105b1610a51565b60096105bd828261175b565b5050565b5f818152600260205260408120546001600160a01b0316806103565760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b6044820152606401610483565b600b545f906001600160a01b031633146106a25760405162461bcd60e51b815260206004820152603f60248201527f54686973206d6574686f642063616e206f6e6c792062652063616c6c6564206260448201527f792074686520736361747465722070726f746f636f6c20636f6e7472616374006064820152608401610483565b6106b0600c80546001019055565b5f6106ba600c5490565b90506106c68382610c89565b6106d08185610ca2565b9392505050565b5f6001600160a01b0382166107405760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b6064820152608401610483565b506001600160a01b03165f9081526003602052604090205490565b610763610a51565b61076c5f610cac565b565b60606001805461036a9061168b565b6105bd338383610cfd565b6107923383610aab565b6107ae5760405162461bcd60e51b8152600401610483906116c3565b6107ba84848484610dca565b50505050565b60606107cb82610986565b5f82815260066020526040812080546107e39061168b565b80601f016020809104026020016040519081016040528092919081815260200182805461080f9061168b565b801561085a5780601f106108315761010080835404028352916020019161085a565b820191905f5260205f20905b81548152906001019060200180831161083d57829003601f168201915b505050505090505f61086a610dfd565b905080515f0361087b575092915050565b8151156108ad57808260405160200161089592919061181b565b60405160208183030381529060405292505050919050565b6108b684610e0c565b949350505050565b6108c6610a51565b6001600160a01b03811661092b5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610483565b61093481610cac565b50565b5f6001600160e01b031982166380ac58cd60e01b148061096757506001600160e01b03198216635b5e139f60e01b145b8061035657506301ffc9a760e01b6001600160e01b0319831614610356565b5f818152600260205260409020546001600160a01b03166109345760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b6044820152606401610483565b5f81815260046020526040902080546001600160a01b0319166001600160a01b0384169081179091558190610a18826105c1565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6007546001600160a01b0316331461076c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610483565b5f80610ab6836105c1565b9050806001600160a01b0316846001600160a01b03161480610afc57506001600160a01b038082165f9081526005602090815260408083209388168352929052205460ff165b806108b65750836001600160a01b0316610b15846103eb565b6001600160a01b031614949350505050565b826001600160a01b0316610b3a826105c1565b6001600160a01b031614610b605760405162461bcd60e51b815260040161048390611849565b6001600160a01b038216610bc25760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b6064820152608401610483565b826001600160a01b0316610bd5826105c1565b6001600160a01b031614610bfb5760405162461bcd60e51b815260040161048390611849565b5f81815260046020908152604080832080546001600160a01b03199081169091556001600160a01b038781168086526003855283862080545f1901905590871680865283862080546001019055868652600290945282852080549092168417909155905184937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6105bd828260405180602001604052805f815250610e6f565b6105bd8282610ea1565b600780546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b816001600160a01b0316836001600160a01b031603610d5e5760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c6572000000000000006044820152606401610483565b6001600160a01b038381165f81815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b610dd5848484610b27565b610de184848484610f6a565b6107ba5760405162461bcd60e51b81526004016104839061188e565b60606009805461036a9061168b565b6060610e1782610986565b5f610e20610dfd565b90505f815111610e3e5760405180602001604052805f8152506106d0565b80610e4884611067565b604051602001610e5992919061181b565b6040516020818303038152906040529392505050565b610e7983836110f7565b610e855f848484610f6a565b6105245760405162461bcd60e51b81526004016104839061188e565b5f828152600260205260409020546001600160a01b0316610f1b5760405162461bcd60e51b815260206004820152602e60248201527f45524337323155524953746f726167653a2055524920736574206f66206e6f6e60448201526d32bc34b9ba32b73a103a37b5b2b760911b6064820152608401610483565b5f828152600660205260409020610f32828261175b565b506040518281527ff8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce79060200160405180910390a15050565b5f6001600160a01b0384163b1561105c57604051630a85bd0160e11b81526001600160a01b0385169063150b7a0290610fad9033908990889088906004016118e0565b6020604051808303815f875af1925050508015610fe7575060408051601f3d908101601f19168201909252610fe49181019061191c565b60015b611042573d808015611014576040519150601f19603f3d011682016040523d82523d5f602084013e611019565b606091505b5080515f0361103a5760405162461bcd60e51b81526004016104839061188e565b805181602001fd5b6001600160e01b031916630a85bd0160e11b1490506108b6565b506001949350505050565b60605f6110738361127f565b60010190505f8167ffffffffffffffff81111561109257611092611491565b6040519080825280601f01601f1916602001820160405280156110bc576020820181803683370190505b5090508181016020015b5f19016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a85049450846110c657509392505050565b6001600160a01b03821661114d5760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f20616464726573736044820152606401610483565b5f818152600260205260409020546001600160a01b0316156111b15760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610483565b5f818152600260205260409020546001600160a01b0316156112155760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610483565b6001600160a01b0382165f81815260036020908152604080832080546001019055848352600290915280822080546001600160a01b0319168417905551839291907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b5f8072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b83106112bd5772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef810000000083106112e9576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061130757662386f26fc10000830492506010015b6305f5e100831061131f576305f5e100830492506008015b612710831061133357612710830492506004015b60648310611345576064830492506002015b600a83106103565760010192915050565b6001600160e01b031981168114610934575f80fd5b5f6020828403121561137b575f80fd5b81356106d081611356565b5f5b838110156113a0578181015183820152602001611388565b50505f910152565b5f81518084526113bf816020860160208601611386565b601f01601f19169290920160200192915050565b602081525f6106d060208301846113a8565b5f602082840312156113f5575f80fd5b5035919050565b80356001600160a01b0381168114611412575f80fd5b919050565b5f8060408385031215611428575f80fd5b611431836113fc565b946020939093013593505050565b5f6020828403121561144f575f80fd5b6106d0826113fc565b5f805f6060848603121561146a575f80fd5b611473846113fc565b9250611481602085016113fc565b9150604084013590509250925092565b634e487b7160e01b5f52604160045260245ffd5b5f67ffffffffffffffff808411156114bf576114bf611491565b604051601f8501601f19908116603f011681019082821181831017156114e7576114e7611491565b816040528093508581528686860111156114ff575f80fd5b858560208301375f602087830101525050509392505050565b5f82601f830112611527575f80fd5b6106d0838335602085016114a5565b5f60208284031215611546575f80fd5b813567ffffffffffffffff81111561155c575f80fd5b6108b684828501611518565b5f8060408385031215611579575f80fd5b823567ffffffffffffffff81111561158f575f80fd5b61159b85828601611518565b9250506115aa602084016113fc565b90509250929050565b5f80604083850312156115c4575f80fd5b6115cd836113fc565b9150602083013580151581146115e1575f80fd5b809150509250929050565b5f805f80608085870312156115ff575f80fd5b611608856113fc565b9350611616602086016113fc565b925060408501359150606085013567ffffffffffffffff811115611638575f80fd5b8501601f81018713611648575f80fd5b611657878235602084016114a5565b91505092959194509250565b5f8060408385031215611674575f80fd5b61167d836113fc565b91506115aa602084016113fc565b600181811c9082168061169f57607f821691505b6020821081036116bd57634e487b7160e01b5f52602260045260245ffd5b50919050565b6020808252602d908201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560408201526c1c881bdc88185c1c1c9bdd9959609a1b606082015260800190565b601f82111561052457805f5260205f20601f840160051c810160208510156117355750805b601f840160051c820191505b81811015611754575f8155600101611741565b5050505050565b815167ffffffffffffffff81111561177557611775611491565b61178981611783845461168b565b84611710565b602080601f8311600181146117bc575f84156117a55750858301515b5f19600386901b1c1916600185901b178555611813565b5f85815260208120601f198616915b828110156117ea578886015182559484019460019091019084016117cb565b508582101561180757878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b5f835161182c818460208801611386565b835190830190611840818360208801611386565b01949350505050565b60208082526025908201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060408201526437bbb732b960d91b606082015260800190565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b6001600160a01b03858116825284166020820152604081018390526080606082018190525f90611912908301846113a8565b9695505050505050565b5f6020828403121561192c575f80fd5b81516106d08161135656fea2646970667358221220c10ba8c6a3295acfc812354f2a95c269d70819b2dea37a2fb3f56d2d00eb5bb964736f6c63430008170033",
}

// TrainingtokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TrainingtokenMetaData.ABI instead.
var TrainingtokenABI = TrainingtokenMetaData.ABI

// TrainingtokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TrainingtokenMetaData.Bin instead.
var TrainingtokenBin = TrainingtokenMetaData.Bin

// DeployTrainingtoken deploys a new Ethereum contract, binding an instance of Trainingtoken to it.
func DeployTrainingtoken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Trainingtoken, error) {
	parsed, err := TrainingtokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TrainingtokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Trainingtoken{TrainingtokenCaller: TrainingtokenCaller{contract: contract}, TrainingtokenTransactor: TrainingtokenTransactor{contract: contract}, TrainingtokenFilterer: TrainingtokenFilterer{contract: contract}}, nil
}

// Trainingtoken is an auto generated Go binding around an Ethereum contract.
type Trainingtoken struct {
	TrainingtokenCaller     // Read-only binding to the contract
	TrainingtokenTransactor // Write-only binding to the contract
	TrainingtokenFilterer   // Log filterer for contract events
}

// TrainingtokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TrainingtokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrainingtokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TrainingtokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrainingtokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TrainingtokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrainingtokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TrainingtokenSession struct {
	Contract     *Trainingtoken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TrainingtokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TrainingtokenCallerSession struct {
	Contract *TrainingtokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TrainingtokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TrainingtokenTransactorSession struct {
	Contract     *TrainingtokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TrainingtokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TrainingtokenRaw struct {
	Contract *Trainingtoken // Generic contract binding to access the raw methods on
}

// TrainingtokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TrainingtokenCallerRaw struct {
	Contract *TrainingtokenCaller // Generic read-only contract binding to access the raw methods on
}

// TrainingtokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TrainingtokenTransactorRaw struct {
	Contract *TrainingtokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrainingtoken creates a new instance of Trainingtoken, bound to a specific deployed contract.
func NewTrainingtoken(address common.Address, backend bind.ContractBackend) (*Trainingtoken, error) {
	contract, err := bindTrainingtoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Trainingtoken{TrainingtokenCaller: TrainingtokenCaller{contract: contract}, TrainingtokenTransactor: TrainingtokenTransactor{contract: contract}, TrainingtokenFilterer: TrainingtokenFilterer{contract: contract}}, nil
}

// NewTrainingtokenCaller creates a new read-only instance of Trainingtoken, bound to a specific deployed contract.
func NewTrainingtokenCaller(address common.Address, caller bind.ContractCaller) (*TrainingtokenCaller, error) {
	contract, err := bindTrainingtoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TrainingtokenCaller{contract: contract}, nil
}

// NewTrainingtokenTransactor creates a new write-only instance of Trainingtoken, bound to a specific deployed contract.
func NewTrainingtokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TrainingtokenTransactor, error) {
	contract, err := bindTrainingtoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TrainingtokenTransactor{contract: contract}, nil
}

// NewTrainingtokenFilterer creates a new log filterer instance of Trainingtoken, bound to a specific deployed contract.
func NewTrainingtokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TrainingtokenFilterer, error) {
	contract, err := bindTrainingtoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TrainingtokenFilterer{contract: contract}, nil
}

// bindTrainingtoken binds a generic wrapper to an already deployed contract.
func bindTrainingtoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TrainingtokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trainingtoken *TrainingtokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trainingtoken.Contract.TrainingtokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trainingtoken *TrainingtokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trainingtoken.Contract.TrainingtokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trainingtoken *TrainingtokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trainingtoken.Contract.TrainingtokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trainingtoken *TrainingtokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trainingtoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trainingtoken *TrainingtokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trainingtoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trainingtoken *TrainingtokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trainingtoken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Trainingtoken *TrainingtokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Trainingtoken.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Trainingtoken *TrainingtokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Trainingtoken.Contract.BalanceOf(&_Trainingtoken.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Trainingtoken *TrainingtokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Trainingtoken.Contract.BalanceOf(&_Trainingtoken.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Trainingtoken *TrainingtokenCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Trainingtoken.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Trainingtoken *TrainingtokenSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Trainingtoken.Contract.GetApproved(&_Trainingtoken.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Trainingtoken *TrainingtokenCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Trainingtoken.Contract.GetApproved(&_Trainingtoken.CallOpts, tokenId)
}

// GetTokenURI is a free data retrieval call binding the contract method 0x3bb3a24d.
//
// Solidity: function getTokenURI(uint256 tokenId) view returns(string)
func (_Trainingtoken *TrainingtokenCaller) GetTokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Trainingtoken.contract.Call(opts, &out, "getTokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetTokenURI is a free data retrieval call binding the contract method 0x3bb3a24d.
//
// Solidity: function getTokenURI(uint256 tokenId) view returns(string)
func (_Trainingtoken *TrainingtokenSession) GetTokenURI(tokenId *big.Int) (string, error) {
	return _Trainingtoken.Contract.GetTokenURI(&_Trainingtoken.CallOpts, tokenId)
}

// GetTokenURI is a free data retrieval call binding the contract method 0x3bb3a24d.
//
// Solidity: function getTokenURI(uint256 tokenId) view returns(string)
func (_Trainingtoken *TrainingtokenCallerSession) GetTokenURI(tokenId *big.Int) (string, error) {
	return _Trainingtoken.Contract.GetTokenURI(&_Trainingtoken.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Trainingtoken *TrainingtokenCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Trainingtoken.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Trainingtoken *TrainingtokenSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Trainingtoken.Contract.IsApprovedForAll(&_Trainingtoken.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Trainingtoken *TrainingtokenCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Trainingtoken.Contract.IsApprovedForAll(&_Trainingtoken.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Trainingtoken *TrainingtokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Trainingtoken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Trainingtoken *TrainingtokenSession) Name() (string, error) {
	return _Trainingtoken.Contract.Name(&_Trainingtoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Trainingtoken *TrainingtokenCallerSession) Name() (string, error) {
	return _Trainingtoken.Contract.Name(&_Trainingtoken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trainingtoken *TrainingtokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Trainingtoken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trainingtoken *TrainingtokenSession) Owner() (common.Address, error) {
	return _Trainingtoken.Contract.Owner(&_Trainingtoken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trainingtoken *TrainingtokenCallerSession) Owner() (common.Address, error) {
	return _Trainingtoken.Contract.Owner(&_Trainingtoken.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Trainingtoken *TrainingtokenCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Trainingtoken.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Trainingtoken *TrainingtokenSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Trainingtoken.Contract.OwnerOf(&_Trainingtoken.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Trainingtoken *TrainingtokenCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Trainingtoken.Contract.OwnerOf(&_Trainingtoken.CallOpts, tokenId)
}

// ProtocolDeployer is a free data retrieval call binding the contract method 0x330c093c.
//
// Solidity: function protocolDeployer() view returns(address)
func (_Trainingtoken *TrainingtokenCaller) ProtocolDeployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Trainingtoken.contract.Call(opts, &out, "protocolDeployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolDeployer is a free data retrieval call binding the contract method 0x330c093c.
//
// Solidity: function protocolDeployer() view returns(address)
func (_Trainingtoken *TrainingtokenSession) ProtocolDeployer() (common.Address, error) {
	return _Trainingtoken.Contract.ProtocolDeployer(&_Trainingtoken.CallOpts)
}

// ProtocolDeployer is a free data retrieval call binding the contract method 0x330c093c.
//
// Solidity: function protocolDeployer() view returns(address)
func (_Trainingtoken *TrainingtokenCallerSession) ProtocolDeployer() (common.Address, error) {
	return _Trainingtoken.Contract.ProtocolDeployer(&_Trainingtoken.CallOpts)
}

// ScatterContractAddress is a free data retrieval call binding the contract method 0x8672c212.
//
// Solidity: function scatterContractAddress() view returns(address)
func (_Trainingtoken *TrainingtokenCaller) ScatterContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Trainingtoken.contract.Call(opts, &out, "scatterContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ScatterContractAddress is a free data retrieval call binding the contract method 0x8672c212.
//
// Solidity: function scatterContractAddress() view returns(address)
func (_Trainingtoken *TrainingtokenSession) ScatterContractAddress() (common.Address, error) {
	return _Trainingtoken.Contract.ScatterContractAddress(&_Trainingtoken.CallOpts)
}

// ScatterContractAddress is a free data retrieval call binding the contract method 0x8672c212.
//
// Solidity: function scatterContractAddress() view returns(address)
func (_Trainingtoken *TrainingtokenCallerSession) ScatterContractAddress() (common.Address, error) {
	return _Trainingtoken.Contract.ScatterContractAddress(&_Trainingtoken.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Trainingtoken *TrainingtokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Trainingtoken.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Trainingtoken *TrainingtokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Trainingtoken.Contract.SupportsInterface(&_Trainingtoken.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Trainingtoken *TrainingtokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Trainingtoken.Contract.SupportsInterface(&_Trainingtoken.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Trainingtoken *TrainingtokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Trainingtoken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Trainingtoken *TrainingtokenSession) Symbol() (string, error) {
	return _Trainingtoken.Contract.Symbol(&_Trainingtoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Trainingtoken *TrainingtokenCallerSession) Symbol() (string, error) {
	return _Trainingtoken.Contract.Symbol(&_Trainingtoken.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Trainingtoken *TrainingtokenCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Trainingtoken.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Trainingtoken *TrainingtokenSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Trainingtoken.Contract.TokenURI(&_Trainingtoken.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Trainingtoken *TrainingtokenCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Trainingtoken.Contract.TokenURI(&_Trainingtoken.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Trainingtoken *TrainingtokenTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Trainingtoken.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Trainingtoken *TrainingtokenSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Trainingtoken.Contract.Approve(&_Trainingtoken.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Trainingtoken *TrainingtokenTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Trainingtoken.Contract.Approve(&_Trainingtoken.TransactOpts, to, tokenId)
}

// PublishTrainingJob is a paid mutator transaction binding the contract method 0x6838a4f6.
//
// Solidity: function publishTrainingJob(string tokenURI, address recipient) returns(uint256)
func (_Trainingtoken *TrainingtokenTransactor) PublishTrainingJob(opts *bind.TransactOpts, tokenURI string, recipient common.Address) (*types.Transaction, error) {
	return _Trainingtoken.contract.Transact(opts, "publishTrainingJob", tokenURI, recipient)
}

// PublishTrainingJob is a paid mutator transaction binding the contract method 0x6838a4f6.
//
// Solidity: function publishTrainingJob(string tokenURI, address recipient) returns(uint256)
func (_Trainingtoken *TrainingtokenSession) PublishTrainingJob(tokenURI string, recipient common.Address) (*types.Transaction, error) {
	return _Trainingtoken.Contract.PublishTrainingJob(&_Trainingtoken.TransactOpts, tokenURI, recipient)
}

// PublishTrainingJob is a paid mutator transaction binding the contract method 0x6838a4f6.
//
// Solidity: function publishTrainingJob(string tokenURI, address recipient) returns(uint256)
func (_Trainingtoken *TrainingtokenTransactorSession) PublishTrainingJob(tokenURI string, recipient common.Address) (*types.Transaction, error) {
	return _Trainingtoken.Contract.PublishTrainingJob(&_Trainingtoken.TransactOpts, tokenURI, recipient)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Trainingtoken *TrainingtokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trainingtoken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Trainingtoken *TrainingtokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _Trainingtoken.Contract.RenounceOwnership(&_Trainingtoken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Trainingtoken *TrainingtokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Trainingtoken.Contract.RenounceOwnership(&_Trainingtoken.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Trainingtoken *TrainingtokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Trainingtoken.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Trainingtoken *TrainingtokenSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Trainingtoken.Contract.SafeTransferFrom(&_Trainingtoken.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Trainingtoken *TrainingtokenTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Trainingtoken.Contract.SafeTransferFrom(&_Trainingtoken.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Trainingtoken *TrainingtokenTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Trainingtoken.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Trainingtoken *TrainingtokenSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Trainingtoken.Contract.SafeTransferFrom0(&_Trainingtoken.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Trainingtoken *TrainingtokenTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Trainingtoken.Contract.SafeTransferFrom0(&_Trainingtoken.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Trainingtoken *TrainingtokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Trainingtoken.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Trainingtoken *TrainingtokenSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Trainingtoken.Contract.SetApprovalForAll(&_Trainingtoken.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Trainingtoken *TrainingtokenTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Trainingtoken.Contract.SetApprovalForAll(&_Trainingtoken.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI_) returns()
func (_Trainingtoken *TrainingtokenTransactor) SetBaseURI(opts *bind.TransactOpts, baseURI_ string) (*types.Transaction, error) {
	return _Trainingtoken.contract.Transact(opts, "setBaseURI", baseURI_)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI_) returns()
func (_Trainingtoken *TrainingtokenSession) SetBaseURI(baseURI_ string) (*types.Transaction, error) {
	return _Trainingtoken.Contract.SetBaseURI(&_Trainingtoken.TransactOpts, baseURI_)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI_) returns()
func (_Trainingtoken *TrainingtokenTransactorSession) SetBaseURI(baseURI_ string) (*types.Transaction, error) {
	return _Trainingtoken.Contract.SetBaseURI(&_Trainingtoken.TransactOpts, baseURI_)
}

// SetScatterContractAddress is a paid mutator transaction binding the contract method 0x126d5b6a.
//
// Solidity: function setScatterContractAddress(address contractAddress) returns()
func (_Trainingtoken *TrainingtokenTransactor) SetScatterContractAddress(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _Trainingtoken.contract.Transact(opts, "setScatterContractAddress", contractAddress)
}

// SetScatterContractAddress is a paid mutator transaction binding the contract method 0x126d5b6a.
//
// Solidity: function setScatterContractAddress(address contractAddress) returns()
func (_Trainingtoken *TrainingtokenSession) SetScatterContractAddress(contractAddress common.Address) (*types.Transaction, error) {
	return _Trainingtoken.Contract.SetScatterContractAddress(&_Trainingtoken.TransactOpts, contractAddress)
}

// SetScatterContractAddress is a paid mutator transaction binding the contract method 0x126d5b6a.
//
// Solidity: function setScatterContractAddress(address contractAddress) returns()
func (_Trainingtoken *TrainingtokenTransactorSession) SetScatterContractAddress(contractAddress common.Address) (*types.Transaction, error) {
	return _Trainingtoken.Contract.SetScatterContractAddress(&_Trainingtoken.TransactOpts, contractAddress)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Trainingtoken *TrainingtokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Trainingtoken.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Trainingtoken *TrainingtokenSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Trainingtoken.Contract.TransferFrom(&_Trainingtoken.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Trainingtoken *TrainingtokenTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Trainingtoken.Contract.TransferFrom(&_Trainingtoken.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Trainingtoken *TrainingtokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Trainingtoken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Trainingtoken *TrainingtokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Trainingtoken.Contract.TransferOwnership(&_Trainingtoken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Trainingtoken *TrainingtokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Trainingtoken.Contract.TransferOwnership(&_Trainingtoken.TransactOpts, newOwner)
}

// TrainingtokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Trainingtoken contract.
type TrainingtokenApprovalIterator struct {
	Event *TrainingtokenApproval // Event containing the contract specifics and raw log

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
func (it *TrainingtokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrainingtokenApproval)
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
		it.Event = new(TrainingtokenApproval)
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
func (it *TrainingtokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrainingtokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrainingtokenApproval represents a Approval event raised by the Trainingtoken contract.
type TrainingtokenApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Trainingtoken *TrainingtokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*TrainingtokenApprovalIterator, error) {

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

	logs, sub, err := _Trainingtoken.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TrainingtokenApprovalIterator{contract: _Trainingtoken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Trainingtoken *TrainingtokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TrainingtokenApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Trainingtoken.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrainingtokenApproval)
				if err := _Trainingtoken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Trainingtoken *TrainingtokenFilterer) ParseApproval(log types.Log) (*TrainingtokenApproval, error) {
	event := new(TrainingtokenApproval)
	if err := _Trainingtoken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrainingtokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Trainingtoken contract.
type TrainingtokenApprovalForAllIterator struct {
	Event *TrainingtokenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *TrainingtokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrainingtokenApprovalForAll)
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
		it.Event = new(TrainingtokenApprovalForAll)
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
func (it *TrainingtokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrainingtokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrainingtokenApprovalForAll represents a ApprovalForAll event raised by the Trainingtoken contract.
type TrainingtokenApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Trainingtoken *TrainingtokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*TrainingtokenApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Trainingtoken.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &TrainingtokenApprovalForAllIterator{contract: _Trainingtoken.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Trainingtoken *TrainingtokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *TrainingtokenApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Trainingtoken.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrainingtokenApprovalForAll)
				if err := _Trainingtoken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Trainingtoken *TrainingtokenFilterer) ParseApprovalForAll(log types.Log) (*TrainingtokenApprovalForAll, error) {
	event := new(TrainingtokenApprovalForAll)
	if err := _Trainingtoken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrainingtokenBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the Trainingtoken contract.
type TrainingtokenBatchMetadataUpdateIterator struct {
	Event *TrainingtokenBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *TrainingtokenBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrainingtokenBatchMetadataUpdate)
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
		it.Event = new(TrainingtokenBatchMetadataUpdate)
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
func (it *TrainingtokenBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrainingtokenBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrainingtokenBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the Trainingtoken contract.
type TrainingtokenBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Trainingtoken *TrainingtokenFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*TrainingtokenBatchMetadataUpdateIterator, error) {

	logs, sub, err := _Trainingtoken.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &TrainingtokenBatchMetadataUpdateIterator{contract: _Trainingtoken.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Trainingtoken *TrainingtokenFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *TrainingtokenBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Trainingtoken.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrainingtokenBatchMetadataUpdate)
				if err := _Trainingtoken.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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
func (_Trainingtoken *TrainingtokenFilterer) ParseBatchMetadataUpdate(log types.Log) (*TrainingtokenBatchMetadataUpdate, error) {
	event := new(TrainingtokenBatchMetadataUpdate)
	if err := _Trainingtoken.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrainingtokenMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the Trainingtoken contract.
type TrainingtokenMetadataUpdateIterator struct {
	Event *TrainingtokenMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *TrainingtokenMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrainingtokenMetadataUpdate)
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
		it.Event = new(TrainingtokenMetadataUpdate)
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
func (it *TrainingtokenMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrainingtokenMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrainingtokenMetadataUpdate represents a MetadataUpdate event raised by the Trainingtoken contract.
type TrainingtokenMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Trainingtoken *TrainingtokenFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*TrainingtokenMetadataUpdateIterator, error) {

	logs, sub, err := _Trainingtoken.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &TrainingtokenMetadataUpdateIterator{contract: _Trainingtoken.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Trainingtoken *TrainingtokenFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *TrainingtokenMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Trainingtoken.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrainingtokenMetadataUpdate)
				if err := _Trainingtoken.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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
func (_Trainingtoken *TrainingtokenFilterer) ParseMetadataUpdate(log types.Log) (*TrainingtokenMetadataUpdate, error) {
	event := new(TrainingtokenMetadataUpdate)
	if err := _Trainingtoken.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrainingtokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Trainingtoken contract.
type TrainingtokenOwnershipTransferredIterator struct {
	Event *TrainingtokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TrainingtokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrainingtokenOwnershipTransferred)
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
		it.Event = new(TrainingtokenOwnershipTransferred)
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
func (it *TrainingtokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrainingtokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrainingtokenOwnershipTransferred represents a OwnershipTransferred event raised by the Trainingtoken contract.
type TrainingtokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Trainingtoken *TrainingtokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TrainingtokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Trainingtoken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TrainingtokenOwnershipTransferredIterator{contract: _Trainingtoken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Trainingtoken *TrainingtokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TrainingtokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Trainingtoken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrainingtokenOwnershipTransferred)
				if err := _Trainingtoken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Trainingtoken *TrainingtokenFilterer) ParseOwnershipTransferred(log types.Log) (*TrainingtokenOwnershipTransferred, error) {
	event := new(TrainingtokenOwnershipTransferred)
	if err := _Trainingtoken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrainingtokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Trainingtoken contract.
type TrainingtokenTransferIterator struct {
	Event *TrainingtokenTransfer // Event containing the contract specifics and raw log

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
func (it *TrainingtokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrainingtokenTransfer)
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
		it.Event = new(TrainingtokenTransfer)
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
func (it *TrainingtokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrainingtokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrainingtokenTransfer represents a Transfer event raised by the Trainingtoken contract.
type TrainingtokenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Trainingtoken *TrainingtokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*TrainingtokenTransferIterator, error) {

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

	logs, sub, err := _Trainingtoken.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TrainingtokenTransferIterator{contract: _Trainingtoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Trainingtoken *TrainingtokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TrainingtokenTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Trainingtoken.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrainingtokenTransfer)
				if err := _Trainingtoken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Trainingtoken *TrainingtokenFilterer) ParseTransfer(log types.Log) (*TrainingtokenTransfer, error) {
	event := new(TrainingtokenTransfer)
	if err := _Trainingtoken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
