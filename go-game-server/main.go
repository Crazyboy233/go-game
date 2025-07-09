package main

import (
	"go-game-server/internal/command"
)

func main() {
	input := command.Parser()
	command.Dispatch(input)
}
