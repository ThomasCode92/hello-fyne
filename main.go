package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	w.SetContent(widget.NewLabel("Hello, world!"))
	w.ShowAndRun()

	tidy()
}

func tidy() {
	fmt.Println("would tidy up")
}
