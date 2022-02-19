package main

import "fmt"

type InputFromStd struct {
}

func (*InputFromStd) GetInput() Person {
	var name string
	fmt.Print("姓名：")
	fmt.Scanln(&name)

	var weight float64
	fmt.Print("体重（千克）：")
	fmt.Scanln(&weight)

	var tall float64
	fmt.Print("身高（米）：")
	fmt.Scanln(&tall)

	var age int
	fmt.Print("年龄：")
	fmt.Scanln(&age)

	//var sexWeight int
	var sex = "男"
	fmt.Print("性别（男/女）：")
	fmt.Scanln(&sex)

	//if sex == "男" {
	//	sexWeight = 1
	//} else {
	//	sexWeight = 0
	//}
	return Person{
		name:   name,
		sex:    sex,
		tall:   tall,
		weight: weight,
		age:    age,
	}
}
