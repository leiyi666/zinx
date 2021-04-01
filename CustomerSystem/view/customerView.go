package view

import (
	"fmt"
	"go_code/CustomerSystem/model"
	"go_code/CustomerSystem/service"
)

type CustomerView struct {
	flag            bool
	key             string
	customerService *service.CustomerService
}

func NewCustomerView() CustomerView {
	cs := service.NewCustomerService()
	return CustomerView{
		flag:            true,
		key:             "",
		customerService: cs,
	}
}

func (cv *CustomerView) Menu() {
	for cv.flag == true {
		fmt.Println("-----------------------客户信息管理软件-----------------------")
		fmt.Println("-----------------------功能1： 添加客户-----------------------")
		fmt.Println("-----------------------功能2： 修改客户-----------------------")
		fmt.Println("-----------------------功能3： 删除客户-----------------------")
		fmt.Println("-----------------------功能4： 客户列表-----------------------")
		fmt.Println("-----------------------功能5： 退出软件-----------------------")
		fmt.Println("选择功能1-5：")

		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			cv.add()
		case "2":
			cv.update()
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.quit()
		default:
			fmt.Println("输入无效！请重新输入: ")
		}
	}
}

func (cv *CustomerView) add() {
	fmt.Println("--------------------------添加客户--------------------------")
	fmt.Println("姓名")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer1(name, gender, age, phone, email)
	if cv.customerService.Add(customer) {
		fmt.Println("--------------------------添加完成--------------------------")
	} else {
		fmt.Println("--------------------------添加失败--------------------------")
	}
}

func (cv *CustomerView) update() {
	fmt.Println("--------------------------修改客户--------------------------")
	fmt.Println("请选择待修改客户编号(-1退出): ")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("姓名")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer1(name, gender, age, phone, email)
	if cv.customerService.Update(id, customer) {
		fmt.Println("--------------------------修改完成--------------------------")
	} else {
		fmt.Println("--------------------------修改失败--------------------------")
	}
}

func (cv *CustomerView) delete() {
	fmt.Println("--------------------------删除客户--------------------------")
	fmt.Println("请选择待删除客户编号(-1退出): ")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确认是否删除(Y/N): ")
	c := ""
	fmt.Scanln(&c)
	for c != "y" && c != "Y" && c != "n" && c != "N" {
		fmt.Println("确认是否删除(Y/N): ")
		fmt.Scanln(&c)
	}
	if c == "N" || c == "n" {
		return
	}
	if cv.customerService.Delete(id) {
		fmt.Println("--------------------------删除完成--------------------------")
	} else {
		fmt.Println("--------------------------删除失败--------------------------")
	}
}

func (cv *CustomerView) list() {
	customers := cv.customerService.List()
	fmt.Println("--------------------------客户列表--------------------------")
	fmt.Println("编号\t\t姓名\t\t性别\t\t年龄\t\t电话\t\t邮箱")
	for _, c := range customers {
		fmt.Println(c.GetInfo())
	}
	fmt.Println("------------------------客户列表完成-------------------------")
}

func (cv *CustomerView) quit() {
	fmt.Println("是否确定退出程序Y/N")
	c := ""
	fmt.Scanln(&c)
	for c != "y" && c != "Y" && c != "n" && c != "N" {
		fmt.Println("是否确定退出程序Y/N")
		fmt.Scanln(&c)
	}
	if c == "y" || c == "y" {
		cv.flag = false
	}
}
