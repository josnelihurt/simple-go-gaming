package game

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/josnelihurt/simple-go-gaming/pkg/engine/util"
)

const (
	resFont             = "fonts/Starjout.ttf"
	resSpriteEnemy      = "sprites/basic_enemy.png"
	resSpriteBullet     = "sprites/bullet.png"
	resSpritePlayer     = "sprites/player.png"
	resSpriteBackground = "sprites/background_space.png"
	resSoundLaser       = "sounds/NFF-laser.wav"
	resSoundExplosion   = "sounds/explosion.wav"
	resMusicBackground  = "sounds/scene.mp3"
)

func LoadResources() {
	util.Logger <- fmt.Sprintf("%v", AssetNames())
	// remember you must run
	// $ go get -u github.com/jteeuwen/go-bindata/...
	// $ go-bindata sprites/... fonts/... sounds/...
	unPack(resFont)
	unPack(resSpriteBackground)
	unPack(resSpriteBullet)
	unPack(resSpriteEnemy)
	unPack(resSpritePlayer)
	unPack(resSoundExplosion)
	unPack(resSoundLaser)
	unPack(resMusicBackground)
}
func unPack(file string) {
	folder, _ := filepath.Split(file)
	os.Mkdir(folder, os.ModePerm)
	if _, err := os.Stat(file); os.IsNotExist(err) {
		fileBytes, erro := Asset(file)
		if erro != nil {
			util.Logger <- fmt.Sprintf("resource %v error %v ", file, erro)
		}
		fileIO, _ := os.Create(file)
		if _, err := fileIO.Write(fileBytes); err == nil {
			util.Logger <- fmt.Sprintf("%v unpacked from resources", file)
		}
		fileIO.Sync()
		fileIO.Close()
	} else {
		util.Logger <- fmt.Sprintf("%v already in fs", file)
	}
}
