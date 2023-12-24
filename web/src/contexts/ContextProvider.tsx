"use client";
import React, { useEffect, useState } from "react";
import { ScreensURLs, Step, Workspace } from "@/utils/types";
import { ScreenContext } from "@/contexts/ScreenContext";
import { ProtocolContext, StepContext } from "./ProtocolContext";

export default function ContextProvider({
    children
}: {
    children: React.ReactNode
}) {
    const [isMobile, setIsMobile] = useState<boolean>(false);
    const [currentScreen, setCurrentScreen] = useState<ScreensURLs>(ScreensURLs.HOME)
    const [currentWorkspace, setWorkspace] = useState<Workspace>({ name: undefined, _id: undefined });
    const [stepUpdateKey, setUpdateKey] = useState<number>(0);
    const [editMode, setEditMode] = useState<boolean>(false);
    const [stepInEdit, setStepInEdit] = useState<Step | null>(null);
    useEffect(() => {
        const handleResize = () => {
            const mobileWidthThreshold = 768;
            setIsMobile(window.innerWidth < mobileWidthThreshold);

        };

        handleResize();
        window.addEventListener('resize', handleResize);

        const currentWorkspace = localStorage.getItem("currentWorkspace");
        if (currentWorkspace) {
            setWorkspace(JSON.parse(currentWorkspace))
        }

        return () => {
            window.removeEventListener('resize', handleResize);
        };

    }, [])

    return (
        <ScreenContext.Provider value={{
            isMobile, setIsMobile: (mobile: boolean) => {
                setIsMobile(mobile)
            },
            currentScreen, setCurrentScreen: (screen: ScreensURLs) => {
                setCurrentScreen(screen);
            }
        }}>
            <ProtocolContext.Provider value={{
                currentWorkspace,
                setCurrentWorkspace: (workspace: Workspace) => { setWorkspace({ ...workspace }) }
            }}>
                <StepContext.Provider value={{
                    stepUpdateKey, setStepUpdateKey: (key: number) => setUpdateKey(key),
                    editMode, setEditMode: (mode: boolean) => { setEditMode(mode) },
                    stepInEdit, setStepInEdit: (step: null | Step) => {
                        if (step) {
                            setStepInEdit({ ...step })
                        } else {
                            setStepInEdit(null);
                        }
                    }
                }}>
                    {children}
                </StepContext.Provider>

            </ProtocolContext.Provider>
        </ScreenContext.Provider>)
}