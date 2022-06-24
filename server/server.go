package server

import (
	"log"
	"modulo/server/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	port   string
	server *echo.Echo
}

func NewServer() Server {
	return Server{
		port:   ":1323",
		server: echo.New(),
	}
}

func (base *Server) Run() {
	s := base.server
	s.Use(middleware.Logger())
	s.Use(middleware.Recover())
	routes.ConfigRoutes(s)
	s.Logger.Fatal(s.Start(base.port))

	log.Print("Server rodando na porta: " + base.port)
}
