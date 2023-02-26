package resizer

import (
	"image"

	"github.com/disintegration/imaging"
)

type Resizer struct {
	image  image.Image
	width  int
	height int
}

func NewResizer(image image.Image, width int, height int) *Resizer {
	resizer := Resizer{
		image:  image,
		width:  width,
		height: height,
	}

	return &resizer
}

func (r *Resizer) Resize() image.Image {
	image := imaging.Resize(r.image, r.width, r.height, imaging.Lanczos)

	return image
}
