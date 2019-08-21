package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

type CollisionDetecter struct {
	parent       *Element
	vulnerableTo []string
}

func NewCollisionDetecter(parent *Element, elementsActives ...string) *CollisionDetecter {
	return &CollisionDetecter{
		parent:       parent,
		vulnerableTo: elementsActives,
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
			Sender:    context.parent,
			Code:      MsgCollision,
			RelatedTo: []*Element{other},
		})
	}
	return nil
}
func (context *CollisionDetecter) OnMessage(message *Message) error {
	return nil
}
func isSingleAndEmpty(a []string) bool {
	return len(a) == 1 && a[0] == ""
}
func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
