package utility

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-PIPELINE/types"
)

func TestFetchContent(t *testing.T) {
	testTeamsResponse := []types.Team{
		{
			Id:        111,
			Name:      "Test Team 1",
			ShortCode: "tt1",
			ImgPath:   "path/to/img",
			CountryId: 1,
		},
		{
			Id:        222,
			Name:      "Teest Team 2",
			ShortCode: "tt2",
			ImgPath:   "path/to/img",
			CountryId: 1,
		},
	}

	// TODO: Create test helper function to marshal the test data.
	expectedBody, err := json.Marshal(testTeamsResponse)
	if err != nil {
		t.Fatalf("Failed to marshal the testTeamsResponse: %v", err)
	}

	// TODO: Create extpected body for tesing the schedules entity.

	testCases := []struct {
		name         string
		endpoint     string
		expectedBody []byte
	}{
		{
			name:         "Should return a list of teams",
			endpoint:     "/teams?api_token=%v",
			expectedBody: expectedBody,
		},
	}

	// Instantiate a test server.
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/teams":
			token := r.URL.Query().Get("api_token")
			if token == "token" {
				w.Write(expectedBody)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
			}
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer testServer.Close()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			url := testServer.URL + tc.endpoint
			gotTeamResponse, err := FetchContent(url, "teams", "token")
			if err != nil {
				t.Fatalf("Failed to fetch content: %v", err)
			}

			if !reflect.DeepEqual(gotTeamResponse, tc.expectedBody) {
				t.Fatalf("Expected team data: got %v, want %v", string(gotTeamResponse), string(tc.expectedBody))
			}
		})
	}
}
