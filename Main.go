package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 800
)

var logger chan string

func doLog(input <-chan string) {
	for line := range input {
		fmt.Println(line)
	}
}

func main() {
	logger = make(chan string, 1024)
	defer close(logger)
	go doLog(logger)
	logger <- "Starting up.."
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
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		logger <- fmt.Sprintln("inititalizing renderer:", err)
	}
	defer renderer.Destroy()

	player1 := newPlayer(renderer)
	elements = append(elements, player1)

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			x := (float64(i)/5)*screenWidth + (basicEnemySize / 2.0)
			y := float64(j)*basicEnemySize + (basicEnemySize / 2.0)

			enemy := newBasicEnemy(renderer, vector{x: x, y: y})
			enemy.tag = fmt.Sprintf("x:%v,y:%v", i, j)
			if err != nil {
				logger <- fmt.Sprintln("creating enemy:", err)
				return
			}
			elements = append(elements, enemy)
		}
	}

	elements = append(elements, initBulletPool(renderer)...)

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				logger <- fmt.Sprintln("exiting:")
				return
			}
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		for _, currentElement := range elements {
			if currentElement.active {
				if err := currentElement.update(); err != nil {
					logger <- fmt.Sprintln("updating fail:", err)
				}
				if err := currentElement.draw(renderer); err != nil {
					logger <- fmt.Sprintln("drawing fail:", err)
				}
			}
		}
		if err := checkColisions(); err != nil {
			logger <- fmt.Sprintln("checking collisions:", err)
		}

		renderer.Present()
	}

}
