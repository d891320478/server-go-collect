package biliservice

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Akegarasu/blivedm-go/client"
	"github.com/Akegarasu/blivedm-go/message"
	baselog "github.com/d891320478/server-go-collect/base-log"
	"github.com/d891320478/server-go-collect/bili-danmu/constants"
	"github.com/d891320478/server-go-collect/bili-danmu/domain"
	"github.com/d891320478/server-go-collect/bili-danmu/redisService"
)

const redis_bili_avator_format = "bili-danmu-bili-avator-%d"
const bili_avator_get_url = "https://tenapi.cn/bilibili/?uid=%d"
const redis_avator_expire = 4
const bili_gift_list_url = "https://api.live.bilibili.com/gift/v3/live/gift_config"

func Register(channel chan domain.DanMuVO, roomId int64) (*client.Client, error) {
	giftmap := BiliGiftList()
	c := client.NewClient(int(roomId))
	c.SetCookie(getCookieFromFile())
	// 弹幕事件
	c.OnDanmaku(func(danmaku *message.Danmaku) {
		dm := domain.DanMuVO{
			Content: danmaku.Content,
			Name:    danmaku.Sender.Uname,
			Sc:      false,
			Uid:     danmaku.Sender.Uid,
			Empty:   false,
			Type:    danmaku.Type,
			Gift:    false,
			Guard:   false,
		}
		if danmaku.Type == message.EmoticonDanmaku {
			dm.EmoticonUrl = danmaku.Emoticon.Url
		}
		channel <- dm
	})
	// 醒目留言
	c.OnSuperChat(func(superChat *message.SuperChat) {
		channel <- domain.DanMuVO{
			Content: superChat.Message,
			Name:    superChat.UserInfo.Uname,
			Sc:      true,
			Uid:     superChat.Uid,
			Avatar:  superChat.UserInfo.Face,
			Empty:   false,
			Gift:    false,
			Guard:   false,
		}
	})
	// 礼物
	c.OnGift(func(gift *message.Gift) {
		str, _ := json.Marshal(gift)
		baselog.InfoLog().Info(string(str))
		channel <- domain.DanMuVO{
			Content:  gift.GiftName,
			Name:     gift.Uname,
			Sc:       false,
			Uid:      gift.Uid,
			Avatar:   gift.Face,
			Empty:    false,
			Gift:     true,
			GiftNum:  gift.Num,
			GiftType: gift.CoinType,
			Guard:    false,
			Price:    float64(gift.Price*gift.Num) / 1000.0,
			GiftUrl:  giftmap[int64(gift.GiftId)],
		}
	})
	// TODO 上舰事件
	c.OnGuardBuy(func(guardBuy *message.GuardBuy) {
		str, _ := json.Marshal(guardBuy)
		baselog.InfoLog().Info(string(str))
		channel <- domain.DanMuVO{
			Name:    guardBuy.Username,
			Sc:      false,
			Uid:     guardBuy.Uid,
			Empty:   false,
			Gift:    false,
			GiftNum: guardBuy.Num,
			Guard:   true,
			Content: guard(guardBuy.GuardLevel),
			Price:   float64(guardBuy.Price*guardBuy.Num) / 1000.0,
		}
	})
	return c, c.Start()
}

func CloseClient(c *client.Client) {
	if c != nil {
		c.Stop()
	}
	c = nil
}

func GetAvatar(uid int) string {
	redisKey := fmt.Sprintf(redis_bili_avator_format, uid)
	url, _ := redisService.Get(redisKey)
	if len(url) == 0 {
		resp, err := http.Get(fmt.Sprintf(bili_avator_get_url, uid))
		if err == nil && resp.StatusCode == http.StatusOK {
			var rlt domain.TenapiResult[domain.BiliUserDTO]
			jsonStr, _ := io.ReadAll(resp.Body)
			json.Unmarshal(jsonStr, &rlt)
			if rlt.Code == constants.TENAPI_RESULT_SUCCESS_CODE {
				url = rlt.Data.Avatar
				if len(url) > 0 {
					redisService.Set(redisKey, url, redis_avator_expire, time.Hour)
				}
			}
		}
	}
	return url
}

func BiliGiftList() map[int64]string {
	mp := make(map[int64]string)
	resp, err := http.Get(bili_gift_list_url)
	if err == nil && resp.StatusCode == http.StatusOK {
		var rlt domain.TenapiResult[[]domain.BiliGiftDetailDTO]
		jsonStr, _ := io.ReadAll(resp.Body)
		json.Unmarshal(jsonStr, &rlt)
		if rlt.Code == constants.TENAPI_RESULT_SUCCESS_CODE {
			for _, iter := range rlt.Data {
				mp[iter.Id] = iter.ImgBasic
			}
		}
	}
	return mp
}

func getCookieFromFile() string {
	b, err := os.ReadFile("/root/conf/biliCookie.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}

func guard(level int) string {
	if level == 1 {
		return "总督"
	} else if level == 2 {
		return "提督"
	} else {
		return "舰长"
	}
}
