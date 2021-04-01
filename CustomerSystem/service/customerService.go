package service

import (
	"go_code/CustomerSystem/model"
)

type CustomerService struct {
	customers   []model.Customer
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "Tom", "male", 18, "110", "Tom@qq.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

func (cs *CustomerService) Add(customer model.Customer) bool {
	cs.customerNum++
	customer.Id = cs.customerNum
	cs.customers = append(cs.customers, customer)
	return true
}

func (cs *CustomerService) FindById(id int) int {
	for i := 0; i < len(cs.customers); i++ {
		if cs.customers[i].Id == id {
			return i
		}
	}
	return -1
}

func (cs *CustomerService) Update(id int, customer model.Customer) bool {
	index := cs.FindById(id)
	if index == -1 {
		return false
	}
	customer.Id = id
	cs.customers[id-1] = customer
	return true
}

func (cs *CustomerService) Delete(id int) bool {
	index := cs.FindById(id)
	if index == -1 {
		return false
	}
	cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
	return true
}

func (cs *CustomerService) List() []model.Customer {
	return cs.customers
}
