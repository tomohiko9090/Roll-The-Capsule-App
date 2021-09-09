package gachahandler

import (
	controller "GachaAPI/app/controller/gacha"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/gommon/log"

	"github.com/labstack/echo"
)

// DrowCharacters -> 2. ガチャの実行
func DrowCharacters(c echo.Context) error {
	/*
		①tokenとtimesを受け取る(ハンドラー)
		②キャラ数取得(モデル)
		③1回以上ガチャを回す(コントローラ)
		④当たったキャラをDBにfor文でインサート(モデル)
		⑤キャラの名前、レア度などをfor文で取得(モデル)
		⑥1個以上当たったキャラクターをレスポンス(ハンドラー)
	*/
	// ①
	token := c.Request().Header.Get("X-token")

	// Bodyを読む
	jsonBlob, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, "error：ServerError")
	}
	// マッピング
	var gachaDrawRequest = new(GachaDrawRequest)
	if err := json.Unmarshal(jsonBlob, gachaDrawRequest); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, "error：ServerError")
	}

	if gachaDrawRequest.Times == 0 {
		return c.JSON(http.StatusBadRequest, "error：0 times drawing is bad")
	}

	// ②
	characterLength, err := controller.GetCharacterLength()
	if err != nil {
		log.Error(err) // ターミナル上にエラーを表示する
		return c.JSON(http.StatusInternalServerError, "error:ServerError")
	}

	// ③④⑤
	resultCharacterIDs, status, err := controller.DrowCharacter(characterLength, token, gachaDrawRequest.Times)
	if err != nil {
		log.Error(err) // ターミナル上にエラーを表示する
		return c.JSON(http.StatusNotAcceptable, "error：Do not exist the user")
	}

	var GachaDrawLIst []GachaResult
	for _, resultCharacterID := range resultCharacterIDs {
		// マッピング
		gachaResult := GachaResult{
			CharacterID: status[resultCharacterID].CharacterID,
			Name:        status[resultCharacterID].CharacterName,
		}
		GachaDrawLIst = append(GachaDrawLIst, gachaResult)
	}
	return c.JSON(http.StatusOK, GachaDrawResponse{GachaDrawLIst}) //⑥
}
