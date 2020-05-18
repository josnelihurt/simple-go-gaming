package engine

import (
	"fmt"
	"io"
	"os"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

// It doesn't work on my laptop
// mix.Init(mix.INIT_MP3)
// mix.OpenAudio(44100, //mix.DEFAULT_FREQUENCY,
// 	16, 2, 4096)
// music, err := mix.LoadMUS("sounds/scene.mp3")
// if err != nil {
// 	fmt.Println(err)
// }
// music.Play(-1)

//PlayMusicLoop plays a filename continuously
func PlayMusicLoop(fileName string) error {
	mp3File, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer mp3File.Close()

	decoder, err := mp3.NewDecoder(mp3File)
	if err != nil {
		return err
	}

	player, err := oto.NewPlayer(decoder.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}
	defer player.Close()

	for {
		fmt.Printf("Length: %d[bytes]\n", decoder.Length())
		if _, err := io.Copy(player, decoder); err != nil {
			return err
		}
		decoder.Seek(0, 0)
	}
}
