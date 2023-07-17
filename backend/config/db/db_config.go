package db

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

const dbHostEnv = "DB_HOST"
const dbPortEnv = "DB_PORT"
const dbUserEnv = "DB_USER"
const dbPasswordEnv = "DB_PASSWORD"
const dbDatabaseEnv = "DB_DATABASE"
const dbSslmodeEnv = "DB_SSLMODE"

// DatabaseConf contains the configuration related to the database.
type DatabaseConf struct {
	Host     string `validate:"required"`
	Port     string `validate:"required"`
	User     string `validate:"required"`
	Password string `validate:"required"`
	Database string `validate:"required"`
	Sslmode  string `validate:"required"`
}

// NewDatabaseConf returns a new DatabaseConf
func NewDatabaseConf() DatabaseConf {
	return DatabaseConf{}
}

// GetStorageConfiguration returns our database configuration
func (d DatabaseConf) GetStorageConfiguration() (DatabaseConf, error) {
	d.Host = os.Getenv(dbHostEnv)
	d.Port = os.Getenv(dbPortEnv)
	d.User = os.Getenv(dbUserEnv)
	d.Password = os.Getenv(dbPasswordEnv)
	d.Database = os.Getenv(dbDatabaseEnv)
	d.Sslmode = os.Getenv(dbSslmodeEnv)

	validate := validator.New()

	err := validate.Struct(d)
	if err != nil {
		log.Fatalf("Invalid database env configuration: %v", err)
		return d, err
	}

	println("Database configuration loaded successfully")
	return d, nil
}

func (d DatabaseConf) GetConnection(conf DatabaseConf) (*sqlx.DB, error) {
	databaseURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		conf.Host, conf.Port, conf.User, conf.Password, conf.Database, conf.Sslmode)

	db, err := sqlx.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
