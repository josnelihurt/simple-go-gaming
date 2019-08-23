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
		currentValue: 30,
		textRenderer: textRenderer,
	}
}
func (context *playerLifeCounter) OnUpdate() error {
	context.textRenderer.SetNewText(fmt.Sprintf("life:%03d", context.currentValue))
	if context.currentValue == 0 {
		context.parent.BroadcastMessage(&engine.Message{
			Code:                msgPlayerDead,
			Sender:              context.parent,
			Data:                "playerDead",
			SendToOtherElements: true,
		})
	}
	return nil
}
func (context *playerLifeCounter) OnMessage(message *engine.Message) error {
	util.Logger <- fmt.Sprintf("life hit:%v", message)
	if (message.Code == engine.MsgCollision &&
		message.Sender.Tag == tagEnemy &&
		len(message.RelatedTo) > 0 &&
		message.RelatedTo[0].Tag == tagPlayer) || message.Code == msgHitPlayer {
		context.currentValue--
	}

	return nil
}

func (context *playerLifeCounter) OnDraw(renderer *sdl.Renderer) error     { return nil }
func (context *playerLifeCounter) OnCollision(other *engine.Element) error { return nil }
