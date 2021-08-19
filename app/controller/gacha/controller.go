package controller

import (
	"GachaAPI/app/models/gacha"
	"math/rand"
	"time"
)

var Error gacha.Characters

func DrowCharacters(token string, drows int) (gacha.Characters, error) {
	// キャラクターユニーク数の取得
	total, err := gacha.GetCharacterUnique()
	if err != nil {
		return Error, err
	}

	// ガチャを回す
	ans_id := getCharacters(drows, total)

	// 当たったキャラクターをインサートする
	err = gacha.InsertCharacters(token, ans_id)
	if err != nil {
		return Error, err
	}
	// 当たったキャラ情報の取得
	results, err := gacha.GetCharactersData(ans_id)
	return results, nil
}

func getCharacters(times int, total int) []int {
	// レア度が低いものの方が当たりやすくする
	rangePlus := 0
	for i := 1; i <= total; i++ {
		rangePlus += i // キャラid全て足しあわせる
	}

	m := make(map[int]int, 0) // goらしい書き方で定義
	count := 0
	max := 0
	for i := 1; i <= total; i++ {
		b := 100 * (total + 1 - i) / rangePlus // 当たり範囲を生成
		max = count + b
		for j := count; j <= max; j++ { // 範囲をmapで定義する
			m[j] = i
		}
		count += b
	}
	t := time.Now().UnixNano() // 現在時刻でランダム性を担保
	rand.Seed(t)
	var ans []int
	for i := 1; i <= times; i++ {
		n := rand.Intn(count + 1) // 乱数を発生させる
		ans = append(ans, m[n])
	}
	return ans
}
