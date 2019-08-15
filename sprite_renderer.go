package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type spriteRenderer struct {
	parent                       *element
	texture                      *sdl.Texture
	originalWidth, orginalHeight float64
	scaledWidth, scaledHeight    float64
}

func textureFromBMP(renderer *sdl.Renderer, filename string) (texture *sdl.Texture) {
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		panic(fmt.Errorf("loading %v: %v", filename, err))
	}
	defer img.Free()

	texture, err = renderer.CreateTextureFromSurface(img)
	if err != nil {

		panic(fmt.Errorf("creating  basic enemy texture: %v", err))
	}
	return texture
}
func newSpriteRenderer(parent *element, renderer *sdl.Renderer, filename string, scale float64) *spriteRenderer {
	texture := textureFromBMP(renderer, filename)
	_, _, width, height, err := texture.Query()
	if err != nil {
		panic(fmt.Errorf("querying texture:%v", err))
	}
	result := &spriteRenderer{
		parent:        parent,
		texture:       texture,
		originalWidth: float64(width),
		orginalHeight: float64(height),
		scaledWidth:   float64(width) * scale,
		scaledHeight:  float64(height) * scale,
	}

	//fmt.Println("new spriteRenderer:", result)

	return result
}
func (context *spriteRenderer) onDraw(renderer *sdl.Renderer) error {
	// Converting coordinates to top left of sprite
	x := context.parent.position.x - context.scaledWidth/2.0
	y := context.parent.position.y - context.scaledHeight/2.0
	//fmt.Println(x, y)
	renderer.CopyEx(context.texture,
		&sdl.Rect{X: 0, Y: 0, W: int32(context.originalWidth), H: int32(context.orginalHeight)},
		&sdl.Rect{X: int32(x), Y: int32(y), W: int32(context.scaledWidth), H: int32(context.scaledHeight)},
		context.parent.rotation,
		&sdl.Point{X: int32(context.scaledWidth / 2.0), Y: int32(context.scaledHeight / 2.0)},
		sdl.FLIP_NONE)

	return nil
}
func (context *spriteRenderer) onUpdate() error {
	return nil
}
func (context *spriteRenderer) onCollision(other *element) error {
	return nil
}
