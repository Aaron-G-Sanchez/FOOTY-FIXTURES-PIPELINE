package api

import (
	"errors"
	"log"
	"regexp"

	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-PIPELINE/types"
)

// Call GetSchedules and process the data to return only a list of matches.
func ProcessScheduleResponse() ([]types.Match, error) {
	schdeduleResponse := GetSchedules()

	if len(schdeduleResponse.Data) < 1 {
		return nil, errors.New("No stage for the given season.")
	}

	// Loop through the stages to find the regular season fixtures/matches.
	for _, stage := range schdeduleResponse.Data {
		matched, err := regexp.MatchString(`(?i)regular season`, stage.Name)
		if err != nil {
			log.Fatal("Error matching stage name: ", err)
		}

		if matched {
			return stage.Fixtures, nil
		}
	}

	return nil, errors.New("No fixtures for the given stage")
}
