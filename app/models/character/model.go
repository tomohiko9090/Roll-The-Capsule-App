package character

import (
	"GachaAPI/app/models"
	_ "bytes"
	_ "fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// 3. キャラクター一覧取得
type Characters struct {
	Results []Results `json:"characters"`
}

type Results struct {
	UserCharacterID int    `json:"userID"`
	CharacterID     int    `json:"userName"`
	Name            string `json:"characterName"`
	Rarity          string `json:"rarity"`
}

func GetCharacters(token string) (Characters, error) {

	// トークンからユーザーidの取得
	var userid string
	err := models.DB.QueryRow("SELECT id FROM capsule.User WHERE token = ?", token).Scan(&userid)
	//if err != nil {
	//	fmt.Println(err)
	//}

	// ユーザーidからキャラidの一覧とuserCharacterIDを取得
	var (
		userCharacterID int
		characterid     int
		name            string
		rarity          string
	)

	result_list := []Results{}

	rows, err := models.DB.Query("SELECT userCharacterID, character_id FROM capsule.Possess WHERE user_id = ?", userid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&userCharacterID, &characterid)
		//　キャラidからCharacterテーブルのname, rarityを取得
		err = models.DB.QueryRow("SELECT name, rarity  FROM capsule.Character WHERE id = ?", characterid).
			Scan(&name, &rarity)
		if err != nil {
			log.Fatal(err)
		}

		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

		result := Results{
			UserCharacterID: userCharacterID,
			CharacterID:     characterid,
			Name:            name,
			Rarity:          rarity,
		}
		result_list = append(result_list, result)
	}
	results := Characters{result_list}
	return results, nil
}
