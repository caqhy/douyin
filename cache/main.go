package main

import (
	"fmt"
	_ "github.com/RaymondCode/simple-demo/dal"
	"github.com/RaymondCode/simple-demo/dal/db"
	"strconv"
	"strings"
	"time"
)

type Fn func()

type MyTicker struct {
	MyTick *time.Ticker
	Runner Fn
}

func NewMyTick(interval int, f Fn) *MyTicker {
	return &MyTicker{
		MyTick: time.NewTicker(time.Duration(interval) * time.Second),
		Runner: f,
	}
}

func (t *MyTicker) Start() {
	for {
		select {
		case <-t.MyTick.C:
			t.Runner()
		}
	}
}

func FavoriteToDb() {
	videos, _ := db.Redis.Scan(0, "douyin:favorite:*user1", 100).Val()
	for _, video := range videos {
		start := strings.Index(video, "video")
		end := strings.Index(video, ":user")
		video = video[start+5 : end]
		videoId, _ := strconv.ParseInt(video, 10, 64)
		fmt.Println(videoId)
	}

}

func main() {
	t := NewMyTick(2, FavoriteToDb)
	t.Start()
}
