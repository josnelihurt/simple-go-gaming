package main

import (
	"github.com/josnelihurt/simple-go-gaming/engine"
	"github.com/veandco/go-sdl2/sdl"
)

type playerStarter struct {
	parent          *engine.Element
	defaultPosition engine.Vector
}

func newPlayerStarter(parent *engine.Element, defaultPosition engine.Vector) *playerStarter {
	return &playerStarter{
		parent:          parent,
		defaultPosition: defaultPosition,
	}
}
func (context *playerStarter) OnDraw(renderer *sdl.Renderer) error     { return nil }
func (context *playerStarter) OnCollision(other *engine.Element) error { return nil }
func (context *playerStarter) OnUpdate() error {
	return nil
}
func (context *playerStarter) OnMessage(message *engine.Message) error {
	if message.Code == engine.MsgCollision {
		context.parent.Position = context.defaultPosition
	}
	return nil
}