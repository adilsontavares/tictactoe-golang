package server

import (
	"net"
	"fmt"
)

var (
	
	clients []*Client

	initialized bool

	ln net.Listener
	conn net.Conn
)

func Init(laddr string) bool {

	var err error
	if ln, err = net.Listen("tcp", laddr); err != nil {
		
		fmt.Println("Listen error.")
		return false
	}

	initialized = true

	return true
}

func loop() {

	for {

		if conn, err := ln.Accept(); err != nil {
			fmt.Println("Error accepting a new client.")
		} else {
			go handleClient(conn)
		}
	}
}

func Run() bool {

	if !initialized {
		
		fmt.Println("Server could not run because it has not been initialized.")
		return false
	}

	loop()

	return true
}

func handleClient(conn net.Conn) {

	client := NewClient(conn)
	clients = append(clients, client)

	client.log("CONNECTED (%v total connected).", len(clients))
	client.startNewGame()

	for !client.Socket.IsClosed() {
		client.Socket.ReadMessage()
	}

	for i, c := range clients {
		if c == client {
			clients = append(clients[:i], clients[i+1:]...)
			break
		}
	}

	client.log("DISCONNECTED (%v total connected).", len(clients))
}