package bililive

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/Akegarasu/blivedm-go/client"
	"github.com/Akegarasu/blivedm-go/message"
)

const roomId = 222272

var c *client.Client

func Register(channel chan int, total int) {
	limit := make(map[int]int)
	lock := new(sync.RWMutex)
	c = client.NewClient(roomId)
	c.SetCookie(getCookieFromFile())
	//弹幕事件
	c.OnDanmaku(func(danmaku *message.Danmaku) {
		fmt.Printf("[弹幕] %s[%d]：%s\n", danmaku.Sender.Uname, danmaku.Sender.Uid, danmaku.Content)
		if danmaku.Type != message.EmoticonDanmaku {
			val, err := strconv.Atoi(danmaku.Content)
			if err == nil {
				if val > 0 && val <= total {
					lock.Lock()
					limit[danmaku.Sender.Uid]++
					if limit[danmaku.Sender.Uid] <= 3 {
						channel <- val
					}
					lock.Unlock()
				}
			}
		}
	})
	err := c.Start()
	if err != nil {
		fmt.Println("Register")
		fmt.Println(err)
		panic(err)
	}
}

func CloseClient() {
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
