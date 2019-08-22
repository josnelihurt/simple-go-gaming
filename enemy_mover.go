package main

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/josnelihurt/simple-go-gaming/engine"
	"github.com/josnelihurt/simple-go-gaming/engine/util"
)

type enemyMover struct {
	active bool
	parent *engine.Element
}

func newEnemyMover(parent *engine.Element) *enemyMover {
	return &enemyMover{
		active: false,
		parent: parent,
	}
}
func (context *enemyMover) OnUpdate() error {
	if context.parent.Position.Y >= screenHeight {
		util.Logger <- "enemy has finished his race"
		context.parent.Active = false
		context.parent.BroadcastMessageToComponents(&engine.Message{
			Code:                msgHitPlayer,
			Sender:              context.parent,
			SendToOtherElements: true,
			Data:                "HitPlayer",
			RelatedTo:           []*engine.Element{context.parent},
		})
	}
	if context.active {
		context.parent.Position.Y += enemySpeed * delta
	}
	return nil
}
func (context *enemyMover) OnDraw(renderer *sdl.Renderer) error     { return nil }
func (context *enemyMover) OnCollision(other *engine.Element) error { return nil }
func (context *enemyMover) OnMessage(message *engine.Message) error { return nil }
