package table

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"go-qn2management/internal/pkg/app/config"
	"go-qn2management/internal/pkg/app/i18n"
	"go-qn2management/internal/pkg/app/model"
	"go-qn2management/internal/pkg/app/repository"
	"go-qn2management/internal/pkg/app/util"
	"log"
	"strconv"
)

type Widget struct {
	SessionMap   map[string][]*repository.SessionItem
	SessionSlice [][]interface{}
	TableWidget  *widget.Table
}

var SessionWidget Widget

func (t *table) SessionTable() *widget.Table {
	SessionWidget.SessionMap, SessionWidget.SessionSlice = t.render.GetSessionProps()

	SessionWidget.TableWidget = t.getSessionsTable()
	t.render.SetSessionWidget(SessionWidget.TableWidget)

	return SessionWidget.TableWidget
}

func (t *table) getSessionsTable() *widget.Table {
	table := widget.NewTable(
		func() (int, int) {
			return len(SessionWidget.SessionSlice), len(SessionWidget.SessionSlice[0])
		},
		func() fyne.CanvasObject {
			ctr := container.NewVBox(widget.NewLabel(""))
			return ctr
		},
		func(cellID widget.TableCellID, o fyne.CanvasObject) {
			sessionId := SessionWidget.SessionSlice[cellID.Row][0].(string)

			addItemColIndex := len(SessionWidget.SessionSlice[0]) - 2
			delItemColIndex := len(SessionWidget.SessionSlice[0]) - 1

			// Check index in slice
			itemSlice := SessionWidget.SessionMap[sessionId]
			itemIndex := util.SliceIndex(len(itemSlice), func(i int) bool {
				return itemSlice[i].ID == SessionWidget.SessionSlice[cellID.Row][1].(string)
			})

			if cellID.Row == 0 {
				header := widget.NewButton(SessionWidget.SessionSlice[cellID.Row][cellID.Col].(string), func() {})
				header.Importance = widget.DangerImportance
				o.(*fyne.Container).Objects = []fyne.CanvasObject{
					header,
				}
			} else {
				if cellID.Col == 0 {
					if SessionWidget.SessionSlice[cellID.Row][addItemColIndex] == true {
						o.(*fyne.Container).Objects = []fyne.CanvasObject{
							widget.NewLabel(SessionWidget.SessionSlice[cellID.Row][delItemColIndex].(string)), // replace session id by index
						}
					} else {
						o.(*fyne.Container).Objects = []fyne.CanvasObject{
							widget.NewLabel(""),
						}
					}
				} else if cellID.Col == 1 {
					o.(*fyne.Container).Objects = []fyne.CanvasObject{
						widget.NewLabel(strconv.Itoa(itemIndex + 1)), // replace session id by item index
					}
				} else if cellID.Col == 3 { // put in a delete session button
					if SessionWidget.SessionSlice[cellID.Row][addItemColIndex] == true {
						w := widget.NewButtonWithIcon(
							i18n.Messages["delete_session"][config.GlobalConfig.Language],
							theme.ContentRemoveIcon(),
							func() {
								t.deleteSessionDialog(sessionId)
							},
						)
						w.Importance = widget.DangerImportance

						o.(*fyne.Container).Objects = []fyne.CanvasObject{
							w,
						}

					} else {
						o.(*fyne.Container).Objects = []fyne.CanvasObject{
							widget.NewLabel(""),
						}
					}
				} else if cellID.Col == (delItemColIndex) { // Last cell - put in a delete item button
					if itemIndex != -1 {
						w := widget.NewButtonWithIcon(
							i18n.Messages["delete_item"][config.GlobalConfig.Language],
							theme.DeleteIcon(),
							func() {
								t.deleteItemDialog(SessionWidget.SessionSlice[cellID.Row][1].(string)) // item id
							},
						)
						w.Importance = widget.DangerImportance

						o.(*fyne.Container).Objects = []fyne.CanvasObject{
							w,
						}
					} else {
						o.(*fyne.Container).Objects = []fyne.CanvasObject{
							widget.NewLabel(""),
						}
					}
				} else if cellID.Col == addItemColIndex { // Last n - 1 cell - put in an add item button
					if SessionWidget.SessionSlice[cellID.Row][cellID.Col] == true {
						w := widget.NewButtonWithIcon(
							i18n.Messages["add"][config.GlobalConfig.Language],
							theme.ContentAddIcon(),
							func() {
								t.addItemDialog(sessionId)
							},
						)
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
					val := SessionWidget.SessionSlice[cellID.Row][cellID.Col].(string)
					l := widget.NewLabel(val)
					if len(val) > 50 {
						l.Wrapping = fyne.TextWrapWord
					}
					o.(*fyne.Container).Objects = []fyne.CanvasObject{
						l,
					}
				}
			}
		})

	colWidths := []float32{
		35,
		35,
		50,
		95,
		330,
		100,
		350,
		100,
		95,
		95,
	}
	for i := 0; i < len(colWidths); i++ {
		table.SetColumnWidth(i, colWidths[i])
	}

	for i := 1; i < len(SessionWidget.SessionSlice); i++ {
		table.SetRowHeight(i, 50)
	}
	return table
}

func (t *table) addItemDialog(sessionId string) dialog.Dialog {
	emptyValidator := func(s string) error {
		if len(s) == 0 {
			return errors.New(i18n.Messages["empty_value"][config.GlobalConfig.Language])
		}
		return nil
	}

	title := widget.NewEntry()
	title.Validator = emptyValidator

	extension := widget.NewEntry()
	extension.Validator = emptyValidator

	// Create a dialog
	addItemForm := dialog.NewForm(
		i18n.Messages["add_item"][config.GlobalConfig.Language],
		i18n.Messages["add"][config.GlobalConfig.Language],
		i18n.Messages["cancel"][config.GlobalConfig.Language],
		[]*widget.FormItem{
			{Text: i18n.Messages["title"][config.GlobalConfig.Language], Widget: title},
			{Text: i18n.Messages["youtube_extension"][config.GlobalConfig.Language], Widget: extension},
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

				dialog.ShowInformation(
					i18n.Messages["notify_message"][config.GlobalConfig.Language],
					i18n.Messages["update_successful"][config.GlobalConfig.Language],
					config.GlobalConfig.MainWindow)
				// Refresh sessions table
				t.refreshSessionsContent()
			}
		},
		config.GlobalConfig.MainWindow,
	)

	// Size and show the dialog
	addItemForm.Resize(fyne.Size{Width: 500})
	addItemForm.Show()

	return addItemForm
}

