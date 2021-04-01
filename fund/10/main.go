/*
	模拟火箭发射程序
*/
package main

import (
	"fmt"
	"os"
	"time"
)

//倒计时期间不能中断
/*func main() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	launch()
}*/

//倒计时期间可以中断
/*func main() {
	fmt.Println("Commencing countdown. Press return to abort.")
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	select {
		case <- time.After(10 * time.Second):

		case <- abort :
				fmt.Println("Launch aborted!")
				return
	}
	launch()
}*/

func main() {
	// ...create abort channel...
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	fmt.Println("Commencing countdown.  Press return to abort.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// Do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}

func launch() {
	fmt.Println("Go!")
}
