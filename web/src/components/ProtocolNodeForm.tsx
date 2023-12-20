import React, { useContext, useState } from 'react';
import { PeerType, ProtocolNode } from '@/utils/types';
import { capitalizeWords } from '@/utils/format';
import { ROLE_TO_COLOR_MAPPING } from '@/utils/constants';
import { ProtocolContext } from '@/contexts/ProtocolContext';
const ProtocolNodeForm = ({ node }: { node: ProtocolNode | null }) => {
    const [peerType, setPeerType] = useState<string | PeerType>("")
    const { currentWorkspace } = useContext(ProtocolContext);
    const [error, setError] = useState<string>("")

    const [apiPort, setApiPort] = useState<string>("")
    const [postgresUsername, setPostgresUsername] = useState<string>("")
    const [postgresPassword, setPostgresPassword] = useState<string>("")
    const [databasePort, setDatabasePort] = useState<string>("")
    const [initialTokenSupply, setInitialTokenSupply] = useState<number>(0)
    const [blockchainAddress, setBlockchainAddress] = useState<string>("")
    const [privateKey, setPrivateKey] = useState<string>("")
    const [useDummyData, setUseDummyData] = useState<boolean>(true)
    const [useMdns, setUseMdns] = useState<boolean>(true)
    const [isProtocolOwner, setIsProtocolOwner] = useState<boolean>(true)

    const createNode = async () => {
        setError("");
        if (!currentWorkspace) {
            setError("Must select a workspace before creating a node")
            return;
        }
    }

    const updateNode = async () => {
        setError("");
        if (!currentWorkspace) {
            setError("Must select a workspace before updating a node")
            return;
        }
    }

    return (
        <div className='h-full flex flex-col justify-between overflow-y-scroll'>
            <div className='flex flex-col gap-x-2 gap-y-3 w-full'>
                <div className='flex flex-row gap-4 w-full items-center break-all flex-wrap'>
                    <label className='text-black font-semibold'>Node Type:</label>
                    {Object.values(PeerType).map((type: PeerType, index: number) => {
                        return (<button
                            onClick={() => { setPeerType(type) }}
                            key={index}
                            className={`text-white p-1 bg-${ROLE_TO_COLOR_MAPPING[type]}-500 rounded-md shrink grow ${peerType == type ? "border-2 border-black" : ""}`}>{capitalizeWords(type)}</button>)
                    })}
                </div>
                <div className='flex flex-row gap-4 w-full items-center flex-1 flex-wrap'>
                    <div className='flex flex-col gap-2 shrink grow'>
                        <label className="text-black font-semibold">API Port:</label>
                        <input className='text-black p-2 border-none outline-none bg-gray-100 rounded-md' placeholder='API Port'></input>
                    </div>
                    <div className='flex flex-col gap-2 shrink grow'>
                        <label className="text-black font-semibold">Postgres Username:</label>
                        <input className='text-black p-2 border-none outline-none bg-gray-100 rounded-md' placeholder='Postgres Username'></input>
                    </div>
                    <div className='flex flex-col gap-2 shrink grow'>
                        <label className="text-black font-semibold">Postgres Password:</label>
                        <input className='text-black p-2 border-none outline-none bg-gray-100 rounded-md' placeholder='Postgres Password'></input>
                    </div>
                </div>
                <div className='flex flex-row gap-4 w-full items-center flex-1 flex-wrap'>
                    <div className='flex flex-col gap-2 shrink grow'>
                        <label className="text-black font-semibold">Database Port:</label>
                        <input className='text-black p-2 border-none outline-none bg-gray-100 rounded-md' placeholder='Database Port'></input>
                    </div>
                    <div className='flex flex-col gap-2 shrink grow'>
                        <label className="text-black font-semibold">Initial Token Supply:</label>
                        <input className='text-black p-2 border-none outline-none bg-gray-100 rounded-md' placeholder='Initial Token Supply'></input>
                    </div>
                </div>

                <div className='flex flex-row gap-4 w-full items-center flex-1 flex-wrap'>
                    <div className='flex flex-col gap-2 shrink grow'>
                        <label className="text-black font-semibold">Blockchain Address:</label>
                        <input className='text-black p-2 border-none outline-none bg-gray-100 rounded-md' placeholder='Blockchain Address'></input>
                    </div>
                    <div className='flex flex-col gap-2 shrink grow'>
                        <label className="text-black font-semibold">Private Key:</label>
                        <input className='text-black p-2 border-none outline-none bg-gray-100 rounded-md' placeholder='Private Key'></input>
                    </div>
                </div>
                <div className='flex flex-row gap-4 w-full items-center flex-1 flex-wrap'>
                    <div className='flex flex-row gap-2 shrink grow'>
                        <label className="text-black font-semibold">Use Dummy Data?</label>
                        <input type="checkbox"></input>
                    </div>
                    <div className='flex flex-row gap-2 shrink grow'>
                        <label className="text-black font-semibold">Use mDNS?</label>
                        <input type="checkbox"></input>
                    </div>
                    <div className='flex flex-row gap-2 shrink grow'>
                        <label className="text-black font-semibold">Is Protocol Owner?</label>
                        <input type="checkbox"></input>
                    </div>
                </div>
            </div>
            <div className='flex flex-col gap-2'>
                <button
                    className="p-2 text-white w-full border-2 rounded-md border-black bg-black text-center hover:bg-white hover:text-black cursor-pointer mt-4"
                    onClick={async () => { if (node) { await createNode() } else { await updateNode() } }}
                >
                    {!node ? "Create Node" : "Update Node"}</button>
                {error && <p className='text-red-500 font-semibold text-center'>{error}</p>}
            </div>
        </div>
    )
};

export default ProtocolNodeForm;
