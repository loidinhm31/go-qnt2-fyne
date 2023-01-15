package service

import (
	"errors"
	"go-qn2management/internal/pkg/app/model"
	"go-qn2management/internal/pkg/app/repository"
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
	if len(sessionSubmit.SessionName) > 0 &&
		len(sessionSubmit.SessionKey) > 0 {
		session := repository.Session{
			SessionName: sessionSubmit.SessionName,
			SessionKey:  sessionSubmit.SessionKey,
			Order:       sessionSubmit.SessionOrder,
		}

		err := s.mongoRepository.InsertSession(&session)
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	}
	return errors.New("fields cannot empty")
}

func (s *service) GetSessionById(id string) (*repository.Session, error) {
	session, err := s.mongoRepository.FindSessionById(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return session, nil
}

func (s *service) RemoveSessionById(id string, sessionMap map[string][]*repository.SessionItem) error {
	// Cannot remove the session that had items
	if len(sessionMap[id]) > 0 {
		return errors.New("cannot remove this session")
	}

	err := s.mongoRepository.DeleteSessionById(id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
