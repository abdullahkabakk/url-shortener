package http

import (
	"context"
	"fmt"
	_ "net/http"
	auth_handler "url-shortener/internal/app/handlers/auth"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server represents the HTTP server.
type Server struct {
	echo *echo.Echo
	host string
	port string
}

// NewServer creates a new instance of the HTTP server.
func NewServer(host, port string, userHandler *auth_handler.Handler) *Server {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	authGroup := e.Group("/auth")

	authRouter(authGroup, userHandler)

	return &Server{
		echo: e,
		host: host,
		port: port,
	}
}

// Start starts the HTTP server.
func (s *Server) Start() error {
	addr := fmt.Sprintf("%s:%s", s.host, s.port)
	return s.echo.Start(addr)
}

// Shutdown shuts down the HTTP server.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}

func authRouter(group *echo.Group, userHandler *auth_handler.Handler) {
	group.POST("/register/", userHandler.CreateUserHandler)
	group.POST("/login/", userHandler.LoginUserHandler)
}
