package main

import (
	"fmt"

	"github.com/josnelihurt/simple-go-gaming/engine/util"
	"github.com/veandco/go-sdl2/mix"
)

func main() {
	mix.OpenAudio(mix.DEFAULT_FREQUENCY, 16, 2, 4096)
	music, _ := mix.LoadMUS("sounds/scene.mp3")
	music.Play(-1)
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
			break
		}
	}
}
