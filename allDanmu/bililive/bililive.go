package bililive

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Akegarasu/blivedm-go/client"
	"github.com/Akegarasu/blivedm-go/message"
)

const danmuFilePath = "/data/biliDanMu%d/%d-%s-%s.log"

var c *client.Client

func AllDanMu(roomId int) {
	c := client.NewClient(roomId)
	c.SetCookie(getCookieFromFile())
	// 弹幕事件
	c.OnDanmaku(func(danmaku *message.Danmaku) {
		if danmaku.Type != message.EmoticonDanmaku {
			writeToFile(time.Unix(danmaku.Timestamp/1000, 0).Format("2006-01-02 15:04:05"), danmaku.Sender.Uname, danmaku.Content, roomId, danmaku.Sender.Uid)
		}
	})
	// 醒目留言
	c.OnSuperChat(func(superChat *message.SuperChat) {
		jsonStr, _ := json.Marshal(superChat)
		fmt.Println(string(jsonStr))
		writeToFile(time.Unix(int64(superChat.StartTime), 0).Format("2006-01-02 15:04:05")+"[SC]", superChat.UserInfo.Uname, superChat.Message, roomId, superChat.Uid)
	})
	err := c.Start()
	if err != nil {
		panic(err)
	}
	for {
		time.Sleep(5 * time.Minute)
	}
}

func writeToFile(tm, uname, content string, roomId, uid int) {
	now := time.Now()
	filePath := fmt.Sprintf(danmuFilePath, roomId, now.Year(), now.Format("01"), now.Format("02"))
	file, _ := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(fmt.Sprintf("[%s] %s[%d]: %s\n", tm, uname, uid, content))
	write.Flush()
	file.Close()
}

func getCookieFromFile() string {
	b, err := os.ReadFile("/root/conf/biliCookie.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}
