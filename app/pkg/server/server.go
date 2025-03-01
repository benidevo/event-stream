package server

import (
	"net"
	"strings"

	"github.com/codecrafters-io/kafka-starter-go/app/pkg/constants"
	"github.com/codecrafters-io/kafka-starter-go/app/pkg/request"
	"github.com/codecrafters-io/kafka-starter-go/app/pkg/response"
)

type Server struct {
	Host       string
	Port       string
	Connection net.Conn
}

// NewServer returns a new Server given a host and a port.
func NewServer(host, port string) *Server {
	return &Server{
		Host: host,
		Port: port,
	}

}

// Run the server and accept a connection.
func (s *Server) Run() {
	var address strings.Builder
	address.WriteString(s.Host)
	address.WriteString(":")
	address.WriteString(s.Port)

	listener, err := net.Listen(constants.PROTOCOL, address.String())
	if err != nil {
		panic(err)
	}

	connection, err := listener.Accept()
	if err != nil {
		panic(err)
	}

	s.Connection = connection
}

// Request reads a message from the server's connection and deserializes it.
//
// It reads data from the connection until it reads a full message, and then
// deserializes the message from the data.
//
// If reading or deserialization fails, it returns an error.
func (s *Server) Request() (*request.Message, error) {
	var data []byte
	var err error

	data = make([]byte, constants.DEFAULT_BUFFER_SIZE)
	_, err = s.Connection.Read(data)
	if err != nil {
		return nil, err
	}

	message, err := request.DeserializeMessage(data)

	return message, err
}

// Respond sends a serialized message over the server's connection.
//
// It serializes the provided message and writes the serialized bytes to the connection.
// If serialization or writing fails, it returns an error.
func (s *Server) Respond(message *response.Message) error {

	serializedMessage, err := message.Serialize()
	if err != nil {
		return err
	}

	_, err = s.Connection.Write(serializedMessage)
	if err != nil {
		return err
	}

	return nil

}

// Shutdown closes the connection of the server.
func (s *Server) Shutdown() {
	s.Connection.Close()
}
