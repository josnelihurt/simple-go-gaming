package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type textRenderer struct {
	currenValue string
	newValue    string
	width       int32
	height      int32
	position    *vector
	texture     *sdl.Texture
	font        *ttf.Font
	textColor   sdl.Color
}

func newTextRenderer(position *vector, fontSize int, textColor sdl.Color) *textRenderer {
	font, err := loadFont(fontSize)
	if err != nil {
		logger <- fmt.Sprintf("loading font ttf:%v", err)
		panic(err)
	}

	return &textRenderer{
		font:      font,
		textColor: textColor,
		position:  position,
	}
}
func (context *textRenderer) onDraw(renderer *sdl.Renderer) error {
	if context.currenValue != context.newValue {
		surface, err := context.font.RenderUTF8Solid(context.newValue, context.textColor)
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
	x := context.position.x
	y := context.position.y

	renderer.Copy(context.texture,
		&sdl.Rect{X: 0, Y: 0, W: int32(context.width), H: int32(context.height)},
		&sdl.Rect{X: int32(x), Y: int32(y), W: int32(context.width), H: int32(context.height)})
	return nil
}
func (context *textRenderer) onUpdate() error {
	return nil
}
func (context *textRenderer) onCollision(other *element) error {
	return nil
}
