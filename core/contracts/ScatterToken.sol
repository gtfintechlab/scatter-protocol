// SPDX-License-Identifier: MIT

pragma solidity >=0.8.17;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Capped.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "hardhat/console.sol";

import "./Shared.sol";
import "./Math.sol";

import "./IScatterProtocol.sol";
import "./IScatterToken.sol";
import "./IReputationManager.sol";
import "./IEvaluationJobToken.sol";

contract ScatterToken is ERC20Capped, ERC20Burnable, IScatterToken {
    uint256 public requiredModelValidatorStake;
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
    uint256 lotteryPercentage = 10;

    // Punishment constants
    uint256 validatorPunishmentPercentage = 10;

    // Reward Factors
    uint256 trainerRewardProportion = 80;
    uint256 validatorRewardProportion = 20;

    IScatterProtocol scatterProtocolContract;
    IReputationManager reputationManagerContract;
    IEvaluationJobToken evaluationJobTokenContract;

    event TokensStaked(address from, uint256 amount);
    event TokensUnstaked(address to, uint256 amount);

    constructor(
        uint256 cap
    ) ERC20("ScatterToken", "ST") ERC20Capped(cap * (10 ** decimals())) {
        owner = payable(msg.sender);
        _mint(owner, cap * (10 ** decimals()));

        requiredModelValidatorStake = 10000;
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

    function stakeToken(uint256 amount) public {
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

    function removeStake(uint256 amount) public {
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

    function donateToLottery(
        address requestorAddress,
        string memory topicName
    ) external onlyScatterProtocolContract {
        uint256 requestorPooledReward = requestorLockedTokenForTrainingJob[
            requestorAddress
        ][topicName];

        uint256 toLottery = (requestorPooledReward * lotteryPercentage) / 100;
        lotteryPool += toLottery;
        requestorLockedTokenForTrainingJob[requestorAddress][
            topicName
        ] -= toLottery;
    }

    function punishRogueTrainers(
        address requestorAddress,
        string memory topicName
    ) external onlyScatterProtocolContract {
        address[] memory trainers = reputationManagerContract
            .getMalevolentTrainers(requestorAddress, topicName);
        for (uint i = 0; i < trainers.length; i++) {
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
            uint256 punishmentAmount = (addressToStake[validators[i]] *
                validatorPunishmentPercentage) / 100;
            lotteryPool += punishmentAmount;

            addressToStake[validators[i]] -= punishmentAmount;
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
            uint performanceWeight = evaluationJobTokenContract
                .getAverageScoreForTrainerForJob(
                    requestorAddress,
                    topicName,
                    trainers[i]
                ) ** 2;

            uint stakedToken = trainerLockedTokenForTrainingJob[
                requestorAddress
            ][topicName][trainers[i]];

            uint stakeWeight = Math.floorSqrt(stakedToken);

            totalWeight += performanceWeight * stakeWeight;
        }

        for (uint i = 0; i < trainers.length; i++) {
            uint256 performanceWeight = evaluationJobTokenContract
                .getAverageScoreForTrainerForJob(
                    requestorAddress,
                    topicName,
                    trainers[i]
                ) ** 2;

            uint256 stakedToken = trainerLockedTokenForTrainingJob[
                requestorAddress
            ][topicName][trainers[i]];

            uint256 stakeWeight = Math.floorSqrt(stakedToken);

            uint256 tokenTransferred = (performanceWeight *
                stakeWeight *
                totalRewardPool) / totalWeight;
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
            totalStaked += addressToStake[validators[i]];
        }
        uint totalRewardPool = (requestorLockedTokenForTrainingJob[
            requestorAddress
        ][topicName] * validatorRewardProportion) / 100;

        for (uint i = 0; i < validators.length; i++) {
            uint256 tokenTransferred = (addressToStake[validators[i]] *
                totalRewardPool) / totalStaked;
            _mint(validators[i], tokenTransferred);
        }
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

            _mint(trainers[i], tokenTransferred);
        }
    }

    function getLotteryPool() public view returns (uint256) {
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
