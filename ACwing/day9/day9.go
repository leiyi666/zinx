package main

import (
	"fmt"
	"sort"
)

var n int

type Student struct {
	Id      int
	Chinese int
	Math    int
	English int
	sum     int
}

func main() {
	fmt.Scanln(&n)
	students := make([]Student, n)
	for i := 0; i < n; i++ {
		students[i].Id = i + 1
		fmt.Scanln(&students[i].Chinese, &students[i].Math, &students[i].English)
		students[i].sum = students[i].Chinese + students[i].Math + students[i].English
	}
	sort.Slice(students, func(i, j int) bool {
		if students[i].sum != students[j].sum {
			return students[i].sum > students[j].sum
		} else if students[i].Chinese != students[j].Chinese {
			return students[i].Chinese > students[j].Chinese
		} else {
			return students[i].Id < students[j].Id
		}
	})
	for i := 0; i < 5; i++ {
		fmt.Println(students[i].Id, students[i].sum)
	}
}

func show(arr []Student) {
	for _, v := range arr {
		fmt.Println(v)
	}
}
