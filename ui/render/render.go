package render

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"go-qn2management/service"
	"log"
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
