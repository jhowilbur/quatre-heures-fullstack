package db

import (
	databaseConfig "github.com/jhowilbur/event-backend/config/db"
	"github.com/jhowilbur/event-backend/pkg/entity"
	"github.com/jhowilbur/event-backend/pkg/storage"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

// DatabaseConnection returns a new database connection
type defaultEventRepository struct {
	db       sqlx.Ext
	dbConfig databaseConfig.DatabaseConf
}

// NewEventRepository returns a new instance of the event repository
func NewEventRepository(db sqlx.Ext) storage.EventRepository {
	config, err := databaseConfig.NewDatabaseConf().GetStorageConfiguration()
	if err != nil {
		return nil
	}

	return defaultEventRepository{
		db:       db,
		dbConfig: config,
	}
}

// FindOneById finds an event by its id
func (d defaultEventRepository) FindOneById(id uint) (*entity.Event, error) {
	db, _ := d.dbConfig.GetConnection(d.dbConfig)

	var event entity.Event
	err := db.Get(&event, "SELECT * FROM events WHERE id = $1", id)
	if err != nil {
		log.Printf("can't find event id: %d error: [%v]", id, err)
		return nil, err
	}

	return &event, nil
}

func (d defaultEventRepository) FindAll() ([]entity.Event, error) {
	db, _ := d.dbConfig.GetConnection(d.dbConfig)

	var events []entity.Event
	err := db.Select(&events, "SELECT * FROM events")
	if err != nil {
		log.Println(err)
	}

	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}(db)

	return events, nil
}

// Create creates a new event
func (d defaultEventRepository) Create(event *entity.NewEvent) error {
	db, _ := d.dbConfig.GetConnection(d.dbConfig)

	_, err := db.Exec("INSERT INTO events (title, description, initialdate, finaldate) VALUES ($1, $2, $3, $4)", event.Title, event.Description, event.InitialDate, event.FinalDate)
	if err != nil {
		log.Println(err)
	}

	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}(db)

	return nil
}
