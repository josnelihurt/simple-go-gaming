package engine

import (
	"fmt"

	"github.com/josnelihurt/simple-go-gaming/engine/util"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type SDLComponents struct {
	Renderer *sdl.Renderer
	Window   *sdl.Window
	AudioDev sdl.AudioDeviceID
	GameName string
}

func openAudioDevice() sdl.AudioDeviceID {
	currenAudioDriver := sdl.GetCurrentAudioDriver()
	util.Logger <- currenAudioDriver
	soundInfo := &sdl.AudioSpec{
		Freq:     44100,
		Format:   32784,
		Channels: 2,
		Silence:  0,
		Samples:  4096,
	}
	dev, err := sdl.OpenAudioDevice("", false, soundInfo, nil, 0)
	if err != nil {
		util.Logger <- fmt.Sprintf("error opeing audio dev:%v", err)
	}
	return dev
}
func createRenderer(screenWidth, screenHeight int32, gameName string) (*sdl.Renderer, *sdl.Window, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		util.Logger <- fmt.Sprintln("initializing SDL:", err)
		panic(err)
	}
	window, err := sdl.CreateWindow(
		gameName,
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_SHOWN)
	if err != nil {
		util.Logger <- fmt.Sprintln("initializing window:", err)
		return nil, nil, err
	}
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		util.Logger <- fmt.Sprintf("inititalizing renderer:%v", err)
		return nil, nil, err
	}

	if err := ttf.Init(); err != nil {
		util.Logger <- fmt.Sprintf("initializing ttf:%v", err)
		return nil, nil, err
	}
	return renderer, window, nil
}

func NewSDLComponents(screenWidth, screenHeight int32, gameName string) (this *SDLComponents, err error) {
	this = &SDLComponents{}
	if this.Renderer, this.Window, err = createRenderer(screenWidth, screenHeight, gameName); err != nil {
		util.Logger <- fmt.Sprintf("init errro%v:", err)
		return nil, err
	}
	this.AudioDev = openAudioDevice()

	return this, nil
}

func (context *SDLComponents) Release() {
	context.Renderer.Destroy()
	context.Window.Destroy()
	sdl.CloseAudioDevice(context.AudioDev)
}
