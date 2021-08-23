package gachahandler

import (
	controller "GachaAPI/app/controller/gacha"
	"encoding/json"
	"fmt"
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
		fmt.Println("ioutil ReadAll error:", err)
		return err
	}
	// マッピング
	var gachaDrawRequest = new(GachaDrawRequest)
	if err := json.Unmarshal(jsonBlob, gachaDrawRequest); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return err
	}

	// nameに入力がなかったらエラーを返す
	if gachaDrawRequest.Times == 0 {
		return c.JSON(http.StatusInternalServerError, "エラー:ガチャが実行されませんでした")
	}

	//　②
	characterLength, err := controller.GetCharacterLength()
	if err != nil {
		log.Error(err) // ターミナル上にエラーを表示する
		return c.JSON(http.StatusInternalServerError, "エラー:ガチャが実行されませんでした")
	}

	// ③④⑤
	var resultCharacterIDs []int
	resultCharacterIDs, err = controller.DrowCharacter(characterLength, token, gachaDrawRequest.Times)
	if err != nil {
		log.Error(err) // ターミナル上にエラーを表示する
		return c.JSON(http.StatusInternalServerError, "エラー:ガチャが実行されませんでした")
	}

	var GachaDrawLIst []GachaResult
	for i := 0; i < gachaDrawRequest.Times; i++ {

		// キャラクター情報の取得
		character, err := controller.GetCharacter(resultCharacterIDs[i])
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "エラー:ガチャが実行されませんでした")
		}
		//マッピング
		gachaResult := GachaResult{
			CharacterID: character.CharacterID,
			Name:        character.CharacterName,
		}
		GachaDrawLIst = append(GachaDrawLIst, gachaResult)
	}
	gachaDrawResponse := GachaDrawResponse{GachaDrawLIst}
	return c.JSON(http.StatusOK, gachaDrawResponse) //⑥
}
