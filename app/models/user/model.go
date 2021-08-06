package user

import (
	"GachaAPI/app/models"
	_ "database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// InsertUser ユーザー作成
func InsertUser(name string, token string) error {

	// idの発行
	var total int
	err := models.DB.QueryRow("SELECT COUNT(*) FROM capsule.User").Scan(&total)
	if err != nil {
		log.Fatal(err)
	}
	id := total + 1

	// nameとtokenをインサート
	ins, err := models.DB.Prepare("INSERT INTO User(id,name,token) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	ins.Exec(id, name, token)
	return nil
}

// 1.2. ユーザー取得
func SelectUser(token string) string {

	// 特定のユーザーnameを取得
	var name string
	err := models.DB.QueryRow("SELECT name FROM capsule.User WHERE token = ?", token).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	return name
}

// 1.3. ユーザー更新
func UpdateUser(name string, token string) {

	// Userテーブルのnameを新しい名前に変更
	ins, err := models.DB.Prepare("UPDATE capsule.User SET name = ? WHERE token = ?")

	if err != nil {
		log.Fatal(err)
	}

	t1, err := ins.Exec(name, token)
	if err != nil {
		fmt.Printf("type: %T\n", err)
		fmt.Printf("value: %v\n", err)
		return
	}
	fmt.Println(t1)
}
