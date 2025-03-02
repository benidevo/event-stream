package v4

import (
	"bytes"
	"encoding/binary"
)

type Message struct {
	Size          int32
	CorrelationId int32
	ErrorCode     int16
}

// NewMessage creates a new Message with the specified size, correlationId, and errorCode.
// It returns a pointer to the newly created Message instance.
func NewMessage(size, correlationId int32, errorCode int16) *Message {
	return &Message{size, correlationId, errorCode}
}

// Serialize encodes the Message into a byte array in BigEndian format.
// The byte array contains the Size, CorrelationId, and ErrorCode fields
// in sequence. Returns an error if the binary.Write operation fails.
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

	err = binary.Write(&buffer, binary.BigEndian, m.ErrorCode)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
