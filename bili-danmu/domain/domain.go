package domain

type DanMuVO struct {
	Content string `json:"content"`
	Name    string `json:"name"`
	Sc      bool   `json:"sc"`
	Uid     int    `json:"uid"`
	Avatar  string `json:"avatar"`
	Empty   bool   `json:"empty"`
	Type    int    `json:"type"`
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
