package main

import "fmt"

func main() {
	Solution()
}

func Solution() {
	var m int
	var n int
	fmt.Scanf("%d %d", &m, &n)
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}
	dx := [4]int{-1, 0, 1, 0}
	dy := [4]int{0, 1, 0, -1}
	x, y, d := 0, 0, 1
	for i := 1; i <= n*m; i++ {
		result[x][y] = i
		a := x + dx[d]
		b := y + dy[d]
		if a < 0 || a >= m || b < 0 || b >= n || result[a][b] != 0 {
			d = (d + 1) % 4
			a = x + dx[d]
			b = y + dy[d]
		}
		x, y = a, b
	}
	show(result)
}

func show(arr [][]int) {
	for _, v := range arr {
		for _, value := range v {
			fmt.Print(value, " ")
		}
		fmt.Println()
	}
}
