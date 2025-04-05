package utility

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Fetch data for the given entity from the Sportmonk's API.
func FetchContent(endpoint, entity, token string) ([]byte, error) {
	client := &http.Client{}

	endpointWithToken := fmt.Sprintf(endpoint, token)
	request, err := http.NewRequest("GET", endpointWithToken, nil)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("Error fetching %s: %v", entity, err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("Server returned non-OK status code")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
