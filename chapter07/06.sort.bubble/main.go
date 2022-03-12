package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	length := 10
	arry := []int{}
	rand.Seed(time.Now().Unix())
	for i := 0; i < length; i++ {
		arry = append(arry, rand.Int()%50)
	}
	fmt.Println(arry)
	bubbleSort(arry)
}

func bubbleSort(arry []int) {
	for i := 0; i < len(arry); i++ {
		for j := 0; j < len(arry)-i-1; j++ {
			if arry[j] > arry[j+1] {
				arry[j], arry[j+1] = arry[j+1], arry[j]
			}
		}
		//fmt.Println("中间排序:", arry)
	}

	fmt.Println("最终排序:", arry)
}
