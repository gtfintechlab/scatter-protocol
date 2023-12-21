import { ROLE_TO_COLOR_MAPPING } from "@/utils/constants";
import { capitalizeWords } from "@/utils/format";
import { ProtocolNode } from "@/utils/types";
import React from "react";
import { useState } from "react";
import { IoChevronDownOutline, IoChevronUpOutline } from "react-icons/io5";


function NodeDropdown({
    items,
    selectedCallback
}: { items: Record<string, ProtocolNode>, selectedCallback?: (itemKey: string) => void | Promise<void> }) {
    const [isOpen, setIsOpen] = useState<boolean>(false);
    const [selectedOption, setSelectedOption] = useState<string | null>(null)
    const [selectedDisplay, setSelectedDisplay] = useState<ProtocolNode | null>(null);

    const selectOption = async (itemKey: string, itemValue: ProtocolNode) => {
        setSelectedOption(itemKey);
        setSelectedDisplay({ ...itemValue });
        setIsOpen(!isOpen)

        if (selectedCallback) {
            await selectedCallback(itemValue._id.toString())
        }
    }
    return (
        <div className="relative">
            <div className="p-2 border-none outline-none bg-gray-100 text-sm rounded-md flex flex-row gap-x-2 w-64 items-center cursor-pointer justify-between" onClick={() => setIsOpen(!isOpen)}>
                {!selectedDisplay && <div className="text-gray-500">Select One</div>}
                {selectedDisplay && <div className="text-gray-500 flex flex-row w-full">
                    <div className="truncate w-[50%]">{selectedDisplay.blockchainAddress}</div>
                    <div className={`truncate w-[50%] text-center text-white p-y-1 rounded-md bg-${ROLE_TO_COLOR_MAPPING[selectedDisplay.peerType]}-500`}>{capitalizeWords(selectedDisplay.peerType)}</div>
                </div>}

                {!isOpen && <IoChevronDownOutline className="text-gray-500" />}
                {isOpen && <IoChevronUpOutline className="text-gray-500" />}
            </div>
            {isOpen && <div className="absolute left-0 top-8 mt-3 ml-1 w-64 rounded-md shadow-sm bg-white ring-1 ring-black ring-opacity-5 absolute z-50">
                <div
                    className="max-h-48 overflow-y-scroll"
                    role="menu"
                    aria-orientation="vertical"
                    aria-labelledby="options-menu"
                >

                    {Object.entries(items).map(([itemKey, itemValue], index: number) => {
                        return (
                            <div
                                onClick={async () => await selectOption(itemKey, itemValue)}
                                key={index}
                                className="block px-4 py-3 text-sm text-gray-700 hover:bg-gray-50 cursor-pointer flex flex-row justify-between items-center gap-x-2"
                                role="menuitem"
                            >
                                <div className="truncate w-[50%]">
                                    {itemValue.blockchainAddress}
                                </div>
                                <div className={`truncate w-[50%] text-center text-white p-y-1 rounded-md bg-${ROLE_TO_COLOR_MAPPING[itemValue.peerType]}-500`}>
                                    {capitalizeWords(itemValue.peerType)}
                                </div>

                            </div>
                        )
                    })}
                </div>
            </div>}

        </div>
    );
}

export default NodeDropdown;