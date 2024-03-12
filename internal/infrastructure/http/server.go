package http

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "net/http"
	"url-shortener/internal/app/handlers/auth"
	clicks_handler "url-shortener/internal/app/handlers/clicks"
	"url-shortener/internal/app/handlers/url"
)

// Server represents the HTTP server.
type Server struct {
	echo *echo.Echo
	host string
	port string
}

// NewServer creates a new instance of the HTTP server.
func NewServer(host, port string, userHandler *auth_handler.Handler, urlHandler *url_handler.Handler, clickHandler *clicks_handler.Handler) *Server {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	authGroup := e.Group("/auth")

	urlGroup := e.Group("/url")

	clicksGroup := e.Group("/clicks")

	authRouter(authGroup, userHandler)

	urlRoute(urlGroup, urlHandler)

	clicksRoute(clicksGroup, clickHandler)

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
	group.GET("/refresh-token/", userHandler.RefreshTokenHandler)
}

func urlRoute(group *echo.Group, urlHandler *url_handler.Handler) {
	group.POST("/shorten/", urlHandler.ShortenURLHandler)
	group.GET("/", urlHandler.GetUserUrlsHandler)
}

func clicksRoute(group *echo.Group, clickHandler *clicks_handler.Handler) {
	group.GET("/:id", clickHandler.CreateClickHandler)
	group.GET("/:id/details/", clickHandler.GetUserClickDetailsHandler)
}
