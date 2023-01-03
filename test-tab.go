package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func (app *AppConfig) testTab() *fyne.Container {
	testImage := app.getImage()

	testContainer := container.NewVBox(testImage)
	app.TestContainer = testContainer

	return testContainer
}

func (app *AppConfig) getImage() *canvas.Image {
	img := canvas.NewImageFromResource(resourceUnreachablePng)
	img.SetMinSize(fyne.Size{
		Width:  770,
		Height: 410,
	})
	img.FillMode = canvas.ImageFillOriginal

	return img
}
