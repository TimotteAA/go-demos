package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	concurrent()
}

func concurrent() {
	for i := 0; i <= 20; i++ {
		go safeWork()
	}

	time.Sleep(time.Second * 1)
	log.Println("程序退出")
}

func safeWork() {
	defer func ()  {
		log.Println("一个go程挂逼了")
		if err := recover(); err != nil {
			log.Println("呜呜呜呜")
		}	
	}()
	badWork2()
}

func badWork2() {
	rand := rand.Float32()
	if rand >= 0.8 {
		panic("我挂了")
	} else {
		log.Println("执行任务成功")
	}
}