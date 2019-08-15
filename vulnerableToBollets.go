package main

import (
	"log"
	"reflect"

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
	log.Println(reflect.TypeOf(other), reflect.TypeOf(&bullet{}))
	//if reflect.TypeOf(other) == reflect.TypeOf(&bullet{}) {
	context.parent.active = false
	//}
	return nil
}
