package chat

import (
	"github.com/ganselmo/go-first-project/internal/config"
	"github.com/jmoiron/sqlx"
)

//Message ...
type Message struct {
	ID   int64
	Text string
}

//ChatService ...
type Service interface {
	AddMessage(Message) error
	FindById(int) *Message
	FindAll() []*Message
}

type service struct {
	db     *sqlx.DB
	config *config.Config
}

func NewChatService(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

func (s service) AddMessage(Message) error {
	return nil
}
func (s service) FindById(int) *Message {
	return nil
}
func (s service) FindAll() []*Message {
	var list []*Message
	if err := s.db.Select(&list, "SELECT * FROM messages"); err != nil {
		panic(err)
	}
	return list
}
