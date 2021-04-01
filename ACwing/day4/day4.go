package main

import (
	"fmt"
)

var w, h int

func main() {
	startx := 0
	starty := 0
	for {
		fmt.Scanln(&w, &h)
		if w == 0 && h == 0 {
			break
		}
		house := make([][]byte, h)
		for i := 0; i < h; i++ {
			house[i] = make([]byte, w)
			fmt.Scanln(&house[i])
			for j := 0; j < w; j++ {
				if house[i][j] == '@' {
					startx = i
					starty = j
				}
			}
		}
		result := bfs(startx, starty, house)
		fmt.Println(result)
	}
}

func bfs(i, j int, house [][]byte) int {
	cnt := 0
	q1 := make([]int, 0)
	q2 := make([]int, 0)
	q1 = append(q1, i)
	q2 = append(q2, j)
	dx := [4]int{-1, 0, 1, 0}
	dy := [4]int{0, 1, 0, -1}
	for len(q1) != 0 {
		x := q1[0]
		y := q2[0]
		q1 = q1[1:]
		q2 = q2[1:]
		cnt++
		for k := 0; k < 4; k++ {
			a := x + dx[k]
			b := y + dy[k]
			if a >= 0 && a < h && b >= 0 && b < w && house[a][b] == '.' {
				q1 = append(q1, a)
				q2 = append(q2, b)
				house[a][b] = '#'
			}
		}
	}
	return cnt
}

func show(arr [][]byte) {
	for _, v := range arr {
		for _, value := range v {
			fmt.Printf("%c\t", value)
		}
		fmt.Println()
	}
}
