package main

import (
	"fmt"
	"go_code/zinx/ziface"
	"go_code/zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (pr *PingRouter) Handle(req ziface.IRequest) {
	fmt.Println("Call Router Handle")
	fmt.Println("Message from client, MsgId :", req.GetMsgId(), "MsgData :", string(req.GetData()))
	err := req.GetConnection().SendMsg(1, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Errorf("server Handle SendMsg err :%v\n", err)
		return
	}
}

func main() {
	s := znet.NewServer("[zinx V0.5]")
	s.AddRouter(&PingRouter{})
	s.Serve()
}