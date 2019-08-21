package main

import (
	"github.com/josnelihurt/simple-go-gaming/engine"
	"github.com/veandco/go-sdl2/sdl"
)

type vulnerableToElement struct {
	parent        *engine.Element
	vulnerableTo  []string
	onHitCallback func(*engine.Element)
}

func newVulnerableToElement(parent *engine.Element, onHitCallback func(*engine.Element), vulnerableTo ...string) *vulnerableToElement {
	return &vulnerableToElement{
		parent:        parent,
		vulnerableTo:  vulnerableTo,
		onHitCallback: onHitCallback,
	}
}
func (context *vulnerableToElement) OnDraw(renderer *sdl.Renderer) error {
	return nil
}
func (context *vulnerableToElement) OnUpdate() error {
	return nil
}
func (context *vulnerableToElement) OnCollision(other *engine.Element) error {
	if contains(context.vulnerableTo, other.Tag) {
		context.parent.Active = false
		context.onHitCallback(other)
	}
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
