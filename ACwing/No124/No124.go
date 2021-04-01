package main

import (
	"fmt"
)

func main() {
	//fmt.Println(string(6 + '0'))
	var n int
	fmt.Scanln(&n)
	for ; n > 0; n-- {
		var a, b int32
		var aLine, bLine string
		fmt.Scanln(&a, &b, &aLine)
		temp := make([]int32, 0)
		res := make([]int32, 0)
		for _, aa := range aLine {
			if aa >= '0' && aa <= '9' {
				temp = append(temp, aa-'0')
			} else if aa >= 'A' && aa <= 'Z' {
				temp = append(temp, aa-'A'+10)
			} else if aa >= 'a' && aa <= 'z' {
				temp = append(temp, aa-'a'+36)
			}
		}
		for len(temp) > 0 {
			r := int32(0)
			for i := 0; i < len(temp); i++ {
				temp[i] += r * a
				r = temp[i] % b
				temp[i] = temp[i] / b
			}
			res = append(res, r)
			for len(temp) > 0 && temp[0] == 0 {
				temp = temp[1:]
			}
		}
		for i := len(res) - 1; i >= 0; i-- {
			if res[i] >= 0 && res[i] <= 9 {
				bLine += string(res[i] + '0')
			} else if res[i] >= 10 && res[i] <= 35 {
				bLine += string(res[i] - 10 + 'A')
			} else if res[i] >= 36 && res[i] <= 61 {
				bLine += string(res[i] - 36 + 'a')
			}
		}
		fmt.Println(a, aLine)
		fmt.Println(b, bLine)
		fmt.Println()
	}
}
func show(arr []int32) {
	for _, v := range arr {
		fmt.Print(v, "\t")
	}
	fmt.Println()
}
