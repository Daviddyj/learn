package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	result := make(chan int, 2000000)
	workNumber := 16
	baseNumCh := make(chan int, 1024)
	wg := sync.WaitGroup{}
	wg.Add(workNumber)
	for i := 0; i < workNumber; i++ {
		go func() {
			defer wg.Done()
			for numCh := range baseNumCh {
				if isPrime(numCh) {
					//result = append(result, numCh) //多个goroutine写入数据到result,有风险
					result <- numCh
				}
			}
		}()
	}

	for i := 2; i < 2000000; i++ {
		baseNumCh <- i
	}
	close(baseNumCh)
	//time.Sleep(15 * time.Second)
	finishTime := time.Now()
	//time.Sleep(3 * time.Second)
	wg.Wait()
	fmt.Println(len(result))
	fmt.Println("共耗时：", finishTime.Sub(startTime))

}

func collectPrime(start, end int) []int {
	var result []int
	for num := start; num <= end; num++ {
		if isPrime(num) {
			result = append(result, num)
		}
	}
	return result
}

func isPrime(num int) (isPrime bool) {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			return
		}
	}
	isPrime = true
	return
}
