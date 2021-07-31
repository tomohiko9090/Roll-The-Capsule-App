package controller

import (
	"math/rand"
	"time"
)

func Gacha(total int, times int) []int {
	// レア度が低いものの方が当たりやすくする
	range_plus := 0
	for i := 1; i <= total; i++{
		range_plus += i // キャラid全て足しあわせる
	}
	m := map[int]int{}
	count := 0
	max := 0
	for i := 1; i <= total; i++{
		b := 100*(total+1-i)/range_plus // 当たり範囲を生成
		max = count + b
		for j := count; j <= max; j++{ // 範囲をmapで定義する
			m[j] = i
		}
		count += b
	}
	t := time.Now().UnixNano() // 現在時刻でランダム性を担保
	rand.Seed(t)
	var ans []int
	for i:=1; i<=times; i++{
		n := rand.Intn(count+1) // 乱数を発生させる
		ans = append(ans, m[n])
	}
	return ans
}