package server

import (
	"net"
	"strings"

	"github.com/codecrafters-io/kafka-starter-go/internal/protocol/request"
	"github.com/codecrafters-io/kafka-starter-go/pkg/constants"
	"github.com/codecrafters-io/kafka-starter-go/pkg/interfaces"
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

// Request reads data from the server's connection and constructs a request.Message.
//
// It attempts to read a byte array from the connection using a predefined buffer size.
// The byte array is then deserialized into a request.Message using the MessageBuilder function.
// Returns the constructed request.Message or an error if reading from the connection or deserialization fails.
func (s *Server) Request() (*request.Message, error) {
	var data []byte
	var err error

	data = make([]byte, constants.DEFAULT_BUFFER_SIZE)
	_, err = s.Connection.Read(data)
	if err != nil {
		return nil, err
	}
	request_message, err := request.MessageBuilder(data)

	return request_message, nil
}

// Respond sends a serialized response message over the server's connection.
//
// It takes a ResponseMessage interface, serializes it into a byte array,
// and writes the byte array to the server's connection. If serialization
// or writing to the connection fails, it returns an error.
//
// Returns nil on success or an error on failure.
func (s *Server) Respond(message interfaces.ResponseMessage) error {

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
