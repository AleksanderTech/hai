package repository

import (
	"context"
	"fmt"

	"bitbucket.org/oaroz/hai/app/common"

	"bitbucket.org/oaroz/hai/app/domain"
	"github.com/jackc/pgx/v4/pgxpool"
)

type MessageRepository interface {
	Get(email string) []domain.Message
	Create(message domain.Message) domain.Message
	Delete(id int64, code string)
}

type messageRepository struct {
	dbCon *pgxpool.Pool
}

func NewMessageRepository(dbCon *pgxpool.Pool) MessageRepository {
	return messageRepository{dbCon: dbCon}
}

func (r messageRepository) Get(email string) []domain.Message {
	messages := []domain.Message{}
	sqlStatement := "SELECT id, title, code, content, email FROM MESSAGES WHERE messages.email=$1"
	rows, err := r.dbCon.Query(context.Background(), sqlStatement, email)
	if err != nil {
		panic(fmt.Sprintf("Query cannot be executed. Error: %v\n", err))
	}
	var message domain.Message
	for rows.Next() {
		err = rows.Scan(&message.ID, &message.Title, &message.Code, &message.Content, &message.Email)
		if err != nil {
			panic(fmt.Sprintf("Sql row cannot be mapped to struct. Error: %v\n", err))
		}
		messages = append(messages, message)
	}
	return messages
}

func (r messageRepository) Create(msg domain.Message) domain.Message {
	msg.Code = common.RandomCode(20)
	sqlStatement := "INSERT INTO messages(code, title, content, email) VALUES($1, $2, $3, $4) RETURNING id"
	err := r.dbCon.QueryRow(context.Background(), sqlStatement, msg.Code, msg.Title, msg.Content, msg.Email).Scan(&msg.ID)

	if err != nil {
		panic(fmt.Sprintf("Insert cannot be executed. Error: %v\n", err))
	}
	return msg
}

func (r messageRepository) Delete(id int64, code string) {
	sqlStatement := "DELETE FROM messages WHERE id = $1 and code = $2"
	_, err := r.dbCon.Exec(context.Background(), sqlStatement, id, code)
	if err != nil {
		panic(fmt.Sprintf("Delete cannot be executed. Error: %v\n", err))
	}
}
