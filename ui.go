package main

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

func (app *AppConfig) makeUI() {

	grey := color.NRGBA{R: 155, G: 155, B: 155, A: 255}

	a := canvas.NewText("a", grey)
	b := canvas.NewText("b", grey)
	c := canvas.NewText("c", grey)

	// Put information into container
	sessionContent := container.NewGridWithColumns(4,
		a,
		b,
		c,
	)
	app.SessionContainer = sessionContent

	// Get toolbar
	toolbar := app.getToolBar(app.MainWindow)
	app.Toolbar = toolbar

	testTabContent := app.testTab()

	// Get app tabs
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon(
			"Session",
			theme.HomeIcon(),
			testTabContent),
		container.NewTabItemWithIcon("Add Session", theme.ContentAddIcon(), canvas.NewText("Add session here", nil)),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	// Add container to window
	finalContent := container.NewVBox(sessionContent, toolbar, tabs)

	app.MainWindow.SetContent(finalContent)
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
