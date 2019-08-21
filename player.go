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
	player := &element{}
	player.z = 10
	player.active = true
	player.rotation = 180
	player.tag = "player"

	currentSpriteRenderer := newSpriteRenderer(player, renderer, "sprites/player.png", playerScale)
	player.position = vector{
		x: screenWidth / 2.0,
		y: screenHeight - currentSpriteRenderer.scaledHeight/2.0,
	}
	player.addComponent(currentSpriteRenderer)
	player.addComponent(newKeyboardMover(player, playerSpeed))
	player.addComponent(newKeyboardShooter(player, playerShotCooldown, audioDev))
	player.addComponent(newVulnerableToElement(player, func(*engine.Element) {
		player.active = true
		explosion := newSoundPlayer("sounds/explosion.wav", audioDev)
		explosion.play()
	}, "enemy"))
	player.collisions = append(player.collisions,
		circle{
			center: &player.position,
			radius: math.Max(currentSpriteRenderer.scaledWidth, currentSpriteRenderer.scaledHeight) / 2,
		})
	//player.addComponent(newPlayerLife(player))

	return player
}

func init() {

}
