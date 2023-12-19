import APIWrapper from "server/utils/APIWrapper";
import { APIWrapperType } from "@/utils/types"

const route: APIWrapperType = APIWrapper({
    GET: {
        config: {
        },
        handler: async () => {
            return {
                Hello: "World",
                Version: 2.0,
            };
        },
    },
});

export let GET: APIWrapperType, POST: APIWrapperType, PATCH: APIWrapperType, DELETE: APIWrapperType, PUT: APIWrapperType;
GET = POST = PATCH = POST = DELETE = PUT = route;