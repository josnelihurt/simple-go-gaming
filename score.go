package main

import (
	"github.com/josnelihurt/simple-go-gaming/engine"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	scoreFontSize = 34
	scoreY        = 15
	scoreX        = 90
)

func newScore() *engine.Element {
	score := &engine.Element{Active: true, Tag: "score", Z: 99}
	score.AddComponent(engine.NewTextRenderer(
		&engine.Vector{X: (screenWidth - scoreX), Y: scoreY},
		scoreFontSize,
		sdl.Color{R: 255, G: 255, B: 255}))
	score.AddComponent(newScoreCounter(score))
	return score
}
