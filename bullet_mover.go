package main

import (
	"github.com/josnelihurt/simple-go-gaming/engine"
	"github.com/veandco/go-sdl2/sdl"
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

func (context *bulletMover) OnUpdate() error {
	parent := context.parent
	parent.Position.Y -= bulletSpeed * delta

	if parent.Position.X > screenWidth || parent.Position.X < 0 ||
		parent.Position.Y < 0 {
		parent.Active = false
	}

	return nil
}

func (context *bulletMover) OnDraw(renderer *sdl.Renderer) error {
	return nil
}

func (context *bulletMover) OnCollision(other *engine.Element) error {
	if other.Tag == "bullet" {
		return nil
	}
	context.parent.Active = false
	context.onCollisionCallback()
	//logger <- fmt.Sprintf("bullet has crashed with %v :", other)
	return nil
}
