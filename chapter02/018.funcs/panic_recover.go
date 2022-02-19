package main

import "fmt"

func panicAndRecover() {
	defer coverPainUpgraded() //这样是能抓住严重错误的     工作中用这个函数
	//defer func() {
	//	coverPainUpgraded() //还是抓不住异常
	//}()
	defer coverPain()
	var nameScore map[string]int = nil
	nameScore["一杰"] = 100
}

func coverPain() { //不在同一个栈内，程序无法recover住
	func() {
		if r := recover(); r != nil {
			fmt.Println("系统出现严重故障") //该问题功能模块因报错不会实现执行下去，但会让系统main函数继续执行下去，非则会panic宝座，主程序直接退出
		}
	}()
}

func coverPainUpgraded() { //这样是能抓住严重错误的

	if r := recover(); r != nil {
		fmt.Println("系统出现严重故障") //该问题功能模块因报错不会实现执行下去，但会让系统main函数继续执行下去，非则会panic宝座，主程序直接退出
	}

}
