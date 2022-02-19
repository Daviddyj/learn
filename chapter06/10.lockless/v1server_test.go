package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	WebServer := &WebServer{config: Config{
		"小王是大帅笔",
	}}
	startTime := time.Now()
	Wg := sync.WaitGroup{}
	Wg.Add(10000)
	go WebServer.ReloadWork()

	for i := 0; i < 10000; i++ {
		go func() {
			defer Wg.Done()
			for i := 0; i < 10000; i++ {
				WebServer.visit()
			}
		}()
		//defer Wg.Done()
		//for i := 0; i < 10000; i++ {
		//	go WebServer2.visit()
		//}
	}
	Wg.Wait()
	finishTime := time.Now()
	fmt.Println("本次加锁编程总耗时为:", finishTime.Sub(startTime))
	fmt.Println(WebServer.config.context)
	//fmt.Println(WebServer2.config.context)
	//time.Sleep(2 * time.Second)

}
