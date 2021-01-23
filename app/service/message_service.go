package service

import (
	"context"
	"errors"
	"fmt"

	"bitbucket.org/oaroz/hai/app/domain"
	"github.com/jackc/pgx/v4/pgxpool"
)

type MessageService interface {
	Get(email string) ([]domain.Message, error)
	Create(message domain.Message) (domain.Message, error)
	Delete(id string, code string) (domain.Message, error)
}

type messageService struct {
	dbCon *pgxpool.Pool
}

func NewService(dbCon *pgxpool.Pool) MessageService {
	return messageService{dbCon: dbCon}
}

func (s messageService) Get(email string) ([]domain.Message, error) {
	messages := []domain.Message{}
	rows, err := s.dbCon.Query(context.Background(), "SELECT * FROM MESSAGES WHERE messages.email=$1", email)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Query cannot be executed. Error: %v\n", err))
	}
	var message domain.Message
	for rows.Next() {
		err = rows.Scan(&message.ID, &message.Title, &message.Content, &message.Email)
		messages = append(messages, message)
	}
	return messages, nil
}

func (s messageService) Create(message domain.Message) (domain.Message, error) {
	m := domain.Message{ID: 1, Title: "Title", Content: "Content", Email: "Email"} // mocked
	return m, nil
}

func (s messageService) Delete(id string, code string) (domain.Message, error) { // mocked
	return domain.Message{}, nil
}
