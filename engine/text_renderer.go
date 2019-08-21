package engine

import (
	"fmt"

	"github.com/josnelihurt/simple-go-gaming/engine/util"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type TextRenderer struct {
	currenValue string
	newValue    string
	width       int32
	height      int32
	position    *Vector
	texture     *sdl.Texture
	font        *ttf.Font
	textColor   sdl.Color
}

func NewTextRenderer(position *Vector, fontSize int, textColor sdl.Color) *TextRenderer {
	font, err := LoadFont(fontSize)
	if err != nil {
		util.Logger <- fmt.Sprintf("loading font ttf:%v", err)
		panic(err)
	}

	return &TextRenderer{
		font:      font,
		textColor: textColor,
		position:  position,
	}
}
func (context *TextRenderer) onDraw(renderer *sdl.Renderer) error {
	if context.currenValue != context.newValue {
		util.Logger <- fmt.Sprintf("n:%v,o:%v", context.newValue, context.currenValue)
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
func (context *TextRenderer) onUpdate() error {
	return nil
}
func (context *TextRenderer) onCollision(other *Element) error {
	return nil
}
