package main

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/josnelihurt/simple-go-gaming/engine"
)

const (
	basicEnemySize = 110
	enemyScale     = 1.0
	enemySpeed     = 6.0
)

func newBasicEnemy(renderer *sdl.Renderer, position engine.Vector) *engine.Element {
	context := &engine.Element{}
	context.Z = 10

	context.Position = position
	context.Rotation = 0
	context.Active = true

	context.AddComponent(engine.NewSpriteRenderer(context, renderer, "sprites/basic_enemy.png", enemyScale))
	context.AddComponent(engine.NewCollisionDetecter(context, true, "bullet", "player"))
	context.AddComponent(newEnemyMover(context))
	context.AddComponent(engine.NewComponentDestroyerOnMessage(context, engine.MsgCollision, "bullet", "player"))
	context.Collisions = append(context.Collisions,
		engine.Circle{
			Center: &context.Position,
			Radius: basicEnemySize / 2,
		})
	return context
}
