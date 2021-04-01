package ziface

type IServer interface {
	//start server
	Start()
	//stop server
	Stop()
	//run server
	Serve()

	AddRouter(msgId uint32, router IRouter)

	GetConnMgr() IConnManager

	SetOnConnStart(hookFunc func(conn IConnection))
	SetOnConnStop(hookFunc func(conn IConnection))

	CallOnConnStart(conn IConnection)
	CallOnConnStop(conn IConnection)
}
