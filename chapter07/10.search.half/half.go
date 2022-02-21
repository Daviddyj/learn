package main

import (
	"fmt"
	"time"
)

var totalCompare int = 0

func main() {
	// size := 1000
	//arr := sampleData // generateRandomData(size)
	arr := []int{1, 5, 3, 6}
	//fmt.Println(arr)
	// 1. 501 在不在里边？
	// 2. 888 在不在？
	// 3. 900 ？
	// 4. 3 ?
	// ……

	startTime := time.Now()
	//for i := 0; i < 1000000; i++ { // 4.838486846s，次数：31 6500 0000
	fmt.Println(search(arr, 501))
	fmt.Println(search(arr, 3))
	fmt.Println(search(arr, 6))

	//search(arr, 888)
	//search(arr, 900)
	//}

	finishTime := time.Now()

	fmt.Println("总比较次数：", totalCompare)
	fmt.Println("总用时：", finishTime.Sub(startTime))
}

func search(arrP []int, targetNum int) bool {
	quickSort(arrP, 0, len(arrP)-1)
	return half(arrP, 0, len(arrP)-1, targetNum)

}

func half(p []int, left int, right int, targetNum int) bool {
	totalCompare++
	middle := (left + right) / 2
	key := p[middle]

	if key == targetNum {
		return true
	} else {
		if left == right {
			return false
		}
		if targetNum < key {
			half(p, 0, middle, targetNum)
		}
		if targetNum > key {
			half(p, middle+1, right, targetNum)
		}

	}
	return false

}

func searchHarf(arrP []int, left, right int, targetNum int) bool {
	middleIndex := (left + right) / 2
	data := (arrP)[middleIndex]

	totalCompare++ // 增加计数器数值

	if data < targetNum {
		if left == right {
			return false
		}
		return searchHarf(arrP, middleIndex+1, right, targetNum)
	} else if data > targetNum {
		if left == right {
			return false
		}
		return searchHarf(arrP, left, middleIndex-1, targetNum)
	} else {
		return true
	}
}
