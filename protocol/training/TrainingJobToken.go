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
)

// TrainingtokenMetaData contains all meta data concerning the Trainingtoken contract.
var TrainingtokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getTokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolDeployer\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"publishTrainingJob\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scatterContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setScatterContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060400160405280601e81526020017f536361747465722050726f746f636f6c20547261696e696e67204a6f627300008152506040518060400160405280600481526020016329a82a2560e11b8152508160009081620000769190620001b5565b506001620000858282620001b5565b505050620000a26200009c620000ba60201b60201c565b620000be565b600a80546001600160a01b0319163317905562000281565b3390565b600780546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806200013b57607f821691505b6020821081036200015c57634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620001b057600081815260208120601f850160051c810160208610156200018b5750805b601f850160051c820191505b81811015620001ac5782815560010162000197565b5050505b505050565b81516001600160401b03811115620001d157620001d162000110565b620001e981620001e2845462000126565b8462000162565b602080601f831160018114620002215760008415620002085750858301515b600019600386901b1c1916600185901b178555620001ac565b600085815260208120601f198616915b82811015620002525788860151825594840194600190910190840162000231565b5085821015620002715787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6119d380620002916000396000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c80636838a4f6116100b857806395d89b411161007c57806395d89b41146102a9578063a22cb465146102b1578063b88d4fde146102c4578063c87b56dd146102d7578063e985e9c5146102ea578063f2fde38b1461032657600080fd5b80636838a4f61461024957806370a082311461026a578063715018a61461027d5780638672c212146102855780638da5cb5b1461029857600080fd5b806323b872dd1161010a57806323b872dd146101d7578063330c093c146101ea5780633bb3a24d146101fd57806342842e0e1461021057806355f804b3146102235780636352211e1461023657600080fd5b806301ffc9a71461014757806306fdde031461016f578063081812fc14610184578063095ea7b3146101af578063126d5b6a146101c4575b600080fd5b61015a6101553660046113a1565b610339565b60405190151581526020015b60405180910390f35b610177610364565b604051610166919061140e565b610197610192366004611421565b6103f6565b6040516001600160a01b039091168152602001610166565b6101c26101bd366004611456565b61041d565b005b6101c26101d2366004611480565b610537565b6101c26101e536600461149b565b610561565b600a54610197906001600160a01b031681565b61017761020b366004611421565b610592565b6101c261021e36600461149b565b61059d565b6101c2610231366004611583565b6105b8565b610197610244366004611421565b6105d0565b61025c6102573660046115b8565b610630565b604051908152602001610166565b61025c610278366004611480565b6106e9565b6101c261076f565b600b54610197906001600160a01b031681565b6007546001600160a01b0316610197565b610177610783565b6101c26102bf366004611606565b610792565b6101c26102d2366004611642565b61079d565b6101776102e5366004611421565b6107d5565b61015a6102f83660046116be565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b6101c2610334366004611480565b6108d8565b60006001600160e01b03198216632483248360e11b148061035e575061035e82610951565b92915050565b606060008054610373906116e8565b80601f016020809104026020016040519081016040528092919081815260200182805461039f906116e8565b80156103ec5780601f106103c1576101008083540402835291602001916103ec565b820191906000526020600020905b8154815290600101906020018083116103cf57829003601f168201915b5050505050905090565b6000610401826109a1565b506000908152600460205260409020546001600160a01b031690565b6000610428826105d0565b9050806001600160a01b0316836001600160a01b03160361049a5760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b60648201526084015b60405180910390fd5b336001600160a01b03821614806104b657506104b681336102f8565b6105285760405162461bcd60e51b815260206004820152603d60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206f7220617070726f76656420666f7220616c6c0000006064820152608401610491565b6105328383610a00565b505050565b61053f610a6e565b600b80546001600160a01b0319166001600160a01b0392909216919091179055565b61056b3382610ac8565b6105875760405162461bcd60e51b815260040161049190611722565b610532838383610b46565b606061035e826107d5565b6105328383836040518060200160405280600081525061079d565b6105c0610a6e565b60096105cc82826117bd565b5050565b6000818152600260205260408120546001600160a01b03168061035e5760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b6044820152606401610491565b600b546000906001600160a01b031633146106b35760405162461bcd60e51b815260206004820152603f60248201527f54686973206d6574686f642063616e206f6e6c792062652063616c6c6564206260448201527f792074686520736361747465722070726f746f636f6c20636f6e7472616374006064820152608401610491565b6106c1600c80546001019055565b60006106cc600c5490565b90506106d88382610caa565b6106e28185610cc4565b9392505050565b60006001600160a01b0382166107535760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b6064820152608401610491565b506001600160a01b031660009081526003602052604090205490565b610777610a6e565b6107816000610cce565b565b606060018054610373906116e8565b6105cc338383610d20565b6107a73383610ac8565b6107c35760405162461bcd60e51b815260040161049190611722565b6107cf84848484610dee565b50505050565b60606107e0826109a1565b600082815260066020526040812080546107f9906116e8565b80601f0160208091040260200160405190810160405280929190818152602001828054610825906116e8565b80156108725780601f1061084757610100808354040283529160200191610872565b820191906000526020600020905b81548152906001019060200180831161085557829003601f168201915b505050505090506000610883610e21565b90508051600003610895575092915050565b8151156108c75780826040516020016108af92919061187d565b60405160208183030381529060405292505050919050565b6108d084610e30565b949350505050565b6108e0610a6e565b6001600160a01b0381166109455760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610491565b61094e81610cce565b50565b60006001600160e01b031982166380ac58cd60e01b148061098257506001600160e01b03198216635b5e139f60e01b145b8061035e57506301ffc9a760e01b6001600160e01b031983161461035e565b6000818152600260205260409020546001600160a01b031661094e5760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b6044820152606401610491565b600081815260046020526040902080546001600160a01b0319166001600160a01b0384169081179091558190610a35826105d0565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6007546001600160a01b031633146107815760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610491565b600080610ad4836105d0565b9050806001600160a01b0316846001600160a01b03161480610b1b57506001600160a01b0380821660009081526005602090815260408083209388168352929052205460ff165b806108d05750836001600160a01b0316610b34846103f6565b6001600160a01b031614949350505050565b826001600160a01b0316610b59826105d0565b6001600160a01b031614610b7f5760405162461bcd60e51b8152600401610491906118ac565b6001600160a01b038216610be15760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b6064820152608401610491565b826001600160a01b0316610bf4826105d0565b6001600160a01b031614610c1a5760405162461bcd60e51b8152600401610491906118ac565b600081815260046020908152604080832080546001600160a01b03199081169091556001600160a01b0387811680865260038552838620805460001901905590871680865283862080546001019055868652600290945282852080549092168417909155905184937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6105cc828260405180602001604052806000815250610e96565b6105cc8282610ec9565b600780546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b816001600160a01b0316836001600160a01b031603610d815760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c6572000000000000006044820152606401610491565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b610df9848484610b46565b610e0584848484610f94565b6107cf5760405162461bcd60e51b8152600401610491906118f1565b606060098054610373906116e8565b6060610e3b826109a1565b6000610e45610e21565b90506000815111610e6557604051806020016040528060008152506106e2565b80610e6f84611095565b604051602001610e8092919061187d565b6040516020818303038152906040529392505050565b610ea08383611128565b610ead6000848484610f94565b6105325760405162461bcd60e51b8152600401610491906118f1565b6000828152600260205260409020546001600160a01b0316610f445760405162461bcd60e51b815260206004820152602e60248201527f45524337323155524953746f726167653a2055524920736574206f66206e6f6e60448201526d32bc34b9ba32b73a103a37b5b2b760911b6064820152608401610491565b6000828152600660205260409020610f5c82826117bd565b506040518281527ff8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce79060200160405180910390a15050565b60006001600160a01b0384163b1561108a57604051630a85bd0160e11b81526001600160a01b0385169063150b7a0290610fd8903390899088908890600401611943565b6020604051808303816000875af1925050508015611013575060408051601f3d908101601f1916820190925261101091810190611980565b60015b611070573d808015611041576040519150601f19603f3d011682016040523d82523d6000602084013e611046565b606091505b5080516000036110685760405162461bcd60e51b8152600401610491906118f1565b805181602001fd5b6001600160e01b031916630a85bd0160e11b1490506108d0565b506001949350505050565b606060006110a2836112b3565b600101905060008167ffffffffffffffff8111156110c2576110c26114d7565b6040519080825280601f01601f1916602001820160405280156110ec576020820181803683370190505b5090508181016020015b600019016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a85049450846110f657509392505050565b6001600160a01b03821661117e5760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f20616464726573736044820152606401610491565b6000818152600260205260409020546001600160a01b0316156111e35760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610491565b6000818152600260205260409020546001600160a01b0316156112485760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610491565b6001600160a01b038216600081815260036020908152604080832080546001019055848352600290915280822080546001600160a01b0319168417905551839291907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b60008072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b83106112f25772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef8100000000831061131e576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061133c57662386f26fc10000830492506010015b6305f5e1008310611354576305f5e100830492506008015b612710831061136857612710830492506004015b6064831061137a576064830492506002015b600a831061035e5760010192915050565b6001600160e01b03198116811461094e57600080fd5b6000602082840312156113b357600080fd5b81356106e28161138b565b60005b838110156113d95781810151838201526020016113c1565b50506000910152565b600081518084526113fa8160208601602086016113be565b601f01601f19169290920160200192915050565b6020815260006106e260208301846113e2565b60006020828403121561143357600080fd5b5035919050565b80356001600160a01b038116811461145157600080fd5b919050565b6000806040838503121561146957600080fd5b6114728361143a565b946020939093013593505050565b60006020828403121561149257600080fd5b6106e28261143a565b6000806000606084860312156114b057600080fd5b6114b98461143a565b92506114c76020850161143a565b9150604084013590509250925092565b634e487b7160e01b600052604160045260246000fd5b600067ffffffffffffffff80841115611508576115086114d7565b604051601f8501601f19908116603f01168101908282118183101715611530576115306114d7565b8160405280935085815286868601111561154957600080fd5b858560208301376000602087830101525050509392505050565b600082601f83011261157457600080fd5b6106e2838335602085016114ed565b60006020828403121561159557600080fd5b813567ffffffffffffffff8111156115ac57600080fd5b6108d084828501611563565b600080604083850312156115cb57600080fd5b823567ffffffffffffffff8111156115e257600080fd5b6115ee85828601611563565b9250506115fd6020840161143a565b90509250929050565b6000806040838503121561161957600080fd5b6116228361143a565b91506020830135801515811461163757600080fd5b809150509250929050565b6000806000806080858703121561165857600080fd5b6116618561143a565b935061166f6020860161143a565b925060408501359150606085013567ffffffffffffffff81111561169257600080fd5b8501601f810187136116a357600080fd5b6116b2878235602084016114ed565b91505092959194509250565b600080604083850312156116d157600080fd5b6116da8361143a565b91506115fd6020840161143a565b600181811c908216806116fc57607f821691505b60208210810361171c57634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252602d908201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560408201526c1c881bdc88185c1c1c9bdd9959609a1b606082015260800190565b601f82111561053257600081815260208120601f850160051c810160208610156117965750805b601f850160051c820191505b818110156117b5578281556001016117a2565b505050505050565b815167ffffffffffffffff8111156117d7576117d76114d7565b6117eb816117e584546116e8565b8461176f565b602080601f83116001811461182057600084156118085750858301515b600019600386901b1c1916600185901b1785556117b5565b600085815260208120601f198616915b8281101561184f57888601518255948401946001909101908401611830565b508582101561186d5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6000835161188f8184602088016113be565b8351908301906118a38183602088016113be565b01949350505050565b60208082526025908201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060408201526437bbb732b960d91b606082015260800190565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b6001600160a01b0385811682528416602082015260408101839052608060608201819052600090611976908301846113e2565b9695505050505050565b60006020828403121561199257600080fd5b81516106e28161138b56fea26469706673582212201832085bbb529268a9d500546d7e1be46c2ad7c45ad0876d78818cf97455f25164736f6c63430008110033",
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
	parsed, err := abi.JSON(strings.NewReader(TrainingtokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
