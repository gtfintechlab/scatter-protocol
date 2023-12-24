import React, { useEffect } from "react";
import { useState } from "react";
import { IoChevronDownOutline, IoChevronUpOutline } from "react-icons/io5";


function GenericDropdownMenu({
    items,
    selectedCallback,
    initialValue,
    disabled
}: { items: Record<string, string>, selectedCallback?: (itemKey: string) => void | Promise<void>, initialValue?: string, disabled?: boolean }) {
    const [isOpen, setIsOpen] = useState<boolean>(false);
    const [selectedOption, setSelectedOption] = useState<string | null>(null)

    const selectOption = async (itemKey: string, itemValue: string) => {
        setSelectedOption(itemKey);
        setIsOpen(!isOpen)

        if (selectedCallback) {
            await selectedCallback(itemKey)
        }
    }

    useEffect(() => {
        console.log(initialValue)
        if (initialValue) {
            selectOption(initialValue, items[initialValue])
            setIsOpen(false)
        }
    }, [])
    return (
        <div className="flex flex-col gap-y-1 relative">
            <div className="p-2 border-none outline-none bg-gray-100 text-sm rounded-md flex flex-row gap-x-2 w-fit items-center cursor-pointer justify-between" onClick={() => setIsOpen(!isOpen)}>
                <div className="text-gray-500">{selectedOption ?? "Select One"}</div>
                {(!isOpen || disabled) && <IoChevronDownOutline className="text-gray-500" />}
                {(isOpen && !disabled) && <IoChevronUpOutline className="text-gray-500" />}
            </div>
            {(isOpen && !disabled) && <div className="absolute left-0 top-8 mt-3 ml-1 w-fit rounded-md shadow-sm bg-white ring-1 ring-black ring-opacity-5 z-50">
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
                                className="block px-4 py-3 text-sm text-gray-700 hover:bg-gray-50 cursor-pointer flex flex-row justify-between items-center"
                                role="menuitem"
                            >
                                <div className="truncate">
                                    {itemValue}
                                </div>
                            </div>
                        )
                    })}
                </div>
            </div>}

        </div>
    );
}

export default GenericDropdownMenu;