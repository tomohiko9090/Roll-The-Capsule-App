package character

import (
	"GachaAPI/app/models/character"
)

func GetCharacters(token string) (string character.Characters) {
	// キャラクター一覧を取得する
	characters, _ := character.GetCharacters(token)
	// トークンからユーザーidの取得

	return characters
}
