package usecase

import (
	"context"
	"fmt"
	"sluck/model"
	"sluck/repository"
)

// UserUsecaseはユーザー関連のビジネスロジックを抽象化するインターフェースです。
type UserUsecase interface {
	// Createは新しいユーザーを作成するためのビジネスロジックを実行します。
	Create(ctx context.Context, user *model.User) error
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
func (u *userUsecase) Create(ctx context.Context, user *model.User) error {
	fmt.Println("usecase creating...")
	return nil
}
