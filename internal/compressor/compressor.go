package compressor

import "image"

type Compressor struct {
	Image   image.Image
	Quality int
}
