package main

import (
	"github.com/josnelihurt/simple-go-gaming/engine"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	tagLevel = "level"
)

func newLevel() (context *engine.Element) {
	context.Tag = tagLevel
	context.Z = 99
	position := engine.Vector{X: 0, Y: upperTextY}
	textRenderer := engine.NewTextRenderer(&position, defaultFontSize, sdl.Color{R: 255, G: 255, B: 255})
	textRenderer.SetNewText("Level ")
	//context.Position =
	return context
}
