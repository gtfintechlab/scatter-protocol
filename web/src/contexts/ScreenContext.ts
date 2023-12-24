"use client"
import { ScreensURLs } from "@/utils/types";
import { createContext } from "react";

export const ScreenContext = createContext({
    isMobile: false, setIsMobile: (isMobile: boolean) => { },
    currentScreen: ScreensURLs.HOME, setCurrentScreen: (screen: ScreensURLs) => { }
});