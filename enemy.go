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

func newBasicEnemy(renderer *sdl.Renderer, position engine.Vector, onDistroyed func()) *engine.Element {
	basicEnemy := &engine.Element{}
	basicEnemy.Z = 10

	basicEnemy.Position = position
	basicEnemy.Rotation = 0
	basicEnemy.Active = true

	basicEnemy.AddComponent(engine.NewSpriteRenderer(basicEnemy, renderer, "sprites/basic_enemy.png", enemyScale))
	basicEnemy.AddComponent(engine.NewVulnerableToElement(basicEnemy, func(origin *engine.Element) {
		if origin.Tag == "bullet" {
			onDistroyed()
		}
	}, "bullet", "player"))
	basicEnemy.AddComponent(newEnemyMover(basicEnemy))

	basicEnemy.Collisions = append(basicEnemy.Collisions,
		engine.Circle{
			Center: &basicEnemy.Position,
			Radius: basicEnemySize / 2,
		})
	return basicEnemy
}
