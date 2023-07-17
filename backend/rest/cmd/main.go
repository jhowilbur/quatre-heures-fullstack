package cmd

import (
	"github.com/jhowilbur/event-backend/config/server"
	"github.com/jhowilbur/event-backend/rest"
	"github.com/spf13/cobra"
)

func MainCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "web",
		Short: "Starts the Web server",
		Run:   bootWebServer,
	}
}

func bootWebServer(cmd *cobra.Command, args []string) {
	rest.NewServer(rest.Options{
		Web: server.NewWebServerConf(),
	})
}
