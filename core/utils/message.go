package utils

import (
	"bytes"
	"encoding/gob"
)

func (message *Message) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	enc := gob.NewEncoder(b)

	if err := enc.Encode(message.MessageType); err != nil {
		return nil, err
	}

	if err := enc.Encode(message.Payload); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (message *Message) UnmarshalBinary(data []byte) error {
	//Use default gob decoder
	reader := bytes.NewReader(data)
	dec := gob.NewDecoder(reader)
	if err := dec.Decode(message); err != nil {
		return err
	}

	return nil
}
