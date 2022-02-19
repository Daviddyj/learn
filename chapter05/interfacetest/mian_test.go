package main

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {
	r := TestBox{}
	//var b Box = r
	var c Close = r

	//我们做了一个类型的判断  这个判断的过程就叫做断言
	{
		switch cDetail := c.(type) {
		case Refrigerator:
			fmt.Println("是预期的冰箱 ")
			fmt.Println(cDetail.Size)
		case TestBox:
			fmt.Println("这是一个box,不能当作冰箱用")
		}
	}
	{
		refrigerator, ok := checkIsRefrigerator(c)
		if ok {
			fmt.Println("是个冰箱", refrigerator)
		} else {
			fmt.Println("这不是个冰箱")
		}

	}

}

func checkIsRefrigerator(c Close) (Refrigerator, bool) {
	r, ok := c.(Refrigerator)
	return r, ok
}

type TestBox struct {
}

func (tb TestBox) Close() string {
	return "nil"
}
