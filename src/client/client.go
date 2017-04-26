package client

import (
	"../sockets"
	"../console"
	"fmt"
	"os"
	"net"
)

var (
	
	initialized bool
	game *Game
)

func Init(laddr string) bool {

	if initialized {
		return false
	}
	
	conn, err := net.Dial("tcp", laddr)
	if err != nil {
		return false
	}

	game = NewGame()
	game.Socket = sockets.New(conn)
	game.Socket.InterpretClosure = game.interpretMessage

	initialized = true

	return true
}

func listen() {

	for !game.Socket.IsClosed() {
		game.Socket.ReadMessage()
	}

	console.End()
	fmt.Println("Connection with server was closed.")
	os.Exit(1)
}

func Run() {

	if !initialized {
		return
	}

	go listen()
	game.Loop()
}