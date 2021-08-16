package gacha

import (
	"GachaAPI/app/models"
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Error     Characters
	character Character
	user      User
)

// ②キャラ数取得
func GetCharacterUnique() (int, error) {

	// characterテーブルからキャラ数取得
	var total int
	err := models.DB.QueryRow("SELECT COUNT(*) FROM capsule.Character").Scan(&total)
	if err != nil {
		return -1, err
	}
	return total, nil
}

// ④当たったキャラのインサート
func InsertCharacters(token string, ans_id []int) error {

	// tokenからidを引っ張ってくる
	err := models.DB.QueryRow("SELECT * FROM capsule.User WHERE token = ?", token).
		Scan(&user.ID, &user.Name, &user.Token)
	if err != nil {
		return err
	}

	// usercharacterIDの長さを取得
	var total int
	err = models.DB.QueryRow("SELECT COUNT(*) FROM capsule.Possess").Scan(&total)
	if err != nil {
		return err
	}
	usercharacterID := total + 1

	// 複数回当たったキャラをインサート
	for i := 0; i < len(ans_id); i++ {
		stmtInsert, err := models.DB.Prepare("INSERT INTO Possess(user_id, usercharacterID, character_id) VALUES(?, ?, ?)")
		if err != nil {
			return err
		}
		defer stmtInsert.Close()

		result1, err := stmtInsert.Exec(user.ID, usercharacterID+i, ans_id[i])

		if err != nil {
			return err
		}
		_, err = result1.LastInsertId()
		if err != nil {
			return err
		}
	}
	return nil // ここまで実行できていれば、nilを返す
}

// ⑤当たったキャラ情報を取得
func GetCharactersData(ans_id []int) (Characters, error) {

	resultlist := []GachaResults{}

	// 複数のキャラnameとrarityを取得
	for i := 0; i < len(ans_id); i++ {

		rows, err := models.DB.Query("SELECT * FROM capsule.Character WHERE id = ?", ans_id[i])
		if err != nil {
			return Error, err
		}
		defer rows.Close()

		rows.Next()
		rows.Scan(&character.CharacterID, &character.Name, &character.Rarity, &character.Attack, &character.Defence, &character.Recovery)

		result := GachaResults{
			UserID:        user.ID,
			UserName:      user.Name,
			CharacterName: character.Name,
			Rarity:        character.Rarity,
			Attack:        character.Attack,
			Defence:       character.Defence,
			Recovery:      character.Recovery,
		}
		resultlist = append(resultlist, result)
	}

	results := Characters{resultlist}
	return results, nil
}
