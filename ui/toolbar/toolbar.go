package toolbar

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
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
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			t.addItemsDialog()
		}),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			t.refreshSessionsContent()
		}),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {}),
	)

	return toolbar
}

func (t *toolbar) refreshSessionsContent() {
	log.Println("Refreshing...")

	renderConfig := t.render.GetRenderConfig()
	//
	slice := t.render.GetSessionSlice()
	t.coordinator.SetSlice(slice)
	//
	renderConfig.SessionWidget.Refresh()
}

func (t *toolbar) addItemsDialog() dialog.Dialog {
	emptyValidator := func(s string) error {
		if len(s) == 0 {
			return errors.New("empty value, add value for this field")
		}
		return nil
	}

	title := widget.NewEntry()
	title.Validator = emptyValidator

	extension := widget.NewEntry()
	extension.Validator = emptyValidator

	// Create a dialog
	addForm := dialog.NewForm(
		"Add Item",
		"Add",
		"Cancel",
		[]*widget.FormItem{
			{Text: "Title", Widget: title},
			{Text: "Youtube extension", Widget: extension},
		},

		func(valid bool) {
			if valid {

				t.refreshSessionsContent()
			}
		},
		t.render.GetRenderConfig().MainWindow,
	)

	// Size and show the dialog
	addForm.Resize(fyne.Size{Width: 500})
	addForm.Show()

	return addForm
}
