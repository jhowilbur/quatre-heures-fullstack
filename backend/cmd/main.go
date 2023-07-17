package cmd

import (
	dbCmd "github.com/jhowilbur/event-backend/db/cmd"
	restCmd "github.com/jhowilbur/event-backend/rest/cmd"
	"github.com/joho/godotenv"
	"log"

	"github.com/spf13/cobra"
)

var (
	cfgFile         string
	restCommandFunc = restCmd.MainCommand
	dbCommandFunc   = dbCmd.MainCommand
)

// RootCommand returns the main command, properly configured
func RootCommand() *cobra.Command {
	var rootCommand = &cobra.Command{
		Use:              "App",
		PersistentPreRun: initConfig,
	}

	rootCommand.PersistentFlags().StringVar(&cfgFile, "config", "", "/path/to/config.yml")
	rootCommand.AddCommand(restCommandFunc())
	rootCommand.AddCommand(dbCommandFunc())

	return rootCommand
}

// initConfig reads in config file and ENV variables if set.
func initConfig(cmd *cobra.Command, args []string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
