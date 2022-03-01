package service

import (
	"customerManager/model"
	"fmt"
) //写customerManager下的相对路径

type CustomerService struct {
	customers   []model.Customer
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	newCustomer := model.NewCustomer(1, "戴一节", "男", 26, "13913271775", "1165207594@qq.com")
	customerService.customerNum = 1
	customerService.customers = append(customerService.customers, newCustomer)
	return customerService
}

func (c *CustomerService) List() []model.Customer {
	return c.customers
}

func (c *CustomerService) Add(customer model.Customer) bool {
	c.customerNum++
	customer.Id = c.customerNum
	c.customers = append(c.customers, customer)
	return true
}

func (c *CustomerService) Delete() {
	fmt.Println("请输入你要删除的用户ID号:")
	var id int
	fmt.Scanln(&id)
	for index, customer := range c.customers {
		if customer.Id == id {
			c.customers = append(c.customers[:index], c.customers[index+1:]...)
		}
	}
}
