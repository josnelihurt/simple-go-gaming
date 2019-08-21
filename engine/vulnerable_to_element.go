package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

type VulnerableToElement struct {
	parent        *Element
	vulnerableTo  []string
	onHitCallback func(*Element)
}

func NewVulnerableToElement(parent *Element, onHitCallback func(*Element), vulnerableTo ...string) *VulnerableToElement {
	return &VulnerableToElement{
		parent:        parent,
		vulnerableTo:  vulnerableTo,
		onHitCallback: onHitCallback,
	}
}
func (context *VulnerableToElement) OnDraw(renderer *sdl.Renderer) error {
	return nil
}
func (context *VulnerableToElement) OnUpdate() error {
	return nil
}
func (context *VulnerableToElement) OnCollision(other *Element) error {
	if contains(context.vulnerableTo, other.Tag) {
		context.parent.Active = false // Replace this for a message
		context.onHitCallback(other)
	}
	return nil
}
func (context *VulnerableToElement) OnMessage(message *Message) error {
	return nil
}
func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
