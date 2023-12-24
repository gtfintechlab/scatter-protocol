"use client"
import { Step, Workspace } from "@/utils/types";
import { createContext } from "react";

export const ProtocolContext = createContext({
    currentWorkspace: { name: undefined, _id: undefined } as Workspace, setCurrentWorkspace: (workpace: Workspace) => { }
});

interface StepContextType {
    stepUpdateKey: number;
    setStepUpdateKey: (key: number) => void,
    editMode: boolean,
    setEditMode: (mode: boolean) => void,
    stepInEdit: null | Step,
    setStepInEdit: (step: Step | null) => void,

}
export const StepContext = createContext<StepContextType>({
    stepUpdateKey: 0, setStepUpdateKey: (key: number) => { },
    editMode: false, setEditMode: (mode: boolean) => { },
    stepInEdit: null, setStepInEdit: (step: null | Step) => { },
});