import React, { useContext, useEffect, useRef, useState } from 'react';
import { EthereumAccount, PeerType, ProtocolNode, ProtocolNodeState } from '@/utils/types';
import { capitalizeWords } from '@/utils/format';
import { DEFAULT_NODE_OPTIONS, ROLE_TO_COLOR_MAPPING } from '@/utils/constants';
import { ProtocolContext } from '@/contexts/ProtocolContext';
import { createProtocolNode, getNodesForWorkspace, updateProtocolNode } from '@/actions/ProtocolNode';
import { getAllWeb3Accounts } from '@/actions/Web3';
import { Types } from 'mongoose';

const ProtocolNodeForm = ({ node, nodeCreationCallback }: { node: ProtocolNode | null, nodeCreationCallback?: () => Promise<void> | void }) => {
    const [peerType, setPeerType] = useState<string | PeerType>("")
    const { currentWorkspace } = useContext(ProtocolContext);
    const [error, setError] = useState<string>("")

    const [apiPort, setApiPort] = useState<number>(0)
    const [postgresUsername, setPostgresUsername] = useState<string>("")
    const [postgresPassword, setPostgresPassword] = useState<string>("")
    const [databasePort, setDatabasePort] = useState<number | undefined>(undefined)
    const [initialTokenSupply, setInitialTokenSupply] = useState<number>(0)
    const [blockchainAddress, setBlockchainAddress] = useState<string>("")
    const [privateKey, setPrivateKey] = useState<string>("")
    const [useDummyData, setUseDummyData] = useState<boolean>(true)
    const [useMdns, setUseMdns] = useState<boolean>(true)
    const [isProtocolOwner, setIsProtocolOwner] = useState<boolean>(true)
    const [isMalicious, setIsMalicious] = useState<boolean>(false)
    const [settingValues, setSettingValues] = useState<boolean>(false);

    const setDefaultValues = async () => {
        const nodes = await getNodesForWorkspace(currentWorkspace._id?.toString() as string);
        const web3Accounts = await getAllWeb3Accounts();
        const web3AccountPrivateKeys = new Set(web3Accounts.map((account: EthereumAccount) => { return account.privateKey }))
        const occupiedPrivateKeys = new Set(nodes.map((node: ProtocolNode) => { return node.privateKey }))
        const occupiedApiPorts = new Set(nodes.map((node: ProtocolNode) => { return node.apiPort }))

        if (occupiedApiPorts.size == 0) {
            setApiPort(DEFAULT_NODE_OPTIONS.apiPort)
        } else {
            setApiPort(Math.max(...occupiedApiPorts) + 1)
        }

        if (occupiedPrivateKeys.size == 0) {
            setPrivateKey(web3Accounts[0]?.privateKey ?? "")
            setBlockchainAddress(web3Accounts[0]?.address ?? "")
        } else {
            const freePrivateKeys = [...web3AccountPrivateKeys].filter(element => !occupiedPrivateKeys.has(element));
            if (freePrivateKeys.length) {
                const account = web3Accounts.find((acc: EthereumAccount) => { return acc.privateKey == freePrivateKeys[0] })
                setPrivateKey(account?.privateKey ?? "")
                setBlockchainAddress(account?.address ?? "")
            }
        }

        setInitialTokenSupply(DEFAULT_NODE_OPTIONS.intialTokenSupply)
        setPostgresUsername(DEFAULT_NODE_OPTIONS.postgresUsername);
        setPostgresPassword(DEFAULT_NODE_OPTIONS.postgresPassword);
        setUseDummyData(DEFAULT_NODE_OPTIONS.useDummyData);
        setUseMdns(DEFAULT_NODE_OPTIONS.useMdns);
        setPeerType("")
        setDatabasePort(undefined);
        if (nodes.length === 0) {
            setIsProtocolOwner(true);
        } else {
            setIsProtocolOwner(DEFAULT_NODE_OPTIONS.isProtocolOwner);
        }
    }

    const setPeerTypeAndDbPort = (peerType: PeerType) => {
        setPeerType(peerType)
        setDatabasePort(DEFAULT_NODE_OPTIONS.databasePorts[peerType])
    }

    const createNode = async () => {
        setError("");
        if (!currentWorkspace) {
            setError("Must select a workspace before creating a node")
            return;
        }
        if (databasePort === undefined || !peerType) {
            console.log(databasePort, peerType)
            setError("Must complete all fields before creating node")
            return;
        }


        const newNode: Omit<ProtocolNode, "_id"> = {
            apiPort,
            peerType: peerType as PeerType,
            postgresUsername,
            postgresPassword,
            databasePort: databasePort as number,
            initialTokenSupply,
            tokenSupply: 0,
            blockchainAddress,
            privateKey,
            dummyLoad: useDummyData,
            useMdns,
            state: ProtocolNodeState.STOPPED,
            isProtocolOwner,
            isMalicious,
            workspaceId: currentWorkspace._id as Types.ObjectId
        }

        await createProtocolNode(newNode);

        if (nodeCreationCallback) {
            await nodeCreationCallback()
        }

        await setDefaultValues()
    }

    const updateNode = async () => {
        setError("");
        if (!currentWorkspace) {
            setError("Must select a workspace before updating a node")
            return;
        }

        const newNode: ProtocolNode = {
            _id: (node as ProtocolNode)._id,
            apiPort,
            peerType: peerType as PeerType,
            postgresUsername,
            postgresPassword,
            databasePort: databasePort as number,
            initialTokenSupply,
            tokenSupply: 0,
            blockchainAddress,
            privateKey,
            dummyLoad: useDummyData,
            useMdns,
            state: ProtocolNodeState.STOPPED,
            isProtocolOwner,
            isMalicious,
            workspaceId: currentWorkspace._id as Types.ObjectId
        }

        await updateProtocolNode(newNode);
        if (nodeCreationCallback) {
            await nodeCreationCallback()
        }

        setPreexistingValues();

    }

    const setPreexistingValues = () => {
        const { apiPort, peerType, postgresPassword, postgresUsername, databasePort, initialTokenSupply, blockchainAddress, privateKey, dummyLoad, useMdns,
            isProtocolOwner, isMalicious
        } = node as ProtocolNode

        setApiPort(apiPort);
        setPeerType(peerType)
        setPostgresPassword(postgresPassword)
        setPostgresUsername(postgresUsername);
        setDatabasePort(databasePort);
        setInitialTokenSupply(initialTokenSupply);
        setBlockchainAddress(blockchainAddress);
        setPrivateKey(privateKey);
        setUseDummyData(dummyLoad)
        setUseMdns(useMdns);
        setIsMalicious(isMalicious);
        setIsProtocolOwner(isProtocolOwner);
    }

    useEffect(() => {
        if (!currentWorkspace._id) {
            return;
        }
        setSettingValues(true);
        if (!node) {
            setDefaultValues().then(() => setSettingValues(false)).catch()
        } else {
            setPreexistingValues()
            setSettingValues(false);
        }
    }, [currentWorkspace, node])

    return (
        <div className='h-full flex flex-col justify-between overflow-y-scroll'>
            <div className='flex flex-col gap-x-2 gap-y-3 w-full'>
                <div className='flex flex-row gap-4 w-full items-center break-all flex-wrap'>
                    <label className='text-black font-semibold'>Node Type:</label>
                    {Object.values(PeerType).map((type: PeerType, index: number) => {
                        return (<button
                            onClick={() => { setPeerTypeAndDbPort(type) }}
                            key={index}
                            className={`text-white p-1 bg-${ROLE_TO_COLOR_MAPPING[type]}-500 rounded-md shrink grow ${peerType == type ? "border-2 border-black" : ""} ${settingValues ? "opacity-50" : ""}`} disabled={settingValues}>{capitalizeWords(type)}</button>)
                    })}
                </div>
                <div className='flex flex-row gap-4 w-full items-center flex-1 flex-wrap'>
                    <div className='flex flex-col gap-2 shrink grow'>
                        <label className="text-black font-semibold">API Port:</label>
                        <input className='text-black p-2 border-none outline-none bg-gray-100 rounded-md' placeholder='API Port'
                            value={apiPort} onChange={(event) => setApiPort(parseInt(event?.target.value))} type="number" disabled={settingValues}></input>
                    </div>
                    <div className='flex flex-col gap-2 shrink grow'>
                        <label className="text-black font-semibold">Postgres Username:</label>
                        <input className='text-black p-2 border-none outline-none bg-gray-100 rounded-md' placeholder='Postgres Username'
                            value={postgresUsername} onChange={(event) => setPostgresUsername(event?.target.value)} disabled={settingValues}></input>
                    </div>
                    <div className='flex flex-col gap-2 shrink grow'>
                        <label className="text-black font-semibold">Postgres Password:</label>
                        <input className='text-black p-2 border-none outline-none bg-gray-100 rounded-md' placeholder='Postgres Password'
                            value={postgresPassword} onChange={(event) => setPostgresPassword(event?.target.value)} disabled={settingValues}></input>
                    </div>
                </div>
                <div className='flex flex-row gap-4 w-full items-center flex-1 flex-wrap'>
                    <div className='flex flex-col gap-2 shrink grow'>
                        <label className="text-black font-semibold">Database Port:</label>
                        <input className='text-black p-2 border-none outline-none bg-gray-100 rounded-md' placeholder='Database Port'
                            value={databasePort !== undefined ? databasePort : ''} onChange={(event) => setDatabasePort(parseInt(event?.target.value))} type='number' disabled={settingValues}></input>
                    </div>
                    <div className='flex flex-col gap-2 shrink grow'>
                        <label className="text-black font-semibold">Initial Token Supply:</label>
                        <input className='text-black p-2 border-none outline-none bg-gray-100 rounded-md' placeholder='Initial Token Supply'
                            value={initialTokenSupply} onChange={(event) => setInitialTokenSupply(parseInt(event?.target.value))} type='number' disabled={settingValues}></input>
                    </div>
                </div>

                <div className='flex flex-row gap-4 w-full items-center flex-1 flex-wrap'>
                    <div className='flex flex-col gap-2 shrink grow'>
                        <label className="text-black font-semibold">Blockchain Address:</label>
                        <input className='text-black p-2 border-none outline-none bg-gray-100 rounded-md' placeholder='Blockchain Address'
                            value={blockchainAddress} onChange={(event) => setBlockchainAddress(event.target.value)} disabled={settingValues}></input>
                    </div>
                    <div className='flex flex-col gap-2 shrink grow'>
                        <label className="text-black font-semibold">Private Key:</label>
                        <input className='text-black p-2 border-none outline-none bg-gray-100 rounded-md' placeholder='Private Key'
                            value={privateKey} onChange={(event) => setPrivateKey(event.target.value)} disabled={settingValues}></input>
                    </div>
                </div>
                <div className='flex flex-row gap-4 w-full items-center flex-1 flex-wrap'>
                    <div className='flex flex-row gap-2 shrink grow'>
                        <label className="text-black font-semibold">Use Dummy Data?</label>
                        <input type="checkbox" checked={useDummyData} onChange={() => setUseDummyData(!useDummyData)} disabled={settingValues}></input>
                    </div>
                    <div className='flex flex-row gap-2 shrink grow'>
                        <label className="text-black font-semibold">Use mDNS?</label>
                        <input type="checkbox" checked={useMdns} onChange={() => setUseMdns(!useMdns)} disabled={settingValues}></input>
                    </div>
                    <div className='flex flex-row gap-2 shrink grow'>
                        <label className="text-black font-semibold">Is Protocol Owner?</label>
                        <input type="checkbox" checked={isProtocolOwner} onChange={() => setIsProtocolOwner(!isProtocolOwner)} disabled={settingValues}></input>
                    </div>
                    <div className='flex flex-row gap-2 shrink grow'>
                        <label className="text-black font-semibold">Is Malicious?</label>
                        <input type="checkbox" checked={isMalicious} onChange={() => setIsMalicious(!isMalicious)} disabled={settingValues}></input>
                    </div>

                </div>
            </div>
            <div className='flex flex-col gap-2'>
                <button
                    className={`${settingValues ? "opacity-50" : ""} p-2 text-white w-full border-2 rounded-md border-black bg-black text-center hover:bg-white hover:text-black cursor-pointer mt-4`}
                    disabled={settingValues}
                    onClick={async () => { if (!node) { await createNode() } else { await updateNode() } }}
                >
                    {!node ? "Create Node" : "Update Node"}</button>
                {error && <p className='text-red-500 font-semibold text-center'>{error}</p>}
            </div>
        </div>
    )
};

export default ProtocolNodeForm;
