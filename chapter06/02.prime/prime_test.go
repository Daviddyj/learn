package _2_prime

import (
	"fmt"
	"testing"
	"time"
)

//map 很危险
func TestRunPrime(t *testing.T) {
	startTime := time.Now()
	result := []int{}
	for num := 2; num <= 200000; num++ {
		if isPrime(num) {
			result = append(result, num)
		}
	}
	finishTime := time.Now()
	fmt.Println(len(result))
	fmt.Println("共耗时：", finishTime.Sub(startTime))
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
func TestRunPrime3(t *testing.T) {
	startTime := time.Now()
	result := []int{}
	go func() {
		fmt.Println("第一个work开始计算", time.Now())
		result = append(result, collectPrime(2, 100000)...)
		fmt.Println("第一个work完成计算", time.Now())

	}()
	go func() {
		fmt.Println("第二个work开始计算", time.Now())
		result = append(result, collectPrime(100001, 200000)...)
		fmt.Println("第二个work完成计算", time.Now())
	}()
	time.Sleep(15 * time.Second)
	finishTime := time.Now()
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
