package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

// CollisionDetecter is a component that reacta to a message with given conditions
type CollisionDetecter struct {
	parent                      *Element
	vulnerableTo                []string
	sendMessagesToOtherElements bool
}
type localType = CollisionDetecter

// NewCollisionDetecter creates a CollisionDetecter instance with given parameters
func NewCollisionDetecter(parent *Element, sendMessagesToOtherElements bool, elementsActives ...string) *CollisionDetecter {
	return &CollisionDetecter{
		parent:                      parent,
		vulnerableTo:                elementsActives,
		sendMessagesToOtherElements: sendMessagesToOtherElements,
	}
}

func (context *CollisionDetecter) OnDraw(renderer *sdl.Renderer) error {
	return nil
}
func (context *CollisionDetecter) OnUpdate() error {
	return nil
}
func (context *CollisionDetecter) OnCollision(other *Element) error {
	if isSingleAndEmpty(context.vulnerableTo) || contains(context.vulnerableTo, other.Tag) {
		context.parent.BroadcastMessageToComponents(&Message{
			Sender:              context.parent,
			Code:                MsgCollision,
			RelatedTo:           []*Element{other},
			SendToOtherElements: context.sendMessagesToOtherElements,
			Data:                "newCollision",
		})
	}
	return nil
}
func (context *CollisionDetecter) OnMessage(message *Message) error {
	return nil
}
