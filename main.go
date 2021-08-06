package main

import (
	character "GachaAPI/app/handler/character"
	gacha "GachaAPI/app/handler/gacha"
	user "GachaAPI/app/handler/user"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// 1. ユーザー関連API
	e.POST("/user/create", user.CreateUser) // 1.1. ユーザー情報作成API
	e.GET("/user/get", user.GetUser)        // 1.2. ユーザー情報取得API
	e.PUT("/user/update", user.UpdateUser)  // 1.3. ユーザー情報更新API

	// 2. ガチャ実行API
	e.POST("/gacha/draw", gacha.DrowCharacters)

	// 3. ユーザ所持キャラクター一覧取得API
	e.GET("/character/list", character.GetCharacters)

	e.Logger.Fatal(e.Start(":8080"))
}
