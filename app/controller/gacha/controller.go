package controller

import (
	"GachaAPI/app/models/character"
	user "GachaAPI/app/models/user"
	"GachaAPI/app/models/userCharacters"
	"database/sql"
	"math/rand"
	"time"
)

// GetCharacterLength なし -> キャラクターの種類の数
func GetCharacterLength() (int, error) {
	characterLength, err := character.GetCharacterLength()
	if err != nil {
		return 0, err
	}
	return characterLength, nil
}

// DrowCharacter キャラクターの種類の数, トークン, ガチャ回数 -> 当たったキャラクターIDの配列, キャラクターテーブルのmap
func DrowCharacter(characterLength int, token string, numberOfTimes int) ([]int, map[int]character.Character, error) {

	var (
		resultCharacterIDs []int
		characterRows      *sql.Rows
	)
	status := make(map[int]character.Character)

	// ユーザー情報の取得
	user, err := user.SelectUser(token)
	if err != nil {
		return resultCharacterIDs, status, err
	}

	// ガチャで得られたキャラクターIDの配列
	resultCharacterIDs = turnCharacterID(characterLength, numberOfTimes)

	for times := 0; times < numberOfTimes; times++ {

		// 当たったキャラクターIDをDBにインサート
		err := userCharacters.InsertUserCharacter(user.ID, resultCharacterIDs[times])
		if err != nil {
			return resultCharacterIDs, status, err
		}

	}

	// キャラクターテーブルをまるごと取得
	characterRows, err = character.SelectCharacterTable()
	if err != nil {
		return resultCharacterIDs, status, err
	}
	defer characterRows.Close()

	// キャラクターテーブルをmapにする(キャラクターidをキー, キャラクターの詳細情報をバリューに)
	for characterRows.Next() {
		var character character.Character
		err = characterRows.Scan(
			&character.CharacterID,
			&character.CharacterName,
			&character.Rarity,
			&character.Attack,
			&character.Defence,
			&character.Recovery)
		if err != nil {
			return resultCharacterIDs, status, err
		}
		status[character.CharacterID] = character
	}
	return resultCharacterIDs, status, err
}

// turnCharacterID キャラクターの種類の数, ガチャ回数 ->　当たったキャラクターIDの配列
func turnCharacterID(characterLength int, times int) []int {
	rangePlus := 0
	for i := 1; i <= characterLength; i++ {
		rangePlus += i // キャラid全て足しあわせる
	}
	characterRange := make(map[int]int, 0) // goらしい書き方で定義
	count := 0
	max := 0
	for i := 1; i <= characterLength; i++ {
		b := 100 * (characterLength + 1 - i) / rangePlus // 当たり範囲を生成
		max = count + b
		// 範囲をmapで定義する
		for j := count; j <= max; j++ {
			characterRange[j] = i
		}
		count += b
	}

	t := time.Now().UnixNano() // 現在時刻でランダム性を担保
	rand.Seed(t)

	var resultCharacterIDs []int

	for i := 0; i < times; i++ {
		n := rand.Intn(count + 1) // 乱数を発生させる
		resultCharacterIDs = append(resultCharacterIDs, characterRange[n])
	}
	return resultCharacterIDs
}
