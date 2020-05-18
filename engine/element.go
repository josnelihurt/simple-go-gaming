package engine

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/josnelihurt/simple-go-gaming/engine/util"
	"github.com/veandco/go-sdl2/sdl"
)

//Element represents a game concept
type Element struct {
	components     []Component
	messageEmitter []func(*Message) error
	Active         bool
	Collisions     []Circle
	Position       Vector
	Rotation       float64
	Tag            string
	Z              uint8
}

// RegisterEmmitterCallback insert a function callback in the observers list to call when a message is rised
func (context *Element) RegisterEmitterrCallback(callback func(*Message) error) {
	context.messageEmitter = append(context.messageEmitter, callback)
}

// BroadcastMessage sends a message to other components
func (context *Element) BroadcastMessage(message *Message) error {
	for _, currentComponent := range context.components {
		if err := currentComponent.OnMessage(message); err != nil {
			return err
		}
	}
	if len(context.messageEmitter) > 0 && message.SendToOtherElements {
		for _, currentCallback := range context.messageEmitter {
			if err := currentCallback(message); err != nil {
				util.Logger <- fmt.Sprintf("Error in msg %v", err)
			}
		}
	}

	return nil
}

//Draw calls the OnDraw method on all components
func (context *Element) Draw(renderer *sdl.Renderer) error {
	for _, currentComponent := range context.components {
		if err := currentComponent.OnDraw(renderer); err != nil {
			return err
		}
	}
	return nil
}

// Update calls the OnUpdate method on all components
func (context *Element) Update() error {
	for _, currentComponent := range context.components {
		if err := currentComponent.OnUpdate(); err != nil {
			return err
		}
	}
	return nil
}

//AddComponent inserts a component in the internal data structure
func (context *Element) AddComponent(new Component) {
	context.components = append(context.components, new)
}

//GetFirstComponent obtains the firts component by its type
func (context *Element) GetFirstComponent(withType Component) Component {
	componentType := reflect.TypeOf(withType)
	for _, currentComponent := range context.components {
		if reflect.TypeOf(currentComponent) == componentType {
			return currentComponent
		}
	}

	panic(fmt.Sprintf("there is not such component %v", componentType))
}

//Collision calls OnCollision method in all components
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
