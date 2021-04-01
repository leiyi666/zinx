package main

import (
	"fmt"
	"go_code/zinx/ziface"
	"go_code/zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (pr *PingRouter) PreHandle(res ziface.IRequest) {
	fmt.Println("Call Router PreHandle")
	_, err := res.GetConnection().GetTCPConnection().Write([]byte(" Before Ping...\n"))
	if err != nil {
		fmt.Errorf("PreHandle call back err :%v\n", err)
	}
}
func (pr *PingRouter) Handle(res ziface.IRequest) {
	fmt.Println("Call Router Handle")
	_, err := res.GetConnection().GetTCPConnection().Write([]byte(" Ping... Ping...\n"))
	if err != nil {
		fmt.Errorf("Handle call back err :%v\n", err)
	}
}
func (pr *PingRouter) PostHandle(res ziface.IRequest) {
	fmt.Println("Call Router PostHandle")
	_, err := res.GetConnection().GetTCPConnection().Write([]byte(" After Ping...\n"))
	if err != nil {
		fmt.Errorf("PostHandle call back err :%v\n", err)
	}
}

func main() {
	s := znet.NewServer("[zinx V0.3]")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
