package main

import (
	"fmt"
	"sort"
)

var L, m int

func main() {
	fmt.Scanln(&L, &m)
	trees := make([][]int, m)
	for i := 0; i < m; i++ {
		trees[i] = make([]int, 2)
		fmt.Scanln(&trees[i][0], &trees[i][1])
	}
	show(trees)
	sort.Slice(trees, func(i, j int) bool {
		return trees[i][0] < trees[j][0]
	})
	show(trees)
	result := 0
	start := 0
	end := 0
	for i := 0; i < m; i++ {
		if start == 0 && end == 0 {
			start = trees[i][0]
			end = trees[i][1]
		} else if trees[i][0] <= end && trees[i][1] >= end {
			end = trees[i][1]
		} else if trees[i][0] > end {
			result += end - start + 1
			start = trees[i][0]
			end = trees[i][1]
		}
	}
	result += end - start + 1
	fmt.Println(L + 1 - result)
}

func show(arr [][]int) {
	for _, v := range arr {
		for _, value := range v {
			fmt.Print(value, "\t")
		}
		fmt.Println()
	}
}
