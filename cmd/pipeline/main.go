package main

import (
	"log"

	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-PIPELINE/config"
	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-PIPELINE/db"
)

func init() {
	config.LoadConfig()
}

func main() {
	database := db.Connect(config.DatabaseConfig.StagingURI)

	// TODO: place behind command flag.
	err := db.PopulateDB(database)
	if err != nil {
		log.Fatal("Error populating database: ", err)
	}
}
