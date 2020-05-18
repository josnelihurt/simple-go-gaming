package main

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/josnelihurt/simple-go-gaming/engine"
)

const (
	swarmRows    = 3
	SwarmColumns = 6
)

type enemySwarm struct {
	enemies []*engine.Element
}

func newEnemySwarm(renderer *sdl.Renderer) *enemySwarm {
	context := &enemySwarm{}
	context.enemies = make([]*engine.Element, SwarmColumns*swarmRows)
	for i := 0; i < SwarmColumns*swarmRows; i++ {
		context.enemies[i] = newBasicEnemy(renderer, engine.Vector{X: 0, Y: 0})
		context.enemies[i].Tag = tagEnemy
	}
	context.activateAll()
	context.setInitialPositions()
	return context
}
func (context *enemySwarm) setInitialPositions() {
	for j := 0; j < swarmRows; j++ {
		for i := 0; i < SwarmColumns; i++ {
			index := i + j*SwarmColumns
			context.enemies[index].Position.X = (float64(i)/SwarmColumns)*screenWidth + (basicEnemySize / 2.0)
			context.enemies[index].Position.Y = float64(j)*basicEnemySize + (basicEnemySize / 2.0) + 50
			context.enemies[index].GetFirstComponent(&enemyMover{}).(*enemyMover).active = false
		}
	}
}
func (context *enemySwarm) activateAll() {
	for _, currentElement := range context.enemies {
		currentElement.Active = true
	}
}
func (context *enemySwarm) deActivateAll() {
	for _, currentElement := range context.enemies {
		currentElement.Active = false
	}
}
func (context *enemySwarm) getActives() (result []*engine.Element) {
	for _, currentElement := range context.enemies {
		if currentElement.Active {
			result = append(result, currentElement)
		}
	}
	return result
}
