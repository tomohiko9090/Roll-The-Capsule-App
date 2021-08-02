package main

import (
	"GachaAPI/view"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	//1. ユーザー関連API
	e.POST("/user/create", view.UserPostHandler)	// 1.1. ユーザー情報作成API
	e.GET("/user/get", view.UserGetHandler)	    // 1.2. ユーザー情報取得API
	e.PUT("/user/update", view.UserPutHandler)	    // 1.3. ユーザー情報更新API

	//2. ガチャ実行API
	e.POST("/gacha/draw", view.GachaPostHandler)

	//3. ユーザ所持キャラクター一覧取得API
	e.GET("/character/list", view.CharacterGetHandler)

	e.Logger.Fatal(e.Start(":8080"))
}


