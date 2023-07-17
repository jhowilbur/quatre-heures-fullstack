package cmd

import (
	databaseConfig "github.com/jhowilbur/event-backend/config/db"
	migrate "github.com/jhowilbur/event-backend/db"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"log"
)

// MainCommand performs database migrations
func MainCommand() *cobra.Command {
	var rootCommand = &cobra.Command{
		Use:   "db",
		Short: "Database tooling",
	}

	rootCommand.AddCommand(&cobra.Command{
		Use:   "migrate",
		Short: "Performs database migrations or read migrations info",
		Run:   performMigrations,
	})
	return rootCommand
}

func performMigrations(cmd *cobra.Command, args []string) {
	configuration, _ := databaseConfig.NewDatabaseConf().GetStorageConfiguration()
	connection, err := configuration.GetConnection(configuration)
	if err != nil {
		log.Fatalf("failed to connect to database: %w", err)
		return
	}

	err = migrate.ApplySQLScript(connection)
	if err != nil {
		log.Fatalf("failed to apply SQL script: %w", err)
	}

	defer func(connection *sqlx.DB) {
		err := connection.Close()
		if err != nil {
			return
		}
	}(connection)

	log.Println("Successfully applied migration SQL script")
}
