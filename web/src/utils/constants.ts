import { PeerType, ScreensURLs } from "@/utils/types";

// Home page for steps the simulations go through
// Environment to configure blockchain env
// Nodes page to configure node states
// Data page to see all data/graphs
export const NAVBAR_ITEMS = [
    { text: "Home", url: ScreensURLs.HOME }, { text: "Environment", url: ScreensURLs.ENVIRONMENT }, { text: 'Nodes', url: ScreensURLs.NODES }, { text: "Data", url: ScreensURLs.DATA }
]

export const ROLE_TO_COLOR_MAPPING = {
    [PeerType.REQUESTOR]: "red",
    [PeerType.TRAINER]: "blue",
    [PeerType.VALIDATOR]: "green",
    [PeerType.CHALLENGER]: "pink",
    [PeerType.NO_ROLE]: "gray",

}