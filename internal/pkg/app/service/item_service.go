package service

import (
	"errors"
	"go-qn2management/internal/pkg/app/model"
	"go-qn2management/internal/pkg/app/repository"
	"log"
)

func (s *service) GetAllItems() ([]*repository.SessionItem, error) {
	sessionItems, err := s.mongoRepository.FindAllItems()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return sessionItems, nil
}

func (s *service) GetItemsBySessionId(sessionId string) ([]*repository.SessionItem, error) {
	sessionItems, err := s.mongoRepository.FindItemsBySessionId(sessionId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return sessionItems, nil
}

func (s *service) AddItem(sessionItemSubmit *model.SessionItemSubmit) error {
	if len(sessionItemSubmit.SessionID) > 0 &&
		len(sessionItemSubmit.Title) > 0 &&
		len(sessionItemSubmit.Extension) > 0 {
		sessionItem := repository.SessionItem{
			Title:     sessionItemSubmit.Title,
			Extension: sessionItemSubmit.Extension,
			SessionID: sessionItemSubmit.SessionID,
		}

		err := s.mongoRepository.InsertItem(&sessionItem)
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		return errors.New("fields cannot empty")
	}
	return nil
}
