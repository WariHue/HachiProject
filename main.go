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
	newFileItem := fyne.NewMenuItem("New File", nil)
	newWindowItem := fyne.NewMenuItem("New Window", nil)
	openFileItem := fyne.NewMenuItem("Open File", nil)
	openFolderItem := fyne.NewMenuItem("Open Folder", nil)
	file := fyne.NewMenu("File", newFileItem, newWindowItem, fyne.NewMenuItemSeparator(), openFileItem, openFolderItem)
	edit := fyne.NewMenu("Edit", newFileItem, newWindowItem, fyne.NewMenuItemSeparator(), openFileItem)
	main := fyne.NewMainMenu(file, edit)
	return main
}
