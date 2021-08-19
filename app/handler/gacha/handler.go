package gachahandler

import (
	controller "GachaAPI/app/controller/gacha"
	"net/http"
	"strconv"

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
		⑤キャラの名前とレア度とユーザー名をfor文で取得(モデル)
		⑥1個以上当たったキャラクターをレスポンス(ハンドラー)
	*/
	token := c.Request().Header.Get("Token")                //①
	times := c.FormValue("times")                           //①
	drows, _ := strconv.Atoi(times)                         //　文字列から整数に変換
	results, err := controller.DrowCharacters(token, drows) //②③④⑤
	if err != nil {
		log.Error(err) // ターミナル上にエラーを表示する
		return c.JSON(http.StatusInternalServerError, "エラー:ガチャが実行されませんでした")
	}

	return c.JSON(http.StatusOK, results) //⑥
}
