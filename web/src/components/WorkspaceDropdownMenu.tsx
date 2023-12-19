import { getWorkspaces, userCreateWorkspace, userDeleteWorkspace } from '@/actions/Workspace';
import { ProtocolContext } from '@/contexts/ProtocolContext';
import { Workspace } from '@/utils/types';
import React, { useContext, useEffect, useState } from 'react';
import TrashIcon from './icons/TrashIcon';

const WorkspaceDropdownMenu = () => {
    const [isOpen, setIsOpen] = useState(false);
    const [workspaces, setWorkspaces] = useState<Workspace[]>([]);
    const { currentWorkspace, setCurrentWorkspace } = useContext(ProtocolContext);
    const [inCreation, setInCreation] = useState<boolean>(false);
    const [currentWorkspaceName, setCurrentWorkspaceName] = useState<string>("");
    const [creatingWorkspace, setCreatingWorkspace] = useState<boolean>(false);

    useEffect(() => {
        const getAllWorkspaces = async () => {
            const workspaces = await getWorkspaces()
            setWorkspaces([...workspaces])
        }
        getAllWorkspaces().then().catch()

    }, [])

    const toggleDropdown = () => {
        setIsOpen(!isOpen);
        setInCreation(false);
        setCurrentWorkspaceName("");
    };

    const createWorkspace = async () => {
        setCreatingWorkspace(true);
        const workspace = await userCreateWorkspace(currentWorkspaceName);
        setCurrentWorkspace(workspace);
        setCreatingWorkspace(false);
        setCurrentWorkspaceName("");
        setInCreation(false);
        setWorkspaces([...workspaces, workspace])
    }

    const setActiveWorkspace = (workspace: Workspace) => {
        setCurrentWorkspace(workspace)
        toggleDropdown();
    }

    const deleteWorkspaceFromList = async (workspace: Workspace) => {
        await userDeleteWorkspace(workspace._id?.toString() as string)
        toggleDropdown()

        const newWorkspaces = await getWorkspaces();
        setWorkspaces([...newWorkspaces])
        setActiveWorkspace({ name: undefined, _id: undefined });
    }

    return (
        <div className="relative inline-block text-left">
            <button
                onClick={toggleDropdown}
                type="button"
                className="inline-flex justify-center w-full px-4 py-2 text-sm font-medium text-white bg-gray-800 rounded-md focus:outline-none"
                id="options-menu"
                aria-haspopup="true"
                aria-expanded="true"
            >
                {currentWorkspace.name ? currentWorkspace.name : "Select a Workspace"}
            </button>

            {isOpen && (
                <div className="origin-top-right absolute right-0 mt-2 w-64 rounded-md shadow-lg bg-white ring-1 ring-black ring-opacity-5">
                    <div
                        className=""
                        role="menu"
                        aria-orientation="vertical"
                        aria-labelledby="options-menu"
                    >
                        {workspaces.map((workspace: Workspace, index: number) => {
                            return (
                                <div
                                    onClick={() => setActiveWorkspace(workspace)}
                                    key={index}
                                    className="block px-4 py-3 text-sm text-gray-700 hover:bg-gray-50 cursor-pointer flex flex-row justify-between items-center"
                                    role="menuitem"
                                >
                                    <div >
                                        {workspace.name}
                                    </div>
                                    <TrashIcon className='w-6 h-6 text-red-500 z-50 p-1 hover:bg-gray-200 rounded-md' onClick={() => { deleteWorkspaceFromList(workspace) }}></TrashIcon>
                                </div>
                            )
                        })}
                        {!inCreation && <div
                            className="block px-4 py-3 text-sm text-gray-700 hover:bg-gray-100 cursor-pointer"
                            role="menuitem"
                            onClick={() => setInCreation(true)}
                        >
                            Create a workspace
                        </div>}
                        {inCreation &&
                            <div className='w-full flex flex-row gap-2 p-2'>
                                <input
                                    className="block px-2 text-sm w-[75%] text-gray-700 focus:outline-none"
                                    role="menuitem"
                                    placeholder='Workspace Name'
                                    onChange={(event) => { setCurrentWorkspaceName(event?.target.value) }}
                                />
                                <button className={`text-white bg-black rounded-md px-1 hover:text-black hover:bg-white border-2 border-black ${creatingWorkspace ? 'opacity-50' : ''}`} disabled={creatingWorkspace} onClick={createWorkspace}>Create</button>
                            </div>}

                    </div>
                </div>
            )}
        </div>
    );
};

export default WorkspaceDropdownMenu;
