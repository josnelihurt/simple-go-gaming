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

func newBullet(renderer *sdl.Renderer, onCollisionCallback func()) *element {
	bullet := &element{}
	bullet.z = 10

	sr := newSpriteRenderer(bullet, renderer, "sprites/bullet.png", bulletScale)
	bullet.addCompoenent(sr)

	mover := newBulletMover(bullet, bulletSpeed, onCollisionCallback)
	bullet.addCompoenent(mover)

	bullet.collisions = append(bullet.collisions,
		circle{
			center: &bullet.position,
			radius: 5,
		})

	bullet.active = false
	bullet.rotation = 0.0
	bullet.tag = "bullet"

	return bullet
}

var bulletPool []*element

func initBulletPool(renderer *sdl.Renderer, onCollisionCallback func()) []*element {
	for i := 0; i < 30; i++ {
		b := newBullet(renderer, onCollisionCallback)
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
