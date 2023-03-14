package colorize

import (
	"image"
)

type DecomposedImage struct {
	Red   image.Image
	Green image.Image
	Blue  image.Image
}

func DecomposeGRB(img image.Image) DecomposedImage {
	return DecomposedImage{}
}

func ComposeRGB(dec DecomposedImage) image.Image {
	return nil
}
