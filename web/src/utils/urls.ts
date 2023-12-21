function getBaseURL() {
    return "http://localhost:3000";
}

export const urls = {
    baseUrl: getBaseURL(),
    api: {
        nodes: {
            general: "/api/nodes"
        },
        workspaces: {
            general: "/api/workspaces"
        },
        web3: {
            accounts: "/api/web3/accounts"
        },
        steps: {
            general: "/api/steps"
        }
    }
};