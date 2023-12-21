"use client"
import { Workspace } from "@/utils/types";
import { createContext } from "react";

export const ProtocolContext = createContext({
    currentWorkspace: { name: undefined, _id: undefined } as Workspace, setCurrentWorkspace: (workpace: Workspace) => { }
});

export const StepContext = createContext({ stepUpdateKey: 0, setStepUpdateKey: (key: number) => { } });