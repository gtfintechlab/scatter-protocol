// SPDX-License-Identifier: MIT

pragma solidity ^0.8.17;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Capped.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "./ITrainingJobToken.sol";
import "./IScatterToken.sol";

import "./Shared.sol";

// Model Validator: 500,000 Scatter Token Staked
contract ScatterProtocol {
    struct TrainingJobInfo {
        string trainingJobCid;
        address[] trainers;
    }

    address payable public owner;
    address public trainingJobContract;
    address public evaluationJobContract;
    address public scatterTokenContract;

    mapping(address => uint256) internal addressToStake;
    mapping(address => uint256) internal addressToStakeTime;
    mapping(address => roles) public addressToRoles;

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
    // We specifically do order and not index because default value of a mapping is 0
    mapping(address => uint256) internal networkRequestorOrder;
    address[] public networkRequestors;

    // Trainer Information
    // We specifically do order and not index because default value of a mapping is 0
    mapping(address => uint256) internal networkTrainerOrder;
    address[] public networkTrainers;

    // Validator Information
    // We specifically do order and not index because default value of a mapping is 0
    mapping(address => uint256) internal networkValidatorOrder;
    address[] public networkValidators;

    // Information for P2P communication
    mapping(address => string) public addressToNodeId;

    // An event that is fire every time a requestor initializes the training procedure
    event TrainingInitialized(address requestor, string topicName);

    /**
     *  @dev Deploys the Protocol with associated contracts
     *  @param trainingTokenContractAddress Contract address for the training job token contract
     *  @param evaluationTokenContractAddress Contract address for the evaluation job token contract
     *  @param scatterTokenContractAddress Contract address for the scatter token contract
     */
    constructor(
        address trainingTokenContractAddress,
        address evaluationTokenContractAddress,
        address scatterTokenContractAddress
    ) {
        owner = payable(msg.sender);
        trainingJobContract = trainingTokenContractAddress;
        evaluationJobContract = evaluationTokenContractAddress;
        scatterTokenContract = scatterTokenContractAddress;
    }

    /**
     *  @dev maps the address to a p2p node id
     *  @param nodeId Node id to set the value to
     */
    function setNodeId(string memory nodeId) external {
        addressToNodeId[msg.sender] = nodeId;
    }

    /**
     *  @dev Sets the training job contract address
     *  @param newAddress The new address value of the training job contract
     */
    function setTrainingJobContractAddress(
        address newAddress
    ) external onlyOwner {
        trainingJobContract = newAddress;
    }

    /**
     *  @dev Sets the evaluation job contract address
     *  @param newAddress The new address value of the evaluation job contract
     */
    function setEvaluationJobContractAddress(
        address newAddress
    ) external onlyOwner {
        evaluationJobContract = newAddress;
    }

    /**
     *  @dev Sets the scatter token contract address
     *  @param newAddress The new address value of the scatter token contract
     */
    function setScatterTokenContractAddress(
        address newAddress
    ) external onlyOwner {
        scatterTokenContract = newAddress;
    }

    /**
     *  @dev Set the current user role to be a requestor
     */
    function initRequestorNode() public {
        // Nothing will happen if they do not have the respective role
        _removeFromOrderedArray(
            networkTrainers,
            networkTrainerOrder,
            msg.sender
        );
        _removeFromOrderedArray(
            networkValidators,
            networkValidatorOrder,
            msg.sender
        );
        addressToRoles[msg.sender] = roles.Requestor;
        _addToOrderedArray(
            networkRequestors,
            networkRequestorOrder,
            msg.sender
        );
    }

    /**
     *  @dev Set the current user role to be a trainer
     */
    function initTrainerNode() public {
        // Nothing will happen if they do not have the respective role
        _removeFromOrderedArray(
            networkValidators,
            networkValidatorOrder,
            msg.sender
        );
        _removeFromOrderedArray(
            networkRequestors,
            networkRequestorOrder,
            msg.sender
        );
        addressToRoles[msg.sender] = roles.Trainer;
        _addToOrderedArray(networkTrainers, networkTrainerOrder, msg.sender);
    }

    /**
     *  @dev Set the current user role to be a trainer
     */
    function initModelValidator() public {
        bool canBePromoted = IScatterToken(scatterTokenContract)
            .canBecomeModelValidator();
        require(
            canBePromoted,
            "Your node is not eligible to become a model validator"
        );
        _removeFromOrderedArray(
            networkRequestors,
            networkRequestorOrder,
            msg.sender
        );
        _removeFromOrderedArray(
            networkTrainers,
            networkTrainerOrder,
            msg.sender
        );
        addressToRoles[msg.sender] = roles.ModelValidator;
        _addToOrderedArray(
            networkValidators,
            networkValidatorOrder,
            msg.sender
        );
    }

    /**
     *  @dev Start the training procedure for a specific topic
     *  @param topicName The topic we want to train the model for
     */
    function startTraining(string memory topicName) public {
        if (checkIfTopicExistsForRequestor(msg.sender, topicName)) {
            emit TrainingInitialized(msg.sender, topicName);
        }
    }

    /**
     *  @dev Set a role for a specific address (only used by scatter token contract)
     *  @param addressToUpdate The address we want to set the role for
     *  @param newRole The new role we want to set this address to
     */
    function setRole(
        address addressToUpdate,
        roles newRole
    ) external onlyScatterTokenContract {
        addressToRoles[addressToUpdate] = newRole;
    }

    /**
     *  @dev Get a role for a specific address (only used by scatter token contract)
     *  @param addressToView The address we want to view
     *  @return role The role of that specific address
     */
    function getEnumRoleByAddress(
        address addressToView
    ) external view onlyScatterTokenContract returns (roles) {
        return addressToRoles[addressToView];
    }

    /**
     *  @dev Get the role for a specific address
     *  @param addressToFind The address we want to get the role for
     *  @return role The role of the specified address
     */
    function getRoleByAddress(
        address addressToFind
    ) public view returns (string memory) {
        if (addressToRoles[addressToFind] == roles.Requestor) {
            return "requestor";
        }
        if (addressToRoles[addressToFind] == roles.Trainer) {
            return "trainer";
        }
        if (addressToRoles[addressToFind] == roles.ModelValidator) {
            return "validator";
        }

        return "no role";
    }

    /**
     *  @dev Create a training job NFT for a specific topic
     *  @param tokenURI The Content ID Hash for the training job zip file
     *  @param topicName The name of the topic
     */
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

    /**
     *  @dev Add a trainer to a specific topic
     *  @param requestorAddress The address of the node that requested the topic
     *  @param requestorTopic The name of the topic
     */
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

    /**
     *  @dev Check if a requestor has a specific topic
     *  @param nodeAddress The address of the node
     *  @param topicName The name of the topic
     */
    function checkIfTopicExistsForRequestor(
        address nodeAddress,
        string memory topicName
    ) public view returns (bool) {
        return
            bytes(
                addressToTrainingJobInfo[nodeAddress][topicName].trainingJobCid
            ).length > 0;
    }

    /**
     *  @dev Return a new line concatenated string with a list of trainers by address and topic
     *  @param requestorAddress The address of the requestor
     *  @param topicName The name of the topic
     *  @param skip number of elements to skip in array
     */
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

    /**
     *  @dev Add a new training job to the state
     *  @param topicName The name of the topic
     *  @param jobCid The content ID hash of the training job files
     *  @param requestorAddress The address of the requestor
     */
    function processTrainingJobToken(
        string memory topicName,
        string memory jobCid,
        address requestorAddress
    ) external onlyTrainingJobTokenContract {
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
    }

    /**
     *  @dev Return a new line concatenated string with a list of addresses
     *  @param skip number of elements to skip in array
     */
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

    /**
     *  @dev Return a new line concatenated string with a list of topics by requestor address
     *  @param requestorAddress The address of the requestor
     *  @param skip number of elements to skip in array
     */
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

    /**
     *  @dev Modifier to ensure only the training job token contract can run a function
     */
    modifier onlyTrainingJobTokenContract() {
        require(
            msg.sender == trainingJobContract,
            "This method can only be called by the training job token contract"
        );
        _;
    }
    /**
     *  @dev Modifier to ensure only the scatter token contract can run a function
     */
    modifier onlyScatterTokenContract() {
        require(
            msg.sender == scatterTokenContract,
            "This method can only be called by the scatter token contract"
        );
        _;
    }
    /**
     *  @dev Modifier to ensure only the owner can run a function
     */
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

    /**
     *  @dev destroy the smart contract
     */
    function destroy() public onlyOwner {
        selfdestruct(owner);
    }

    /**
     *  @dev Helper function to add to an array easily
     * @param array The array we want to add to
     * @param orderMap Keeps track of the order in which the elements were added
     * @param addressToAdd The address we want to add
     */
    function _addToOrderedArray(
        address[] storage array,
        mapping(address => uint256) storage orderMap,
        address addressToAdd
    ) private {
        if (orderMap[addressToAdd] != 0) {
            return;
        }

        array.push(addressToAdd);
        orderMap[addressToAdd] = array.length;
    }

    /**
     *  @dev Helper function to remove from an array easily
     * @param array The array we want to remove from
     * @param orderMap Keeps track of the order in which the elements were added
     * @param addressToRemove The address we want to remove
     */
    function _removeFromOrderedArray(
        address[] storage array,
        mapping(address => uint256) storage orderMap,
        address addressToRemove
    ) private {
        uint256 orderToDelete = orderMap[addressToRemove];
        if (orderToDelete == 0) {
            return;
        }

        uint256 indexToDelete = orderToDelete - 1;

        address lastValue = array[array.length - 1];
        array[indexToDelete] = lastValue;
        orderMap[lastValue] = orderToDelete;

        array.pop();
        delete orderMap[addressToRemove];
    }
}
