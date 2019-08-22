package engine

import "github.com/veandco/go-sdl2/sdl"

type ComponentDestroyerOnMessage struct {
	parent       *Element
	relatedToTag []string
	messageCode  int
}

func NewComponentDestroyerOnMessage(parent *Element, messageCode int, relatedToTag ...string) *ComponentDestroyerOnMessage {
	return &ComponentDestroyerOnMessage{
		parent:       parent,
		relatedToTag: relatedToTag,
		messageCode:  messageCode,
	}
}

func (context *ComponentDestroyerOnMessage) OnUpdate() error                     { return nil }
func (context *ComponentDestroyerOnMessage) OnDraw(renderer *sdl.Renderer) error { return nil }
func (context *ComponentDestroyerOnMessage) OnCollision(other *Element) error    { return nil }
func (context *ComponentDestroyerOnMessage) OnMessage(message *Message) error {
	if context.messageCode == message.Code {
		if isSingleAndEmpty(context.relatedToTag) || contains(context.relatedToTag, message.RelatedTo[0].Tag) {
			context.parent.Active = false
		}
	}
	return nil
}
