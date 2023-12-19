"use client";
import React, { useEffect, useState } from "react";
import { ScreensURLs } from "@/utils/types";
import { ScreenContext } from "@/contexts/ScreenContext";

export default function ContextProvider({
    children
}: {
    children: React.ReactNode
}) {
    const [isMobile, setIsMobile] = useState<boolean>(false);
    const [currentScreen, setCurrentScreen] = useState<ScreensURLs>(ScreensURLs.HOME)

    useEffect(() => {
        const handleResize = () => {
            const mobileWidthThreshold = 768;
            setIsMobile(window.innerWidth < mobileWidthThreshold);

        };

        handleResize();
        window.addEventListener('resize', handleResize);

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
            {children}
        </ScreenContext.Provider>)
}