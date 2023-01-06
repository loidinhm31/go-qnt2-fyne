package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"go-qn2management/internal/pkg/app/repository"
	"go-qn2management/internal/pkg/app/service"
	"go-qn2management/internal/pkg/app/ui"
	"go-qn2management/internal/pkg/app/ui/render"
)

type AppConfig struct {
	App fyne.App
}

var appConfig AppConfig

func main() {
	mongoLayer := repository.New()

	serviceLayer := service.New(mongoLayer)

	// Create a fyne app
	fyneApp := app.NewWithID("vn.flo.qnt2.preferences")
	appConfig.App = fyneApp

	// Create and size a fyne window
	appMainWindow := fyneApp.NewWindow("QN2 Management")

	// Get user interface
	renderComponent := render.New(serviceLayer)
	uiComponent := ui.New(serviceLayer, renderComponent)

	renderComponent.SetMainWindow(appMainWindow)
	renderComponent.SetAppSize(fyne.Size{Width: 600, Height: 600})

	finalContainer := uiComponent.MakeUI(renderComponent)
	appMainWindow.SetContent(finalContainer)

	// Create Menu
	uiComponent.CreateMenuItems()

	renderConfig := renderComponent.GetRenderConfig()
	renderConfig.MainWindow.Resize(renderConfig.AppSize)
	renderConfig.MainWindow.CenterOnScreen()
	renderConfig.MainWindow.ShowAndRun()

	repository.DeferDisconnect()
}
