package utility

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-PIPELINE/config"
	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-PIPELINE/types"
)

var urls = map[string]string{
	"teamsBySeasonId":    "https://api.sportmonks.com/v3/football/teams/seasons/24962?api_token=%v",
	"scheduleBySeasonId": "https://api.sportmonks.com/v3/football/schedules/seasons/24962?api_token=%v",
}

// TODO: Refactor functions to be one fetch method with entity specific params.

func GetTeams() types.GetTeamsResponse {
	client := &http.Client{}

	endpoint := fmt.Sprintf(urls["teamsBySeasonId"], config.AppConfig.APIToken)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatal("err: ", err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error fetching teams: ", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error reading body: ", err)
	}

	var teamsRes types.GetTeamsResponse
	err = json.Unmarshal(body, &teamsRes)
	if err != nil {
		log.Fatal("Error unmashalling response: ", err)
	}

	return teamsRes
}

func GetSchedule() types.GetScheduleResponse {
	client := http.Client{}

	endpoint := fmt.Sprintf(urls["scheduleBySeasonId"], config.AppConfig.APIToken)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatal("err: ", err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error fetching teams: ", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error reading body: ", err)
	}

	var scheduleRes types.GetScheduleResponse
	err = json.Unmarshal(body, &scheduleRes)
	if err != nil {
		log.Fatal("err unmarshalling response: ", err)
	}

	return scheduleRes
}
