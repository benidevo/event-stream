package v4

import (
	"bytes"
	"encoding/binary"
)

type ApiVersion struct {
	ApiKey     int16
	MinVersion int16
	MaxVersion int16
}

type Message struct {
	CorrelationId int32
	ErrorCode     int16
	ApiVersions   []ApiVersion
	ThrottleTime  int32
}

// NewMessage creates a new Message with the given correlationId and errorCode.
// The ApiVersions list is fixed to a single entry: ApiKey = 18, MinVersion = 0, MaxVersion = 4.
// The ThrottleTime is set to 0.
func NewMessage(correlationId int32, errorCode int16) *Message {
	apiVersions := []ApiVersion{
		{
			ApiKey:     18,
			MinVersion: 0,
			MaxVersion: 4,
		},
	}

	return &Message{
		CorrelationId: correlationId,
		ErrorCode:     errorCode,
		ApiVersions:   apiVersions,
		ThrottleTime:  0,
	}
}

// Serialize encodes the Message into a byte slice following a specific binary format.
// The serialization includes the following fields in order:
// 1. CorrelationId: Written in BigEndian format.
// 2. ErrorCode: Written in BigEndian format.
// 3. ApiVersions: Serialized as a compact array, including:
//   - The length of the ApiVersions array, encoded as a variable-length unsigned integer.
//   - Each ApiVersion entry, containing:
//   - ApiKey, MinVersion, MaxVersion: Each written in BigEndian format.
//   - A compact tag buffer (currently empty), encoded as a variable-length unsigned integer.
//
// 4. ThrottleTime: Written in BigEndian format.
// 5. A final compact tag buffer for the entire response (currently empty).
// The method returns the serialized byte slice or an error if any serialization step fails.
func (m *Message) Serialize() ([]byte, error) {
	var contentBuffer bytes.Buffer

	if err := binary.Write(&contentBuffer, binary.BigEndian, m.CorrelationId); err != nil {
		return nil, err
	}

	if err := binary.Write(&contentBuffer, binary.BigEndian, m.ErrorCode); err != nil {
		return nil, err
	}

	if err := writeUnsignedVarInt(&contentBuffer, uint32(len(m.ApiVersions)+1)); err != nil {
		return nil, err
	}

	for _, version := range m.ApiVersions {
		if err := binary.Write(&contentBuffer, binary.BigEndian, version.ApiKey); err != nil {
			return nil, err
		}
		if err := binary.Write(&contentBuffer, binary.BigEndian, version.MinVersion); err != nil {
			return nil, err
		}
		if err := binary.Write(&contentBuffer, binary.BigEndian, version.MaxVersion); err != nil {
			return nil, err
		}

		if err := writeUnsignedVarInt(&contentBuffer, 0); err != nil {
			return nil, err
		}
	}

	if err := binary.Write(&contentBuffer, binary.BigEndian, m.ThrottleTime); err != nil {
		return nil, err
	}

	if err := writeUnsignedVarInt(&contentBuffer, 0); err != nil {
		return nil, err
	}

	contentSize := int32(contentBuffer.Len())
	finalBuffer := bytes.NewBuffer(make([]byte, 0, 4+contentBuffer.Len()))

	if err := binary.Write(finalBuffer, binary.BigEndian, contentSize); err != nil {
		return nil, err
	}

	if _, err := finalBuffer.Write(contentBuffer.Bytes()); err != nil {
		return nil, err
	}

	return finalBuffer.Bytes(), nil
}

// writeUnsignedVarInt writes a variable length unsigned int to a buffer.
//
// The varint encoding is as follows:
//   - The most significant bit of each byte indicates if there is a next byte
//     (if set) or not (if unset).
//   - The lower 7 bits of each byte contain a part of the value.
//   - The values are concatenated in big-endian order.
//   - The value is terminated by a byte with the most significant bit unset.
func writeUnsignedVarInt(buffer *bytes.Buffer, value uint32) error {
	for {
		b := byte(value & 0x7f)
		value >>= 7
		if value != 0 {
			b |= 0x80
		}
		err := buffer.WriteByte(b)
		if err != nil {
			return err
		}
		if value == 0 {
			break
		}
	}
	return nil
}
