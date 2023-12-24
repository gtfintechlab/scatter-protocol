export const urls = {
    baseUrl: "http://localhost:3000",
    externalUrl: "http://localhost:2000",
    external: {
        startNode: "/start-node",
        stopNode: "/stop-node"
    },
    api: {
        nodes: {
            general: "/api/nodes",
            single: "/api/nodes/single",
            reset: "/api/nodes/reset",
            addressWorkspace: "/api/nodes/address-workspace",

        },
        workspaces: {
            general: "/api/workspaces"
        },
        web3: {
            accounts: "/api/web3/accounts"
        },
        steps: {
            general: "/api/steps",
            byNode: "/api/steps/by-node"
        },
        events: {
            logs: "/api/events/logs"
        }
    }
};