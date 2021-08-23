package possess

import "GachaAPI/app/models"

// SelectPossessList ユーザーID -> possessテーブル構造体の配列
func SelectPossessList(userId int) ([]Possess, error) {
	var possessList []Possess
	rows, err := models.DB.Query("SELECT * FROM Possess WHERE user_id = ?", userId)
	if err != nil {
		return possessList, err
	}
	defer rows.Close()

	var possess Possess
	for rows.Next() {
		err = rows.Scan(&possess.UserID, &possess.UserCharacterID, &possess.CharacterID)
		if err != nil {
			return possessList, err
		}
		possessList = append(possessList, possess)
	}
	return possessList, nil
}

// GetPossessLength なし -> 所持テーブルのレコード数
func GetPossessLength() (int, error) {
	var possessLength int
	err := models.DB.QueryRow("SELECT COUNT(*) FROM Possess").Scan(&possessLength)
	if err != nil {
		return possessLength, err
	}
	return possessLength, err
}

// InsertPossess ユーザーID, 所持テーブル数, ガチャ途中実行数, 当たったキャラクターID -> なし
func InsertPossess(userId int, possessLength int, i int, resultCharacterID int) error {
	stmtInsert, err := models.DB.Prepare("INSERT INTO Possess(user_id, usercharacterID, character_id) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmtInsert.Close()

	result1, err := stmtInsert.Exec(userId, possessLength+i, resultCharacterID)

	if err != nil {
		return err
	}
	_, err = result1.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}
