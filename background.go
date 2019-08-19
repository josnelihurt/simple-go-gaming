package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	backgroundSpeed = 0.5
)

func newBackground(renderer *sdl.Renderer) *element {
	background := &element{
		active: true,
	}
	background.addCompoenent(newBackgroundMover(renderer))
	return background
}

type backgroundMover struct {
	parent       *element
	texture      *sdl.Texture
	currentY     float64
	bitmapWidth  int32
	bitmapHeight int32
}

func newBackgroundMover(renderer *sdl.Renderer) *backgroundMover {
	texture := textureFromFile(renderer, "sprites/background_space.png")
	_, _, width, height, err := texture.Query()
	if err != nil {
		panic(fmt.Errorf("querying texture:%v", err))
	}
	return &backgroundMover{
		texture:      texture,
		bitmapHeight: height,
		bitmapWidth:  width,
	}
}
func (context *backgroundMover) onUpdate() error {
	context.currentY += backgroundSpeed * delta

	if int32(context.currentY) >= context.bitmapHeight {
		context.currentY = 0
	}
	return nil
}
func (context *backgroundMover) onDraw(renderer *sdl.Renderer) error {
	x := 0
	y := context.currentY
	remainingY := context.bitmapHeight - int32(y)
	if remainingY < int32(screenHeight) {
		renderer.Copy(context.texture,
			&sdl.Rect{X: int32(x), Y: int32(y), W: screenWidth, H: int32(remainingY)},
			&sdl.Rect{X: 0, Y: 0, W: screenWidth, H: remainingY},
		)
		renderer.Copy(context.texture,
			&sdl.Rect{X: int32(x), Y: int32(0), W: screenWidth, H: int32(screenHeight - remainingY)},
			&sdl.Rect{X: 0, Y: int32(remainingY), W: screenWidth, H: int32(screenHeight - remainingY)},
		)
	} else {
		renderer.Copy(context.texture,
			&sdl.Rect{X: int32(x), Y: int32(y), W: screenWidth, H: int32(screenHeight)},
			&sdl.Rect{X: 0, Y: 0, W: screenWidth, H: screenHeight},
		)
	}

	return nil
}
func (context *backgroundMover) onCollision(other *element) error {
	return nil
}
