package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed        = 5
	playerShotCooldown = time.Millisecond * 200
	playerScale        = 0.7
)

func newPlayer(renderer *sdl.Renderer) *element {
	player := &element{}
	player.z = 10
	player.active = true
	player.rotation = 180

	currentSpriteRenderer := newSpriteRenderer(player, renderer, "sprites/player.png", playerScale)
	player.position = vector{
		x: screenWidth / 2.0,
		y: screenHeight - currentSpriteRenderer.scaledHeight/2.0,
	}
	player.addCompoenent(currentSpriteRenderer)

	mover := newKeyboardMover(player, playerSpeed)
	player.addCompoenent(mover)

	shooter := newKeyboardShooter(player, playerShotCooldown)
	player.addCompoenent(shooter)

	return player
}
