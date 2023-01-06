package toolbar

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"go-qn2management/service"
	"go-qn2management/ui/coordinator"
	"go-qn2management/ui/render"
)

type toolbar struct {
	service     service.Service
	render      render.Render
	coordinator coordinator.Coordinator
}

func New(service service.Service, render render.Render, coordinator coordinator.Coordinator) *toolbar {
	return &toolbar{
		service:     service,
		render:      render,
		coordinator: coordinator,
	}
}

func (t *toolbar) ToolBar() *widget.Toolbar {
	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			t.refreshSessionsContent()
		}),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			t.addSessionDialog()
		}),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {}),
	)
	return toolbar
}
