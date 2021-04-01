package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	i, j := 1, 2
	fmt.Println(i, j)
	i, j = j, i
	fmt.Println(i, j)
	s := "hello"
	h := s[0:1]
	fmt.Printf("%T", h)
	s1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	s2 := s1[6:9]
	fmt.Println(s2)
	s3 := s2[:5]
	fmt.Println(s3)
}
