package util

import (
	"fmt"

	"github.com/veandco/go-sdl2/ttf"
)

// LoadFont create a new font element with the given font size
func LoadFont(fontSize int) (font *ttf.Font, err error) {
	font, err = ttf.OpenFont("fonts/Starjout.ttf", fontSize)
	if err != nil {
		return nil, fmt.Errorf("initializing font:%v", err)
	}
	return font, nil
}
