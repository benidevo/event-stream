package request

// MessageBuilder constructs a Request Message from a byte array.
//
// It uses DeserializeMessage to parse the byte array into a Request Message.
// Returns the constructed Request Message or an error if deserialization fails.
func MessageBuilder(data []byte) (*Message, error) {
	message, error := DeserializeMessage(data)
	if error != nil {
		return nil, error

	}

	return message, nil
}
