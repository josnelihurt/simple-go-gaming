package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	basicEnemySize = 110
	enemyScale     = 0.1
)

func newBasicEnemy(renderer *sdl.Renderer, position vector) *element {
	basicEnemy := &element{}

	basicEnemy.position = position
	basicEnemy.rotation = 180
	basicEnemy.active = true

	sr := newSpriteRenderer(basicEnemy, renderer, "sprites/basic_enemy.bmp", enemyScale)
	basicEnemy.addCompoenent(sr)
	return basicEnemy
}
