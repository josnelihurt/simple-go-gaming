package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	basicEnemySize = 110
	enemyScale     = 1.0
	enemySpeed     = 1.0
)

func newBasicEnemy(renderer *sdl.Renderer, position vector) *element {
	basicEnemy := &element{}
	basicEnemy.z = 10

	basicEnemy.position = position
	basicEnemy.rotation = 0
	basicEnemy.active = true

	basicEnemy.addCompoenent(newSpriteRenderer(basicEnemy, renderer, "sprites/basic_enemy.png", enemyScale))
	basicEnemy.addCompoenent(newVulnerableToBullets(basicEnemy))
	basicEnemy.addCompoenent(newEnemyMover())

	basicEnemy.collisions = append(basicEnemy.collisions,
		circle{
			center: basicEnemy.position,
			radius: basicEnemySize / 2,
		})
	return basicEnemy
}

func createEnemySwarm(renderer *sdl.Renderer) (swarm []*element) {
	const rows = 3
	const colums = 6
	for i := 0; i < colums; i++ {
		for j := 0; j < rows; j++ {
			x := (float64(i)/colums)*screenWidth + (basicEnemySize / 2.0)
			y := float64(j)*basicEnemySize + (basicEnemySize / 2.0) + 50

			enemy := newBasicEnemy(renderer, vector{x: x, y: y})
			enemy.tag = fmt.Sprintf("x:%v,y:%v", i, j)
			swarm = append(swarm, enemy)
		}
	}
	return swarm
}
