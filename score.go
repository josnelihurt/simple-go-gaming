package main

import (
	"fmt"

	"github.com/josnelihurt/simple-go-gaming/engine"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	scoreFontSize = 27
)

func newScore() *engine.Element {
	score := &element{active: true, tag: "score", z: 99}
	score.addComponent(newTextRenderer(
		&vector{x: (screenWidth - 70), y: 15},
		scoreFontSize,
		sdl.Color{R: 255, G: 255, B: 255}))
	score.addComponent(newScoreCounter(score))
	return score
}

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

func (context *scoreCounter) onUpdate() error {
	context.parent.getComponent(&textRenderer{}).(*textRenderer).newValue = fmt.Sprintf("%03d", context.currentValue)
	return nil
}
func (context *scoreCounter) onDraw(renderer *sdl.Renderer) error {
	return nil
}
func (context *scoreCounter) onCollision(other *engine.Element) error {
	return nil
}

func (context *scoreCounter) increase() {
	context.currentValue++
}
