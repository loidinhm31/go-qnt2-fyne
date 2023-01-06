package table

import (
	"go-qn2management/service"
	"go-qn2management/ui/render"
)

type Tab interface {
	SetSessionSlice(sessionSlice [][]interface{})
}

type tab struct {
	service service.Service
	render  render.Render
}

func New(service service.Service, render render.Render) *tab {
	return &tab{
		service: service,
		render:  render,
	}
}
