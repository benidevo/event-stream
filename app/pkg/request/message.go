package request

import (
	"bytes"
	"encoding/binary"
)

type Message struct {
	Size          int32
	ApiKey        int16
	ApiVersion    int16
	CorrelationId int32
}

// NewMessage creates a new Message with the given size, apiKey, apiVersion, and
// correlationId.
func NewMessage(apiKey, apiVersion int16, size, correlationId int32) *Message {
	return &Message{size, apiKey, apiVersion, correlationId}
}

// DeserializeMessage reads a Message from the given byte array.
//
// It reads the apiKey, apiVersion, size, and correlationId from the byte array
// and returns a Message containing the values.
//
// The method returns an error if deserialization fails.
func DeserializeMessage(data []byte) (*Message, error) {
	var message Message

	err := binary.Read(bytes.NewReader(data), binary.BigEndian, &message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}
