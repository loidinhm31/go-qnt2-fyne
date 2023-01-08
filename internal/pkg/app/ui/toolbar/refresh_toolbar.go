package toolbar

import "log"

func (t *toolbar) refreshSessionsContent() {
	log.Println("Refreshing...")

	sessionMap, sessionSlice := t.render.GetSessionProps()
	t.coordinator.SetSessionProps(sessionMap, sessionSlice)

	renderConfig := t.render.GetRenderConfig()
	renderConfig.SessionWidget.Refresh()
}
