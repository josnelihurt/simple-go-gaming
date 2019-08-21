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
	scaledWidth, scaledHeight    float64
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
		scaledWidth:   float64(width) * scale,
		scaledHeight:  float64(height) * scale,
	}

	//logger <- fmt.Sprintln("new spriteRenderer:", result)

	return result
}
func (context *SpriteRenderer) onDraw(renderer *sdl.Renderer) error {
	// Converting coordinates to top left of sprite
	x := context.parent.position.x - context.scaledWidth/2.0
	y := context.parent.position.y - context.scaledHeight/2.0

	renderer.CopyEx(context.texture,
		&sdl.Rect{X: 0, Y: 0, W: int32(context.originalWidth), H: int32(context.orginalHeight)},
		&sdl.Rect{X: int32(x), Y: int32(y), W: int32(context.scaledWidth), H: int32(context.scaledHeight)},
		context.parent.rotation,
		&sdl.Point{X: int32(context.scaledWidth / 2.0), Y: int32(context.scaledHeight / 2.0)},
		sdl.FLIP_NONE)
	return nil
}
func (context *SpriteRenderer) onUpdate() error {
	return nil
}
func (context *SpriteRenderer) onCollision(other *Element) error {
	return nil
}
