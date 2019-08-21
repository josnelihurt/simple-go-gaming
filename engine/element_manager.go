package engine

import (
	"fmt"
	"math"

	"github.com/josnelihurt/simple-go-gaming/engine/util"
	"github.com/veandco/go-sdl2/sdl"
)

func collides(c1, c2 Circle) bool {
	dist := math.Sqrt(math.Pow(c2.Center.X-c1.Center.X, 2) + math.Pow(c2.Center.Y-c1.Center.Y, 2))
	return dist <= c1.Radius+c2.Radius
}

type ElementManager struct {
	elements []*Element
}

func NewElementManager() ElementManager {
	var elements []*Element
	return ElementManager{
		elements: elements,
	}
}

func (context *ElementManager) InsertSlice(newChunk []*Element) {
	for _, item := range newChunk {
		context.elements = insertSort(context.elements, item)
	}
}
func (context *ElementManager) GetElementsByTag(tag string) []*Element {
	var elements []*Element
	for _, currentElement := range context.elements {
		if currentElement.Tag == tag {
			elements = append(elements, currentElement)
		}
	}
	return elements
}

func (context *ElementManager) UpdateElements(renderer *sdl.Renderer) {
	for _, currentElement := range context.elements {
		if currentElement.Active {
			if err := currentElement.Update(); err != nil {
				util.Logger <- fmt.Sprintf("updating fail:%v", err)
			}
			if err := currentElement.Draw(renderer); err != nil {
				util.Logger <- fmt.Sprintf("drawing fail:%v", err)
			}
		}
	}
}

func (context *ElementManager) CheckColisions() error {
	elements := context.elements // I don't like it

	for i := 0; i < len(elements)-1; i++ {
		for j := 0; j < len(elements); j++ {
			for _, currenCollision1 := range elements[i].Collisions {
				for _, currenCollision2 := range elements[j].Collisions {
					if elements[i].Active && elements[j].Active && collides(currenCollision1, currenCollision2) {
						if err := elements[i].Collision(elements[j]); err != nil {
							return err
						}
						if err := elements[j].Collision(elements[i]); err != nil {
							return err
						}
					}
				}
			}
		}
	}
	return nil
}