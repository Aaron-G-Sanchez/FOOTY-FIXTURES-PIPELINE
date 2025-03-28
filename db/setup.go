package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-PIPELINE/api"
	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-PIPELINE/types"
)

func PopulateDB(database *sql.DB) error {
	// Fetch all teams from the Sportmonk API.
	getTeamsResponse := api.GetTeams()
	// Fetch seasonal schedule from the Sportmonk API.
	matches, err := api.ProcessScheduleResponse()
	if err != nil {
		return err
	}

	// Insert teams into the database.
	err = insertTeams(getTeamsResponse.Data, database)
	if err != nil {
		return err
	}
	fmt.Println("DB populated: teams")

	// Insert schedule into the databse.
	err = insertSchedule(matches)
	if err != nil {
		return err
	}
	fmt.Println("DB populated: schedule")

	return nil
}

func insertTeams(teams []types.Team, db *sql.DB) error {
	// Begin database transaction.
	trx, err := db.Begin()
	if err != nil {
		return err
	}

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

	query := fmt.Sprintf(
		"INSERT INTO teams (id, name, short_code, img_path, country_id) VALUES %s",
		strings.Join(valueStrings, ","),
	)

	// Execute the query(INSERT).
	_, err = db.Exec(query, valueArgs...)
	if err != nil {
		_ = trx.Rollback()
		return err
	}

	// Commit the transaction to the db.
	if err := trx.Commit(); err != nil {
		return err
	}

	return nil
}

func insertSchedule(matches []types.Match) error {
	for _, match := range matches {
		fmt.Println("Match:", match.Id)
	}

	return nil
}
