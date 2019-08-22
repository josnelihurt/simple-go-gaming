package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

// SoundPlayer represents a player for a given track
type SoundPlayer struct {
	parent       *Element
	soundRaw     []byte
	audioDevice  sdl.AudioDeviceID
	relatedToTag []string
	messageCodes []int
}

// NewSoundPlayer creates new SoundPlayer
func NewSoundPlayer(parent *Element, file string, audioDev sdl.AudioDeviceID, messageCodes []int, relatedToTag ...string) *SoundPlayer {
	soundRaw, _ := sdl.LoadWAV(file)
	return &SoundPlayer{
		parent:       parent,
		audioDevice:  audioDev,
		soundRaw:     soundRaw,
		relatedToTag: relatedToTag,
		messageCodes: messageCodes,
	}
}

// Play sends the audio to the main device to playback the sound
func (context *SoundPlayer) Play() {
	sdl.QueueAudio(context.audioDevice, context.soundRaw)
	sdl.PauseAudioDevice(context.audioDevice, false)
}

func (context *SoundPlayer) OnDraw(enderer *sdl.Renderer) error { return nil }
func (context *SoundPlayer) OnUpdate() error                    { return nil }
func (context *SoundPlayer) OnCollision(other *Element) error   { return nil }
func (context *SoundPlayer) OnMessage(message *Message) error {
	if containsInt(context.messageCodes, message.Code) {
		if isSingleAndEmpty(context.relatedToTag) || contains(context.relatedToTag, message.RelatedTo[0].Tag) {
			context.Play()
		}
	}
	return nil
}
