package bililive

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var total int
var mp map[int]int = make(map[int]int)
var channel chan int = make(chan int)
var lockMp = new(sync.RWMutex)
var stopTime int64

func StartBiliHttp() {
	stopTime = time.Now().Unix() + 100
	http.HandleFunc("/startGetDanMu", startGetDanMu)
	http.HandleFunc("/getCountRlt", getCountRlt)
	go http.ListenAndServe(":9777", nil)
	for {
		time.Sleep(60 * time.Second)
		// 没有请求了，停止
		now := time.Now().Unix()
		fmt.Printf("%d %d\n", now, stopTime)
		if now > stopTime {
			os.Exit(0)
		}
	}
}

func startGetDanMu(w http.ResponseWriter, req *http.Request) {
	total, _ = strconv.Atoi(req.URL.Query().Get("total"))
	Register(channel, total)
	for i := 0; i < total; i++ {
		mp[i] = 0
	}
	go count()
}

func count() {
	for {
		val := <-channel
		lockMp.Lock()
		mp[val-1]++
		lockMp.Unlock()
	}
}

func getCountRlt(w http.ResponseWriter, req *http.Request) {
	stopTime = time.Now().Unix() + 50
	lockMp.Lock()
	defer lockMp.Unlock()
	jsonStr, _ := json.Marshal(mp)
	w.Write(jsonStr)
}
