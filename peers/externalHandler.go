package peers

import (
	"net/http"

	utils "github.com/gtfintechlab/scatter-protocol/utils"
)

func switchPeerNodeRole(node *utils.PeerNode) http.HandlerFunc {
	return func(http.ResponseWriter, *http.Request) {
		if node.PeerType == utils.PEER_REQUESTOR {
			node.PeerType = utils.PEER_TRAINER
		} else {
			node.PeerType = utils.PEER_REQUESTOR
		}
	}
}
