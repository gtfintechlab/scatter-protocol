// SPDX-License-Identifier: MIT

pragma solidity >=0.8.17;

enum roles {
    NoRole,
    Requestor,
    Trainer,
    Validator,
    Challenger
}

enum Checkpoints {
    ModelSubmission,
    FederatedJobEnd
}

enum FederatedJobStatus {
    NotStarted,
    InProgress,
    DistributingRewards,
    Complete
}

enum ValidationStatus {
    InvalidModel,
    ValidModel,
    ModelNotValidated
}
