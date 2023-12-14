package usecase

import (
	"context"
	"sluck/repository"
)

// UserUsecaseはユーザー関連のビジネスロジックを抽象化するインターフェースです。
type UserUsecase interface {
	// Createは新しいユーザーを作成するためのビジネスロジックを実行します。
	Create(ctx context.Context) error
}

// userUsecaseはUserUsecaseの実装です。
type userUsecase struct {
	u repository.UserRepository
}

// NewUserUsecaseはUserUsecaseの新しいインスタンスを作成します。
func NewUserUsecase(r repository.UserRepository) UserUsecase {
	return &userUsecase{r}
}

// Createは新しいユーザーを作成するためのビジネスロジックを実行します。
func (u *userUsecase) Create(ctx context.Context) error {
	// 実際のユーザー作成のビジネスロジックはここに追加します。
	return nil
}
