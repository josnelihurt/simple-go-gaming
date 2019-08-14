package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSize  = 32
	bulletSpeed = 2
)

type bullet struct {
	texture *sdl.Texture
	x, y    float64
	active  bool
	angle   float64
}

func newBullet(renderer *sdl.Renderer) (bul bullet) {
	bul.texture = textureFromBMP(renderer, "sprites/bullet.bmp")
	return bul
}

func (b *bullet) draw(renderer *sdl.Renderer) {
	if !b.active {
		return
	}

	x := b.x - bulletSize/2.0
	y := b.y - bulletSize/2.0
	renderer.Copy(b.texture,
		&sdl.Rect{X: 0, Y: 0, W: 200, H: 200},
		&sdl.Rect{X: int32(x), Y: int32(y), W: bulletSize, H: bulletSize})
}

func (b *bullet) update() {
	b.x += bulletSpeed * math.Cos(b.angle)
	b.y += bulletSpeed * math.Sin(b.angle)
	if b.x > screenWidth || b.x < 0 ||
		b.y < 0 {
		b.active = false
	}
}

var bulletPool []*bullet

func initBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		b := newBullet(renderer)
		bulletPool = append(bulletPool, &b)
	}
}

func bulletFromPool() (*bullet, bool) {
	for _, b := range bulletPool {
		if !b.active {
			return b, true
		}
	}
	return nil, false
}
