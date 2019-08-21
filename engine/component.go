package engine

import "github.com/veandco/go-sdl2/sdl"

type Component interface {
	OnUpdate() error
	OnDraw(renderer *sdl.Renderer) error
	OnCollision(other *Element) error
	OnMessage(message *Message) error
}
