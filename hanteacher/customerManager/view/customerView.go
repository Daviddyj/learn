package main

import (
	"customerManager/service"
	"fmt"
)

func main() {
	NewCustomerView().ForMain()
}

type customerView struct {
	key      string
	loop     bool
	customer service.CustomerService
}

func NewCustomerView() *customerView {
	return &customerView{
		key:  "",
		loop: true,
	}
}

func (cv *customerView) ForMain() {
	for {
		fmt.Println("-----------------客户信息管理软件-----------------")
		fmt.Println(" 1 添 加 客 户")
		fmt.Println(" 2 修 改 客 户")
		fmt.Println(" 3 删 除 客 户")
		fmt.Println(" 4 客 户 列 表")
		fmt.Println(" 5 退 出")
		fmt.Print("请选择(1-5)：")
		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			fmt.Println("添 加 客 户")
		case "2":
			fmt.Println("修 改 客 户")
		case "3":
			fmt.Println("删 除 客 户")
		case "4":
			fmt.Println("客 户 列 表")
			cv.List()
		case "5":
			cv.loop = false
		default:
			fmt.Println("输入有误，请重新输入...")

		}
		if cv.loop == false {
			break
		}
	}
	fmt.Println("退出客户关系管理系统")

}

//显示所有的客户信息

func (c *customerView) List() {
	customers := service.List()

	fmt.Println("---------------------------客户列表---------------------------")
	fmt.Println("编号\t 姓名\t 性别\t 年龄\t 电话\t 邮箱")
	for i := 0; i < len(customers); i++ {
		//fmt.Println(customers[i].Id,"\t", customers[i].Name...)
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("\n-------------------------客户列表完成-------------------------\n\n")

}
