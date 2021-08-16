package gacha

// DB取得構造体
type User struct {
	ID    int
	Name  string
	Token string
}
type Character struct {
	CharacterID int
	Name        string
	Rarity      string
	Attack      int
	Defence     int
	Recovery    int
}

// レスポンス構造体
type GachaResults struct {
	UserID        int
	UserName      string
	CharacterName string
	Rarity        string
	Attack        int
	Defence       int
	Recovery      int
}

type Characters struct {
	Results []GachaResults `json:"results"`
}
