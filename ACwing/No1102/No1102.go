package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for ; n > 0; n-- {
		var length int
		var x1, y1 int
		var x2, y2 int
		fmt.Scanf("%d", &length)
		fmt.Scanf("%d %d", &x1, &y1)
		fmt.Scanf("%d %d", &x2, &y2)
		res := bfs(length, x1, y1, x2, y2)
		fmt.Println(res)
	}
}

func bfs(n int, x1 int, y1 int, x2 int, y2 int) int {
	var vis [301][301]int
	dx := [8]int{-2, -1, 1, 2, 2, 1, -1, -2}
	dy := [8]int{1, 2, 2, 1, -1, -2, -2, -1}
	qx := make([]int, 0)
	qy := make([]int, 0)
	qx = append(qx, x1)
	qy = append(qy, y1)
	for len(qx) != 0 {
		nowx := qx[0]
		nowy := qy[0]
		qx = qx[1:]
		qy = qy[1:]
		if nowx == x2 && nowy == y2 {
			break
		} else {
			for i := 0; i < 8; i++ {
				x := nowx + dx[i]
				y := nowy + dy[i]
				if x >= 0 && x < n && y >= 0 && y < n && vis[x][y] == 0 {
					vis[x][y] = vis[nowx][nowy] + 1
					qx = append(qx, x)
					qy = append(qy, y)
				}
			}
		}
	}
	return vis[x2][y2]
}
