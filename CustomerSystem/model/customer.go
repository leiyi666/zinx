package model

import "fmt"

type Customer struct {
	Id     int
	name   string
	gender string
	age    int
	phone  string
	email  string
}

func NewCustomer(id int, name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		Id:     id,
		name:   name,
		gender: gender,
		age:    age,
		phone:  phone,
		email:  email,
	}
}

func NewCustomer1(name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		name:   name,
		gender: gender,
		age:    age,
		phone:  phone,
		email:  email,
	}
}

func (c *Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v", c.Id, c.name, c.gender, c.age, c.phone, c.email)
	return info
}
