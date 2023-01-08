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
	MongoUsername           = "mongodb"   // example, must change
	MongoPassword           = "mongodbpw" // example, must change
	MongoURL                = "mongodb://mongodb:mongodbpw@localhost:27017/?authSource=admin"
	MongoDB                 = "qnt2"
	MongoSessionsCollection = "sessions"
	MongoItemsCollection    = "items"
)
