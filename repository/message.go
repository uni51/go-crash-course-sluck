package repository

import (
	"context"
	"database/sql"
)

type MessageRepository interface {
	Delete(ctx context.Context, userId string) error
}

type messageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) MessageRepository {
	return &messageRepository{db}
}

func (r *messageRepository) Delete(ctx context.Context, userId string) error {
	_, err := r.db.Exec("DELETE FROM messages WHERE user_id = ?", userId)
	if err != nil {
		return err
	}

	return nil
}
