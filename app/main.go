package main

import (
	"fmt"
	"github.com/codecrafters-io/kafka-starter-go/app/pkg/constants"
	"github.com/codecrafters-io/kafka-starter-go/app/pkg/response"
	"github.com/codecrafters-io/kafka-starter-go/app/pkg/server"
	// "github.com/codecrafters-io/kafka-starter-go/app/pkg/server"
	"os"
)

func main() {
	server := server.NewServer("0.0.0.0", "9092")
	server.Run()

	requestMessage, err := server.Request()
	if err != nil {
		fmt.Println("Error reading message: ", err.Error())
		os.Exit(1)
	}

	api_key := requestMessage.ApiKey
	if constants.ApiKeyNames[api_key] != "" {
		fmt.Println("API Key: ", constants.ApiKeyNames[api_key])
	}

	response_message := response.MessageBuilder(requestMessage)

	err = server.Respond(response_message)
	if err != nil {
		fmt.Println("Error sending message: ", err.Error())
		os.Exit(1)
	}

	fmt.Println("Message sent")

}
