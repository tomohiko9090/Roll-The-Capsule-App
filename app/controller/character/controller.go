package character

import (
	"GachaAPI/app/models/character"
)

var Error character.Characters

func GetCharacters(token string) (character.Characters, error) {
	// キャラクター一覧を取得する
	characters, err := character.GetCharacters(token)
	if err != nil {
		return Error, err
	}
	// トークンからユーザーidの取得

	return characters, nil
}
