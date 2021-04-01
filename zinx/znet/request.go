package znet

import (
	"go_code/zinx/ziface"
)

type Request struct {
	Conn ziface.IConnection
	Msg  ziface.IMessage
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.Conn
}

func (r *Request) GetMsgId() uint32 {
	return r.Msg.GetMsgId()
}

func (r *Request) GetData() []byte {
	return r.Msg.GetData()
}
