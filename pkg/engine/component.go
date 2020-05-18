package engine

import "github.com/veandco/go-sdl2/sdl"

// Component represents any part of an element
type Component interface {
	// OnUpdate
	OnUpdate() error
	OnDraw(renderer *sdl.Renderer) error
	OnCollision(other *Element) error
	OnMessage(message *Message) error
}
