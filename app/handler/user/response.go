package handler

// UserCreateResponse 1.1 作成
type UserCreateResponse struct {
	Token string `json:"token"`
}

// UserGetResponse 1.2 取得
type UserGetResponse struct {
	Name string `json:"name"`
}
