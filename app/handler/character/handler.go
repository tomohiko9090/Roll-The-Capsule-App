package characterResponse

import (
	character2 "GachaAPI/app/controller/character"
	"GachaAPI/app/models/character"
	"GachaAPI/app/models/userCharacters"
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

	var (
		userCharacterList []userCharacters.UserCharacter
		status            map[int]character.Character
		characterList     []CharacterResponse
	)

	userCharacterList, status, err := character2.GetCharacterDetail(token)
	if err != nil {
		log.Error(err) // ターミナル上にエラーを表示する
		return c.JSON(http.StatusNotAcceptable, "error：Do not exist the user")
	}

	for _, userCharacter := range userCharacterList {
		// マッピング
		userCharacter := CharacterResponse{
			UserCharacterID: userCharacter.UserCharacterID,
			CharacterID:     userCharacter.CharacterID,
			Name:            status[userCharacter.CharacterID].CharacterName,
			Rarity:          status[userCharacter.CharacterID].Rarity,
			Attack:          status[userCharacter.CharacterID].Attack,
			Defence:         status[userCharacter.CharacterID].Defence,
			Recovery:        status[userCharacter.CharacterID].Recovery,
		}
		characterList = append(characterList, userCharacter)
	}
	return c.JSON(http.StatusOK, CharacterListResponse{characterList}) //③
}
