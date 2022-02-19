package main

import "fmt"

var counter int

func calSum(nums ...int) (sum int) {
	for _, item := range nums {
		sum += item

	}
	counter++

	return
}
func showUsedTime() {
	fmt.Println(counter)
}

func genImprovementFunc() func(percentage float64) {
	base := 1000.0
	return func(percentage float64) {
		base = base * (1 + percentage)
		fmt.Println("new:", base)
	}
}

func closeMain() {
	imp := genImprovementFunc()
	imp(0.1)
	imp(0.1)
	imp(0.1)
	imp(0.1)
	imp(0.1)

}
