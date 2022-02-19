package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	result := []int{}

	result = append(result, collectPrime(2, 2000000)...)

	//time.Sleep(15 * time.Second)
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
