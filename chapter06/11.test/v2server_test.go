package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWebServerV2(t *testing.T) {
	v2 := &WebServerV2{
		config: &Config{Content: "a"},
	}
	go v2.ReloadWorker()
	start := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				v2.Visit() //在这里打印v2.config.Content 为time.Now().UnixNano()，但是在程序最后一行输出却还是a,然后通过输出内存地址，发现内存地址发生了改变
				if i == 9999 {
					fmt.Println("当i==99999时:", v2.config.Content)
				}
			}
		}()
	}
	wg.Wait()
	finish := time.Now()
	fmt.Println("总时间：", finish.Sub(start))
	fmt.Println("运行结束之后v2版WebServer中配置文件的内容为:", v2.config.Content)
}
