package networking

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	utils "github.com/gtfintechlab/scatter-protocol/utils"
	"github.com/libp2p/go-libp2p/core/network"
)

func SendMessage(stream *network.Stream, message utils.Message) {
	defer (*stream).Close()
	payload, _ := json.Marshal(&message)
	(*stream).Write(payload)
}

func DecodeMessage(stream *network.Stream) utils.Message {
	buffer := StreamReader(1024, stream)
	// Read the message as a JSON
	var message map[string]interface{}
	err := json.Unmarshal(buffer, &message)
	if err != nil {
		log.Fatal(err)
	}

	return utils.Message{
		MessageType: message["MessageType"].(string),
		Payload:     message["Payload"].(map[string]interface{}),
	}
}

func ReadFileBytes(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return []byte{}
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var data []byte
	for {
		chunk := make([]byte, 1024)
		n, err := reader.Read(chunk)
		if err != nil {
			break
		}
		data = append(data, chunk[:n]...)
	}

	return data
}

func WriteBytesToFile(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %s", err)
	}
	defer file.Close()

	// Write the data to the file.
	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("error writing to file: %s", err)
	}

	return nil
}
