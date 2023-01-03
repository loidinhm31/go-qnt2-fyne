package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"log"
	"os"
)

type AppConfig struct {
	App      fyne.App
	InfoLog  *log.Logger
	ErrorLog *log.Logger

	MainWindow fyne.Window
	Toolbar    *widget.Toolbar

	SessionContainer *fyne.Container
	TestContainer    *fyne.Container

	SessionName *widget.Entry

	PreviewWidget *widget.RichText
}

var appConfig AppConfig

func main() {
	// Create a fyne app
	fyneApp := app.NewWithID("vn.flo.qnt2.preferences")
	appConfig.App = fyneApp

	// Create loggers
	appConfig.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	appConfig.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Create and size a fyne window
	appConfig.MainWindow = fyneApp.NewWindow("QN2 Management")
	appConfig.MainWindow.Resize(fyne.Size{Width: 1200, Height: 800})
	appConfig.MainWindow.CenterOnScreen()

	// Create Menu
	appConfig.createMenuItems(appConfig.MainWindow)

	// Get user interface
	appConfig.makeUI()

	// Set the content of the window
	//appConfig.MainWindow.SetContent(container.NewHSplit(entry, preview))

	appConfig.MainWindow.ShowAndRun()
}
