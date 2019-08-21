package main

import (
	"github.com/josnelihurt/simple-go-gaming/engine"
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
}

func newBullet(renderer *sdl.Renderer, onCollisionCallback func()) *engine.Element {
	bullet := &engine.Element{}
	bullet.Z = 10

	sr := engine.NewSpriteRenderer(bullet, renderer, "sprites/bullet.png", bulletScale)
	bullet.AddComponent(sr)

	mover := newBulletMover(bullet, bulletSpeed, onCollisionCallback)
	bullet.AddComponent(mover)

	bullet.Collisions = append(bullet.Collisions,
		engine.Circle{
			Center: &bullet.Position,
			Radius: 5,
		})

	bullet.Active = false
	//bullet.rotation = 0.0
	bullet.Tag = "bullet"

	return bullet
}

var bulletPool []*engine.Element

func initBulletPool(renderer *sdl.Renderer, onCollisionCallback func()) []*engine.Element {
	for i := 0; i < 30; i++ {
		b := newBullet(renderer, onCollisionCallback)
		bulletPool = append(bulletPool, b)
	}
	return bulletPool
}

func bulletFromPool() (*engine.Element, bool) {
	for _, b := range bulletPool {
		if !b.Active {
			return b, true
		}
	}
	return nil, false
}
