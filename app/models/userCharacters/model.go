package userCharacters

import (
	"GachaAPI/app/models"
)

// SelectUserCharacterList ユーザーID -> UserCharacterテーブル構造体の配列
func SelectUserCharacterList(userId int) ([]UserCharacter, error) {
	var userCharacterList []UserCharacter
	rows, err := models.DB.Query("SELECT * FROM UserCharacters WHERE user_id = ?", userId)
	if err != nil {
		return userCharacterList, err
	}
	defer rows.Close()

	var userCharacter UserCharacter
	for rows.Next() {
		err = rows.Scan(&userCharacter.UserCharacterID, &userCharacter.UserID, &userCharacter.CharacterID)
		if err != nil {
			return userCharacterList, err
		}
		userCharacterList = append(userCharacterList, userCharacter)
	}
	return userCharacterList, nil
}

// InsertUserCharacter ユーザーID, ガチャ途中実行数, 当たったキャラクターID -> なし
func InsertUserCharacter(userId int, resultCharacterID int) error {
	stmtInsert, err := models.DB.Prepare("INSERT INTO UserCharacters(user_id, character_id) VALUES(?, ?)")
	if err != nil {
		return err
	}
	defer stmtInsert.Close()

	result1, err := stmtInsert.Exec(userId, resultCharacterID)

	if err != nil {
		return err
	}
	_, err = result1.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}
