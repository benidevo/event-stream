package response

import (
	"github.com/codecrafters-io/kafka-starter-go/app/pkg/constants/error_codes"
	"github.com/codecrafters-io/kafka-starter-go/app/pkg/interfaces"
	"github.com/codecrafters-io/kafka-starter-go/app/pkg/request"
	"github.com/codecrafters-io/kafka-starter-go/app/pkg/response/v4"
)

const (
	V1 int16 = 1
	V2 int16 = 2
	V3 int16 = 3
	V4 int16 = 4
)

func MessageBuilder(requestMessage *request.Message) interfaces.ResponseMessage {
	switch requestMessage.ApiVersion {
	case V4:
		return v4.NewMessage(requestMessage.Size, requestMessage.CorrelationId, error_codes.NONE)
	default:
		return v4.NewMessage(requestMessage.Size, requestMessage.CorrelationId, error_codes.UNSUPPORTED_VERSION)
	}
}
