import { HttpMethod, PeerType, ScreensURLs, StepTypes } from "@/utils/types";
import { enumToDictionary, getRandomNumber } from "./format";

// Home page for steps the simulations go through
// Environment to configure blockchain env
// Nodes page to configure node states
// Data page to see all data/graphs
export const NAVBAR_ITEMS = [
    { text: "Home", url: ScreensURLs.HOME }, { text: 'Nodes', url: ScreensURLs.NODES }, { text: "Data", url: ScreensURLs.DATA }
]

export const ROLE_TO_COLOR_MAPPING = {
    [PeerType.REQUESTOR]: "red",
    [PeerType.TRAINER]: "blue",
    [PeerType.VALIDATOR]: "green",
    [PeerType.CHALLENGER]: "pink",
    [PeerType.NO_ROLE]: "gray",

}

export const WEB3_NUM_ACCOUNTS = 100;
export const WEB3_MNEMONIC = "test test test test test test test test test test test junk"
export const WEB3_PATH = "m/44'/60'/0'/0"

export const DEFAULT_NODE_OPTIONS = {
    apiPort: 5002,
    intialTokenSupply: 100000,
    useMdns: true,
    useDummyData: true,
    isProtocolOwner: false,
    postgresUsername: "postgres",
    postgresPassword: "postgres",
    databasePorts: {
        [PeerType.REQUESTOR]: 8701,
        [PeerType.TRAINER]: 8702,
        [PeerType.VALIDATOR]: 8703,
        [PeerType.NO_ROLE]: 0,
        [PeerType.CHALLENGER]: 0
    }
}

export const DEFAULT_STEP_OPTIONS: { [key: string]: any } = {
    topicName: "cats and dogs",
    trainingJobPath: "training/requestor/train",
    reward: (): number => {
        return getRandomNumber(1000, 2000)
    },
    validationThreshold: 40,
    evaluationJobPath: "training/requestor/validate",
    evaluationJobDataPath: "training/data/requestor/validate",
    trainingDataPath: "training/data/trainer/example",
    stake: (): number => {
        return getRandomNumber(500, 1500)
    },
}
export const STEPS_DICTIONARY = enumToDictionary(StepTypes);

export const STEPS_CONFIG = {
    [StepTypes.DEPLOY_PROTOCOL]: {
        apiPath: "/deploy-protocol",
        apiMethod: HttpMethod.POST
    },
    [StepTypes.PROTOCOL_OWNER_TRANSFER_TOKEN]: {
        apiPath: "/transfer-initial-supply",
        apiMethod: HttpMethod.POST
    },
    [StepTypes.INITIALIZE_ROLES]: {
        apiPath: "/initialize-roles",
        apiMethod: HttpMethod.POST
    },
    [StepTypes.REQUESTOR_ADD_TOPIC]: {
        apiPath: "/requestor/add-topic",
        apiMethod: HttpMethod.POST
    },
    [StepTypes.TRAINER_ADD_TOPIC]: {
        apiPath: "/trainer/add-topic",
        apiMethod: HttpMethod.POST
    },
    [StepTypes.REQUESTOR_START_TRAINING]: {
        apiPath: "/requestor/start-training",
        apiMethod: HttpMethod.POST
    },
}
