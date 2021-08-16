package character

import (
	"GachaAPI/app/models"
	_ "bytes"
	_ "fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Error Characters

// 3. キャラクター一覧取得
func GetCharacters(token string) (Characters, error) {

	var user User

	// トークンからユーザーidの取得
	err := models.DB.QueryRow("SELECT * FROM capsule.User WHERE token = ?", token).Scan(&user.ID, &user.Name, &user.Token)
	if err != nil {
		return Error, err
	}

	resultlist := []CharacterResults{}

	var (
		possess   Possess
		character Character
	)

	rows, err := models.DB.Query("SELECT * FROM capsule.Possess WHERE user_id = ?", user.ID)
	if err != nil {
		return Error, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&possess.UserID, &possess.UserCharacterID, &possess.CharacterID)
		//　キャラidからCharacterテーブルのname, rarityを取得
		err = models.DB.QueryRow("SELECT * FROM capsule.Character WHERE id = ?", possess.CharacterID).
			Scan(&character.CharacterID, &character.CharacterName, &character.Rarity, &character.Attack, &character.Defence, &character.Recovery)
		if err != nil {
			return Error, err
		}
		err = rows.Err()
		if err != nil {
			return Error, err
		}

		result := CharacterResults{
			UserCharacterID: possess.UserCharacterID,
			CharacterID:     possess.CharacterID,
			Name:            character.CharacterName,
			Rarity:          character.Rarity,
			Attack:          character.Attack,
			Defence:         character.Defence,
			Recovery:        character.Recovery,
		}
		resultlist = append(resultlist, result)
	}

	results := Characters{resultlist}
	return results, nil
}
