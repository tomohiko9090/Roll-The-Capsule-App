package gacha

import (
	"GachaAPI/app/models"
	_ "database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// ②キャラ数取得
func GetCharacterUnique() int {

	// characterテーブルからキャラ数取得
	var total int
	err := models.DB.QueryRow("SELECT COUNT(*) FROM capsule.Character").Scan(&total)
	if err != nil {
		log.Fatal(err)
	}
	return total
}

// ④当たったキャラのインサート
func InsertCharacters(token string, ans_id []int) {

	// tokenからidを引っ張ってくる
	var userid string
	if err := models.DB.QueryRow("SELECT id FROM capsule.User WHERE token = ?", token).Scan(&userid); err != nil {
		fmt.Println(err)
	}

	// usercharacterIDの長さを取得
	var total int
	err := models.DB.QueryRow("SELECT COUNT(*) FROM capsule.Possess").Scan(&total)
	if err != nil {
		panic(err.Error())
	}
	usercharacterID := total + 1

	// 複数回当たったキャラをインサート
	for i := 0; i < len(ans_id); i++ {
		stmtInsert, err := models.DB.Prepare("INSERT INTO Possess(user_id, usercharacterID, character_id) VALUES(?, ?, ?)")
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

// ⑤当たったキャラ情報を取得
type Characters struct {
	Results []Results `json:"results"`
}

type Results struct {
	UserID        int    `json:"userID"`
	UserName      string `json:"userName"`
	CharacterName string `json:"characterName"`
	Rarity        string `json:"rarity"`
}

func GetCharactersData(ans_id []int, token string) Characters {

	var (
		userid        int
		username      string
		charactername string
		rarity        string
	)

	result_list := []Results{}

	// 複数のキャラnameとrarityを取得
	for i := 0; i < len(ans_id); i++ {
		rows, _ := models.DB.Query("SELECT name,rarity FROM capsule.Character WHERE id = ?", ans_id[i])
		defer rows.Close()

		rows.Next()
		rows.Scan(&charactername, &rarity)

		rows2, _ := models.DB.Query("SELECT id,name FROM capsule.User WHERE token = ?", token)
		defer rows2.Close()

		rows2.Next()
		rows2.Scan(&userid, &username)

		result := Results{
			UserID:        userid,
			UserName:      username,
			CharacterName: charactername,
			Rarity:        rarity,
		}

		result_list = append(result_list, result)
	}

	results := Characters{result_list}
	return results
}
