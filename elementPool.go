package main

import "sort"

type elementPool struct {
	elements []*element
}

func newElementPool() elementPool {
	var elements []*element
	return elementPool{
		elements: elements,
	}
}
func insertSort(data []*element, el *element) []*element {
	index := sort.Search(len(data), func(i int) bool { return data[i].z > el.z })
	data = append(data, &element{})
	copy(data[index+1:], data[index:])
	data[index] = el
	return data
}
func (context *elementPool) insertElement(newElement *element) {
	context.elements = insertSort(context.elements, newElement)
}

func (context *elementPool) insertSlice(newChunk []*element) {
	for _, item := range newChunk {
		context.insertElement(item)
	}
}
