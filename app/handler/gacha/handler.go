package handler

import (
	controller "GachaAPI/app/controller/gacha"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// 2. ガチャの実行
func DrowCharacters(c echo.Context) error {
	/*
		①tokenとtimesを受け取る(ハンドラー)
		②キャラ数取得(モデル)
		③1回以上ガチャを回す(コントローラ)
		④当たったキャラをDBにfor文でインサート(モデル)
		⑤キャラの名前とレア度とユーザー名をfor文で取得(モデル)
		⑥1個以上当たったキャラクターをレスポンス(ハンドラー)
	*/
	token := c.Request().Header.Get("Token")           //①
	times := c.FormValue("times")                      //①
	drows, _ := strconv.Atoi(times)                    //　文字列から整数に変換
	results := controller.DrowCharacters(token, drows) //②③④⑤

	return c.JSON(http.StatusOK, results) //⑥
}
