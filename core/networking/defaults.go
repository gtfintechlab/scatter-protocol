package networking

import utils "github.com/gtfintechlab/scatter-protocol/core/utils"

var MESSAGE_JOIN_NETWORK utils.Message = utils.Message{
	MessageType: utils.MESSAGE_CODE_JOIN_NETWORK,
	Payload: map[string]interface{}{
		"Hello": "World",
	},
}
