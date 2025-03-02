package main

import (
	"fmt"
	"github.com/codecrafters-io/kafka-starter-go/internal/protocol/response"
	"github.com/codecrafters-io/kafka-starter-go/internal/server"
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

	response_message := response.MessageBuilder(requestMessage)

	err = server.Respond(response_message)
	if err != nil {
		fmt.Println("Error sending message: ", err.Error())
		os.Exit(1)
	}

	fmt.Println("Message sent")
}
