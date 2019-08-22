package main

import (
	"fmt"

	"github.com/josnelihurt/simple-go-gaming/engine"
	"github.com/josnelihurt/simple-go-gaming/engine/util"
	"github.com/veandco/go-sdl2/sdl"
)

type playerLifeCounter struct {
	parent       *engine.Element
	currentValue int
	textRenderer *engine.TextRenderer
}

func newPlayerLifeCounter(parent *engine.Element, textRenderer *engine.TextRenderer) *playerLifeCounter {
	return &playerLifeCounter{
		parent:       parent,
		currentValue: 3,
		textRenderer: textRenderer,
	}
}
func (context *playerLifeCounter) OnDraw(renderer *sdl.Renderer) error     { return nil }
func (context *playerLifeCounter) OnCollision(other *engine.Element) error { return nil }
func (context *playerLifeCounter) OnUpdate() error {
	context.textRenderer.SetNewText(fmt.Sprintf("life:%03d", context.currentValue))
	if context.currentValue == 0 {
		context.parent.BroadcastMessageToComponents(&engine.Message{
			Code:                msgPlayerDead,
			Sender:              context.parent,
			Data:                "playerDead",
			SendToOtherElements: true,
		})
	}
	return nil
}
func (context *playerLifeCounter) OnMessage(message *engine.Message) error {
	util.Logger <- fmt.Sprintf("msg:%v", message)
	if message.Code == engine.MsgCollision || message.Code == msgHitPlayer {
		context.currentValue--
	}

	return nil
}
