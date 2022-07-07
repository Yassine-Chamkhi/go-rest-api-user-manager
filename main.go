package main

import (
	"fmt"
	"os"
	"target/onboarding-assignment/http/handlers"
	"target/onboarding-assignment/repository"
	"target/onboarding-assignment/server"
	"target/onboarding-assignment/services"

	"github.com/gin-gonic/gin"
)

func main() {
	//Initializing database connection
	conn, err := repository.ConnectToDatabase()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer conn.Close()

	//Initializing dependency chain
	userRepo := repository.UserRepository{Db: conn}
	userService := services.UserService{Repo: &userRepo}
	userHandler := handlers.UserHandler{UserSvc: &userService}
	server := server.Server{HTTPHandler: &userHandler, Router: gin.Default(), Port: ":8080"}

	server.InitRoutes()
	server.Run()
}
