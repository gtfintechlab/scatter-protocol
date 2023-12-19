"use client"

import { ScreenContext } from "@/contexts/ScreenContext";
import { ScreensURLs } from "@/utils/types";
import { useContext, useEffect } from "react";

export default function EnvironmentPage() {
    const { setCurrentScreen } = useContext(ScreenContext);
    useEffect(() => {
        setCurrentScreen(ScreensURLs.ENVIRONMENT)
    }, [])
    return (
        <main>
        </main>
    )
}
