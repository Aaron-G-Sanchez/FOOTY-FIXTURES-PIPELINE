package types

type GetTeamsResponse struct {
	Data []Team `json:"data"`
}

type Team struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	ShortCode string `json:"short_code"`
	ImgPath   string `json:"image_path"`
	CountryId int    `json:"country_id"`
}

// TODO: Add type definition for GetScheduleResponse.
type GetScheduleResponse struct {
	Data []Stage `json:"data"`
}

type Stage struct {
	Id         int     `json:"id"`
	LeagueId   int     `json:"league_id"`
	SeasonId   int     `json:"season_id"`
	Name       string  `json:"name"`
	Finished   bool    `json:"finished"`
	IsCurrent  bool    `json:"is_current"`
	StartingAt string  `json:"starting_at"`
	EndingAt   string  `json:"ending_at"`
	Fixtures   []Match `json:"fixtures"`
}

type Match struct {
	Id           int     `json:"id"`
	LeagueId     int     `json:"league_id"`
	SeasonId     int     `json:"season_id"`
	Name         string  `json:"name"`
	StartingAt   string  `json:"starting_at"` // TODO: Figure out how to transform into proper Date type.
	ResultInfo   string  `json:"result_info"`
	Participants []Team  `json:"participants"`
	Scores       []Score `json:"scores"`
}

type Score struct {
	Id            int `json:"id"`
	MatchId       int `json:"fixture_id"`
	ParticipantId int `json:"participant_id"`
	Score         struct {
		Goals       int    `json:"goals"`
		Participant string `json:"participant"`
	} `json:"score"`
	Description string `json:"description"`
}
