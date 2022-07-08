package main

import (
	"fmt"
	"os"
	"target/onboarding-assignment/http/handlers"
	"target/onboarding-assignment/repository"
	"target/onboarding-assignment/server"
	"target/onboarding-assignment/services"

	migrate "github.com/rubenv/sql-migrate"

	"github.com/gin-gonic/gin"
)

func main() {
	//Initializing database connection
	fmt.Println("Connecting to database")
	conn, err := repository.ConnectToDatabase()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer conn.Close()

	fmt.Println("Defining migrations path")
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations/postgres",
	}

	fmt.Println("Starting migrations")
	n, err := migrate.Exec(conn, "postgres", migrations, migrate.Up)
	if err != nil {
		fmt.Println("ERROR WHEN MIGRATING UP", err)
	}
	fmt.Printf("Applied %d migrations!\n", n)

	//Initializing dependency chain
	fmt.Println("Initializing dependency chain")
	userRepo := repository.UserRepository{Db: conn}
	userService := services.UserService{Repo: &userRepo}
	userHandler := handlers.UserHandler{UserSvc: &userService}
	server := server.Server{HTTPHandler: &userHandler, Router: gin.Default(), Port: ":8080"}

	server.InitRoutes()
	server.Run()
}
