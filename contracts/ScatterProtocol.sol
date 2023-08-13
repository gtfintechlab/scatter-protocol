// SPDX-License-Identifier: GPL-2.0-or-later

pragma solidity ^0.8.17;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Capped.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";

// Model Validator: 500,000 Scatter Token Staked
// Cosmos Validator: 100,000 Scatter Token Staked
contract ScatterProtocol is ERC20Capped, ERC20Burnable {
    enum roles {
        Peer,
        Celestial,
        ModelValidator,
        CosmosValidator
    }

    struct TrainingJobInfo {
        string trainingJobCid;
        address[] trainers;
    }

    address payable public owner;
    address trainingJobContract;
    address evaluationJobContract;

    uint256 public blockReward;
    uint256 public requiredModelValidatorStake;
    uint256 public requiredCosmosValidatorStake;

    mapping(address => uint256) internal addressToStake;
    mapping(address => uint256) internal addressToStakeTime;
    mapping(address => roles) internal addressToRoles;

    /*
        Example addressToTrainingJobInfo Mapping:
        {
            requestor 1 address: {
                topic 1: {
                    trainingJobCid: CID of the training job location, 
                    trainers: [trainer 1 address, trainer 2 address]
                }
            }
            requestor 2 address: {...}
        }
     */
    mapping(address => mapping(string => TrainingJobInfo))
        public addressToTrainingJobInfo;

    constructor(
        uint256 cap,
        uint256 reward,
        address trainingTokenContractAddress,
        address evaluationTokenContractAddress
    ) ERC20("ScatterToken", "ST") ERC20Capped(cap * (10 ** decimals())) {
        owner = payable(msg.sender);
        _mint(owner, (cap / 100) * 70 * (10 ** decimals()));
        blockReward = reward * (10 ** decimals());

        requiredModelValidatorStake = 500000;
        requiredCosmosValidatorStake = 100000;

        trainingJobContract = trainingTokenContractAddress;
        evaluationJobContract = evaluationTokenContractAddress;
    }

    function setTrainingJobContractAddress(
        address newAddress
    ) external onlyOwner {
        trainingJobContract = newAddress;
    }

    function setEvaluationJobContractAddress(
        address newAddress
    ) external onlyOwner {
        evaluationJobContract = newAddress;
    }

    function initPeerNode() public {
        addressToRoles[msg.sender] = roles.Peer;
    }

    function initCelestialNode() public {
        addressToRoles[msg.sender] = roles.Celestial;
    }

    function elevateToModelValidator() public {
        require(
            addressToStake[msg.sender] >= requiredModelValidatorStake,
            "To become a model validator, you must have 500,000 Scatter Token staked"
        );

        addressToRoles[msg.sender] = roles.ModelValidator;
    }

    function elevateToCosmosValidator() public {
        require(
            addressToStake[msg.sender] >= requiredCosmosValidatorStake,
            "To become a cosmos validator, you must have 100,000 Scatter Token staked"
        );

        addressToRoles[msg.sender] = roles.CosmosValidator;
    }

    function addStake(uint256 amount) public {
        require(amount > 0, "Stake amount must be positive");
        _burn(msg.sender, amount);
        addressToStake[msg.sender] += amount;
        // Anytime more stake is added, stake timer increases
        addressToStakeTime[msg.sender] = block.timestamp + 30 days;
    }

    function removeStake(uint256 amount) public {
        require(
            addressToStakeTime[msg.sender] >= block.timestamp,
            "You must wait 30 days before being able to unstake Scatter Token"
        );
        require(
            amount <= addressToStake[msg.sender],
            "Amount must be less than staked amount"
        );
        require(amount > 0, "Stake amount must be positive");

        _mint(msg.sender, amount);
        addressToStake[msg.sender] -= amount;

        if (
            addressToRoles[msg.sender] == roles.ModelValidator &&
            addressToStake[msg.sender] < requiredModelValidatorStake
        ) {
            addressToRoles[msg.sender] = roles.Peer;
        }

        if (
            addressToRoles[msg.sender] == roles.CosmosValidator &&
            addressToStake[msg.sender] < requiredCosmosValidatorStake
        ) {
            addressToRoles[msg.sender] = roles.Peer;
        }
    }

    function getOwnStake() public view returns (uint256) {
        return addressToStake[msg.sender];
    }

    function getAccountStake(address account) public view returns (uint256) {
        return addressToStake[account];
    }

    function _mint(
        address account,
        uint256 amount
    ) internal virtual override(ERC20Capped, ERC20) {
        require(
            ERC20.totalSupply() + amount <= cap(),
            "ERC20Capped: cap exceeded"
        );
        super._mint(account, amount);
    }

    function setBlockReward(uint256 reward) public onlyOwner {
        blockReward = reward * (10 ** decimals());
    }

    function _mintMinerReward() internal {
        _mint(block.coinbase, blockReward);
    }

    function _beforeTokenTransfer(
        address from,
        address to,
        uint256 value
    ) internal virtual override {
        if (
            from != address(0) &&
            to != block.coinbase &&
            block.coinbase != address(0)
        ) {
            _mintMinerReward();
        }
        super._beforeTokenTransfer(from, to, value);
    }

    function addTopicForRequestor(
        string memory topicName,
        string memory jobCid,
        address requestorAddress,
        uint maxTrainerCount
    ) external onlyTrainingJobTokenContract {
        address[] memory emptyAddressArray = new address[](maxTrainerCount);

        TrainingJobInfo memory trainingInfo = TrainingJobInfo(
            jobCid,
            emptyAddressArray
        );
        addressToTrainingJobInfo[requestorAddress][topicName] = trainingInfo;
    }

    modifier onlyTrainingJobTokenContract() {
        require(
            msg.sender == trainingJobContract,
            "This method can only be called by the training job token contract"
        );
        _;
    }

    modifier onlyOwner() {
        require(
            payable(msg.sender) == owner,
            "Only the owner can call this function"
        );
        _;
    }

    function destroy() public onlyOwner {
        selfdestruct(owner);
    }
}
