package main

import (
	"github.com/josnelihurt/simple-go-gaming/engine"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	scoreX = 220
)

func newScore() *engine.Element {
	context := &engine.Element{Active: true, Tag: "score", Z: 99}
	textRenderer := engine.NewTextRenderer(
		&engine.Vector{X: (screenWidth - scoreX), Y: upperTextY},
		defaultFontSize,
		sdl.Color{R: 255, G: 255, B: 255})
	context.AddComponent(textRenderer)
	context.AddComponent(newScoreCounter(context, textRenderer))
	return context
}
