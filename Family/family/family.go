package family

import "fmt"

type family struct {
	flag    bool
	key     string
	balance float64
	money   float64
	note    string
	details string
}

func NewFamily() *family {
	return &family{
		flag:    true,
		key:     "",
		balance: 0.0,
		money:   0.0,
		note:    "",
		details: "收支\t余额\t金额\t说明",
	}
}

func (this *family) MainMenu() {
	for this.flag == true {
		fmt.Println("-----------------------家庭收支管理系统-----------------------")
		fmt.Println("-----------------------功能1： 查看明细-----------------------")
		fmt.Println("-----------------------功能2： 存入钱钱-----------------------")
		fmt.Println("-----------------------功能3： 取出钱钱-----------------------")
		fmt.Println("-----------------------功能4： 退出系统-----------------------")
		fmt.Println("选择功能1-4：")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.show()
		case "2":
			this.income()
		case "3":
			this.withdraw()
		case "4":
			this.quit()
		default:
			fmt.Println("输入无效！请重新输入: ")
		}
	}
}

func (this *family) show() {
	if this.note != "" {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有任何操作记录，快去存笔钱钱吧！")
	}
	fmt.Println()
}

func (this *family) income() {
	fmt.Print("请输入存钱的金额: ")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Print("请输入存钱的理由: ")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
}

func (this *family) withdraw() {
	fmt.Print("请输入取钱的金额: ")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("你哪有这么多钱啊，想屁吃呢！")
		return
	}
	this.balance -= this.money
	fmt.Print("请输入取钱的理由: ")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.note)
}

func (this *family) quit() {
	fmt.Println("请输入： y/n")
	c := ""
	fmt.Scanln(&c)
	for c != "y" && c != "n" {
		fmt.Println("请输入： y/n")
		fmt.Scanln(&c)
	}
	if c == "y" {
		this.flag = false
	}
}
