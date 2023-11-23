package utils

import (
	"io"
	"log"
	"os"

	ipfs "github.com/ipfs/go-ipfs-api"
	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")

// func RunIpfsNode(p2pPort int, gatewayPort int, rpcApiPort int) {
// 	exec.Command(
// 		"docker", "run",
// 		"-d", "--name", "ipfs_host",
// 		"-v", fmt.Sprintf("%s:/export", os.Getenv("IPFS_STAGING_PATH")),
// 		"-v", fmt.Sprintf("%s:/data/ipfs", os.Getenv("IPFS_DATA_PATH")),
// 		"-p", fmt.Sprintf("%d:4001", p2pPort),
// 		"-p", fmt.Sprintf("%d:4001/udp", p2pPort),
// 		"-p", fmt.Sprintf("127.0.0.1:%d:8080", gatewayPort),
// 		"-p", fmt.Sprintf("127.0.0.1:%d:5001", rpcApiPort),
// 		"ipfs/kubo:latest",
// 	).Run()

// 	exec.Command(
// 		"docker", "exec", "ipfs_host",
// 		"ipfs", "config",
// 		"--json", "API.HTTPHeaders.Access-Control-Allow-Origin",
// 		"'[\"http://localhost:8080\", \"http://localhost:3000\", \"http://127.0.0.1:5001\", \"https://webui.ipfs.io\"]'",
// 	).Run()

// 	exec.Command(
// 		"docker", "exec", "ipfs_host",
// 		"ipfs", "config",
// 		"--json", "API.HTTPHeaders.Access-Control-Allow-Methods",
// 		"'[\"PUT\", \"POST\"]'",
// 	).Run()

// 	exec.Command(
// 		"docker", "exec", "ipfs_host",
// 		"ipfs", "daemon",
// 	).Run()

// 	exec.Command(
// 		"docker", "exec",
// 		"ipfs_host", "ipfs",
// 		"swarm", "peers",
// 	).Run()
// }

func UploadFileToIpfs(filePath string) string {
	shell := ipfs.NewShell(os.Getenv("IPFS_NODE_URL"))
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	cid, err := shell.Add(file)
	if err != nil {
		log.Fatal("Failed to upload file with error:", err)
	}

	shell.Pin(cid)

	return cid
}

func GetFileBytesFromIPFS(hash string) ([]byte, error) {
	shell := ipfs.NewShell(os.Getenv("IPFS_NODE_URL"))
	file, err := shell.Cat(hash)
	if err != nil {
		return []byte{}, err
	}
	defer file.Close()
	bytes, _ := io.ReadAll(file)

	return bytes, nil
}
