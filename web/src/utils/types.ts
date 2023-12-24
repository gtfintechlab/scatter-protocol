import { Types } from "mongoose";
import { NextApiRequest } from "next";
import { NextRequest, NextResponse } from "next/server";

export type Optional<T, K extends keyof T> = Pick<Partial<T>, K> & Omit<T, K>;

export enum ScreensURLs {
    HOME = "",
    ENVIRONMENT = "environment",
    NODES = "nodes",
    DATA = "data"
}

export interface NavbarItem {
    text: string;
    url: string;
}

export interface EthereumAccount {
    _id: Types.ObjectId
    privateKey: string;
    address: string;
    mnemonic: string;
    path: string;
    index: number;
}

export enum ProtocolNodeState {
    STARTED = "started",
    STOPPED = "stopped"
}

export interface ProtocolNode {
    _id: Types.ObjectId
    peerType: PeerType;
    apiPort: number;
    postgresUsername: string;
    postgresPassword: string;
    databasePort: number;
    blockchainAddress: string;
    privateKey: string;
    dummyLoad: boolean;
    useMdns: boolean;
    tokenSupply: number;
    initialTokenSupply: number;
    isProtocolOwner: boolean;
    state: ProtocolNodeState;
    workspaceId: Types.ObjectId;
}

export enum PeerType {
    REQUESTOR = "requestor",
    TRAINER = "trainer",
    VALIDATOR = "validator",
    CHALLENGER = "challenger",
    NO_ROLE = "no role"
}

export interface Workspace {
    _id?: Types.ObjectId;
    name?: string;
}


export enum HttpMethod {
    GET = "GET",
    POST = "POST",
    PATCH = "PATCH",
    PUT = "PUT",
    DELETE = "DELETE",
}

export enum StepTypes {
    DEPLOY_PROTOCOL = "Deploy Protocol",
    PROTOCOL_OWNER_TRANSFER_TOKEN = "Protocol Owner - Transfer Initial Supply",
    INITIALIZE_ROLES = "Initialize Roles",
    REQUESTOR_ADD_TOPIC = "Requestor - Add Topic",
    TRAINER_ADD_TOPIC = "Trainer - Add Topic",
    REQUESTOR_START_TRAINING = "Requestor - Start Training"
}

export interface Step {
    _id: Types.ObjectId;
    type: StepTypes;
    apiPath: string;
    apiMethod: HttpMethod;
    body: Record<string, string | number | Record<string, string | number>>;
    workspaceId: Types.ObjectId | string;
    nodeId: Types.ObjectId | string;
    order: number;
}

export interface InternalRequest extends NextApiRequest {
    body: { [key: string]: unknown };
}

export interface InternalRequestData {
    url: string;
    method: HttpMethod;
    body?: { [key: string]: unknown };
    queryParams?: { [key: string]: string | number | boolean | undefined };
}

export interface InternalResponseData<T> {
    success: boolean;
    message?: string;
    payload?: T;

}

export interface LogEvent {
    _id?: Types.ObjectId;
    logType: string;
    workspaceId: Types.ObjectId;
    blockchainAddress: string;
    xDataPoint: number;
    yDataPoint: number;
    createdAt: Date;
}

export enum Nodes {
    ALL = "all"
}

export type APIWrapperResponse = Promise<NextResponse<{ success: boolean; message: string; }> | NextResponse<{ success: boolean; payload: unknown; }> | undefined>;
export type APIWrapperType = (req: NextRequest) => APIWrapperResponse;