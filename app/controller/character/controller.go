package character

import (
	"GachaAPI/app/models/character"
	"GachaAPI/app/models/possess"
	user2 "GachaAPI/app/models/user"
)

// GetPossessList トークン -> possessテーブル構造体の配列
func GetPossessList(token string) ([]possess.Possess, error) {
	var possessList []possess.Possess
	// キャラクター情報を取得
	user, err := user2.SelectUser(token)
	if err != nil {
		return possessList, err
	}
	possessList, err = possess.SelectPossessList(user.ID)
	if err != nil {
		return possessList, err
	}
	return possessList, nil
}

// GetCharacter possessテーブル構造体 -> characterテーブル構造体
func GetCharacter(possess possess.Possess) (character.Character, error) {
	character, err := character.SelectCharacter(possess.CharacterID)
	if err != nil {
		return character, err
	}
	return character, nil
}
