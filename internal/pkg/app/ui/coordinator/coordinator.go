package coordinator

import (
	"go-qn2management/internal/pkg/app/ui/render"
	"go-qn2management/internal/pkg/app/ui/table"
)

// Coordinator receives config from @Render to distribute function for UI component working together
type Coordinator interface {
	SetSlice(sessionSlice [][]interface{})
}

type coordinator struct {
	render render.Render
	tab    table.Tab
}

func New(render render.Render, tab table.Tab) *coordinator {
	return &coordinator{
		render: render,
		tab:    tab,
	}
}

func (c *coordinator) SetSlice(sessionSlice [][]interface{}) {
	c.tab.SetSessionSlice(sessionSlice)
}
