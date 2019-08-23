package main

import (
	"fmt"
	"math"
	"time"

	"github.com/josnelihurt/simple-go-gaming/engine"
	"github.com/josnelihurt/simple-go-gaming/engine/util"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed        = 5.0
	playerShotCooldown = time.Millisecond * 200
	playerScale        = 0.7
	tagPlayer          = "player"
)

func newPlayer(components *engine.SDLComponents) *engine.Element {
	context := &engine.Element{}
	context.Z = 10
	context.Active = true
	context.Rotation = 180
	context.Tag = tagPlayer

	currentSpriteRenderer := engine.NewSpriteRenderer(context, components.Renderer, resSpritePlayer, playerScale)
	playerDefaultPosition := engine.Vector{
		X: screenWidth / 2.0,
		Y: screenHeight - currentSpriteRenderer.ScaledHeight/2.0,
	}
	context.Position = playerDefaultPosition
	context.AddComponent(currentSpriteRenderer)
	context.AddComponent(newPlayerStarter(context, playerDefaultPosition))
	context.AddComponent(newKeyboardShooter(context, playerShotCooldown, components.AudioDev))
	textRenderer := engine.NewTextRenderer(
		&engine.Vector{X: upperTextY, Y: upperTextY},
		defaultFontSize,
		sdl.Color{R: 255, G: 255, B: 255})
	context.AddComponent(textRenderer)
	context.AddComponent(engine.NewCounterText(context, textRenderer, 3, -1, 0, "life:%03d", incrementConditionPlayer, resetPlayerLife, onPlayerLifeEmpty))
	allowedRect := &engine.Rect{
		X:      currentSpriteRenderer.ScaledWidth / 2.0,
		Y:      4 * screenHeight / 5,
		Width:  screenWidth - currentSpriteRenderer.ScaledWidth,
		Height: screenHeight/5 - currentSpriteRenderer.ScaledHeight/2.0}
	context.AddComponent(engine.NewKeyboardMover(context, allowedRect, &delta, playerSpeed))
	context.AddComponent(engine.NewCollisionDetecter(context, true, tagEnemy))
	context.AddComponent(engine.NewSoundPlayer(context, resSoundExplosion, components.AudioDev, []int{engine.MsgCollision, msgHitPlayer}, tagEnemy))
	context.Collisions = append(context.Collisions,
		engine.Circle{
			Center: &context.Position,
			Radius: math.Max(currentSpriteRenderer.ScaledWidth, currentSpriteRenderer.ScaledHeight) / 2,
		})

	return context
}

func incrementConditionPlayer(message *engine.Message, context *engine.Element) bool {
	return (message.Code == engine.MsgCollision &&
		message.Sender.Tag == tagEnemy &&
		len(message.RelatedTo) > 0 &&
		message.RelatedTo[0].Tag == tagPlayer) || message.Code == msgHitPlayer
}
func resetPlayerLife(message *engine.Message, context *engine.Element) bool {
	return message.Code == msgPlayerDead
}

func onPlayerLifeEmpty(context *engine.Element) {
	util.Logger <- fmt.Sprintf(":/")
	context.BroadcastMessage(&engine.Message{
		Code:                msgPlayerDead,
		Sender:              context,
		Data:                "playerDead",
		SendToOtherElements: true,
	})
}
