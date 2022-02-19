package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestName2(t *testing.T) {
	WebServer := &WebServer2{config: &Config{
		"小王是大帅笔",
	}}
	startTime := time.Now()
	Wg := sync.WaitGroup{}
	Wg.Add(10000)
	go WebServer.ReloadWork()
	fmt.Println("1", &WebServer.config.context)

	for i := 0; i < 1000; i++ {
		go func() {
			defer Wg.Done()
			for j := 0; j < 1000; j++ {
				//time.Sleep(1 * time.Nanosecond)
				WebServer.visit()
				if j == 5 {
					fmt.Println(&WebServer.config.context)
				}
				//fmt.Println(a)
			}
		}()
		//defer Wg.Done()
		//for i := 0; i < 10000; i++ {
		//	go WebServer2.visit()
		//}
	}
	Wg.Wait()
	finishTime := time.Now()
	fmt.Println("本次不加锁编程总耗时为:", finishTime.Sub(startTime))
	fmt.Println("2", &WebServer.config.context)
	//fmt.Println(WebServer2.config.context)
	//time.Sleep(2 * time.Second)

}
