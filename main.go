package main

import (
	"fmt"
	"os"

	"github.com/Mira1702-git/tic-tac-toe-arena/internal/cli"
	"github.com/Mira1702-git/tic-tac-toe-arena/internal/game"
)

func main() {
	config, err := cli.Parse(os.Args[1:])
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Print(cli.Usage())
		os.Exit(1)
	}

	if config.Help {
		fmt.Print(cli.Usage())
		return
	}

	game.Run(config)
}
