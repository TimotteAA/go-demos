package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// 简单的定时任务
func main() {
	// 重复1s的一个tick
	t1 := time.Tick(time.Second * 1)
	t3 := time.Tick(time.Second * 3)
	t5 := time.Tick(time.Second * 5)

	//	轮询
	for {
		select {
		case <-t1:
			println("1 sec")
		case <-t3:
			println("3 sec")
		case <-t5:
			println("5 sec")
		}
	}
}

/** select配合channel */
func test5() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 2)

	ch3 := make(chan string, 1)
	ch4 := make(chan string, 2)

	ch1 <- "select channel1"
	ch2 <- "select channel2"

	// 两个channel都满足，会随机选择一个执行
	// 无缓冲channel，得立刻读
	go func() {
		for {
			select {
			case msg1, ok := <-ch1:
				if !ok {
					println("channel1被关闭")
				}
				println("channel1收到数据 ", msg1)
			case msg2, ok := <-ch2:
				if !ok {
					println("channel2被关闭")
				}
				println("channel2收到数据 ", msg2)
				//default:
				//	// 一直阻塞，进入此处
				//	println("select default")
				//	break
				//}
			case ch3 <- "select 3":
				fmt.Println("往channel3里写了")
			case ch4 <- "select 4":
				fmt.Println("往channel4里写了")
			}
		}

	}()

	fmt.Println(<-ch3)
	fmt.Println(<-ch4)

	time.Sleep(time.Second * 1)
}

func receive(ch chan string) {
	for {
		msg := <-ch
		println("读取到的消息 ", msg)
	}
}

/*
* 有缓冲channel * /

	func test4() {
		ch := make(chan string, 5)

		// 也许读的go程应该在前面
		go receive(ch)

		for i := 0; i < 10; i++ {
			// 写到第6个阻塞
			ch <- fmt.Sprintf("send %d", i)
			fmt.Println("i ", i)
		}

		fmt.Println("写完了")

		time.Sleep(time.Second * 1)
	}

/*** 无缓冲channel
*/
func test3() {
	ch := make(chan string)

	//go func() {
	//	msg := <-ch
	//	log.Println("读取到的数据 ", msg)
	//}()

	// 写完立刻阻塞了（缓冲为0）
	// 无缓冲channel是同步的，此处直接写入，但是没有人和他同步
	ch <- "send1"

	go func() {
		msg := <-ch
		log.Println("读取到的数据 ", msg)
	}()

	//msg := <-ch
	//log.Println("读取到的数据", msg)
}

/*****创建协程 * */
func test1() {
	// 等待所有协程执行完毕
	var wg sync.WaitGroup

	// 协程是由线程管控的
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("携程")
		time.Sleep(time.Second * 3)
	}()

	wg.Wait()
}

/*  多协程用锁避免脏读、脏写 */
var count = 5
var wg sync.WaitGroup
var lock sync.Mutex

func buy() {
	// 进来先加锁
	lock.Lock()
	defer wg.Done()
	// 退出前把锁放了
	defer lock.Unlock()

	// 模拟业务逻辑
	time.Sleep(time.Millisecond * 2)

	if count > 0 {
		log.Println("抢购成功")
	} else {
		log.Println("抢购失败")
		return
	}
	count -= 1
}

func test2() {
	//	数据加锁：防止多个协程同时读写数据时，造成数据不一致
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go buy()
	}

	wg.Wait()

	// 打印最终结果
	log.Println("库存数量 ", count)
}
