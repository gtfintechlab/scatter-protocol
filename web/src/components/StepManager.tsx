import { useContext, useEffect, useState } from "react"
import { StepAdder } from "./StepAdder";
import { getStepsByWorkspace } from "@/actions/Step";
import { ProtocolContext } from "@/contexts/ProtocolContext";
import { Step } from "@/utils/types";

export default function StepManager({ className }: { className?: string }) {
    const [stepAddMode, setStepAddMode] = useState<boolean>(false);
    const [steps, setSteps] = useState<Step[]>([]);
    const { currentWorkspace } = useContext(ProtocolContext);

    useEffect((
    ) => {
        const stepSetter = async () => {
            const workspaceSteps = await getStepsByWorkspace(currentWorkspace._id?.toString() as string)
            setSteps([...workspaceSteps])
        }
        stepSetter().then().catch();
    }, [currentWorkspace])
    return (
        <main className={`h-full bg-white w-full rounded-md p-4 ${className} flex flex-col gap-2 h-full`}>
            <div className="flex items-center">
                <button className="justify-center p-1 bg-black text-sm rounded-md hover:bg-white hover:text-black border-2 border-black" type="button" onClick={() => setStepAddMode(!stepAddMode)}>
                    <span className="text-xs">{stepAddMode ? "Back" : "Add New Step"}</span>
                </button>
            </div>
            {!stepAddMode && <div className="relative w-full overflow-auto h-full overflow-y-scroll pb-12 overflow-x-scroll">
                <table className="w-full caption-bottom text-sm border-collapse">
                    <thead className="[&amp;_tr]:border-b">
                        <tr className="border-b transition-colors hover:bg-muted/50 data-[state=selected]:bg-muted">
                            <th className="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-500">Step Number</th>
                            <th className="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-500">Address</th>
                            <th className="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-500">Node Role</th>
                            <th className="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-500">Step Name</th>
                            <th className="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-500">View Step Details</th>
                            <th className="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-500">Delete Step</th>
                        </tr>
                    </thead>
                    <tbody>
                        {steps.map((step: Step, index: number) => {
                            return (
                                <tr className="border-b transition-colors hover:bg-muted/50 data-[state=selected]:bg-muted" key={index}>
                                    <td className="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-300">{step.order}</td>
                                    <td className="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-300">{""}</td>
                                    <td className="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-300">{""}</td>
                                    <td className="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-300">{""}</td>
                                    <td className="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-300">{""}</td>
                                    <td className="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-300">{""}</td>
                                </tr>
                            )
                        })}
                    </tbody>
                </table>
            </div>}
            {stepAddMode && <StepAdder completeAdditionCallback={() => setStepAddMode(false)}></StepAdder>}
        </main>
    )
}