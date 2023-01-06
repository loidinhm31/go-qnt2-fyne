package model

type SessionSubmit struct {
	SessionName string
	SessionKey  string
}

type SessionItemSubmit struct {
	Title     string
	Extension string
	SessionID string
}
