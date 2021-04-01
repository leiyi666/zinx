package main

import "go_code/zinx/znet"

func main() {
	s := znet.NewServer("[zinx V0.1]")
	s.Serve()
}
