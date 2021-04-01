package main

import (
	"fmt"
	"math"
)

func minCostClimbingStairs(cost []int) int {
	var dp []int = make([]int, len(cost)+1)
	for i := 2; i < len(cost); i++ {
		dp[i] = int(math.Min(float64(dp[i-2]+cost[i-2]), float64(dp[i-1]+cost[i-1])))
	}
	return dp[len(cost)-1]
}

func main() {
	var cost []int
	var result int
	cost = []int{0, 2, 2, 1}
	result = minCostClimbingStairs(cost)
	fmt.Print("result ", result)
}
