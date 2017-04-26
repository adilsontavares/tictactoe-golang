package main

import (
	"./console"
	"./sockets"
	"./client"
	"fmt"
	"os"
)

func fatal(format string, a...interface{}) {

	console.End()
	fmt.Printf(format, a...)
	fmt.Println()
	os.Exit(1)
}

func main() {

	console.Begin()

	if !client.InitLog("../logs/client-log.txt") {
		fatal("Could not initialize log.")
	}

	if len(os.Args) != 3 {
		fatal("USAGE: <IP_ADDR> <PORT>")
	}

	addr := os.Args[1]
	port := os.Args[2]

	if !client.Init(addr + ":" + port) {
		fatal("Could not connect to server.")
	}

	defer console.End()

	sockets.PrintsLog = false

	client.Run()
}