package main

import (
	"fmt"
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type keyboardMover struct {
	parent *element
	speed  float64
	sr     *spriteRenderer
}

func newKeyboardMover(parent *element, speed float64) *keyboardMover {
	return &keyboardMover{
		parent: parent,
		speed:  speed,
		sr:     parent.getComponent(&spriteRenderer{}).(*spriteRenderer),
	}
}
func (context *keyboardMover) onUpdate() error {
	keys := sdl.GetKeyboardState()
	parent := context.parent

	if keys[sdl.SCANCODE_LEFT] == 1 {
		if parent.position.x-(context.sr.scaledWidth/2.0) > 0 {
			parent.position.x -= context.speed
		} else {
			fmt.Println("limit l:", parent.position)
		}
	}
	if keys[sdl.SCANCODE_RIGHT] == 1 {
		if parent.position.x-(context.sr.scaledWidth/2.0) < screenWidth {
			parent.position.x += context.speed
		} else {
			fmt.Println("limit r:", parent.position)
		}
	}
	return nil
}
func (context *keyboardMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

type keyboardShooter struct {
	parent   *element
	cooldown time.Duration
	lastShot time.Time
}

func newKeyboardShooter(parent *element, cooldown time.Duration) *keyboardShooter {
	return &keyboardShooter{
		parent:   parent,
		cooldown: cooldown,
	}
}
func (context *keyboardShooter) onUpdate() error {
	keys := sdl.GetKeyboardState()

	parent := context.parent

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(context.lastShot) > context.cooldown {
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
func (context *keyboardShooter) shoot(x, y float64) {
	if currentBullet, ok := bulletFromPool(); ok {
		currentBullet.active = true
		currentBullet.position.x = x
		currentBullet.position.y = y
		currentBullet.rotation = 270 * (math.Pi / 180)
	}
}
