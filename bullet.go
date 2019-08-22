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
	context := &engine.Element{}
	context.Z = 10

	context.AddComponent(engine.NewSpriteRenderer(context, renderer, "sprites/bullet.png", bulletScale))
	context.AddComponent(newBulletMover(context, bulletSpeed))
	context.AddComponent(engine.NewCollisionDetecter(context, false, ""))
	context.AddComponent(engine.NewComponentDestroyerOnMessage(context, engine.MsgCollision, "enemy"))

	context.Collisions = append(context.Collisions,
		engine.Circle{
			Center: &context.Position,
			Radius: 5,
		})

	context.Active = false
	context.Rotation = 0.0
	context.Tag = "bullet"

	return context
}
