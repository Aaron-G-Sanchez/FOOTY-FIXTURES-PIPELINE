package api

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-PIPELINE/types"
	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-PIPELINE/utility"
)

// TODO: Convert into ENUMs.
var urls = map[string]string{
	"teamsBySeasonId":    "https://api.sportmonks.com/v3/football/teams/seasons/24962?api_token=%v",
	"scheduleBySeasonId": "https://api.sportmonks.com/v3/football/schedules/seasons/24962?api_token=%v",
}

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

	data := schedulesResponse.Data
	fmt.Println(data)
	return schedulesResponse
}
