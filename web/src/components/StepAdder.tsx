import GenericDropdownMenu from "@/components/GenericDropdownMenu";
import { STEPS_DICTIONARY } from "@/utils/constants";
import { StepTypes } from "@/utils/types";
import { useState } from "react";
import { RequestorAddTopic } from "./steps/RequestorAddTopic";
import DeployProtocol from "./steps/DeployProtocol";
import ProtocolOwnerTransferInitialSupply from "./steps/ProtocolOwnerTransferInitialSupply";
import { TrainerAddTopic } from "./steps/TrainerAddTopic";
import RequestorStartTraining from "./steps/RequestorStartTraining";

export function StepAdder({ completeAdditionCallback }: { completeAdditionCallback: () => void }) {
    const [selectedStep, setSelectedStep] = useState<null | StepTypes | string>(null);

    return (
        <div className="h-full">
            <div className="flex flex-col gap-y-2 h-full overflow-y-scroll">
                <label className="text-black text-sm font-semibold">Step Name:</label>
                <GenericDropdownMenu items={STEPS_DICTIONARY} selectedCallback={setSelectedStep}></GenericDropdownMenu>
                {selectedStep === StepTypes.REQUESTOR_ADD_TOPIC && <RequestorAddTopic completionCallback={completeAdditionCallback} />}
                {selectedStep === StepTypes.DEPLOY_PROTOCOL && <DeployProtocol completionCallback={completeAdditionCallback} />}
                {selectedStep === StepTypes.PROTOCOL_OWNER_TRANSFER_TOKEN && <ProtocolOwnerTransferInitialSupply completionCallback={completeAdditionCallback} />}
                {selectedStep === StepTypes.TRAINER_ADD_TOPIC && <TrainerAddTopic completionCallback={completeAdditionCallback} />}
                {selectedStep === StepTypes.REQUESTOR_START_TRAINING && <RequestorStartTraining completionCallback={completeAdditionCallback} />}

            </div>
        </div>)
}