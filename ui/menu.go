package ui

import (
	"fyne.io/fyne/v2"
)

func (u *ui) CreateMenuItems(w fyne.Window) {
	saveMenuItem := fyne.NewMenuItem("Save", func() {})

	appMenu := fyne.NewMenu("Tools", saveMenuItem)

	menu := fyne.NewMainMenu(appMenu)

	w.SetMainMenu(menu)
}
