package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/jacksongodsey/golearning/hangman/getword"
)

func main() {
	out := make(chan string)
	app := app.New()
	w := app.NewWindow("Hello World")

	w.SetContent(widget.NewLabel("Hello World"))
	w.Resize(fyne.NewSize(500, 500))
	w.ShowAndRun()
	temp := <-out
	go getword.Getword(out)
	w.SetContent(widget.NewLabeL(temp))
}
