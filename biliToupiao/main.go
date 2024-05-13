package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"github.com/d891320478/server-go-collect/util/file"
)

func Throwable() {
	err := recover()
	if err == nil {
		return
	}
	fmt.Println(err)
	fmt.Println(string(debug.Stack()))
}

// CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-X github.com/d891320478/server-go-collect/biliToupiao/main.Host=$ALIYUN_HOST" -o build/usr.exe -x src/Main.go
func main() {
	defer Throwable()
	biliToupiao()
}

const song_list_file = "list.txt"

var Host string

func biliToupiao() {
	if !file.Exists(song_list_file) {
		file, _ := os.Create(song_list_file)
		defer file.Close()
		w := bufio.NewWriter(file)
		w.WriteString("")
		w.Flush()
		file.Close()
	}
	stdinReader := bufio.NewReader(os.Stdin)
	fmt.Println("歌单放到list.txt，保存，然后按回车。。。。。。")
	stdinReader.ReadString('\n')
	// 读取歌单内容，生成编号，初始化票数
	mp := make(map[int]int)

	f, _ := os.Open("list.txt")
	defer f.Close()
	a, _ := io.ReadAll(f)
	a = append(a, 13)
	var list []string

	tmp := make([]byte, 0)
	for _, v := range a {
		if v == byte(10) || v == byte(13) {
			ss := string(tmp)
			fmt.Println(tmp)
			fmt.Printf("ss = %s\n", ss)
			if len(strings.TrimSpace(ss)) > 0 {
				list = append(list, strings.TrimSpace(ss))
			}
			tmp = make([]byte, 0)
		} else {
			tmp = append(tmp, v)
		}
	}
	total := len(list)

	for i := 0; i < total; i++ {
		mp[i] = 0
	}
	fmt.Println(total)
	fmt.Println(list)
	fmt.Println(mp)
	// 回写文件
	writeToListFile(mp, list, total)
	// 发请求start
	http.Get("http://" + Host + ":9961/htdong/liveVote/startVote")
	time.Sleep(3 * time.Second)
	resp, err := http.Get(fmt.Sprintf("http://"+Host+":9961/startLive/startGetDanMu?total=%d", total))
	for {
		fmt.Println("startGetDanMu")
		if err != nil {
			fmt.Printf("startGetDanMu err. %v\n", err)
		} else {
			if resp.StatusCode == 200 {
				break
			}
			fmt.Printf("startGetDanMu code = %d\n", resp.StatusCode)
		}
		time.Sleep(1 * time.Second)
		resp, err = http.Get(fmt.Sprintf("http://"+Host+":9961/startLive/startGetDanMu?total=%d", total))
	}
	mp = make(map[int]int)
	for {
		resp, err := http.Get("http://" + Host + ":9961/startLive/getCountRlt")
		if err != nil {
			fmt.Printf("getCountRlt err. %v\n", err)
		} else if resp.StatusCode == 200 {
			jsonStr, _ := io.ReadAll(resp.Body)
			json.Unmarshal(jsonStr, &mp)
			fmt.Println(mp)
			writeToListFile(mp, list, total)
		} else {
			fmt.Printf("getCountRlt code = %d\n", resp.StatusCode)
		}
		time.Sleep(3 * time.Second)
	}
}

func writeToListFile(mp map[int]int, list []string, total int) {
	file, _ := os.OpenFile(song_list_file, os.O_WRONLY, os.ModeAppend)
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString("规则：征集时请把想听的歌打在弹幕 选曲没任何限制0w0\r\n")
	write.WriteString("投票时每个人最多三票\r\n")
	write.WriteString("从会唱的里面选两首的票最高的今天唱/吹\r\n")
	write.WriteString("不会的里最高的下周唱（不含卡祖笛曲）\r\n")
	write.WriteString("想听卡祖笛版也可以，在歌名前面加上卡祖笛三字。下周翻唱\r\n")
	for i := 0; i < total; i++ {
		val := strconv.Itoa(i+1) + ". " + strings.TrimSpace(list[i]) + "   " + strconv.Itoa(mp[i]) + " 票\r\n"
		write.WriteString(val)
	}
	write.Flush()
}
