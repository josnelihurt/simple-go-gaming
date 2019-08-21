package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type score struct {
	textRenderer
	currentValue int
}

func newScore() *score {
	score := newTextRenderer(
		&vector{x: (screenWidth - 30), y: 10},
		scoreFontSize,
		sdl.Color{R: 255, G: 255, B: 255})
	return score
}

func (context *score) increase() {
	context.currentValue++
	context.currenValue = fmt.Sprintf("%03d", context.currentValue)
}
