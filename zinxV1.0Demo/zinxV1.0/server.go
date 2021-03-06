package main

import (
	"fmt"
	"go_code/zinx/ziface"
	"go_code/zinx/znet"
)

func DoConnBegin(conn ziface.IConnection) {
	fmt.Println("======>DoConnBegin is called")
	if err := conn.SendMsg(202, []byte("DoConnBegin begin")); err != nil {
		fmt.Errorf("DoConnBegin err :%v\n", err)
	}

	fmt.Println("Set Property...")
	conn.SetProperty("Name", "The Bug")
	conn.SetProperty("GitHub", "https://github.com/leiyi666")
}

func DoConnLast(conn ziface.IConnection) {
	fmt.Println("======>DoConnLast is called")
	fmt.Println("connID :", conn.GetConnID(), " is stop")
	name, err := conn.GetProperty("Name")
	if err != nil {
		fmt.Println("GetProperty err :", err)
	}
	fmt.Println("Name :", name)
	GitHub, err := conn.GetProperty("GitHub")
	if err != nil {
		fmt.Println("GetProperty err :", err)
	}
	fmt.Println("GitHub :", GitHub)
}

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

type HelloRouter struct {
	znet.BaseRouter
}

func (hr *HelloRouter) Handle(req ziface.IRequest) {
	fmt.Println("Call Router Handle")
	fmt.Println("Message from client, MsgId :", req.GetMsgId(), "MsgData :", string(req.GetData()))
	err := req.GetConnection().SendMsg(1, []byte("hello...hello...hello"))
	if err != nil {
		fmt.Errorf("server Handle SendMsg err :%v\n", err)
		return
	}
}

func main() {
	s := znet.NewServer("[zinx V1.0]")
	s.SetOnConnStart(DoConnBegin)
	s.SetOnConnStop(DoConnLast)
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})
	s.Serve()
}
