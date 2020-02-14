package main

import (
	"github.com/NBR41/gosnake"
)

const (
	size = 20
	col  = 20
	row  = 10
)

func main() {
	game, err := gosnake.NewGame(size, col, row)
	if err != nil {
		panic(err)
	}
	if err := game.Run(); err != nil {
		panic(err)
	}
}
