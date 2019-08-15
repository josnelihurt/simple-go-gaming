package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed        = 2
	playerShotCooldown = time.Millisecond * 200
	playerScale        = 0.6
)

func newPlayer(renderer *sdl.Renderer) *element {
	player := &element{}
	player.active = true
	player.rotation = 180

	currentSpriteRenderer := newSpriteRenderer(player, renderer, "sprites/player.bmp", playerScale)
	player.position = vector{
		x: screenWidth / 2.0,
		y: screenHeight - 100, //currentSpriteRenderer.scaledHeight/2.0,
	}
	player.addCompoenent(currentSpriteRenderer)

	mover := newKeyboardMover(player, playerSpeed)
	player.addCompoenent(mover)

	shooter := newKeyboardShooter(player, playerShotCooldown)
	player.addCompoenent(shooter)

	return player
}
