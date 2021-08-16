package characterResponse

import (
	controller "GachaAPI/app/controller/character"
	"net/http"

	"github.com/labstack/gommon/log"

	"github.com/labstack/echo"
)

// 3. キャラクター関連API
func GetCharacters(c echo.Context) error {
	/*
		①tokenを受け取る(ハンドラー)
		②ユーザーIDの取得, ユーザーの所持キャラクター取得(モデル)
		③userCharacterID, characterID, name, rarity情報をレスポンス(ハンドラー)
	*/
	token := c.Request().Header.Get("Token")           //①
	characters, err := controller.GetCharacters(token) // ②
	if err != nil {
		log.Error(err) // ターミナル上にエラーを表示する
		return c.JSON(http.StatusInternalServerError, "エラー：キャラクター一覧を取得できませんでした")
	}
	return c.JSON(http.StatusOK, characters) //③
}
