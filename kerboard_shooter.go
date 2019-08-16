package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type keyboardShooter struct {
	parent        *element
	cooldown      time.Duration
	lastShot      time.Time
	shootSound    *sdl.AudioSpec
	shootSoundRaw []byte
	audioDevice   sdl.AudioDeviceID
}

func newKeyboardShooter(parent *element, cooldown time.Duration) *keyboardShooter {

	shootSoundRaw, shootSound := sdl.LoadWAV("sounds/NFF-laser.wav")
	sdl.MixAudio(shootSoundRaw, shootSoundRaw, shootSound.Size, 50)

	logger <- fmt.Sprintf("sound %v ", len(shootSoundRaw))
	currenAudioDriver := sdl.GetCurrentAudioDriver()
	logger <- currenAudioDriver
	dev, err := sdl.OpenAudioDevice("", false, shootSound, nil, 0)
	if err != nil {
		logger <- fmt.Sprintf("error opeing audio dev:%v", err)
	}

	return &keyboardShooter{
		parent:        parent,
		cooldown:      cooldown,
		shootSound:    shootSound,
		shootSoundRaw: shootSoundRaw,
		audioDevice:   dev,
	}
}
func (context *keyboardShooter) onUpdate() error {
	keys := sdl.GetKeyboardState()

	parent := context.parent

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(context.lastShot) > context.cooldown {
			sdl.QueueAudio(context.audioDevice, context.shootSoundRaw)
			sdl.PauseAudioDevice(context.audioDevice, false)
			sprite := parent.getComponent(&spriteRenderer{}).(*spriteRenderer)
			context.shoot(parent.position.x+23, parent.position.y-sprite.scaledHeight/2)
			context.shoot(parent.position.x-23, parent.position.y-sprite.scaledHeight/2)
			context.lastShot = time.Now()
		}
	}
	return nil
}
func (context *keyboardShooter) onDraw(enderer *sdl.Renderer) error {
	return nil
}

func (context *keyboardShooter) onCollision(other *element) error {
	return nil
}
func (context *keyboardShooter) shoot(x, y float64) {
	if currentBullet, ok := bulletFromPool(); ok {
		currentBullet.active = true
		currentBullet.position.x = x
		currentBullet.position.y = y
	}
}
