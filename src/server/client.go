package server

import (
	"net"
	"../board"
	"../sockets"
)

type Client struct {

	waitingPlay bool

	Item int

	Socket *sockets.Socket
	Board *board.Board

	PlayerScore int
	ComputerScore int
}

func (client *Client) log(message string, format ...interface{}) {
	client.Socket.Log(message, format...)
}

func (client *Client) Reset() {

	client.waitingPlay = false
	client.Board.Reset()
}

func (client *Client) sendMessage(data interface{}) bool {
	return client.Socket.SendMessage(data)
}

func NewClient(conn net.Conn) (*Client) {

	client := Client{}
	client.Item = board.ItemX
	client.Board = board.New()
	client.Socket = sockets.New(conn)
	client.Socket.ActorName = "Client"
	client.Socket.InterpretClosure = client.interpretMessage

	client.startNewGame()

	return &client
}