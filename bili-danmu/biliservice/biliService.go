package biliservice

import (
	"os"
	"strings"

	"github.com/Akegarasu/blivedm-go/client"
	"github.com/Akegarasu/blivedm-go/message"
)

func Register(channel chan string, roomId int64) (*client.Client, error) {
	c := client.NewClient(int(roomId))
	c.SetCookie(getCookieFromFile())
	// 弹幕事件
	c.OnDanmaku(func(danmaku *message.Danmaku) {
		// TODO
		channel <- danmaku.Content
	})
	// 醒目留言
	c.OnSuperChat(func(superChat *message.SuperChat) {
		// TODO
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
