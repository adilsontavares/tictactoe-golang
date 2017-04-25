package main

import (
	"./server"
	"./sockets"
	"fmt"
)

func main() {

	fmt.Println("Starting server.")
	defer fmt.Println("Server stopped.")

	sockets.PrintsLog = true

	server.Init(":8080")
	fmt.Println("Server is running.")

	server.Run()
}