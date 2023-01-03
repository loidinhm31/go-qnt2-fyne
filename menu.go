package main

import "fyne.io/fyne/v2"

func (app *AppConfig) createMenuItems(w fyne.Window) {
	saveMenuItem := fyne.NewMenuItem("Save", func() {})

	appMenu := fyne.NewMenu("Tools", saveMenuItem)

	menu := fyne.NewMainMenu(appMenu)

	w.SetMainMenu(menu)
}
