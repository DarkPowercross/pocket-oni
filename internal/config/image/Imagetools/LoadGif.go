package Imagetools

import (
	"image"
	"image/gif"
	"os"
)

func LoadGIF(path string) ([]image.Image, []int, int, int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, 0, 0, err
	}
	defer f.Close()

	g, err := gif.DecodeAll(f)
	if err != nil {
		return nil, nil, 0, 0, err
	}

	frames := make([]image.Image, len(g.Image))
	for i, img := range g.Image {
		frames[i] = img
	}

	return frames, g.Delay, g.Config.Height, g.Config.Width, nil

}
