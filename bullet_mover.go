package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type bulletMover struct {
	parent *element
	speed  float64
}

func newBulletMover(parent *element, speed float64) *bulletMover {
	return &bulletMover{
		parent: parent,
		speed:  speed,
	}
}

func (context *bulletMover) onUpdate() error {
	parent := context.parent
	parent.position.x += bulletSpeed * math.Cos(parent.rotation)
	parent.position.y += bulletSpeed * math.Sin(parent.rotation)

	if parent.position.x > screenWidth || parent.position.x < 0 ||
		parent.position.y < 0 {
		parent.active = false
	}

	context.parent.collisions[0].center = parent.position

	return nil
}

func (context *bulletMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (context *bulletMover) onCollision(other *element) error {
	return nil
}
