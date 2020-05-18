package engine

import (
	"fmt"

	"github.com/josnelihurt/simple-go-gaming/engine/util"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

//SDLComponents contains sdl components used in the engine
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
		util.Logger <- fmt.Sprintf("initializing renderer:%v", err)
		return nil, nil, err
	}

	if err := ttf.Init(); err != nil {
		util.Logger <- fmt.Sprintf("initializing ttf:%v", err)
		return nil, nil, err
	}
	return renderer, window, nil
}

//NewSDLComponents creates SDLComponents
func NewSDLComponents(screenWidth, screenHeight int32, gameName string) (context *SDLComponents, err error) {
	context = &SDLComponents{}
	if context.Renderer, context.Window, err = createRenderer(screenWidth, screenHeight, gameName); err != nil {
		util.Logger <- fmt.Sprintf("init error%v:", err)
		return nil, err
	}
	context.AudioDev = openAudioDevice()

	return context, nil
}

//Release all sdl components in the engine
func (context *SDLComponents) Release() {
	context.Renderer.Destroy()
	context.Window.Destroy()
	sdl.CloseAudioDevice(context.AudioDev)
}
