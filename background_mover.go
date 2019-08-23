package main

import (
	"fmt"

	"github.com/josnelihurt/simple-go-gaming/engine"

	"github.com/veandco/go-sdl2/sdl"
)

type backgroundMover struct {
	parent       *engine.Element
	texture      *sdl.Texture
	currentY     float64
	bitmapWidth  int32
	bitmapHeight int32
}

func newBackgroundMover(renderer *sdl.Renderer) *backgroundMover {
	texture := engine.TextureFromFile(renderer, resSpriteBackground)
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
func (context *backgroundMover) OnUpdate() error {
	context.currentY += backgroundSpeed * delta

	if int32(context.currentY) >= context.bitmapHeight {
		context.currentY = 0
	}
	return nil
}
func (context *backgroundMover) OnDraw(renderer *sdl.Renderer) error {
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
func (context *backgroundMover) OnCollision(other *engine.Element) error {
	return nil
}
func (context *backgroundMover) OnMessage(message *engine.Message) error {
	return nil
}
