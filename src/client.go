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

	if !client.Init("127.0.0.1:8080") {
		fatal("Could not start client.")
	}

	defer console.End()

	sockets.PrintsLog = false

	client.Run()
}