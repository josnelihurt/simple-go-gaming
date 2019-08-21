package main

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/josnelihurt/simple-go-gaming/engine"
)

type enemyMover struct {
	active bool
	parent *engine.Element
}

func newEnemyMover(parent *engine.Element) *enemyMover {
	return &enemyMover{
		active: false,
		parent: parent,
	}
}
func (context *enemyMover) onUpdate() error {
	if context.parent.position.y >= screenHeight {
		context.parent.active = false
	}
	if context.active {
		context.parent.position.y += enemySpeed * delta
	}
	return nil
}
func (context *enemyMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}
func (context *enemyMover) onCollision(other *engine.Element) error {
	return nil
}
