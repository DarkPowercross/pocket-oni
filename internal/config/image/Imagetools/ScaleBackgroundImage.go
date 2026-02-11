package Imagetools

import (
	"fmt"
	"image"
	"strings"
)

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
