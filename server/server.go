package server

import (
	"target/onboarding-assignment/http/handlers"

	"github.com/gin-gonic/gin"
)

type ServerInterface interface {
	initRoutes()
	Run()
}

type Server struct {
	HTTPHandler handlers.UserHandlerInterface
	Router      *gin.Engine
	Port        string
}

func (srv *Server) InitRoutes() {
	srv.Router.GET("/", srv.HTTPHandler.Greet())
	srv.Router.GET("/users/:id", srv.HTTPHandler.GetUserFromPath())
	srv.Router.GET("/users", srv.HTTPHandler.GetAllUsers())
	srv.Router.POST("/users", srv.HTTPHandler.PostUser())
	srv.Router.DELETE("/users", srv.HTTPHandler.DeleteUser())
}

func (srv *Server) Run() {
	srv.Router.Run(srv.Port)
}
