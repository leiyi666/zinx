package main

import (
	"fmt"
)

func uniquePaths(m int, n int) int {
	var res int
	dfs(0, 0, m, n, &res)
	return res
}

func dfs(x int, y int, m int, n int, count *int) {
	if x == m-1 && y == n-1 {
		*count++
	}
	if x+1 < m {
		dfs(x+1, y, m, n, count)
	}
	if y+1 < n {
		dfs(x, y+1, m, n, count)
	}
}

func main() {
	m, n := 23, 12
	res := uniquePaths(m, n)
	fmt.Print("res: ", res)
}
