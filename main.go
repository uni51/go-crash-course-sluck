package main

import (
	"fmt"
	"sluck/controller"
	"sluck/repository"
	"sluck/usecase"
)

func main() {
	fmt.Println("Hello World!")

	ur := repository.NewUserRepository(nil)
	uu := usecase.NewUserUsecase(ur)
	uc := controller.NewUserController(uu)

	uc.Create(nil)
}
