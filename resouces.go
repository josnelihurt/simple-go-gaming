package main

import (
	"fmt"
	"os"

	"github.com/josnelihurt/simple-go-gaming/engine/util"
)

const (
	resFont             = "fonts/Starjout.ttf"
	resSpriteEnemy      = "sprites/basic_enemy.png"
	resSpriteBullet     = "sprites/bullet.png"
	resSpritePlayer     = "sprites/player.png"
	resSpriteBackground = "sprites/background_space.png"
	resSoundLaser       = "sounds/NFF-laser.wav"
	resSoundExplosion   = "sounds/explosion.wav"
)

func unPackFileFromAsset(folder, filename string) {
	os.Mkdir(folder, os.ModePerm)
	out := fmt.Sprintf("%v/%v", folder, filename)
	if _, err := os.Stat(out); os.IsNotExist(err) {
		file, _ := Asset(out)
		fileIO, _ := os.Create(out)
		if _, err := fileIO.Write(file); err == nil {
			util.Logger <- fmt.Sprintf("%v unpacked from resources", out)
		}
		fileIO.Sync()
		fileIO.Close()
	} else {
		util.Logger <- fmt.Sprintf("%v already in fs", out)
	}
}

func loadResources() {
	// remember you must run
	// $ go get -u github.com/jteeuwen/go-bindata/...
	// $ go-bindata sprites/... fonts/... sounds/...
	unPackFileFromAsset("fonts", "Starjout.ttf")
	unPackFileFromAsset("sprites", "basic_enemy.png")
	unPackFileFromAsset("sprites", "bullet.png")
	unPackFileFromAsset("sprites", "player.png")
	unPackFileFromAsset("sprites", "background_space.png")
	unPackFileFromAsset("sounds", "NFF-laser.wav")
	unPackFileFromAsset("sounds", "explosion.wav")
}
