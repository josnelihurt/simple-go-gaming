package game

import (
	"github.com/josnelihurt/simple-go-gaming/pkg/engine"
	"github.com/veandco/go-sdl2/sdl"
)

var bulletPool []*engine.Element

func initBulletPool(renderer *sdl.Renderer) []*engine.Element {
	for i := 0; i < 30; i++ {
		b := newBullet(renderer)
		bulletPool = append(bulletPool, b)
	}
	return bulletPool
}

func bulletFromPool() (*engine.Element, bool) {
	for _, b := range bulletPool {
		if !b.Active {
			return b, true
		}
	}
	return nil, false
}
