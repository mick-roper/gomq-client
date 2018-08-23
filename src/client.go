package gomq

import (
	"errors"
	"net"

	"./core"
)

// AmqpClient contains the methods required to interract with an AMQP server
type AmqpClient interface {
}

// Client that communicates with an AMQP server
type Client struct {
	conn   net.Conn
	recvCh chan core.Frame
	sendCh chan core.Frame
}

// NewClient creates a new client
func NewClient(addr string) (*Client, error) {
	return nil, errors.New("not implemented")
}
