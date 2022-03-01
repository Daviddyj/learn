package main

import (
	"customerManager/model"
	"customerManager/service"
	"fmt"
)

func main() {
	NewCustomerView().ForMain()
}

type customerView struct {
	key      string
	loop     bool
	customer *service.CustomerService
}

func NewCustomerView() *customerView {
	return &customerView{
		key:      "",
		loop:     true,
		customer: service.NewCustomerService(),
	}
}

func (cv *customerView) ForMain() {
	for {
		fmt.Println("-----------------客户信息管理软件-----------------")
		fmt.Println(" 1 添 加 客 户:")
		fmt.Println(" 2 修 改 客 户:")
		fmt.Println(" 3 删 除 客 户:")
		fmt.Println(" 4 客 户 列 表:")
		fmt.Println(" 5 退 出")

		fmt.Print("请选择(1-5)：")
		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			cv.add()
		case "2":
			fmt.Println("修 改 客 户:")
		case "3":
			cv.Delete()
		case "4":
			cv.List()
		case "5":
			cv.Exit()
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
	customers := c.customer.List()

	fmt.Println("---------------------------客户列表---------------------------")
	fmt.Println("编号\t 姓名\t 性别\t 年龄\t 电话\t 邮箱")
	for i := 0; i < len(customers); i++ {
		//fmt.Println(customers[i].Id,"\t", customers[i].Name...)
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("\n-------------------------客户列表完成-------------------------\n\n")

}

func (cv *customerView) add() {

	Name := ""
	Gender := ""
	Age := 0
	Phone := ""
	Email := ""

	fmt.Print("姓名:")
	fmt.Scanln(&Name)
	fmt.Print("性别:")
	fmt.Scanln(&Gender)
	fmt.Print("年龄:")
	fmt.Scanln(&Age)
	fmt.Print("手机号码:")
	fmt.Scanln(&Phone)
	fmt.Print("邮箱:")
	fmt.Scanln(&Email)

	newpeople := model.NewCustomer2(Name, Gender, Age, Phone, Email)
	if cv.customer.Add(newpeople) {
		fmt.Println("---添加成功---")
	} else {
		fmt.Println("---添加失败---")
	}
}

func (c *customerView) Delete() {
	c.customer.Delete()
}

func (cv *customerView) Exit() {
	cv.loop = false
}
