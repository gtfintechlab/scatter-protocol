// SPDX-License-Identifier: MIT

pragma solidity ^0.8.17;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Capped.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

interface ITrainingJobToken {
    function publishTrainingJob(
        string memory tokenURI,
        string memory topicName,
        address recipient
    ) external;
}

// Model Validator: 500,000 Scatter Token Staked
// Cosmos Validator: 100,000 Scatter Token Staked
contract ScatterProtocol is ERC20Capped, ERC20Burnable {
    enum roles {
        NoRole,
        Requestor,
        Trainer,
        Celestial,
        ModelValidator,
        CosmosValidator
    }

    struct TrainingJobInfo {
        string trainingJobCid;
        address[] trainers;
    }

    address payable public owner;
    address public trainingJobContract;
    address public evaluationJobContract;

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

    mapping(address => string[]) public addressToTopics;

    // Requestor Information
    mapping(address => bool) internal requestorDuplicationChecker;
    address[] public networkRequestors;

    // Trainer Information
    mapping(address => bool) internal trainerDuplicationChecker;
    address[] public networkTrainers;

    // Information for P2P communication
    mapping(address => string) public addressToNodeId;

    constructor(
        uint256 cap,
        uint256 reward,
        address trainingTokenContract,
        address evaluationTokenContract
    ) ERC20("ScatterToken", "ST") ERC20Capped(cap * (10 ** decimals())) {
        owner = payable(msg.sender);
        _mint(owner, (cap / 100) * 70 * (10 ** decimals()));
        blockReward = reward * (10 ** decimals());

        requiredModelValidatorStake = 500000;
        requiredCosmosValidatorStake = 100000;

        trainingJobContract = trainingTokenContract;
        evaluationJobContract = evaluationTokenContract;
    }

    function setNodeId(string memory nodeId) external {
        addressToNodeId[msg.sender] = nodeId;
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

    function initRequestorNode() public {
        addressToRoles[msg.sender] = roles.Requestor;
        updateNetworkParticipants(msg.sender, roles.Requestor);
    }

    function initTrainerNode() public {
        addressToRoles[msg.sender] = roles.Trainer;
        updateNetworkParticipants(msg.sender, roles.Trainer);
    }

    function initCelestialNode() public {
        addressToRoles[msg.sender] = roles.Celestial;
    }

    event TrainingInitialized(address requestor, string topicName);

    function startTraining(string memory topicName) public {
        if (checkIfTopicExistsForRequestor(msg.sender, topicName)) {
            emit TrainingInitialized(msg.sender, topicName);
        }
    }

    function getRoleByAddress(
        address addressToFind
    ) public view returns (string memory) {
        if (addressToRoles[addressToFind] == roles.Requestor) {
            return "requestor";
        }
        if (addressToRoles[addressToFind] == roles.Trainer) {
            return "trainer";
        }
        if (addressToRoles[addressToFind] == roles.Celestial) {
            return "celestial";
        }
        if (addressToRoles[addressToFind] == roles.ModelValidator) {
            return "validator";
        }

        return "no role";
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
            addressToRoles[msg.sender] = roles.NoRole;
        }

        if (
            addressToRoles[msg.sender] == roles.CosmosValidator &&
            addressToStake[msg.sender] < requiredCosmosValidatorStake
        ) {
            addressToRoles[msg.sender] = roles.NoRole;
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

    function requestorAddTopic(
        string memory tokenURI,
        string memory topicName
    ) external {
        if (addressToRoles[msg.sender] == roles.Requestor) {
            ITrainingJobToken(trainingJobContract).publishTrainingJob(
                tokenURI,
                topicName,
                msg.sender
            );
        }
    }

    function trainerAddTopic(
        address requestorAddress,
        string memory requestorTopic
    ) external {
        if (addressToRoles[msg.sender] == roles.Trainer) {
            addressToTrainingJobInfo[requestorAddress][requestorTopic]
                .trainers
                .push(msg.sender);
        }
    }

    function checkIfTopicExistsForRequestor(
        address nodeAddress,
        string memory topicName
    ) public view returns (bool) {
        return
            bytes(
                addressToTrainingJobInfo[nodeAddress][topicName].trainingJobCid
            ).length > 0;
    }

    function getTrainersByAddressAndTopic(
        address requestorAddress,
        string memory topicName,
        int256 skip
    ) external view returns (string memory) {
        string memory trainerList = "";
        TrainingJobInfo storage trainingInfo = addressToTrainingJobInfo[
            requestorAddress
        ][topicName];
        int256 minIndex = int256(trainingInfo.trainers.length) < skip + 10
            ? int256(trainingInfo.trainers.length)
            : skip + 10;

        for (int256 i = skip; i < int256(minIndex); i++) {
            if (trainingInfo.trainers[uint256(i)] == address(0)) {
                continue;
            }

            trainerList = string.concat(
                trainerList,
                Strings.toHexString(
                    uint256(uint160(trainingInfo.trainers[uint256(i)])),
                    20
                )
            );

            trainerList = string.concat(trainerList, "\n");
        }

        return trainerList;
    }

    function addTopicForRequestor(
        string memory topicName,
        string memory jobCid,
        address requestorAddress
    ) external onlyTrainingJobTokenContract hasPeerNodeConfigured {
        require(
            addressToRoles[requestorAddress] == roles.Requestor,
            "You must change your node's role to requestor before calling this method"
        );

        address[] memory emptyAddressArray = new address[](0);

        TrainingJobInfo memory trainingInfo = TrainingJobInfo(
            jobCid,
            emptyAddressArray
        );
        addressToTrainingJobInfo[requestorAddress][topicName] = trainingInfo;

        // Enables trainers to view all topics by a specific network participant
        addressToTopics[requestorAddress].push(topicName);

        updateNetworkParticipants(requestorAddress, roles.Requestor);
    }

    // Displays 10 addresses at a time
    function getRequestorAddresses(
        int256 skip
    ) public view returns (string memory) {
        string memory addressList = "";
        int256 minIndex = int256(networkRequestors.length) < skip + 10
            ? int256(networkRequestors.length)
            : skip + 10;

        for (int256 i = skip; i < int256(minIndex); i++) {
            addressList = string.concat(
                addressList,
                Strings.toHexString(
                    uint256(uint160(networkRequestors[uint256(i)])),
                    20
                )
            );

            addressList = string.concat(addressList, "\n");
        }

        return addressList;
    }

    function getTopicsByRequestorAddress(
        address requestorAddress,
        int256 skip
    ) public view returns (string memory) {
        string memory topicList = "";
        int256 minIndex = int256(addressToTopics[requestorAddress].length) <
            skip + 10
            ? int256(addressToTopics[requestorAddress].length)
            : skip + 10;

        for (int256 i = skip; i < int256(minIndex); i++) {
            topicList = string.concat(
                topicList,
                addressToTopics[requestorAddress][uint256(i)]
            );

            topicList = string.concat(topicList, "\n");
        }

        return topicList;
    }

    function updateNetworkParticipants(
        address participant,
        roles role
    ) internal {
        if (role == roles.Requestor) {
            if (!requestorDuplicationChecker[participant]) {
                networkRequestors.push(participant);
                requestorDuplicationChecker[participant] = true;
            }
        } else if (role == roles.Trainer) {
            if (!trainerDuplicationChecker[participant]) {
                networkTrainers.push(participant);
                trainerDuplicationChecker[participant] = true;
            }
        }
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

    modifier hasPeerNodeConfigured() {
        // require(
        //     bytes(addressToNodeId[msg.sender]).length > 0,
        //     "You must have a node id set for other nodes to be able to communicate with you"
        // );
        _;
    }

    function destroy() public onlyOwner {
        selfdestruct(owner);
    }
}
