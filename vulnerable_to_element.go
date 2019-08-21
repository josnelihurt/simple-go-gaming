package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/josnelihurt/simple-go-gaming/engine"
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
func (context *vulnerableToElement) onDraw(renderer *sdl.Renderer) error {
	return nil
}
func (context *vulnerableToElement) onUpdate() error {
	return nil
}
func (context *vulnerableToElement) onCollision(other *engine.Element) error {
	if contains(context.vulnerableTo, other.tag) {
		context.parent.active = false
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
