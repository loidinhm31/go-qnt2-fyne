package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"go-qn2management/service"
	"log"
)

func (u *ui) sessionTab() *fyne.Container {

	sessionsTable := u.getSessionsTable()

	sessionsContainer := container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		container.NewAdaptiveGrid(1, sessionsTable),
	)
	return sessionsContainer
}

func (u *ui) getSessionsTable() *widget.Table {
	slice := getSessionSlice(u.service)

	t := widget.NewTable(
		func() (int, int) {
			return len(slice), len(slice[0])
		},
		func() fyne.CanvasObject {
			ctr := container.NewVBox(widget.NewLabel(""))
			return ctr
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {

			o.(*fyne.Container).Objects = []fyne.CanvasObject{
				widget.NewLabel(slice[i.Row][i.Col].(string)),
			}
		})

	colWidths := []float32{200, 200}
	for i := 0; i < len(colWidths); i++ {
		t.SetColumnWidth(i, colWidths[i])
	}

	return t
}

func getSessionSlice(service service.Service) [][]interface{} {
	var slice [][]interface{}

	sessions, err := service.GetAllSessions()
	if err != nil {
		log.Println(err)
	}

	slice = append(slice, []interface{}{"Session Name", "Session Key"})

	for _, x := range sessions {
		var currentRow []interface{}

		currentRow = append(currentRow, x.SessionName)
		currentRow = append(currentRow, x.SessionKey)

		slice = append(slice, currentRow)
	}
	return slice
}
