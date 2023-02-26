package main

import (
	_ "image/jpeg"
	"log"

	ui "squashbucket/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(ui.NewModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

	// resizer := resizer.NewResizer(image, flags.Width, flags.Height)
	// resizedImage := resizer.Resize()

	// fmt.Println(resizedImage)

}
