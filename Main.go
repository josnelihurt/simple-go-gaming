package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 800
)

func textureFromBMP(renderer *sdl.Renderer, filename string) (texture *sdl.Texture) {
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		panic(fmt.Errorf("loading %v: %v", filename, err))
	}
	defer img.Free()

	texture, err = renderer.CreateTextureFromSurface(img)
	if err != nil {

		panic(fmt.Errorf("creating  basic enemy texture: %v", err))
	}
	return texture
}

func main() {
	fmt.Println("Starting up..")
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL:", err)
		panic(err)
	}
	window, err := sdl.CreateWindow(
		"Demo game 1",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("initializing window:", err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("inititalizing renderer:", err)
	}
	defer renderer.Destroy()

	player1, err := newPlayer(renderer)
	if err != nil {
		fmt.Println("creating player:", err)
	}
	defer player1.destroy()

	var enemies []basicEnemy
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			x := (float64(i)/5)*screenWidth + (basicEnemySize / 2.0) + 30
			y := float64(j)*basicEnemySize + (basicEnemySize / 2.0)

			enemy, err := newBasicEnemy(renderer, x, y)
			if err != nil {
				fmt.Println("creating enemy:", err)
				return
			}
			defer enemy.destroy()
			enemies = append(enemies, enemy)
		}
	}
	initBulletPool(renderer)

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				fmt.Println("exiting:")
				return
			}
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()
		player1.draw(renderer)
		player1.update()

		for _, enemy := range enemies {
			enemy.draw(renderer)
		}

		for _, b := range bulletPool {
			b.draw(renderer)
			b.update()
		}

		renderer.Present()
	}

}
