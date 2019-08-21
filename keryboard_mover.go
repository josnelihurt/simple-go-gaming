package main

import (
	"github.com/josnelihurt/simple-go-gaming/engine"
	"github.com/veandco/go-sdl2/sdl"
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
		sr:     parent.GetComponent(&engine.SpriteRenderer{}).(*engine.SpriteRenderer),
	}
}
func (context *keyboardMover) OnUpdate() error {
	keys := sdl.GetKeyboardState()
	parent := context.parent

	if keys[sdl.SCANCODE_LEFT] == 1 {
		if parent.Position.X-(context.sr.ScaledWidth/2.0) > 0 {
			parent.Position.X -= context.speed * delta
		}
	}
	if keys[sdl.SCANCODE_RIGHT] == 1 {
		if parent.Position.X+(context.sr.ScaledWidth/2.0) < screenWidth {
			parent.Position.X += context.speed * delta
		}
	}
	if keys[sdl.SCANCODE_UP] == 1 {
		if parent.Position.Y+(context.sr.ScaledHeight/2.0) > 4*screenHeight/5.0 {
			parent.Position.Y -= context.speed * delta
		}
	}
	if keys[sdl.SCANCODE_DOWN] == 1 {
		if parent.Position.Y+(context.sr.ScaledHeight/2.0) < screenHeight {
			parent.Position.Y += context.speed * delta
		}
	}
	return nil
}
func (context *keyboardMover) OnDraw(renderer *sdl.Renderer) error {
	return nil
}
func (context *keyboardMover) OnCollision(other *engine.Element) error {
	return nil
}
func (context *keyboardMover) OnMessage(message *engine.Message) error {
	return nil
}
