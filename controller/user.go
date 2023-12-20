package controller

import (
	"net/http"
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
	var req UserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	// fmt.Println("creating...", req.Name)
	u := toModel(req)
	// fmt.Println(u.Name, u.Age, u.Email, u.CreatedAt, u.UpdatedAt)
	c.u.Create(ctx.Request().Context(), u)

	return nil
}
