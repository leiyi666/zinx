package main

import (
	"fmt"
	"math"
)

func main() {
	n := 0
	fmt.Scanln(&n)
	w := make([][]int, n)
	for i := 0; i < n; i++ {
		w[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Scanln(&w[i][j])
		}
	}
	res := Solution(n, w)
	fmt.Println(res)
}

func Solution(n int, w [][]int) int {
	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			w[i][j] = int(math.Max(float64(w[i+1][j]), float64(w[i+1][j+1]))) + w[i][j]
		}
	}
	return w[0][0]
}

func show(arr [][]int) {
	for _, v := range arr {
		for _, value := range v {
			fmt.Print(value, "\t")
		}
		fmt.Println()
	}
}
