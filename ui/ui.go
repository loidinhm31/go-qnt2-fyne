package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"go-qn2management/repository"
	"go-qn2management/service"
)

type DrawUI struct {
	Toolbar *widget.Toolbar
}

type UI interface {
	GetAllSessions() ([]*repository.Session, error)
}

type ui struct {
	service service.Service
}

var drawUI DrawUI

func New(service service.Service) *ui {
	return &ui{
		service: service,
	}
}

func (u *ui) MakeUI(mainWindow fyne.Window) {
	// Get toolbar
	toolbar := u.getToolBar()
	drawUI.Toolbar = toolbar

	sessionTabContent := u.sessionTab()
	testTabContent := u.testTab()

	// Get app tabs
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon(
			"Session",
			theme.HomeIcon(),
			sessionTabContent,
		),
		container.NewTabItemWithIcon(
			"Add Session",
			theme.ContentAddIcon(),
			testTabContent,
		),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	// Add container to window
	finalContent := container.NewVBox(toolbar, tabs)

	mainWindow.SetContent(finalContent)
}

//func (app *AppConfig) makeUI() (*widget.Label, *widget.Entry, *widget.Button, *widget.RichText) {
//	outputLabel := widget.NewLabel("Session")
//
//	// Preview
//	preview := widget.NewRichTextWithText("")
//	app.PreviewWidget = preview
//
//	// Input
//	entrySessionName := widget.NewEntry()
//	app.SessionName = entrySessionName
//
//	entrySessionName.OnChanged = preview.ParseMarkdown
//
//	// Button
//	btn := widget.NewButton("Enter", func() {
//		app.SessionName.SetText(entrySessionName.Text)
//	})
//	btn.Importance = widget.HighImportance
//
//	return outputLabel, entrySessionName, btn, preview
//}
