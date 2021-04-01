package main

import (
	"fmt"
)

var start, aim []byte

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	start = []byte(s1)
	aim = []byte(s2)
	result := 0
	for i := 0; i < len(aim)-1; i++ {
		if start[i] != aim[i] {
			turn(i)
			turn(i + 1)
			result++
		}
	}
	fmt.Println(result)
}

func turn(i int) {
	if start[i] == '*' {
		start[i] = 'o'
	} else {
		start[i] = '*'
	}
}
