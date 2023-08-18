package protocol

import (
	"fmt"
	"strings"

	"github.com/gtfintechlab/scatter-protocol/networking"
	"github.com/gtfintechlab/scatter-protocol/utils"
)

func downloadEvaluationJob(requestorAddress string, ipfsCid string) {
	filePath := fmt.Sprintf("training/validator/jobs/%s/%s/", strings.ToLower(requestorAddress), strings.ToLower(ipfsCid))
	fileBytes, _ := utils.GetFileBytesFromIPFS(ipfsCid, filePath)
	networking.UnzipFolder(fileBytes, filePath)
}

func buildEvaluationImage(requestorId string, ipfsCid string) {
}
