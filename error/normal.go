package main

import (
	"log"
	"time"
)

func Normal() {
	log.Printf("Calling test\r\n")
	test()
	log.Printf("Test completed\r\n")
}	

func test() {
	defer func ()  {
		if err := recover(); err != nil {
			log.Println("发生了错误 ", err)
		}
	}()
	badWork()
	log.Println("能看我吗？")
}

func badWork() {
	time.Sleep(time.Second * 3)
	panic("我死了")
}