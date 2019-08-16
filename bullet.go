package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSpeed = 10
	bulletScale = 1.0
)

type bullet struct {
	texture *sdl.Texture
	x, y    float64
	active  bool
	angle   float64
	score   *scoreRenderer
}

func newBullet(renderer *sdl.Renderer, scoreRenderer *scoreRenderer) *element {
	bullet := &element{}

	sr := newSpriteRenderer(bullet, renderer, "sprites/bullet.bmp", bulletScale)
	bullet.addCompoenent(sr)

	mover := newBulletMover(bullet, bulletSpeed)
	bullet.addCompoenent(mover)
	bullet.addCompoenent(scoreRenderer)

	bullet.collisions = append(bullet.collisions,
		circle{
			center: bullet.position,
			radius: 5,
		})

	bullet.active = false
	bullet.rotation = 0.0
	bullet.tag = "bullet"

	return bullet
}

var bulletPool []*element

func initBulletPool(renderer *sdl.Renderer, scoreRenderer *scoreRenderer) []*element {
	for i := 0; i < 30; i++ {
		b := newBullet(renderer, scoreRenderer)
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
