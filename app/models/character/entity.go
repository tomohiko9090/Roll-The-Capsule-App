package character

// DB取得構造体
type User struct {
	ID    int
	Name  string
	Token string
}
type Possess struct {
	UserID          int
	UserCharacterID int
	CharacterID     int
}

type Character struct {
	CharacterID   int
	CharacterName string
	Rarity        string
	Attack        int
	Defence       int
	Recovery      int
}

// レスポンス構造体
type CharacterResults struct {
	UserCharacterID int    // `json:"userID"`
	CharacterID     int    // `json:"userName"`
	Name            string // `json:"characterName"`
	Rarity          string // `json:"rarity"`
	Attack          int    // `json:"attack"`
	Defence         int    // `json:"defence"`
	Recovery        int    // `json:"recovery"`
}

type Characters struct {
	Results []CharacterResults `json:"characters"`
}
