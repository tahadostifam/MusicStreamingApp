package music_duration

import (
	"io"
	"os"
	"time"

	"github.com/tcolgate/mp3"
)

func MusicDuration(mp3FilePath string) (*time.Duration, error) {
	file, err := os.Open(mp3FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := mp3.NewDecoder(file)
	totalDuration := 0.0

	var frame mp3.Frame
	var skipped int
	for {
		if err := decoder.Decode(&frame, &skipped); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		totalDuration += frame.Duration().Seconds()
	}

	duration := time.Second * time.Duration(totalDuration)

	return &duration, nil
}
