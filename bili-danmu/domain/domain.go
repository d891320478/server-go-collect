package domain

type DanMuVO struct {
	Empty       bool    `json:"empty"`
	Guard       bool    `json:"guard"`
	Gift        bool    `json:"gift"`
	GiftUrl     string  `json:"giftUrl"`
	Sc          bool    `json:"sc"`
	Content     string  `json:"content"`
	Name        string  `json:"name"`
	Uid         int     `json:"uid"`
	Avatar      string  `json:"avatar"`
	Type        int     `json:"type"`
	EmoticonUrl string  `json:"emoticonUrl"`
	GiftNum     int     `json:"giftNum"`
	GiftType    string  `json:"giftType"`
	Price       float64 `json:"price"`
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

type BiliGiftDetailDTO struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	CoinType string `json:"coin_type"` // gold, silver
	ImgBasic string `json:"img_basic"`
}
