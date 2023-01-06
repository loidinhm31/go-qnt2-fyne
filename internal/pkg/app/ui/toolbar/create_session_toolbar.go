package toolbar

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"go-qn2management/internal/pkg/app/config"
	"go-qn2management/internal/pkg/app/i18n"
	"go-qn2management/internal/pkg/app/model"
	"log"
)

func (t *toolbar) addSessionDialog() dialog.Dialog {
	emptyValidator := func(s string) error {
		if len(s) == 0 {
			return errors.New(i18n.Messages["empty_value"][config.GlobalConfig.Language])
		}
		return nil
	}

	sessionNameValue := widget.NewEntry()
	sessionNameValue.Validator = emptyValidator

	sessionKeyValue := widget.NewEntry()
	sessionKeyValue.Validator = emptyValidator

	addSessionForm := dialog.NewForm(
		i18n.Messages["add_session"][config.GlobalConfig.Language],
		i18n.Messages["add"][config.GlobalConfig.Language],
		i18n.Messages["cancel"][config.GlobalConfig.Language],
		[]*widget.FormItem{
			{Text: i18n.Messages["session_name"][config.GlobalConfig.Language], Widget: sessionNameValue},
			{Text: i18n.Messages["session_key"][config.GlobalConfig.Language], Widget: sessionKeyValue},
		},

		func(valid bool) { // optional, handle sessionForm submission
			if valid {
				sessionSubmit := model.SessionSubmit{
					SessionName: sessionNameValue.Text,
					SessionKey:  sessionKeyValue.Text,
				}
				err := t.handleSubmit(&sessionSubmit)
				if err != nil {
					log.Println(err)
					return
				}

				t.refreshSessionsContent()
			}
		},
		config.GlobalConfig.MainWindow,
	)
	// Size and show the dialog
	addSessionForm.Resize(fyne.Size{Width: 500})
	addSessionForm.Show()

	return addSessionForm
}

func (t *toolbar) handleSubmit(sessionSubmit *model.SessionSubmit) error {
	// Get sessions from mongo
	err := t.service.AddSession(sessionSubmit)
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
