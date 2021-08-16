package handler

import (
	controller "GachaAPI/app/controller/user"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

// 1.1. ユーザー作成
func CreateUser(c echo.Context) error {
	/*
		①POSTされたnameを受け取る(ハンドラー)
		②ランダムでtokenを発行(コントローラ)
		③idの発行(モデル)
		④SQLにnameとtokenをインサート(モデル)
		⑤JSONでレスポンス(ハンドラー)
	*/
	name := c.FormValue("name")               //①
	token, err := controller.CreateUser(name) //②③④
	if err != nil {
		log.Error(err) // ターミナル上にエラーを表示する
		return c.JSON(http.StatusInternalServerError, "エラー：ユーザーが作成されませんでした")
	}
	tokenStruct := PostToken{token}
	return c.JSON(http.StatusOK, tokenStruct) //⑤
}

// 1.2. ユーザー取得
func GetUser(c echo.Context) error {
	/*
		①tokenを受け取る(ハンドラー)
		②Userテーブルからnameを取得(モデル)
		③nameをレスポンス(ハンドラー)
	*/
	token := c.Request().Header.Get("Token") //①
	user, err := controller.GetUser(token)   //②
	if err != nil {
		log.Error(err) // ターミナル上にエラーを表示する
		return c.JSON(http.StatusInternalServerError, "エラー：ユーザーが取得できませんでした")
	}
	nameStruct := GetName{user.Name}
	c.JSON(http.StatusOK, nameStruct) //③
	return nil
}

// 1.3. ユーザー更新
func UpdateUser(c echo.Context) error {
	/*
		①tokenと新しいnameを受け取る(ハンドラー)
		②Userテーブルのnameを新しい名前に変更(モデル)
		③code200をレスポンス(ハンドラー)
	*/
	token := c.Request().Header.Get("Token")     //①
	newName := c.FormValue("name")               //①
	err := controller.UpdateUser(token, newName) //②
	if err != nil {
		log.Error(err) // ターミナル上にエラーを表示する
		return c.JSON(http.StatusInternalServerError, "エラー：ユーザーが更新できませんでした")
	}
	return c.JSON(http.StatusOK, http.StatusOK) //③
}
