package rest

import (
	"github.com/jhowilbur/event-backend/config/server"
	"github.com/jhowilbur/event-backend/pkg/context"
	"github.com/jhowilbur/event-backend/pkg/routes/http"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server is the interface that wraps the Start and Close methods.
type Server interface {
	Close() error
	Start(address string) error
}

// defaultServer is the default implementation of Server
type defaultServer struct {
	e        *echo.Echo
	database *sqlx.DB

	routes struct {
		event http.EventRouting
	}
}

// Options for configuring the web server app
type Options struct {
	Web server.WebServerConf
}

// NewServer builds the web server
func NewServer(options Options) (Server, error) {
	var database *sqlx.DB
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderContentType},
	}))

	prepareServer := &defaultServer{
		e:        e,
		database: database,
	}

	prepareServer.routes.event = context.NewEventRoute(database)
	prepareServer.e.GET("/events/:id", prepareServer.routes.event.RoutingRecoverById)
	prepareServer.e.GET("/events", prepareServer.routes.event.RoutingRecoverAll)
	prepareServer.e.POST("/events", prepareServer.routes.event.RoutingRegister)

	// TODO middleware for security

	serverConf := options.Web.GetWebServerConfiguration()
	e.Logger.Fatal(prepareServer.Start(serverConf.Port))

	return prepareServer, nil
}

// Close shuts down the web server
func (server defaultServer) Close() error {
	return server.e.Close()
}

// Start starts the web server
func (server defaultServer) Start(address string) error {
	return server.e.Start(address)
}
