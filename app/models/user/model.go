package user

import (
	"GachaAPI/app/models"
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// SelectUser トークン -> Userテーブル構造体
func SelectUser(token string) (User, error) {
	var user User
	// tokenからidを引っ張ってくる
	err := models.DB.QueryRow("SELECT * FROM User WHERE token = ?", token).
		Scan(&user.ID, &user.Name, &user.Token)
	if err != nil {
		return user, err
	}
	return user, nil
}

// InsertUser ユーザーネーム, トークン -> なし
func InsertUser(name string, token string) error {
	// nameとtokenをインサート
	ins, err := models.DB.Prepare("INSERT INTO User(id, name,token) VALUES(?,?,?)")
	if err != nil {
		return err
	}
	ins.Exec(1, name, token)
	return nil
}

// UpdateUser トークン, ユーザーネーム ->　なし
func UpdateUser(token string, newName string) error {

	upd, err := models.DB.Prepare("UPDATE User SET name = ? WHERE token = ?")
	if err != nil {
		return err
	}
	defer upd.Close()

	_, err = upd.Exec(newName, token)
	if err != nil {
		return err
	}
	return nil
}
