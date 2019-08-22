package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

type SoundPlayer struct {
	parent       *Element
	soundRaw     []byte
	audioDevice  sdl.AudioDeviceID
	relatedToTag []string
	messageCode  int
}

func NewSoundPlayer(parent *Element, file string, audioDev sdl.AudioDeviceID, messageCode int, relatedToTag ...string) SoundPlayer {
	soundRaw, _ := sdl.LoadWAV(file)
	return SoundPlayer{
		parent:       parent,
		audioDevice:  audioDev,
		soundRaw:     soundRaw,
		relatedToTag: relatedToTag,
		messageCode:  messageCode,
	}
}
func (context *SoundPlayer) Play() {
	sdl.QueueAudio(context.audioDevice, context.soundRaw)
	sdl.PauseAudioDevice(context.audioDevice, false)
}
func (context *SoundPlayer) OnUpdate() error {
	return nil
}
func (context *SoundPlayer) OnCollision(other *Element) error {
	return nil
}
func (context *SoundPlayer) OnMessage(message *Message) error {
	if context.messageCode == message.Code {
		if isSingleAndEmpty(context.relatedToTag) || contains(context.relatedToTag, message.RelatedTo[0].Tag) {
			context.Play()
		}
	}
	return nil
}
