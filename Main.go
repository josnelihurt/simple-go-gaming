package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	screenWidth          = 720
	screenHeight         = 800
	targetTicksPerSecond = 60
)

var logger chan string
var delta float64

func doLog(input <-chan string) {
	for line := range input {
		fmt.Println(line)
	}
}
func createRenderer() (*sdl.Renderer, *sdl.Window, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		logger <- fmt.Sprintln("initializing SDL:", err)
		panic(err)
	}
	window, err := sdl.CreateWindow(
		"Demo game 1",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_SHOWN)
	if err != nil {
		logger <- fmt.Sprintln("initializing window:", err)
		return nil, nil, err
	}
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		logger <- fmt.Sprintf("inititalizing renderer:%v", err)
		return nil, nil, err
	}

	if err := ttf.Init(); err != nil {
		logger <- fmt.Sprintf("initializing ttf:%v", err)
		return nil, nil, err
	}
	return renderer, window, nil
}
func loadElements(elementPool *elementPool, renderer *sdl.Renderer) {
	scoreRenderer := newScoreRenderer()
	score := &element{active: true}
	score.addCompoenent(scoreRenderer)
	score.z = 10
	elementPool.insertElement(score)
	elementPool.insertElement(newPlayer(renderer))
	elementPool.insertSlice(initBulletPool(renderer, scoreRenderer))
	elementPool.insertSlice(createEnemySwarm(renderer))
	elementPool.insertElement(newBackground(renderer))
}

func main() {
	logger = make(chan string, 1024)
	go doLog(logger)
	logger <- "Starting up.."
	defer close(logger)

	loadResources()
	renderer, window, err := createRenderer()
	if err != nil {
		logger <- "unable to create renderer"
	}

	elementPool := newElementPool()
	loadElements(&elementPool, renderer)
	for _, currentElement := range elementPool.elements {
		logger <- fmt.Sprintf("Element :%v", currentElement)
	}

	defer window.Destroy()
	defer renderer.Destroy()

	for {
		frameStartTimer := time.Now()
		if continueFlag := inputHandler(); continueFlag == false {
			logger <- fmt.Sprintf("exiting gameLoop:")
			return
		}
		renderer.SetDrawColor(255, 255, 0, 0)
		renderer.Clear()
		elementPool.updateElements(renderer)
		if err := checkColisions(&elementPool); err != nil {
			logger <- fmt.Sprintf("checking collisions:%v", err)
		}

		renderer.Present()
		delta = time.Since(frameStartTimer).Seconds() * targetTicksPerSecond
	}
}
