package main

import (
	"net/http"
	"sluck/controller"
	"sluck/infra"
	"sluck/repository"
	"sluck/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	db := infra.Connect()
	ur := repository.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur)
	uc := controller.NewUserController(uu)

	e.POST("/users", uc.Create)
	e.Start(":8080")
}
