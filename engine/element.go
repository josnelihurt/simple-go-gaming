package engine

import (
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Vector struct {
	X, Y float64
}

type Component interface {
	OnUpdate() error
	OnDraw(renderer *sdl.Renderer) error
	OnCollision(other *Element) error
}

type Element struct {
	Position   Vector
	Rotation   float64
	Active     bool
	Collisions []Circle
	components []Component
	Tag        string
	Z          uint8
}

func (context *Element) RunOnAllComponents(callback func(Component, *sdl.Renderer) error, renderer *sdl.Renderer) error {
	for _, currentComponent := range context.components {
		if err := callback(currentComponent, renderer); err != nil {
			return err
		}
	}
	return nil
}

func (context *Element) Draw(renderer *sdl.Renderer) error {
	for _, currentComponent := range context.components {
		if err := currentComponent.OnDraw(renderer); err != nil {
			return err
		}
	}
	return nil
}
func (context *Element) Update() error {
	for _, currentComponent := range context.components {
		if err := currentComponent.OnUpdate(); err != nil {
			return err
		}
	}
	return nil
}

func (context *Element) AddComponent(new Component) {
	for _, existing := range context.components {
		if reflect.TypeOf(existing) == reflect.TypeOf(new) {
			panic(fmt.Sprintf("attempt to add new component with existing type %v",
				reflect.TypeOf(new)))
		}
	}
	context.components = append(context.components, new)
}

func (context *Element) GetComponent(withType Component) Component {
	componentType := reflect.TypeOf(withType)
	for _, currentComponent := range context.components {
		if reflect.TypeOf(currentComponent) == componentType {
			return currentComponent
		}
	}

	panic(fmt.Sprintf("there is not such component %v", componentType))
}

func (context *Element) Collision(other *Element) error {
	for _, currentComponent := range context.components {
		if err := currentComponent.OnCollision(other); err != nil {
			return err
		}
	}
	return nil
}

func LoadFont(fontSize int) (font *ttf.Font, err error) {
	font, err = ttf.OpenFont("fonts/Starjout.ttf", fontSize)
	if err != nil {
		return nil, fmt.Errorf("initializing font:%v", err)
	}
	return font, nil
}
