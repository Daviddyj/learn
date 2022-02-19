package main

import "fmt"

func main() {
	var left, right int = 1, 2
	//var op string = "+"

	c := Calculator{
		left:  left,
		right: right,
		//op:    op,
	}
	fmt.Printf("%p\n", &c)
	fmt.Println(c.Add())
	fmt.Println(c.result)

}
