package response

import (
	"github.com/codecrafters-io/kafka-starter-go/app/pkg/constants/error_codes"
	"github.com/codecrafters-io/kafka-starter-go/app/pkg/request"
	"github.com/codecrafters-io/kafka-starter-go/app/pkg/response/v4"
)

func MessageBuilder(requestMessage *request.Message) *v4.Message {
	switch requestMessage.ApiVersion {
	case 4:
		return v4.NewMessage(requestMessage.Size, requestMessage.CorrelationId, error_codes.NONE)
	default:
		return v4.NewMessage(requestMessage.Size, requestMessage.CorrelationId, error_codes.NONE)
	}
}
