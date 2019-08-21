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
	basicEnemy.AddComponent(engine.NewCollisionDetecter(basicEnemy, "bullet", "player"))
	basicEnemy.AddComponent(newEnemyMover(basicEnemy))
	basicEnemy.AddComponent(NewComponentDestroyerOnCollision(basicEnemy))
	basicEnemy.Collisions = append(basicEnemy.Collisions,
		engine.Circle{
			Center: &basicEnemy.Position,
			Radius: basicEnemySize / 2,
		})
	return basicEnemy
}

type componentDestroyerOnCollision struct {
	parent *engine.Element
}

func NewComponentDestroyerOnCollision(parent *engine.Element) *componentDestroyerOnCollision {
	return &componentDestroyerOnCollision{
		parent: parent,
	}
}

func (context *componentDestroyerOnCollision) OnUpdate() error                         { return nil }
func (context *componentDestroyerOnCollision) OnDraw(renderer *sdl.Renderer) error     { return nil }
func (context *componentDestroyerOnCollision) OnCollision(other *engine.Element) error { return nil }
func (context *componentDestroyerOnCollision) OnMessage(message *engine.Message) error {
	context.parent.Active = false
	if message.RelatedTo[0].Tag == "player" {
		message.SendToOtherElements = true
	}
	return nil
}
