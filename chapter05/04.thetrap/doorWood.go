package main

import "fmt"

var _ Door = &WoodDoor{} //以后想要实现这个Door接口先实现这个匿名变量，来确保接口的成功实现，如过对象的成员函数没有实现接口定义的方法，则程序会编译报错

type WoodDoor struct {
}

func (d *WoodDoor) Unlock() {
	fmt.Println("WoodDoor Unlock ")

}

func (d *WoodDoor) Lock() {
	fmt.Println("WoodDoor Lock ")

}

func (*WoodDoor) Open() {
	fmt.Println("WoodDoor Open ")
}
func (*WoodDoor) Close() {
	fmt.Println("WoodDoor Close")
}
