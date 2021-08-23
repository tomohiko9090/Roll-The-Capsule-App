package handler

import (
	controller "GachaAPI/app/controller/user"
	"encoding/json"
	"fmt"
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
		fmt.Println("ioutil ReadAll error:", err)
		return err
	}

	// マッピング
	var userCreateRequest = new(UserCreateRequest)
	if err := json.Unmarshal(jsonBlob, userCreateRequest); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return err
	}

	// nameに入力がなかったらエラーを返す
	if len(userCreateRequest.Name) == 0 {
		return c.JSON(http.StatusInternalServerError, "エラー：ユーザーが作成されませんでした")
	}

	token, err := controller.CreateUser(userCreateRequest.Name) //②③④

	if err != nil {
		log.Error(err) // ターミナル上にエラーを表示する
		return c.JSON(http.StatusInternalServerError, "エラー：ユーザーが作成されませんでした")
	}
	userCreateResponse := UserCreateResponse{token}
	return c.JSON(http.StatusOK, userCreateResponse) //⑤
}

// GetUser 1.2. ユーザー取得
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
	nameStruct := UserGetResponse{user.Name}
	return c.JSON(http.StatusOK, nameStruct)
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
		fmt.Println("ioutil ReadAll error:", err)
		return err
	}
	// マッピング
	var userUpdateRequest = new(UserUpdateRequest)
	if err := json.Unmarshal(jsonBlob, userUpdateRequest); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return err
	}

	// nameに入力がなかったらエラーを返す
	if len(userUpdateRequest.Name) == 0 {
		return c.JSON(http.StatusInternalServerError, "エラー：ユーザーが作成されませんでした")
	}

	err = controller.UpdateUser(token, userUpdateRequest.Name) //②
	if err != nil {
		log.Error(err) // ターミナル上にエラーを表示する
		return c.JSON(http.StatusInternalServerError, "エラー：ユーザーが更新できませんでした")
	}
	return c.JSON(http.StatusOK, http.StatusOK) //③
}
