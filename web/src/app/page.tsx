"use client"

import NodeManager from "@/components/NodeManager";
import { ScreenContext } from "@/contexts/ScreenContext";
import { ProtocolNode, ScreensURLs } from "@/utils/types";
import { useContext, useEffect, useState } from "react";
import StepManager from "@/components/StepManager";

export default function Home() {
  const { setCurrentScreen, isMobile } = useContext(ScreenContext);
  useEffect(() => {
    setCurrentScreen(ScreensURLs.HOME)
  }, [])
  const selectedNodeCallback = (node: ProtocolNode | null) => {
  }


  return (
    <main className={`flex flex-1 flex-col gap-4 p-4 md:gap-8 md:p-10 bg-gray-100 h-full ${isMobile ? 'overflow-y-scroll' : ''}`}>
      <div className="grid grid-cols-1 md:grid-cols-[360px,1fr] gap-5 items-start h-full">
        <NodeManager nodeCallback={selectedNodeCallback}></NodeManager>
        <StepManager></StepManager>
        {isMobile && <div className="bg-gray-100 text-gray-100">{"Hello World"}</div>}
      </div>
    </main>
  )
}
