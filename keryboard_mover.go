package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/josnelihurt/simple-go-gaming/engine"
)

type keyboardMover struct {
	parent *engine.Element
	speed  float64
	sr     *engine.SpriteRenderer
}

func newKeyboardMover(parent *engine.Element, speed float64) *keyboardMover {
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
			parent.position.x -= context.speed * delta
		}
	}
	if keys[sdl.SCANCODE_RIGHT] == 1 {
		if parent.position.x+(context.sr.scaledWidth/2.0) < screenWidth {
			parent.position.x += context.speed * delta
		}
	}
	if keys[sdl.SCANCODE_UP] == 1 {
		if parent.position.y+(context.sr.scaledHeight/2.0) > 4*screenHeight/5.0 {
			parent.position.y -= context.speed * delta
		}
	}
	if keys[sdl.SCANCODE_DOWN] == 1 {
		if parent.position.y+(context.sr.scaledHeight/2.0) < screenHeight {
			parent.position.y += context.speed * delta
		}
	}
	return nil
}
func (context *keyboardMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}
func (context *keyboardMover) onCollision(other *engine.Element) error {
	return nil
}
