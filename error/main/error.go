package main

import (
	"errors"
	"fmt"
)

func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误")
	}
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err = ", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res = ", res)
}

func test1() {
	err := readConf("config.ini")
	if err != nil {
		panic(err)
	}

	fmt.Println("test1()继续执行")
}

func main() {
	test()
	test1()
	fmt.Println("main()下面的代码和逻辑")
}
