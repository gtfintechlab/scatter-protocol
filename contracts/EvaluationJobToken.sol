// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./IScatterProtocol.sol";

contract EvaluationJobToken is ERC721URIStorage, Ownable {
    using Counters for Counters.Counter;
    mapping(uint256 => string) private _tokenURIs;
    mapping(string => address) private jobToAddress;
    mapping(string => string) evaluationJobToDataMapping;
    string baseURI;
    address payable public protocolDeployer;
    address public scatterContractAddress;

    Counters.Counter private _tokenIds;

    constructor() ERC721("Scatter Protocol Evaluation Jobs", "SPEJ") {
        protocolDeployer = payable(msg.sender);
    }

    function setScatterContractAddress(
        address contractAddress
    ) public onlyOwner {
        scatterContractAddress = contractAddress;
    }

    function publishEvaluationJob(
        address recipient,
        string memory tokenURI
    ) external onlyScatterProtocolContract returns (uint256) {
        _tokenIds.increment();

        uint256 newItemId = _tokenIds.current();
        _safeMint(recipient, newItemId);
        setTokenURI(newItemId, tokenURI);
        jobToAddress[tokenURI] = recipient;
        return newItemId;
    }

    function publishEvaluationData(
        address dataPublisher,
        string memory jobURI,
        string memory evaluationDataURI
    ) external onlyScatterProtocolContract {
        require(
            jobToAddress[jobURI] == dataPublisher,
            "This evaluation job does not belong to the data publisher"
        );

        require(
            keccak256(bytes(evaluationJobToDataMapping[jobURI])) ==
                keccak256(bytes("")),
            "This job already has a mapping to it"
        );

        evaluationJobToDataMapping[jobURI] = evaluationDataURI;
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

    modifier onlyScatterProtocolContract() {
        require(
            msg.sender == scatterContractAddress,
            "This method can only be called by the scatter protocol contract"
        );
        _;
    }
}
