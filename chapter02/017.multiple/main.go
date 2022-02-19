package main

import "fmt"

func main() {
	//bmis := []float64{1.23, 5.62, 9.66}
	avgbmis := calulate()
	fmt.Println(avgbmis)
}

func calulate(bmis ...float64) interface{} {
	var total float64 = 0
	for _, item := range bmis {
		total += item
	}
	return total / float64(len(bmis))
}
