package gachahandler

type GachaResult struct {
	CharacterID int    `json:"characterID"`
	Name        string `json:"name"`
}

type GachaDrawResponse struct {
	Results []GachaResult `json:"results"`
}
