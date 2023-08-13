// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract EvaluationJobToken is ERC721URIStorage, Ownable {
    using Counters for Counters.Counter;
    mapping(uint256 => string) private _tokenURIs;
    string baseURI;
    address payable public protocolDeployer;

    Counters.Counter private _tokenIds;

    constructor() ERC721("Scatter Protocol Evaluation Jobs", "SPTJ") {
        protocolDeployer = payable(msg.sender);
    }

    function publishEvaluationJob(
        address recipient,
        string memory tokenURI
    ) external returns (uint256) {
        _tokenIds.increment();

        uint256 newItemId = _tokenIds.current();
        _safeMint(recipient, newItemId);
        setTokenURI(newItemId, tokenURI);
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

    function setBaseURI(string memory baseURI_) external {
        require(
            msg.sender == protocolDeployer,
            "Only the owner can access this method"
        );

        baseURI = baseURI_;
    }

    function _baseURI() internal view override returns (string memory) {
        return baseURI;
    }
}
