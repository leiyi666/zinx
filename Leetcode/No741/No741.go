package No741

import (
	"fmt"
	"math"
)

/*
不会
*/
func cherryPickup(grid [][]int) int {
	n := len(grid)
	w := make([][]int, n)
	for i := 0; i < n; i++ {
		w[i] = make([]int, n)
	}
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != -1 {
				if (i-1 >= 0 && grid[i-1][j] != -1) && (j-1 >= 0 && grid[i][j-1] != -1) {
					w[i][j] = int(math.Max(float64(w[i-1][j]), float64(w[i][j-1])))
				} else if (i-1 >= 0 && grid[i-1][j] != -1) && (j-1 < 0 || grid[i][j-1] == -1) {
					w[i][j] = w[i-1][j]
				} else if (i-1 < 0 || grid[i-1][j] == -1) && (j-1 >= 0 && grid[i][j-1] != -1) {
					w[i][j] = w[i][j-1]
				}
				w[i][j] += grid[i][j]
				grid[i][j] = 0
			}
		}
	}
	show(grid)
	show(w)
	fmt.Println()
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] != -1 {
				if (i+1 < n && grid[i+1][j] != -1) && (j+1 < n && grid[i][j+1] != -1) {
					w[i][j] = int(math.Max(float64(w[i+1][j]), float64(w[i][j+1])))
				} else if (i+1 < n && grid[i+1][j] != -1) && (j+1 >= n || grid[i][j+1] == -1) {
					w[i][j] = w[i+1][j]
				} else if (i+1 >= n || grid[i+1][j] == -1) && (j+1 < n && grid[i][j+1] != -1) {
					w[i][j] = w[i][j+1]
				}
				w[i][j] += grid[i][j]
				grid[i][j] = 0
			}
		}
	}
	show(grid)
	show(w)
	return w[0][0] + count
}

func show(arr [][]int) {
	for _, v := range arr {
		for _, value := range v {
			fmt.Print(value, "\t")
		}
		fmt.Println()
	}
}
