package domain

type DanMuVO struct {
	Content string
	Name    string
	Sc      bool
	Uid     int
	Avatar  string
}

type TenapiResult[T any] struct {
	Code int `json:"code"`
	Data T   `json:"data"`
}

type BiliUserDTO struct {
	Uid    string `json:"uid"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}
