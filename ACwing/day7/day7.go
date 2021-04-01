package main

import (
	"fmt"
)

var n, k int

var H [100010]int
var W [100010]int

func main() {
	fmt.Scanln(&n, &k)
	for i := 0; i < n; i++ {
		fmt.Scanln(&H[i], &W[i])
	}
	l, r := 1, 100000
	for l < r {
		mid := (l + r + 1) >> 1
		if check(mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	fmt.Println(l)
}

func check(x int) bool {
	var cnt int
	for i := 0; i < n; i++ {
		a := H[i] / x
		b := W[i] / x
		cnt += a * b
	}
	if cnt >= k {
		return true
	} else {
		return false
	}
}
