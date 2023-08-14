// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

interface IScatterProtocol {
    function addTopicForRequestor(
        string memory topicName,
        string memory jobCid,
        address requestorAddress,
        uint maxTrainerCount
    ) external;
}

contract TrainingJobToken is ERC721URIStorage, Ownable {
    using Counters for Counters.Counter;
    mapping(uint256 => string) private _tokenURIs;
    string baseURI;
    address payable public protocolDeployer;
    address public scatterContractAddress;

    Counters.Counter private _tokenIds;

    constructor() ERC721("Scatter Protocol Training Jobs", "SPTJ") {
        protocolDeployer = payable(msg.sender);
    }

    function setScatterContractAddress(
        address contractAddress
    ) public onlyOwner {
        scatterContractAddress = contractAddress;
    }

    function publishTrainingJob(
        string memory tokenURI,
        string memory topicName
    ) external returns (uint256) {
        _tokenIds.increment();

        uint256 newItemId = _tokenIds.current();
        _safeMint(msg.sender, newItemId);
        setTokenURI(newItemId, tokenURI);

        IScatterProtocol(scatterContractAddress).addTopicForRequestor(
            topicName,
            tokenURI,
            msg.sender,
            100
        );
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
}
