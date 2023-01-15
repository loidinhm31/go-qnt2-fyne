package config

import (
	"fyne.io/fyne/v2"
	"go-qn2management/internal/pkg/app/i18n"
)

type AppConfig struct {
	App        fyne.App
	Language   i18n.Language
	MainWindow fyne.Window
	AppSize    fyne.Size
}

var GlobalConfig AppConfig

const (
	MongoUsername           = "root_mongo"       // example, must change
	MongoPassword           = "9ujHM5F7yBYop1I6" // example, must change
	MongoURL                = "mongodb+srv://root_mongo:9ujHM5F7yBYop1I6@cluster0.buv2a7p.mongodb.net/?retryWrites=true&w=majority"
	MongoDB                 = "qnt2"
	MongoSessionsCollection = "sessions"
	MongoItemsCollection    = "items"
)
