package engine

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/veandco/go-sdl2/sdl"
)

type Element struct {
	Position   Vector
	Rotation   float64
	Active     bool
	Collisions []Circle
	components []Component
	Tag        string
	Z          uint8
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
func insertSort(data []*Element, el *Element) []*Element {
	index := sort.Search(len(data), func(i int) bool { return data[i].Z > el.Z })
	data = append(data, &Element{})
	copy(data[index+1:], data[index:])
	data[index] = el
	return data
}
