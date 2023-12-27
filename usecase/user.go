package usecase

import (
	"context"
	"sluck/model"
	"sluck/repository"
	"sluck/transaction"
)

// UserUsecaseはユーザー関連のビジネスロジックを抽象化するインターフェースです。
type UserUsecase interface {
	GetById(ctx context.Context, id string) (*model.User, error)
	Create(ctx context.Context, user *model.User) (string, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id string) error
}

// userUsecaseはUserUsecaseの実装です。
type userUsecase struct {
	r           repository.UserRepository
	mr          repository.MessageRepository
	transaction transaction.Transaction
}

// NewUserUsecaseはUserUsecaseの新しいインスタンスを作成します。
func NewUserUsecase(r repository.UserRepository, mr repository.MessageRepository, transaction transaction.Transaction) UserUsecase {
	return &userUsecase{r, mr, transaction}
}

func (u *userUsecase) GetById(ctx context.Context, id string) (*model.User, error) {
	user, err := u.r.Read(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Createは新しいユーザーを作成するためのビジネスロジックを実行します。
func (u *userUsecase) Create(ctx context.Context, user *model.User) (string, error) {
	id, err := u.r.Create(ctx, user)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (u *userUsecase) Update(ctx context.Context, user *model.User) error {
	err := u.r.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) Delete(ctx context.Context, id string) error {
	u.transaction.DoInTx(ctx, func(ctx context.Context) (any, error) {
		err := u.r.Delete(ctx, id)
		if err != nil {
			return nil, err
		}

		// ユーザーのメッセージも削除する
		err = u.mr.Delete(ctx, id)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})

	return nil
}
