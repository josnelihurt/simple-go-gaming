package main

import (
	"fmt"

	"github.com/josnelihurt/simple-go-gaming/engine"
	"github.com/josnelihurt/simple-go-gaming/engine/util"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	tagLevel = "level"
)

func newLevel(sdlComponents *engine.SDLComponents) *engine.Element {
	context := &engine.Element{
		Active:   true,
		Tag:      tagLevel,
		Z:        99,
		Position: engine.Vector{X: screenWidth / 2, Y: upperTextY},
	}
	textRenderer := engine.NewTextRenderer(&context.Position, defaultFontSize, sdl.Color{R: 255, G: 255, B: 255})
	levelCounter := engine.NewCounterText(context, textRenderer, 0, 1, 99, "level:%02d", incrementConditionLevel, resetConditionLevel, nil)
	levelCounter.OnUpdate()
	textRenderer.RenderNewValue(sdlComponents.Renderer)
	_, _, width, _ := textRenderer.GetCurrentRenderInfo()
	context.Position.X -= width / 2
	context.AddComponent(textRenderer)
	context.AddComponent(levelCounter)
	util.Logger <- fmt.Sprintf("Level:%v", context)
	return context
}

func incrementConditionLevel(message *engine.Message, context *engine.Element) bool {
	return message.Code == msgLevelUp
}

func resetConditionLevel(message *engine.Message, context *engine.Element) bool {
	return message.Code == msgPlayerDead
}
