package rest_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jhowilbur/event-backend/config/server"
	"github.com/jhowilbur/event-backend/rest"
	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// TODO - Finish the test preparation
func TestServerEndpoint(t *testing.T) {
	e := echo.New()

	options := rest.Options{
		Web: server.WebServerConf{
			Port: "", // remember to change the port in .env file
		},
	}

	s, err := rest.NewServer(options)
	assert.NoError(t, err)

	go func() {
		err := s.Start(":" + options.Web.Port)
		assert.NoError(t, err)
	}()

	// Wait for the server to start
	// You can add a delay or any other mechanism to ensure the server is ready

	req := httptest.NewRequest(http.MethodGet, "/events", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	err = s.Close()
	assert.NoError(t, err)
}
