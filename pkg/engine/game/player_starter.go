package game

import (
	"github.com/josnelihurt/simple-go-gaming/pkg/engine"
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
func (context *playerStarter) OnUpdate() error {
	return nil
}
func (context *playerStarter) OnMessage(message *engine.Message) error {
	if message.Code == engine.MsgCollision &&
		message.Sender.Tag == tagEnemy &&
		len(message.RelatedTo) > 0 &&
		message.RelatedTo[0].Tag == tagPlayer {
		context.parent.Position = context.defaultPosition
	}
	return nil
}

func (context *playerStarter) OnDraw(renderer *sdl.Renderer) error     { return nil }
func (context *playerStarter) OnCollision(other *engine.Element) error { return nil }
