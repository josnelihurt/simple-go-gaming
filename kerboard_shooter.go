package main

import (
	"time"

	"github.com/josnelihurt/simple-go-gaming/engine"
	"github.com/veandco/go-sdl2/sdl"
)

type keyboardShooter struct {
	parent   *engine.Element
	cooldown time.Duration
	lastShot time.Time
	player   *engine.SoundPlayer
}

func newKeyboardShooter(parent *engine.Element, cooldown time.Duration, audioDev sdl.AudioDeviceID) *keyboardShooter {

	return &keyboardShooter{
		parent:   parent,
		cooldown: cooldown,
		player:   engine.NewSoundPlayer(parent, resSoundLaser, audioDev, []int{engine.MsgNone}, ""),
	}
}
func (context *keyboardShooter) OnUpdate() error {
	keys := sdl.GetKeyboardState()

	parent := context.parent

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(context.lastShot) > context.cooldown {
			context.player.Play()
			sprite := parent.GetComponent(&engine.SpriteRenderer{}).(*engine.SpriteRenderer)
			context.shoot(parent.Position.X+23, parent.Position.Y-sprite.ScaledHeight/2)
			context.shoot(parent.Position.X-23, parent.Position.Y-sprite.ScaledHeight/2)
			context.lastShot = time.Now()
		}
	}
	return nil
}
func (context *keyboardShooter) OnDraw(enderer *sdl.Renderer) error      { return nil }
func (context *keyboardShooter) OnCollision(other *engine.Element) error { return nil }
func (context *keyboardShooter) OnMessage(message *engine.Message) error { return nil }
func (context *keyboardShooter) shoot(x, y float64) {
	if currentBullet, ok := bulletFromPool(); ok {
		currentBullet.Active = true
		currentBullet.Position.X = x
		currentBullet.Position.Y = y
	}
}
