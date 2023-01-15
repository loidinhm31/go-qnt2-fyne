package table

import (
	"go-qn2management/internal/pkg/app/repository"
	"go-qn2management/internal/pkg/app/service"
	"go-qn2management/internal/pkg/app/ui/render"
)

type Table interface {
	SetSessionProps(sessionMap map[string][]*repository.SessionItem, sessionSlice [][]interface{})
}

type table struct {
	service service.Service
	render  render.Render
}

func New(service service.Service, render render.Render) *table {
	return &table{
		service: service,
		render:  render,
	}
}
