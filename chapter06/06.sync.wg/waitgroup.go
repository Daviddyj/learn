package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	//numCount := 10
	var Runners []Runner
	wg := sync.WaitGroup{}
	wg.Add(10)
	startPoint := sync.WaitGroup{}
	startPoint.Add(1)
	for i := 0; i < 10; i++ {
		Runners = append(Runners, Runner{
			fmt.Sprintf("%d", i),
		})
	}
	for _, item := range Runners {
		go item.Run(&wg, &startPoint)
	}
	fmt.Println("预备，跑")
	startPoint.Done()
	wg.Wait()

}

type Runner struct {
	name string
}

func (r Runner) Run(wg, startPoint *sync.WaitGroup) {
	defer wg.Done()
	startPoint.Wait()
	startTime := time.Now()
	fmt.Println(r.name, "开始跑@", startTime)
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Uint64()%10) * time.Second)
	finishTime := time.Now()
	fmt.Println(r.name, "用时", finishTime.Sub(startTime), "秒")

}
