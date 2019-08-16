package main

import (
	"fmt"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	screenWidth          = 600
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
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			x := (float64(i)/5)*screenWidth + (basicEnemySize / 2.0)
			y := float64(j)*basicEnemySize + (basicEnemySize / 2.0) + 50

			enemy := newBasicEnemy(renderer, vector{x: x, y: y})
			enemy.tag = fmt.Sprintf("x:%v,y:%v", i, j)
			swarm = append(swarm, enemy)
		}
	}
	return swarm
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func unPackFileFromAsset(folder, filename string) {
	os.Mkdir(folder, os.ModePerm)
	out := fmt.Sprintf("%v/%v", folder, filename)
	if _, err := os.Stat(out); os.IsNotExist(err) {
		file, _ := Asset(out)
		fileIO, _ := os.Create(out)
		if _, err := fileIO.Write(file); err == nil {
			logger <- fmt.Sprintf("%v unpacked from resources", out)
		}
		fileIO.Sync()
		fileIO.Close()
	} else {
		logger <- fmt.Sprintf("%v already in fs", out)
	}
}

func loadResources() {
	// remember you must run
	// $ go get -u github.com/jteeuwen/go-bindata/...
	// $ go-bindata sprites/... fonts/... sounds/...
	unPackFileFromAsset("fonts", "Starjout.ttf")
	unPackFileFromAsset("sprites", "basic_enemy.bmp")
	unPackFileFromAsset("sprites", "bullet.bmp")
	unPackFileFromAsset("sprites", "player.bmp")
	unPackFileFromAsset("sounds", "NFF-laser.wav")
}
func main() {
	logger = make(chan string, 1024)
	go doLog(logger)
	logger <- "Starting up.."
	defer close(logger)

	loadResources()

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
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		logger <- fmt.Sprintf("inititalizing renderer:%v", err)
		return
	}
	defer renderer.Destroy()

	if err := ttf.Init(); err != nil {
		logger <- fmt.Sprintf("initializing ttf:%v", err)
		return
	}

	scoreRenderer := newScoreRenderer()
	score := &element{active: true}
	score.addCompoenent(scoreRenderer)
	elements = append(elements, score)
	elements = append(elements, newPlayer(renderer))
	elements = append(elements, initBulletPool(renderer, scoreRenderer)...)
	elements = append(elements, createEnemySwarm(renderer)...)

	for {
		frameStartTimer := time.Now()
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
		delta = time.Since(frameStartTimer).Seconds() * targetTicksPerSecond

	}

}
