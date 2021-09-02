package character

import (
	"GachaAPI/app/models"
	_ "bytes"
	_ "fmt"

	_ "github.com/go-sql-driver/mysql"
)

// SelectCharacter キャラクターID -> Characterテーブル構造体
func SelectCharacter(characterId int) (Character, error) {
	var character Character
	err := models.DB.QueryRow("SELECT * FROM Characters WHERE id = ?", characterId).
		Scan(&character.CharacterID, &character.CharacterName, &character.Rarity, &character.Attack, &character.Defence, &character.Recovery)

	if err != nil {
		return character, err
	}
	return character, nil
}

// GetCharacterLength なし -> キャラクターのユニーク数
func GetCharacterLength() (int, error) {
	var characterLength int
	err := models.DB.QueryRow("SELECT COUNT(*) FROM Characters").Scan(&characterLength)
	if err != nil {
		return characterLength, err
	}
	return characterLength, nil
}
