package engine

import (
	"fmt"

	"github.com/josnelihurt/simple-go-gaming/engine/util"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

//TextRenderer shows a new text on screen
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

//NewTextRenderer creates a new instance of TextRenderer
func NewTextRenderer(position *Vector, fontSize int, textColor sdl.Color) *TextRenderer {
	font, err := util.LoadFont(fontSize)
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

// GetCurrentRenderInfo obtains information about the internal text rendering
func (context *TextRenderer) GetCurrentRenderInfo() (x, y, width, height float64) {
	return (context.position.X), (context.position.Y), float64(context.width), float64(context.height)
}

//SetNewText update the internal text value
func (context *TextRenderer) SetNewText(text string) {
	context.newValue = text
}

//RenderNewValue update the internal surface, this operation could be expensive,
// and this is called automatiacally if the internal text is not the same as the text represented on screen (OnDraw)
func (context *TextRenderer) RenderNewValue(renderer *sdl.Renderer) error {
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
	return nil
}

func (context *TextRenderer) OnDraw(renderer *sdl.Renderer) error {
	if context.currenValue != context.newValue {
		if err := context.RenderNewValue(renderer); err != nil {
			return err
		}
	}
	x := context.position.X
	y := context.position.Y

	renderer.Copy(context.texture,
		&sdl.Rect{X: 0, Y: 0, W: int32(context.width), H: int32(context.height)},
		&sdl.Rect{X: int32(x), Y: int32(y), W: int32(context.width), H: int32(context.height)})
	return nil
}
func (context *TextRenderer) OnUpdate() error                  { return nil }
func (context *TextRenderer) OnCollision(other *Element) error { return nil }
func (context *TextRenderer) OnMessage(message *Message) error { return nil }
