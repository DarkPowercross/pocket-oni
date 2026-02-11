package Imagetools

import (
	"image"

	"golang.org/x/image/draw"
)

const characterdir = "internal/characters/"
const character = "oni"
const imagesize = "65x65"

func GenerateGifBackground(gifimage, background image.Image) image.Image {

	result := image.NewRGBA(background.Bounds())

	draw.Draw(result, result.Bounds(), background, image.Point{0, 0}, draw.Src)

	draw.Draw(result, gifimage.Bounds().Add(image.Point{0, 0}), gifimage, image.Point{0, 0}, draw.Over)

	return result

}
