package main

import "fmt"

var ch8 = make(chan int, 6)

/*func mm1() {
	for i := 0; i < 10; i++ {
		ch8 <- 8 * i
	}

}
func main() {
	fmt.Println("ch8 begin :", len(ch8))
	go mm1()
	for i:=0;i<10;i++{
		fmt.Print(<-ch8, "\t")
	}
	fmt.Println()
	fmt.Println("ch8 end :", len(ch8))
}*/

func mm1() {
	for i := 0; i < 10; i++ {
		ch8 <- 8 * i
	}
	close(ch8)
}

func main() {
	fmt.Println("ch8 begin :", len(ch8))
	go mm1()
	for {
		for data := range ch8 {
			fmt.Print(data, "\t")
		}
		break
	}
	fmt.Println("ch8 end :", len(ch8))
}
