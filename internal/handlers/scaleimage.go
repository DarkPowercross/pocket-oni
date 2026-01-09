package handlers

import (
	"image"

	"golang.org/x/image/draw"
)

func ScaleImage(src image.Image, w, h int) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, w, h))

	draw.ApproxBiLinear.Scale(
		dst,
		dst.Bounds(),
		src,
		src.Bounds(),
		draw.Over,
		nil,
	)

	return dst
}
