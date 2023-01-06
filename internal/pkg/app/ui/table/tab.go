package table

import (
	"go-qn2management/internal/pkg/app/repository"
	"go-qn2management/internal/pkg/app/service"
	"go-qn2management/internal/pkg/app/ui/render"
)

type Tab interface {
	SetSessionProps(sessionMap map[string][]*repository.SessionItem, sessionSlice [][]interface{})
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
