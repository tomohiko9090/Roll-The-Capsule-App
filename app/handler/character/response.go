package characterResponse

type UserCharacter struct {
	UserCharacterID int    `json:"userCharacterID"`
	CharacterID     int    `json:"characterID"`
	Name            string `json:"name"`
	Rarity          string `json:"rarity"`
	Attack          int    `json:"attack"`
	Defence         int    `json:"defence"`
	Recovery        int    `json:"recovery"`
}

type CharacterListResponse struct {
	Characters []UserCharacter `json:"characters"`
}
