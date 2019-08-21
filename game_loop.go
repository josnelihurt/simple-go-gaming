package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func inputHandler() (continueFlag bool) {

	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_ESCAPE] == 1 {
		return false
	}
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			logger <- fmt.Sprintf("exit requested:")
			return false
		}
	}

	return true
}
