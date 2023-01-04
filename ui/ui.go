package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"go-qn2management/service"
	"go-qn2management/ui/coordinator"
	"go-qn2management/ui/render"
	"go-qn2management/ui/tab"
	toolbar "go-qn2management/ui/toolbar"
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
	tabComponent := tab.New(u.service, u.render)
	sessionTabContent := tabComponent.SessionTab()
	addSessionContent := tabComponent.AddSessionTab()

	coordinatorComponent := coordinator.New(render, tabComponent)

	// Get toolbar
	toolbarComponent := toolbar.New(u.service, u.render, coordinatorComponent)
	toolbarRender := toolbarComponent.ToolBar()

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon(
			"Session",
			theme.HomeIcon(),
			sessionTabContent,
		),
		container.NewTabItemWithIcon(
			"Add Session",
			theme.ContentAddIcon(),
			addSessionContent,
		),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	// Add container to window
	finalContent := container.NewVBox(toolbarRender, tabs)

	return finalContent
}