func (t *table) refreshSessionsContent() {
	log.Println("Refreshing...")

	// Update sessions
	sessionMap, sessionSlice := t.render.GetSessionProps()
	t.SetSessionProps(sessionMap, sessionSlice)

	renderConfig := t.render.GetRenderConfig()
	renderConfig.SessionWidget.Refresh()
}

func (t *table) SetSessionProps(sessionMap map[string][]*repository.SessionItem, sessionSlice [][]interface{}) {
	SessionWidget.SessionMap = sessionMap
	SessionWidget.SessionSlice = sessionSlice
}

func (t *table) deleteSessionDialog(id string) {
	label := widget.NewLabel(i18n.Messages["delete_session_question"][config.GlobalConfig.Language])

	addItemForm := dialog.NewForm(
		i18n.Messages["delete_session"][config.GlobalConfig.Language],
		i18n.Messages["yes"][config.GlobalConfig.Language],
		i18n.Messages["no"][config.GlobalConfig.Language],
		[]*widget.FormItem{
			{Text: "", Widget: label},
		},

		func(valid bool) {
			if valid {
				err := t.service.RemoveSessionById(id, SessionWidget.SessionMap)
				if err != nil {
					dialog.ShowInformation(
						i18n.Messages["warning_message"][config.GlobalConfig.Language],
						i18n.Messages["invalid_session"][config.GlobalConfig.Language],
						config.GlobalConfig.MainWindow)
					return
				}

				dialog.ShowInformation(
					i18n.Messages["notify_message"][config.GlobalConfig.Language],
					i18n.Messages["update_successful"][config.GlobalConfig.Language],
					config.GlobalConfig.MainWindow)
				// Refresh session table
				t.refreshSessionsContent()
			}
		},
		config.GlobalConfig.MainWindow,
	)

	// Size and show the dialog
	addItemForm.Resize(fyne.Size{Width: 500})
	addItemForm.Show()
}

func (t *table) deleteItemDialog(id string) {
	label := widget.NewLabel(i18n.Messages["delete_item_question"][config.GlobalConfig.Language])

	addItemForm := dialog.NewForm(
		i18n.Messages["delete_item"][config.GlobalConfig.Language],
		i18n.Messages["yes"][config.GlobalConfig.Language],
		i18n.Messages["no"][config.GlobalConfig.Language],
		[]*widget.FormItem{
			{Text: "", Widget: label},
		},

		func(valid bool) {
			if valid {
				err := t.service.RemoveItemById(id)
				if err != nil {
					return
				}
				dialog.ShowInformation(
					i18n.Messages["notify_message"][config.GlobalConfig.Language],
					i18n.Messages["update_successful"][config.GlobalConfig.Language],
					config.GlobalConfig.MainWindow)
				// Refresh sessions table
				t.refreshSessionsContent()
			}
		},
		config.GlobalConfig.MainWindow,
	)

	// Size and show the dialog
	addItemForm.Resize(fyne.Size{Width: 500})
	addItemForm.Show()
}
