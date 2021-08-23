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
	err := models.DB.QueryRow("SELECT * FROM capsule.User WHERE token = ?", token).
		Scan(&user.ID, &user.Name, &user.Token)
	if err != nil {
		return user, err
	}
	return user, nil
}

// InsertUser ユーザーネーム, トークン -> なし
func InsertUser(name string, token string) error {

	// 新しいidを作成
	userLength, err := getUserLength()
	if err != nil {
		return err
	}
	newId := userLength + 1

	// nameとtokenをインサート
	ins, err := models.DB.Prepare("INSERT INTO capsule.User(id,name,token) VALUES(?,?,?)")
	if err != nil {
		return err
	}
	ins.Exec(newId, name, token)
	return nil
}

// GetUserLength なし -> ユーザーユニーク数
func getUserLength() (int, error) {
	var userLength int
	err := models.DB.QueryRow("SELECT COUNT(*) FROM capsule.User").Scan(&userLength)
	if err != nil {
		return userLength, err
	}
	return userLength, nil
}

// UpdateUser トークン, ユーザーネーム ->　なし
func UpdateUser(token string, newName string) error {

	upd, err := models.DB.Prepare("UPDATE capsule.User SET name = ? WHERE token = ?")
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
