package main

import (
	"log"
	"time"
)

func test() {
	// 1s的ticker
	ticker1s := time.NewTicker(time.Second * 1)
	// 3s的ticker
	ticker3s := time.NewTicker(time.Second * 3)
	boom := time.After(time.Second * 15)
	stop := false;
	for {
		select {
		case <- ticker1s.C:
			log.Println("...")
		case <- ticker3s.C: {
			log.Println("...tick tick")
		}
		// 15s后爆炸 
		// 写在下面卵子用没有，因为每当select监听到后，都相当于重新创建了一个 time.After的定时器
		// case <- := time.After(time.Second * 15)
		case <- boom:
			log.Println("爆炸啦")
			stop = true;
			// break的是select语句
			break;
		}
		if stop {
			break;
		}
	}
}