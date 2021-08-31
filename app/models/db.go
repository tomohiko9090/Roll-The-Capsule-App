package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// 実行時にDBへ接続する(サーバーが立ち上がった段階だけコネクションが作られる)
func init() {
	var err error
	// DB, err = sql.Open("mysql", "root:taitasu2@/capsule")
	DB, err = sql.Open("mysql", "root:taitasu2@tcp(db:3307)/capsule")
	// DB, err = sql.Open("mysql", "root:taitasu2@tcp(db:3306)/capsule")
	if err != nil {
		panic(err)
	}
}
