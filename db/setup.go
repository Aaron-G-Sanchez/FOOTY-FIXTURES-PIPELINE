package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-PIPELINE/api"
	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-PIPELINE/types"
)

// Populate the database when a new db is created.
func PopulateDB(database *sql.DB) error {
	getTeamsResponse := api.GetTeams()
	getScheduleResponse := api.GetSchedules()

	matches, err := api.ProcessScheduleResponse(getScheduleResponse)
	if err != nil {
		return err
	}

	// TODO: Look into logging the error if data is not inserted rather than
	// returning the error.
	err = insertTeams(getTeamsResponse.Data, database)
	if err != nil {
		return err
	}
	fmt.Println("DB populated: teams")

	err = insertMatches(matches, database)
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

func insertMatches(matches []types.Match, db *sql.DB) error {
	// TODO: Insert matches into the matches table.

	// Begin DB insertion.
	trx, err := db.Begin()
	if err != nil {
		return err
	}

	valueStrings := make([]string, 0, len(matches))
	valueArgs := make([]any, 0, len(matches)*6)

	for i, match := range matches {
		valueStrings = append(valueStrings, fmt.Sprintf(
			"($%d, $%d, $%d, $%d, $%d, $%d)",
			i*6+1, i*6+2, i*6+3, i*6+4, i*6+5, i*6+6,
		))
		valueArgs = append(valueArgs, match.Id)
		valueArgs = append(valueArgs, match.LeagueId)
		valueArgs = append(valueArgs, match.SeasonId)
		valueArgs = append(valueArgs, match.Name)
		valueArgs = append(valueArgs, match.StartingAt)
		valueArgs = append(valueArgs, match.ResultInfo)
	}

	query := fmt.Sprintf(
		"INSERT INTO matches (id, league_id, season_id, name, starting_at, result_info) VALUES %s",
		strings.Join(valueStrings, ","),
	)

	// Exectute the query(INSERT)
	_, err = db.Exec(query, valueArgs...)
	if err != nil {
		_ = trx.Rollback()
		return err
	}

	if err := trx.Commit(); err != nil {
		return err
	}

	return nil
}
