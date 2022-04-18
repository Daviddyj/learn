package main

import "fmt"

//来自百度的归并排序
func mergeSort(r []int) []int {
	length := len(r)
	if length <= 1 {
		return r
	}
	num := length / 2
	left := mergeSort(r[:num])
	right := mergeSort(r[num:])
	return merge(left, right)
}
func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
		if l == len(left) {
			result = append(result, right[r:]...)
			return result
		}
		if r == len(right) {
			result = append(result, left[l:]...)
			return result
		}
	}
}

func main() {
	arr := []int{45, 12, 42, 33, 10, 44, 0, 27, 27, 20}
	fmt.Println(mergeSort(arr))
}
