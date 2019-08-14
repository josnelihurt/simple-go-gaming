package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	basicEnemySize = 65
)

type basicEnemy struct {
	texture *sdl.Texture
	x, y    float64
}

func newBasicEnemy(renderer *sdl.Renderer, x, y float64) (be basicEnemy, err error) {
	be.texture = textureFromBMP(renderer, "sprites/basic_enemy.bmp")
	be.x = x
	be.y = y
	return be, nil
}

func (be *basicEnemy) draw(renderer *sdl.Renderer) {
	x := be.x - basicEnemySize/2.0
	y := be.y - basicEnemySize/2.0
	renderer.CopyEx(be.texture,
		&sdl.Rect{X: 0, Y: 0, W: 200, H: 200},
		&sdl.Rect{X: int32(x), Y: int32(y), W: basicEnemySize, H: basicEnemySize},
		180,
		&sdl.Point{X: basicEnemySize / 2, Y: basicEnemySize / 2},
		sdl.FLIP_NONE)

}

func (be *basicEnemy) destroy() {
	be.texture.Destroy()
}
