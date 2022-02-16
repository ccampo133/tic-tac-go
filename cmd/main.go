package main

import (
	"github.com/ccampo133/tic-tac-toe/internal/game"
)

func main() {
	g := game.NewGame()
	if err := g.MainLoop(); err != nil {
		panic(err)
	}
}
