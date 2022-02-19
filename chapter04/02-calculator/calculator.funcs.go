package main

import "fmt"

func (c *Calculator) Add() int {
	fmt.Printf("函数里%p\n", &c)
	tempResult := c.left + c.right
	c.result = tempResult
	fmt.Println("调用result函数,c的result=", c.result)
	return tempResult
}
func (c Calculator) Sub() int {
	return c.left - c.right
}
func (c Calculator) Multiple() int {
	return c.left * c.right
}
func (c Calculator) Dicide() int {
	return c.left / c.right
}
func (c Calculator) Reminder() int {
	return c.left % c.right
}
