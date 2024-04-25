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
	w.SetMainMenu(makeMenu())
	// splitSource := container.NewVBox(
	// 	widget.NewButton("quit", func() {
	// 		a.Quit()
	// 	}),
	// )
	listSource := widget.NewList(
		func() int { return 1 },
		func() fyne.CanvasObject { return widget.NewLabel("a") },
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText("T")
		},
	)
	splitInst := container.NewHSplit(listSource, nil)
	w.SetContent(splitInst)
	// w.SetContent(splitSource)
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
