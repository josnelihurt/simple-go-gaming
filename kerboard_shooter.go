package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type keyboardShooter struct {
	parent   *element
	cooldown time.Duration
	lastShot time.Time
	player   soundPlayer
}

func newKeyboardShooter(parent *element, cooldown time.Duration, audioDev sdl.AudioDeviceID) *keyboardShooter {

	return &keyboardShooter{
		parent:   parent,
		cooldown: cooldown,
		player:   newSoundPlayer("sounds/NFF-laser.wav", audioDev),
	}
}
func (context *keyboardShooter) onUpdate() error {
	keys := sdl.GetKeyboardState()

	parent := context.parent

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(context.lastShot) > context.cooldown {
			context.player.play()
			sprite := parent.getComponent(&spriteRenderer{}).(*spriteRenderer)
			context.shoot(parent.position.x+23, parent.position.y-sprite.scaledHeight/2)
			context.shoot(parent.position.x-23, parent.position.y-sprite.scaledHeight/2)
			context.lastShot = time.Now()
		}
	}
	return nil
}
func (context *keyboardShooter) onDraw(enderer *sdl.Renderer) error {
	return nil
}

func (context *keyboardShooter) onCollision(other *element) error {
	return nil
}
func (context *keyboardShooter) shoot(x, y float64) {
	if currentBullet, ok := bulletFromPool(); ok {
		currentBullet.active = true
		currentBullet.position.x = x
		currentBullet.position.y = y
	}
}
