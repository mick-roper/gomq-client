package gomq

import "errors"

// Message that is sent to or received from an AMQP server
type Message struct {
	Headers map[string]string
	Payload []byte
}

// NewMessage creates a new message
func NewMessage(payload []byte) (*Message, error) {
	h := make(map[string]string, 0)
	return NewMessageWithHeaders(payload, h)
}

// NewMessageWithHeaders creates a new message with custom headers
func NewMessageWithHeaders(payload []byte, headers map[string]string) (*Message, error) {
	return nil, errors.New("not implemented")
}
