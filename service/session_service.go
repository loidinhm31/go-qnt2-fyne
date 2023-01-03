package service

import (
	"go-qn2management/repository"
	"log"
)

func (s *service) GetAllSessions() ([]*repository.Session, error) {
	sessions, err := s.mongoRepository.FindAllSessions()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return sessions, nil
}
