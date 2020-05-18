package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/josnelihurt/simple-go-gaming/pkg/engine"
	"github.com/josnelihurt/simple-go-gaming/pkg/engine/util"
	"github.com/veandco/go-sdl2/sdl"
)

var delta float64 // <-- where may I put you???

type gameLogic struct {
	sdlComponents         *engine.SDLComponents
	elementManager        engine.ElementManager
	gameOverScreen        *engine.Element
	player                *engine.Element
	swarm                 *enemySwarm
	swarmActivationTime   int
	pressAnyKeyToContinue chan struct{}
}

func NewGameLogic() *gameLogic {
	context := &gameLogic{}
	context.swarmActivationTime = 2000
	context.initSDLComponents()
	context.initElementManager()
	context.finishCondition()
	go engine.PlayMusicLoop(resMusicBackground)
	context.pressAnyKeyToContinue = make(chan struct{}, 2)
	return context
}
func (context *gameLogic) Release() {
	context.sdlComponents.Release()
}
func (context *gameLogic) initSDLComponents() {
	var err error
	context.sdlComponents, err = engine.NewSDLComponents(screenWidth, screenHeight, "simple-game")
	if err != nil {
		panic("Unable to start sdl components")
	}
}
func (context *gameLogic) initElementManager() {
	context.elementManager = engine.NewElementManager()
	context.elementManager.InsertSlice(context.createElements())
}
func (context *gameLogic) enemyAwaker() {
	for {
		actives := context.swarm.getActives()
		if len(actives) == 0 {
			context.elementManager.BroadcastMessage(&engine.Message{
				Code: msgLevelUp,
			})
			context.swarm.activateAll()
			context.swarm.setInitialPositions()
			context.swarmActivationTime /= 2

			continue
		}
		rand.Seed(time.Now().UnixNano())
		actives[rand.Intn(len(actives))].GetFirstComponent(&enemyMover{}).(*enemyMover).active = true
		select {
		case <-context.pressAnyKeyToContinue:
			context.swarm.setInitialPositions()
			context.swarm.deActivateAll()
			return
		case <-time.After(time.Duration(context.swarmActivationTime) * time.Millisecond):
			break
		}
	}
}
func (context *gameLogic) finishCondition() {
	context.player.RegisterEmitterrCallback(func(message *engine.Message) error {
		context.elementManager.GetElementsByTag(tagLevel)[0].BroadcastMessage(message)
		if message.Code == msgPlayerDead {
			context.elementManager.DisableElementsByTag(tagEnemy, tagPlayer, tagScore)
			context.gameOverScreen.Active = true
			context.pressAnyKeyToContinue <- struct{}{}
		}
		return nil
	})
}
func (context *gameLogic) createElements() (elements []*engine.Element) {
	score := newScore()
	context.player = newPlayer(context.sdlComponents)
	background := newBackground(context.sdlComponents)
	bulletPool := initBulletPool(context.sdlComponents.Renderer)
	context.swarm = newEnemySwarm(context.sdlComponents.Renderer)
	level := newLevel(context.sdlComponents)
	engine.BindMessages(context.swarm.enemies, context.player)
	engine.BindMessages(context.swarm.enemies, score)
	context.gameOverScreen = newGameOverScreen(context.sdlComponents)
	go context.enemyAwaker()

	elements = append(elements, score, context.player, background, context.gameOverScreen, level)
	elements = append(elements, bulletPool...)
	elements = append(elements, context.swarm.enemies...)
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
func (context *gameLogic) Loop() (continueFlag bool) {
	frameStartTimer := time.Now()
	if continueFlag = inputHandler(); continueFlag == false {
		util.Logger <- fmt.Sprintf("exiting gameLoop:")
		return continueFlag
	}

	if context.gameOverScreen.Active {

		keys := sdl.GetKeyboardState()
		if keys[sdl.SCANCODE_SPACE] == 1 {
			context.gameOverScreen.Active = false
			context.swarmActivationTime = 2000
			context.elementManager.EnableElementsByTag(tagEnemy, tagPlayer, tagScore)
			go context.enemyAwaker()
		}
	}

	context.updateRenderer()
	delta = time.Since(frameStartTimer).Seconds() * targetTicksPerSecond
	return true

}
