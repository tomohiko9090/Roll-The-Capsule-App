package controller

import (
	"math/rand"
	"time"
)

// ユーザー作成 -> ランダムでtokenを発行
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
func RandomString(len int) string {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}