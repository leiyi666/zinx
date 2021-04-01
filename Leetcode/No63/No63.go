package main

import (
	"fmt"
)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	var dp [][]int
	for i := 0; i < m; i++ {
		d := make([]int, n)
		dp = append(dp, d)
	}
	dp[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				if j-1 >= 0 && obstacleGrid[i][j-1] == 0 {
					dp[i][j] += dp[i][j-1]
				}
				if i-1 >= 0 && obstacleGrid[i-1][j] == 0 {
					dp[i][j] += dp[i-1][j]
				}
			}
		}
	}
	return dp[m-1][n-1]
}

func main() {
	obstacleGrid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	res := uniquePathsWithObstacles(obstacleGrid)
	fmt.Print("res: ", res)
}
