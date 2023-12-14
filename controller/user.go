package controller

import (
	"fmt"
	"sluck/usecase"

	"github.com/labstack/echo"
)

// UserControllerはユーザー関連のHTTPリクエストを処理するコントローラーのインターフェースです。
type UserController interface {
	// Createは新しいユーザーを作成するためのHTTPリクエストを処理します。
	Create(ctx echo.Context) error
}

// userControllerはUserControllerの実装です。
type userController struct {
	u usecase.UserUsecase
}

// NewUserControllerはUserControllerの新しいインスタンスを作成します。
func NewUserController(u usecase.UserUsecase) UserController {
	return &userController{u}
}

// Createは新しいユーザーを作成するためのHTTPリクエストを処理します。
func (c *userController) Create(ctx echo.Context) error {
	// 実際のHTTPリクエスト処理のロジックはここに追加します。
	fmt.Println("creating...")
	return nil
}
