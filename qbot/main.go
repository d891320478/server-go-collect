package main

import (
	"fmt"
	"math/big"
	"strings"
	"time"

	"crypto/rand"

	"github.com/ProtobufBot/go-pbbot"
	"github.com/ProtobufBot/go-pbbot/proto_gen/onebot"
	"github.com/d891320478/server-go-collect/qbot/functions"
	"github.com/gin-gonic/gin"
)

// GOOS=linux GOARCH=amd64 go build -o build/qbot -x main.go
func main() {

	pbbot.HandleConnect = func(bot *pbbot.Bot) {
	}

	var lastMaiDieTime int64 = 0
	var bignum_19 = big.NewInt(19)
	var bignum_1 = big.NewInt(1)

	pbbot.HandleGroupMessage = func(bot *pbbot.Bot, event *onebot.GroupMessageEvent) {
		rawMsg := event.RawMessage
		groupId := event.GroupId
		userId := event.UserId
		if rawMsg == "#help" {
			replyMsg := pbbot.NewMsg().Text(functions.FunctionList())
			_, _ = bot.SendGroupMessage(groupId, replyMsg, false)
		} else if strings.HasPrefix(rawMsg, "#命运示数器") {
			target := strings.TrimSpace(strings.Replace(rawMsg, "#命运示数器", "", 1))
			a, _ := rand.Int(rand.Reader, bignum_19)
			b, _ := rand.Int(rand.Reader, bignum_19)
			a = a.Add(a, bignum_1)
			b = b.Add(b, bignum_1)
			rlt := ""
			if b.Cmp(a) >= 0 {
				rlt = "检定成功"
			} else {
				rlt = "检定失败"
			}
			replyMsg := pbbot.NewMsg().At(userId).Text(fmt.Sprintf("\r\n检定目标：%s\r\n检定数值：%s\r\n检定结果：%s\r\n%s", target, a.String(), b.String(), rlt))
			_, _ = bot.SendGroupMessage(groupId, replyMsg, false)
		}
		if strings.Contains(rawMsg, "碟") && groupId != 151118379 {
			if time.Now().Unix()-lastMaiDieTime > 90 {
				lastMaiDieTime = time.Now().Unix()
				replyMsg := pbbot.NewMsg().At(userId).Text(" 买碟吗").Image("http://htdong-n4:9961/img/weidian.jpg")
				_, _ = bot.SendGroupMessage(groupId, replyMsg, false)
			}
		}
	}

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		if err := pbbot.UpgradeWebsocket(c.Writer, c.Request); err != nil {
			panic("创建机器人失败")
		}
	})

	if err := router.Run(":9999"); err != nil {
		panic(err)
	}
	select {}
}
