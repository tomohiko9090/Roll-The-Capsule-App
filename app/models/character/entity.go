package character

// User DB取得構造体
type User struct {
	ID    int
	Name  string
	Token string
}

// Possess DB取得構造体
type Possess struct {
	UserID          int
	UserCharacterID int
	CharacterID     int
}

// Character DB取得構造体
type Character struct {
	CharacterID   int
	CharacterName string
	Rarity        string
	Attack        int
	Defence       int
	Recovery      int
}

// Results レスポンス構造体
type Results struct {
	UserCharacterID int
	CharacterID     int
	Name            string
	Rarity          string
	Attack          int
	Defence         int
	Recovery        int
}

// Characters レスポンス構造体
type Characters struct {
	Results []Results `json:"characters"`
}
