package coordinator

import (
	"go-qn2management/ui/render"
	"go-qn2management/ui/tab"
)

// Coordinator receives config from @Render to distribute function for UI component working together
type Coordinator interface {
	SetSlice(sessionSlice [][]interface{})
}

type coordinator struct {
	render render.Render
	tab    tab.Tab
}

func New(render render.Render, tab tab.Tab) *coordinator {
	return &coordinator{
		render: render,
		tab:    tab,
	}
}

func (c *coordinator) SetSlice(sessionSlice [][]interface{}) {
	c.tab.SetSessionSlice(sessionSlice)
}
