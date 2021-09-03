package character

import (
	"GachaAPI/app/models/character"
	user2 "GachaAPI/app/models/user"
	"GachaAPI/app/models/userCharacters"
)

// GetUserCharacterList トークン -> UserCharacterテーブル構造体の配列
func GetUserCharacterList(token string) ([]userCharacters.UserCharacter, error) {
	var userCharacterList []userCharacters.UserCharacter
	// キャラクター情報を取得
	user, err := user2.SelectUser(token)
	if err != nil {
		return userCharacterList, err
	}
	userCharacterList, err = userCharacters.SelectUserCharacterList(user.ID)
	if err != nil {
		return userCharacterList, err
	}
	return userCharacterList, nil
}

// GetCharacter UserCharacterテーブル構造体 -> characterテーブル構造体
func GetCharacter(userCharacter userCharacters.UserCharacter) (character.Character, error) {
	character, err := character.SelectCharacter(userCharacter.CharacterID)
	if err != nil {
		return character, err
	}
	return character, nil
}
