package main

import (
	"fmt"
)

func main() {
	var index int
	fmt.Scanf("%d", &index)
	Solution(index)
}

func Solution(index int) {
	result := make([]int, 0)
	for i := 1; i <= 300; i++ {
		if judge(transfor(i*i, index)) {
			result = append(result, i)
		}
	}
	show(result, index)
}

func transfor(num int, index int) string {
	table := [21]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"}
	s := ""
	for num/index != 0 {
		s = fmt.Sprint(table[num%index], s)
		num = num / index
	}
	s = fmt.Sprint(table[num%index], s)
	return s
}

func judge(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func show(arr []int, index int) {
	for _, v := range arr {
		fmt.Println(transfor(v, index), transfor(v*v, index))
	}
}
