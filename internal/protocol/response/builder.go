package response

import (
	"github.com/codecrafters-io/kafka-starter-go/internal/protocol/request"
	"github.com/codecrafters-io/kafka-starter-go/internal/protocol/response/v4"
	"github.com/codecrafters-io/kafka-starter-go/pkg/constants/error_codes"
	"github.com/codecrafters-io/kafka-starter-go/pkg/interfaces"
)

const (
	V1 int16 = 1
	V2 int16 = 2
	V3 int16 = 3
	V4 int16 = 4
)

// MessageBuilder constructs a Response Message based on the given Request Message.
//
// It takes the ApiVersion from the request message and determines which
// version of the response message to build. If the ApiVersion is not
// supported, it returns a Response Message with the UNSUPPORTED_VERSION
// error code.
func MessageBuilder(requestMessage *request.Message) interfaces.ResponseMessage {
	switch requestMessage.ApiVersion {
	case V4:
		return v4.NewMessage(requestMessage.CorrelationId, error_codes.NONE)
	default:
		return v4.NewMessage(requestMessage.CorrelationId, error_codes.UNSUPPORTED_VERSION)
	}
}
