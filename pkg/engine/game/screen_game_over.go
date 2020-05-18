package game

import (
	"github.com/josnelihurt/simple-go-gaming/pkg/engine"
	"github.com/veandco/go-sdl2/sdl"
)

func newGameOverScreen(components *engine.SDLComponents) *engine.Element {
	context := &engine.Element{
		Active: false,
		Tag:    "GameOverScreen",
	}
	textPosition := &engine.Vector{X: screenWidth / 2, Y: screenHeight / 2}
	textRenderer := engine.NewTextRenderer(textPosition, 40, sdl.Color{R: 0, G: 255, B: 0})
	textRenderer.SetNewText("Game over !!")
	textRenderer.RenderNewValue(components.Renderer)
	_, _, width, _ := textRenderer.GetCurrentRenderInfo()
	textPosition.X -= width / 2
	context.AddComponent(textRenderer)
	return context
}
