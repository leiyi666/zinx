package znet

import (
	"fmt"
	"go_code/zinx/utils"
	"go_code/zinx/ziface"
	"net"
)

//IServer 接口的实现
type Server struct {
	Name        string
	IPVersion   string
	IP          string
	Port        int
	MsgHandler  ziface.IMsgHandler
	ConnManager ziface.IConnManager
	OnConnStart func(conn ziface.IConnection)
	OnConnStop  func(conn ziface.IConnection)
}

func NewServer(name string) ziface.IServer {
	return &Server{
		Name:        utils.GlobalObject.Name,
		IPVersion:   "tcp4",
		IP:          utils.GlobalObject.Host,
		Port:        utils.GlobalObject.TcpPort,
		MsgHandler:  NewMsgHandler(),
		ConnManager: NewConnManager(),
	}
}

func (s *Server) Start() {
	fmt.Printf("[Start Server Listener at IP :%s, Port :%d], is starting.\n", s.IP, s.Port)
	go func() {
		s.MsgHandler.StartWorkerPool()
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Errorf("Resolve tcp addr err :%v\n", err)
			return
		}
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Errorf("Resolve tcp listen err :%v\n", err)
			return
		}
		fmt.Println("[Start Zinx server succ] ", s.Name)
		var cID uint32
		cID = 0
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Errorf("Resolve tcp accept err :%v\n", err)
				continue
			}

			if s.ConnManager.Len() > utils.GlobalObject.MaxConn {
				fmt.Println("===============>beyond maxConn")
				conn.Close()
				continue
			}

			fmt.Println("当前客户端连接数 :", s.ConnManager.Len())

			dealConn := NewConnection(s, conn, cID, s.MsgHandler)
			cID++
			dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[Stop] server name :", s.Name)
	s.ConnManager.ClearConn()
}

func (s *Server) Serve() {
	s.Start()
	select {}
}

func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgId, router)
}

func (s *Server) GetConnMgr() ziface.IConnManager {
	return s.ConnManager
}

func (s *Server) SetOnConnStart(hookFunc func(conn ziface.IConnection)) {
	s.OnConnStart = hookFunc
}

func (s *Server) SetOnConnStop(hookFunc func(conn ziface.IConnection)) {
	s.OnConnStop = hookFunc
}

func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("CallOnConnStart...")
		s.OnConnStart(conn)
	}
}

func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("OnConnStop...")
		s.OnConnStop(conn)
	}
}
