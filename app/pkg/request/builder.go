package request

import (
	"encoding/binary"

	"github.com/codecrafters-io/kafka-starter-go/app/pkg/constants"
	"github.com/codecrafters-io/kafka-starter-go/app/pkg/constants/error_codes"
)

func MessageBuilder(data []byte) (*Message, int16) {
	apiVersion := extractApiVersion(data)

	if apiVersion < constants.MIN_SUPPORTED_API_VERSION || apiVersion > constants.MAX_SUPPORTED_API_VERSION {
		return nil, error_codes.UNSUPPORTED_VERSION
	}

	message, error := DeserializeMessage(data)
	if error != nil {
		return nil, error_codes.INVALID_REQUEST
	}

	return message, error_codes.NONE
}

func extractApiVersion(data []byte) int16 {
	return int16(binary.BigEndian.Uint16(data[6:8]))
}
