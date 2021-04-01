package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("[Client starting...]")
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Errorf("Client start err :%v\n", err)
		return
	}
	for {
		_, err := conn.Write([]byte("Hello Zinx V0.1.."))
		if err != nil {
			fmt.Errorf("Client write err %v\n", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Errorf("Client read err %v\n", err)
			return
		}
		fmt.Printf("server call back :%s, cnt :%d\n", buf[:cnt], cnt)
		time.Sleep(1 * time.Second)
	}
}
