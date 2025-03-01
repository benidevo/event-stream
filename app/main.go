package main

import (
	"fmt"
	"github.com/codecrafters-io/kafka-starter-go/app/pkg/response"
	"github.com/codecrafters-io/kafka-starter-go/app/pkg/server"
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

	messageSize := 4

	err = server.Respond(response.NewMessage(int32(messageSize), requestMessage.CorrelationId))
	if err != nil {
		fmt.Println("Error sending message: ", err.Error())
		os.Exit(1)
	}

	fmt.Println("Message sent")

}
