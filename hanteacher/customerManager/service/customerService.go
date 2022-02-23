package service

import "customerManager/model" //写customerManager下的相对路径

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

func List() []model.Customer {
	return NewCustomerService().customers
}
