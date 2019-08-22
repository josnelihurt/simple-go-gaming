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
	basicEnemy := &engine.Element{}
	basicEnemy.Z = 10

	basicEnemy.Position = position
	basicEnemy.Rotation = 0
	basicEnemy.Active = true

	basicEnemy.AddComponent(engine.NewSpriteRenderer(basicEnemy, renderer, "sprites/basic_enemy.png", enemyScale))
	basicEnemy.AddComponent(engine.NewCollisionDetecter(basicEnemy, true, "bullet", "player"))
	basicEnemy.AddComponent(newEnemyMover(basicEnemy))
	basicEnemy.AddComponent(engine.NewComponentDestroyerOnMessage(basicEnemy, engine.MsgCollision, "bullet"))
	basicEnemy.Collisions = append(basicEnemy.Collisions,
		engine.Circle{
			Center: &basicEnemy.Position,
			Radius: basicEnemySize / 2,
		})
	return basicEnemy
}
