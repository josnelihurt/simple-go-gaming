package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	"github.com/josnelihurt/simple-go-gaming/engine"
	"github.com/josnelihurt/simple-go-gaming/engine/util"
)

var delta float64 // <-- where may I put you???

type gameLogic struct {
	sdlComponents  *engine.SDLComponents
	elementManager engine.ElementManager
	gameOverScreen *engine.Element
	player         *engine.Element
}

func newGameLogic() *gameLogic {
	context := &gameLogic{}
	context.initSDLComponents()
	context.initElementManager()
	context.finishCondition()
	go playMusic()
	return context
}
func (context *gameLogic) Release() {
	context.sdlComponents.Release()
}
func playMusic() error {
	mp3File, err := os.Open("sounds/scene.mp3")
	if err != nil {
		return err
	}
	defer mp3File.Close()

	decoder, err := mp3.NewDecoder(mp3File)
	if err != nil {
		return err
	}

	player, err := oto.NewPlayer(decoder.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}
	defer player.Close()

	for {
		fmt.Printf("Length: %d[bytes]\n", decoder.Length())
		if _, err := io.Copy(player, decoder); err != nil {
			return err
		}
		decoder.Seek(0, 0)
	}
}

// It doesn't work in my laptop
// mix.Init(mix.INIT_MP3)
// mix.OpenAudio(44100, //mix.DEFAULT_FREQUENCY,
// 	16, 2, 4096)
// music, err := mix.LoadMUS("sounds/scene.mp3")
// if err != nil {
// 	fmt.Println(err)
// }
// music.Play(-1)
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
