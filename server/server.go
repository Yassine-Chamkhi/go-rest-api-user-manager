package server

import (
	"go-rest-api/http/handlers"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type ServerInterface interface {
	InitRoutes()
	Run()
}

type Server struct {
	HTTPHandler handlers.UserHandlerInterface
	Router      *gin.Engine
	Port        string
}

func (srv *Server) InitRoutes() {
	srv.Router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	srv.Router.GET("/", srv.HTTPHandler.Greet())
	srv.Router.GET("/users/:id", srv.HTTPHandler.GetUserFromPath())
	srv.Router.GET("/users", srv.HTTPHandler.GetAllUsers())
	srv.Router.POST("/users", srv.HTTPHandler.PostUser())
	srv.Router.DELETE("/users", srv.HTTPHandler.DeleteUser())
}

func (srv *Server) Run() {
	srv.Router.Run(srv.Port)
}
