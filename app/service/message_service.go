package service

import (
	"bitbucket.org/oaroz/hai/app/domain"
	"bitbucket.org/oaroz/hai/app/repository"
)

type MessageService interface {
	Get(email string) ([]domain.Message, error)
	Create(message domain.Message) (domain.Message, error)
	Delete(id int64, code string) error
}

type messageService struct {
	repo repository.MessageRepository
}

func NewMessageService(r repository.MessageRepository) MessageService {
	return messageService{repo: r}
}

func (s messageService) Get(email string) ([]domain.Message, error) {
	if email == "" {
		return s.repo.GetAll()
	} else {
		return s.repo.Get(email)
	}
}

func (s messageService) Create(message domain.Message) (domain.Message, error) {
	// business logic goes here
	return s.repo.Create(message)
}

func (s messageService) Delete(id int64, code string) error {
	// business logic goes here
	return s.repo.Delete(id, code)
}
