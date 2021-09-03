package handler

import (
	controller "GachaAPI/app/controller/user"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

// CreateUser 1.1. ユーザー作成
func CreateUser(c echo.Context) error {
	/*
		①POSTされたnameを受け取る(ハンドラー)
		②ランダムでtokenを発行(コントローラ)
		③idの発行(モデル)
		④SQLにnameとtokenをインサート(モデル)
		⑤JSONでレスポンス(ハンドラー)
	*/

	// ①
	// Bodyを読む
	jsonBlob, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, "error：ServerError1")
	}

	// マッピング
	var userCreateRequest = new(UserCreateRequest)
	if err := json.Unmarshal(jsonBlob, userCreateRequest); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, "error：ServerError2")
	}

	// nameに入力がなかったらエラーを返す
	if len(userCreateRequest.Name) <= 2 || 10 <= len(userCreateRequest.Name) {
		err := "error：Could not create user because length of a character string is bad"
		log.Error(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	token, err := controller.CreateUser(userCreateRequest.Name) //②③④

	if err != nil {
		log.Error(err) // ターミナル上にエラーを表示する
		return c.JSON(http.StatusInternalServerError, "error：ServerError3")
	}
	return c.JSON(http.StatusOK, UserCreateResponse{token}) //⑤
}

// GetUser 1.2. ユーザー取得
func GetUser(c echo.Context) error {
	/*
		①tokenを受け取る(ハンドラー)
		②Userテーブルからnameを取得(モデル)
		③nameをレスポンス(ハンドラー)
	*/
	token := c.Request().Header.Get("X-token") //①
	user, err := controller.GetUser(token)     //②
	if err != nil {
		log.Error(err) // ターミナル上にエラーを表示する
		return c.JSON(http.StatusNotAcceptable, "error：Do not exist the user")
	}
	return c.JSON(http.StatusOK, UserGetResponse{user.Name})
}

// UpdateUser 1.3. ユーザー更新
func UpdateUser(c echo.Context) error {
	/*
		①tokenと新しいnameを受け取る(ハンドラー)
		②Userテーブルのnameを新しい名前に変更(モデル)
		③code200をレスポンス(ハンドラー)
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
	var userUpdateRequest = new(UserUpdateRequest)
	if err := json.Unmarshal(jsonBlob, userUpdateRequest); err != nil {
		log.Error(err)
		return err
	}

	// nameに入力がなかったらエラーを返す
	if len(userUpdateRequest.Name) <= 2 || 10 <= len(userUpdateRequest.Name) {
		err := "error：Could not update because length of a character string is bad"
		log.Error(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	err = controller.UpdateUser(token, userUpdateRequest.Name) //②
	if err != nil {
		log.Error(err) // ターミナル上にエラーを表示する
		return c.JSON(http.StatusUnauthorized, "error：Could not Authenticate")
	}
	return c.JSON(http.StatusOK, http.StatusOK) //③
}
