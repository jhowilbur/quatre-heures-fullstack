package main

import (
	"github.com/jhowilbur/event-backend/cmd"
	"log"
)

var rootCommand = cmd.RootCommand

func main() {
	err := rootCommand().Execute()

	if err != nil {
		log.Println("Error to start the root command: ", err)
		return
	}
}
