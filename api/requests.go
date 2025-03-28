package api

import (
	"encoding/json"
	"log"

	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-PIPELINE/types"
	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-PIPELINE/utility"
)

// TODO: Move to a JSON config file within the config directory.
var urls = map[string]string{
	"teamsBySeasonId":    "https://api.sportmonks.com/v3/football/teams/seasons/24962?api_token=%v",
	"scheduleBySeasonId": "https://api.sportmonks.com/v3/football/schedules/seasons/24962?api_token=%v",
}

// Get all teams participating in the 24962 season.
func GetTeams() types.GetTeamsResponse {
	var teamsResponse types.GetTeamsResponse
	responseBody, err := utility.FetchContent(urls["teamsBySeasonId"], "teams")
	if err != nil {
		log.Fatal("Error reading response body: ", err)
	}

	err = json.Unmarshal(responseBody, &teamsResponse)
	if err != nil {
		log.Fatal("Error unmarshalling teams data: ", err)
	}

	return teamsResponse
}

// Get all fixtures/matches for the 24962 season.
func GetSchedules() types.GetScheduleResponse {
	var schedulesResponse types.GetScheduleResponse
	responseBody, err := utility.FetchContent(urls["scheduleBySeasonId"], "schedules")
	if err != nil {
		log.Fatal("Error reading response body: ", err)
	}

	err = json.Unmarshal(responseBody, &schedulesResponse)
	if err != nil {
		log.Fatal("Error unmarshalling schedule data: ", err)
	}

	return schedulesResponse
}
