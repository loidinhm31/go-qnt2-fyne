package main

import (
	"fyne.io/fyne/v2"
	"net/http"
)

type SessionManagement struct {
	Session Session
	Client  *http.Client
}

type Session struct {
	SessionName string `json:"session_name"`
}

//func (m *SessionManagement) GetSession() (*Session, error) {
//	if m.Client == nil {
//		m.Client = &http.Client{}
//	}
//	client := m.Client
//
//	url := mongoUrl
//
//	req, _ := http.NewRequest("GET", url, nil)
//
//	resp, err := client.Do(req)
//	if err != nil {
//		log.Println("error", err)
//		return nil, err
//	}
//	defer resp.Body.Close()
//
//	body, err := io.ReadAll(resp.Body)
//
//	session := Session{}
//
//	err = json.Unmarshal(body, session)
//
//	return &session, nil
//}

func (app *AppConfig) refreshSessionContent() {
	app.InfoLog.Print("Refreshing session content")

	app.SessionContainer.Objects = []fyne.CanvasObject{} // TODO object here
	app.SessionContainer.Refresh()
}
