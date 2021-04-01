package main

import (
	"fmt"
	"math"
)

var n, m int
var res []int

func main() {
	fmt.Scanln(&n, &m)
	res = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &res[i])
	}
	l := 0.0
	r := 1e9
	for r-l >= 1e-4 {
		mid := (l + r) / 2
		if check(mid) {
			l = mid
		} else {
			r = mid
		}
	}
	fmt.Printf("%.2f\n", l)
}

func check(x float64) bool {
	cnt := 0
	for i := 0; i < n; i++ {
		cnt += int(math.Floor(float64(res[i]) / x))
	}
	if cnt >= m {
		return true
	} else {
		return false
	}
}

func show(arr []int) {
	for _, v := range arr {
		fmt.Print(v, "\t")
	}
}
