package config

import (
	"github.com/jhowilbur/event-backend/config/db"
	"github.com/jhowilbur/event-backend/config/server"
	"github.com/jmoiron/sqlx"
)

// WebServerConf contains the configuration related to the web server exposing our REST API.
type WebServerConf interface {
	GetWebServerConfiguration() server.WebServerConf
}

// StorageConf contains the configuration related to the storage like SQL database or S3 bucket.
type StorageConf interface {
	GetStorageConfiguration() (db.DatabaseConf, error)
	GetConnection(conf db.DatabaseConf) (*sqlx.DB, error)
}
