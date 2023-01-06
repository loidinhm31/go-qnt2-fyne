package toolbar

import "log"

func (t *toolbar) refreshSessionsContent() {
	log.Println("Refreshing...")

	slice := t.render.GetSessionSlice()
	t.coordinator.SetSlice(slice)

	renderConfig := t.render.GetRenderConfig()
	renderConfig.SessionWidget.Refresh()
}
