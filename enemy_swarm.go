package main

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/josnelihurt/simple-go-gaming/engine"
)

func createEnemySwarm(renderer *sdl.Renderer, onEnemyDistroyed func()) (swarm []*engine.Element) {
	const rows = 3
	const colums = 6
	for i := 0; i < colums; i++ {
		for j := 0; j < rows; j++ {
			x := (float64(i)/colums)*screenWidth + (basicEnemySize / 2.0)
			y := float64(j)*basicEnemySize + (basicEnemySize / 2.0) + 50

			enemy := newBasicEnemy(renderer, engine.Vector{X: x, Y: y}, onEnemyDistroyed)
			enemy.Tag = "enemy"
			swarm = append(swarm, enemy)
		}
	}
	return swarm
}
