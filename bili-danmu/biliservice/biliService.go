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
	"github.com/d891320478/server-go-collect/bili-danmu/constants"
	"github.com/d891320478/server-go-collect/bili-danmu/domain"
	"github.com/d891320478/server-go-collect/bili-danmu/redisService"
)

const redis_bili_avator_format = "bili-danmu-bili-avator-%d"
const bili_avator_get_url = "https://tenapi.cn/bilibili/?uid=%d"
const redis_avator_expire = 4

func Register(channel chan domain.DanMuVO, roomId int64) (*client.Client, error) {
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
		}
	})
	// TODO 礼物
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

func getCookieFromFile() string {
	b, err := os.ReadFile("/root/conf/biliCookie.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}
