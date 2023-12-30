// SPDX-License-Identifier: MIT

pragma solidity >=0.8.17;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./IScatterProtocol.sol";
import "./IModelToken.sol";

contract ModelToken is ERC721URIStorage, Ownable, IModelToken {
    using Counters for Counters.Counter;
    mapping(uint256 => string) private _tokenURIs;
    string baseURI;
    address payable public protocolDeployer;
    address public scatterContractAddress;

    Counters.Counter private _tokenIds;
    /*
        Example modelLogger Mapping:
        {
            requestor address: {
                topic name: {
                    trainer 1 address: CID 1,
                    trainer 2 address: CID 1,
                }
            }
        }
    */
    // Allows the protocol to keep track of model URIs for trainers
    mapping(address => mapping(string => mapping(address => string)))
        public modelLogger;

    constructor() ERC721("Scatter Protocol Models", "SPM") {
        protocolDeployer = payable(msg.sender);
    }

    function setScatterContractAddress(
        address contractAddress
    ) public onlyOwner {
        scatterContractAddress = contractAddress;
    }

    function publishModel(
        string memory tokenURI,
        address recipient,
        address requestorAddress,
        string memory topicName
    ) external onlyScatterProtocolContract returns (uint256) {
        _tokenIds.increment();

        uint256 newItemId = _tokenIds.current();
        _safeMint(recipient, newItemId);
        setTokenURI(newItemId, tokenURI);

        modelLogger[requestorAddress][topicName][recipient] = tokenURI;
        return newItemId;
    }

    function setTokenURI(uint256 tokenId, string memory tokenURI) internal {
        _setTokenURI(tokenId, tokenURI);
    }

    function getTokenURI(
        uint256 tokenId
    ) external view returns (string memory) {
        return tokenURI(tokenId);
    }

    function setBaseURI(string memory baseURI_) external onlyOwner {
        baseURI = baseURI_;
    }

    function _baseURI() internal view override returns (string memory) {
        return baseURI;
    }

    function getModelCidForTrainer(
        address requestorAddress,
        string memory topicName,
        address trainer
    ) external view returns (string memory) {
        return modelLogger[requestorAddress][topicName][trainer];
    }

    modifier onlyScatterProtocolContract() {
        require(
            msg.sender == scatterContractAddress,
            "This method can only be called by the scatter protocol contract"
        );
        _;
    }
}
