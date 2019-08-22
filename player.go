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

func newPlayer(renderer *sdl.Renderer, audioDev sdl.AudioDeviceID) *engine.Element {
	player := &engine.Element{}
	player.Z = 10
	player.Active = true
	player.Rotation = 180
	player.Tag = "player"

	currentSpriteRenderer := engine.NewSpriteRenderer(player, renderer, "sprites/player.png", playerScale)
	player.Position = engine.Vector{
		X: screenWidth / 2.0,
		Y: screenHeight - currentSpriteRenderer.ScaledHeight/2.0,
	}
	player.AddComponent(currentSpriteRenderer)
	player.AddComponent(newKeyboardShooter(player, playerShotCooldown, audioDev))
	allowedRect := &engine.Rect{
		X:      currentSpriteRenderer.ScaledWidth / 2.0,
		Y:      4 * screenHeight / 5,
		Width:  screenWidth - currentSpriteRenderer.ScaledWidth,
		Height: screenHeight/5 - currentSpriteRenderer.ScaledHeight/2.0}
	player.AddComponent(engine.NewKeyboardMover(player, allowedRect, &delta, playerSpeed))
	player.AddComponent(engine.NewCollisionDetecter(player, true, "enemy"))
	player.AddComponent(engine.NewSoundPlayer(player, "sounds/explosion.wav", audioDev, engine.MsgCollision, "enemy"))
	player.Collisions = append(player.Collisions,
		engine.Circle{
			Center: &player.Position,
			Radius: math.Max(currentSpriteRenderer.ScaledWidth, currentSpriteRenderer.ScaledHeight) / 2,
		})

	return player
}
