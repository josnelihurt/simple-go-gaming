package main

import (
	"fmt"
	"time"

	"github.com/josnelihurt/simple-go-gaming/engine"
	"github.com/josnelihurt/simple-go-gaming/engine/util"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth          = 720
	screenHeight         = 800
	targetTicksPerSecond = 60
)

var delta float64

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

func main() {
	util.Logger = make(chan string, 1024)
	go util.DoLog(util.Logger)
	util.Logger <- "Starting up.."
	defer close(util.Logger)

	loadResources()
	engineComponents, err := engine.NewSDLComponents(screenWidth, screenHeight, "simple-game")
	if err != nil {
		util.Logger <- "Unable to start"
		return
	}

	defer engineComponents.Release()

	elementPool := engine.NewElementPool()
	loadElements(&elementPool, engineComponents.Renderer, engineComponents.AudioDev)

	for {
		frameStartTimer := time.Now()
		if continueFlag := inputHandler(); continueFlag == false {
			util.Logger <- fmt.Sprintf("exiting gameLoop:")
			return
		}
		engineComponents.Renderer.SetDrawColor(255, 255, 0, 0)
		engineComponents.Renderer.Clear()
		elementPool.UpdateElements(engineComponents.Renderer)
		if err := engine.CheckColisions(&elementPool); err != nil {
			util.Logger <- fmt.Sprintf("checking collisions:%v", err)
		}

		engineComponents.Renderer.Present()
		delta = time.Since(frameStartTimer).Seconds() * targetTicksPerSecond
	}
}
