package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// 実行時にDBへ接続する(サーバーが立ち上がった段階だけコネクションが作られる)
func init() {
	var err error
	// ローカル環境時
	// user:password@tcp(container-name:port)/dbname
	//DB, err = sql.Open("mysql", "root:taitasu2@/capsule")
	// DB, err = sql.Open("mysql", "root:xxxx@tcp(mysql-container:3307)/capsule")
	DB, err = sql.Open("mysql", "root:xxxx@tcp(mysql:3306)/capsule?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		panic(err)

	}
}
