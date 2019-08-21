package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed        = 5
	playerShotCooldown = time.Millisecond * 200
	playerScale        = 0.7
)

func newPlayer(renderer *sdl.Renderer, audioDev sdl.AudioDeviceID) *element {
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
	player.addCompoenent(currentSpriteRenderer)
	player.addCompoenent(newKeyboardMover(player, playerSpeed))
	player.addCompoenent(newKeyboardShooter(player, playerShotCooldown, audioDev))
	player.addCompoenent(newVulnerableToElement(player, func(*element) {
		player.active = true
		explosion := newSoundPlayer("sounds/explosion.wav", audioDev)
		explosion.play()
	}, "enemy"))
	player.collisions = append(player.collisions,
		circle{
			center: &player.position,
			radius: math.Max(currentSpriteRenderer.scaledWidth, currentSpriteRenderer.scaledHeight) / 2,
		})
	//player.addCompoenent(newPlayerLife(player))

	return player
}
