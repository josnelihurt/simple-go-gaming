package main

import (
	"fmt"

	"github.com/josnelihurt/simple-go-gaming/engine"
	"github.com/veandco/go-sdl2/sdl"
)

type scoreCounter struct {
	parent       *engine.Element
	currentValue int
}

func newScoreCounter(parent *engine.Element) *scoreCounter {
	return &scoreCounter{
		parent:       parent,
		currentValue: 0,
	}
}

func (context *scoreCounter) OnUpdate() error {
	context.parent.GetComponent(&engine.TextRenderer{}).(*engine.TextRenderer).SetNewText(fmt.Sprintf("%03d", context.currentValue))
	return nil
}
func (context *scoreCounter) OnDraw(renderer *sdl.Renderer) error {
	return nil
}
func (context *scoreCounter) OnCollision(other *engine.Element) error {
	return nil
}
func (context *scoreCounter) OnMessage(message *engine.Message) error {
	return nil
}

func (context *scoreCounter) increase() {
	context.currentValue++
}
