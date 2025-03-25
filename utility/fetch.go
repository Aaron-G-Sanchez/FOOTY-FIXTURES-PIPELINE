package utility

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-PIPELINE/config"
)

var urls = map[string]string{
	"teamsBySeasonId":    "https://api.sportmonks.com/v3/football/teams/seasons/24962?api_token=%v",
	"scheduleBySeasonId": "https://api.sportmonks.com/v3/football/schedules/seasons/24962?api_token=%v",
}

func FetchContent(endpoint, entity string) ([]byte, error) {
	client := &http.Client{}

	endpointWithToken := fmt.Sprintf(endpoint, config.AppConfig.APIToken)

	request, err := http.NewRequest("GET", endpointWithToken, nil)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("Error fetching %s: %v", entity, err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
