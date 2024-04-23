package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(1680, 960))
	hello := widget.NewLabel("Hello Fyne!")
	w.SetMainMenu(makeMenu())
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
		widget.NewButton("quit", func() {
			a.Quit()
		}),
	))

	w.ShowAndRun()
}

func makeMenu() *fyne.MainMenu {
	newItem := fyne.NewMenuItem("New", nil)
	checkedItem := fyne.NewMenuItem("Checked", nil)
	checkedItem.Checked = true
	disabledItem := fyne.NewMenuItem("Disabled", nil)
	disabledItem.Disabled = true
	file := fyne.NewMenu("File", newItem, checkedItem, disabledItem)
	main := fyne.NewMainMenu(file)
	return main
}
