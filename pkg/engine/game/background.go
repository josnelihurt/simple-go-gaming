package game

import (
	"github.com/josnelihurt/simple-go-gaming/pkg/engine"
)

const (
	backgroundSpeed = 0.9
)

func newBackground(components *engine.SDLComponents) *engine.Element {
	background := &engine.Element{
		Active: true,
	}
	background.AddComponent(newBackgroundMover(components.Renderer))
	return background
}
