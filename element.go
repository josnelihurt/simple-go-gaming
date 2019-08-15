package main

import (
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

type vector struct {
	x, y float64
}

type component interface {
	onUpdate() error
	onDraw(renderer *sdl.Renderer) error
	onCollision(other *element) error
}

type element struct {
	position   vector
	rotation   float64
	active     bool
	collisions []circle
	components []component
}

func (context *element) runOnAllComponents(callback func(component, *sdl.Renderer) error, renderer *sdl.Renderer) error {
	for _, currentComponent := range context.components {
		if err := callback(currentComponent, renderer); err != nil {
			return err
		}
	}
	return nil
}

func (context *element) draw(renderer *sdl.Renderer) error {
	for _, currentComponent := range context.components {
		if err := currentComponent.onDraw(renderer); err != nil {
			return err
		}
	}
	return nil
}
func (context *element) update() error {
	for _, currentComponent := range context.components {
		if err := currentComponent.onUpdate(); err != nil {
			return err
		}
	}
	return nil
}

func (context *element) addCompoenent(new component) {
	for _, existing := range context.components {
		if reflect.TypeOf(existing) == reflect.TypeOf(new) {
			panic(fmt.Sprintf("attempt to add new component with existing type %v",
				reflect.TypeOf(new)))
		}
	}
	context.components = append(context.components, new)
}

func (context *element) getComponent(withType component) component {
	componentType := reflect.TypeOf(withType)
	for _, currentComponent := range context.components {
		if reflect.TypeOf(currentComponent) == componentType {
			return currentComponent
		}
	}

	panic(fmt.Sprintf("there is not such component %v", componentType))
}

func (context *element) collision(other *element) error {
	for _, currentComponent := range context.components {
		if err := currentComponent.onCollision(other); err != nil {
			return err
		}
	}
	return nil
}

var elements []*element
