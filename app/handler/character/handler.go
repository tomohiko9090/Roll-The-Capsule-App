package characterResponse

import (
	controller "GachaAPI/app/controller/character"
	"net/http"

	"github.com/labstack/gommon/log"

	"github.com/labstack/echo"
)

// GetCharacters -> 3. キャラクター関連API
func GetCharacters(c echo.Context) error {
	/*
		①tokenを受け取る(ハンドラー)
		②ユーザーIDの取得, ユーザーの所持キャラクター取得(モデル)
		③レスポンス(ハンドラー)
	*/

	// ①
	token := c.Request().Header.Get("X-token")

	// ②
	possessList, err := controller.GetPossessList(token)
	if err != nil {
		log.Error(err) // ターミナル上にエラーを表示する
		return c.JSON(http.StatusInternalServerError, "エラー：キャラクター一覧を取得できませんでした")
	}
	possessLength := len(possessList)

	var characterList []UserCharacter
	for i := 0; i < possessLength; i++ {
		character, err := controller.GetCharacter(possessList[i])
		if err != nil {
			log.Error(err) // ターミナル上にエラーを表示する
			return c.JSON(http.StatusInternalServerError, "エラー：キャラクター一覧を取得できませんでした")
		}
		// マッピング
		userCharacter := UserCharacter{
			UserCharacterID: possessList[i].UserCharacterID,
			CharacterID:     possessList[i].CharacterID,
			Name:            character.CharacterName,
			Rarity:          character.Rarity,
			Attack:          character.Attack,
			Defence:         character.Defence,
			Recovery:        character.Recovery,
		}
		characterList = append(characterList, userCharacter)
	}
	characterListResponse := CharacterListResponse{characterList}

	return c.JSON(http.StatusOK, characterListResponse) //③
}
