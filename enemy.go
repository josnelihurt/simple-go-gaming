package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	basicEnemySize = 110
	enemyScale     = 1.0
	enemySpeed     = 6.0
)

func newBasicEnemy(renderer *sdl.Renderer, position vector, onDistroyed func()) *element {
	basicEnemy := &element{}
	basicEnemy.z = 10

	basicEnemy.position = position
	basicEnemy.rotation = 0
	basicEnemy.active = true

	basicEnemy.addCompoenent(newSpriteRenderer(basicEnemy, renderer, "sprites/basic_enemy.png", enemyScale))
	basicEnemy.addCompoenent(newVulnerableToElement(basicEnemy, func(origin *element) {
		if origin.tag == "bullet" {
			onDistroyed()
		}
	}, "bullet", "player"))
	basicEnemy.addCompoenent(newEnemyMover(basicEnemy))

	basicEnemy.collisions = append(basicEnemy.collisions,
		circle{
			center: &basicEnemy.position,
			radius: basicEnemySize / 2,
		})
	return basicEnemy
}
func createEnemySwarm(renderer *sdl.Renderer, onEnemyDistroyed func()) (swarm []*element) {
	const rows = 3
	const colums = 6
	for i := 0; i < colums; i++ {
		for j := 0; j < rows; j++ {
			x := (float64(i)/colums)*screenWidth + (basicEnemySize / 2.0)
			y := float64(j)*basicEnemySize + (basicEnemySize / 2.0) + 50

			enemy := newBasicEnemy(renderer, vector{x: x, y: y}, onEnemyDistroyed)
			enemy.tag = "enemy"
			swarm = append(swarm, enemy)
		}
	}
	return swarm
}
