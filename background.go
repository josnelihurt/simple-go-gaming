package main

import (
	"github.com/josnelihurt/simple-go-gaming/engine"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	backgroundSpeed = 0.5
)

func newBackground(renderer *sdl.Renderer) *engine.Element {
	background := &engine.Element{
		Active: true,
	}
	background.AddComponent(newBackgroundMover(renderer))
	return background
}
