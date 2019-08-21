package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type soundPlayer struct {
	soundRaw    []byte
	audioDevice sdl.AudioDeviceID
}

func newSoundPlayer(file string, audioDev sdl.AudioDeviceID) (player soundPlayer) {
	soundRaw, _ := sdl.LoadWAV(file)
	//sdl.MixAudio(shootSoundRaw, shootSoundRaw, shootSound.Size, 50)

	return soundPlayer{
		audioDevice: audioDev,
		soundRaw:    soundRaw,
	}
}
func (context *soundPlayer) play() {
	sdl.QueueAudio(context.audioDevice, context.soundRaw)
	sdl.PauseAudioDevice(context.audioDevice, false)
}
