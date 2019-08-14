package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSize  = 32
	bulletSpeed = 2
	bulletScale = 0.1
)

type bullet struct {
	texture *sdl.Texture
	x, y    float64
	active  bool
	angle   float64
}

func newBullet(renderer *sdl.Renderer) *element {
	bullet := &element{}

	sr := newSpriteRenderer(bullet, renderer, "sprites/bullet.bmp", bulletScale)
	bullet.addCompoenent(sr)

	mover := newBulletMover(bullet, bulletSpeed)
	bullet.addCompoenent(mover)

	bullet.active = false
	bullet.rotation = 90

	return bullet
}

var bulletPool []*element

func initBulletPool(renderer *sdl.Renderer) []*element {
	for i := 0; i < 30; i++ {
		b := newBullet(renderer)
		bulletPool = append(bulletPool, b)
	}
	return bulletPool
}

func bulletFromPool() (*element, bool) {
	for _, b := range bulletPool {
		if !b.active {
			return b, true
		}
	}
	return nil, false
}
