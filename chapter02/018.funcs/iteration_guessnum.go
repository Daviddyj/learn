package main

import "fmt"

func guess(left, right uint) {
	guessd := (left + right) / 2
	var getFromInput int
	fmt.Println("我猜的是", guessd)
	fmt.Println("如果高了，输入1，如果低了，输入0；对了，输入9")
	fmt.Println("请输入你要输入的数字:")
	fmt.Scanln(&getFromInput)
	switch getFromInput {
	case 0:
		if left == right {
			fmt.Println("你是不是该数字了？")
			return
		}

		guess(guessd+1, right)
	case 1:
		if left == right {
			fmt.Println("你是不是该数字了？")
			return

		}
		guess(left, guessd-1)
	case 9:
		fmt.Println("你心里想的数字是：", guessd)
	}

}
