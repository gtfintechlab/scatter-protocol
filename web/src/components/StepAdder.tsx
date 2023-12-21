import GenericDropdownMenu from "@/components/GenericDropdownMenu";
import { STEPS_DICTIONARY } from "@/utils/constants";
import { StepTypes } from "@/utils/types";
import { useState } from "react";
import { RequestorAddTopic } from "./steps/RequestorAddTopic";

export function StepAdder({ completeAdditionCallback }: { completeAdditionCallback: () => void }) {
    const [selectedStep, setSelectedStep] = useState<null | StepTypes | string>(null);

    return (
        <div className="h-full">
            <div className="flex flex-col gap-y-2 h-full overflow-y-scroll">
                <label className="text-black text-sm font-semibold">Step Name:</label>
                <GenericDropdownMenu items={STEPS_DICTIONARY} selectedCallback={setSelectedStep}></GenericDropdownMenu>
                {selectedStep === StepTypes.REQUESTOR_ADD_TOPIC && <RequestorAddTopic completionCallback={completeAdditionCallback} />}
            </div>
        </div>)
}