package main

import (
	"math"
	"time"

	"github.com/josnelihurt/simple-go-gaming/engine"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed        = 5
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
	player.AddComponent(newKeyboardMover(player, playerSpeed))
	player.AddComponent(newKeyboardShooter(player, playerShotCooldown, audioDev))
	player.AddComponent(engine.NewCollisionDetecter(player, "enemy"))
	player.Collisions = append(player.Collisions,
		engine.Circle{
			Center: &player.Position,
			Radius: math.Max(currentSpriteRenderer.ScaledWidth, currentSpriteRenderer.ScaledHeight) / 2,
		})

	return player
}
