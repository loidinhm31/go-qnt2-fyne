package service

import (
	"go-qn2management/internal/pkg/app/model"
	"go-qn2management/internal/pkg/app/repository"
)

type Service interface {
	GetAllSessions() ([]*repository.Session, error)
	AddSession(sessionSubmit *model.SessionSubmit) error
	GetSessionById(id string) (*repository.Session, error)
	RemoveSessionById(id string, sessionMap map[string][]*repository.SessionItem) error

	GetAllItems() ([]*repository.SessionItem, error)
	GetItemsBySessionId(sessionId string) ([]*repository.SessionItem, error)
	AddItem(sessionItem *model.SessionItemSubmit) error
	RemoveItemById(id string) error
}

type service struct {
	mongoRepository repository.MongoRepository
}

func New(mongoRepository repository.MongoRepository) *service {
	return &service{
		mongoRepository: mongoRepository,
	}
}
