package main

import (
	"github.com/NBR41/gosnake"
)

func main() {
	game, err := gosnake.NewGame(20, 20, 10)
	if err != nil {
		panic(err)
	}
	if err := game.Run(); err != nil {
		panic(err)
	}
}
