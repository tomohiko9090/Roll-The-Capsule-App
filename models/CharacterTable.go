package models

import (
	_ "bytes"
	"database/sql"
	"encoding/json"
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
type Animal struct {
	UserID int
	UserName string
	CharacterName string
	Rarity string
}
type Animals []Animal

func DBcontrollerCharaGet(ans_id []int, token string) string {
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

	var animals Animals
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

		kame := Animal {
			UserID: userid,
			UserName: username,
			CharacterName: charactername,
			Rarity: rarity,
		}

		animals = append(animals, kame)
	}
	str, _ := json.Marshal(animals)
	//a := string(str)
	//b := a[:strings.Index(a, "\\")]
	return string(str)
}

// 3. キャラクター一覧取得
type Animala struct {
	UserCharacterID int
	CharacterID int
	Name string
	Rarity string
}
type Animalsa []Animala

func DBcontrollerCatalog(token string) string {

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
		characterid int
		name   string
		rarity string
	)

	rows, err := db.Query("SELECT userCharacterID, character_id FROM capsule.Possess WHERE user_id = ?", userid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var animalsa Animalsa

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

		kamea := Animala {
			UserCharacterID: userCharacterID,
			CharacterID: characterid,
			Name: name,
			Rarity: rarity,
		}

		animalsa = append(animalsa, kamea)
	}

	str2, _ := json.Marshal(animalsa)
	return string(str2)
}