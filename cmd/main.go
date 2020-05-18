package main

import (
	"fmt"
	"os"

	"github.com/josnelihurt/simple-go-gaming/pkg/engine/game"
	"github.com/josnelihurt/simple-go-gaming/pkg/engine/util"
)

func main() {
	util.Logger = make(chan string, 1024)
	go util.DoLog(util.Logger)
	util.Logger <- "Starting up.."
	defer close(util.Logger)
	game.LoadResources()

	logic := game.NewGameLogic()
	defer logic.Release()

	for {
		if continueFlag := logic.Loop(); continueFlag == false {
			util.Logger <- fmt.Sprintf("exiting main:")
			os.Exit(0)
		}
	}
}
