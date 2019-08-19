package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func inputHandler() (continueFlag bool) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			logger <- fmt.Sprintf("exit requested:")
			return false
		}
	}
	return true
}
func gameLoop(renderer *sdl.Renderer) (continueFlag bool) {
	return true
}
