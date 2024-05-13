package bililive

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
)

var total int
var mp map[int]int
var limit map[int]int
var channel chan TouPiao = make(chan TouPiao, 10)
var lockMp = new(sync.RWMutex)

func StartBiliHttp() {
	http.HandleFunc("/startGetDanMu", startGetDanMu)
	http.HandleFunc("/getCountRlt", getCountRlt)
	Register(channel)
	go count()
	http.ListenAndServe(":9777", nil)
}

func startGetDanMu(w http.ResponseWriter, req *http.Request) {
	lockMp.Lock()
	defer lockMp.Unlock()
	total, _ = strconv.Atoi(req.URL.Query().Get("total"))
	mp = make(map[int]int)
	limit = make(map[int]int)
	for i := 0; i < total; i++ {
		mp[i] = 0
	}
}

func count() {
	for {
		val := <-channel
		if val.Val <= total && val.Val > 0 {
			lockMp.Lock()
			limit[val.Uid]++
			if limit[val.Uid] <= 3 {
				mp[val.Val]++
			}
			lockMp.Unlock()
		}
	}
}

func getCountRlt(w http.ResponseWriter, req *http.Request) {
	lockMp.Lock()
	defer lockMp.Unlock()
	jsonStr, _ := json.Marshal(mp)
	w.Write(jsonStr)
}
