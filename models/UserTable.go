package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// ユーザー作成
var connection *sql.DB

// 1.1. ユーザー作成
func DBcontrollerPost(name string, token string){
	db, err := sql.Open("mysql", "root:taitasu2@/capsule")
	if err != nil {
	panic(err.Error())
	}
	connection = db
	defer db.Close()

	// idの発行
	var total int
	err = db.QueryRow("SELECT COUNT(*) FROM capsule.User").Scan(&total)
	if err != nil {
		panic(err.Error())
	}
	id := total + 1

	// nameとtokenをインサート
	ins, err := db.Prepare("INSERT INTO User(id,name,token) VALUES(?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	ins.Exec(id, name, token)
}

// 1.2. ユーザー取得
func DBcontrollerGet(token string) string {
	db, err := sql.Open("mysql", "root:taitasu2@/capsule")
	if err != nil {
		panic(err.Error())
	}
	connection = db
	defer db.Close()

	// 特定のユーザーnameを取得
	var name string
	err = db.QueryRow("SELECT name FROM capsule.User WHERE token = ?", token).Scan(&name)
	if err != nil {
		panic(err.Error())
	}
	return name
}

// 1.3. ユーザー更新
func DBcontrollerPut(name string, token string) {
	db, err := sql.Open("mysql", "root:taitasu2@/capsule")
	if err != nil {
		panic(err.Error())
	}
	connection = db
	defer db.Close()

	// Userテーブルのnameを新しい名前に変更
	ins, err := db.Prepare("UPDATE capsule.User SET name = ? WHERE token = ?")
	if err != nil {
		panic(err.Error())
	}
	ins.Exec(name, token)
}