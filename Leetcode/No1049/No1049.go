package main

import (
	"fmt"
	"sort"
)

func lastStoneWeightII(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}
	n := len(stones)
	sort.Sort(sort.Reverse(sort.IntSlice(stones)))
	for n > 1 {
		index := findNext(stones, 0)
		if index == -1 {
			if n%2 == 0 {
				return 0
			} else {
				return stones[0]
			}
		} else {
			stones = append(stones, stones[0]-stones[index])
			stones = delete(stones, index)
			stones = delete(stones, 0)
			sort.Sort(sort.Reverse(sort.IntSlice(stones)))
			n = len(stones)
		}
	}
	return stones[0]
}

func findNext(stones []int, index int) int {
	for i := index + 1; i < len(stones); i++ {
		if stones[i] != stones[index] {
			return i
		}
	}
	return -1
}

func delete(stones []int, index int) []int {
	stones = append(stones[:index], stones[index+1:]...)
	fmt.Printf("delete: %p\n", stones)
	return stones
}

func delete1(stones []int, index int) {
	stones = append(stones[:index], stones[index+1:]...)
	fmt.Printf("delete: %p\n", stones)
}

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	stones1 := []int{2, 7, 4, 1, 8, 1}
	//stones := []int{6, 8, 4, 3, 8, 3}
	//res := lastStoneWeightII(stones)
	//fmt.Print("res: ",res)
	fmt.Printf("delete: %p\n", stones)
	stones = delete(stones, 4)
	fmt.Println("delete stones: ", stones)
	fmt.Printf("delete: %p\n", stones1)
	delete1(stones1, 4)
	fmt.Println("delete1 stones: ", stones1)
}
