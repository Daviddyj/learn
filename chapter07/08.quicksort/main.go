package main

import "fmt"

func quickSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		if start < j {
			quickSort(arr, start, j)
		}
		if end > i {
			quickSort(arr, i, end)
		}
	}
}

func main() {
	arr := []int{45, 12, 42, 33, 10, 44, 0, 27, 27, 20}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
