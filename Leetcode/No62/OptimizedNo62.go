package main

import (
	"fmt"
)

func uniquePaths1(m int, n int) int {
	var dp [][]int
	for i := 0; i < m; i++ {
		d := make([]int, n)
		dp = append(dp, d)
	}
	dp[0][0] = 1
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	return dp[m-1][n-1]
}

func main() {
	m, n := 3, 2
	res := uniquePaths1(m, n)
	fmt.Print("res: ", res)
}
