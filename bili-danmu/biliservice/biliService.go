package biliservice

import (
	"os"
	"strings"

	"github.com/Akegarasu/blivedm-go/client"
	"github.com/Akegarasu/blivedm-go/message"
	"github.com/d891320478/server-go-collect/bili-danmu/domain"
)

func Register(channel chan domain.DanMuVO, roomId int64) (*client.Client, error) {
	c := client.NewClient(int(roomId))
	c.SetCookie(getCookieFromFile())
	// 弹幕事件
	c.OnDanmaku(func(danmaku *message.Danmaku) {
		// TODO
		channel <- domain.DanMuVO{
			Content: danmaku.Content,
			Name:    danmaku.Sender.Uname,
			Sc:      false,
			Uid:     danmaku.Sender.Uid,
		}
	})
	// 醒目留言
	c.OnSuperChat(func(superChat *message.SuperChat) {
		channel <- domain.DanMuVO{
			Content: superChat.Message,
			Name:    superChat.UserInfo.Uname,
			Sc:      true,
			Uid:     superChat.Uid,
			Avatar:  superChat.UserInfo.Face,
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

func getCookieFromFile() string {
	b, err := os.ReadFile("/root/conf/biliCookie.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}
