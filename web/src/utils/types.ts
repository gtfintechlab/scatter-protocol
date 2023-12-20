import { Types } from "mongoose";
import { NextApiRequest } from "next";
import { NextRequest, NextResponse } from "next/server";

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
    isProtocolOwner: boolean;
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

export type APIWrapperResponse = Promise<NextResponse<{ success: boolean; message: string; }> | NextResponse<{ success: boolean; payload: unknown; }> | undefined>;
export type APIWrapperType = (req: NextRequest) => APIWrapperResponse;