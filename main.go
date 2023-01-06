package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"go-qn2management/internal/pkg/app/config"
	"go-qn2management/internal/pkg/app/i18n"
	"go-qn2management/internal/pkg/app/repository"
	"go-qn2management/internal/pkg/app/service"
	"go-qn2management/internal/pkg/app/ui"
	"go-qn2management/internal/pkg/app/ui/render"
)

func main() {
	mongoLayer := repository.New()

	serviceLayer := service.New(mongoLayer)

	// Default language
	config.GlobalConfig.Language = i18n.Vietnam

	// Create a fyne app
	fyneApp := app.NewWithID("vn.flo.qnt2.preferences")
	config.GlobalConfig.App = fyneApp

	// Create and size a fyne window
	appMainWindow := fyneApp.NewWindow(i18n.Messages["qn2_management"][config.GlobalConfig.Language])
	config.GlobalConfig.MainWindow = appMainWindow
	config.GlobalConfig.AppSize = fyne.Size{Width: 600, Height: 600}

	// Get user interface
	renderComponent := render.New(serviceLayer)
	uiComponent := ui.New(serviceLayer, renderComponent)

	finalContainer := uiComponent.MakeUI(renderComponent)
	appMainWindow.SetContent(finalContainer)

	// Create Menu
	uiComponent.CreateMenuItems()

	config.GlobalConfig.MainWindow.Resize(config.GlobalConfig.AppSize)
	config.GlobalConfig.MainWindow.CenterOnScreen()
	config.GlobalConfig.MainWindow.ShowAndRun()

	repository.DeferDisconnect()
}
