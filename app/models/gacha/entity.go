package gacha

// User DB取得構造体
type User struct {
	ID    int
	Name  string
	Token string
}

// Character DB取得構造体
type Character struct {
	CharacterID int
	Name        string
	Rarity      string
	Attack      int
	Defence     int
	Recovery    int
}

// Results レスポンス構造体
type Results struct {
	UserID        int
	UserName      string
	CharacterName string
	Rarity        string
	Attack        int
	Defence       int
	Recovery      int
}

// Characters レスポンス構造体
type Characters struct {
	Results []Results `json:"results"`
}
