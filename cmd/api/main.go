package main

import (
	"fmt"
	"log"

	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-PIPELINE/config"
	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-PIPELINE/db"
	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-PIPELINE/utility"
)

func init() {
	config.LoadConfig()
}

func main() {
	database := db.Connect(config.DatabaseConfig.StagingURI)

	err := db.PopulateDB(database)
	if err != nil {
		log.Fatal("Error populating database: ", err)
	}

	getScheduleResponse := utility.GetSchedule()

	fmt.Println(getScheduleResponse.Data)
}
