// SPDX-License-Identifier: MIT

pragma solidity >=0.8.17;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Capped.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/utils/math/Math.sol";
import "hardhat/console.sol";

import "./Shared.sol";
import "./ScatterMath.sol";

import "./IScatterProtocol.sol";
import "./IScatterToken.sol";
import "./IReputationManager.sol";
import "./IEvaluationJobToken.sol";
import "./IVoteManager.sol";

contract ScatterToken is
    ERC20Capped,
    ERC20Burnable,
    IScatterToken,
    ReentrancyGuard
{
    uint256 public requiredModelValidatorStake;
    uint256 public requiredModelChallengerStake;
    address payable public owner;

    // Long Term Stake
    mapping(address => uint256) public addressToStake;
    mapping(address => uint256) internal addressToStakeTime;

    // Short Term Stakes for trainers
    mapping(address => mapping(string => mapping(address => uint256))) trainerLockedTokenForTrainingJob;

    // Reward pool for requestors
    mapping(address => mapping(string => uint256)) requestorLockedTokenForTrainingJob;

    // Lottery for challenges who prove a validator wrong
    uint256 lotteryPool = 0;

    // Reward Factors - should add up to 100
    uint256 lotteryPercentage = 10;
    uint256 trainerRewardProportion = 70;
    uint256 validatorRewardProportion = 20;

    // Punishment constants
    uint256 validatorBasePenalty = 200;
    uint256 validatorPenaltyMultiplier = 2;
    mapping(address => uint256) validatorPenaltyCount;

    IScatterProtocol scatterProtocolContract;
    IReputationManager reputationManagerContract;
    IEvaluationJobToken evaluationJobTokenContract;
    IVoteManager voteManagerContract;

    event TokensStaked(address from, uint256 amount);
    event TokensUnstaked(address to, uint256 amount);

    constructor(uint256 cap) ERC20("ScatterToken", "ST") ERC20Capped(cap) {
        owner = payable(msg.sender);
        _mint(owner, cap);

        requiredModelValidatorStake = 25000;
        requiredModelChallengerStake = 20000;
    }

    function setScatterProtocolAddress(address newAddress) public onlyOwner {
        scatterProtocolContract = IScatterProtocol(newAddress);
    }

    function setEvaluationJobTokenContract(
        address newAddress
    ) public onlyOwner {
        evaluationJobTokenContract = IEvaluationJobToken(newAddress);
    }

    function setReputationManagerContract(address newAddress) public onlyOwner {
        reputationManagerContract = IReputationManager(newAddress);
    }

    function setVoteManagerContract(address newAddress) public onlyOwner {
        voteManagerContract = IVoteManager(newAddress);
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

    function requestorLockToken(
        address requestorAddress,
        string memory topicName,
        uint256 amount
    ) external onlyScatterProtocolContract {
        require(
            amount <= this.balanceOf(requestorAddress),
            "Cannot lock more tokens than you own"
        );
        _burn(requestorAddress, amount);
        requestorLockedTokenForTrainingJob[requestorAddress][
            topicName
        ] = amount;
    }

    function trainerLockToken(
        address trainerAddress,
        address requestorAddress,
        string memory topicName,
        uint256 amount
    ) external onlyScatterProtocolContract {
        require(
            amount <= this.balanceOf(trainerAddress),
            "Cannot lock more tokens than you own"
        );

        require(amount > 0, "Must lock some amount of tokens");

        _burn(trainerAddress, amount);
        trainerLockedTokenForTrainingJob[requestorAddress][topicName][
            trainerAddress
        ] += amount;
    }

    function stakeToken(uint256 amount) public nonReentrant {
        require(
            amount <= this.balanceOf(msg.sender),
            "Cannot stake more tokens than you own"
        );
        _burn(msg.sender, amount);
        addressToStake[msg.sender] += amount;

        // Anytime more stake is added, stake timer increases
        addressToStakeTime[msg.sender] = block.timestamp + 30 days;
        emit TokensStaked(msg.sender, amount);
    }

    function removeStake(uint256 amount) public nonReentrant {
        require(
            addressToStakeTime[msg.sender] >= block.timestamp,
            "You must wait before being able to unstake Scatter Token"
        );
        require(
            amount < addressToStake[msg.sender],
            "Amount must be less than staked amount"
        );

        _mint(msg.sender, amount);
        addressToStake[msg.sender] -= amount;

        roles messageSenderRole = scatterProtocolContract.getEnumRoleByAddress(
            msg.sender
        );
        if (
            messageSenderRole == roles.Validator &&
            addressToStake[msg.sender] < requiredModelValidatorStake
        ) {
            scatterProtocolContract.setRole(msg.sender, roles.NoRole);
        }
    }

    function getOwnStake() public view returns (uint256) {
        return addressToStake[msg.sender];
    }

    function getAccountStake(address account) public view returns (uint256) {
        return addressToStake[account];
    }

    function canBecomeValidator(address account) external view returns (bool) {
        return addressToStake[account] >= requiredModelValidatorStake;
    }

    function canBecomeChallenger(address account) external view returns (bool) {
        return addressToStake[account] >= requiredModelChallengerStake;
    }

    function donateToLottery(
        address requestorAddress,
        string memory topicName
    ) external onlyScatterProtocolContract {
        uint256 requestorPooledReward = requestorLockedTokenForTrainingJob[
            requestorAddress
        ][topicName];

        uint256 toLottery = (requestorPooledReward * lotteryPercentage) / 100;
        lotteryPool += toLottery;
    }

    function punishRogueTrainers(
        address requestorAddress,
        string memory topicName
    ) external onlyScatterProtocolContract {
        address[] memory trainers = reputationManagerContract
            .getMalevolentTrainers(requestorAddress, topicName);

        for (uint i = 0; i < trainers.length; i++) {
            if (trainers[i] == address(0x0)) {
                break;
            }

            lotteryPool += trainerLockedTokenForTrainingJob[requestorAddress][
                topicName
            ][trainers[i]];
            trainerLockedTokenForTrainingJob[requestorAddress][topicName][
                trainers[i]
            ] = 0;
        }
    }

    function punishRogueValidators(
        address requestorAddress,
        string memory topicName
    ) external onlyScatterProtocolContract {
        address[] memory validators = reputationManagerContract
            .getMalevolentValidators(requestorAddress, topicName);

        for (uint i = 0; i < validators.length; i++) {
            if (validators[i] == address(0)) {
                break;
            }
            uint256 calculatedPenalty = validatorBasePenalty *
                (validatorPenaltyMultiplier **
                    validatorPenaltyCount[validators[i]]);

            uint256 punishmentAmount = Math.min(
                addressToStake[validators[i]],
                calculatedPenalty
            );

            lotteryPool += punishmentAmount;
            addressToStake[validators[i]] -= punishmentAmount;
            validatorPenaltyCount[validators[i]] += 1;

            // Automatically demote validators once their stake reaches 0 tokens
            if (addressToStake[validators[i]] <= 0) {
                scatterProtocolContract.setRole(msg.sender, roles.NoRole);
            }
        }
    }

    function rewardBenevolentTrainers(
        address requestorAddress,
        string memory topicName
    ) external onlyScatterProtocolContract {
        address[] memory trainers = reputationManagerContract
            .getBenevolentTrainers(requestorAddress, topicName);

        uint totalRewardPool = (requestorLockedTokenForTrainingJob[
            requestorAddress
        ][topicName] * trainerRewardProportion) / 100;

        uint totalWeight = 0;
        for (uint i = 0; i < trainers.length; i++) {
            if (trainers[i] == address(0x0)) {
                break;
            }
            uint performanceWeight = evaluationJobTokenContract
                .getAverageScoreForTrainerForJob(
                    requestorAddress,
                    topicName,
                    trainers[i]
                ) ** 2;

            uint stakedToken = trainerLockedTokenForTrainingJob[
                requestorAddress
            ][topicName][trainers[i]];

            uint stakeWeight = ScatterMath.floorSqrt(stakedToken);

            totalWeight += performanceWeight * stakeWeight;
        }

        for (uint i = 0; i < trainers.length; i++) {
            if (trainers[i] == address(0x0)) {
                break;
            }
            uint256 performanceWeight = evaluationJobTokenContract
                .getAverageScoreForTrainerForJob(
                    requestorAddress,
                    topicName,
                    trainers[i]
                ) ** 2;

            uint256 stakedToken = trainerLockedTokenForTrainingJob[
                requestorAddress
            ][topicName][trainers[i]];

            uint256 stakeWeight = ScatterMath.floorSqrt(stakedToken);

            uint256 tokenTransferred = (performanceWeight *
                stakeWeight *
                totalRewardPool) / (totalWeight);

            if (trainers[i] == address(0)) {
                continue;
            }

            _mint(trainers[i], tokenTransferred);
        }
    }

    function rewardBenevolentValidators(
        address requestorAddress,
        string memory topicName
    ) external onlyScatterProtocolContract {
        address[] memory validators = reputationManagerContract
            .getBenevolentValidators(requestorAddress, topicName);
        uint totalStaked = 0;
        for (uint i = 0; i < validators.length; i++) {
            if (validators[i] == address(0)) {
                continue;
            }
            totalStaked += addressToStake[validators[i]];
        }
        uint totalRewardPool = (requestorLockedTokenForTrainingJob[
            requestorAddress
        ][topicName] * validatorRewardProportion) / 100;
        for (uint i = 0; i < validators.length; i++) {
            if (validators[i] == address(0)) {
                continue;
            }
            uint256 tokenTransferred = (addressToStake[validators[i]] *
                totalRewardPool) / (totalStaked + 1);

            _mint(validators[i], tokenTransferred);
        }
    }

    function rewardChallengers(
        address requestorAddress,
        string memory topicName
    ) external onlyScatterProtocolContract {
        address[] memory trainers = reputationManagerContract
            .getMalevolentTrainers(requestorAddress, topicName);
        address[] memory validators = reputationManagerContract
            .getMalevolentValidators(requestorAddress, topicName);

        address[] memory rewardedChallengers = new address[](
            trainers.length + validators.length
        );
        uint numChallengers = 0;
        for (uint256 i = 0; i < trainers.length; i++) {
            if (trainers[i] == address(0x0)) {
                break;
            }
            bool trainerChallengeSuccess = voteManagerContract
                .isChallengeSuccessfulTrainer(
                    requestorAddress,
                    topicName,
                    trainers[i]
                );
            bool validatorValidationStatus = voteManagerContract
                .getModelValidationStatus(
                    requestorAddress,
                    topicName,
                    trainers[i]
                ) == ValidationStatus.InvalidModel;

            // If the trainer was challenged successfully BUT was missed by the validator, then reward the challenger owner
            if (trainerChallengeSuccess && !validatorValidationStatus) {
                rewardedChallengers[numChallengers] = scatterProtocolContract
                    .getChallengeOwner(
                        requestorAddress,
                        topicName,
                        trainers[i]
                    );
                numChallengers += 1;
            }
        }

        for (uint256 j = 0; j < validators.length; j++) {
            if (validators[j] == address(0x0)) {
                break;
            }
            bool validatorChallengeSuccess = voteManagerContract
                .isChallengeSuccessfulValidator(
                    requestorAddress,
                    topicName,
                    validators[j]
                );

            // If the trainer was challenged successfully BUT was missed by the validator, then reward the challenger owner
            if (validatorChallengeSuccess) {
                rewardedChallengers[numChallengers] = scatterProtocolContract
                    .getChallengeOwner(
                        requestorAddress,
                        topicName,
                        validators[j]
                    );
                numChallengers += 1;
            }
        }

        for (uint256 k = 0; k < numChallengers; k++) {
            _mint(rewardedChallengers[k], lotteryPool / numChallengers);
        }
        lotteryPool = 0;
    }

    function returnTokensToTrainers(
        address requestorAddress,
        string memory topicName
    ) external onlyScatterProtocolContract {
        address[] memory trainers = reputationManagerContract
            .getBenevolentTrainers(requestorAddress, topicName);

        for (uint i = 0; i < trainers.length; i++) {
            uint tokenTransferred = trainerLockedTokenForTrainingJob[
                requestorAddress
            ][topicName][trainers[i]];

            if (trainers[i] == address(0)) {
                continue;
            }
            _mint(trainers[i], tokenTransferred);
            trainerLockedTokenForTrainingJob[requestorAddress][topicName][
                trainers[i]
            ] = 0;
        }
    }

    function returnTokensToRequestor(
        address requestorAddress,
        string memory topicName
    ) external onlyScatterProtocolContract {
        address[] memory benevolentTrainers = reputationManagerContract
            .getBenevolentTrainers(requestorAddress, topicName);
        address[] memory benevolentValidators = reputationManagerContract
            .getBenevolentValidators(requestorAddress, topicName);

        uint numBenevolentTrainers = 0;
        uint numBenevolentValidators = 0;
        for (uint i = 0; i < benevolentTrainers.length; i++) {
            if (benevolentTrainers[i] == address(0x0)) {
                break;
            }
            numBenevolentTrainers += 1;
        }

        for (uint i = 0; i < benevolentValidators.length; i++) {
            if (benevolentValidators[i] == address(0x0)) {
                break;
            }
            numBenevolentValidators += 1;
        }

        if (numBenevolentTrainers == 0) {
            _mint(
                requestorAddress,
                (trainerRewardProportion *
                    requestorLockedTokenForTrainingJob[requestorAddress][
                        topicName
                    ]) / 100
            );
        }

        if (numBenevolentValidators == 0) {
            _mint(
                requestorAddress,
                (validatorRewardProportion *
                    requestorLockedTokenForTrainingJob[requestorAddress][
                        topicName
                    ]) / 100
            );
        }
    }

    function getLotteryPool() public view returns (uint256) {
        return lotteryPool;
    }

    function getLotteryPoolExternal() external view returns (uint256) {
        return lotteryPool;
    }

    modifier onlyScatterProtocolContract() {
        require(
            msg.sender == address(scatterProtocolContract),
            "This method can only be called by the scatter protocol contract"
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
}
