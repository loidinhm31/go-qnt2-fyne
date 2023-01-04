package render

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"go-qn2management/repository"
	"go-qn2management/service"
	"log"
	"sort"
)

// Config contains all UI configs, and distribute service for components working together
type Config struct {
	MainWindow    fyne.Window
	AppSize       fyne.Size
	Toolbar       *widget.Toolbar
	SessionWidget *widget.Table
}

type Render interface {
	GetRenderConfig() Config
	SetMainWindow(mainWindow fyne.Window)
	SetAppSize(size fyne.Size)
	SetSessionWidget(sessionWidget *widget.Table)
	GetSessionSlice() [][]interface{}
}

type render struct {
	service service.Service
}

type Item struct {
	sessionItem repository.SessionItem
}

func New(service service.Service) *render {
	return &render{
		service: service,
	}
}

var DrawConfig Config

func (r *render) GetRenderConfig() Config {
	return DrawConfig
}

func (r *render) SetMainWindow(mainWindow fyne.Window) {
	DrawConfig.MainWindow = mainWindow
}

func (r *render) SetAppSize(size fyne.Size) {
	DrawConfig.AppSize = size
}

func (r *render) SetSessionWidget(widgetTable *widget.Table) {
	DrawConfig.SessionWidget = widgetTable
}

func (r *render) GetSessionSlice() [][]interface{} {
	var slice [][]interface{}

	// Get sessions from mongo
	sessions, err := r.service.GetAllSessions()
	if err != nil {
		log.Println(err)
		return nil
	}
	sort.Slice(sessions, func(i, j int) bool {
		return sessions[i].CreatedAt.After(sessions[j].CreatedAt)
	})

	// Get items for session from mongo
	items, err := r.service.GetAllItems()
	if err != nil {
		log.Println(err)
		return nil
	}

	sessionMap := make(map[string][]*repository.SessionItem)

	slice = append(slice, []interface{}{"ID", "Session Name", "Session Key", "Title", "Youtube Extension", "Add Item?"})

	for _, session := range sessions {
		sessionMap[session.ID] = []*repository.SessionItem{}

		for _, item := range items {
			if session.ID == item.SessionID {
				sessionMap[session.ID] = append(sessionMap[session.ID], item)
			}
		}
	}

	for _, session := range sessions {
		if len(sessionMap[session.ID]) > 0 {
			for index, item := range sessionMap[session.ID] {
				var currentRow []interface{}

				currentRow = append(currentRow, session.ID)
				currentRow = append(currentRow, session.SessionName)
				currentRow = append(currentRow, session.SessionKey)
				currentRow = append(currentRow, item.Title)
				currentRow = append(currentRow, item.Extension)

				if index == 0 {
					currentRow = append(currentRow, true)
				} else {
					currentRow = append(currentRow, false)

				}
				slice = append(slice, currentRow)
			}
		} else {
			var currentRow []interface{}

			currentRow = append(currentRow, session.ID)
			currentRow = append(currentRow, session.SessionName)
			currentRow = append(currentRow, session.SessionKey)
			currentRow = append(currentRow, "")
			currentRow = append(currentRow, "")
			currentRow = append(currentRow, true)
			slice = append(slice, currentRow)
		}
	}

	return slice
}
