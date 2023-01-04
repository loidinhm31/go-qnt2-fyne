package tab

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Widget struct {
	SessionSlice [][]interface{}
	TableWidget  *widget.Table
}

var SessionWidget Widget

func (t *tab) SessionTab() *fyne.Container {
	SessionWidget.SessionSlice = t.render.GetSessionSlice()

	SessionWidget.TableWidget = t.getSessionsTable()
	t.render.SetSessionWidget(SessionWidget.TableWidget)

	sessionsContainer := container.NewGridWrap(
		t.render.GetRenderConfig().AppSize,
		container.NewAdaptiveGrid(1, SessionWidget.TableWidget))

	//sessionsContainer := container.NewAdaptiveGrid(1, sessionsTable)

	return sessionsContainer
}

func (t *tab) getSessionsTable() *widget.Table {
	table := widget.NewTable(
		func() (int, int) {
			return len(SessionWidget.SessionSlice), len(SessionWidget.SessionSlice[0])
		},
		func() fyne.CanvasObject {
			ctr := container.NewVBox(widget.NewLabel(""))
			return ctr
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Row == 0 {
				header := widget.NewButton(SessionWidget.SessionSlice[i.Row][i.Col].(string), func() {})
				header.Importance = widget.HighImportance
				o.(*fyne.Container).Objects = []fyne.CanvasObject{
					header,
				}
			} else {
				o.(*fyne.Container).Objects = []fyne.CanvasObject{
					widget.NewLabel(SessionWidget.SessionSlice[i.Row][i.Col].(string)),
				}
			}
		})

	colWidths := []float32{200, 200}
	for i := 0; i < len(colWidths); i++ {
		table.SetColumnWidth(i, colWidths[i])
	}

	return table
}

func (t *tab) SetSessionSlice(sessionSlice [][]interface{}) {
	SessionWidget.SessionSlice = sessionSlice
}
