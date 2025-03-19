package types

type TeamResponse struct {
	Data []Team `json:"data"`
}

type Team struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	ShortCode string `json:"short_code"`
	ImgPath   string `json:"image_path"`
	CountryId int    `json:"country_id"`
}
