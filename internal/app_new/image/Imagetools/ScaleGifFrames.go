package Imagetools

import (
	"fmt"
	"image"
	"strings"
)

func ScaleGifFrames(gifname string) ([]image.Image, []int) {
	gifname = strings.ToLower(gifname)
	frames, delay, height, width, err := LoadGIF(characterdir + character + "/" + imagesize + "/" + gifname)
	if err != nil {
		fmt.Println("ending application due to the following error: ", err)
		return nil, nil
	}

	scaledFrames := make([]image.Image, len(frames))
	for i, frame := range frames {
		scaledFrames[i] = ScaleImage(frame, width, height)
	}

	return scaledFrames, delay
}
