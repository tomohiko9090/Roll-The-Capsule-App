package user

import (
	"GachaAPI/app/models"
	_ "database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Error User

// 1.1. ユーザー作成
func InsertUser(name string, token string) error {

	// idの発行
	var total int
	err := models.DB.QueryRow("SELECT COUNT(*) FROM capsule.User").Scan(&total)
	if err != nil {
		return err
	}
	id := total + 1

	// nameとtokenをインサート
	ins, err := models.DB.Prepare("INSERT INTO User(id,name,token) VALUES(?,?,?)")
	if err != nil {
		return err
	}
	ins.Exec(id, name, token)
	return nil
}

// 1.2. ユーザー取得
func SelectUser(token string) (User, error) {
	var user User
	err := models.DB.QueryRow("SELECT * FROM capsule.User WHERE token = ?", token).Scan(&user.ID, &user.Name, &user.Token)
	if err != nil {
		return Error, err
	}
	return user, nil
}

// 1.3. ユーザー更新
func UpdateUser(name string, token string) error {

	// Userテーブルのnameを新しい名前に変更
	ins, err := models.DB.Prepare("UPDATE capsule.User SET name = ? WHERE token = ?")

	if err != nil {
		return err
	}

	t1, err := ins.Exec(name, token)
	if err != nil {
		return err
	}
	fmt.Println(t1)
	return nil
}
