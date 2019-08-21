package engine

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type SpriteRenderer struct {
	parent                       *Element
	texture                      *sdl.Texture
	originalWidth, orginalHeight float64
	ScaledWidth, ScaledHeight    float64
}

func TextureFromFile(renderer *sdl.Renderer, filename string) (texture *sdl.Texture) {
	var image *sdl.Surface
	var err error

	image, err = img.Load(filename)
	if err != nil {
		panic(fmt.Errorf("loading %v: %v", filename, err))
	}
	defer image.Free()

	texture, err = renderer.CreateTextureFromSurface(image)
	if err != nil {
		panic(fmt.Errorf("creating  basic enemy texture: %v", err))
	}
	return texture
}
func NewSpriteRenderer(parent *Element, renderer *sdl.Renderer, filename string, scale float64) *SpriteRenderer {
	texture := TextureFromFile(renderer, filename)
	_, _, width, height, err := texture.Query()
	if err != nil {
		panic(fmt.Errorf("querying texture:%v", err))
	}
	result := &SpriteRenderer{
		parent:        parent,
		texture:       texture,
		originalWidth: float64(width),
		orginalHeight: float64(height),
		ScaledWidth:   float64(width) * scale,
		ScaledHeight:  float64(height) * scale,
	}

	//logger <- fmt.Sprintln("new spriteRenderer:", result)

	return result
}
func (context *SpriteRenderer) OnDraw(renderer *sdl.Renderer) error {
	// Converting coordinates to top left of sprite
	x := context.parent.Position.X - context.ScaledWidth/2.0
	y := context.parent.Position.Y - context.ScaledHeight/2.0

	renderer.CopyEx(context.texture,
		&sdl.Rect{X: 0, Y: 0, W: int32(context.originalWidth), H: int32(context.orginalHeight)},
		&sdl.Rect{X: int32(x), Y: int32(y), W: int32(context.ScaledWidth), H: int32(context.ScaledHeight)},
		context.parent.Rotation,
		&sdl.Point{X: int32(context.ScaledWidth / 2.0), Y: int32(context.ScaledHeight / 2.0)},
		sdl.FLIP_NONE)
	return nil
}
func (context *SpriteRenderer) OnUpdate() error {
	return nil
}
func (context *SpriteRenderer) OnCollision(other *Element) error {
	return nil
}
