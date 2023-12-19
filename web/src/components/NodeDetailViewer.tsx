import { ProtocolNode } from "@/utils/types";
import ProtocolNodeForm from "./ProtocolNodeForm";

export default function NodeDetailViewer({ node }: { node: ProtocolNode | null }) {
    return (
        <div className="h-full bg-white w-full rounded-md p-4">
            {!node && <ProtocolNodeForm></ProtocolNodeForm>}
        </div>
    )
}