package main

import (
	"github.com/josnelihurt/simple-go-gaming/engine"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSpeed = 10
	bulletScale = 1.0
)

func newBullet(renderer *sdl.Renderer) *engine.Element {
	bullet := &engine.Element{}
	bullet.Z = 10

	bullet.AddComponent(engine.NewSpriteRenderer(bullet, renderer, "sprites/bullet.png", bulletScale))
	bullet.AddComponent(newBulletMover(bullet, bulletSpeed))
	bullet.AddComponent(engine.NewCollisionDetecter(bullet, false, ""))
	bullet.AddComponent(engine.NewComponentDestroyerOnMessage(bullet, engine.MsgCollision, "enemy"))

	bullet.Collisions = append(bullet.Collisions,
		engine.Circle{
			Center: &bullet.Position,
			Radius: 5,
		})

	bullet.Active = false
	bullet.Rotation = 0.0
	bullet.Tag = "bullet"

	return bullet
}
