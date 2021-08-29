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
	userCharacterList, err := controller.GetUserCharacterList(token)
	if err != nil {
		log.Error(err) // ターミナル上にエラーを表示する
		return c.JSON(http.StatusNotAcceptable, "error：Do not exist the user")
	}

	var characterList []CharacterResponse
	for _, userCharacter := range userCharacterList {
		character, err := controller.GetCharacter(userCharacter)
		if err != nil {
			log.Error(err) // ターミナル上にエラーを表示する
			return c.JSON(http.StatusInternalServerError, "error：ServerError")
		}
		// マッピング
		userCharacter := CharacterResponse{
			UserCharacterID: userCharacter.UserCharacterID,
			CharacterID:     userCharacter.CharacterID,
			Name:            character.CharacterName,
			Rarity:          character.Rarity,
			Attack:          character.Attack,
			Defence:         character.Defence,
			Recovery:        character.Recovery,
		}
		characterList = append(characterList, userCharacter)
	}
	return c.JSON(http.StatusOK, CharacterListResponse{characterList}) //③
}
