package gacha

import (
	"GachaAPI/app/models/gacha"
	"math/rand"
	"time"
)

func DrowCharacters(token string, drows int) gacha.Characters {
	// キャラクターユニーク数の取得
	total := gacha.GetCharacterUnique()

	// ガチャを回す
	ans_id := getCharacters(drows, total)

	// 当たったキャラクターをインサートする
	gacha.InsertCharacters(token, ans_id)

	// 当たったキャラ情報の取得
	results := gacha.GetCharactersData(ans_id, token)
	return results
}

func getCharacters(times int, total int) []int {
	// レア度が低いものの方が当たりやすくする
	range_plus := 0
	for i := 1; i <= total; i++ {
		range_plus += i // キャラid全て足しあわせる
	}
	m := map[int]int{}
	count := 0
	max := 0
	for i := 1; i <= total; i++ {
		b := 100 * (total + 1 - i) / range_plus // 当たり範囲を生成
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
