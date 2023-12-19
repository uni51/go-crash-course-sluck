package main

import (
	"sluck/controller"
	"sluck/repository"
	"sluck/usecase"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	ur := repository.NewUserRepository(nil)
	uu := usecase.NewUserUsecase(ur)
	uc := controller.NewUserController(uu)

	e.POST("/users", uc.Create)
	e.Start(":8080")
}
