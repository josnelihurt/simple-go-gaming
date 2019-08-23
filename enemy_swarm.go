package main

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/josnelihurt/simple-go-gaming/engine"
)

const (
	swarmRows   = 3
	swarmColums = 6
)

type enemySrawm struct {
	enemies []*engine.Element
}

func newEnemySwarm(renderer *sdl.Renderer) *enemySrawm {
	context := &enemySrawm{}
	context.enemies = make([]*engine.Element, swarmColums*swarmRows)
	for i := 0; i < swarmColums*swarmRows; i++ {
		context.enemies[i] = newBasicEnemy(renderer, engine.Vector{X: 0, Y: 0})
		context.enemies[i].Tag = tagEnemy
	}
	context.activateAll()
	context.setInitialPositions()
	return context
}
func (context *enemySrawm) setInitialPositions() {
	for j := 0; j < swarmRows; j++ {
		for i := 0; i < swarmColums; i++ {
			index := i + j*swarmColums
			context.enemies[index].Position.X = (float64(i)/swarmColums)*screenWidth + (basicEnemySize / 2.0)
			context.enemies[index].Position.Y = float64(j)*basicEnemySize + (basicEnemySize / 2.0) + 50
			context.enemies[index].GetFirstComponent(&enemyMover{}).(*enemyMover).active = false
		}
	}
}
func (context *enemySrawm) activateAll() {
	for _, currentElement := range context.enemies {
		currentElement.Active = true
	}
}
func (context *enemySrawm) deActivateAll() {
	for _, currentElement := range context.enemies {
		currentElement.Active = false
	}
}
func (context *enemySrawm) getActives() (result []*engine.Element) {
	for _, currentElement := range context.enemies {
		if currentElement.Active {
			result = append(result, currentElement)
		}
	}
	return result
}
