package main

import "fmt"

func main() {
	x := 9
	word, bit := x/64, uint(x%64)
	y := 1 << bit
	fmt.Println(word, bit, y)
}
