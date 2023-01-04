package tab

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"go-qn2management/model"
	"go-qn2management/service"
	"golang.org/x/image/colornames"
	"log"
)

func (t *tab) AddSessionTab() *fyne.Container {
	sessionForm := t.getSessionForm()

	addSessionContainer := container.NewAdaptiveGrid(1,
		sessionForm,
	)

	return addSessionContainer
}

func (t *tab) getSessionForm() *widget.Form {
	emptyValidator := func(s string) error {
		if len(s) == 0 {
			return errors.New("empty value, add value for this field")
		}
		return nil
	}

	sessionNameValue := widget.NewEntry()
	sessionNameValue.Validator = emptyValidator

	sessionKeyValue := widget.NewEntry()
	sessionKeyValue.Validator = emptyValidator

	resultText := canvas.NewText("Create Session successfully", colornames.Green)
	resultText.Hidden = true

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Session Name", Widget: sessionNameValue},
			{Text: "Session Key", Widget: sessionKeyValue},
			{Text: "", Widget: resultText},
		},
		OnSubmit: func() { // optional, handle form submission
			sessionSubmit := model.SessionSubmit{
				SessionName: sessionNameValue.Text,
				SessionKey:  sessionKeyValue.Text,
			}
			err := handleSubmit(t.service, &sessionSubmit)
			if err != nil {
				log.Println(err)
				return
			}

			resultText.Hidden = false
			sessionNameValue.SetText("")
			sessionKeyValue.SetText("")

			t.refreshSessionsContent()
		},
	}
	return form
}

func handleSubmit(service service.Service, sessionSubmit *model.SessionSubmit) error {
	// Get sessions from mongo
	err := service.AddSession(sessionSubmit)
	if err != nil {
		return err
	}
	return nil
}

func getImage() *canvas.Image {
	img := canvas.NewImageFromResource(resourceUnreachablePng)
	img.SetMinSize(fyne.Size{
		Width:  770,
		Height: 410,
	})
	img.FillMode = canvas.ImageFillOriginal

	return img
}

func (t *tab) refreshSessionsContent() {
	log.Println("Refreshing...")

	renderConfig := t.render.GetRenderConfig()
	//
	slice := t.render.GetSessionSlice()
	t.SetSessionSlice(slice)
	//
	renderConfig.SessionWidget.Refresh()
}
