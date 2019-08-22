package main

import (
	"fmt"

	"github.com/josnelihurt/simple-go-gaming/engine"
	"github.com/veandco/go-sdl2/sdl"
)

type scoreCounter struct {
	parent       *engine.Element
	currentValue int
	textRenderer *engine.TextRenderer
}

func newScoreCounter(parent *engine.Element, textRenderer *engine.TextRenderer) *scoreCounter {
	return &scoreCounter{
		parent:       parent,
		currentValue: 0,
		textRenderer: textRenderer,
	}
}
func (context *scoreCounter) OnDraw(renderer *sdl.Renderer) error     { return nil }
func (context *scoreCounter) OnCollision(other *engine.Element) error { return nil }

func (context *scoreCounter) OnUpdate() error {
	context.textRenderer.SetNewText(fmt.Sprintf("score:%03d", context.currentValue))
	return nil
}
func (context *scoreCounter) OnMessage(message *engine.Message) error {
	if message.Code == engine.MsgCollision && message.Sender.Tag == "enemy" && message.RelatedTo[0].Tag == "bullet" {
		context.currentValue++
	}
	return nil
}
