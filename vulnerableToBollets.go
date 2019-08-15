package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type vulnerableToBullets struct {
	parent *element
}

func newVulnerableToBullets(parent *element) *vulnerableToBullets {
	return &vulnerableToBullets{parent: parent}
}
func (context *vulnerableToBullets) onDraw(renderer *sdl.Renderer) error {
	return nil
}
func (context *vulnerableToBullets) onUpdate() error {
	return nil
}
func (context *vulnerableToBullets) onCollision(other *element) error {
	if other.tag == "bullet" {
		context.parent.active = false
		logger <- fmt.Sprintf("%v was hit by a bullet %v", context.parent, other)
	}
	return nil
}
