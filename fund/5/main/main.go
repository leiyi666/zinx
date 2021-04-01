package main

import (
	"log"
	"time"
)

func main() {
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work…
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { //相当于defer的函数
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}
