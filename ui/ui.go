package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"go-qn2management/service"
	"go-qn2management/ui/coordinator"
	"go-qn2management/ui/render"
	"go-qn2management/ui/table"
	"go-qn2management/ui/toolbar"
)

type ui struct {
	service service.Service
	render  render.Render
}

func New(service service.Service, render render.Render) *ui {
	return &ui{
		service: service,
		render:  render,
	}
}

func (u *ui) MakeUI(render render.Render) *fyne.Container {
	// Get app tabs
	tabComponent := table.New(u.service, u.render)
	sessionTabContent := tabComponent.SessionTable()
	//addSessionContent := tabComponent.createSessionTab()

	coordinatorComponent := coordinator.New(render, tabComponent)

	// Get toolbar
	toolbarComponent := toolbar.New(u.service, u.render, coordinatorComponent)
	toolbarRender := toolbarComponent.ToolBar()

	// Add container to window
	label := widget.NewLabel("Toolbox:")
	hBox := container.NewHBox(label, toolbarRender)
	finalContent := container.NewBorder(hBox, nil, nil, nil, sessionTabContent)

	return finalContent
}
