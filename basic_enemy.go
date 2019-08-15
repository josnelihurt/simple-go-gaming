package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	basicEnemySize = 110
	enemyScale     = 0.7
)

func newBasicEnemy(renderer *sdl.Renderer, position vector) *element {
	basicEnemy := &element{}

	basicEnemy.position = position
	basicEnemy.rotation = 180
	basicEnemy.active = true

	basicEnemy.addCompoenent(newSpriteRenderer(basicEnemy, renderer, "sprites/basic_enemy.bmp", enemyScale))
	basicEnemy.addCompoenent(newVulnerableToBullets(basicEnemy))

	basicEnemy.collisions = append(basicEnemy.collisions,
		circle{
			center: basicEnemy.position,
			radius: 10,
		})
	return basicEnemy
}
