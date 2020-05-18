package game

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/josnelihurt/simple-go-gaming/pkg/engine"
)

const (
	basicEnemySize = 110
	enemyScale     = 1.0
	enemySpeed     = 3.0
	tagEnemy       = "enemy"
)

func newBasicEnemy(renderer *sdl.Renderer, position engine.Vector) *engine.Element {
	context := &engine.Element{}
	context.Z = 10

	context.Position = position
	context.Rotation = 0
	context.Active = true
	context.Tag = tagEnemy

	context.AddComponent(engine.NewSpriteRenderer(context, renderer, resSpriteEnemy, enemyScale))
	context.AddComponent(engine.NewCollisionDetecter(context, true, tagBullet, tagPlayer))
	context.AddComponent(newEnemyMover(context))
	context.AddComponent(engine.NewComponentDestroyerOnMessage(context, engine.MsgCollision, tagBullet, tagPlayer))
	context.Collisions = append(context.Collisions,
		engine.Circle{
			Center: &context.Position,
			Radius: basicEnemySize / 2,
		})
	return context
}
