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

func createEnemySwarm(renderer *sdl.Renderer) (swarm []*element) {
	const rows = 3
	const colums = 6
	for i := 0; i < colums; i++ {
		for j := 0; j < rows; j++ {
			x := (float64(i)/colums)*screenWidth + (basicEnemySize / 2.0)
			y := float64(j)*basicEnemySize + (basicEnemySize / 2.0) + 50

			enemy := newBasicEnemy(renderer, vector{x: x, y: y})
			enemy.tag = fmt.Sprintf("x:%v,y:%v", i, j)
			swarm = append(swarm, enemy)
		}
	}
	return swarm
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
func loadElements(renderer *sdl.Renderer) {
	scoreRenderer := newScoreRenderer()
	score := &element{active: true}
	score.addCompoenent(scoreRenderer)
	elements = append(elements, score)
	elements = append(elements, newPlayer(renderer))
	elements = append(elements, initBulletPool(renderer, scoreRenderer)...)
	elements = append(elements, createEnemySwarm(renderer)...)
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

	loadElements(renderer)
	defer window.Destroy()
	defer renderer.Destroy()

	backround := newBackground(renderer)
	for {
		//if continueFlag := gameLoop(renderer); continueFlag == false {
		//	break
		//}Texture dimensions are limited to 8192x8192

		frameStartTimer := time.Now()
		if continueFlag := inputHandler(); continueFlag == false {
			logger <- fmt.Sprintf("exiting gameLoop:")
			return
		}
		renderer.SetDrawColor(255, 255, 0, 0)
		renderer.Clear()
		if err := backround.update(); err != nil {
			logger <- fmt.Sprintf("updating fail:%v", err)
		}
		if err := backround.draw(renderer); err != nil {
			logger <- fmt.Sprintf("drawing fail:%v", err)
		}

		for _, currentElement := range elements {
			if currentElement.active {
				if err := currentElement.update(); err != nil {
					logger <- fmt.Sprintf("updating fail:%v", err)
				}
				if err := currentElement.draw(renderer); err != nil {
					logger <- fmt.Sprintf("drawing fail:%v", err)
				}
			}
		}
		if err := checkColisions(); err != nil {
			logger <- fmt.Sprintf("checking collisions:%v", err)
		}

		renderer.Present()
		delta = time.Since(frameStartTimer).Seconds() * targetTicksPerSecond
	}
}
