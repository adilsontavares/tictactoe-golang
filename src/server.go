package main

import (
	"./server"
	"fmt"
)

func main() {

	fmt.Println("Starting server.")
	defer fmt.Println("Server stopped.")

	server.Init(":8080")
	fmt.Println("Server is running.")

	server.Run()
}