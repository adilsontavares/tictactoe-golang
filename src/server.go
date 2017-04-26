package main

import (
	"./server"
	"./sockets"
	"fmt"
	"os"
)

func fatal(format string, a...interface{}) {

	fmt.Printf(format, a...)
	fmt.Println()
	os.Exit(1)
}

func main() {

	fmt.Println("Starting server.")
	defer fmt.Println("Server stopped.")

	sockets.PrintsLog = true

	if len(os.Args) != 2 {
		fatal("USAGE: <PORT>")
	}

	port := os.Args[1]
	
	if !server.Init(":" + port) {
		fatal("Could not start server.")
	}

	fmt.Println("Server is running.")

	server.Run()
}