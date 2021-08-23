package character

// Character DB取得構造体
type Character struct {
	CharacterID   int    `json:"characterID"`
	CharacterName string `json:"characterName"`
	Rarity        string `json:"rarity"`
	Attack        int    `json:"attack"`
	Defence       int    `json:"defence"`
	Recovery      int    `json:"recovery"`
}
