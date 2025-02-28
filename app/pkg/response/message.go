package response

import (
	"bytes"
	"encoding/binary"
)

type Message struct {
	Size          int32
	CorrelationId int32
}

// NewMessage creates a new Message with the given size and correlationId.
func NewMessage(size, correlationId int32) *Message {
	return &Message{size, correlationId}
}

// Serialize the message into a byte array.
//
// The first four bytes contain the Size of the message, and the second four
// bytes contain the CorrelationId of the message.
//
// The method returns an error if either of the binary.Write calls fail.
func (m *Message) Serialize() ([]byte, error) {
	var buffer bytes.Buffer
	var err error

	err = binary.Write(&buffer, binary.BigEndian, m.Size)
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.BigEndian, m.CorrelationId)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
