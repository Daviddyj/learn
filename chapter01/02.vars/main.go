package main

import (
	"fmt"
	"math"
)

func main() {
	var hello string = "hello,world"
	var world = "world"
	fmt.Println(hello, world)
	小鼠 := 123
	fmt.Println(小鼠)
	var int3, int4 int8 = 33, 44
	fmt.Println(int4, int3)
	var int5 uint = math.MaxInt64
	fmt.Println(int5)
	var int6 int = int(int5)

	fmt.Println(int6)

}
