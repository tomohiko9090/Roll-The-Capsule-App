package user

import (
	"GachaAPI/app/models/user"
	"math/rand"
	"time"
)

var Error user.User

// 1.1. ユーザー作成
func CreateUser(name string) (string, error) {

	// tokenを作成する
	token := randomString(10)
	// userModelをDBにINSERTする
	err := user.InsertUser(name, token)
	if err != nil {
		return "", err
	}
	return token, nil
}

// ランダムでtokenを発行
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
func randomString(len int) string {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}

// 1.2. ユーザー取得
func GetUser(token string) (user.User, error) {
	// ユーザーnameを取得する
	getUser, err := user.SelectUser(token)
	if err != nil {
		return Error, err
	}
	return getUser, nil
}

// 1.3. ユーザー更新
func UpdateUser(token string, newName string) error {

	// tokenで認証し、ユーザーnameを変更する
	err := user.UpdateUser(token, newName)
	if err != nil {
		return err
	}
	return err
}
