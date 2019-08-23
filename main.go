package main

import (
	"fmt"

	"os"

	"github.com/josnelihurt/simple-go-gaming/engine/util"
)

func main() {
	util.Logger = make(chan string, 1024)
	go util.DoLog(util.Logger)
	util.Logger <- "Starting up.."
	defer close(util.Logger)
	loadResources()

	logic := newGameLogic()
	defer logic.Release()

	for {
		if continueFlag := logic.loop(); continueFlag == false {
			util.Logger <- fmt.Sprintf("exiting main:")
			os.Exit(0)
		}
	}
}
