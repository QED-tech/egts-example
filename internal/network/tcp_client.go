package network

import (
	"fmt"
	"net"
	"time"
)

type Client struct {
	connection  net.Conn
	idleTimeout time.Duration
}

func NewTCPClient(address string, idleTimeout time.Duration) (*Client, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to egts server: %v", err)
	}

	return &Client{
		connection:  conn,
		idleTimeout: idleTimeout,
	}, nil
}

func (c *Client) Send(message []byte) ([]byte, error) {
	if err := c.connection.SetDeadline(time.Now().Add(c.idleTimeout)); err != nil {
		return nil, err
	}

	if _, err := c.connection.Write(message); err != nil {
		return nil, err
	}

	response := make([]byte, 2048)
	count, err := c.connection.Read(response)
	if err != nil {
		return nil, fmt.Errorf("failed to read message, err: %v", err)
	}

	return response[:count], nil
}
