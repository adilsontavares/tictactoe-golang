package server

import (
	"bufio"
	"net"
)

type Client struct {

	closed bool

	Conn net.Conn
}

func (client *Client) HandleMessage() {

	if client.closed {
		return
	}

	message, err := bufio.NewReader(client.Conn).ReadString('\n')

	if err != nil {

		client.closed = true
		return
	}

	client.handleMessage(message)
}

func (client *Client) handleMessage(message string) {


}

func NewClient(conn net.Conn) (*Client) {

	client := Client{}
	client.Conn = conn

	return &client
}

func (client *Client) Closed() bool {
	return client.closed
}