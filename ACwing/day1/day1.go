package main

import (
	"fmt"
	"sort"
)

func main() {
	//method 1: 寻找中位数
	fmt.Println("Method 1")
	Median()
	//method 2: 三分法
	fmt.Println("Method 2")
	Trisection()
	fmt.Println("Method 3")
	Median1()
}

func Median() {
	var n int
	fmt.Scanln(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	sort.Ints(arr)
	res := distance(n, arr[n>>1], arr)
	fmt.Println(res)
}

func Trisection() {
	var n int
	fmt.Scanf("%d", &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	sort.Ints(arr)
	l := arr[0]
	r := arr[n-1]
	for l < r-1 {
		midl := (l + r) / 2
		midr := (midl + r) / 2
		if distance(n, midl, arr) > distance(n, midr, arr) {
			l = midl
		} else {
			r = midr
		}
	}
	res1 := distance(n, l, arr)
	res2 := distance(n, r, arr)
	if res1 <= res2 {
		fmt.Print(res1)
	} else {
		fmt.Print(res2)
	}
}

func Median1() {
	var n int
	fmt.Scanln(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	temp := Qchose(0, n-1, n/2, arr)
	res := distance(n, temp, arr)
	fmt.Println(res)
}

func Qchose(l int, r int, k int, arr []int) int {
	pivot := arr[l]
	i, j := l, r
	for i < j {
		for i < j && arr[j] >= pivot {
			j--
		}
		if i < j {
			arr[i] = arr[j]
		}
		for i < j && arr[i] <= pivot {
			i++
		}
		if i < j {
			arr[j] = arr[i]
		}
		if i >= j {
			arr[i] = pivot
		}
	}
	if i == j && i == k {
		return pivot
	} else if k < j {
		return Qchose(l, j-1, k, arr)
	} else {
		return Qchose(j+1, r, k, arr)
	}
}

func distance(n int, num int, arr []int) int {
	res := 0
	for i := 0; i < n; i++ {
		res += Abs(arr[i], num)
	}
	return res
}

func Abs(x int, y int) int {
	z := x - y
	if z >= 0 {
		return z
	} else {
		return -z
	}
}

func show(arr []int) {
	for _, v := range arr {
		fmt.Print(v, "\t")
	}
	fmt.Println()
}
