package networking

import (
	"bufio"
	"bytes"
	"io"

	"github.com/libp2p/go-libp2p/core/network"
)

func StreamReader(bufferSize int, stream *network.Stream) []byte {
	var buffer []byte
	reader := bufio.NewReader(*stream)

	for {
		tempBuffer := make([]byte, bufferSize)
		_, err := reader.Read(tempBuffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return []byte{}
		}

		tempBuffer = bytes.Trim(tempBuffer, "\x00") // Remove extra zeroes from buffer
		// Append to global buffer
		buffer = append(buffer, tempBuffer...)
	}
	return buffer
}
