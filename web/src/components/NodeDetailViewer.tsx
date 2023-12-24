import { ProtocolNode } from "@/utils/types";
import ProtocolNodeForm from "@/components/ProtocolNodeForm";
import { useState } from "react";

export default function NodeDetailViewer({ node, nodesUpdateCallback }: { node: ProtocolNode | null, nodesUpdateCallback?: () => void }) {

    const [formClearKey, setFormClearKey] = useState<number>(0)

    const clearForm = () => {
        setFormClearKey(formClearKey + 1)
        if (nodesUpdateCallback) {
            nodesUpdateCallback();
        }
    }
    return (
        <div className="h-full w-full bg-white rounded-md p-4 min-w-0 shrink overflow-y-scroll">
            <ProtocolNodeForm node={node} key={formClearKey} nodeCreationCallback={clearForm}></ProtocolNodeForm>
        </div>
    )
}