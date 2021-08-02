package view

import (
	"GachaAPI/controller"
	"GachaAPI/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// 1.1. ユーザー作成
type PostToken struct {
	Token  string  `json:"token"`
}
func UserPostHandler(c echo.Context)error{
	/*
	①POSTされたnameを受け取る(ハンドラー)
	②ランダムでtokenを発行(コントローラ)
	③idの発行)(モデル)
	④SQLにnameとtokenをインサート(モデル)
	⑤JSONでレスポンス(ハンドラー)
	*/
	name := c.FormValue("name") //①
	token_before := controller.RandomString(10) //②
	models.DBcontrollerPost(name, token_before) //③④
	token := PostToken{token_before}
	return c.JSON(http.StatusOK, token) //⑤
}

// 1.2. ユーザー取得
type GetName struct {
	Name  string  `json:"name"`
}
func UserGetHandler(c echo.Context) error {
	/*
	①tokenを受け取る(ハンドラー)
	②Userテーブルからnameを取得(モデル)
	③nameをレスポンス(ハンドラー)
	*/
	token := c.Request().Header.Get("Token") //①
	name_before := models.DBcontrollerGet(token) //②
	name := GetName{name_before}
	return c.JSON(http.StatusOK, name) //③
}

// 1.3. ユーザー更新
func UserPutHandler(c echo.Context) error {
	/*
	①tokenと新しいnameを受け取る(ハンドラー)
	②Userテーブルのnameを新しい名前に変更(モデル)
	③code200をレスポンス(ハンドラー)
	 */
	token := c.Request().Header.Get("Token") //①
	name := c.FormValue("name") //①
	models.DBcontrollerPut(token, name) //②
	return c.JSON(http.StatusOK, http.StatusOK) //③
}

// 2. ガチャの実行
func GachaPostHandler(c echo.Context) error{
	/*
	①tokenとtimesを受け取る(ハンドラー)
	②キャラ数取得(モデル)
	③1回以上ガチャを回す(コントローラ)
	④当たったキャラをDBにfor文でインサート(モデル)
	⑤キャラの名前とレア度とユーザー名をfor文で取得(モデル)
	⑥1個以上当たったキャラクターをレスポンス(ハンドラー)
	 */
	token := c.Request().Header.Get("Token") //①
	times := c.FormValue("times") //①
	number,_ := strconv.Atoi(times)
	total := models.DBcontrollerCaracterNo() //②
	ans_id := controller.Gacha(total, number) //③
	models.DBcontrollerInsert(token, ans_id)  //④
	results := models.DBcontrollerCharaGet(ans_id, token) //⑤
	return c.JSON(http.StatusOK, results) //⑥
}

// 3. キャラクター関連API
func CharacterGetHandler(c echo.Context) error{
	/*
	①tokenを受け取る(ハンドラー)
	②ユーザーIDの取得, ユーザーの所持キャラクター取得(モデル)
	③userCharacterID, characterID, name, rarity情報をレスポンス(ハンドラー)
	*/
	token := c.Request().Header.Get("Token") //①
	result := models.DBcontrollerCatalog(token)// ②
	return c.JSON(http.StatusOK, result) //③
}


