package main

import (
	"fmt"
	"sort"
)

var n, m int

func main() {
	fmt.Scanln(&n, &m)
	coins := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &coins[i])
	}
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] < coins[j]
	})
	left := 0
	right := n - 1
	for left < right {
		cur := coins[left] + coins[right]
		if cur == m {
			break
		} else if cur < m {
			left++
		} else {
			right--
		}
	}
	if left < right {
		fmt.Println(coins[left], coins[right])
	} else {
		fmt.Println("No Solution")
	}
}

func show(arr []int) {
	for _, v := range arr {
		fmt.Print(v, "\t")
	}
}
