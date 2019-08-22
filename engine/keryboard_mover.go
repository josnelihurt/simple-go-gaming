package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

type KeyboardMover struct {
	parent      *Element
	allowedRect *Rect
	delta       *float64
	speed       float64
}

func NewKeyboardMover(parent *Element, allowedRect *Rect, delta *float64, speed float64) *KeyboardMover {
	return &KeyboardMover{
		parent:      parent,
		speed:       speed,
		allowedRect: allowedRect,
		delta:       delta,
	}
}
func (context *KeyboardMover) getCurrentSpeed() float64 {
	return context.speed * (*context.delta)
}
func (context *KeyboardMover) OnUpdate() error {
	keys := sdl.GetKeyboardState()
	parent := context.parent

	if keys[sdl.SCANCODE_LEFT] == 1 {
		if parent.Position.X > context.allowedRect.X {
			parent.Position.X -= context.getCurrentSpeed()
		}
	}
	if keys[sdl.SCANCODE_RIGHT] == 1 {
		if parent.Position.X < context.allowedRect.Width+context.allowedRect.X {
			parent.Position.X += context.getCurrentSpeed()
		}
	}
	if keys[sdl.SCANCODE_UP] == 1 {
		if parent.Position.Y > context.allowedRect.Y {
			parent.Position.Y -= context.getCurrentSpeed()
		}
	}
	if keys[sdl.SCANCODE_DOWN] == 1 {
		if parent.Position.Y < context.allowedRect.Height+context.allowedRect.Y {
			parent.Position.Y += context.getCurrentSpeed()
		}
	}
	return nil
}
func (context *KeyboardMover) OnDraw(renderer *sdl.Renderer) error {
	return nil
}
func (context *KeyboardMover) OnCollision(other *Element) error {
	return nil
}
func (context *KeyboardMover) OnMessage(message *Message) error {
	return nil
}
