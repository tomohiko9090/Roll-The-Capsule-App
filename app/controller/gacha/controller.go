package controller

import (
	"GachaAPI/app/models/character"
	"GachaAPI/app/models/possess"
	user "GachaAPI/app/models/user"
	"fmt"
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

// DrowCharacter キャラクターの種類の数, トークン, ガチャ回数 -> 当たったキャラクターIDの配列
func DrowCharacter(characterLength int, token string, times int) ([]int, error) {

	var resultCharacterIDs []int

	// ユーザー情報の取得
	user, err := user.SelectUser(token)
	if err != nil {
		return resultCharacterIDs, err
	}

	// ガチャで得られたキャラクターIDの配列
	resultCharacterIDs = turnCharacterID(characterLength, times)

	// possessの長さ
	possessLength, err := possess.GetPossessLength()
	if err != nil {
		return resultCharacterIDs, err
	}

	for i := 0; i < times; i++ {

		// 当たったキャラクターIDをDBにインサート
		err := possess.InsertPossess(user.ID, possessLength, i, resultCharacterIDs[i])
		if err != nil {
			return resultCharacterIDs, err
		}

	}
	return resultCharacterIDs, err
}

// GetCharacter キャラクターID -> characterテーブル構造体
func GetCharacter(characteId int) (character.Character, error) {
	character, err := character.SelectCharacter(characteId)
	if err != nil {
		return character, err
	}
	return character, nil
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

	for i := 0; i <= times; i++ {
		n := rand.Intn(count + 1) // 乱数を発生させる
		resultCharacterIDs = append(resultCharacterIDs, characterRange[n])
	}
	fmt.Println(resultCharacterIDs)
	return resultCharacterIDs
}
