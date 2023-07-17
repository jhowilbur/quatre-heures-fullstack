package server

import (
	"fmt"
	"log"
	"os"
)

// defaultPort is the default port our web server will listen on
const defaultPort = "8080"
const webPortEnv = "WEB_PORT"

// WebServerConf is the default implementation of our web server configuration
type WebServerConf struct {
	Port string
}

// NewWebServerConf returns our web server configuration
func NewWebServerConf() WebServerConf {
	return WebServerConf{
		Port: defaultPort,
	}
}

// GetWebServerConfiguration returns our web server configuration
func (d WebServerConf) GetWebServerConfiguration() WebServerConf {
	portEnv := os.Getenv(webPortEnv)
	if portEnv == "" {
		log.Fatalf("Error to get PORT from env file using default: %s", defaultPort)
	}
	d.Port = fmt.Sprintf(":%s", portEnv)

	return d
}
