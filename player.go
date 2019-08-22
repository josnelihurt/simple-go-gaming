package main

import (
	"math"
	"time"

	"github.com/josnelihurt/simple-go-gaming/engine"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed        = 5.0
	playerShotCooldown = time.Millisecond * 200
	playerScale        = 0.7
)

func newPlayer(components *engine.SDLComponents) *engine.Element {
	context := &engine.Element{}
	context.Z = 10
	context.Active = true
	context.Rotation = 180
	context.Tag = "player"

	currentSpriteRenderer := engine.NewSpriteRenderer(context, components.Renderer, "sprites/player.png", playerScale)
	playerDefaultPosition := engine.Vector{
		X: screenWidth / 2.0,
		Y: screenHeight - currentSpriteRenderer.ScaledHeight/2.0,
	}
	context.Position = playerDefaultPosition
	context.AddComponent(currentSpriteRenderer)
	context.AddComponent(newPlayerStarter(context, playerDefaultPosition))
	context.AddComponent(newKeyboardShooter(context, playerShotCooldown, components.AudioDev))
	textRenderer := engine.NewTextRenderer(
		&engine.Vector{X: upperTextY, Y: upperTextY},
		defaultFontSize,
		sdl.Color{R: 255, G: 255, B: 255})
	context.AddComponent(textRenderer)
	context.AddComponent(newPlayerLifeCounter(context, textRenderer))
	allowedRect := &engine.Rect{
		X:      currentSpriteRenderer.ScaledWidth / 2.0,
		Y:      4 * screenHeight / 5,
		Width:  screenWidth - currentSpriteRenderer.ScaledWidth,
		Height: screenHeight/5 - currentSpriteRenderer.ScaledHeight/2.0}
	context.AddComponent(engine.NewKeyboardMover(context, allowedRect, &delta, playerSpeed))
	context.AddComponent(engine.NewCollisionDetecter(context, true, "enemy"))
	context.AddComponent(engine.NewSoundPlayer(context, "sounds/explosion.wav", components.AudioDev, []int{engine.MsgCollision, msgHitPlayer}, "enemy"))
	context.Collisions = append(context.Collisions,
		engine.Circle{
			Center: &context.Position,
			Radius: math.Max(currentSpriteRenderer.ScaledWidth, currentSpriteRenderer.ScaledHeight) / 2,
		})

	return context
}
