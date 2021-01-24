package repository

import (
	"context"
	"log"

	"bitbucket.org/oaroz/hai/app/common"
	"bitbucket.org/oaroz/hai/app/domain"
	"bitbucket.org/oaroz/hai/app/errors"
	"github.com/jackc/pgx/v4/pgxpool"
)

type MessageRepository interface {
	Get(email string) ([]domain.Message, error)
	GetAll() ([]domain.Message, error)
	Create(message domain.Message) (domain.Message, error)
	Delete(id int64, code string) error
}

type messageRepository struct {
	dbCon *pgxpool.Pool
}

func NewMessageRepository(dbCon *pgxpool.Pool) MessageRepository {
	return messageRepository{dbCon: dbCon}
}

func (r messageRepository) Get(email string) ([]domain.Message, error) {
	messages := []domain.Message{}
	sqlStatement := "SELECT id, title, code, content, email FROM MESSAGES WHERE messages.email=$1"
	rows, err := r.dbCon.Query(context.Background(), sqlStatement, email)
	if err != nil {
		log.Printf("Query cannot be executed. Error: %v\n", err)
		return nil, errors.New(errors.InternalDbError)
	}
	var message domain.Message
	for rows.Next() {
		err = rows.Scan(&message.ID, &message.Title, &message.Code, &message.Content, &message.Email)
		if err != nil {
			log.Printf("Sql row cannot be mapped to struct. Error: %v\n", err)
			return nil, errors.New(errors.InternalDbError)
		}
		messages = append(messages, message)
	}
	return messages, nil
}

func (r messageRepository) GetAll() ([]domain.Message, error) {
	messages := []domain.Message{}
	sqlStatement := "SELECT id, title, code, content, email FROM MESSAGES"
	rows, err := r.dbCon.Query(context.Background(), sqlStatement)
	if err != nil {
		log.Printf("Query cannot be executed. Error: %v\n", err)
		return []domain.Message{}, errors.New(errors.InternalDbError)
	}
	var message domain.Message
	for rows.Next() {
		err = rows.Scan(&message.ID, &message.Title, &message.Code, &message.Content, &message.Email)
		if err != nil {
			log.Printf("Sql row cannot be mapped to struct. Error: %v\n", err)
			return []domain.Message{}, errors.New(errors.InternalDbError)
		}
		messages = append(messages, message)
	}
	return messages, nil
}

func (r messageRepository) Create(msg domain.Message) (domain.Message, error) {
	msg.Code = common.RandomCode(20)
	sqlStatement := "INSERT INTO messages(code, title, content, email) VALUES($1, $2, $3, $4) RETURNING id"
	err := r.dbCon.QueryRow(context.Background(), sqlStatement, msg.Code, msg.Title, msg.Content, msg.Email).Scan(&msg.ID)
	if err != nil {
		log.Printf("Insert cannot be executed. Error: %v\n", err)
		return domain.Message{}, errors.HaiError{ErrorCode: errors.InternalDbError}
	}
	return msg, nil
}

func (r messageRepository) Delete(id int64, code string) error {
	sqlStatement := "DELETE FROM messages WHERE id = $1 and code = $2"
	res, err := r.dbCon.Exec(context.Background(), sqlStatement, id, code)
	if err != nil {
		return errors.New(errors.InternalDbError)
	} else {
		count := res.RowsAffected()
		if count == 0 {
			log.Printf("Message with provided id: %v and code %v does not exist\n", id, code)
			return errors.New(errors.MessageNotFound)
		}
	}
	return nil
}
