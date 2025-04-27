package main

import (
	"INF368/internal/controller"
	"INF368/internal/repository"
	"INF368/internal/models"
)

func main() {
	repo := repository.NewUserRepository()
	svc := service.NewUserService(repo)
	h := handler.NewUserHandler(svc)

	router := h.InitRoutes()

	router.Run(":8080")
}
