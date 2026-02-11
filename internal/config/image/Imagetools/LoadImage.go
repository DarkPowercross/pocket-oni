package Imagetools

import (
	"image"
	"image/png"
	"os"
)

func LoadImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	image, err := png.Decode(f)
	if err != nil {
		return nil, err
	}

	return image, nil

}
