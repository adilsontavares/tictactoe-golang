package main

import (
	"./console"
	"./tictactoe"
)

func main() {

	console.Begin()
	defer console.End()

	game := tictactoe.New()
	game.Loop()
}