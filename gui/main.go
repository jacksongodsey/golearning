package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	w := app.NewWindow("Hello World")
	button := widget.NewButton("Press me", func() {
		app.Quit()
	})

	w.SetContent(widget.NewLabel("Hello World"))
	w.SetContent(button)
	w.Resize(fyne.NewSize(500, 500))
	w.ShowAndRun()
}
