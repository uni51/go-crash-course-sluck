package repository

import (
	"context"
	"database/sql"
	"fmt"
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
	db, ok := GetTx(ctx)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	result, err := db.Exec("DELETE FROM messages WHERE user_id = ?", userId)
	if err != nil {
		return err
	}

	rowsAffect, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffect == 0 {
		return fmt.Errorf("no rows affected: %s", userId)
	}

	return nil
}
