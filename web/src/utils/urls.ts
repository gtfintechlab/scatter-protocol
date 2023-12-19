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
        }
    }
};