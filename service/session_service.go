package service

import (
	"go-qn2management/model"
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

func (s *service) AddSession(sessionSubmit *model.SessionSubmit) error {
	session := repository.Session{
		SessionName: sessionSubmit.SessionName,
		SessionKey:  sessionSubmit.SessionKey,
	}

	err := s.mongoRepository.InsertSession(&session)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *service) GetSessionById(id string) (*repository.Session, error) {
	session, err := s.mongoRepository.FindById(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return session, nil
}
