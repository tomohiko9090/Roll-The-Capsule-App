package handler

// 1.1 作成
type PostToken struct {
	Token string `json:"token"`
}

// 1.2 取得
type GetName struct {
	Name string `json:"name"`
}
