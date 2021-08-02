package models

import (
	_ "bytes"
	"database/sql"
	_ "fmt"
	"log"
)

// 2. ガチャ実行(②キャラ数取得)
func DBcontrollerCaracterNo() int {
	db, err := sql.Open("mysql", "root:taitasu2@/capsule")
	if err != nil {
		panic(err.Error())
	}
	connection = db
	defer db.Close()

	// characterテーブルからキャラ数取得
	var total int
	err = db.QueryRow("SELECT COUNT(*) FROM capsule.Character").Scan(&total)
	if err != nil {
		log.Fatal(err)
	}
	return total
}

// 2. ガチャ実行(④当たったキャラのインサート)
func DBcontrollerInsert(token string, ans_id []int) {
	db, err := sql.Open("mysql", "root:taitasu2@/capsule")
	if err != nil {
		panic(err.Error())
	}
	connection = db
	defer db.Close()

	// tokenからidを引っ張ってくる
	var userid string
	err = db.QueryRow("SELECT id FROM capsule.User WHERE token = ?", token).Scan(&userid)
	if err != nil {
		panic(err.Error())
	}

	// usercharacterIDの長さを取得
	var total int
	err = db.QueryRow("SELECT COUNT(*) FROM capsule.Possess").Scan(&total)
	if err != nil {
		panic(err.Error())
	}
	usercharacterID := total + 1

	// 複数回当たったキャラをインサート
	for i := 0; i < len(ans_id); i++{
		stmtInsert, err := db.Prepare("INSERT INTO Possess(user_id, usercharacterID, character_id) VALUES(?, ?, ?)")
		if err != nil {
			panic(err.Error())
		}
		defer stmtInsert.Close()

		result1, err := stmtInsert.Exec(userid, usercharacterID+i, ans_id[i])
		if err != nil {
			panic(err.Error())
		}
		_, err = result1.LastInsertId()
		if err != nil {
			panic(err.Error())
		}
	}
}

// 2. ガチャ実行(⑤当たったキャラ情報を取得)
type Characters struct {
	Results     []Results   `json:"results"`
}

type Results struct {
	UserID        int    `json:"userID"`
	UserName      string `json:"userName"`
	CharacterName string `json:"characterName"`
	Rarity        string `json:"rarity"`
}

func DBcontrollerCharaGet(ans_id []int, token string) Characters {
	db, err := sql.Open("mysql", "root:taitasu2@/capsule")
	if err != nil {
		panic(err.Error())
	}
	connection = db
	defer db.Close()

	var (
		userid int
		username string
		charactername string
		rarity string
	)

	result_list := []Results{}

	// 複数のキャラnameとrarityを取得
	for i := 0; i < len(ans_id); i++ {
		rows, _ := db.Query("SELECT name,rarity FROM capsule.Character WHERE id = ?", ans_id[i])
		defer rows.Close()

		rows.Next()
		rows.Scan(&charactername, &rarity)

		rows2, _ := db.Query("SELECT id,name FROM capsule.User WHERE token = ?", token)
		defer rows2.Close()

		rows2.Next()
		rows2.Scan(&userid, &username)

		result := Results {
			UserID: userid,
			UserName: username,
			CharacterName: charactername,
			Rarity:rarity,
		}

		result_list = append(result_list, result)
	}

	results := Characters{result_list}
	return results
}

// 3. キャラクター一覧取得
type Characters2 struct {
	Results2 []Results2  `json:"characters"`
}

type Results2 struct {
	UserCharacterID int    `json:"userID"`
	CharacterID     int    `json:"userName"`
	Name            string `json:"characterName"`
	Rarity          string `json:"rarity"`
}

func DBcontrollerCatalog(token string) Characters2 {

	db, err := sql.Open("mysql", "root:taitasu2@/capsule")
	if err != nil {
		panic(err.Error())
	}
	connection = db
	defer db.Close()

	// トークンからユーザーidの取得
	var userid string
	err = db.QueryRow("SELECT id FROM capsule.User WHERE token = ?", token).Scan(&userid)
	if err != nil {
		panic(err.Error())
	}

	// ユーザーidからキャラidの一覧とuserCharacterIDを取得
	var (
		userCharacterID int
		characterid     int
		name            string
		rarity          string
	)

	result_list2 := []Results2{}

	rows, err := db.Query("SELECT userCharacterID, character_id FROM capsule.Possess WHERE user_id = ?", userid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&userCharacterID, &characterid)
		//　キャラidからCharacterテーブルのname, rarityを取得
		err = db.QueryRow("SELECT name, rarity  FROM capsule.Character WHERE id = ?", characterid).
			Scan(&name, &rarity)
		if err != nil {
			log.Fatal(err)
		}

		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

		result2 := Results2{
			UserCharacterID: userCharacterID,
			CharacterID: characterid,
			Name: name,
			Rarity: rarity,
		}
		result_list2 = append(result_list2, result2)
	}
	results2 := Characters2{result_list2}
	return results2
}