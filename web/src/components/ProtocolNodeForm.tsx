import React, { useState } from 'react';
import { PeerType } from '@/utils/types';
import { capitalizeWords } from '@/utils/format';
import { ROLE_TO_COLOR_MAPPING } from '@/utils/constants';
const ProtocolNodeForm: React.FC = () => {
    const [peerType, setPeerType] = useState<string | PeerType>("")

    return (
        <div>
            <div className='flex flex-row gap-4 w-full items-center'>
                <label className='text-black font-semibold'>Node Type:</label>
                {Object.values(PeerType).map((type: PeerType) => {
                    return (<button
                        onClick={() => { setPeerType(type) }}
                        className={`text-white p-1 bg-${ROLE_TO_COLOR_MAPPING[type]}-500 rounded-md ${peerType == type ? "border-2 border-black" : ""}`}>{capitalizeWords(type)}</button>)
                })}
            </div>
        </div>
    )
};

export default ProtocolNodeForm;
