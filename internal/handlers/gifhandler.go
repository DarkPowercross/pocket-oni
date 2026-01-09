package handlers

import (
	"fmt"
	"image"
	"image/gif"
	"image/png"
	"os"
	"strings"

	"golang.org/x/image/draw"
)

const characterdir = "internal/characters/"
const character = "oni"
const imagesize = "65x65"

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

func ScaleFrames(gifname string) ([]image.Image, []int) {
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

func ScalebackgroundImage(imagename string) image.Image {
	imagename = strings.ToLower(imagename)
	Image, err := LoadImage(characterdir + "" + "/backgrounds/" + imagename)
	if err != nil {
		fmt.Println("ending application due to the following error: ", err)
		return nil
	}

	scaledFrame := ScaleImage(Image, 100, 100)

	return scaledFrame
}

func GenerateGifBackground(gifimage, background image.Image) image.Image {

	result := image.NewRGBA(background.Bounds())

	draw.Draw(result, result.Bounds(), background, image.Point{0, 0}, draw.Src)

	draw.Draw(result, gifimage.Bounds().Add(image.Point{20, 30}), gifimage, image.Point{0, 0}, draw.Over)

	return result

}
