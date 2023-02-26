package args

import (
	"image"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
)

type Args struct {
	Path    textinput.Model
	Width   textinput.Model
	Height  textinput.Model
	Bucket  textinput.Model
	Quality textinput.Model
	err     error
}

func readImage(path string) image.Image {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	image, _, err := image.Decode(file)

	if err != nil {
		log.Fatal(err)
	}

	return image
}
