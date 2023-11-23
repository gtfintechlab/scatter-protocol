// SPDX-License-Identifier: MIT

pragma solidity ^0.8.17;

import "@openzeppelin/contracts/utils/Strings.sol";
import "./ITrainingJobToken.sol";
import "./IEvaluationJobToken.sol";
import "./IScatterToken.sol";
import "./IModelToken.sol";
import "./Shared.sol";

// Model Validator: 10,000 Scatter Token Staked
contract ScatterProtocol {
    struct FederatedJob {
        string trainingJobCid;
        address[] trainers;
        uint256 publishedTrainerCount;
        uint256 pooledReward;
        string evaluationJobCid;
        string evaluationJobDataCid;
        uint256 jobTerminationDate;
        uint256 validatorValidationCount;
        uint256 jobStartDate;
    }

    address payable public owner;
    address public trainingJobContract;
    address public evaluationJobContract;
    address public scatterTokenContract;
    address public modelTokenContract;

    mapping(address => uint256) internal addressToStake;
    mapping(address => uint256) internal addressToStakeTime;
    mapping(address => roles) public addressToRoles;

    /*
        Example addressToFederatedJob Mapping:
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
    mapping(address => mapping(string => FederatedJob))
        public addressToFederatedJob;

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

    /*
        Example trainerTrainingMap & validatorTrainingMap Mapping:
        {
            requestor 1 address: {
                topic 1: {
                    trainer / validator 1: true,
                    trainer 2 / validator 2: false,
                }
            }
            requestor 2 address: {...}
        }
     */
    // Used to check if a trainer is already subscribed to a topic
    mapping(address => mapping(string => mapping(address => bool)))
        public trainerTrainingMap;
    mapping(address => mapping(string => address[])) trainerTrainingMapKeys;

    // Mapping of validators for each topic
    mapping(address => mapping(string => mapping(address => bool)))
        public validatorTrainingMap;
    mapping(address => mapping(string => address[])) validatorTrainingMapKeys;

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

    // Mappings used to keep track of evaluaation metrics
    /*
        Example evaluationScore Mapping:
        {
            requestor address: {
                topic name: {
                    validator 1 address: {
                        trainer 1 address: score 1,
                        trainer 2 address: score 2
                    },
                    validator 2 address: {...}
                    }
                }
            }
        }
    */
    mapping(address => mapping(string => mapping(address => mapping(address => uint256)))) evaluationScore;
    mapping(address => mapping(string => mapping(address => mapping(address => bool)))) evaluationScoreSet;

    // Rogue + Benevolent trainers used for reward mechanism
    mapping(address => mapping(string => address[])) rogueTrainers;
    mapping(address => mapping(string => mapping(address => bool))) rogueTrainerMapping;
    mapping(address => mapping(string => address[])) benevolentTrainers;
    mapping(address => mapping(string => mapping(address => bool))) benevolentTrainerMapping;

    // Rogue + Benevolent validators used for reward mechanism
    mapping(address => mapping(string => address[])) rogueValidators;
    mapping(address => mapping(string => mapping(address => bool))) rogueValidatorMapping;
    mapping(address => mapping(string => address[])) benevolentValidators;
    mapping(address => mapping(string => mapping(address => bool))) benevolentValidatorMapping;

    // Random nonce to generate pseudorandom number
    uint randomNonce = 1;

    // Percent of validators we want to use to validate models
    uint256 modelValidationPercentage = 100;

    // An event that is fire every time a requestor initializes the training procedure
    event TrainingInitialized(address requestor, string topicName);

    // An event fired by the trainer when they are ready to send their model over to the validators
    event ModelReadyToValidate(address requestor, string topicName);

    // An event that is fired when all of the trainers have submitted their models and are ready to be evaluated
    event RequestForEvaluationSet(address requestor, string topicName);

    // An event used for debugging purposes
    event DebugEvent(string message);

    /**
     *  @dev Deploys the Protocol with associated contracts
     *  @param trainingTokenContractAddress Contract address for the training job token contract
     *  @param evaluationTokenContractAddress Contract address for the evaluation job token contract
     *  @param scatterTokenContractAddress Contract address for the scatter token contract
     */
    constructor(
        address trainingTokenContractAddress,
        address evaluationTokenContractAddress,
        address scatterTokenContractAddress,
        address modelTokenContractAddress
    ) {
        owner = payable(msg.sender);
        trainingJobContract = trainingTokenContractAddress;
        evaluationJobContract = evaluationTokenContractAddress;
        scatterTokenContract = scatterTokenContractAddress;
        modelTokenContract = modelTokenContractAddress;
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
    function initValidatorNode() public {
        bool canBePromoted = IScatterToken(scatterTokenContract)
            .canBecomeValidator(msg.sender);

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
        addressToRoles[msg.sender] = roles.Validator;
        _addToOrderedArray(
            networkValidators,
            networkValidatorOrder,
            msg.sender
        );
    }

    function rewardDistributor(
        address requestorAddress,
        string memory topicName
    ) public {
        bool distributeRewards = false;
        // First we must check if the job's time limit is up
        uint256 terminationTime = addressToFederatedJob[requestorAddress][
            topicName
        ].jobTerminationDate;

        if (block.timestamp >= terminationTime && terminationTime != 0) {
            distributeRewards = true;
        }

        // Next we must check if all of the trainers have submitted their models
        bool allModelsSubmitted = _checkTrainerModelSubmissions(
            requestorAddress,
            topicName
        );

        // Next we must check that all of the validators have validated the models
        bool allModelsValidated = _checkValidatorModelValidations(
            requestorAddress,
            topicName
        );

        // We distribute rewards when all models are validated and submitted or the job is over time
        distributeRewards =
            distributeRewards ||
            (allModelsSubmitted && allModelsValidated);

        // 2% towards lottery for challengers
        // Get Rogue Trainers - Slash 100% of what they staked --> lottery
        // Get Rogue Validators - Slash 10% of their stake --> lottery
        // Reward benevolent validators - reward should be proportional to their stake
        // Return trainer staked token to benevolent trainers
        // Reward benevolent trainers - reward should be proportional to their short-term stake & model score
        if (distributeRewards) {
            _populateTrainerStatus(requestorAddress, topicName);
            _populateValidatorStatus(requestorAddress, topicName);

            address[] memory rogueTrainerList = rogueTrainers[requestorAddress][
                topicName
            ];
            address[] memory rogueValidatorList = rogueValidators[
                requestorAddress
            ][topicName];

            address[] memory benevolentTrainerList = benevolentTrainers[
                requestorAddress
            ][topicName];
            address[] memory benevolentValidatorList = benevolentValidators[
                requestorAddress
            ][topicName];

            IScatterToken(scatterTokenContract).donateToLottery(
                requestorAddress,
                topicName
            );

            IScatterToken(scatterTokenContract).punishRogueTrainers(
                requestorAddress,
                topicName,
                rogueTrainerList
            );

            IScatterToken(scatterTokenContract).punishRogueValidators(
                rogueValidatorList
            );

            IScatterToken(scatterTokenContract).rewardBenevolentTrainers(
                requestorAddress,
                topicName,
                benevolentTrainerList
            );

            IScatterToken(scatterTokenContract).rewardBenevolentValidators(
                requestorAddress,
                topicName,
                benevolentValidatorList
            );
        }
    }

    /**
     *  @dev Start the training procedure for a specific topic - chooses random validators
     *  @param topicName The topic we want to train the model for
     */
    function startTraining(string memory topicName) external isRequestor {
        require(
            checkIfTopicExistsForRequestor(msg.sender, topicName),
            "This topic has not been added and therefore cannot start a training job"
        );

        uint256 validatorCount = uint256(
            (networkValidators.length * modelValidationPercentage) / 100
        );

        address[] memory chosenValidators = _getRandomAddressSubset(
            networkValidators,
            int256(validatorCount)
        );

        for (uint i = 0; i < chosenValidators.length; i++) {
            validatorTrainingMap[msg.sender][topicName][
                chosenValidators[i]
            ] = true;
            validatorTrainingMapKeys[msg.sender][topicName].push(
                chosenValidators[i]
            );
        }

        emit TrainingInitialized(msg.sender, topicName);
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
        if (addressToRoles[addressToFind] == roles.Validator) {
            return "validator";
        }

        if (addressToRoles[addressToFind] == roles.Challenger) {
            return "challenge";
        }

        return "no role";
    }

    /**
     *  @dev Create a training job NFT for a specific topic + locks up the respective amount of token
     *  @param trainingTokenURI The Content ID Hash for the training job zip file
     *  @param evaluationTokenURI The Content ID Hash for the evaluation job zip file
     *  @param topicName The name of the topic
     *  @param pooledReward The reward given across all validators, trainers, etc.
     */
    function requestorAddTopic(
        string memory trainingTokenURI,
        string memory evaluationTokenURI,
        string memory topicName,
        uint256 pooledReward
    ) external isRequestor {
        ITrainingJobToken(trainingJobContract).publishTrainingJob(
            trainingTokenURI,
            msg.sender
        );

        IEvaluationJobToken(evaluationJobContract).publishEvaluationJob(
            msg.sender,
            evaluationTokenURI
        );

        address[] memory emptyAddressArray = new address[](0);
        FederatedJob memory trainingInfo = FederatedJob(
            trainingTokenURI,
            emptyAddressArray,
            0,
            pooledReward,
            evaluationTokenURI, // evaluation job does not include the evaluation data --> just the job itself
            "",
            block.timestamp +
                getTimeFromCheckpoint(Checkpoints.FederatedJobEnd), // All jobs must be completed within 30 days of submission
            0,
            block.timestamp
        );
        addressToFederatedJob[msg.sender][topicName] = trainingInfo;
        // Enables trainers to view all topics by a specific network participant
        addressToTopics[msg.sender].push(topicName);
        IScatterToken(scatterTokenContract).requestorLockToken(
            msg.sender,
            topicName,
            pooledReward
        );
    }

    /**
     *  @dev Add a trainer to a specific topic
     *  @param requestorAddress The address of the node that requested the topic
     *  @param requestorTopic The name of the topic
     */
    function trainerAddTopic(
        address requestorAddress,
        string memory requestorTopic,
        uint256 stakeAmount
    ) external isTrainer {
        bool alreadyInTraining = trainerTrainingMap[requestorAddress][
            requestorTopic
        ][msg.sender];

        if (alreadyInTraining) {
            return;
        }

        IScatterToken(scatterTokenContract).trainerLockToken(
            msg.sender,
            requestorAddress,
            requestorTopic,
            stakeAmount
        );
        addressToFederatedJob[requestorAddress][requestorTopic].trainers.push(
            msg.sender
        );

        trainerTrainingMap[requestorAddress][requestorTopic][msg.sender] = true;
        trainerTrainingMapKeys[requestorAddress][requestorTopic].push(
            msg.sender
        );
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
            bytes(addressToFederatedJob[nodeAddress][topicName].trainingJobCid)
                .length > 0;
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
        FederatedJob storage trainingInfo = addressToFederatedJob[
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

    function publishModelToProtocol(
        string memory modelURI,
        address requestorAddress,
        string memory topicName
    ) external isTrainer {
        require(
            Strings.equal(
                modelLogger[requestorAddress][topicName][msg.sender],
                ""
            ),
            "You have already published a model - cannot publish again!"
        );

        bool isSubscribed = trainerTrainingMap[requestorAddress][topicName][
            msg.sender
        ];
        require(
            isSubscribed,
            "You cannot submit a model to this topic when you are not subscribed to it"
        );

        IModelToken(modelTokenContract).publishModel(modelURI, msg.sender);
        modelLogger[requestorAddress][topicName][msg.sender] = modelURI;

        FederatedJob storage info = addressToFederatedJob[requestorAddress][
            topicName
        ];
        info.publishedTrainerCount += 1;

        // Check if we are ready to validate all the models from all of the trainers
        // If the current block timestamp is after the model submission threshold
        // we will emit the event.
        address[] memory trainers = info.trainers;
        bool trainersNotReady = (trainers.length == 0 ||
            trainers.length != info.publishedTrainerCount);
        bool withinTimeThreshold = block.timestamp <=
            info.jobStartDate +
                getTimeFromCheckpoint(Checkpoints.ModelSubmission);

        if (trainersNotReady && withinTimeThreshold) {
            return;
        }

        // Only emit a request for the evaluation data once we know all the trainers have trained their model
        // This ensures that trainers do not exploit the evaluation data to train their model beforehand
        emit RequestForEvaluationSet(requestorAddress, topicName);
    }

    function getTimeFromCheckpoint(
        Checkpoints checkpoint
    ) internal pure returns (uint256) {
        if (checkpoint == Checkpoints.ModelSubmission) {
            return 15 days;
        }

        if (checkpoint == Checkpoints.FederatedJobEnd) {
            return 30 days;
        }

        return 0 days;
    }

    function submitEvaluationSet(
        string memory topicName,
        string memory evaluationDataURI
    ) public isRequestor {
        IEvaluationJobToken(evaluationJobContract).publishEvaluationData(
            msg.sender,
            addressToFederatedJob[msg.sender][topicName].evaluationJobCid,
            evaluationDataURI
        );
        addressToFederatedJob[msg.sender][topicName]
            .evaluationJobDataCid = evaluationDataURI;
        emit ModelReadyToValidate(msg.sender, topicName);
    }

    function submitEvaluationScore(
        address requestorAddress,
        string memory topicName,
        address trainerAddress,
        uint256 score
    ) public isValidator {
        require(
            !evaluationScoreSet[requestorAddress][topicName][msg.sender][
                trainerAddress
            ],
            "Cannot resubmit a score for a trainer that has previously been submitted"
        );

        require(
            validatorTrainingMap[requestorAddress][topicName][msg.sender],
            "You are not an assigned validator for this training job"
        );

        addressToFederatedJob[requestorAddress][topicName]
            .validatorValidationCount += 1;

        evaluationScore[requestorAddress][topicName][msg.sender][
            trainerAddress
        ] = score;

        evaluationScoreSet[requestorAddress][topicName][msg.sender][
            trainerAddress
        ] = true;
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

    modifier isRequestor() {
        require(
            addressToRoles[msg.sender] == roles.Requestor,
            "You must change your node's role to requestor before calling this method"
        );
        _;
    }

    modifier isTrainer() {
        require(
            addressToRoles[msg.sender] == roles.Trainer,
            "You must change your node's role to trainer before calling this method"
        );
        _;
    }

    modifier isValidator() {
        require(
            addressToRoles[msg.sender] == roles.Validator,
            "You must change your node's role to validator before calling this method"
        );
        _;
    }

    modifier isChallenger() {
        require(
            addressToRoles[msg.sender] == roles.Challenger,
            "You must change your node's role to challenger before calling this method"
        );
        _;
    }

    /**
     *  @dev destroy the smart contract
     */
    function destroy() public onlyOwner {
        selfdestruct(owner);
    }

    function _checkTrainerModelSubmissions(
        address requestorAddress,
        string memory topicName
    ) private view returns (bool) {
        address[] memory trainerList = addressToFederatedJob[requestorAddress][
            topicName
        ].trainers;
        uint256 trainersPublishedCount = addressToFederatedJob[
            requestorAddress
        ][topicName].publishedTrainerCount;
        return trainerList.length == trainersPublishedCount;
    }

    function _checkValidatorModelValidations(
        address requestorAddress,
        string memory topicName
    ) private view returns (bool) {
        address[] memory validatorList = validatorTrainingMapKeys[
            requestorAddress
        ][topicName];
        address[] memory trainerList = trainerTrainingMapKeys[requestorAddress][
            topicName
        ];

        uint256 validatorValidationCount = addressToFederatedJob[
            requestorAddress
        ][topicName].validatorValidationCount;

        return
            validatorValidationCount ==
            (validatorList.length * trainerList.length);
    }

    function _populateTrainerStatus(
        address requestorAddress,
        string memory topicName
    ) private {
        address[] memory trainers = addressToFederatedJob[requestorAddress][
            topicName
        ].trainers;
        mapping(address => string) storage trainerModelMap = modelLogger[
            requestorAddress
        ][topicName];

        for (uint i = 0; i < trainers.length; i++) {
            if (Strings.equal(trainerModelMap[trainers[i]], "")) {
                rogueTrainers[requestorAddress][topicName].push(trainers[i]);
                rogueTrainerMapping[requestorAddress][topicName][
                    trainers[i]
                ] = true;
            } else {
                benevolentTrainers[requestorAddress][topicName].push(
                    trainers[i]
                );
                benevolentTrainerMapping[requestorAddress][topicName][
                    trainers[i]
                ] = true;
            }
        }
    }

    function _populateValidatorStatus(
        address requestorAddress,
        string memory topicName
    ) private {
        address[] memory validatorList = validatorTrainingMapKeys[
            requestorAddress
        ][topicName];
        address[] memory trainerList = trainerTrainingMapKeys[requestorAddress][
            topicName
        ];

        for (uint i = 0; i < validatorList.length; i++) {
            for (uint j = 0; j < trainerList.length; j++) {
                // We do not care about rogue trainers - skip them entirely
                if (
                    rogueTrainerMapping[requestorAddress][topicName][
                        trainerList[j]
                    ]
                ) {
                    continue;
                }

                // For a non-rogue trainer, check if their evaluation scores have been set
                // Mark them as rogue if not
                if (
                    !evaluationScoreSet[requestorAddress][topicName][
                        validatorList[i]
                    ][trainerList[j]]
                ) {
                    rogueValidators[requestorAddress][topicName].push(
                        validatorList[i]
                    );
                    rogueValidatorMapping[requestorAddress][topicName][
                        validatorList[i]
                    ] = true;
                } else {
                    benevolentValidators[requestorAddress][topicName].push(
                        validatorList[i]
                    );
                    benevolentValidatorMapping[requestorAddress][topicName][
                        validatorList[i]
                    ] = true;
                }
            }
        }
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

    function _pseudoRandomNumber() private returns (uint) {
        randomNonce++;
        return
            uint(
                keccak256(
                    abi.encodePacked(
                        block.difficulty,
                        block.timestamp,
                        networkValidators,
                        randomNonce
                    )
                )
            );
    }

    function _getRandomAddressSubset(
        address[] memory arr,
        int256 count
    ) private returns (address[] memory) {
        if (count <= 0) {
            return new address[](0);
        }

        require(
            uint256(count) <= arr.length,
            "Count should be less than or equal to array length"
        );

        address[] memory shuffled = _shuffle(arr);
        address[] memory result = new address[](uint256(count));

        for (uint256 i = 0; i < uint256(count); i++) {
            result[i] = shuffled[i];
        }

        return result;
    }

    // Fisher-Yates shuffle algorithm
    function _shuffle(address[] memory arr) private returns (address[] memory) {
        address[] memory shuffled = arr;
        uint256 n = shuffled.length;

        for (uint256 i = n - 1; i > 0; i--) {
            uint256 j = _pseudoRandomNumber() % (i + 1);
            (shuffled[i], shuffled[j]) = (shuffled[j], shuffled[i]);
        }

        return shuffled;
    }
}
