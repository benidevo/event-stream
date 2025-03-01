# Event Stream

This project implements a basic Kafka-compatible event streaming server in Go, focusing on the binary protocol used by Kafka for communication between clients and servers.

## Key Components

### Server

The `server` package provides a basic TCP server implementation that can:

- Listen on a specified host and port
- Accept connections
- Read incoming messages (requests)
- Send responses

### Message Handling

The application implements binary message serialization and deserialization following the Kafka protocol format:

- **Request Messages**: Contain fields for Size, ApiKey, ApiVersion, and CorrelationId
- **Response Messages**: Contain fields for Size and CorrelationId

## Protocol Implementation

The current implementation focuses on the binary protocol used in Kafka:

1. **Message Format**: All messages start with a size (int32) and correlation ID (int32)
2. **Serialization**: Messages are serialized to binary using Go's binary package with BigEndian byte order
3. **Deserialization**: Incoming binary data is parsed into structured message types

## Development

### Running the Server

```bash
# Run locally
./run.sh
```

### Formatting Code

```bash
make fmt
```

## Current Status

This project is still in development. The current implementation:

- Establishes a basic TCP server
- Handles simple message serialization and deserialization
- Provides the foundation for implementing more Kafka protocol features

Next steps include implementing:

- Full request handling for different Kafka API types
- Topic management
- Message storage and retrieval
- Consumer group coordination

## References

- [Apache Kafka Protocol Guide](https://kafka.apache.org/protocol.html)
- [Kafka Wire Format](https://kafka.apache.org/protocol#protocol_messages)
