package handler

import (
	controller "GachaAPI/app/controller/user"
	"net/http"

	"github.com/labstack/echo"
)

// 1.1. ユーザー作成
type PostToken struct {
	Token string `json:"token"`
}

func CreateUser(c echo.Context) error {
	/*
		①POSTされたnameを受け取る(ハンドラー)
		②ランダムでtokenを発行(コントローラ)
		③idの発行(モデル)
		④SQLにnameとtokenをインサート(モデル)
		⑤JSONでレスポンス(ハンドラー)
	*/
	name := c.FormValue("name")             //①
	token, _ := controller.CreateUser(name) //②③④
	tokenStruct := PostToken{token}
	return c.JSON(http.StatusOK, tokenStruct) //⑤
}

// 1.2. ユーザー取得
type GetName struct {
	Name string `json:"name"`
}

func GetUser(c echo.Context) error {
	/*
		①tokenを受け取る(ハンドラー)
		②Userテーブルからnameを取得(モデル)
		③nameをレスポンス(ハンドラー)
	*/
	token := c.Request().Header.Get("Token") //①
	name := controller.GetUser(token)        //②
	nameStruct := GetName{name}
	return c.JSON(http.StatusOK, nameStruct) //③
}

// 1.3. ユーザー更新
func UpdateUser(c echo.Context) error {
	/*
		①tokenと新しいnameを受け取る(ハンドラー)
		②Userテーブルのnameを新しい名前に変更(モデル)
		③code200をレスポンス(ハンドラー)
	*/
	token := c.Request().Header.Get("Token")    //①
	newName := c.FormValue("name")              //①
	controller.UpdateUser(token, newName)       //②
	return c.JSON(http.StatusOK, http.StatusOK) //③
}
