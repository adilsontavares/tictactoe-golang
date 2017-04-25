package client

import (
	"github.com/nsf/termbox-go"
	"../sockets"
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
		termbox.Flush()
	}
}

func Run() {

	if !initialized {
		return
	}

	go listen()
	game.Loop()
}