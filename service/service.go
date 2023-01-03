package service

import "go-qn2management/repository"

type Service interface {
	GetAllSessions() ([]*repository.Session, error)
}

type service struct {
	mongoRepository repository.MongoRepository
}

func New(mongoRepository repository.MongoRepository) *service {
	return &service{
		mongoRepository: mongoRepository,
	}
}
