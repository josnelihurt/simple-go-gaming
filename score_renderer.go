package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	scoreFontSize = 30
)

type scoreRenderer struct {
	currenValue    int32
	newValue       int32
	width          int32
	height         int32
	texture        *sdl.Texture
	font           *ttf.Font
	textColor      sdl.Color
	backgoundColor sdl.Color
}

func newScoreRenderer() *scoreRenderer {
	font, err := loadFont()
	if err != nil {
		logger <- fmt.Sprintf("loading font ttf:%v", err)
		panic(err)
	}

	return &scoreRenderer{
		currenValue:    -1,
		font:           font,
		textColor:      sdl.Color{R: 255, G: 255, B: 255},
		backgoundColor: sdl.Color{R: 255, G: 255, B: 255},
	}
}
func (context *scoreRenderer) increase() {
	context.newValue++
	logger <- fmt.Sprintf("newScore:%v", context.newValue)
}
func (context *scoreRenderer) onDraw(renderer *sdl.Renderer) error {
	if context.currenValue != context.newValue {
		logger <- fmt.Sprintf("update texture to:%v", context.newValue)

		surface, err := context.font.RenderUTF8Solid(fmt.Sprintf("%03d", context.newValue), context.textColor)
		if err != nil {
			return fmt.Errorf("creating surface from font %v", err)
		}

		context.texture, err = renderer.CreateTextureFromSurface(surface)
		if err != nil {
			return fmt.Errorf("creating texture from surface %v", err)
		}

		_, _, context.width, context.height, err = context.texture.Query()
		if err != nil {
			return fmt.Errorf("querying texture:%v", err)
		}

		context.currenValue = context.newValue
	}
	x := (screenWidth - context.width) / 2
	y := 10

	renderer.Copy(context.texture,
		&sdl.Rect{X: 0, Y: 0, W: int32(context.width), H: int32(context.height)},
		&sdl.Rect{X: int32(x), Y: int32(y), W: int32(context.width), H: int32(context.height)})
	return nil
}
func (context *scoreRenderer) onUpdate() error {
	return nil
}
func (context *scoreRenderer) onCollision(other *element) error {
	return nil
}
func loadFont() (font *ttf.Font, err error) {
	font, err = ttf.OpenFont("fonts/Starjout.ttf", scoreFontSize)
	if err != nil {
		return nil, fmt.Errorf("initializing font:%v", err)
	}
	return font, nil
}
