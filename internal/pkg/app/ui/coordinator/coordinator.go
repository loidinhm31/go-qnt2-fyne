package coordinator

import (
	"go-qn2management/internal/pkg/app/repository"
	"go-qn2management/internal/pkg/app/ui/render"
	"go-qn2management/internal/pkg/app/ui/table"
)

// Coordinator receives config from @Render to distribute function for UI component working together
type Coordinator interface {
	SetSessionProps(sessionMap map[string][]*repository.SessionItem, sessionSlice [][]interface{})
}

type coordinator struct {
	render render.Render
	tab    table.Table
}

func New(render render.Render, tab table.Table) *coordinator {
	return &coordinator{
		render: render,
		tab:    tab,
	}
}

func (c *coordinator) SetSessionProps(sessionMap map[string][]*repository.SessionItem, sessionSlice [][]interface{}) {
	c.tab.SetSessionProps(sessionMap, sessionSlice)
}
