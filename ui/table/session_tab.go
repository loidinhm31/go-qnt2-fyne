package table

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"go-qn2management/model"
	"log"
)

type Widget struct {
	SessionSlice [][]interface{}
	TableWidget  *widget.Table
}

var SessionWidget Widget

func (t *tab) SessionTable() *widget.Table {
	SessionWidget.SessionSlice = t.render.GetSessionSlice()

	SessionWidget.TableWidget = t.getSessionsTable()
	t.render.SetSessionWidget(SessionWidget.TableWidget)

	return SessionWidget.TableWidget
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
			//if i.Col != 0 { // ignore column 0 - id
			if i.Row == 0 {
				header := widget.NewButton(SessionWidget.SessionSlice[i.Row][i.Col].(string), func() {})
				header.Importance = widget.DangerImportance
				o.(*fyne.Container).Objects = []fyne.CanvasObject{
					header,
				}
			} else {

				// Last cell - put in a button
				if i.Col == (len(SessionWidget.SessionSlice[0]) - 1) {

					if SessionWidget.SessionSlice[i.Row][i.Col] == true {

						w := widget.NewButtonWithIcon("Add", theme.ContentAddIcon(), func() {
							t.addItemDialog(SessionWidget.SessionSlice[i.Row][0].(string))
						})
						w.Importance = widget.HighImportance

						o.(*fyne.Container).Objects = []fyne.CanvasObject{
							w,
						}

					} else {
						o.(*fyne.Container).Objects = []fyne.CanvasObject{
							widget.NewLabel(""),
						}
					}

				} else {
					val := SessionWidget.SessionSlice[i.Row][i.Col].(string)
					l := widget.NewLabel(val)
					//l.Wrapping = fyne.TextWrapBreak

					o.(*fyne.Container).Objects = []fyne.CanvasObject{
						l,
					}
				}
			}
			//}
		})

	colWidths := []float32{190, 350, 90, 300, 140, 100}
	for i := 0; i < len(colWidths); i++ {
		table.SetColumnWidth(i, colWidths[i])
	}

	for i := 1; i < len(SessionWidget.SessionSlice); i++ {
		table.SetRowHeight(i, 50)
	}

	return table
}

func (t *tab) addItemDialog(sessionId string) dialog.Dialog {
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
	addItemForm := dialog.NewForm(
		"Add Item",
		"Add",
		"Cancel",
		[]*widget.FormItem{
			{Text: "Title", Widget: title},
			{Text: "Youtube extension", Widget: extension},
		},

		func(valid bool) {
			if valid {

				err := t.service.AddItem(&model.SessionItemSubmit{
					Title:     title.Text,
					Extension: extension.Text,
					SessionID: sessionId,
				})
				if err != nil {
					return
				}

				// Refresh sessions table
				t.refreshSessionsContent()
			}
		},
		t.render.GetRenderConfig().MainWindow,
	)

	// Size and show the dialog
	addItemForm.Resize(fyne.Size{Width: 500})
	addItemForm.Show()

	return addItemForm
}

func (t *tab) refreshSessionsContent() {
	log.Println("Refreshing...")

	slice := t.render.GetSessionSlice()
	t.SetSessionSlice(slice)

	renderConfig := t.render.GetRenderConfig()
	renderConfig.SessionWidget.Hidden = true
	renderConfig.SessionWidget.Refresh()
	renderConfig.SessionWidget.Hidden = false
}

func (t *tab) SetSessionSlice(sessionSlice [][]interface{}) {
	SessionWidget.SessionSlice = sessionSlice
}
