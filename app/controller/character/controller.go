package character2

import (
	"GachaAPI/app/models/character"
	user2 "GachaAPI/app/models/user"
	"GachaAPI/app/models/userCharacters"
	"database/sql"
)

// GetCharacterDetail トークン -> UserCharacterテーブル構造体の配列, characterテーブル
func GetCharacterDetail(token string) ([]userCharacters.UserCharacter, map[int]character.Character, error) {
	var (
		userCharacterList []userCharacters.UserCharacter
		characterRows     *sql.Rows
	)
	status := make(map[int]character.Character)

	// ユーザーIDを取得
	user, err := user2.SelectUser(token)
	if err != nil {
		return userCharacterList, status, err
	}

	// キャラクター所持情報を取得
	userCharacterList, err = userCharacters.SelectUserCharacterList(user.ID)
	if err != nil {
		return userCharacterList, status, err
	}

	// キャラクターテーブルをまるごと取得
	characterRows, err = character.SelectCharacterTable()
	if err != nil {
		return userCharacterList, status, err
	}
	defer characterRows.Close()

	// キャラクターテーブルをmapにする(キャラクターidをキー, キャラクターの詳細情報をバリューに)
	for characterRows.Next() {
		var character character.Character
		err = characterRows.Scan(
			&character.CharacterID,
			&character.CharacterName,
			&character.Rarity,
			&character.Attack,
			&character.Defence,
			&character.Recovery)
		if err != nil {
			return userCharacterList, status, err
		}
		status[character.CharacterID] = character
	}
	return userCharacterList, status, err
}
