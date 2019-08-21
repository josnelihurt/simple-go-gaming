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
	score := &engine.Element{Active: true, Tag: "score", Z: 99}
	score.AddComponent(engine.NewTextRenderer(
		&engine.Vector{X: (screenWidth - 70), Y: 15},
		scoreFontSize,
		sdl.Color{R: 255, G: 255, B: 255}))
	score.AddComponent(newScoreCounter(score))
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

func (context *scoreCounter) increase() {
	context.currentValue++
}
