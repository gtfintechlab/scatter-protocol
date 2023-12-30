package peers

import (
	"context"
	"encoding/json"
	"strings"

	networking "github.com/gtfintechlab/scatter-protocol/core/networking"
	peerDatabase "github.com/gtfintechlab/scatter-protocol/core/peers/db"
	"github.com/gtfintechlab/scatter-protocol/core/protocol"
	utils "github.com/gtfintechlab/scatter-protocol/core/utils"
	network "github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
)

func requestDecryptionKeyHandler(node *utils.PeerNode, stream *network.Stream, message utils.Message) {
	requestorId := (*stream).Conn().RemotePeer().String()
	blockchainAddress := protocol.GetBlockchainAddressByNodeId(node, requestorId)

	var payload utils.DecryptionKeyRequest
	jsonData, _ := json.Marshal(message.Payload)
	json.Unmarshal(jsonData, &payload)

	// If it is the validator or the requestor, send the decryption key
	if protocol.IsValidatorForRequestorAndTopicExternal(node, payload.RequestorAddress, payload.TopicName, blockchainAddress) ||
		utils.RemoveHexPrefix(payload.RequestorAddress) == utils.RemoveHexPrefix(blockchainAddress) {
		private := *node.AESKey
		peerId, _ := peer.Decode(requestorId)
		newStream, _ := (*node.PeerToPeerServer).NewStream(context.Background(), peerId, utils.PROTOCOL_IDENTIFIER)

		networking.SendMessage(
			&newStream,
			utils.Message{
				MessageType: utils.AWKNOWLEDGE_DECRYPTION_KEY_REQUEST,
				Payload: map[string]interface{}{
					"trainerAddress":   *node.BlockchainAddress,
					"requestorAddress": payload.RequestorAddress,
					"privateKey":       private,
					"topicName":        payload.TopicName,
				},
			})
	}

}

func awknowledgeDecryptionKeyRequestHandler(node *utils.PeerNode, stream *network.Stream, message utils.Message) {
	var payload utils.AwknowledgeDecryptionKeyRequest
	jsonData, _ := json.Marshal(message.Payload)
	json.Unmarshal(jsonData, &payload)
	peerDatabase.WriteDecryptionKey(node, payload.TrainerAddress, payload.PrivateKey, payload.TopicName,
		strings.ToLower(payload.RequestorAddress))
	(*node.AESChannels)[strings.ToLower(payload.RequestorAddress)][payload.TopicName][strings.ToLower(payload.TrainerAddress)] <- true
	(*stream).Close()
}
