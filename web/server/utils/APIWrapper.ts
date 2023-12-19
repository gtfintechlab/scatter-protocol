import mongoose from "mongoose";
import { NextRequest, NextResponse } from "next/server";
import { HttpMethod } from "@/utils/types";
import { APIWrapperType } from "@/utils/types"
import * as context from "next/headers";

interface RouteConfig {
    handleResponse?: boolean;
}
interface Route<T> {
    config: RouteConfig;
    handler: (
        req: NextRequest,
    ) => Promise<T>;
}


function APIWrapper(
    routeHandlers: Partial<Record<HttpMethod, Route<unknown>>>
): APIWrapperType {
    return async (req: NextRequest) => {
        const method = req.method;
        const route = routeHandlers[method as HttpMethod];

        if (!method || !route) {
            const errorMessage = method
                ? `Request method ${method} is invalid.`
                : "Missing request method.";

            return NextResponse.json({
                success: false,
                message: errorMessage,
            }, { status: 200 });
        }

        const { handler, config } = route;

        try {
            const data = await handler(req);

            if (config?.handleResponse) {
                return;
            }

            return NextResponse.json({ success: true, payload: data }, { status: 200 });
        } catch (e) {
            if (e instanceof mongoose.Error) {
                return NextResponse.json({
                    success: false,
                    message: "An Internal Server error occurred.",
                }, { status: 500 });
            }

            const error = e as Error;
            return NextResponse.json({ success: false, message: error.message }, { status: 400 });
        }
    };
}

export default APIWrapper;
