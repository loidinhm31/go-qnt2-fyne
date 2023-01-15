package model

type SessionSubmit struct {
	SessionName  string
	SessionKey   string
	SessionOrder int32
}

type SessionItemSubmit struct {
	Title     string
	Extension string
	SessionID string
}
