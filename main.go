package main

import (
	"fmt"
	"time"

	"github.com/josnelihurt/simple-go-gaming/engine"
	"github.com/josnelihurt/simple-go-gaming/engine/util"
)

const (
	screenWidth          = 720
	screenHeight         = 800
	targetTicksPerSecond = 60
)

var delta float64

func bindMessage(source *engine.Element, destination *engine.Element) {
	source.RegisterEmmiterCallback(func(message *engine.Message) error {
		destination.BroadcastMessageToComponents(message)
		return nil
	})
}
func bindMessages(sources []*engine.Element, destination *engine.Element) {
	for _, source := range sources {
		bindMessage(source, destination)
	}
}
func createElements(components *engine.SDLComponents) (elements []*engine.Element) {
	score := newScore()
	player := newPlayer(components.Renderer, components.AudioDev)
	background := newBackground(components.Renderer)
	bulletPool := initBulletPool(components.Renderer)
	bindMessages(bulletPool, player)
	enemySwarm := createEnemySwarm(components.Renderer)
	bindMessages(enemySwarm, score)

	elements = append(elements, score, player, background)
	elements = append(elements, bulletPool...)
	elements = append(elements, enemySwarm...)
	return
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

	manager := engine.NewElementManager()
	manager.InsertSlice(createElements(engineComponents))

	for {
		frameStartTimer := time.Now()
		if continueFlag := inputHandler(); continueFlag == false {
			util.Logger <- fmt.Sprintf("exiting gameLoop:")
			return
		}
		engineComponents.Renderer.SetDrawColor(255, 255, 0, 0)
		engineComponents.Renderer.Clear()
		manager.UpdateElements(engineComponents.Renderer)
		if err := manager.CheckColisions(); err != nil {
			util.Logger <- fmt.Sprintf("checking collisions:%v", err)
		}

		engineComponents.Renderer.Present()
		delta = time.Since(frameStartTimer).Seconds() * targetTicksPerSecond
	}
}
