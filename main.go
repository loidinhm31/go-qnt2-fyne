package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"go-qn2management/repository"
	"go-qn2management/service"
	"go-qn2management/ui"
)

type AppConfig struct {
	App fyne.App

	MainWindow fyne.Window

	SessionContainer *fyne.Container
	TestContainer    *fyne.Container

	ServiceLayer *service.Service

	SessionsTable *widget.Table
}

var appConfig AppConfig

func main() {
	mongoLayer := repository.New()

	serviceLayer := service.New(mongoLayer)

	// Create a fyne app
	fyneApp := app.NewWithID("vn.flo.qnt2.preferences")
	appConfig.App = fyneApp

	// Create and size a fyne window
	appConfig.MainWindow = fyneApp.NewWindow("QN2 Management")
	appConfig.MainWindow.Resize(fyne.Size{Width: 1200, Height: 800})
	appConfig.MainWindow.CenterOnScreen()

	// Get user interface
	uiComponent := ui.New(serviceLayer)
	uiComponent.MakeUI(appConfig.MainWindow)

	// Create Menu
	uiComponent.CreateMenuItems(appConfig.MainWindow)

	appConfig.MainWindow.ShowAndRun()

	repository.DeferDisconnect()
}
