package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed        = 2
	playerSize         = 105
	playerShotCooldown = time.Millisecond * 150
)

type player struct {
	texture  *sdl.Texture
	x, y     float64
	lastShot time.Time
}

func newPlayer(renderer *sdl.Renderer) (p player, err error) {
	p.texture = textureFromBMP(renderer, "sprites/player.bmp")
	p.x = screenWidth / 2
	p.y = screenHeight - playerSize

	return p, nil
}

func (p *player) draw(renderer *sdl.Renderer) {
	x := p.x - playerSize/2.0
	y := p.y - playerSize/2.0
	renderer.Copy(p.texture,
		&sdl.Rect{X: 0, Y: 0, W: 200, H: 200},
		&sdl.Rect{X: int32(x), Y: int32(y), W: playerSize, H: playerSize})
}
func (p *player) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		if p.x-(playerSize/2.0) > 0 {
			p.x = p.x - playerSpeed
		}
	}
	if keys[sdl.SCANCODE_RIGHT] == 1 {
		if p.x+(playerSize/2.0) < screenWidth {
			p.x = p.x + playerSpeed
		}
	}
	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(p.lastShot) > playerShotCooldown {
			p.shoot(p.x+25, p.y-40)
			p.shoot(p.x-25, p.y-40)
			p.lastShot = time.Now()
		}
	}
}
func (p *player) shoot(x, y float64) {
	if b, ok := bulletFromPool(); ok {
		b.active = true
		b.x = x
		b.y = y
		b.angle = 270 * (math.Pi / 180)
	}
}
func (p *player) destroy() {
	p.texture.Destroy()
}
