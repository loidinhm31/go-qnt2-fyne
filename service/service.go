package service

import (
	"go-qn2management/model"
	"go-qn2management/repository"
)

type Service interface {
	GetAllSessions() ([]*repository.Session, error)
	AddSession(sessionSubmit *model.SessionSubmit) error
	GetSessionById(id string) (*repository.Session, error)

	GetAllItems() ([]*repository.SessionItem, error)
	GetItemsBySessionId(sessionId string) ([]*repository.SessionItem, error)
	AddItem(sessionItem *model.SessionItemSubmit) error
}

type service struct {
	mongoRepository repository.MongoRepository
}

func New(mongoRepository repository.MongoRepository) *service {
	return &service{
		mongoRepository: mongoRepository,
	}
}
