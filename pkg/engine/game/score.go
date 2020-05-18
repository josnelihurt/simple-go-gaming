package game

import (
	"github.com/josnelihurt/simple-go-gaming/pkg/engine"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	scoreX   = 220
	tagScore = "score"
)

func newScore() *engine.Element {
	context := &engine.Element{Active: true, Tag: tagScore, Z: 99}
	textRenderer := engine.NewTextRenderer(
		&engine.Vector{X: (screenWidth - scoreX), Y: upperTextY},
		defaultFontSize,
		sdl.Color{R: 255, G: 255, B: 255})
	context.AddComponent(textRenderer)
	context.AddComponent(engine.NewCounterText(context, textRenderer, 0, 1, -1, "score:%03d", incrementCondition, nil, nil))
	return context
}
func incrementCondition(message *engine.Message, parent *engine.Element) bool {
	return message.Code == engine.MsgCollision && message.Sender.Tag == tagEnemy && message.RelatedTo[0].Tag == tagBullet
}
