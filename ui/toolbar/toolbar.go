package toolbar

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"go-qn2management/service"
	"go-qn2management/ui/coordinator"
	"go-qn2management/ui/render"
	"log"
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
		widget.NewToolbarAction(theme.SettingsIcon(), func() {}),
	)

	return toolbar
}

func (t *toolbar) refreshSessionsContent() {
	log.Println("Refreshing...")

	slice := t.render.GetSessionSlice()
	t.coordinator.SetSlice(slice)

	renderConfig := t.render.GetRenderConfig()
	renderConfig.SessionWidget.Refresh()
}
