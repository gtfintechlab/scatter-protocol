"use client"

import { ScreenContext } from "@/contexts/ScreenContext";
import { ScreensURLs } from "@/utils/types";
import { useContext, useEffect } from "react";

export default function DataPage() {
    const { setCurrentScreen } = useContext(ScreenContext);
    useEffect(() => {
        setCurrentScreen(ScreensURLs.DATA)
    }, [])

    return (
        <main>
        </main>
    )
}
