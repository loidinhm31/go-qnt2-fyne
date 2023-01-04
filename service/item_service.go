package service

import (
	"go-qn2management/model"
	"go-qn2management/repository"
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
	return nil
}
