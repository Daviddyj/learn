package main

import (
	"fmt"
)

func main() {
	panicAndRecover()
	//deferGuess()
	//openFile()
	////closeMain()
	//time.Sleep(108 * time.Second)
	//fmt.Println(fib(5))

	//guess(1, 100)
	fmt.Println(calSum(10, 45, 45, 4, 1, 5, 4, 5, 45))

	fmt.Println(calSum(10, 45, 45, 4, 1, 5, 4, 5, 45))
	fmt.Println(calSum(10, 45, 45, 4, 1, 5, 4, 5, 45))
	fmt.Println(calSum(10, 45, 45, 4, 1, 5, 4, 5, 45))
	fmt.Println(calSum(10, 45, 45, 4, 1, 5, 4, 5, 45))
	showUsedTime()
}

func fib(n uint) uint {
	if n == 0 {
		return 00
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)

}
