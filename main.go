package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Metronome - Pokémon Randomizer")

	w.Resize(fyne.NewSize(600, 400))

	title := widget.NewLabelWithStyle(
		"Metronome",
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)

	subtitle := widget.NewLabel("A Go rewrite of the Universal Pokémon Randomizer ")

	content := container.NewVBox(
		title,
		subtitle,
		widget.NewSeparator(),
	)

	w.SetContent(container.NewCenter(content))
	w.ShowAndRun()
}
