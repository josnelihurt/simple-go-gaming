package engine

import (
	"fmt"
	"sort"

	"github.com/josnelihurt/simple-go-gaming/engine/util"
	"github.com/veandco/go-sdl2/sdl"
)

type ElementPool struct {
	elements []*Element
}

func NewElementPool() ElementPool {
	var elements []*Element
	return ElementPool{
		elements: elements,
	}
}
func insertSort(data []*Element, el *Element) []*Element {
	index := sort.Search(len(data), func(i int) bool { return data[i].Z > el.Z })
	data = append(data, &Element{})
	copy(data[index+1:], data[index:])
	data[index] = el
	return data
}
func (context *ElementPool) InsertElement(newElement *Element) {
	context.elements = insertSort(context.elements, newElement)
}

func (context *ElementPool) InsertSlice(newChunk []*Element) {
	for _, item := range newChunk {
		context.InsertElement(item)
	}
}
func (context *ElementPool) GetElementsByTag(tag string) []*Element {
	var elements []*Element
	for _, currentElement := range context.elements {
		if currentElement.Tag == tag {
			elements = append(elements, currentElement)
		}
	}
	return elements
}

func (context *ElementPool) UpdateElements(renderer *sdl.Renderer) {
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
