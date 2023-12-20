import { ProtocolNode } from "@/utils/types";
import ProtocolNodeForm from "./ProtocolNodeForm";

export default function NodeDetailViewer({ node }: { node: ProtocolNode | null }) {
    return (
        <div className="h-full w-full bg-white rounded-md p-4 min-w-0 shrink overflow-y-scroll">
            <ProtocolNodeForm node={node}></ProtocolNodeForm>
        </div>
    )
}