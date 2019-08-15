package main

import (
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
			//logger <- fmt.Sprintln("limit l:", parent.position)
		}
	}
	if keys[sdl.SCANCODE_RIGHT] == 1 {
		if parent.position.x+(context.sr.scaledWidth/2.0) < screenWidth {
			parent.position.x += context.speed
		} else {
			//logger <- fmt.Sprintln("limit r:", parent.position)
		}
	}
	return nil
}
func (context *keyboardMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}
func (context *keyboardMover) onCollision(other *element) error {
	return nil
}
