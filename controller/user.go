package controller

import (
	"net/http"
	"sluck/usecase"

	"github.com/labstack/echo"
)

// UserControllerはユーザー関連のHTTPリクエストを処理するコントローラーのインターフェースです。
type UserController interface {
	Get(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(echo.Context) error
	Delete(echo.Context) error
}

// userControllerはUserControllerの実装です。
type userController struct {
	u usecase.UserUsecase
}

// NewUserControllerはUserControllerの新しいインスタンスを作成します。
func NewUserController(u usecase.UserUsecase) UserController {
	return &userController{u}
}

func (c *userController) Get(ctx echo.Context) error {
	id := ctx.Param("id")
	u, err := c.u.GetById(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, u)
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

	u := toModel(req)
	c.u.Create(ctx.Request().Context(), &u)

	return nil
}

func (c *userController) Update(ctx echo.Context) error {
	var req UserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	u := toModel(req)
	c.u.Update(ctx.Request().Context(), &u)

	return nil
}

func (c *userController) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	c.u.Delete(ctx.Request().Context(), id)

	return nil
}
