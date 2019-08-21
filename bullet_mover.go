package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/josnelihurt/simple-go-gaming/engine"
)

type bulletMover struct {
	parent              *engine.Element
	speed               float64
	onCollisionCallback func()
}

func newBulletMover(parent *engine.Element, speed float64, onCollisionCallback func()) *bulletMover {
	return &bulletMover{
		parent:              parent,
		speed:               speed,
		onCollisionCallback: onCollisionCallback,
	}
}

func (context *bulletMover) onUpdate() error {
	parent := context.parent
	parent.position.y -= bulletSpeed * delta

	if parent.position.x > screenWidth || parent.position.x < 0 ||
		parent.position.y < 0 {
		parent.active = false
	}

	return nil
}

func (context *bulletMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (context *bulletMover) onCollision(other *engine.Element) error {
	if other.tag == "bullet" {
		return nil
	}
	context.parent.active = false
	context.onCollisionCallback()
	//logger <- fmt.Sprintf("bullet has crashed with %v :", other)
	return nil
}
