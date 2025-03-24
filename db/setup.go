package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-PIPELINE/types"
	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-PIPELINE/utility"
)

func PopulateDB(database *sql.DB) error {
	// Fetch all teams from the Sportmonk database.
	getTeamsResponse := utility.GetTeams()

	if err := database.Ping(); err != nil {
		log.Fatal("DB not responding.")
	}

	trx, err := database.Begin()
	if err != nil {
		return err
	}

	// Insert teams into the database.
	err = insertTeams(getTeamsResponse.Data, database)
	if err != nil {
		_ = trx.Rollback()
		return err
	}

	// TODO: Write query to populate seasonal schedule.

	if err := trx.Commit(); err != nil {
		return err
	}

	fmt.Println("DB populated: teams")
	return nil
}

// TODO: Move to a seperate file.
func insertTeams(teams []types.Team, db *sql.DB) error {
	valueStrings := make([]string, 0, len(teams))
	valueArgs := make([]any, 0, len(teams)*5)

	for i, team := range teams {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", i*5+1, i*5+2, i*5+3, i*5+4, i*5+5))
		valueArgs = append(valueArgs, team.Id)
		valueArgs = append(valueArgs, team.Name)
		valueArgs = append(valueArgs, team.ShortCode)
		valueArgs = append(valueArgs, team.ImgPath)
		valueArgs = append(valueArgs, team.CountryId)
	}

	query := fmt.Sprintf("INSERT INTO teams (id, name, short_code, img_path, country_id) VALUES %s", strings.Join(valueStrings, ","))

	_, err := db.Exec(query, valueArgs...)
	if err != nil {
		return err
	}
	return nil
}
