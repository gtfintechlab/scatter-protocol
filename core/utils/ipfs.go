package utils

import (
	"bytes"
	"io"
	"log"
	"os"

	ipfs "github.com/ipfs/go-ipfs-api"
	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")

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

func UploadEncryptedFileToIpfs(filePath string, node *PeerNode) string {
	shell := ipfs.NewShell(os.Getenv("IPFS_NODE_URL"))
	file, err := os.Open(filePath)
	// Encrypted file logic here

	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	// Read the content of the file
	fileData, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading file data:", err)
	}

	// Encrypt the file content
	encryptedData, err := encryptData(fileData, *node.AESKey)
	if err != nil {
		log.Fatal("Error encrypting data:", err)
	}

	// Create a temporary file to hold the encrypted data
	encryptedFile, err := os.CreateTemp("", "encrypted_file_*.bin")
	if err != nil {
		log.Fatal("Error creating temporary file:", err)
	}
	defer os.Remove(encryptedFile.Name())
	defer encryptedFile.Close()

	_, err = encryptedFile.Write(encryptedData)
	if err != nil {
		log.Fatal("Error writing encrypted data to file:", err)
	}
	// Create a reader from the encrypted data
	encryptedDataReader := bytes.NewReader(encryptedData)

	// Upload the encrypted file to IPFS
	cid, err := shell.Add(encryptedDataReader)
	if err != nil {
		log.Fatal("Failed to upload encrypted file with error:", err)
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
