package main

import (
	"fmt"
	"go_code/zinx/znet"
	"io"
	"net"
	"time"
)

func main() {
	fmt.Println("[Client0 starting...]")
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Errorf("Client start err :%v\n", err)
		return
	}
	for {
		dp := znet.NewDataPkg()
		binaryMsg, err := dp.Pack(znet.NewMessage(0, []byte("zinxV1.0 client0 test message")))
		if err != nil {
			fmt.Errorf("client Pack err :%v\n", err)
			return
		}
		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Errorf("client Write err :%v\n", err)
			return
		}
		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Errorf("client msgHead ReadFull err :%v\n", err)
			return
		}
		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Errorf("client msgHead Unpack err :%v\n", err)
			return
		}
		if msgHead.GetMsgLen() > 0 {
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msgHead.GetMsgLen())
			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Errorf("client msg.Data ReadFull err :%v\n", err)
				return
			}
			fmt.Println("zinxV1.0 from server message, MsgId :", msg.Id, "Len :", msg.DataLen, "data :", string(msg.Data))
		}
		time.Sleep(1 * time.Second)
	}
}
