package main

import (
	"fmt"
	"time"

	"github.com/josnelihurt/simple-go-gaming/engine"
	"github.com/josnelihurt/simple-go-gaming/engine/util"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	screenWidth          = 720
	screenHeight         = 800
	targetTicksPerSecond = 60
)

var delta float64

func createRenderer() (*sdl.Renderer, *sdl.Window, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		util.Logger <- fmt.Sprintln("initializing SDL:", err)
		panic(err)
	}
	window, err := sdl.CreateWindow(
		"Demo game 1",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_SHOWN)
	if err != nil {
		util.Logger <- fmt.Sprintln("initializing window:", err)
		return nil, nil, err
	}
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		util.Logger <- fmt.Sprintf("inititalizing renderer:%v", err)
		return nil, nil, err
	}

	if err := ttf.Init(); err != nil {
		util.Logger <- fmt.Sprintf("initializing ttf:%v", err)
		return nil, nil, err
	}
	return renderer, window, nil
}
func loadElements(elementPool *engine.ElementPool, renderer *sdl.Renderer, audioDev sdl.AudioDeviceID) {
	elementPool.InsertElement(newScore())
	elementPool.InsertElement(newPlayer(renderer, audioDev))
	elementPool.InsertSlice(initBulletPool(renderer, func() {
		elementPool.GetElementsByTag("score")[0].GetComponent(&scoreCounter{}).(*scoreCounter).increase()
	}))
	elementPool.InsertSlice(createEnemySwarm(renderer, func() {
		for _, currentEnemy := range elementPool.GetElementsByTag("enemy") {
			if currentEnemy.Active == true {
				currentEnemyMover := currentEnemy.GetComponent(&enemyMover{})
				currentEnemyMover.(*enemyMover).active = true
				break
			}
		}

	}))
	elementPool.InsertElement(newBackground(renderer))
}
func openAudioDevice() sdl.AudioDeviceID {
	currenAudioDriver := sdl.GetCurrentAudioDriver()
	util.Logger <- currenAudioDriver
	soundInfo := &sdl.AudioSpec{
		Freq:     44100,
		Format:   32784,
		Channels: 2,
		Silence:  0,
		Samples:  4096,
	}
	dev, err := sdl.OpenAudioDevice("", false, soundInfo, nil, 0)
	if err != nil {
		util.Logger <- fmt.Sprintf("error opeing audio dev:%v", err)
	}
	return dev
}
func main() {
	util.Logger = make(chan string, 1024)
	go util.DoLog(util.Logger)
	util.Logger <- "Starting up.."
	defer close(util.Logger)

	loadResources()
	renderer, window, err := createRenderer()
	if err != nil {
		util.Logger <- "unable to create renderer"
	}
	defer window.Destroy()
	defer renderer.Destroy()
	audioDev := openAudioDevice()
	defer sdl.CloseAudioDevice(audioDev)

	elementPool := engine.NewElementPool()
	loadElements(&elementPool, renderer, audioDev)

	for {
		frameStartTimer := time.Now()
		if continueFlag := inputHandler(); continueFlag == false {
			util.Logger <- fmt.Sprintf("exiting gameLoop:")
			return
		}
		renderer.SetDrawColor(255, 255, 0, 0)
		renderer.Clear()
		elementPool.UpdateElements(renderer)
		if err := engine.CheckColisions(&elementPool); err != nil {
			util.Logger <- fmt.Sprintf("checking collisions:%v", err)
		}

		renderer.Present()
		delta = time.Since(frameStartTimer).Seconds() * targetTicksPerSecond
	}
}
