package main

import (
	"fmt"
	"testing"
	"time"
)

func TestDefChannel(t *testing.T) {
	var sampleMap map[string]int = map[string]int{}
	var intCh chan int = make(chan int, 10)
	fmt.Println(sampleMap)
	fmt.Println(intCh)

	intCh <- 5

	out := <-intCh
	fmt.Println(intCh)
	fmt.Println(out)
}

func TestChanPutGet(t *testing.T) {
	workerCount := 10
	var intCh = make(chan int)
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			intCh <- i
		}(i)
	}

	for j := 0; j < workerCount; j++ {
		go func(j int) {
			out := <-intCh
			fmt.Printf("j为%d,拿到的数为%d\n", j, out)
		}(j)
	}
	time.Sleep(1 * time.Second)

}
func TestChanPutGet_wait(t *testing.T) {
	workerCount := 5
	var intCh = make(chan int, 5)

	for j := 0; j < workerCount; j++ {
		go func(j int) {
			out := <-intCh
			fmt.Println(j, "拿到out的时间为:", time.Now())
			fmt.Printf("j为%d,拿到的数为%d\n", j, out)
		}(j)
	}
	time.Sleep(5 * time.Second)
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			fmt.Println(i, "工人准备工作:", time.Now())
			intCh <- i
			fmt.Println(i, "工人结束工作:", time.Now())

		}(i)
	}

	time.Sleep(1 * time.Second)

}

//range没有close的话，在去除所有数据后会panic
func TestRangeChannel(t *testing.T) {
	intCh := make(chan int, 10)
	intCh <- 1
	intCh <- 2
	intCh <- 3
	intCh <- 4
	intCh <- 5

	{
		o1, ok := <-intCh
		fmt.Println("直接取数：", o1, ok)
	}

	close(intCh) //这里关水管
	for ch := range intCh {
		fmt.Println("range取数:", ch)
	}

}

//case的优先级大于default    如过case1和case2都要等，则执行default
//close的channel永远是ready状态的

func TestRangeChannelSelect(t *testing.T) {
	intCh := make(chan int, 10)
	intCh2 := make(chan string, 10)
	go func() {
		time.Sleep(2 * time.Second)

		intCh <- 5
	}()
	go func() {
		time.Sleep(2 * time.Second)

		intCh2 <- "haha"

	}()

	time.Sleep(1 * time.Second)
	select {
	case o := <-intCh:
		fmt.Println("int", o)
	case o := <-intCh2:
		fmt.Println("string", o)
	default:
		fmt.Println("哈哈哈")
	}

}
