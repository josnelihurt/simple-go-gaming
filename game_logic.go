package main

import (
	"fmt"
	"time"

	"github.com/josnelihurt/simple-go-gaming/engine"
	"github.com/josnelihurt/simple-go-gaming/engine/util"
)

var delta float64 // <-- where may I put you???

type gameLogic struct {
	sdlComponents   *engine.SDLComponents
	elementManager  engine.ElementManager
	player          *engine.Element
	backgroundSound *engine.SoundPlayer
	gameOverScreen  *engine.Element
}

func newGameLogic() *gameLogic {
	context := &gameLogic{}
	context.initSDLComponents()
	context.initElementManager()
	context.finishCondition()
	context.backgroundSound = engine.NewSoundPlayer(nil, "sound/scene.wav", context.sdlComponents.AudioDev, []int{}, "")
	go context.backgroundSound.Play()
	return context
}
func (context *gameLogic) Release() {
	context.sdlComponents.Release()
}
func (context *gameLogic) initSDLComponents() {
	var err error
	context.sdlComponents, err = engine.NewSDLComponents(screenWidth, screenHeight, "simple-game")
	if err != nil {
		panic("Unable to start sdl commponents")
	}
}
func (context *gameLogic) initElementManager() {
	context.elementManager = engine.NewElementManager()
	context.elementManager.InsertSlice(context.createElements())
}
func (context *gameLogic) enemyAwaker(enemies []*engine.Element) {
	for _, currentElement := range enemies {
		if currentElement.Active {
			currentElement.GetComponent(&enemyMover{}).(*enemyMover).active = true
			time.Sleep(4000 * time.Millisecond)
		}
	}
}
func (context *gameLogic) finishCondition() {
	context.player.RegisterEmmiterCallback(func(message *engine.Message) error {
		if message.Code == msgPlayerDead {
			context.elementManager.DisableElementsByTag("enemy", "player", "score")
			context.gameOverScreen.Active = true
		}
		return nil
	})
}
func (context *gameLogic) createElements() (elements []*engine.Element) {
	score := newScore()
	context.player = newPlayer(context.sdlComponents)
	background := newBackground(context.sdlComponents)
	bulletPool := initBulletPool(context.sdlComponents.Renderer)
	enemySwarm := createEnemySwarm(context.sdlComponents.Renderer)
	go context.enemyAwaker(enemySwarm)
	engine.BindMessages(enemySwarm, context.player)
	engine.BindMessages(enemySwarm, score)
	context.gameOverScreen = newGameOverScreen(context.sdlComponents)

	elements = append(elements, score, context.player, background, context.gameOverScreen)
	elements = append(elements, bulletPool...)
	elements = append(elements, enemySwarm...)
	return
}
func (context *gameLogic) updateRenderer() {
	context.sdlComponents.Renderer.SetDrawColor(255, 255, 0, 0)
	context.sdlComponents.Renderer.Clear()
	context.elementManager.UpdateElements(context.sdlComponents.Renderer)
	if err := context.elementManager.CheckCollisions(); err != nil {
		util.Logger <- fmt.Sprintf("checking collisions:%v", err)
	}
	context.sdlComponents.Renderer.Present()
}
func (context *gameLogic) loop() (continueFlag bool) {
	frameStartTimer := time.Now()
	if continueFlag = inputHandler(); continueFlag == false {
		util.Logger <- fmt.Sprintf("exiting gameLoop:")
		return continueFlag
	}
	context.updateRenderer()
	delta = time.Since(frameStartTimer).Seconds() * targetTicksPerSecond
	return true

}
