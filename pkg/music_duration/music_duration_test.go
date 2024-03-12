package music_duration

import (
	"log"
	"os"
	"testing"
)

func TestMusicDuration(t *testing.T) {
	wd, _ := os.Getwd()
	rootPath := wd + "/../../"
	filePath := rootPath + "/tmp/music.mp3"

	duration, err := MusicDuration(filePath)
	if err != nil {
		log.Fatal(err)
	}

	t.Log(duration)
}
