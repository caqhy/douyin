package main

import (
	"fmt"
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

func testPrint() {
	fmt.Println(" 滴答 1 次")
}

func main() {
	t := NewMyTick(2, testPrint)
	t.Start()
}
