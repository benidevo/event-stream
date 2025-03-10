package main

import (
	"github.com/codecrafters-io/kafka-starter-go/internal/server"
)

func main() {
	server := server.NewServer("0.0.0.0", "9092")
	server.Run()
}
